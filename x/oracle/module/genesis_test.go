package oracle_test

// import (
// 	"testing"

// 	keepertest "github.com/onomyprotocol/reserve/testutil/keeper"
// 	"github.com/onomyprotocol/reserve/testutil/nullify"
// 	oracle "github.com/onomyprotocol/reserve/x/oracle/module"
// 	"github.com/onomyprotocol/reserve/x/oracle/types"

// 	"github.com/stretchr/testify/require"
// )

// func TestGenesis(t *testing.T) {
// 	genesisState := types.GenesisState{
// 		Params: types.DefaultParams(),
// 		// this line is used by starport scaffolding # genesis/test/state
// 	}

// 	k, ctx := keepertest.OracleKeeper(t)
// 	oracle.InitGenesis(ctx, k, genesisState)
// 	got := oracle.ExportGenesis(ctx, k)
// 	require.NotNil(t, got)

// 	nullify.Fill(&genesisState)
// 	nullify.Fill(got)

// 	// this line is used by starport scaffolding # genesis/test/assert
// }
