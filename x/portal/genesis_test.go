package portal_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "reserve/testutil/keeper"
	"reserve/testutil/nullify"
	"reserve/x/portal"
	"reserve/x/portal/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PortalKeeper(t)
	portal.InitGenesis(ctx, *k, genesisState)
	got := portal.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
