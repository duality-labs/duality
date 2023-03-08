package incentives_test

import (
	"testing"

	keepertest "github.com/duality-labs/duality/testutil/keeper"
	"github.com/duality-labs/duality/testutil/nullify"
	"github.com/duality-labs/duality/x/incentives"
	"github.com/duality-labs/duality/x/incentives/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		IncentivePlanList: []types.IncentivePlan{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		UserStakeList: []types.UserStake{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IncentivesKeeper(t)
	incentives.InitGenesis(ctx, *k, genesisState)
	got := incentives.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.IncentivePlanList, got.IncentivePlanList)
	require.ElementsMatch(t, genesisState.UserStakeList, got.UserStakeList)
	// this line is used by starport scaffolding # genesis/test/assert
}
