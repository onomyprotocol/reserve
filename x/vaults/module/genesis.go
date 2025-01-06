package vaults

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/onomyprotocol/reserve/x/vaults/keeper"
	"github.com/onomyprotocol/reserve/x/vaults/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	err := k.InitGenesis(ctx, genState)
	if err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return k.ExportGenesis(ctx)
}
