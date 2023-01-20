package dex_test

import (
	"testing"

	keepertest "github.com/duality-labs/duality/testutil/keeper"
	"github.com/duality-labs/duality/testutil/nullify"
	"github.com/duality-labs/duality/x/dex"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/stretchr/testify/require"
)

var defaultPairId *types.PairId = &types.PairId{Token0: "TokenA", Token1: "TokenB"}

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		TickList: []types.Tick{
			{
				TickIndex: 0,
				PairId:    defaultPairId,
			},
			{
				TickIndex: 1,
				PairId:    defaultPairId,
			},
		},
		TradingPairList: []types.TradingPair{
			{
				PairId: defaultPairId,
			},
			{
				PairId: defaultPairId,
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
				PairId:    defaultPairId,
			},
			{
				TickIndex: 1,
				PairId:    defaultPairId,
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
