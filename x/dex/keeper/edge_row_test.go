package keeper_test

import (
	"testing"

	keepertest "github.com/NicholasDotSol/duality/testutil/keeper"
	"github.com/NicholasDotSol/duality/testutil/nullify"
	"github.com/NicholasDotSol/duality/x/dex/keeper"
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNEdgeRow(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.EdgeRow {
	items := make([]types.EdgeRow, n)
	for i := range items {
		items[i].Id = keeper.AppendEdgeRow(ctx, items[i])
	}
	return items
}

func TestEdgeRowGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNEdgeRow(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetEdgeRow(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestEdgeRowRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNEdgeRow(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveEdgeRow(ctx, item.Id)
		_, found := keeper.GetEdgeRow(ctx, item.Id)
		require.False(t, found)
	}
}

func TestEdgeRowGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNEdgeRow(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllEdgeRow(ctx)),
	)
}

func TestEdgeRowCount(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNEdgeRow(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetEdgeRowCount(ctx))
}
