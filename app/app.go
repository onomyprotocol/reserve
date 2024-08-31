package app

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"github.com/spf13/cast"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	circuit "cosmossdk.io/x/circuit"
	circuitkeeper "cosmossdk.io/x/circuit/keeper"
	circuittypes "cosmossdk.io/x/circuit/types"
	evidence "cosmossdk.io/x/evidence"
	evidencekeeper "cosmossdk.io/x/evidence/keeper"
	evidencetypes "cosmossdk.io/x/evidence/types"
	feegrant "cosmossdk.io/x/feegrant"
	feegrantkeeper "cosmossdk.io/x/feegrant/keeper"
	feegrantmodule "cosmossdk.io/x/feegrant/module"
	upgrade "cosmossdk.io/x/upgrade"
	upgradekeeper "cosmossdk.io/x/upgrade/keeper"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/version"
	"cosmossdk.io/x/tx/signing"
	"github.com/cosmos/cosmos-sdk/codec/address"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/cosmos/gogoproto/proto"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"

	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	txmodule "github.com/cosmos/cosmos-sdk/x/auth/tx/config"
	authcodec "github.com/cosmos/cosmos-sdk/x/auth/codec"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	vesting "github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/authz"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	bank "github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensus "github.com/cosmos/cosmos-sdk/x/consensus"
	consensuskeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisis "github.com/cosmos/cosmos-sdk/x/crisis"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/cosmos/cosmos-sdk/x/group"
	groupkeeper "github.com/cosmos/cosmos-sdk/x/group/keeper"
	groupmodule "github.com/cosmos/cosmos-sdk/x/group/module"
	mint "github.com/cosmos/cosmos-sdk/x/mint"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	params "github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	slashing "github.com/cosmos/cosmos-sdk/x/slashing"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	runtimeservices "github.com/cosmos/cosmos-sdk/runtime/services"
	reflectionv1 "cosmossdk.io/api/cosmos/reflection/v1"
	"github.com/cosmos/cosmos-sdk/std"
	abci "github.com/cometbft/cometbft/abci/types"
	nodeservice "github.com/cosmos/cosmos-sdk/client/grpc/node"
	cmtservice "github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"
	sigtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"

	oraclemodule "github.com/onomyprotocol/reserve/x/oracle/module"
	oraclekeeper "github.com/onomyprotocol/reserve/x/oracle/keeper"
	oracletypes "github.com/onomyprotocol/reserve/x/oracle/types"

	// ibc
	ibc "github.com/cosmos/ibc-go/v8/modules/core"
	porttypes "github.com/cosmos/ibc-go/v8/modules/core/05-port/types"
	capabilitykeeper "github.com/cosmos/ibc-go/modules/capability/keeper"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	icatypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/types"
	icamodule "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts" // import for side-effects
	icacontrollerkeeper "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/keeper"
	icacontrollertypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/types"
	icahost "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host"
	icahostkeeper "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host/keeper"
	icahosttypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host/types"
	ibcfee "github.com/cosmos/ibc-go/v8/modules/apps/29-fee"
	ibcfeetypes "github.com/cosmos/ibc-go/v8/modules/apps/29-fee/types"
	ibcfeekeeper "github.com/cosmos/ibc-go/v8/modules/apps/29-fee/keeper"
	ibctransfer "github.com/cosmos/ibc-go/v8/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v8/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	ibcconnectiontypes "github.com/cosmos/ibc-go/v8/modules/core/03-connection/types"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"
	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"
	ibctm "github.com/cosmos/ibc-go/v8/modules/light-clients/07-tendermint"
)

const (
	AccountAddressPrefix = "onomy"
	Name                 = "reserve"
)

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName:     nil,
		distrtypes.ModuleName:          nil,
		minttypes.ModuleName:           {authtypes.Minter},
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:            {authtypes.Burner},
	}
)

var (
	_ runtime.AppI            = (*App)(nil)
	_ servertypes.Application = (*App)(nil)
)

