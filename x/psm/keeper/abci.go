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
// 			k_in = AdjustmentFeeIn = 0.05
// 			k_out = AdjustmentFeeOut = 0.05
// 			--------------------------------------
// 			case 1:
//			newPrice: 1.01 (1.01$nomUSD = 1USDT)
// 			deltaP = 1.01 - 1 = 0.01
// 			deltaP > 0
//			newfeeIn  = feeIn - k_in * deltaP = 0.01 - 0.5 * 0.01 = 0.005
// 			newfeeOut = feeOut + k_out * deltaP = 0.01 + 0.5 * 0.01 = 0.015
// 			So $USDT swap to $nomUSD will be cheaper than $nomUSD swap to $USDT
// 			--------------------------------------
// 			case 2:
//			newPrice: 0.98 (0.98$nomUSD = 1USDT)
// 			deltaP = 0.98 - 1 = -0.02
// 			deltaP < 0
//			newfeeIn  = feeIn + k_in * deltaP = 0.01 + 0.5 * 0.02 = 0.02
// 			newfeeOut = feeOut - k_out * deltaP = 0.01 - 0.5 * 0.02 = 0.00
// 			So $nomUSD swap to $USDT will be cheaper than $USDT swap to $nomUSD
//

func (k Keeper) stablecoinUpdate(ctx context.Context, newPrice math.LegacyDec, stablecoin types.Stablecoin) types.Stablecoin {
	params, err := k.GetParams(ctx)
	if err != nil {
		panic(err)
	}
	deltaP := newPrice.Sub(math.LegacyMustNewDecFromStr("1"))
	if deltaP.Abs().LT(params.AcceptablePriceRatio) {
		return stablecoin
	}

	feeIn := stablecoin.FeeIn.Sub(params.AdjustmentFeeIn.Mul(deltaP))
	feeOut := stablecoin.FeeOut.Add(params.AdjustmentFeeOut.Mul(deltaP))

	stablecoin.FeeIn = math.LegacyMaxDec(feeIn, math.LegacyZeroDec())
	stablecoin.FeeOut = math.LegacyMaxDec(feeOut, math.LegacyZeroDec())
	return stablecoin
}
