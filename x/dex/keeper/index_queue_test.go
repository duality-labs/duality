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

func createNIndexQueue(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.IndexQueue {
	items := make([]types.IndexQueue, n)
	for i := range items {
		items[i].Queue = append(items[i].Queue, &types.IndexQueueType{
			Price: sdk.NewDec(1),
			Fee:   sdk.ZeroDec(),
			Orderparams: &types.OrderParams{
				"",
				"",
				sdk.ZeroDec(),
			},
		})
		items[i].Index = int32(i)

		keeper.SetIndexQueue(ctx, "Token0", "Token1", items[i])
	}
	//fmt.Println(items)
	return items
}

func TestIndexQueueGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNIndexQueue(keeper, ctx, 10)
	for i := range items {
		rst, found := keeper.GetIndexQueue(ctx,
			"Token0", "Token1",
			items[i].Index,
		)

		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&items[i]),
			nullify.Fill(&rst),
		)
	}
}
func TestIndexQueueRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNIndexQueue(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveIndexQueue(ctx, "Token0", "Token1",
			item.Index,
		)
		_, found := keeper.GetIndexQueue(ctx, "Token0", "Token1",
			item.Index,
		)
		require.False(t, found)
	}
}

func TestIndexQueueGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNIndexQueue(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllIndexQueueByPair(ctx, "Token0", "Token1")),
	)
}