// App extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type App struct {
	*baseapp.BaseApp
	legacyAmino       *codec.LegacyAmino
	appCodec          codec.Codec
	txConfig          client.TxConfig
	interfaceRegistry codectypes.InterfaceRegistry

	// keys to access the substores
	keys  map[string]*storetypes.KVStoreKey
	tkeys map[string]*storetypes.TransientStoreKey
	memKeys map[string]*storetypes.MemoryStoreKey

	// cosmos keepers
	AccountKeeper         authkeeper.AccountKeeper
	BankKeeper            bankkeeper.Keeper
	StakingKeeper         *stakingkeeper.Keeper
	DistrKeeper           distrkeeper.Keeper
	ConsensusParamsKeeper consensuskeeper.Keeper

	SlashingKeeper       slashingkeeper.Keeper
	MintKeeper           mintkeeper.Keeper
	GovKeeper            govkeeper.Keeper
	CrisisKeeper         *crisiskeeper.Keeper
	UpgradeKeeper        *upgradekeeper.Keeper
	ParamsKeeper         paramskeeper.Keeper
	AuthzKeeper          authzkeeper.Keeper
	EvidenceKeeper       evidencekeeper.Keeper
	FeeGrantKeeper       feegrantkeeper.Keeper
	GroupKeeper          groupkeeper.Keeper
	CircuitKeeper        circuitkeeper.Keeper

	// IBC
	IBCKeeper           *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	CapabilityKeeper    *capabilitykeeper.Keeper
	IBCFeeKeeper        ibcfeekeeper.Keeper
	ICAControllerKeeper icacontrollerkeeper.Keeper
	ICAHostKeeper       icahostkeeper.Keeper
	TransferKeeper      ibctransferkeeper.Keeper

	// Scoped IBC
	ScopedIBCKeeper           capabilitykeeper.ScopedKeeper
	ScopedIBCTransferKeeper   capabilitykeeper.ScopedKeeper
	ScopedICAHostKeeper       capabilitykeeper.ScopedKeeper
	ScopedOracleKeeper        capabilitykeeper.ScopedKeeper

	// reserve keeper
	OracleKeeper oraclekeeper.Keeper
	// this line is used by starport scaffolding # stargate/app/keeperDeclaration

	// the module manager
	ModuleManager      *module.Manager
	BasicModuleManager module.BasicManager

	// simulation manager
	sm *module.SimulationManager

	// module configurator
	configurator module.Configurator
}

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, "."+Name)
}

// getGovProposalHandlers return the chain proposal handlers.
func getGovProposalHandlers() []govclient.ProposalHandler {
	var govProposalHandlers []govclient.ProposalHandler
	// this line is used by starport scaffolding # stargate/app/govProposalHandlers

	govProposalHandlers = append(govProposalHandlers,
		paramsclient.ProposalHandler,
		// this line is used by starport scaffolding # stargate/app/govProposalHandler
	)

	return govProposalHandlers
}

