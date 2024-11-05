package keeper

// import (
// 	"testing"

// 	"cosmossdk.io/log"
// 	"cosmossdk.io/store"
// 	"cosmossdk.io/store/metrics"
// 	storetypes "cosmossdk.io/store/types"
// 	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
// 	dbm "github.com/cosmos/cosmos-db"
// 	"github.com/cosmos/cosmos-sdk/codec"
// 	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
// 	"github.com/cosmos/cosmos-sdk/runtime"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
// 	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
// 	capabilitykeeper "github.com/cosmos/ibc-go/modules/capability/keeper"
// 	portkeeper "github.com/cosmos/ibc-go/v8/modules/core/05-port/keeper"
// 	channelkeeper "github.com/cosmos/ibc-go/v8/modules/core/04-channel/keeper"
// 	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"
// 	"github.com/stretchr/testify/require"

// 	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
// 	address "github.com/cosmos/cosmos-sdk/codec/address"

// 	"github.com/onomyprotocol/reserve/x/oracle/keeper"
// 	"github.com/onomyprotocol/reserve/x/oracle/types"
// )

// func OracleKeeper(t testing.TB) (keeper.Keeper, sdk.Context) {
// 	storeKey := storetypes.NewKVStoreKey(types.StoreKey)
// 	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

// 	db := dbm.NewMemDB()
// 	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
// 	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
// 	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
// 	require.NoError(t, stateStore.LoadLatestVersion())

// 	registry := codectypes.NewInterfaceRegistry()
// 	appCodec := codec.NewProtoCodec(registry)
// 	capabilityKeeper := capabilitykeeper.NewKeeper(appCodec, storeKey, memStoreKey)
// 	authority := authtypes.NewModuleAddress(govtypes.ModuleName)

// 	scopedKeeper := capabilityKeeper.ScopeToModule(ibcexported.ModuleName)
// 	portKeeper := portkeeper.NewKeeper(scopedKeeper)
// 	scopeModule := capabilityKeeper.ScopeToModule(types.ModuleName)
// 	channelKeeper := channelkeeper.NewKeeper(appCodec, storeKey,nil, nil, portKeeper, scopeModule)

// 	authKeeper := authkeeper.NewAccountKeeper(appCodec, runtime.NewKVStoreService(storeKey), authtypes.ProtoBaseAccount,
// 		nil,
// 		address.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix()),
// 		sdk.GetConfig().GetBech32AccountAddrPrefix(),
// 		authtypes.NewModuleAddress(govtypes.ModuleName).String(),)

// 	k := keeper.NewKeeper(
// 		appCodec,
// 		runtime.NewKVStoreService(storeKey),
// 		log.NewNopLogger(),
// 		authKeeper,
// 		authority.String(),
// 		channelKeeper,
// 		&portKeeper,
// 		scopedKeeper,
// 	)

// 	ctx := sdk.NewContext(stateStore, cmtproto.Header{}, false, log.NewNopLogger())

// 	// Initialize params
// 	if err := k.SetParams(ctx, types.DefaultParams()); err != nil {
// 		panic(err)
// 	}

// 	return k, ctx
// }
