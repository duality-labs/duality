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

func createNGoodTilRecord(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.GoodTilRecord {
	items := make([]types.GoodTilRecord, n)
	for i := range items {
		items[i].GoodTilDate = time.Unix(int64(i), 10).UTC()
		items[i].TrancheRef = []byte(strconv.Itoa(i))

		keeper.SetGoodTilRecord(ctx, items[i])
	}
	return items
}

func TestGoodTilRecordGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNGoodTilRecord(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetGoodTilRecord(ctx,
			item.GoodTilDate,
			item.TrancheRef,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestGoodTilRecordRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNGoodTilRecord(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveGoodTilRecord(ctx,
			item.GoodTilDate,
			item.TrancheRef,
		)
		_, found := keeper.GetGoodTilRecord(ctx,
			item.GoodTilDate,
			item.TrancheRef,
		)
		require.False(t, found)
	}
}

func TestGoodTilRecordGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNGoodTilRecord(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllGoodTilRecord(ctx)),
	)
}
