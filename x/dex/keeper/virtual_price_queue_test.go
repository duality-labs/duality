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

func createNVirtualPriceQueue(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.VirtualPriceQueue {
	items := make([]types.VirtualPriceQueue, n)
	for i := range items {
		items[i].VPrice = strconv.Itoa(i)
		items[i].Direction = strconv.Itoa(i)
		items[i].OrderType = strconv.Itoa(i)

		keeper.SetVirtualPriceQueue(ctx, items[i])
	}
	return items
}

func TestVirtualPriceQueueGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNVirtualPriceQueue(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetVirtualPriceQueue(ctx,
			item.VPrice,
			item.Direction,
			item.OrderType,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestVirtualPriceQueueRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNVirtualPriceQueue(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveVirtualPriceQueue(ctx,
			item.VPrice,
			item.Direction,
			item.OrderType,
		)
		_, found := keeper.GetVirtualPriceQueue(ctx,
			item.VPrice,
			item.Direction,
			item.OrderType,
		)
		require.False(t, found)
	}
}

func TestVirtualPriceQueueGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNVirtualPriceQueue(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllVirtualPriceQueue(ctx)),
	)
}
