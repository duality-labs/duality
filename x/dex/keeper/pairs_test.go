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

func createNPairs(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Pairs {
	items := make([]types.Pairs, n)
	for i := range items {
		items[i].Token0 = strconv.Itoa(i)
		items[i].Token1 = strconv.Itoa(i)

		keeper.SetPairs(ctx, items[i])
	}
	return items
}

func TestPairsGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNPairs(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPairs(ctx,
			item.Token0,
			item.Token1,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPairsRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNPairs(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePairs(ctx,
			item.Token0,
			item.Token1,
		)
		_, found := keeper.GetPairs(ctx,
			item.Token0,
			item.Token1,
		)
		require.False(t, found)
	}
}

func TestPairsGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNPairs(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPairs(ctx)),
	)
}
