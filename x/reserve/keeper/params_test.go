package keeper_test

import (
	"testing"

	testkeeper "reserve/testutil/keeper"
	"reserve/x/reserve/types"

	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ReserveKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
