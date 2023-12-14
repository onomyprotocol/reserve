package reserve

import (
	"math/rand"

	"reserve/testutil/sample"
	reservesimulation "reserve/x/reserve/simulation"
	"reserve/x/reserve/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = reservesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateVault = "op_weight_msg_create_vault"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateVault int = 100

	opWeightMsgDepositCollateral = "op_weight_msg_deposit_collateral"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDepositCollateral int = 100

	opWeightMsgMintDenom = "op_weight_msg_mint_denom"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMintDenom int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	reserveGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&reserveGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateVault int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateVault, &weightMsgCreateVault, nil,
		func(_ *rand.Rand) {
			weightMsgCreateVault = defaultWeightMsgCreateVault
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateVault,
		reservesimulation.SimulateMsgCreateVault(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDepositCollateral int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDepositCollateral, &weightMsgDepositCollateral, nil,
		func(_ *rand.Rand) {
			weightMsgDepositCollateral = defaultWeightMsgDepositCollateral
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDepositCollateral,
		reservesimulation.SimulateMsgDepositCollateral(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgMintDenom int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgMintDenom, &weightMsgMintDenom, nil,
		func(_ *rand.Rand) {
			weightMsgMintDenom = defaultWeightMsgMintDenom
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMintDenom,
		reservesimulation.SimulateMsgMintDenom(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
