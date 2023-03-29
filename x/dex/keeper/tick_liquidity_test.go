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

func CreateNTickLiquidity(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.TickLiquidity {
	items := make([]types.TickLiquidity, n)
	for i := range items {
		tick := types.TickLiquidity{
			Liquidity: &types.TickLiquidity_LimitOrderTranche{
				LimitOrderTranche: &types.LimitOrderTranche{
					PairID:           &types.PairID{Token0: "TokenA", Token1: "TokenB"},
					TokenIn:          "TokenA",
					TickIndex:        int64(i),
					TrancheKey:       strconv.Itoa(i),
					ReservesTokenIn:  sdk.NewInt(10),
					ReservesTokenOut: sdk.NewInt(10),
					TotalTokenIn:     sdk.NewInt(10),
					TotalTokenOut:    sdk.NewInt(10),
				},
			},
		}
		keeper.SetLimitOrderTranche(ctx, *tick.GetLimitOrderTranche())
		items[i] = tick
	}
	return items
}

func TestTickLiquidityGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := CreateNTickLiquidity(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTickLiquidity(ctx)),
	)
}
