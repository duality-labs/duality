package keeper_test

import (
	"strconv"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/duality-labs/duality/testutil/keeper"
	"github.com/duality-labs/duality/testutil/nullify"
	"github.com/duality-labs/duality/x/dex/keeper"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNGoodTillRecord(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.GoodTillRecord {
	items := make([]types.GoodTillRecord, n)
	for i := range items {
		items[i].GoodTillDate = time.Unix(int64(i), 0)
		items[i].TrancheRef = []byte(strconv.Itoa(i))

		keeper.SetGoodTillRecord(ctx, items[i])
	}
	return items
}

func TestGoodTillRecordGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNGoodTillRecord(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetGoodTillRecord(ctx,
			item.GoodTillDate,
			item.TrancheRef,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestGoodTillRecordRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNGoodTillRecord(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveGoodTillRecord(ctx,
			item.GoodTillDate,
			item.TrancheRef,
		)
		_, found := keeper.GetGoodTillRecord(ctx,
			item.GoodTillDate,
			item.TrancheRef,
		)
		require.False(t, found)
	}
}

func TestGoodTillRecordGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNGoodTillRecord(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllGoodTillRecord(ctx)),
	)
}
