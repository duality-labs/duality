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

func createNAdjMatrix(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.AdjMatrix {
	items := make([]types.AdjMatrix, n)
	for i := range items {
		items[i].Id = keeper.AppendAdjMatrix(ctx, items[i])
	}
	return items
}

func TestAdjMatrixGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNAdjMatrix(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetAdjMatrix(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestAdjMatrixRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNAdjMatrix(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAdjMatrix(ctx, item.Id)
		_, found := keeper.GetAdjMatrix(ctx, item.Id)
		require.False(t, found)
	}
}

func TestAdjMatrixGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNAdjMatrix(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAdjMatrix(ctx)),
	)
}

func TestAdjMatrixCount(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNAdjMatrix(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetAdjMatrixCount(ctx))
}
