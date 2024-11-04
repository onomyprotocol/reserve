package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/math"

	errors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	oracletypes "github.com/onomyprotocol/reserve/x/oracle/types"
	"github.com/onomyprotocol/reserve/x/psm/types"
)

// SwapToStablecoin return receiveAmount, fee, error
func (k Keeper) SwapToStablecoin(ctx context.Context, addr sdk.AccAddress, amount math.Int, toDenom string) (math.Int, sdk.DecCoin, error) {
	denomMint := types.DefaultMintDenom
	asset := k.BankKeeper.GetBalance(ctx, addr, denomMint)

	if asset.Amount.LT(amount) {
		return math.ZeroInt(), sdk.DecCoin{}, fmt.Errorf("insufficient balance")
	}

	multiplier := k.OracleKeeper.GetPrice(ctx, toDenom, denomMint)
	if multiplier == nil || multiplier.IsNil() {
		return math.Int{}, sdk.DecCoin{}, errors.Wrapf(oracletypes.ErrInvalidOracle, "can not get price with base %s quote %s", toDenom, denomMint)
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
	denomMint := types.DefaultMintDenom
	asset := k.BankKeeper.GetBalance(ctx, addr, stablecoin.Denom)

	if asset.Amount.LT(stablecoin.Amount) {
		return math.ZeroInt(), sdk.DecCoin{}, fmt.Errorf("insufficient balance")
	}

	multiplier := k.OracleKeeper.GetPrice(ctx, stablecoin.Denom, denomMint)
	if multiplier == nil || multiplier.IsNil() {
		return math.Int{}, sdk.DecCoin{}, errors.Wrapf(oracletypes.ErrInvalidOracle, "can not get price with base %s quote %s", stablecoin.Denom, denomMint)
	}

	amountnomUSD := multiplier.Mul(stablecoin.Amount.ToLegacyDec())

	fee, err := k.PayFeesIn(ctx, amountnomUSD.RoundInt(), stablecoin.Denom)
	if err != nil {
		return math.Int{}, sdk.DecCoin{}, err
	}

	receiveAmountnomUSD := amountnomUSD.Sub(fee)
	return receiveAmountnomUSD.RoundInt(), sdk.NewDecCoinFromDec(denomMint, fee), nil
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
