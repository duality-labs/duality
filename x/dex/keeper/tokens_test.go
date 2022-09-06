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

func createNTokens(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Tokens {
	items := make([]types.Tokens, n)
	for i := range items {
		items[i].Id = keeper.AppendTokens(ctx, items[i])
	}
	return items
}

func TestTokensGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTokens(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetTokens(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestTokensRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTokens(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTokens(ctx, item.Id)
		_, found := keeper.GetTokens(ctx, item.Id)
		require.False(t, found)
	}
}

func TestTokensGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTokens(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTokens(ctx)),
	)
}

func TestTokensCount(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTokens(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetTokensCount(ctx))
}
