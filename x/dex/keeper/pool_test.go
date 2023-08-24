package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/duality-labs/duality/testutil/keeper"
	"github.com/stretchr/testify/require"
)

func TestPoolInit(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)

	pool, err := keeper.InitPool(ctx, defaultPairID, 0, 1)
	require.NoError(t, err)
	pool.Deposit(sdk.NewInt(100), sdk.NewInt(100), sdk.NewInt(0), true)
	keeper.SetPool(ctx, pool)

	dbPool, found := keeper.GetPool(ctx, defaultPairID, 0, 1)

	require.True(t, found)

	require.Equal(t, pool.ID, dbPool.ID)
	require.Equal(t, pool.LowerTick0, dbPool.LowerTick0)
	require.Equal(t, pool.UpperTick1, dbPool.UpperTick1)
}