// New returns a reference to an initialized App.
func New(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) *App {
	interfaceRegistry, _ := codectypes.NewInterfaceRegistryWithOptions(codectypes.InterfaceRegistryOptions{
		ProtoFiles: proto.HybridResolver,
		SigningOptions: signing.Options{
			AddressCodec: address.Bech32Codec{
				Bech32Prefix: sdk.GetConfig().GetBech32AccountAddrPrefix(),
			},
			ValidatorAddressCodec: address.Bech32Codec{
				Bech32Prefix: sdk.GetConfig().GetBech32ValidatorAddrPrefix(),
			},
		},
	})
	appCodec := codec.NewProtoCodec(interfaceRegistry)
	legacyAmino := codec.NewLegacyAmino()
	txConfig := tx.NewTxConfig(appCodec, tx.DefaultSignModes)

	std.RegisterLegacyAminoCodec(legacyAmino)
	std.RegisterInterfaces(interfaceRegistry)

	// Below we could construct and set an application specific mempool and
	// ABCI 1.0 PrepareProposal and ProcessProposal handlers. These defaults are
	// already set in the SDK's BaseApp, this shows an example of how to override
	// them.
	//
	// Example:
	//
	// bApp := baseapp.NewBaseApp(...)
	// nonceMempool := mempool.NewSenderNonceMempool()
	// abciPropHandler := NewDefaultProposalHandler(nonceMempool, bApp)
	//
	// bApp.SetMempool(nonceMempool)
	// bApp.SetPrepareProposal(abciPropHandler.PrepareProposalHandler())
	// bApp.SetProcessProposal(abciPropHandler.ProcessProposalHandler())
	//
	// Alternatively, you can construct BaseApp options, append those to
	// baseAppOptions and pass them to NewBaseApp.
	//
	// Example:
	//
	// prepareOpt = func(app *baseapp.BaseApp) {
	// 	abciPropHandler := baseapp.NewDefaultProposalHandler(nonceMempool, app)
	// 	app.SetPrepareProposal(abciPropHandler.PrepareProposalHandler())
	// }
	// baseAppOptions = append(baseAppOptions, prepareOpt)

	bApp := baseapp.NewBaseApp(Name, logger, db, txConfig.TxDecoder(), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)
	bApp.SetTxEncoder(txConfig.TxEncoder())

	keys := storetypes.NewKVStoreKeys(
		// SDK keys
		authtypes.StoreKey, banktypes.StoreKey, stakingtypes.StoreKey, crisistypes.StoreKey,
		minttypes.StoreKey, distrtypes.StoreKey, slashingtypes.StoreKey,
		govtypes.StoreKey, paramstypes.StoreKey, consensustypes.StoreKey, upgradetypes.StoreKey, feegrant.StoreKey,
		evidencetypes.StoreKey, circuittypes.StoreKey,
		authzkeeper.StoreKey, group.StoreKey,
		capabilitytypes.StoreKey, ibcexported.StoreKey, ibctransfertypes.StoreKey,
		ibcfeetypes.StoreKey, icahosttypes.StoreKey, icacontrollertypes.StoreKey,

		// Reserve keys
		oracletypes.StoreKey,
	)

	// register streaming services
	if err := bApp.RegisterStreamingServices(appOpts, keys); err != nil {
		panic(err)
	}

	tkeys := storetypes.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := storetypes.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)
	app := &App{
		BaseApp:           bApp,
		legacyAmino:       legacyAmino,
		appCodec:          appCodec,
		txConfig:          txConfig,
		interfaceRegistry: interfaceRegistry,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
	}

	app.ParamsKeeper = initParamsKeeper(appCodec, legacyAmino, keys[paramstypes.StoreKey], tkeys[paramstypes.TStoreKey])

	// set the BaseApp's parameter store
	app.ConsensusParamsKeeper = consensuskeeper.NewKeeper(appCodec, runtime.NewKVStoreService(keys[consensustypes.StoreKey]), authtypes.NewModuleAddress(govtypes.ModuleName).String(), runtime.EventService{})
	bApp.SetParamStore(app.ConsensusParamsKeeper.ParamsStore)

	// add keepers
	app.AccountKeeper = authkeeper.NewAccountKeeper(appCodec, runtime.NewKVStoreService(keys[authtypes.StoreKey]), authtypes.ProtoBaseAccount, maccPerms, authcodec.NewBech32Codec(sdk.Bech32MainPrefix), sdk.Bech32MainPrefix, authtypes.NewModuleAddress(govtypes.ModuleName).String())

	app.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[banktypes.StoreKey]),
		app.AccountKeeper,
		BlockedAddresses(),
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		logger,
	)

	// SDK v0.50
	// Legacy app wiring: to enable SignMode_SIGN_MODE_TEXTUAL app tx config must be updated after bank keeper init
	txConfigOpts := authtx.ConfigOptions{
		EnabledSignModes:           append(authtx.DefaultSignModes, sigtypes.SignMode_SIGN_MODE_TEXTUAL),
		TextualCoinMetadataQueryFn: txmodule.NewBankKeeperCoinMetadataQueryFn(app.BankKeeper),
	}

	txConfig, err := authtx.NewTxConfigWithOptions(app.appCodec, txConfigOpts)
	if err != nil {
		panic("failed to update app tx config: " + err.Error())
	}

	app.txConfig = txConfig

	app.StakingKeeper = stakingkeeper.NewKeeper(
		appCodec, runtime.NewKVStoreService(keys[stakingtypes.StoreKey]), app.AccountKeeper, app.BankKeeper, authtypes.NewModuleAddress(govtypes.ModuleName).String(), authcodec.NewBech32Codec(sdk.Bech32PrefixValAddr), authcodec.NewBech32Codec(sdk.Bech32PrefixConsAddr),
	)
	app.MintKeeper = mintkeeper.NewKeeper(appCodec, runtime.NewKVStoreService(keys[minttypes.StoreKey]), app.StakingKeeper, app.AccountKeeper, app.BankKeeper, authtypes.FeeCollectorName, authtypes.NewModuleAddress(govtypes.ModuleName).String())

	app.DistrKeeper = distrkeeper.NewKeeper(appCodec, runtime.NewKVStoreService(keys[distrtypes.StoreKey]), app.AccountKeeper, app.BankKeeper, app.StakingKeeper, authtypes.FeeCollectorName, authtypes.NewModuleAddress(govtypes.ModuleName).String())

	app.SlashingKeeper = slashingkeeper.NewKeeper(
		appCodec, legacyAmino, runtime.NewKVStoreService(keys[slashingtypes.StoreKey]), app.StakingKeeper, authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	invCheckPeriod := cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod))
	app.CrisisKeeper = crisiskeeper.NewKeeper(appCodec, runtime.NewKVStoreService(keys[crisistypes.StoreKey]), invCheckPeriod,
		app.BankKeeper, authtypes.FeeCollectorName, authtypes.NewModuleAddress(govtypes.ModuleName).String(), app.AccountKeeper.AddressCodec())

	// create evidence keeper with router
	evidenceKeeper := evidencekeeper.NewKeeper(
		appCodec, runtime.NewKVStoreService(keys[evidencetypes.StoreKey]), app.StakingKeeper, app.SlashingKeeper, app.AccountKeeper.AddressCodec(), runtime.ProvideCometInfoService(),
	)
	// If evidence needs to be handled for the app, set routes in router here and seal
	app.EvidenceKeeper = *evidenceKeeper

	app.FeeGrantKeeper = feegrantkeeper.NewKeeper(appCodec, runtime.NewKVStoreService(keys[feegrant.StoreKey]), app.AccountKeeper)

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	app.StakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(app.DistrKeeper.Hooks(), app.SlashingKeeper.Hooks()),
	)

	app.CircuitKeeper = circuitkeeper.NewKeeper(appCodec, runtime.NewKVStoreService(keys[circuittypes.StoreKey]), authtypes.NewModuleAddress(govtypes.ModuleName).String(), app.AccountKeeper.AddressCodec())
	app.BaseApp.SetCircuitBreaker(&app.CircuitKeeper)

	app.AuthzKeeper = authzkeeper.NewKeeper(runtime.NewKVStoreService(keys[authzkeeper.StoreKey]), appCodec, app.MsgServiceRouter(), app.AccountKeeper)

	groupConfig := group.DefaultConfig()
	/*
		Example of setting group params:
		groupConfig.MaxMetadataLen = 1000
	*/
	app.GroupKeeper = groupkeeper.NewKeeper(keys[group.StoreKey], appCodec, app.MsgServiceRouter(), app.AccountKeeper, groupConfig)

	// get skipUpgradeHeights from the app options
	skipUpgradeHeights := map[int64]bool{}
	for _, h := range cast.ToIntSlice(appOpts.Get(server.FlagUnsafeSkipUpgrades)) {
		skipUpgradeHeights[int64(h)] = true
	}
	homePath := cast.ToString(appOpts.Get(flags.FlagHome))
	// set the governance module account as the authority for conducting upgrades
	app.UpgradeKeeper = upgradekeeper.NewKeeper(skipUpgradeHeights, runtime.NewKVStoreService(keys[upgradetypes.StoreKey]), appCodec, homePath, app.BaseApp, authtypes.NewModuleAddress(govtypes.ModuleName).String())

	// Register the proposal types
	// Deprecated: Avoid adding new handlers, instead use the new proposal flow
	// by granting the governance module the right to execute the message.
	// See: https://docs.cosmos.network/main/modules/gov#proposal-messages
	govRouter := govv1beta1.NewRouter()
	govRouter.AddRoute(govtypes.RouterKey, govv1beta1.ProposalHandler).
		AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(app.ParamsKeeper))
	govConfig := govtypes.DefaultConfig()
	/*
		Example of setting gov params:
		govConfig.MaxMetadataLen = 10000
	*/
	govKeeper := govkeeper.NewKeeper(
		appCodec, runtime.NewKVStoreService(keys[govtypes.StoreKey]), app.AccountKeeper, app.BankKeeper,
		app.StakingKeeper, app.DistrKeeper, app.MsgServiceRouter(), govConfig, authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	// Set legacy router for backwards compatibility with gov v1beta1
	govKeeper.SetLegacyRouter(govRouter)

	app.GovKeeper = *govKeeper.SetHooks(
		govtypes.NewMultiGovHooks(
		// register the governance hooks
		),
	)

	app.CapabilityKeeper = capabilitykeeper.NewKeeper(
		appCodec,
		keys[capabilitytypes.StoreKey],
		memKeys[capabilitytypes.MemStoreKey],
	)

	app.ScopedIBCKeeper = app.CapabilityKeeper.ScopeToModule(ibcexported.ModuleName)
	app.IBCKeeper = ibckeeper.NewKeeper(
		app.appCodec,
		app.keys[ibcexported.StoreKey],
		app.GetSubspace(ibcexported.ModuleName),
		app.StakingKeeper,
		app.UpgradeKeeper,
		app.ScopedIBCKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	// TODO: update late for oralce pr
	app.ScopedOracleKeeper = app.CapabilityKeeper.ScopeToModule(oracletypes.ModuleName)
	app.OracleKeeper = oraclekeeper.NewKeeper(
		app.appCodec,
		runtime.NewKVStoreService(keys[oracletypes.StoreKey]),
		app.Logger(),
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		nil,
		nil,
	)

	app.IBCFeeKeeper = ibcfeekeeper.NewKeeper(
		app.appCodec, app.GetKey(ibcfeetypes.StoreKey),
		app.IBCKeeper.ChannelKeeper, // may be replaced with IBC middleware
		app.IBCKeeper.ChannelKeeper,
		app.IBCKeeper.PortKeeper, app.AccountKeeper, app.BankKeeper,
	)

	app.ScopedIBCTransferKeeper = app.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)
	app.TransferKeeper = ibctransferkeeper.NewKeeper(
		app.appCodec,
		app.GetKey(ibctransfertypes.StoreKey),
		app.GetSubspace(ibctransfertypes.ModuleName),
		app.IBCFeeKeeper,
		app.IBCKeeper.ChannelKeeper,
		app.IBCKeeper.PortKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		app.ScopedIBCTransferKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	app.ScopedICAHostKeeper = app.CapabilityKeeper.ScopeToModule(icahosttypes.SubModuleName)
	app.ICAHostKeeper = icahostkeeper.NewKeeper(
		app.appCodec,
		app.GetKey(icahosttypes.StoreKey),
		app.GetSubspace(icahosttypes.SubModuleName),
		app.IBCFeeKeeper, // use ics29 fee as ics4Wrapper in middleware stack
		app.IBCKeeper.ChannelKeeper,
		app.IBCKeeper.PortKeeper,
		app.AccountKeeper,
		app.ScopedICAHostKeeper,
		app.MsgServiceRouter(),
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	// Create Transfer Stack
	// SendPacket, since it is originating from the application to core IBC:
	// transferKeeper.SendPacket -> fee.SendPacket -> channel.SendPacket

	// RecvPacket, message that originates from core IBC and goes down to app, the flow is the other way
	// channel.RecvPacket -> fee.OnRecvPacket -> transfer.OnRecvPacket

	// transfer stack contains (from top to bottom):
	// - IBC Fee Middleware
	// - Transfer

	// create IBC module from bottom to top of stack
	var transferStack porttypes.IBCModule
	transferStack = ibctransfer.NewIBCModule(app.TransferKeeper)
	transferStack = ibcfee.NewIBCMiddleware(transferStack, app.IBCFeeKeeper)

	// Create Interchain Accounts Stack
	// SendPacket, since it is originating from the application to core IBC:
	// icaControllerKeeper.SendTx -> fee.SendPacket -> channel.SendPacket

	// RecvPacket, message that originates from core IBC and goes down to app, the flow is:
	// channel.RecvPacket -> fee.OnRecvPacket -> icaHost.OnRecvPacket
	var icaHostStack porttypes.IBCModule
	icaHostStack = icahost.NewIBCModule(app.ICAHostKeeper)
	icaHostStack = ibcfee.NewIBCMiddleware(icaHostStack, app.IBCFeeKeeper)

	// Create Oracle ibc module
	oracleModule := oraclemodule.NewIBCModule(app.OracleKeeper)
	
	// Create static IBC router, add ibctransfer route, then set and seal it
	ibcRouter := porttypes.NewRouter().
		AddRoute(icahosttypes.SubModuleName, icaHostStack).
		AddRoute(ibctransfertypes.ModuleName, transferStack).
		AddRoute(oracletypes.ModuleName, oracleModule)

	// Setting Router will finalize all routes by sealing router
	// No more routes can be added
	app.IBCKeeper.SetRouter(ibcRouter)

	// Sealing prevents other modules from creating scoped sub-keepers
	app.CapabilityKeeper.Seal()

	/****  Module Options ****/

	// NOTE: we may consider parsing `appOpts` inside module constructors. For the moment
	// we prefer to be more strict in what arguments the modules expect.
	skipGenesisInvariants := cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.
	app.ModuleManager = module.NewManager(
		// SDK app modules
		genutil.NewAppModule(
			app.AccountKeeper, app.StakingKeeper, app,
			txConfig,
		),
		auth.NewAppModule(appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts, app.GetSubspace(authtypes.ModuleName)),
		vesting.NewAppModule(app.AccountKeeper, app.BankKeeper),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper, app.GetSubspace(banktypes.ModuleName)),
		crisis.NewAppModule(app.CrisisKeeper, skipGenesisInvariants, app.GetSubspace(crisistypes.ModuleName)),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		gov.NewAppModule(appCodec, &app.GovKeeper, app.AccountKeeper, app.BankKeeper, app.GetSubspace(govtypes.ModuleName)),
		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper, nil, app.GetSubspace(minttypes.ModuleName)),
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper, app.GetSubspace(slashingtypes.ModuleName), app.interfaceRegistry),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper, app.GetSubspace(distrtypes.ModuleName)),
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper, app.GetSubspace(stakingtypes.ModuleName)),
		upgrade.NewAppModule(app.UpgradeKeeper, app.AccountKeeper.AddressCodec()),
		evidence.NewAppModule(app.EvidenceKeeper),
		params.NewAppModule(app.ParamsKeeper),
		authzmodule.NewAppModule(appCodec, app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		groupmodule.NewAppModule(appCodec, app.GroupKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		consensus.NewAppModule(appCodec, app.ConsensusParamsKeeper),
		circuit.NewAppModule(appCodec, app.CircuitKeeper),
		ibc.NewAppModule(app.IBCKeeper),
		ibctransfer.NewAppModule(app.TransferKeeper),
		ibcfee.NewAppModule(app.IBCFeeKeeper),
		icamodule.NewAppModule(&app.ICAControllerKeeper, &app.ICAHostKeeper),
		ibctm.NewAppModule(),
		icamodule.NewAppModule(nil, &app.ICAHostKeeper),

		// reserve app modules
		oraclemodule.NewAppModule(appCodec, app.OracleKeeper, app.AccountKeeper, app.BankKeeper),
		
	)

	// BasicModuleManager defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration and genesis verification.
	// By default it is composed of all the module from the module manager.
	// Additionally, app module basics can be overwritten by passing them as argument.
	app.BasicModuleManager = module.NewBasicManagerFromManager(
		app.ModuleManager,
		map[string]module.AppModuleBasic{
			genutiltypes.ModuleName: genutil.NewAppModuleBasic(genutiltypes.DefaultMessageValidator),
			govtypes.ModuleName: gov.NewAppModuleBasic(
				[]govclient.ProposalHandler{
					paramsclient.ProposalHandler,
				},
			),
		})
	app.BasicModuleManager.RegisterLegacyAminoCodec(legacyAmino)
	app.BasicModuleManager.RegisterInterfaces(interfaceRegistry)

	// NOTE: upgrade module is required to be prioritized
	app.ModuleManager.SetOrderPreBlockers(
		upgradetypes.ModuleName,
	)
	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	app.ModuleManager.SetOrderBeginBlockers(beginBlockerOrder()...)
	app.ModuleManager.SetOrderEndBlockers(endBlockerOrder()...)

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: The genutils module must also occur after auth so that it can access the params from auth.
	genesisModuleOrder := []string{
		authtypes.ModuleName, banktypes.ModuleName,
		distrtypes.ModuleName, stakingtypes.ModuleName, slashingtypes.ModuleName, govtypes.ModuleName,
		minttypes.ModuleName, crisistypes.ModuleName, genutiltypes.ModuleName, evidencetypes.ModuleName, authz.ModuleName,
		feegrant.ModuleName, group.ModuleName, paramstypes.ModuleName, upgradetypes.ModuleName,
		vestingtypes.ModuleName, consensustypes.ModuleName, circuittypes.ModuleName, oracletypes.ModuleName,
	}
	app.ModuleManager.SetOrderInitGenesis(genesisModuleOrder...)
	app.ModuleManager.SetOrderExportGenesis(genesisModuleOrder...)

	// Uncomment if you want to set a custom migration order here.
	// app.ModuleManager.SetOrderMigrations(custom order)

	app.ModuleManager.RegisterInvariants(app.CrisisKeeper)
	app.configurator = module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	err = app.ModuleManager.RegisterServices(app.configurator)
	if err != nil {
		panic(err)
	}

	// RegisterUpgradeHandlers is used for registering any on-chain upgrades.
	// Make sure it's called after `app.ModuleManager` and `app.configurator` are set.
	// app.RegisterUpgradeHandlers()

	autocliv1.RegisterQueryServer(app.GRPCQueryRouter(), runtimeservices.NewAutoCLIQueryService(app.ModuleManager.Modules))

	reflectionSvc, err := runtimeservices.NewReflectionService()
	if err != nil {
		panic(err)
	}
	reflectionv1.RegisterReflectionServiceServer(app.GRPCQueryRouter(), reflectionSvc)

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(app.memKeys)

	// initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetPreBlocker(app.PreBlocker)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)
	// app.setAnteHandler(txConfig)

	// In v0.46, the SDK introduces _postHandlers_. PostHandlers are like
	// antehandlers, but are run _after_ the `runMsgs` execution. They are also
	// defined as a chain, and have the same signature as antehandlers.
	//
	// In baseapp, postHandlers are run in the same store branch as `runMsgs`,
	// meaning that both `runMsgs` and `postHandler` state will be committed if
	// both are successful, and both will be reverted if any of the two fails.
	//
	// The SDK exposes a default postHandlers chain
	//
	// Please note that changing any of the anteHandler or postHandler chain is
	// likely to be a state-machine breaking change, which needs a coordinated
	// upgrade.
	// app.setPostHandler()

	// At startup, after all modules have been registered, check that all prot
	// annotations are correct.
	// protoFiles, err := proto.MergedRegistry()
	// if err != nil {
	// 	panic(err)
	// }
	// err = msgservice.ValidateProtoAnnotations(protoFiles)
	// if err != nil {
	// 	// Once we switch to using protoreflect-based antehandlers, we might
	// 	// want to panic here instead of logging a warning.
	// 	fmt.Fprintln(os.Stderr, err.Error())
	// }

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			panic(fmt.Errorf("error loading last version: %w", err))
		}
	}

	return app
}

