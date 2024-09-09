package oracle

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/onomyprotocol/reserve/x/oracle/keeper"
	"github.com/onomyprotocol/reserve/x/oracle/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// TODO: update init genesis
	// this line is used by starport scaffolding # genesis/module/init
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
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	// TODO: update export genesis state
	return &types.GenesisState{}
}
