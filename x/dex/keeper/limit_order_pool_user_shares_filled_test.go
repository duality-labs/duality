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

func createNLimitOrderPoolUserSharesFilled(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.LimitOrderPoolUserSharesFilled {
	items := make([]types.LimitOrderPoolUserSharesFilled, n)
	for i := range items {
		items[i].Count = strconv.Itoa(i)
		items[i].Address = strconv.Itoa(i)

		keeper.SetLimitOrderPoolUserSharesFilled(ctx, items[i])
	}
	return items
}

func TestLimitOrderPoolUserSharesFilledGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNLimitOrderPoolUserSharesFilled(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetLimitOrderPoolUserSharesFilled(ctx,
			item.Count,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestLimitOrderPoolUserSharesFilledRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNLimitOrderPoolUserSharesFilled(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveLimitOrderPoolUserSharesFilled(ctx,
			item.Count,
			item.Address,
		)
		_, found := keeper.GetLimitOrderPoolUserSharesFilled(ctx,
			item.Count,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestLimitOrderPoolUserSharesFilledGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNLimitOrderPoolUserSharesFilled(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllLimitOrderPoolUserSharesFilled(ctx)),
	)
}