// func (app *App) setAnteHandler(txConfig client.TxConfig) {
// 	anteHandler, err := NewAnteHandler(
// 		HandlerOptions{
// 			ante.HandlerOptions{
// 				AccountKeeper:   app.AccountKeeper,
// 				BankKeeper:      app.BankKeeper,
// 				SignModeHandler: txConfig.SignModeHandler(),
// 				FeegrantKeeper:  app.FeeGrantKeeper,
// 				SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
// 			},
// 			&app.CircuitKeeper,
// 		},
// 	)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Set the AnteHandler for the app
// 	app.SetAnteHandler(anteHandler)
// }

// PreBlocker application updates every pre block
func (app *App) PreBlocker(ctx sdk.Context, _ *abci.RequestFinalizeBlock) (*sdk.ResponsePreBlock, error) {
	return app.ModuleManager.PreBlock(ctx)
}

// BeginBlocker updates every begin block
func (app *App) BeginBlocker(ctx sdk.Context) (sdk.BeginBlock, error) {
	return app.ModuleManager.BeginBlock(ctx)
}

// EndBlocker application updates every end block
func (app *App) EndBlocker(ctx sdk.Context) (sdk.EndBlock, error) {
	return app.ModuleManager.EndBlock(ctx)
}

// InitChainer application update at chain initialization
func (app *App) InitChainer(ctx sdk.Context, req *abci.RequestInitChain) (*abci.ResponseInitChain, error) {
	var genesisState GenesisState
	app.legacyAmino.MustUnmarshalJSON(req.AppStateBytes, &genesisState)
	if err := app.UpgradeKeeper.SetModuleVersionMap(ctx, app.ModuleManager.GetVersionMap()); err != nil {
		return nil, err
	}

	return app.ModuleManager.InitGenesis(ctx, app.appCodec, genesisState)
}

