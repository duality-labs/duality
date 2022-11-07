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

func createNTokenObject(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.TokenObject {
	items := make([]types.TokenObject, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetTokenObject(ctx, items[i])
	}
	return items
}

func TestTokenObjectGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTokenObject(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetTokenObject(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTokenObjectRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTokenObject(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTokenObject(ctx,
			item.Address,
		)
		_, found := keeper.GetTokenObject(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestTokenObjectGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTokenObject(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTokenObject(ctx)),
	)
}
