package auction

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/onomyprotocol/reserve/x/auction/keeper"
	"github.com/onomyprotocol/reserve/x/auction/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.InitGenesis(ctx, genState)
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return k.ExportGenesis(ctx)
}