func (app *App) RegisterNodeService(clientCtx client.Context, cfg config.Config) {
	nodeservice.RegisterNodeService(clientCtx, app.GRPCQueryRouter(), cfg)
}

// LoadHeight loads a particular height
func (app *App) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

func beginBlockerOrder() []string {
	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: upgrade module must go first to handle software upgrades.
	// NOTE: staking module is required if HistoricalEntries param > 0.
	return []string{
		minttypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		evidencetypes.ModuleName,
		stakingtypes.ModuleName,
		genutiltypes.ModuleName,
		vestingtypes.ModuleName,
		govtypes.ModuleName,
		paramstypes.ModuleName,
		authtypes.ModuleName,
		crisistypes.ModuleName,
		feegrant.ModuleName,
		banktypes.ModuleName,
		authz.ModuleName,
		ibctransfertypes.ModuleName,
		consensustypes.ModuleName,
		capabilitytypes.ModuleName,
		ibcexported.ModuleName,
		icatypes.ModuleName,
		ibcfeetypes.ModuleName,
		oracletypes.ModuleName,
	}
}

func endBlockerOrder() []string {
	return []string{
		crisistypes.ModuleName,
		govtypes.ModuleName,
		stakingtypes.ModuleName,
		genutiltypes.ModuleName,
		vestingtypes.ModuleName,
		paramstypes.ModuleName,
		authtypes.ModuleName,
		feegrant.ModuleName,
		group.ModuleName,
		authz.ModuleName,
		ibctransfertypes.ModuleName,
		consensustypes.ModuleName,
		oracletypes.ModuleName,
		minttypes.ModuleName,
		slashingtypes.ModuleName,
		ibctransfertypes.ModuleName,
		evidencetypes.ModuleName,
		capabilitytypes.ModuleName,
		distrtypes.ModuleName,
		ibcexported.ModuleName,
		icatypes.ModuleName,
		ibcfeetypes.ModuleName,
		upgradetypes.ModuleName,
		banktypes.ModuleName,
	}
}

