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
				Price:     "0",
				Fee:       "0",
				OrderType: "0",
			},
			{
				Price:     "1",
				Fee:       "1",
				OrderType: "1",
			},
		},
		PairsList: []types.Pairs{
			{
				Token0: "0",
				Token1: "0",
			},
			{
				Token0: "1",
				Token1: "1",
			},
		},
		IndexQueueList: []types.IndexQueue{
			{
				Index: 0,
			},
			{
				Index: 1,
			},
		},
		NodesList: []types.Nodes{
			{
				Node: "0",
			},
			{
				Node: "1",
			},
		},
		SharesList: []types.Shares{
			{
				Address:   "0",
				Price:     "0",
				Fee:       "0",
				OrderType: "0",
			},
			{
				Address:   "1",
				Price:     "1",
				Fee:       "1",
				OrderType: "1",
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
	require.ElementsMatch(t, genesisState.PairsList, got.PairsList)
	require.ElementsMatch(t, genesisState.IndexQueueList, got.IndexQueueList)
	require.ElementsMatch(t, genesisState.NodesList, got.NodesList)
	require.ElementsMatch(t, genesisState.SharesList, got.SharesList)
	// this line is used by starport scaffolding # genesis/test/assert
}
