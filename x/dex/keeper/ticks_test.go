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

func createNTicks(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Ticks {
	items := make([]types.Ticks, n)
	for i := range items {
		items[i].Price = strconv.Itoa(i)
		items[i].Fee = strconv.Itoa(i)
		items[i].OrderType = strconv.Itoa(i)
		items[i].PairFee = sdk.ZeroDec()
		items[i].PairPrice = sdk.ZeroDec()
		items[i].TotalShares = sdk.ZeroDec()

		keeper.SetTicks(ctx, "0", "1", items[i])
	}
	return items
}

func TestTicksGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTicks(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetTicks(ctx,
			"0",
			"1",
			item.Price,
			item.Fee,
			item.OrderType,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTicksRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTicks(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTicks(ctx,
			"0",
			"1",
			item.Price,
			item.Fee,
			item.OrderType,
		)
		_, found := keeper.GetTicks(ctx,
			"0",
			"1",
			item.Price,
			item.Fee,
			item.OrderType,
		)
		require.False(t, found)
	}
}

func TestTicksGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTicks(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTicks(ctx)),
	)
}
