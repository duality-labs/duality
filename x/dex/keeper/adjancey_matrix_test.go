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

func createNAdjanceyMatrix(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.AdjanceyMatrix {
	items := make([]types.AdjanceyMatrix, n)
	for i := range items {
		items[i].Id = keeper.AppendAdjanceyMatrix(ctx, items[i])
	}
	return items
}

func TestAdjanceyMatrixGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNAdjanceyMatrix(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetAdjanceyMatrix(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestAdjanceyMatrixRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNAdjanceyMatrix(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAdjanceyMatrix(ctx, item.Id)
		_, found := keeper.GetAdjanceyMatrix(ctx, item.Id)
		require.False(t, found)
	}
}

func TestAdjanceyMatrixGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNAdjanceyMatrix(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAdjanceyMatrix(ctx)),
	)
}

func TestAdjanceyMatrixCount(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNAdjanceyMatrix(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetAdjanceyMatrixCount(ctx))
}
