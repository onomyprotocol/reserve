package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
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

	accAddress, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	// balance check
	asset := k.keeper.BankKeeper.GetBalance(ctx, accAddress, msg.OfferCoin.Denom)
	if asset.Amount.LT(msg.OfferCoin.Amount) {
		return nil, fmt.Errorf("insufficient balance")
	}

	if msg.OfferCoin.Denom == types.ReserveStableCoinDenom {
		err = k.keeper.SwapToOtherStablecoin(ctx, accAddress, msg.OfferCoin, msg.ExpectedDenom)
		if err != nil {
			return nil, err
		}
	} else {
		err = k.keeper.SwapToOnomyStableToken(ctx, accAddress, msg.OfferCoin, msg.ExpectedDenom)
		if err != nil {
			return nil, err
		}
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

	_, err := k.keeper.StablecoinInfos.Get(ctx, msg.Denom)
	if err == nil {
		return &types.MsgAddStableCoinResponse{}, fmt.Errorf("%s has existed", msg.Denom)
	}

	err = k.keeper.StablecoinInfos.Set(ctx, msg.Denom, types.GetMsgStablecoin(msg))
	if err != nil {
		return &types.MsgAddStableCoinResponse{}, err
	}

	err = k.keeper.OracleKeeper.AddNewSymbolToBandOracleRequest(ctx, msg.Symbol, msg.OracleScript)
	if err != nil {
		return &types.MsgAddStableCoinResponse{}, err
	}

	addrPay, err := sdk.AccAddressFromBech32(msg.AddressPayStableInit)
	if err != nil {
		return &types.MsgAddStableCoinResponse{}, err
	}

	err = k.keeper.BankKeeper.SendCoinsFromAccountToModule(ctx, addrPay, types.ModuleName, sdk.NewCoins(sdk.NewCoin(msg.Denom, msg.AmountStableInit)))
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

	oldStablecoin, err := k.keeper.StablecoinInfos.Get(ctx, msg.Denom)
	if err != nil {
		return &types.MsgUpdatesStableCoinResponse{}, fmt.Errorf("%s not existed", msg.Denom)
	}

	newStablecoin := types.GetMsgStablecoin(msg)
	newStablecoin.TotalStablecoinLock = oldStablecoin.TotalStablecoinLock

	err = k.keeper.StablecoinInfos.Set(ctx, newStablecoin.Denom, newStablecoin)
	if err != nil {
		return &types.MsgUpdatesStableCoinResponse{}, err
	}

	err = k.keeper.OracleKeeper.AddNewSymbolToBandOracleRequest(ctx, msg.Symbol, msg.OracleScript)
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

func (k Keeper) checkLimitTotalStablecoin(ctx context.Context, coin sdk.Coin) error {
	totalStablecoinLock, err := k.TotalStablecoinLock(ctx, coin.Denom)
	if err != nil {
		return err
	}

	totalLimit, err := k.GetTotalLimitWithDenomStablecoin(ctx, coin.Denom)
	if err != nil {
		return err
	}
	if (totalStablecoinLock.Add(coin.Amount)).GT(totalLimit) {
		return fmt.Errorf("unable to perform %s token swap transaction: exceeds the allowed limit %s , can only swap up to %s%s", coin.Denom, coin.Denom, (totalLimit).Sub(totalStablecoinLock).String(), coin.Denom)
	}

	return nil
}
