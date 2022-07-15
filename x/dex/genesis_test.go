package dex_test

import (
	"testing"

	keepertest "github.com/NicholasDotSol/duality/testutil/keeper"
	"github.com/NicholasDotSol/duality/testutil/nullify"
	"github.com/NicholasDotSol/duality/x/dex"
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		TicksList: []types.Ticks{
			{
				Token0: "0",
				Token1: "0",
			},
			{
				Token0: "1",
				Token1: "1",
			},
		},
		ShareList: []types.Share{
			{
				Owner:  "0",
				Token0: "0",
				Token1: "0",
				Price:  "0",
				Fee:    "0",
			},
			{
				Owner:  "1",
				Token0: "1",
				Token1: "1",
				Price:  "1",
				Fee:    "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DexKeeper(t)
	dex.InitGenesis(ctx, *k, genesisState)
	got := dex.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.TicksList, got.TicksList)
	require.ElementsMatch(t, genesisState.ShareList, got.ShareList)
	// this line is used by starport scaffolding # genesis/test/assert
}
