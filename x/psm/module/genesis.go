package psm

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/onomyprotocol/reserve/x/psm/keeper"
	"github.com/onomyprotocol/reserve/x/psm/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) error {
	// this line is used by starport scaffolding # genesis/module/init
	for _, sb := range genState.Stablecoins {
		err := k.SetStablecoin(ctx, sb)
		if err != nil {
			return err
		}
	}

	return k.Params.Set(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) (*types.GenesisState, error) {
	var err error

	genesis := types.DefaultGenesis()
	genesis.Params, err = k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	k.IterateStablecoin(ctx, func(red types.Stablecoin) (stop bool) {
		genesis.Stablecoins = append(genesis.Stablecoins, red)
		return false
	})

	return genesis, nil
}
