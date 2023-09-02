package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/duality-labs/duality/testutil/keeper"
	"github.com/duality-labs/duality/testutil/nullify"
	"github.com/duality-labs/duality/x/cwhooks/keeper"
	"github.com/duality-labs/duality/x/cwhooks/types"
	"github.com/stretchr/testify/require"
)

func createNHook(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Hook {
	items := make([]types.Hook, n)
	for i := range items {
		items[i].Id = keeper.AppendHook(ctx, items[i])
	}
	return items
}

func TestHookGet(t *testing.T) {
	keeper, ctx := keepertest.CwhooksKeeper(t)
	items := createNHook(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetHook(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestHookRemove(t *testing.T) {
	keeper, ctx := keepertest.CwhooksKeeper(t)
	items := createNHook(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveHook(ctx, item.Id)
		_, found := keeper.GetHook(ctx, item.Id)
		require.False(t, found)
	}
}

func TestHookGetAll(t *testing.T) {
	keeper, ctx := keepertest.CwhooksKeeper(t)
	items := createNHook(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllHook(ctx)),
	)
}

func TestHookCount(t *testing.T) {
	keeper, ctx := keepertest.CwhooksKeeper(t)
	items := createNHook(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetHookCount(ctx))
}
