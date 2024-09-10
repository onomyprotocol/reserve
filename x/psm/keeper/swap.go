package keeper

import (
	"context"
	"cosmossdk.io/math"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/onomyprotocol/reserve/x/psm/types"
)

func (k Keeper) GetTotalLimitWithDenomStablecoin(ctx context.Context, denom string) (math.Int, error) {
	s, found := k.GetStablecoin(ctx, denom)
	if !found {
		return math.Int{}, fmt.Errorf("not found Stable coin %s", denom)
	}
	return s.LimitTotal, nil
}

func (k Keeper) SwapToStablecoin(ctx context.Context, addr sdk.AccAddress, amount math.Int, toDenom string) (math.Int, sdk.Coin, error) {
	asset := k.bankKeeper.GetBalance(ctx, addr, types.InterStableToken)

	if asset.Amount.LT(amount) {
		return math.ZeroInt(), sdk.Coin{}, fmt.Errorf("insufficient balance")
	}

	multiplier, err := k.GetPrice(ctx, toDenom)
	if err != nil || multiplier.IsZero() {
		return math.Int{}, sdk.Coin{}, err
	}
	amountStablecoin := amount.ToLegacyDec().Quo(multiplier).RoundInt()

	fee, err := k.PayFeesOut(ctx, amountStablecoin, toDenom)
	if err != nil {
		return math.Int{}, sdk.Coin{}, err
	}

	receiveAmount := amountStablecoin.Sub(fee)
	return receiveAmount, sdk.NewCoin(toDenom, fee), nil
}

func (k Keeper) SwaptoIST(ctx context.Context, addr sdk.AccAddress, stablecoin sdk.Coin) (math.Int, sdk.Coin, error) {
	asset := k.bankKeeper.GetBalance(ctx, addr, stablecoin.Denom)

	if asset.Amount.LT(stablecoin.Amount) {
		return math.ZeroInt(), sdk.Coin{}, fmt.Errorf("insufficient balance")
	}

	multiplier, err := k.GetPrice(ctx, stablecoin.Denom)
	if err != nil || multiplier.IsZero() {
		return math.Int{}, sdk.Coin{}, err
	}

	amountIST := multiplier.Mul(stablecoin.Amount.ToLegacyDec()).RoundInt()

	fee, err := k.PayFeesIn(ctx, amountIST, stablecoin.Denom)
	if err != nil {
		return math.Int{}, sdk.Coin{}, err
	}

	receiveAmountIST := amountIST.Sub(fee)
	return receiveAmountIST, sdk.NewCoin(types.InterStableToken, fee), nil
}

func (k Keeper) GetPrice(ctx context.Context, denom string) (math.LegacyDec, error) {
	s, found := k.GetStablecoin(ctx, denom)
	if !found {
		return math.LegacyDec{}, fmt.Errorf("not found Stable coin %s", denom)
	}
	return s.Price, nil
}

func (k Keeper) GetFeeIn(ctx context.Context, denom string) (math.LegacyDec, error) {
	s, found := k.GetStablecoin(ctx, denom)
	if !found {
		return math.LegacyDec{}, fmt.Errorf("not found Stable coin %s", denom)
	}
	return s.FeeIn, nil
}

func (k Keeper) GetFeeOut(ctx context.Context, denom string) (math.LegacyDec, error) {
	s, found := k.GetStablecoin(ctx, denom)
	if !found {
		return math.LegacyDec{}, fmt.Errorf("not found Stable coin %s", denom)
	}
	return s.FeeOut, nil
}

func (k Keeper) PayFeesOut(ctx context.Context, amount math.Int, denom string) (math.Int, error) {
	ratioSwapOutFees, err := k.GetFeeOut(ctx, denom)
	if err != nil {
		return math.Int{}, err
	}

	fee := ratioSwapOutFees.MulInt(amount).RoundInt()
	return fee, nil
}

func (k Keeper) PayFeesIn(ctx context.Context, amount math.Int, denom string) (math.Int, error) {
	ratioSwapInFees, err := k.GetFeeIn(ctx, denom)
	if err != nil {
		return math.Int{}, err
	}
	fee := ratioSwapInFees.MulInt(amount).RoundInt()
	return fee, nil
}
