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
	asset := k.BankKeeper.GetBalance(ctx, addr, types.DenomStable)

	if asset.Amount.LT(amount) {
		return math.ZeroInt(), sdk.DecCoin{}, fmt.Errorf("insufficient balance")
	}

	multiplier := k.OracleKeeper.GetPrice(ctx, toDenom, types.DenomStable)
	if multiplier == nil {
		return math.Int{}, sdk.DecCoin{}, fmt.Errorf("can not get price %s and %s", toDenom, types.DenomStable)
	}
	amountStablecoin := amount.ToLegacyDec().Quo(*multiplier)

	fee, err := k.PayFeesOut(ctx, amountStablecoin.RoundInt(), toDenom)
	if err != nil {
		return math.Int{}, sdk.DecCoin{}, err
	}

	receiveAmount := amountStablecoin.Sub(fee)
	return receiveAmount.RoundInt(), sdk.NewDecCoinFromDec(toDenom, fee), nil
}

func (k Keeper) SwapTonomUSD(ctx context.Context, addr sdk.AccAddress, stablecoin sdk.Coin) (math.Int, sdk.DecCoin, error) {
	asset := k.BankKeeper.GetBalance(ctx, addr, stablecoin.Denom)

	if asset.Amount.LT(stablecoin.Amount) {
		return math.ZeroInt(), sdk.DecCoin{}, fmt.Errorf("insufficient balance")
	}

	multiplier := k.OracleKeeper.GetPrice(ctx, stablecoin.Denom, types.DenomStable)
	if multiplier == nil {
		return math.Int{}, sdk.DecCoin{}, fmt.Errorf("can not get price %s and %s", stablecoin.Denom, types.DenomStable)
	}

	amountnomUSD := multiplier.Mul(stablecoin.Amount.ToLegacyDec())

	fee, err := k.PayFeesIn(ctx, amountnomUSD.RoundInt(), stablecoin.Denom)
	if err != nil {
		return math.Int{}, sdk.DecCoin{}, err
	}

	receiveAmountnomUSD := amountnomUSD.Sub(fee)
	return receiveAmountnomUSD.RoundInt(), sdk.NewDecCoinFromDec(types.DenomStable, fee), nil
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
