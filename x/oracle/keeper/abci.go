package keeper

import (
	"context"
	"time"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/reserve/x/oracle/types"
)

func (k *Keeper) BeginBlocker(ctx context.Context) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	bandParams := k.GetBandParams(ctx)
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Request oracle prices using band IBC in frequent intervals
	if sdkCtx.BlockHeight()%bandParams.IbcRequestInterval == 0 {
		k.RequestAllBandRates(ctx)
	}

	// todo: default cleanup interval (1 day)
	// if sdkCtx.BlockHeight()%86400 == 0 {
	// 	k.CleanUpStaleBandCalldataRecords(sdkCtx)
	// }

	bandPriceState := &types.BandPriceState{
		Symbol:      "ATOM",
		Rate:        math.NewInt(10),
		ResolveTime: 1,
		Request_ID:  1,
		PriceState:  *types.NewPriceState(math.LegacyNewDec(10), 1),
	}

	if sdkCtx.BlockHeight()%5 == 0 {
		k.SetBandPriceState(ctx, "ATOM", bandPriceState)
	}

	bandPriceStates, _ := k.GetAllBandPriceStates(ctx)
	for _, data := range bandPriceStates {
		println("Check data band price state for each begin block: ", data.String())
	}
}

func (k *Keeper) RequestAllBandRates(ctx context.Context) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
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
			sdkCtx.Logger().Error(err.Error())
		}
	}
}

func (k *Keeper) EndBlocker(ctx context.Context) {
}
