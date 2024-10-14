package keeper

import (
	"context"
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	"fmt"
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

func (k msgServer) SwapTonomUSD(ctx context.Context, msg *types.MsgSwapTonomUSD) (*types.MsgSwapTonomUSDResponse, error) {
	// validate msg
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// check stablecoin is suport
	_, found := k.keeper.GetStablecoin(ctx, msg.Coin.Denom)
	if !found {
		return nil, fmt.Errorf("%s not in list stablecoin supported", msg.Coin.Denom)
	}

	// check limit swap
	err := k.keeper.checkLimitTotalStablecoin(ctx, msg.Coin.Denom, msg.Coin.Amount)
	if err != nil {
		return nil, err
	}

	// check balance user and calculate amount of coins received
	addr := sdk.MustAccAddressFromBech32(msg.Address)
	receiveAmountnomUSD, fee_in, err := k.keeper.SwapTonomUSD(ctx, addr, *msg.Coin)
	if err != nil {
		return nil, err
	}

	// lock coin
	totalStablecoinLock, err := k.keeper.totalStablecoinLock.Get(ctx, msg.Coin.Denom)
	if err != nil {
		return nil, err
	}
	newTotalStablecoinLock := totalStablecoinLock.Add(msg.Coin.Amount)
	k.keeper.totalStablecoinLock.Set(ctx, msg.Coin.Denom, newTotalStablecoinLock)

	// send stablecoin to module
	err = k.keeper.BankKeeper.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, sdk.NewCoins(*msg.Coin))
	if err != nil {
		return nil, err
	}

	// mint nomUSD
	coinsMint := sdk.NewCoins(sdk.NewCoin(types.DenomStable, receiveAmountnomUSD))
	err = k.keeper.BankKeeper.MintCoins(ctx, types.ModuleName, coinsMint)
	if err != nil {
		return nil, err
	}
	// send to user
	err = k.keeper.BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, coinsMint)
	if err != nil {
		return nil, err
	}

	// event
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventSwapTonomUSD,
			sdk.NewAttribute(types.AttributeAmount, msg.Coin.String()),
			sdk.NewAttribute(types.AttributeReceive, coinsMint.String()),
			sdk.NewAttribute(types.AttributeFeeIn, fee_in.String()),
		),
	)
	return &types.MsgSwapTonomUSDResponse{}, nil
}

func (k msgServer) SwapToStablecoin(ctx context.Context, msg *types.MsgSwapToStablecoin) (*types.MsgSwapToStablecoinResponse, error) {
	// validate basic
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// check stablecoin is suport
	_, found := k.keeper.GetStablecoin(ctx, msg.ToDenom)
	if !found {
		return nil, fmt.Errorf("%s not in list stablecoin supported", msg.ToDenom)
	}

	// check lock Coin of user
	totalStablecoinLock, err := k.keeper.totalStablecoinLock.Get(ctx, msg.ToDenom)
	if err != nil {
		return nil, err
	}

	// check balace and calculate amount of coins received
	addr := sdk.MustAccAddressFromBech32(msg.Address)
	receiveAmountStablecoin, fee_out, err := k.keeper.SwapToStablecoin(ctx, addr, msg.Amount, msg.ToDenom)
	if err != nil {
		return nil, err
	}

	// locked stablecoin is greater than the amount desired
	if totalStablecoinLock.LT(receiveAmountStablecoin) {
		return nil, fmt.Errorf("amount %s locked lesser than amount desired", msg.ToDenom)
	}

	// burn nomUSD
	coinsBurn := sdk.NewCoins(sdk.NewCoin(types.DenomStable, msg.Amount))
	err = k.keeper.BankKeeper.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, coinsBurn)
	if err != nil {
		return nil, err
	}
	err = k.keeper.BankKeeper.BurnCoins(ctx, types.ModuleName, coinsBurn)
	if err != nil {
		return nil, err
	}

	// unlock
	stablecoinReceive := sdk.NewCoin(msg.ToDenom, receiveAmountStablecoin)
	newTotalStablecoinLock := totalStablecoinLock.Sub(receiveAmountStablecoin)
	k.keeper.totalStablecoinLock.Set(ctx, msg.ToDenom, newTotalStablecoinLock)

	// send stablecoin to user
	err = k.keeper.BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, sdk.NewCoins(stablecoinReceive))
	if err != nil {
		return nil, err
	}

	// event
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventSwapToStablecoin,
			sdk.NewAttribute(types.AttributeAmount, msg.Amount.String()+types.DenomStable),
			sdk.NewAttribute(types.AttributeReceive, stablecoinReceive.String()),
			sdk.NewAttribute(types.AttributeFeeOut, fee_out.String()),
		),
	)
	return &types.MsgSwapToStablecoinResponse{}, nil
}

func (k msgServer) AddStableCoinProposal(ctx context.Context, msg *types.MsgAddStableCoin) (*types.MsgAddStableCoinResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if err := msg.ValidateBasic(); err != nil {
		return &types.MsgAddStableCoinResponse{}, err
	}

	_, found := k.keeper.GetStablecoin(ctx, msg.Denom)
	if found {
		return &types.MsgAddStableCoinResponse{}, fmt.Errorf("%s has existed", msg.Denom)
	}

	err := k.keeper.SetStablecoin(ctx, types.GetMsgStablecoin(msg))
	if err != nil {
		return &types.MsgAddStableCoinResponse{}, err
	}

	err = k.keeper.OracleKeeper.AddNewSymbolToBandOracleRequest(ctx, msg.Denom, 1)
	if err != nil {
		return &types.MsgAddStableCoinResponse{}, err
	}

	k.keeper.totalStablecoinLock.Set(ctx, msg.Denom, math.ZeroInt())
	k.keeper.FeeMaxStablecoin.Set(ctx, msg.Denom, msg.FeeIn.Add(msg.FeeOut).String())

	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventAddStablecoin,
			sdk.NewAttribute(types.AttributeStablecoinName, msg.Denom),
		),
	)
	return &types.MsgAddStableCoinResponse{}, nil
}

func (k msgServer) UpdatesStableCoinProposal(ctx context.Context, msg *types.MsgUpdatesStableCoin) (*types.MsgUpdatesStableCoinResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if err := msg.ValidateBasic(); err != nil {
		return &types.MsgUpdatesStableCoinResponse{}, err
	}

	_, found := k.keeper.GetStablecoin(ctx, msg.Denom)
	if !found {
		return &types.MsgUpdatesStableCoinResponse{}, fmt.Errorf("%s not existed", msg.Denom)
	}

	err := k.keeper.SetStablecoin(ctx, types.GetMsgStablecoin(msg))
	if err != nil {
		return &types.MsgUpdatesStableCoinResponse{}, err
	}
	k.keeper.FeeMaxStablecoin.Set(ctx, msg.Denom, msg.FeeIn.Add(msg.FeeOut).String())

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
