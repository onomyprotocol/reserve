package oracle

import (
	"context"

	"github.com/onomyprotocol/reserve/x/oracle/keeper"
	"github.com/onomyprotocol/reserve/x/oracle/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx context.Context, k keeper.Keeper, genState types.GenesisState) {
	if err := k.SetParams(ctx, genState.Params); err != nil {
		// TODO: should we panic here
		panic(err)
	}

	for _, bandPriceState := range genState.BandPriceStates {
		k.SetBandPriceState(ctx, bandPriceState.Symbol, bandPriceState)
	}

	for _, bandOracleRequest := range genState.BandOracleRequests {
		k.SetBandOracleRequest(ctx, *bandOracleRequest)
	}

	k.SetBandParams(ctx, genState.BandParams)

	if genState.BandParams.IbcPortId != "" {
		k.SetPort(ctx, genState.BandParams.IbcPortId)
		// Only try to bind to port if it is not already bound, since we may already own port capability
		if !k.ShouldBound(ctx, genState.BandParams.IbcPortId) {
			// module binds to the port on InitChain
			// and claims the returned capability
			err := k.BindPort(ctx, genState.BandParams.IbcPortId)
			if err != nil {
				panic(types.ErrBandPortBind.Error() + err.Error())
			}
		}
	}

	if genState.BandLatestClientId != 0 {
		k.SetBandLatestClientID(ctx, genState.BandLatestClientId)
	}

	for _, record := range genState.CalldataRecords {
		k.SetBandCallDataRecord(ctx, record)
	}

	if genState.BandLatestRequestId != 0 {
		k.SetBandLatestRequestID(ctx, genState.BandLatestRequestId)
	}

	k.SetBandOracleRequestParams(ctx, genState.BandOracleRequestParams)
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx context.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		Params:                  k.GetParams(ctx),
		BandParams:              k.GetBandParams(ctx),
		BandPriceStates:         k.GetAllBandPriceStates(ctx),
		BandOracleRequests:      k.GetAllBandOracleRequests(ctx),
		BandLatestClientId:      k.GetBandLatestClientID(ctx),
		CalldataRecords:         k.GetAllBandCalldataRecords(ctx),
		BandLatestRequestId:     k.GetBandLatestRequestID(ctx),
		BandOracleRequestParams: k.GetBandOracleRequestParams(ctx),
	}
}
