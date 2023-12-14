package keeper_test

import (
	"context"
	"testing"

	keepertest "reserve/testutil/keeper"
	"reserve/x/reserve/keeper"
	"reserve/x/reserve/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ReserveKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
