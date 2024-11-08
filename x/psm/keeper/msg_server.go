package keeper

import (
	"context"
	"fmt"
	"strings"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/onomyprotocol/reserve/x/psm/types"
)

type msgServer struct {
	keeper Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) UpdateParams(ctx context.Context, req *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if _, err := k.keeper.addressCodec.StringToBytes(req.Authority); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	if k.keeper.GetAuthority() != req.Authority {
		return nil, errorsmod.Wrapf(types.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.keeper.GetAuthority(), req.Authority)
	}

	if err := req.Params.Validate(); err != nil {
		return nil, err
	}

	if err := k.keeper.SetParams(ctx, req.Params); err != nil {
		return nil, err
	}

	return &types.MsgUpdateParamsResponse{}, nil
}

func (k msgServer) StableSwap(ctx context.Context, msg *types.MsgStableSwap) (_ *types.MsgStableSwapResponse, err error) {
	// validate msg
	if err = msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// balance check

	asset := k.keeper.BankKeeper.GetBalance(ctx, sdk.AccAddress(msg.Address), msg.OfferCoin.Denom)
	if asset.Amount.LT(msg.OfferCoin.Amount) {
		return nil, fmt.Errorf("insufficient balance")
	}

	accAddress, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}
	if strings.Contains(msg.OfferCoin.Denom, types.ReserveStableCoinDenomPrefix) {
		err = k.keeper.SwapToOtherStablecoin(ctx, accAddress, msg.OfferCoin, msg.ExpectedDenom)
	} else {
		err = k.keeper.SwapToOnomyStableToken(ctx, accAddress, msg.OfferCoin, msg.ExpectedDenom)
	}

	if err != nil {
		return nil, err
	}

	return &types.MsgStableSwapResponse{}, nil
}

func (k msgServer) AddStableCoinProposal(ctx context.Context, msg *types.MsgAddStableCoin) (*types.MsgAddStableCoinResponse, error) {
	if k.keeper.authority != msg.Authority {
		return nil, errorsmod.Wrapf(types.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.keeper.authority, msg.Authority)
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if err := msg.ValidateBasic(); err != nil {
		return &types.MsgAddStableCoinResponse{}, err
	}

	_, err := k.keeper.StablecoinInfo.Get(ctx, msg.Denom)
	if err == nil {
		return &types.MsgAddStableCoinResponse{}, fmt.Errorf("%s has existed", msg.Denom)
	}

	err = k.keeper.StablecoinInfo.Set(ctx, msg.Denom, types.GetMsgStablecoin(msg))
	if err != nil {
		return &types.MsgAddStableCoinResponse{}, err
	}

	err = k.keeper.OracleKeeper.AddNewSymbolToBandOracleRequest(ctx, msg.Denom, 1)
	if err != nil {
		return &types.MsgAddStableCoinResponse{}, err
	}

	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventAddStablecoin,
			sdk.NewAttribute(types.AttributeStablecoinName, msg.Denom),
		),
	)
	return &types.MsgAddStableCoinResponse{}, nil
}

func (k msgServer) UpdatesStableCoinProposal(ctx context.Context, msg *types.MsgUpdatesStableCoin) (*types.MsgUpdatesStableCoinResponse, error) {
	if k.keeper.authority != msg.Authority {
		return nil, errorsmod.Wrapf(types.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.keeper.authority, msg.Authority)
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if err := msg.ValidateBasic(); err != nil {
		return &types.MsgUpdatesStableCoinResponse{}, err
	}

	oldStablecoin, err := k.keeper.StablecoinInfo.Get(ctx, msg.Denom)
	if err != nil {
		return &types.MsgUpdatesStableCoinResponse{}, fmt.Errorf("%s not existed", msg.Denom)
	}

	newStablecoin := types.GetMsgStablecoin(msg)
	newStablecoin.TotalStablecoinLock = oldStablecoin.TotalStablecoinLock

	err = k.keeper.StablecoinInfo.Set(ctx, newStablecoin.Denom, newStablecoin)
	if err != nil {
		return &types.MsgUpdatesStableCoinResponse{}, err
	}

	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventAddStablecoin,
			sdk.NewAttribute(types.AttributeStablecoinName, msg.Denom),
		),
	)
	return &types.MsgUpdatesStableCoinResponse{}, nil
}

func (k Keeper) checkLimitTotalStablecoin(ctx context.Context, denom string, amountSwap math.Int) error {
	totalStablecoinLock, err := k.TotalStablecoinLock(ctx, denom)
	if err != nil {
		return err
	}

	totalLimit, err := k.GetTotalLimitWithDenomStablecoin(ctx, denom)
	if err != nil {
		return err
	}
	if (totalStablecoinLock.Add(amountSwap)).GT(totalLimit) {
		return fmt.Errorf("unable to perform %s token swap transaction because the amount of %s you want to swap exceeds the allowed limit, can only swap up to %s%s", denom, denom, (totalLimit).Sub(totalStablecoinLock).String(), denom)
	}

	return nil
}
