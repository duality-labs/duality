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

func createNVirtualPriceTickQueue(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.VirtualPriceTickQueue {
	items := make([]types.VirtualPriceTickQueue, n)
	for i := range items {
		items[i].Fee = sdk.ZeroDec()
		items[i].Price = sdk.ZeroDec()
		items[i].Id = keeper.AppendVirtualPriceTickQueue(ctx, items[i])

	}
	return items
}

func TestVirtualPriceTickQueueGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNVirtualPriceTickQueue(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetVirtualPriceTickQueue(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestVirtualPriceTickQueueRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNVirtualPriceTickQueue(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveVirtualPriceTickQueue(ctx, item.Id)
		_, found := keeper.GetVirtualPriceTickQueue(ctx, item.Id)
		require.False(t, found)
	}
}

func TestVirtualPriceTickQueueGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNVirtualPriceTickQueue(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllVirtualPriceTickQueue(ctx)),
	)
}

func TestVirtualPriceTickQueueCount(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNVirtualPriceTickQueue(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetVirtualPriceTickQueueCount(ctx))
}
