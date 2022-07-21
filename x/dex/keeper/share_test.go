package keeper_test

import (
	//"fmt"
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

func createNShare(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Share {
	items := make([]types.Share, n)
	for i := range items {
		items[i].Owner = strconv.Itoa(i)
		items[i].Token0 = strconv.Itoa(i)
		items[i].Token1 = strconv.Itoa(i)
		items[i].Price = strconv.Itoa(i)
		items[i].Fee = strconv.Itoa(i) 
		items[i].ShareAmount = sdk.ZeroDec()
		keeper.SetShare(ctx, items[i])

	}
	return items
}

func TestShareGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNShare(keeper, ctx, 10)

	for _, item := range items {
		rst, found := keeper.GetShare(ctx,
			item.Owner,
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
func TestShareRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNShare(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveShare(ctx,
			item.Owner,
			item.Token0,
			item.Token1,
			item.Price,
			item.Fee,
		)
		_, found := keeper.GetShare(ctx,
			item.Owner,
			item.Token0,
			item.Token1,
			item.Price,
			item.Fee,
		)
		require.False(t, found)
	}
}

func TestShareGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNShare(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllShare(ctx)),
	)
}
