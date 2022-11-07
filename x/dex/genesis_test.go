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

		TickMapList: []types.TickMap{
			{
				TickIndex: 0,
			},
			{
				TickIndex: 1,
			},
		},
		PairMapList: []types.PairMap{
			{
				PairId: "0",
			},
			{
				PairId: "1",
			},
		},
		TokensList: []types.Tokens{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		TokensCount: 2,
		SharesList: []types.Shares{
			{
				Address:   "0",
				PairId:    "0",
				TickIndex: 0,
				FeeIndex:  0,
			},
			{
				Address:   "1",
				PairId:    "1",
				TickIndex: 1,
				FeeIndex:  uint64(1),
			},
		},
		FeeListList: []types.FeeList{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		FeeListCount: 2,

		LimitOrderPoolUserShareMapList: []types.LimitOrderPoolUserShareMap{
			{
				Count:   0,
				Address: "0",
			},
			{
				Count:   1,
				Address: "1",
			},
		},
		LimitOrderPoolUserSharesWithdrawnList: []types.LimitOrderPoolUserSharesWithdrawn{
			{
				Count:   0,
				Address: "0",
			},
			{
				Count:   1,
				Address: "1",
			},
		},
		LimitOrderPoolTotalSharesMapList: []types.LimitOrderPoolTotalSharesMap{
			{
				Count: 0,
			},
			{
				Count: 1,
			},
		},
		LimitOrderPoolReserveMapList: []types.LimitOrderPoolReserveMap{
			{
				Count: 0,
			},
			{
				Count: 1,
			},
		},
		LimitOrderPoolFillMapList: []types.LimitOrderPoolFillMap{
			{
				Count: 0,
			},
			{
				Count: 1,
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

	require.ElementsMatch(t, genesisState.TickMapList, got.TickMapList)
	require.ElementsMatch(t, genesisState.PairMapList, got.PairMapList)
	require.ElementsMatch(t, genesisState.TokensList, got.TokensList)
	require.Equal(t, genesisState.TokensCount, got.TokensCount)
	require.ElementsMatch(t, genesisState.SharesList, got.SharesList)
	require.ElementsMatch(t, genesisState.FeeListList, got.FeeListList)
	require.Equal(t, genesisState.FeeListCount, got.FeeListCount)
	require.ElementsMatch(t, genesisState.LimitOrderPoolUserShareMapList, got.LimitOrderPoolUserShareMapList)
	require.ElementsMatch(t, genesisState.LimitOrderPoolUserSharesWithdrawnList, got.LimitOrderPoolUserSharesWithdrawnList)
	require.ElementsMatch(t, genesisState.LimitOrderPoolTotalSharesMapList, got.LimitOrderPoolTotalSharesMapList)
	require.ElementsMatch(t, genesisState.LimitOrderPoolReserveMapList, got.LimitOrderPoolReserveMapList)
	require.ElementsMatch(t, genesisState.LimitOrderPoolFillMapList, got.LimitOrderPoolFillMapList)
	// this line is used by starport scaffolding # genesis/test/assert
}
