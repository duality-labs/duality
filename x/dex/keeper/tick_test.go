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

func createNTick(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Tick {
	items := make([]types.Tick, n)
	for i := range items {
		items[i].Token0 = strconv.Itoa(i)
		items[i].Token1 = strconv.Itoa(i)
		items[i].Price = strconv.Itoa(i)
		items[i].Fee = uint64(i)

		keeper.SetTick(ctx, items[i])
	}
	return items
}

func TestTickGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTick(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetTick(ctx,
			item.Token0,
			item.Token1,
			item.Price,
			item.Fee,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTickRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTick(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTick(ctx,
			item.Token0,
			item.Token1,
			item.Price,
			item.Fee,
		)
		_, found := keeper.GetTick(ctx,
			item.Token0,
			item.Token1,
			item.Price,
			item.Fee,
		)
		require.False(t, found)
	}
}

func TestTickGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTick(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTick(ctx)),
	)
}
