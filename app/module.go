package app

import (
	icatypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/types"

	ibcfeetypes "github.com/cosmos/ibc-go/v8/modules/apps/29-fee/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/cosmos/ibc-go/modules/capability"

	ibcfee "github.com/cosmos/ibc-go/v8/modules/apps/29-fee"

	// ibc "github.com/cosmos/ibc-go/v8/modules/core"

	ibctm "github.com/cosmos/ibc-go/v8/modules/light-clients/07-tendermint"

	"cosmossdk.io/x/evidence"

	feegrantmodule "cosmossdk.io/x/feegrant/module"
	"cosmossdk.io/x/upgrade"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"

	// "github.com/cosmos/cosmos-sdk/x/staking"

	"github.com/cosmos/cosmos-sdk/x/auth/vesting"

	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/consensus"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"

	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"

	"github.com/cosmos/cosmos-sdk/x/gov"

	"github.com/cosmos/cosmos-sdk/x/mint"

	sdkparams "github.com/cosmos/cosmos-sdk/x/params"

	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"

	evidencetypes "cosmossdk.io/x/evidence/types"
	"cosmossdk.io/x/feegrant"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/client"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	authz "github.com/cosmos/cosmos-sdk/x/authz"
	genutil "github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	ibc "github.com/cosmos/ibc-go/v8/modules/core"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"

	auction "github.com/onomyprotocol/reserve/x/auction/module"
	auctiontypes "github.com/onomyprotocol/reserve/x/auction/types"
	oracle "github.com/onomyprotocol/reserve/x/oracle/module"
	oracletypes "github.com/onomyprotocol/reserve/x/oracle/types"
	psm "github.com/onomyprotocol/reserve/x/psm/module"
	psmtypes "github.com/onomyprotocol/reserve/x/psm/types"
	vaults "github.com/onomyprotocol/reserve/x/vaults/module"
	vaultstypes "github.com/onomyprotocol/reserve/x/vaults/types"
)

var maccPerms = map[string][]string{
	authtypes.FeeCollectorName:     nil,
	distrtypes.ModuleName:          nil,
	icatypes.ModuleName:            nil,
	minttypes.ModuleName:           {authtypes.Minter},
	stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
	stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
	govtypes.ModuleName:            {authtypes.Burner},
	// liquiditytypes.ModuleName:         {authtypes.Minter, authtypes.Burner},
	ibctransfertypes.ModuleName: {authtypes.Minter, authtypes.Burner},
	ibcfeetypes.ModuleName:      nil,
	psmtypes.ModuleName:         {authtypes.Minter, authtypes.Burner},
	vaultstypes.ModuleName:      {authtypes.Minter, authtypes.Burner},
	auctiontypes.ModuleName:     {authtypes.Minter, authtypes.Burner},
}

func appModules(
	app *App,
	appCodec codec.Codec,
	txConfig client.TxEncodingConfig,
	skipGenesisInvariants bool,
) []module.AppModule {
	return []module.AppModule{
		genutil.NewAppModule(app.AccountKeeper, app.StakingKeeper, app, txConfig),
		auth.NewAppModule(appCodec, app.AccountKeeper, nil, app.GetSubspace(authtypes.ModuleName)),
		vesting.NewAppModule(app.AccountKeeper, app.BankKeeper),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper, app.GetSubspace(banktypes.ModuleName)),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper, false),
		crisis.NewAppModule(app.CrisisKeeper, skipGenesisInvariants, app.GetSubspace(crisistypes.ModuleName)),
		gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper, app.GetSubspace(govtypes.ModuleName)),
		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper, nil, app.GetSubspace(minttypes.ModuleName)),
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper, app.GetSubspace(slashingtypes.ModuleName), app.interfaceRegistry),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper, app.GetSubspace(distrtypes.ModuleName)),
		upgrade.NewAppModule(app.UpgradeKeeper, app.AccountKeeper.AddressCodec()),
		evidence.NewAppModule(app.EvidenceKeeper),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		authzmodule.NewAppModule(appCodec, app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		ibc.NewAppModule(app.IBCKeeper),
		ibcfee.NewAppModule(app.IBCFeeKeeper),
		ibctm.NewAppModule(),
		sdkparams.NewAppModule(app.ParamsKeeper),
		consensus.NewAppModule(appCodec, app.ConsensusParamsKeeper),
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper, app.GetSubspace(stakingtypes.ModuleName)),
		genutil.NewAppModule(app.AccountKeeper, app.StakingKeeper, app, txConfig),
		psm.NewAppModule(appCodec, app.PSMKeeper, app.AccountKeeper, app.BankKeeper),
		oracle.NewAppModule(appCodec, app.OracleKeeper, app.AccountKeeper, app.BankKeeper),
		vaults.NewAppModule(appCodec, app.VaultsKeeper, app.AccountKeeper, app.BankKeeper),
		auction.NewAppModule(appCodec, app.AuctionKeeper, app.AccountKeeper, app.BankKeeper),
	}

}

