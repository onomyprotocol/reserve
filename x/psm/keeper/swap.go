package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/onomyprotocol/reserve/x/psm/types"
	vaultstypes "github.com/onomyprotocol/reserve/x/vaults/types"
)

// SwapToStablecoin return receiveAmount, fee, error
func (k Keeper) SwapToOtherStablecoin(ctx context.Context, addr sdk.AccAddress, offerCoin sdk.Coin, expectedDenom string) error {
	// check stablecoin is suport
	stablecoin, err := k.StablecoinInfos.Get(ctx, expectedDenom)
	if err != nil {
		return fmt.Errorf("%s not in list stablecoin supported", expectedDenom)
	}

	// check lock Coin of user
	totalStablecoinLock, err := k.TotalStablecoinLock(ctx, expectedDenom)
	if err != nil {
		return err
	}

	// check balace and calculate amount of coins received
	receiveAmountStablecoin, fee_out, err := k.calculateSwapToStablecoin(ctx, offerCoin.Amount, stablecoin)
	if err != nil {
		return err
	}

	// locked stablecoin is greater than the amount desired
	if totalStablecoinLock.LT(receiveAmountStablecoin.Add(fee_out.Amount.TruncateInt())) {
		return fmt.Errorf("amount %s locked lesser than amount desired", expectedDenom)
	}

	// burn nomUSD
	coinsBurn := sdk.NewCoins(offerCoin)
	err = k.BankKeeper.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, coinsBurn)
	if err != nil {
		return err
	}
	err = k.BankKeeper.BurnCoins(ctx, types.ModuleName, coinsBurn)
	if err != nil {
		return err
	}

	stablecoinReceive := sdk.NewCoin(expectedDenom, receiveAmountStablecoin)

	// sub total stablecoin lock
	err = k.SubTotalStablecoinLock(ctx, stablecoinReceive)
	if err != nil {
		return err
	}
	// send stablecoin to user
	err = k.BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, sdk.NewCoins(stablecoinReceive))
	if err != nil {
		return err
	}

	coinFee, _ := fee_out.TruncateDecimal()
	err = k.BankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, vaultstypes.ReserveModuleName, sdk.NewCoins(coinFee))
	if err != nil {
		return err
	}

	// event
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventSwapToStablecoin,
			sdk.NewAttribute(types.AttributeAmount, offerCoin.String()),
			sdk.NewAttribute(types.AttributeReceive, stablecoinReceive.String()),
			sdk.NewAttribute(types.AttributeFeeOut, fee_out.String()),
		),
	)
	return nil
}

func (k Keeper) SwapToOnomyStableToken(ctx context.Context, accAddress sdk.AccAddress, offerCoin sdk.Coin, expectedDenom string) error {
	// check stablecoin is suport
	stablecoin, err := k.StablecoinInfos.Get(ctx, offerCoin.Denom)
	if err != nil {
		return fmt.Errorf("%s not in list stablecoin supported", offerCoin.Denom)
	}

	// check limit swap
	err = k.checkLimitTotalStablecoin(ctx, offerCoin)
	if err != nil {
		return err
	}

	// check balance user and calculate amount of coins received
	receiveAmountnomUSD, fee_in, err := k.calculateSwapToOnomyStableToken(ctx, offerCoin, stablecoin.Symbol)
	if err != nil {
		return err
	}

	// add total stablecoin lock
	err = k.AddTotalStablecoinLock(ctx, offerCoin)
	if err != nil {
		return err
	}

	// send stablecoin to module
	err = k.BankKeeper.SendCoinsFromAccountToModule(ctx, accAddress, types.ModuleName, sdk.NewCoins(offerCoin))
	if err != nil {
		return err
	}

	// mint nomUSD
	coinsMint := sdk.NewCoins(sdk.NewCoin(types.ReserveStableCoinDenom, receiveAmountnomUSD))
	err = k.BankKeeper.MintCoins(ctx, types.ModuleName, coinsMint)
	if err != nil {
		return err
	}
	// send to user
	err = k.BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, accAddress, coinsMint)
	if err != nil {
		return err
	}

	// event
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventSwapTonomUSD,
			sdk.NewAttribute(types.AttributeAmount, offerCoin.String()),
			sdk.NewAttribute(types.AttributeReceive, coinsMint.String()),
			sdk.NewAttribute(types.AttributeFeeIn, fee_in.String()),
		),
	)
	return nil
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

// SwapToStablecoin return receiveAmount, fee, error
func (k Keeper) calculateSwapToStablecoin(ctx context.Context, amount math.Int, sc types.StablecoinInfo) (math.Int, sdk.DecCoin, error) {
	multiplier, err := k.OracleKeeper.GetPrice(ctx, sc.Symbol, types.SymBolUSD)
	if err != nil {
		return math.Int{}, sdk.DecCoin{}, err
	}
	amountStablecoin := amount.ToLegacyDec().Quo(multiplier)

	fee, err := k.PayFeesOut(ctx, amountStablecoin.RoundInt(), sc.Denom)
	if err != nil {
		return math.Int{}, sdk.DecCoin{}, err
	}

	receiveAmount := amountStablecoin.Sub(fee)
	return receiveAmount.RoundInt(), sdk.NewDecCoinFromDec(sc.Denom, fee), nil
}

func (k Keeper) calculateSwapToOnomyStableToken(ctx context.Context, stablecoin sdk.Coin, symBol string) (math.Int, sdk.DecCoin, error) {
	multiplier, err := k.OracleKeeper.GetPrice(ctx, symBol, types.SymBolUSD)
	if err != nil {
		return math.Int{}, sdk.DecCoin{}, err
	}

	amountnomUSD := multiplier.Mul(stablecoin.Amount.ToLegacyDec())

	fee, err := k.PayFeesIn(ctx, amountnomUSD.RoundInt(), stablecoin.Denom)
	if err != nil {
		return math.Int{}, sdk.DecCoin{}, err
	}

	receiveAmountnomUSD := amountnomUSD.Sub(fee)
	return receiveAmountnomUSD.RoundInt(), sdk.NewDecCoinFromDec(types.ReserveStableCoinDenom, fee), nil
}
