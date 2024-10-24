package keeper

import (
	"context"

	"cosmossdk.io/math"
	"github.com/onomyprotocol/reserve/x/psm/types"
)

func (k Keeper) BeginBlocker(ctx context.Context) error {
	return k.UpdatesStablecoinEpoch(ctx)
}

func (k Keeper) UpdatesStablecoinEpoch(ctx context.Context) error {
	updatePrice := func(red types.Stablecoin) bool {
		price := k.OracleKeeper.GetPrice(ctx, red.Denom, types.DenomStable)
		if price == nil {
			return false
		}

		sc := k.stablecoinUpdate(ctx, *price, red)
		err := k.SetStablecoin(ctx, sc)
		if err != nil {
			return false
		}
		return false
	}

	return k.IterateStablecoin(ctx, updatePrice)
}

// price is $nomUSD amount to exchange for 1 $stabalecoin
// price taget = 1
//	 	ex:
// 			oldPrice:1
// 			feeIn  : 0.01
// 			feeOut : 0.01
// 			maxFee = 0.02
// 			k = AdjustmentFeeIn = 40
// 			----------------------------------------------------------------------------------------
// 			case 1:
//			newPrice: 1.01 (1.01$nomUSD = 1USDT)
// 			rate = 1/1.01 = 0.990099
// 			newfeeOut = 0.01/(0.990099)**k = 0.01 * (1.01**40)= 0.014888637335882209
//			newfeeIn  = 0.02 - 0.014888637335882209 = 0.005111362664117791

// 			So $USDT swap to $nomUSD will be cheaper than $nomUSD swap to $USDT
// 			----------------------------------------------------------------------------------------
// 			case 2:
//			newPrice: 0.99 (0.98$nomUSD = 1USDT)
// 			rate = 1/0.99 = 1.0101010101
// 			deltaP < 0
//			newfeeIn  = 0.01 * (1.0101010101)**40 = 0.014948314143157351
// 			newfeeOut = 0.02 - 0.014948314143157351 = 0.005051685856842649
// 			So $nomUSD swap to $USDT will be cheaper than $USDT swap to $nomUSD

func (k Keeper) stablecoinUpdate(ctx context.Context, newPrice math.LegacyDec, stablecoin types.Stablecoin) types.Stablecoin {
	params, err := k.GetParams(ctx)
	if err != nil {
		panic(err)
	}
	deltaP := newPrice.Sub(math.LegacyMustNewDecFromStr("1"))
	if deltaP.Abs().LT(params.AcceptablePriceRatio) {
		return stablecoin
	}
	// fee in anf out < fee_in +fee_out
	feeMax, err := k.FeeMaxStablecoin.Get(ctx, stablecoin.Denom)
	if err != nil {
		panic(err)
	}

	rate := math.LegacyOneDec().Quo(newPrice)
	if rate.LT(math.LegacyOneDec()) {
		feeOut := math.LegacyMustNewDecFromStr(feeMax).QuoInt64(2)
		for i := 0; i < int(params.AdjustmentFee); i++ {
			feeOut = feeOut.Quo(rate)
		}
		feeOut = math.LegacyMinDec(feeOut, math.LegacyMustNewDecFromStr(feeMax))
		feeIn := math.LegacyMustNewDecFromStr(feeMax).Sub(feeOut)

		stablecoin.FeeIn = feeIn
		stablecoin.FeeOut = feeOut
	} else {
		feeIn := math.LegacyMustNewDecFromStr(feeMax).QuoInt64(2)
		for i := 0; i < int(params.AdjustmentFee); i++ {
			feeIn = feeIn.Mul(rate)
		}
		feeIn = math.LegacyMinDec(feeIn, math.LegacyMustNewDecFromStr(feeMax))
		feeOut := math.LegacyMustNewDecFromStr(feeMax).Sub(feeIn)

		stablecoin.FeeIn = feeIn
		stablecoin.FeeOut = feeOut
	}
	return stablecoin
}
