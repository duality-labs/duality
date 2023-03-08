package incentives

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/duality-labs/duality/testutil/sample"
	incentivessimulation "github.com/duality-labs/duality/x/incentives/simulation"
	"github.com/duality-labs/duality/x/incentives/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = incentivessimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateIncentivePlan = "op_weight_msg_incentive_plan"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateIncentivePlan int = 100

	opWeightMsgUpdateIncentivePlan = "op_weight_msg_incentive_plan"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateIncentivePlan int = 100

	opWeightMsgDeleteIncentivePlan = "op_weight_msg_incentive_plan"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteIncentivePlan int = 100

	opWeightMsgCreateUserStake = "op_weight_msg_user_stake"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateUserStake int = 100

	opWeightMsgUpdateUserStake = "op_weight_msg_user_stake"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateUserStake int = 100

	opWeightMsgDeleteUserStake = "op_weight_msg_user_stake"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteUserStake int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	incentivesGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		IncentivePlanList: []types.IncentivePlan{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		UserStakeList: []types.UserStake{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&incentivesGenesis)
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

	var weightMsgCreateIncentivePlan int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateIncentivePlan, &weightMsgCreateIncentivePlan, nil,
		func(_ *rand.Rand) {
			weightMsgCreateIncentivePlan = defaultWeightMsgCreateIncentivePlan
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateIncentivePlan,
		incentivessimulation.SimulateMsgCreateIncentivePlan(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateIncentivePlan int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateIncentivePlan, &weightMsgUpdateIncentivePlan, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateIncentivePlan = defaultWeightMsgUpdateIncentivePlan
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateIncentivePlan,
		incentivessimulation.SimulateMsgUpdateIncentivePlan(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteIncentivePlan int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteIncentivePlan, &weightMsgDeleteIncentivePlan, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteIncentivePlan = defaultWeightMsgDeleteIncentivePlan
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteIncentivePlan,
		incentivessimulation.SimulateMsgDeleteIncentivePlan(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateUserStake int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateUserStake, &weightMsgCreateUserStake, nil,
		func(_ *rand.Rand) {
			weightMsgCreateUserStake = defaultWeightMsgCreateUserStake
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateUserStake,
		incentivessimulation.SimulateMsgCreateUserStake(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateUserStake int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateUserStake, &weightMsgUpdateUserStake, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateUserStake = defaultWeightMsgUpdateUserStake
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateUserStake,
		incentivessimulation.SimulateMsgUpdateUserStake(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteUserStake int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteUserStake, &weightMsgDeleteUserStake, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteUserStake = defaultWeightMsgDeleteUserStake
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteUserStake,
		incentivessimulation.SimulateMsgDeleteUserStake(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
