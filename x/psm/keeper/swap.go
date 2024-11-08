package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	oracletypes "github.com/onomyprotocol/reserve/x/oracle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/reserve/x/psm/types"
)

// SwapToStablecoin return receiveAmount, fee, error
func (k Keeper) SwapToOtherStablecoin(ctx context.Context, addr sdk.AccAddress, offerCoin sdk.Coin, expectedDenom string) error {
	// check stablecoin is suport
	stablecoinInfo, err := k.StablecoinInfos.Get(ctx, expectedDenom)
	if err != nil {
		return fmt.Errorf("%s not in list stablecoin supported", expectedDenom)
	}

	// check lock Coin of user
	totalStablecoinLock, err := k.TotalStablecoinLock(ctx, expectedDenom)
	if err != nil {
		return err
	}

	rate := k.OracleKeeper.GetPrice(ctx, expectedDenom, types.USD)
	if rate == nil || rate.IsNil() {
		return errors.Wrapf(oracletypes.ErrInvalidOracle, "can not get price with base %s quote %s", offerCoin.Denom, types.USD)
	}

	expectedAmount := offerCoin.Amount.ToLegacyDec().Quo(*rate).RoundInt()

	// locked stablecoin is greater than the amount desired
	if totalStablecoinLock.LT(expectedAmount) {
		//TODO: register error
		return fmt.Errorf("insufficient balance, PSM module have %d USD but request %d", totalStablecoinLock, expectedAmount)
	}

	fee, err := k.PayFeesOut(ctx, expectedAmount, expectedDenom)
	if err != nil {
		return err
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

	// update stable coin info
	stablecoinReceive := sdk.NewCoin(expectedDenom, math.Int(expectedAmount))
	stablecoinInfo.TotalStablecoinLock = totalStablecoinLock.Sub(expectedAmount)
	err = k.StablecoinInfos.Set(ctx, expectedDenom, stablecoinInfo)
	if err != nil {
		return err
	}

	// send stablecoin to user
	err = k.BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, sdk.NewCoins(stablecoinReceive))
	if err != nil {
		return err
	}

	// event
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventSwapTonomUSD,
			sdk.NewAttribute(types.AttributeAmount, offerCoin.String()),
			sdk.NewAttribute(types.AttributeReceive, expectedAmount.String()),
			sdk.NewAttribute(types.AttributeFeeIn, fee.String()),
		),
	)
	return nil
}

func (k Keeper) SwapToOnomyStableToken(ctx context.Context, accAddress sdk.AccAddress, offerCoin sdk.Coin, expectedDenom string) error {
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
