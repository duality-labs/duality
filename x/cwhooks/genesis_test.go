package cwhooks_test

import (
	"testing"

	keepertest "github.com/duality-labs/duality/testutil/keeper"
	"github.com/duality-labs/duality/testutil/nullify"
	"github.com/duality-labs/duality/x/cwhooks"
	"github.com/duality-labs/duality/x/cwhooks/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		HookList: []types.Hook{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		HookCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CWHooksKeeper(t)
	cwhooks.InitGenesis(ctx, *k, genesisState)
	got := cwhooks.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.HookList, got.HookList)
	require.Equal(t, genesisState.HookCount, got.HookCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
