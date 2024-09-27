package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/onomyprotocol/reserve/testutil/keeper"
	"github.com/onomyprotocol/reserve/x/oracle/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.OracleKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
