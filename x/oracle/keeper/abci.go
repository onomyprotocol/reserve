package keeper

import (
	// "context"
	"time"

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
