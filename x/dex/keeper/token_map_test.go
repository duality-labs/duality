package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/duality-labs/duality/testutil/keeper"
	"github.com/duality-labs/duality/testutil/nullify"
	"github.com/duality-labs/duality/x/dex/keeper"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNTokenMap(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.TokenMap {
	items := make([]types.TokenMap, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetTokenMap(ctx, items[i])
	}
	return items
}

func TestTokenMapGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTokenMap(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetTokenMap(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTokenMapRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTokenMap(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTokenMap(ctx,
			item.Address,
		)
		_, found := keeper.GetTokenMap(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestTokenMapGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTokenMap(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTokenMap(ctx)),
	)
}
