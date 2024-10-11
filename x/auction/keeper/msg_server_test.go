package keeper_test

// import (
// 	"context"
// 	"testing"

// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	"github.com/stretchr/testify/require"

// 	keepertest "github.com/onomyprotocol/reserve/testutil/keeper"
// 	"github.com/onomyprotocol/reserve/x/oracle/keeper"
// 	"github.com/onomyprotocol/reserve/x/oracle/types"
// )

// func setupMsgServer(t testing.TB) (keeper.Keeper, types.MsgServer, context.Context) {
// 	k, ctx := keepertest.OracleKeeper(t)
// 	return k, keeper.NewMsgServerImpl(k), ctx
// }

// func TestMsgServer(t *testing.T) {
// 	k, ms, ctx := setupMsgServer(t)
// 	require.NotNil(t, ms)
// 	require.NotNil(t, ctx)
// 	require.NotEmpty(t, k)
// }

// func TestMsgUpdateParams(t *testing.T) {
// 	k, ms, ctx := setupMsgServer(t)
// 	params := types.DefaultParams()
// 	require.NoError(t, k.SetParams(ctx, params))
// 	wctx := sdk.UnwrapSDKContext(ctx)

// 	// default params
// 	testCases := []struct {
// 		name      string
// 		input     *types.MsgUpdateParams
// 		expErr    bool
// 		expErrMsg string
// 	}{
// 		{
// 			name: "invalid authority",
// 			input: &types.MsgUpdateParams{
// 				Authority: "invalid",
// 				Params:    params,
// 			},
// 			expErr:    true,
// 			expErrMsg: "invalid authority",
// 		},
// 		{
// 			name: "send enabled param",
// 			input: &types.MsgUpdateParams{
// 				Authority: k.GetAuthority(),
// 				Params:    types.Params{},
// 			},
// 			expErr: false,
// 		},
// 		{
// 			name: "all good",
// 			input: &types.MsgUpdateParams{
// 				Authority: k.GetAuthority(),
// 				Params:    params,
// 			},
// 			expErr: false,
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			_, err := ms.UpdateParams(wctx, tc.input)

// 			if tc.expErr {
// 				require.Error(t, err)
// 				require.Contains(t, err.Error(), tc.expErrMsg)
// 			} else {
// 				require.NoError(t, err)
// 			}
// 		})
// 	}
// }
