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

func createNNodes(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Nodes {
	items := make([]types.Nodes, n)
	for i := range items {
		items[i].Node = strconv.Itoa(i)

		keeper.SetNodes(ctx, items[i])
	}
	return items
}

func TestNodesGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNNodes(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNodes(ctx,
			item.Node,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestNodesRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNNodes(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNodes(ctx,
			item.Node,
		)
		_, found := keeper.GetNodes(ctx,
			item.Node,
		)
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
