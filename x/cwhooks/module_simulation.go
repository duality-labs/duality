package cwhooks

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/duality-labs/duality/testutil/sample"
	cwhookssimulation "github.com/duality-labs/duality/x/cwhooks/simulation"
	"github.com/duality-labs/duality/x/cwhooks/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = cwhookssimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateHook = "op_weight_msg_hook"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateHook int = 100

	opWeightMsgDeleteHook = "op_weight_msg_hook"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteHook int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	cwhooksGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		HookList: []types.Hook{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		HookCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&cwhooksGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateHook int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateHook, &weightMsgCreateHook, nil,
		func(_ *rand.Rand) {
			weightMsgCreateHook = defaultWeightMsgCreateHook
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateHook,
		cwhookssimulation.SimulateMsgCreateHook(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteHook int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteHook, &weightMsgDeleteHook, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteHook = defaultWeightMsgDeleteHook
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteHook,
		cwhookssimulation.SimulateMsgDeleteHook(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateHook,
			defaultWeightMsgCreateHook,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cwhookssimulation.SimulateMsgCreateHook(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteHook,
			defaultWeightMsgDeleteHook,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cwhookssimulation.SimulateMsgDeleteHook(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