// LegacyAmino returns App's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) LegacyAmino() *codec.LegacyAmino {
	return app.legacyAmino
}

// AppCodec returns App's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) AppCodec() codec.Codec {
	return app.appCodec
}

// GetKey returns the KVStoreKey for the provided store key.
func (app *App) GetKey(storeKey string) *storetypes.KVStoreKey {
	return app.keys[storeKey]
}

// GetMemKey returns the MemoryStoreKey for the provided store key.
func (app *App) GetMemKey(storeKey string) *storetypes.MemoryStoreKey {
	return app.memKeys[storeKey]
}

func (app *App) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

func (app *App) RegisterTendermintService(clientCtx client.Context) {
	cmtservice.RegisterTendermintService(clientCtx, app.BaseApp.GRPCQueryRouter(), app.interfaceRegistry, app.Query)
}

// GetSubspace returns a param subspace for a given module name.
func (app *App) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// GetIBCKeeper returns the IBC keeper.
func (app *App) GetIBCKeeper() *ibckeeper.Keeper {
	return app.IBCKeeper
}

// GetCapabilityScopedKeeper returns the capability scoped keeper.
func (app *App) GetCapabilityScopedKeeper(moduleName string) capabilitykeeper.ScopedKeeper {
	return app.CapabilityKeeper.ScopeToModule(moduleName)
}

