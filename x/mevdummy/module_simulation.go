package mevdummy

import (
	"math/rand"

	"github.com/NicholasDotSol/duality/testutil/sample"
	mevdummysimulation "github.com/NicholasDotSol/duality/x/mevdummy/simulation"
	"github.com/NicholasDotSol/duality/x/mevdummy/types"
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
	_ = mevdummysimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgSend = "op_weight_msg_send"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSend int = 100

	opWeightMsgWithdrawFunds = "op_weight_msg_withdraw_funds"
	// TODO: Determine the simulation weight value
	defaultWeightMsgWithdrawFunds int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	mevdummyGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&mevdummyGenesis)
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

	var weightMsgSend int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSend, &weightMsgSend, nil,
		func(_ *rand.Rand) {
			weightMsgSend = defaultWeightMsgSend
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSend,
		mevdummysimulation.SimulateMsgSend(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgWithdrawFunds int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgWithdrawFunds, &weightMsgWithdrawFunds, nil,
		func(_ *rand.Rand) {
			weightMsgWithdrawFunds = defaultWeightMsgWithdrawFunds
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgWithdrawFunds,
		mevdummysimulation.SimulateMsgWithdrawFunds(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
