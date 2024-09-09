package oracle

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/onomyprotocol/reserve/x/oracle/keeper"
	"github.com/onomyprotocol/reserve/x/oracle/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	if err := k.SetParams(ctx, genState.Params); err != nil {
		// TODO: should we panic here
		panic(err)
	}

	for _, bandPriceState := range genState.BandPriceStates {
		k.SetBandPriceState(ctx, bandPriceState.Symbol, bandPriceState)
	}

	for _, bandIBCOracleRequest := range genState.BandOracleRequests {
		k.SetBandOracleRequest(ctx, *bandIBCOracleRequest)
	}

	k.SetBandParams(ctx, genState.BandParams)

	if genState.BandParams.IbcPortId != "" {
		k.SetPort(ctx, genState.BandParams.IbcPortId)
		// Only try to bind to port if it is not already bound, since we may already own port capability
		if !k.IsBound(ctx, genState.BandParams.IbcPortId) {
			// module binds to the port on InitChain
			// and claims the returned capability
			err := k.BindPort(ctx, genState.BandParams.IbcPortId)
			if err != nil {
				panic(types.ErrBadIBCPortBind.Error() + err.Error())
			}
		}
	}
	k.SetPort(ctx, types.PortID)
	// Only try to bind to port if it is not already bound, since we may already own
	// port capability from capability InitGenesis
	if k.ShouldBound(ctx, types.PortID) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, types.PortID)
		if err != nil {
			panic("could not claim port capability: " + err.Error())
		}
	}
	if genState.BandLatestClientId != 0 {
		k.SetBandLatestClientID(ctx, genState.BandLatestClientId)
	}

	for _, record := range genState.CalldataRecords {
		k.SetBandCallDataRecord(ctx, record)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		Params:             k.GetParams(ctx),
		BandParams:         k.GetBandParams(ctx),
		BandPriceStates:    k.GetAllBandPriceStates(ctx),
		BandOracleRequests: k.GetAllBandOracleRequests(ctx),
		BandLatestClientId: k.GetBandLatestClientID(ctx),
		CalldataRecords:    k.GetAllBandCalldataRecords(ctx),
	}
}
