package reserve_test

import (
	"testing"

	keepertest "github.com/onomyprotocol/reserve/testutil/keeper"
	"github.com/onomyprotocol/reserve/testutil/nullify"
	reserve "github.com/onomyprotocol/reserve/x/reserve/module"
	"github.com/onomyprotocol/reserve/x/reserve/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx, _ := keepertest.ReserveKeeper(t)
	err := reserve.InitGenesis(ctx, k, genesisState)
	require.NoError(t, err)
	got, err := reserve.ExportGenesis(ctx, k)
	require.NoError(t, err)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
