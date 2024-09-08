package psm_test

import (
	"testing"

	keepertest "github.com/onomyprotocol/reserve/testutil/keeper"
	"github.com/onomyprotocol/reserve/testutil/nullify"
	psm "github.com/onomyprotocol/reserve/x/psm/module"
	"github.com/onomyprotocol/reserve/x/psm/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx, _ := keepertest.PsmKeeper(t)
	err := psm.InitGenesis(ctx, k, genesisState)
	require.NoError(t, err)
	got, err := psm.ExportGenesis(ctx, k)
	require.NoError(t, err)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
