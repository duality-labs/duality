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

func createNNodes(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Nodes {
	items := make([]types.Nodes, n)
	for i := range items {
		items[i].Id = keeper.AppendNodes(ctx, items[i])
	}
	return items
}

func TestNodesGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNNodes(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetNodes(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestNodesRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNNodes(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNodes(ctx, item.Id)
		_, found := keeper.GetNodes(ctx, item.Id)
		require.False(t, found)
	}
}

func TestNodesGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNNodes(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNodes(ctx)),
	)
}

func TestNodesCount(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNNodes(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetNodesCount(ctx))
}
