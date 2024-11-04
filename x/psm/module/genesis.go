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
		err := k.Stablecoins.Set(ctx, sb.Denom, sb)
		if err != nil {
			return err
		}
	}

	for _, nom := range genState.Noms {
		err := k.Noms.Set(ctx, nom, nom)
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

	err = k.Stablecoins.Walk(ctx, nil, func(key string, value types.Stablecoin) (stop bool, err error) {
		genesis.Stablecoins = append(genesis.Stablecoins, value)
		return false, nil
	})
	if err != nil {
		return nil, err
	}

	err = k.Noms.Walk(ctx, nil, func(key string, value string) (stop bool, err error) {
		genesis.Noms = append(genesis.Noms, value)
		return false, nil
	})

	return genesis, err
}
