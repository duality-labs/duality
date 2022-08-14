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

func createNBitArr(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.BitArr {
	items := make([]types.BitArr, n)
	for i := range items {
		items[i].Id = keeper.AppendBitArr(ctx, items[i])
	}
	return items
}

func TestBitArrGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNBitArr(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetBitArr(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestBitArrRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNBitArr(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveBitArr(ctx, item.Id)
		_, found := keeper.GetBitArr(ctx, item.Id)
		require.False(t, found)
	}
}

func TestBitArrGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNBitArr(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllBitArr(ctx)),
	)
}

func TestBitArrCount(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNBitArr(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetBitArrCount(ctx))
}
