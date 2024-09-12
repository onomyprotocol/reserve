package keeper

import (
	"context"
	"cosmossdk.io/math"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/onomyprotocol/reserve/x/psm/types"
)

// SwapToStablecoin return receiveAmount, fee, error
func (k Keeper) SwapToStablecoin(ctx context.Context, addr sdk.AccAddress, amount math.Int, toDenom string) (math.Int, sdk.DecCoin, error) {
	asset := k.BankKeeper.GetBalance(ctx, addr, types.InterStableToken)

	if asset.Amount.LT(amount) {
		return math.ZeroInt(), sdk.DecCoin{}, fmt.Errorf("insufficient balance")
	}

	multiplier, err := k.GetPrice(ctx, toDenom)
	if err != nil || multiplier.IsZero() {
		return math.Int{}, sdk.DecCoin{}, err
	}
	amountStablecoin := amount.ToLegacyDec().Quo(multiplier)

	fee, err := k.PayFeesOut(ctx, amountStablecoin.RoundInt(), toDenom)
	if err != nil {
		return math.Int{}, sdk.DecCoin{}, err
	}

	receiveAmount := amountStablecoin.Sub(fee)
	return receiveAmount.RoundInt(), sdk.NewDecCoinFromDec(toDenom, fee), nil
}

func (k Keeper) SwapToIST(ctx context.Context, addr sdk.AccAddress, stablecoin sdk.Coin) (math.Int, sdk.DecCoin, error) {
	asset := k.BankKeeper.GetBalance(ctx, addr, stablecoin.Denom)

	if asset.Amount.LT(stablecoin.Amount) {
		return math.ZeroInt(), sdk.DecCoin{}, fmt.Errorf("insufficient balance")
	}

	multiplier, err := k.GetPrice(ctx, stablecoin.Denom)
	if err != nil || multiplier.IsZero() {
		return math.Int{}, sdk.DecCoin{}, err
	}

	amountIST := multiplier.Mul(stablecoin.Amount.ToLegacyDec())

	fee, err := k.PayFeesIn(ctx, amountIST.RoundInt(), stablecoin.Denom)
	if err != nil {
		return math.Int{}, sdk.DecCoin{}, err
	}

	receiveAmountIST := amountIST.Sub(fee)
	return receiveAmountIST.RoundInt(), sdk.NewDecCoinFromDec(types.InterStableToken, fee), nil
}

func (k Keeper) PayFeesOut(ctx context.Context, amount math.Int, denom string) (math.LegacyDec, error) {
	ratioSwapOutFees, err := k.GetFeeOut(ctx, denom)
	if err != nil {
		return math.LegacyDec{}, err
	}

	fee := ratioSwapOutFees.MulInt(amount)
	return fee, nil
}

func (k Keeper) PayFeesIn(ctx context.Context, amount math.Int, denom string) (math.LegacyDec, error) {
	ratioSwapInFees, err := k.GetFeeIn(ctx, denom)
	if err != nil {
		return math.LegacyDec{}, err
	}
	fee := ratioSwapInFees.MulInt(amount)
	return fee, nil
}
