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
	bandParams := k.GetBandParams(ctx)

	// Request oracle prices using band IBC in frequent intervals
	if ctx.BlockHeight()%bandParams.IbcRequestInterval == 0 {
		k.RequestAllBandRates(ctx)
	}
}

func (k *Keeper) RequestAllBandRates(ctx sdk.Context) {
	// TODO: check logic flow for this
	bandOracleRequests := k.GetAllBandOracleRequests(ctx)

	if len(bandOracleRequests) == 0 {
		return
	}

	for _, req := range bandOracleRequests {
		err := k.RequestBandOraclePrices(ctx, req)
		if err != nil {
			ctx.Logger().Error(err.Error())
		}
	}
	// TODO: Clean call data record after each 1000 blocks
}

func (k *Keeper) EndBlocker(ctx context.Context) {
}