func newBasicManagerFromManager(app *App) module.BasicManager {
	basicManager := module.NewBasicManagerFromManager(
		app.mm,
		map[string]module.AppModuleBasic{
			genutiltypes.ModuleName: genutil.NewAppModuleBasic(genutiltypes.DefaultMessageValidator),
			govtypes.ModuleName: gov.NewAppModuleBasic(
				[]govclient.ProposalHandler{
					paramsclient.ProposalHandler,
				},
			),
		})
	basicManager.RegisterLegacyAminoCodec(app.legacyAmino)
	basicManager.RegisterInterfaces(app.interfaceRegistry)
	return basicManager
}

// During begin block slashing happens after distr.BeginBlocker so that
// there is nothing left over in the validator fee pool, so as to keep the
// CanWithdrawInvariant invariant.
// NOTE: staking module is required if HistoricalEntries param > 0
// NOTE: capability module's beginblocker must come before any modules using capabilities (e.g. IBC)
func orderBeginBlockers() []string {
	return []string{
		// cosmos sdk modules
		minttypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		evidencetypes.ModuleName,
		stakingtypes.ModuleName,
		authz.ModuleName,
		genutiltypes.ModuleName,
		// ibc modules
		capabilitytypes.ModuleName,
		ibcexported.ModuleName,
		ibctransfertypes.ModuleName,
		icatypes.ModuleName,
		ibcfeetypes.ModuleName,
		// chain modules
		oracletypes.ModuleName,
		vaultstypes.ModuleName,
		psmtypes.ModuleName,
		auctiontypes.ModuleName,
	}
}

func orderEndBlockers() []string {
	return []string{
		// cosmos sdk modules
		crisistypes.ModuleName,
		govtypes.ModuleName,
		stakingtypes.ModuleName,
		feegrant.ModuleName,
		genutiltypes.ModuleName,
		// ibc modules
		ibcexported.ModuleName,
		ibctransfertypes.ModuleName,
		capabilitytypes.ModuleName,
		icatypes.ModuleName,
		ibcfeetypes.ModuleName,
		// chain modules
		oracletypes.ModuleName,
		vaultstypes.ModuleName,
		psmtypes.ModuleName,
		auctiontypes.ModuleName,
	}
}

// NOTE: The genutils module must occur after staking so that pools are
// properly initialized with tokens from genesis accounts.
// NOTE: The genutils module must also occur after auth so that it can access the params from auth.
// NOTE: Capability module must occur first so that it can initialize any capabilities
// so that other modules that want to create or claim capabilities afterwards in InitChain
// can do so safely.
func orderInitBlockers() []string {
	return []string{
		// cosmos-sdk/ibc modules
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		slashingtypes.ModuleName,
		govtypes.ModuleName,
		minttypes.ModuleName,
		crisistypes.ModuleName,
		ibcexported.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		authz.ModuleName,
		ibctransfertypes.ModuleName,
		icatypes.ModuleName,
		ibcfeetypes.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		vestingtypes.ModuleName,
		consensustypes.ModuleName,
		// chain modules
		oracletypes.ModuleName,
		vaultstypes.ModuleName,
		psmtypes.ModuleName,
		auctiontypes.ModuleName,
	}
}

func simulationModules(
	app *App,
	appCodec codec.Codec,
	_ bool,
) []module.AppModuleSimulation {
	return []module.AppModuleSimulation{
		auth.NewAppModule(appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts, app.GetSubspace(authtypes.ModuleName)),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper, app.GetSubspace(banktypes.ModuleName)),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper, false),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper, app.GetSubspace(govtypes.ModuleName)),
		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper, nil, app.GetSubspace(minttypes.ModuleName)),
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper, app.GetSubspace(stakingtypes.ModuleName)),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper, app.GetSubspace(distrtypes.ModuleName)),
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper, app.GetSubspace(slashingtypes.ModuleName), app.interfaceRegistry),
		sdkparams.NewAppModule(app.ParamsKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		authzmodule.NewAppModule(appCodec, app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		ibc.NewAppModule(app.IBCKeeper),
		psm.NewAppModule(appCodec, app.PSMKeeper, app.AccountKeeper, app.BankKeeper),
		// oracle.NewAppModule(appCodec,app.OracleKeeper, app.AccountKeeper, app.BankKeeper),
		// vaults.NewAppModule(appCodec, app.VaultsKeeper, app.AccountKeeper, app.BankKeeper),
		// auction.NewAppModule(appCodec, app.VaultsKeeper, app.AccountKeeper, app.BankKeeper),
	}
}
