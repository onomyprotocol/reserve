package keeper

import (
	"context"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/reserve/x/oracle/types"
)

func (k *Keeper) BeginBlocker(ctx sdk.Context) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	bandIBCParams := k.GetBandParams(ctx)

	// Request oracle prices using band IBC in frequent intervals
	if ctx.BlockHeight()%bandIBCParams.IbcRequestInterval == 0 {
		k.RequestAllBandRates(ctx)
	}
}

func (k *Keeper) RequestAllBandRates(ctx sdk.Context) {
	bandIBCOracleRequests := k.GetAllBandOracleRequests(ctx)

	if len(bandIBCOracleRequests) == 0 {
		return
	}

	for _, req := range bandIBCOracleRequests {
		err := k.RequestBandOraclePrices(ctx, req)
		if err != nil {
			ctx.Logger().Error(err.Error())
		}
	}
}

func (k *Keeper) EndBlocker(ctx context.Context) {
}
