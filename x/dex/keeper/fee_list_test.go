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

func createNFeeList(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.FeeList {
	items := make([]types.FeeList, n)
	for i := range items {
		items[i].Id = keeper.AppendFeeList(ctx, items[i])
	}
	return items
}

func TestFeeListGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNFeeList(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetFeeList(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestFeeListRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNFeeList(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveFeeList(ctx, item.Id)
		_, found := keeper.GetFeeList(ctx, item.Id)
		require.False(t, found)
	}
}

func TestFeeListGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNFeeList(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllFeeList(ctx)),
	)
}

func TestFeeListCount(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNFeeList(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetFeeListCount(ctx))
}
