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

func CreateNTickLiquidityLimitOrder(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.TickLiquidity {
	items := make([]types.TickLiquidity, n)
	for i := range items {
		items[i] = types.TickLiquidity{
			Liquidity: &types.TickLiquidity_LimitOrderTranche{
				LimitOrderTranche: &types.LimitOrderTranche{
					PairId:           &types.PairId{Token0: "TokenA", Token1: "TokenB"},
					TokenIn:          "TokenA",
					TickIndex:        int64(i),
					TrancheIndex:     uint64(i),
					ReservesTokenIn:  sdk.NewInt(1),
					ReservesTokenOut: sdk.NewInt(1),
					TotalTokenIn:     sdk.NewInt(1),
					TotalTokenOut:    sdk.NewInt(1),
				},
			},
		}
		keeper.SetTickLiquidity(ctx, items[i])
	}
	return items
}

func TestTickLiquidityGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := CreateNTickLiquidityLimitOrder(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetTickLiquidity(ctx,
			item.PairId(),
			item.TokenIn(),
			item.TickIndex(),
			item.LiquidityType(),
			item.LiquidityIndex(),
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTickLiquidityRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := CreateNTickLiquidityLimitOrder(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTickLiquidity(ctx,
			item.PairId(),
			item.TokenIn(),
			item.TickIndex(),
			item.LiquidityType(),
			item.LiquidityIndex(),
		)
		_, found := keeper.GetTickLiquidity(ctx,
			item.PairId(),
			item.TokenIn(),
			item.TickIndex(),
			item.LiquidityType(),
			item.LiquidityIndex(),
		)
		require.False(t, found)
	}
}

func TestTickLiquidityGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := CreateNTickLiquidityLimitOrder(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTickLiquidity(ctx)),
	)
}
