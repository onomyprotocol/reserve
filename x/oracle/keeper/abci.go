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

	// Request oracle prices using band IBC in frequent intervals
	if ctx.BlockHeight()%bandParams.IbcRequestInterval == 0 {
		k.RequestAllBandRates(ctx)
	}

	// todo: default cleanup interval (1 day)
	if ctx.BlockHeight()%86400 == 0 {
		k.CleanUpStaleBandCalldataRecords(ctx)
	}

	data := k.GetAllBandPriceStates(ctx)
	println("check band price state in begin block")
	for _, state := range data {
		println("check clmm: ", state.String())
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
	data := k.GetAllBandPriceStates(ctx)
	println("check band price state in end block")
	for _, state := range data {
		println("check clmm: ", state.String())
	}
}
