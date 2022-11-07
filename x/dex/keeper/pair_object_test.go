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

func createNPairObject(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.PairObject {
	items := make([]types.PairObject, n)
	for i := range items {
		items[i].PairId = strconv.Itoa(i)

		keeper.SetPairObject(ctx, items[i])
	}
	return items
}

func TestPairObjectGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNPairObject(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPairObject(ctx,
			item.PairId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPairObjectRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNPairObject(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePairObject(ctx,
			item.PairId,
		)
		_, found := keeper.GetPairObject(ctx,
			item.PairId,
		)
		require.False(t, found)
	}
}

func TestPairObjectGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNPairObject(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPairObject(ctx)),
	)
}
