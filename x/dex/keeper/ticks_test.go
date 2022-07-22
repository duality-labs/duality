package keeper_test

import (
	//"fmt"
	keepertest "github.com/NicholasDotSol/duality/testutil/keeper"
	"github.com/NicholasDotSol/duality/testutil/nullify"
	"github.com/NicholasDotSol/duality/x/dex/keeper"
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNTicks(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Ticks {
	items := make([]types.Ticks, n)
	for i := range items {
		items[i].Token0 = strconv.Itoa(i)
		items[i].Token1 = strconv.Itoa(i)
		items[i].PoolsOneToZero = []*types.Pool{
			&types.Pool{
				Reserve0:    sdk.NewDec(100),
				Reserve1:    sdk.NewDec(100),
				Fee:         sdk.NewDecWithPrec(3, 1),
				Price:       sdk.NewDec(2),
				TotalShares: sdk.NewDec(200),
			},
			&types.Pool{
				Reserve0:    sdk.NewDec(100),
				Reserve1:    sdk.NewDec(100),
				Fee:         sdk.NewDecWithPrec(5, 1),
				Price:       sdk.NewDec(2),
				TotalShares: sdk.NewDec(200),
			},
		}
		items[i].PoolsZeroToOne = []*types.Pool{
			&types.Pool{
				Reserve0:    sdk.NewDec(100),
				Reserve1:    sdk.NewDec(100),
				Fee:         sdk.NewDecWithPrec(3, 1),
				Price:       sdk.NewDec(2),
				TotalShares: sdk.NewDec(200),
			},
			&types.Pool{
				Reserve0:    sdk.NewDec(100),
				Reserve1:    sdk.NewDec(100),
				Fee:         sdk.NewDecWithPrec(5, 1),
				Price:       sdk.NewDec(2),
				TotalShares: sdk.NewDec(200),
			},
		}
		keeper.SetTicks(ctx, items[i])
	}
	return items
}

func TestTicksGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTicks(keeper, ctx, 10)

	for _, item := range items {
		rst, found := keeper.GetTicks(ctx,
			item.Token0,
			item.Token1,
		)
		//fmt.Println(item)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTicksRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTicks(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTicks(ctx,
			item.Token0,
			item.Token1,
		)
		_, found := keeper.GetTicks(ctx,
			item.Token0,
			item.Token1,
		)
		require.False(t, found)
	}
}

func TestTicksGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTicks(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTicks(ctx)),
	)
}

// func TestInit(t *testing.T) {
// 	keeper, ctx := keepertest.DexKeeper(t)

// 	pools := []*types.Pool{
// 		&types.Pool{
// 			Reserve0:    sdk.NewDec(100),
// 			Reserve1:    sdk.NewDec(100),
// 			Fee: sdk.NewDecWithPrec(3, 1),
// 			Price:       sdk.NewDec(2),
// 			TotalShares: sdk.NewDec(200),
// 			Index:       0,
// 		},
// 		&types.Pool{
// 			Reserve0:    sdk.NewDec(100),
// 			Reserve1:    sdk.NewDec(100),
// 			Fee: sdk.NewDecWithPrec(3, 1),
// 			Price:       sdk.NewDec(2),
// 			TotalShares: sdk.NewDec(200),
// 			Index:       0,
// 		},
// 	}

// 	keeper.Init1to0(&pools)
// 	fmt.Println("Intialized pools", pools)

// 	newPool := &types.Pool{
// 		Reserve0:    sdk.NewDec(100),
// 		Reserve1:    sdk.NewDec(100),
// 		Fee: sdk.NewDecWithPrec(2, 1),
// 		Price:       sdk.NewDec(2),
// 		TotalShares: sdk.NewDec(200),
// 		Index:       0,
// 	}

// 	keeper.Push1to0(&pools, newPool)
// 	fmt.Println("(1) Push New Pool:", pools)

