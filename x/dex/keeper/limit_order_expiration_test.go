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

func createNLimitOrderExpiration(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.LimitOrderExpiration {
	items := make([]types.LimitOrderExpiration, n)
	for i := range items {
		items[i].ExpirationTime = time.Unix(int64(i), 10).UTC()
		items[i].TrancheRef = []byte(strconv.Itoa(i))

		keeper.SetLimitOrderExpiration(ctx, items[i])
	}
	return items
}

func TestLimitOrderExpirationGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNLimitOrderExpiration(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetLimitOrderExpiration(ctx,
			item.ExpirationTime,
			item.TrancheRef,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestLimitOrderExpirationRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNLimitOrderExpiration(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveLimitOrderExpiration(ctx,
			item.ExpirationTime,
			item.TrancheRef,
		)
		_, found := keeper.GetLimitOrderExpiration(ctx,
			item.ExpirationTime,
			item.TrancheRef,
		)
		require.False(t, found)
	}
}

func TestLimitOrderExpirationGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNLimitOrderExpiration(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllLimitOrderExpiration(ctx)),
	)
}
