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

func createNTickLiquidity(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.TickLiquidity {
	items := make([]types.TickLiquidity, n)
	for i := range items {
		items[i].PairId = &types.PairId{Token0: "TokenA", Token1: "TokenB"}
		items[i].TokenIn = strconv.Itoa(i)
		items[i].TickIndex = int64(i)
		items[i].LiquidityType = strconv.Itoa(i)
		items[i].LiquidityIndex = uint64(i)

		keeper.SetTickLiquidity(ctx, items[i])
	}
	return items
}

func TestTickLiquidityGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTickLiquidity(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetTickLiquidity(ctx,
			item.PairId,
			item.TokenIn,
			item.TickIndex,
			item.LiquidityType,
			item.LiquidityIndex,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTickLiquidityRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTickLiquidity(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTickLiquidity(ctx, item)
		_, found := keeper.GetTickLiquidity(ctx,
			item.PairId,
			item.TokenIn,
			item.TickIndex,
			item.LiquidityType,
			item.LiquidityIndex,
		)
		require.False(t, found)
	}
}

func TestTickLiquidityGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTickLiquidity(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTickLiquidity(ctx)),
	)
}
