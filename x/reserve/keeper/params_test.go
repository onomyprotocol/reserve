package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "reserve/testutil/keeper"
	"reserve/x/reserve/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ReserveKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.ProviderChannel, k.ProviderChannel(ctx))
	require.EqualValues(t, params.MarketChannel, k.MarketChannel(ctx))
	require.EqualValues(t, params.MarketCollateral, k.MarketCollateral(ctx))
	require.EqualValues(t, params.ReserveCollateral, k.ReserveCollateral(ctx))
	require.EqualValues(t, params.CollateralDeposit, k.CollateralDeposit(ctx))
}
