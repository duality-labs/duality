package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/NicholasDotSol/duality/testutil/keeper"
	"github.com/NicholasDotSol/duality/testutil/nullify"
	"github.com/NicholasDotSol/duality/x/dex/keeper"
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNTradingPair(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.TradingPair {
	items := make([]types.TradingPair, n)
	for i := range items {
		items[i].PairId = strconv.Itoa(i)

		keeper.SetTradingPair(ctx, items[i])
	}
	return items
}

func TestTradingPairGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTradingPair(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetTradingPair(ctx,
			item.PairId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTradingPairRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTradingPair(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTradingPair(ctx,
			item.PairId,
		)
		_, found := keeper.GetTradingPair(ctx,
			item.PairId,
		)
		require.False(t, found)
	}
}

func TestTradingPairGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTradingPair(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTradingPair(ctx)),
	)
}