// 	removedPool := keeper.Pop1to0(&pools)
// 	fmt.Println(" (2) Popped Pool: ", pools)
// 	fmt.Println(" (2.a) Pool Popped: ", removedPool)
// 	removedPool = keeper.Remove1to0(&pools, 0)
// 	fmt.Println(" (3) Post Remove: ", pools)
// 	fmt.Println("(3.a) Pool Removed:", removedPool)
// 	keeper.Push1to0(&pools, newPool)
// 	keeper.Push1to0(&pools, newPool)
// 	fmt.Println("HEAD", pools[1].Index)
// 	fmt.Println("(4) Post Pushes:", pools)
// 	keeper.Update1to0(&pools, pools[1], sdk.NewDec(100), sdk.NewDec(100), sdk.NewDecWithPrec(3, 1), sdk.NewDec(200), sdk.NewDec(4))
// 	fmt.Println("(5) Post Update: ", pools)
// 	removedPool = keeper.Pop1to0(&pools)
// 	fmt.Println("(6) Popped pools:", pools)
// 	fmt.Println("Popped Pool", removedPool)
// 	removedPool = keeper.Pop1to0(&pools)
// 	fmt.Println("(7) Popped pools:", pools)
// 	fmt.Println(" (7.a) Popped Pool", removedPool)

// 	newPool2 := &types.Pool{
// 		Reserve0:    sdk.NewDec(100),
// 		Reserve1:    sdk.NewDec(100),
// 		Fee: sdk.NewDecWithPrec(2, 1),
// 		Price:       sdk.NewDec(4),
// 		TotalShares: sdk.NewDec(200),
// 		Index:       0,
// 	}
// 	newPool3 := &types.Pool{
// 		Reserve0:    sdk.NewDec(100),
// 		Reserve1:    sdk.NewDec(100),
// 		Fee: sdk.NewDecWithPrec(5, 1),
// 		Price:       sdk.NewDec(3),
// 		TotalShares: sdk.NewDec(200),
// 		Index:       0,
// 	}
// 	fmt.Println("New Pushes")
// 	fmt.Println("")
// 	keeper.Push1to0(&pools, newPool)
// 	keeper.Push1to0(&pools, newPool2)
// 	keeper.Push1to0(&pools, newPool3)
// 	keeper.Push1to0(&pools, newPool2)
// 	fmt.Println(pools)
// 	removedPool = keeper.Pop1to0(&pools)
// 	fmt.Println(pools)
// 	fmt.Println("Popped Pool", removedPool)
// 	removedPool = keeper.Pop1to0(&pools)
// 	fmt.Println(pools)
// 	fmt.Println("Popped Pool", removedPool)
// 	removedPool = keeper.Pop1to0(&pools)
// 	fmt.Println(pools)
// 	fmt.Println("Popped Pool", removedPool)
// 	removedPool = keeper.Pop1to0(&pools)
// 	fmt.Println(pools)
// 	fmt.Println("Popped Pool", removedPool)

// 	if _, found := keeper.GetTicks(ctx, "tokenA", "tokenB"); found == false {
// 		keeper.SetTicks(ctx, types.Ticks{
// 			Token0:         "tokenA",
// 			Token1:         "tokenB",
// 			PoolsZeroToOne: pools,
// 			PoolsOneToZero: pools,
// 		})
// 	}

// 	tick, found := keeper.GetTicks(ctx, "tokenA", "tokenB")

// 	require.True(t, found)

// 	fmt.Println(tick)
// 	fmt.Println(tick.PoolsZeroToOne)
// 	removedPool = keeper.Pop1to0(&tick.PoolsZeroToOne)
// 	fmt.Println(&tick.PoolsZeroToOne)
// 	fmt.Println("Popped Pool", removedPool)

// 	keeper.SetTicks(ctx, types.Ticks{
// 		Token0:         "tokenA",
// 		Token1:         "tokenB",
// 		PoolsZeroToOne: tick.PoolsZeroToOne,
// 		PoolsOneToZero: tick.PoolsOneToZero,
// 	})

// 	tick1, found := keeper.GetTicks(ctx, "tokenA", "tokenB")
// 	require.True(t, found)
// 	fmt.Println(tick1.PoolsZeroToOne)

// 	_ = ctx
// }
