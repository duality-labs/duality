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

		TickList: []types.Tick{
			{
				TickIndex: 0,
			},
			{
				TickIndex: 1,
			},
		},
		TradingPairList: []types.TradingPair{
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
		FeeTierList: []types.FeeTier{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		FeeTierCount: 2,
		LimitOrderTrancheUserList: []types.LimitOrderTrancheUser{
			{
				Count:   0,
				Address: "0",
			},
			{
				Count:   1,
				Address: "1",
			},
		},
		LimitOrderTrancheList: []types.LimitOrderTranche{
			{
				TickIndex: 0,
			},
			{
				TickIndex: 1,
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

	require.ElementsMatch(t, genesisState.TickList, got.TickList)
	require.ElementsMatch(t, genesisState.TradingPairList, got.TradingPairList)
	require.ElementsMatch(t, genesisState.TokensList, got.TokensList)
	require.Equal(t, genesisState.TokensCount, got.TokensCount)
	require.ElementsMatch(t, genesisState.FeeTierList, got.FeeTierList)
	require.Equal(t, genesisState.FeeTierCount, got.FeeTierCount)
	require.ElementsMatch(t, genesisState.LimitOrderTrancheUserList, got.LimitOrderTrancheUserList)
	require.ElementsMatch(t, genesisState.LimitOrderTrancheList, got.LimitOrderTrancheList)
	// this line is used by starport scaffolding # genesis/test/assert
}