// SimulationManager implements the SimulationApp interface.
func (app *App) SimulationManager() *module.SimulationManager {
	return app.sm
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *App) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx
	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register new CometBFT queries routes from grpc-gateway.
	cmtservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register node gRPC service for grpc-gateway.
	nodeservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register grpc-gateway routes for all modules.
	app.BasicModuleManager.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// register swagger API from root so that other applications can override easily
	if err := server.RegisterSwaggerAPI(apiSvr.ClientCtx, apiSvr.Router, apiConfig.Swagger); err != nil {
		panic(err)
	}
}

// GetMaccPerms returns a copy of the module account permissions
//
// NOTE: This is solely to be used for testing purposes.
func GetMaccPerms() map[string][]string {
	dup := make(map[string][]string)
	for _, perms := range moduleAccPerms {
		dup[perms.Account] = perms.Permissions
	}
	return dup
}

// BlockedAddresses returns all the app's blocked account addresses.
func BlockedAddresses() map[string]bool {
	result := make(map[string]bool)
	if len(blockAccAddrs) > 0 {
		for _, addr := range blockAccAddrs {
			result[addr] = true
		}
	} else {
		for addr := range GetMaccPerms() {
			result[addr] = true
		}
	}
	return result
}

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(
	appCodec codec.BinaryCodec,
	legacyAmino *codec.LegacyAmino,
	key, tkey storetypes.StoreKey,
) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	// SDK subspaces
	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(stakingtypes.ModuleName)
	paramsKeeper.Subspace(minttypes.ModuleName)
	paramsKeeper.Subspace(distrtypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName)
	paramsKeeper.Subspace(govtypes.ModuleName)
	paramsKeeper.Subspace(crisistypes.ModuleName)

	// register the key tables for legacy param subspaces
	keyTable := ibcclienttypes.ParamKeyTable()
	keyTable.RegisterParamSet(&ibcconnectiontypes.Params{})
	paramsKeeper.Subspace(ibcexported.ModuleName).WithKeyTable(keyTable)
	paramsKeeper.Subspace(ibctransfertypes.ModuleName).WithKeyTable(ibctransfertypes.ParamKeyTable())
	paramsKeeper.Subspace(icacontrollertypes.SubModuleName).WithKeyTable(icacontrollertypes.ParamKeyTable())
	paramsKeeper.Subspace(icahosttypes.SubModuleName).WithKeyTable(icahosttypes.ParamKeyTable())

	// Reserve subspaces
	paramsKeeper.Subspace(oracletypes.ModuleName)

	return paramsKeeper
}
