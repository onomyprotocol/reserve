package app

// import (
// 	storetypes "cosmossdk.io/store/types"
// 	oraclekeeper "github.com/onomyprotocol/reserve/x/oracle/keeper"
// 	oraclemodule "github.com/onomyprotocol/reserve/x/oracle/module"
// 	oracletypes "github.com/onomyprotocol/reserve/x/oracle/types"
// 	"github.com/cosmos/cosmos-sdk/runtime"
// 	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
// 	ibcfee "github.com/cosmos/ibc-go/v8/modules/apps/29-fee"
// 	porttypes "github.com/cosmos/ibc-go/v8/modules/core/05-port/types"
// )

// // registerOracleModule register Oracle keepers and non dependency inject modules.
// func (app *App) registerOracleModule() (porttypes.IBCModule, error) {
// 	// set up non depinject support modules store keys
// 	if err := app.RegisterStores(
// 		storetypes.NewKVStoreKey(oracletypes.StoreKey),
// 	); err != nil {
// 		panic(err)
// 	}

// 	// register the key tables for legacy param subspaces
// 	app.ParamsKeeper.Subspace(oracletypes.ModuleName).WithKeyTable(oracletypes.ParamKeyTable())
// 	// add capability keeper and ScopeToModule for oracle ibc module
// 	scopedOralceKeeper := app.CapabilityKeeper.ScopeToModule(oracletypes.ModuleName)

// 	app.OracleKeeper = oraclekeeper.NewKeeper(
// 		app.AppCodec(),
// 		runtime.NewKVStoreService(app.GetKey(oracletypes.StoreKey)),
// 		app.Logger(),
// 		authtypes.NewModuleAddress(oracletypes.ModuleName).String(),
// 		app.GetIBCKeeper,
// 		scopedOralceKeeper,
// 	)

// 	// register IBC modules
// 	if err := app.RegisterModules(
// 		oraclemodule.NewAppModule(
// 			app.AppCodec(),
// 			app.OracleKeeper,
// 			app.AccountKeeper,
// 			app.BankKeeper,
// 		)); err != nil {
// 		return nil, err
// 	}

// 	app.ScopedOracleKeeper = scopedOralceKeeper

// 	// Create fee enabled ibc stack for oracel
// 	var oracleStack porttypes.IBCModule
// 	oracleStack = oraclemodule.NewIBCModule(app.OracleKeeper)
// 	oracleStack = ibcfee.NewIBCMiddleware(oracleStack, app.IBCFeeKeeper)

// 	return oracleStack, nil
// }
