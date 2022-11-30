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

func createNFeeTier(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.FeeTier {
	items := make([]types.FeeTier, n)
	for i := range items {
		items[i].Id = keeper.AppendFeeTier(ctx, items[i])
	}
	return items
}

func TestFeeTierGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNFeeTier(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetFeeTier(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestFeeTierRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNFeeTier(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveFeeTier(ctx, item.Id)
		_, found := keeper.GetFeeTier(ctx, item.Id)
		require.False(t, found)
	}
}

func TestFeeTierGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNFeeTier(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllFeeTier(ctx)),
	)
}

func TestFeeTierCount(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNFeeTier(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetFeeTierCount(ctx))
}
