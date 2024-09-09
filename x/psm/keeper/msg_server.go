package keeper

import (
	"fmt"
	"github.com/onomyprotocol/reserve/x/psm/types"

	"context"
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

func (k msgServer) SwapToIST(goCtx context.Context, msg *types.MsgSwapToIST) (*types.MsgSwapToISTResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	_, found := k.keeper.GetStablecoin(ctx, msg.Coin.Denom)
	if !found {
		return nil, fmt.Errorf("%s not in list stablecoin supported", msg.Coin.Denom)
	}

	moduleAddr := k.keeper.accountKeeper.GetModuleAddress(types.ModuleName)
	totalStablecoinLock := k.keeper.bankKeeper.GetBalance(goCtx, moduleAddr, msg.Coin.Denom).Amount
	totalLimit, err := k.keeper.GetTotalLimitWithDenomStablecoin(ctx, msg.Coin.Denom)
	if err != nil {
		return nil, err
	}
	if (totalStablecoinLock.Add(msg.Coin.Amount)).GT(totalLimit) {
		return nil, fmt.Errorf("unable to perform %s token swap transaction because the amount of %s you want to swap exceeds the allowed limit, can only swap up to %s%s", msg.Coin.Denom, msg.Coin.Denom, (totalLimit).Sub(totalStablecoinLock).String(), msg.Coin.Denom)
	}

	addr := sdk.MustAccAddressFromBech32(msg.Address)

	receiveAmountIST, _, err := k.keeper.SwaptoIST(ctx, addr, *msg.Coin)
	if err != nil {
		return nil, err
	}

	// lock msg.Coin for addr
	err = k.keeper.bankKeeper.SendCoinsFromAccountToModule(goCtx, addr, types.ModuleName, sdk.NewCoins(*msg.Coin))
	if err != nil {
		return nil, err
	}
	// mint IST
	coinsMint := sdk.NewCoins(sdk.NewCoin(types.InterStableToken, receiveAmountIST))
	err = k.keeper.bankKeeper.MintCoins(goCtx, types.ModuleName, coinsMint)
	if err != nil {
		return nil, err
	}
	err = k.keeper.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, coinsMint)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventSwapToIST,
			sdk.NewAttribute(types.AttributeAmount, msg.Coin.String()),
			sdk.NewAttribute(types.AttributeReceive, receiveAmountIST.String()+"IST"),
		),
	)
	return &types.MsgSwapToISTResponse{}, nil
}

func (k msgServer) SwapToStablecoin(goCtx context.Context, msg *types.MsgSwapToStablecoin) (*types.MsgSwapToStablecoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	_, found := k.keeper.GetStablecoin(ctx, msg.ToDenom)
	if !found {
		return nil, fmt.Errorf("%s not in list stablecoin supported", msg.ToDenom)
	}
	addr := sdk.MustAccAddressFromBech32(msg.Address)

	amount, _, err := k.keeper.SwapToStablecoin(ctx, addr, msg.Amount, msg.ToDenom)
	if err != nil {
		return nil, err
	}
	// lock msg.Coin for addr
	err = k.keeper.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, sdk.NewCoins(sdk.NewCoin(msg.ToDenom, amount)))
	if err != nil {
		return nil, err
	}
	// burn IST
	coinsBurn := sdk.NewCoins(sdk.NewCoin(types.InterStableToken, msg.Amount))
	err = k.keeper.bankKeeper.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, coinsBurn)
	if err != nil {
		return nil, err
	}
	err = k.keeper.bankKeeper.BurnCoins(ctx, types.ModuleName, coinsBurn)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventSwapToIST,
			sdk.NewAttribute(types.AttributeAmount, msg.Amount.String()+types.InterStableToken),
			sdk.NewAttribute(types.AttributeReceive, amount.String()+msg.ToDenom),
		),
	)
	return &types.MsgSwapToStablecoinResponse{}, nil
}
