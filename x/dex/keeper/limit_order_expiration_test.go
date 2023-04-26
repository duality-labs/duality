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

func createLimitOrderExpirationAndTranches(keeper *keeper.Keeper, ctx sdk.Context, expTimes []time.Time) {
	items := make([]types.LimitOrderExpiration, len(expTimes))
	for i := range items {
		tranche := types.LimitOrderTranche{
			PairID:           &types.PairID{Token0: "TokenA", Token1: "TokenB"},
			TokenIn:          "TokenA",
			TickIndex:        0,
			TrancheKey:       strconv.Itoa(i),
			ReservesTokenIn:  sdk.NewInt(10),
			ReservesTokenOut: sdk.NewInt(10),
			TotalTokenIn:     sdk.NewInt(10),
			TotalTokenOut:    sdk.NewInt(10),
			ExpirationTime:   &expTimes[i],
		}
		items[i].ExpirationTime = expTimes[i]
		items[i].TrancheRef = tranche.Ref()

		keeper.SetLimitOrderExpiration(ctx, items[i])
		keeper.SetLimitOrderTranche(ctx, tranche)
	}
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

func TestPurgeExpiredLimitOrders(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	now := time.Now().UTC()
	ctx = ctx.WithBlockTime(now)
	ctx = ctx.WithBlockGasMeter(sdk.NewGasMeter(1000000))

	yesterday := now.AddDate(0, 0, -1)
	tomorrow := now.AddDate(0, 0, 1)
	nextWeek := now.AddDate(0, 0, 7)

	expTimes := []time.Time{
		yesterday,
		yesterday,
		now,
		tomorrow,
		nextWeek,
	}

	createLimitOrderExpirationAndTranches(keeper, ctx, expTimes)
	keeper.PurgeExpiredLimitOrders(ctx, now)

	// Only future LimitOrderExpiration items still exist
	expList := keeper.GetAllLimitOrderExpiration(ctx)
	require.Equal(t, 2, len(expList))
	require.Equal(t, tomorrow, expList[0].ExpirationTime)
	require.Equal(t, nextWeek, expList[1].ExpirationTime)

	// Only future LimitOrderTranches Exist
	trancheList := keeper.GetAllLimitOrderTrancheAtIndex(ctx, defaultPairID, "TokenA", 0)
	require.Equal(t, 2, len(trancheList))
	require.Equal(t, tomorrow, *trancheList[0].ExpirationTime)
	require.Equal(t, nextWeek, *trancheList[1].ExpirationTime)

	// InactiveLimitOrderTranches have been created for the expired tranched
	inactiveTrancheList := keeper.GetAllInactiveLimitOrderTranche(ctx)
	require.Equal(t, 3, len(inactiveTrancheList))
	require.Equal(t, yesterday, *inactiveTrancheList[0].ExpirationTime)
	require.Equal(t, yesterday, *inactiveTrancheList[1].ExpirationTime)
	require.Equal(t, now, *inactiveTrancheList[2].ExpirationTime)
}

func TestPurgeExpiredLimitOrdersAtBlockGasLimit(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	now := time.Now().UTC()
	ctx = ctx.WithBlockTime(now)
	gasLimt := 1000000
	ctx = ctx.WithBlockGasMeter(sdk.NewGasMeter(uint64(gasLimt)))
	gasUsed := gasLimt - types.GoodTilPurgeGasBuffer - 30000

	yesterday := now.AddDate(0, 0, -1)

	expTimes := []time.Time{
		types.JITGoodTilTime(),
		types.JITGoodTilTime(),
		yesterday,
		yesterday,
		yesterday,
	}
	createLimitOrderExpirationAndTranches(keeper, ctx, expTimes)

	// IF blockGasMeter is nearing the GoodTilPurgeBuffer
	ctx = ctx.WithGasMeter(sdk.NewGasMeter(100000))
	ctx.BlockGasMeter().ConsumeGas(uint64(gasUsed), "stub block gas usage")

	// WHEN PurgeExpiredLimitOrders is run
	keeper.PurgeExpiredLimitOrders(ctx, now)

	// THEN GoodTilPurgeHitGasLimit event is emitted
	keepertest.AssertEventEmitted(t, ctx, types.GoodTilPurgeHitGasLimitEventKey, "Gas Limit Event not emitted")

	// All JIT expirations are purged but other expirations remain
	expList := keeper.GetAllLimitOrderExpiration(ctx)
	// NOTE: this test is very brittle because it relies on an estimated cost
	// for deleting expirations. If this cost changes the number of remaining
	// expirations may change
	require.Equal(t, 1, len(expList))
}

func TestPurgeExpiredLimitOrdersAtBlockGasLimitOnlyJIT(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	now := time.Now().UTC()
	ctx = ctx.WithBlockTime(now)
	gasLimt := 1000000
	ctx = ctx.WithBlockGasMeter(sdk.NewGasMeter(uint64(gasLimt)))
	gasUsed := gasLimt - types.GoodTilPurgeGasBuffer - 30000

	expTimes := []time.Time{
		types.JITGoodTilTime(),
		types.JITGoodTilTime(),
		types.JITGoodTilTime(),
		types.JITGoodTilTime(),
		types.JITGoodTilTime(),
		types.JITGoodTilTime(),
		types.JITGoodTilTime(),
	}

	createLimitOrderExpirationAndTranches(keeper, ctx, expTimes)
	ctx = ctx.WithGasMeter(sdk.NewGasMeter(100000))
	ctx.BlockGasMeter().ConsumeGas(uint64(gasUsed), "stub block gas usage")
	keeper.PurgeExpiredLimitOrders(ctx, now)

	// GoodTilPurgeHitGasLimit event is not been emitted
	keepertest.AssertEventNotEmitted(t, ctx, types.GoodTilPurgeHitGasLimitEventGas, "Hit gas limit purging JIT expirations")

	// All JIT expirations are purged
	expList := keeper.GetAllLimitOrderExpiration(ctx)
	require.Equal(t, 0, len(expList))
}
