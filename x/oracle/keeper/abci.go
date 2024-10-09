package keeper

import (
	// "context"
	"time"

	// "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/reserve/x/oracle/types"
)

func (k *Keeper) BeginBlocker(ctx sdk.Context) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	bandParams := k.GetBandParams(ctx)
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Request oracle prices using band IBC in frequent intervals
	if sdkCtx.BlockHeight()%bandParams.IbcRequestInterval == 0 {
		k.RequestAllBandRates(ctx)
	}

	// todo: default cleanup interval (1 day)
	if sdkCtx.BlockHeight()%86400 == 0 {
		k.CleanUpStaleBandCalldataRecords(sdkCtx)
	}

	// bandPriceState := types.BandPriceState{
	// 	Symbol:      "ATOM",
	// 	Rate:        math.NewInt(10),
	// 	ResolveTime: 1,
	// 	Request_ID:  1,
	// 	PriceState:  *types.NewPriceState(math.LegacyNewDec(10), 1),
	// }
	// err := k.BandPriceState.Set(ctx, "ATOM", bandPriceState)
	// if err != nil {
	// 	println("cai dit con me may")
	// }

	data := k.GetAllBandPriceStates(ctx)
	println("check band price state in begin block")
	for _, state := range data {
		println("check state: ", state.String())
	}
}

func (k *Keeper) RequestAllBandRates(ctx sdk.Context) {
	// TODO: check logic flow for this
	bandOracleRequests := k.GetAllBandOracleRequests(ctx)

	if len(bandOracleRequests) == 0 {
		return
	}

	for _, req := range bandOracleRequests {
		println("checking request .......")
		for _, symbol := range req.Symbols {
			println("With symbols: ", symbol)
		}
		err := k.RequestBandOraclePrices(ctx, req)
		if err != nil {
			ctx.Logger().Error(err.Error())
		}
	}
}

func (k *Keeper) EndBlocker(ctx sdk.Context) {
}
