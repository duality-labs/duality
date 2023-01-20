package keeper_test

import (
	"strconv"
	"testing"

	dualityapp "github.com/duality-labs/duality/app"
	keepertest "github.com/duality-labs/duality/testutil/keeper"
	"github.com/duality-labs/duality/testutil/nullify"
	"github.com/duality-labs/duality/x/dex/keeper"
	"github.com/duality-labs/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

type TickTestSuite struct {
	suite.Suite
	app *dualityapp.App
	ctx sdk.Context
}

func TestTickTestSuite(t *testing.T) {
	suite.Run(t, new(TickTestSuite))
}

func (s *TickTestSuite) SetupTest() {
	s.app = dualityapp.Setup(false)
	s.ctx = s.app.BaseApp.NewContext(false, tmproto.Header{})
}

func (s *TickTestSuite) TestLimitHasToken0() {
	pairId := keeper.CreatePairId("TokenA", "TokenB")
	tick := keeper.NewTick(pairId, 0, 6)
	s.app.DexKeeper.SetLimitOrderTranche(s.ctx, types.LimitOrderTranche{
		PairId:          pairId,
		TokenIn:         "TokenA",
		TickIndex:       tick.TickIndex,
		TrancheIndex:    0,
		ReservesTokenIn: sdk.NewInt(5),
	})
	s.Equal(true, s.app.DexKeeper.TickHasToken0(s.ctx, &tick))
}

func (s *TickTestSuite) TestEmptyHasToken0() {
	pairId := keeper.CreatePairId("TokenA", "TokenB")
	tick := keeper.NewTick(pairId, 0, 6)
	s.app.DexKeeper.SetLimitOrderTranche(s.ctx, types.LimitOrderTranche{
		PairId:          pairId,
		TokenIn:         "TokenA",
		TickIndex:       tick.TickIndex,
		TrancheIndex:    0,
		ReservesTokenIn: sdk.NewInt(0),
	})
	s.Equal(false, s.app.DexKeeper.TickHasToken0(s.ctx, &tick))
}

func (s *TickTestSuite) TestLiquidityHasToken0() {
	pairId := keeper.CreatePairId("TokenA", "TokenB")
	tick := keeper.NewTick(pairId, 0, 6)
	tick.TickData.Reserve0[0] = sdk.NewInt(100)
	tick.TickData.Reserve1[0] = sdk.NewInt(20)
	s.app.DexKeeper.SetLimitOrderTranche(s.ctx, types.LimitOrderTranche{
		PairId:          pairId,
		TokenIn:         "TokenA",
		TickIndex:       tick.TickIndex,
		TrancheIndex:    0,
		ReservesTokenIn: sdk.NewInt(0),
	})
	s.Equal(true, s.app.DexKeeper.TickHasToken0(s.ctx, &tick))
}

func (s *TickTestSuite) TestLiquidityHighFeeHasToken0() {
	pairId := keeper.CreatePairId("TokenA", "TokenB")
	tick := keeper.NewTick(pairId, 0, 6)
	tick.TickData.Reserve0[4] = sdk.NewInt(100)
	tick.TickData.Reserve1[4] = sdk.NewInt(20)
	s.app.DexKeeper.SetLimitOrderTranche(s.ctx, types.LimitOrderTranche{
		PairId:          pairId,
		TokenIn:         "TokenA",
		TickIndex:       tick.TickIndex,
		TrancheIndex:    0,
		ReservesTokenIn: sdk.NewInt(0),
	})
	s.Equal(true, s.app.DexKeeper.TickHasToken0(s.ctx, &tick))
}

func (s *TickTestSuite) TestNoLiquidityOnOneSideHasToken0() {
	pairId := keeper.CreatePairId("TokenA", "TokenB")
	tick := keeper.NewTick(pairId, 0, 6)
	tick.TickData.Reserve0[4] = sdk.NewInt(0)
	tick.TickData.Reserve1[4] = sdk.NewInt(20)
	s.app.DexKeeper.SetLimitOrderTranche(s.ctx, types.LimitOrderTranche{
		PairId:          pairId,
		TokenIn:         "TokenA",
		TickIndex:       tick.TickIndex,
		TrancheIndex:    0,
		ReservesTokenIn: sdk.NewInt(0),
	})
	s.Equal(false, s.app.DexKeeper.TickHasToken0(s.ctx, &tick))
}

func (s *TickTestSuite) TestLimitHasToken1() {
	pairId := keeper.CreatePairId("TokenA", "TokenB")
	tick := keeper.NewTick(pairId, 0, 6)
	s.app.DexKeeper.SetLimitOrderTranche(s.ctx, types.LimitOrderTranche{
		PairId:          pairId,
		TokenIn:         "TokenB",
		TickIndex:       tick.TickIndex,
		TrancheIndex:    0,
		ReservesTokenIn: sdk.NewInt(5),
	})
	s.Equal(true, s.app.DexKeeper.TickHasToken1(s.ctx, &tick))
}

func (s *TickTestSuite) TestEmptyHasToken1() {
	pairId := keeper.CreatePairId("TokenA", "TokenB")
	tick := keeper.NewTick(pairId, 0, 6)
	s.app.DexKeeper.SetLimitOrderTranche(s.ctx, types.LimitOrderTranche{
		PairId:          pairId,
		TokenIn:         "TokenB",
		TickIndex:       tick.TickIndex,
		TrancheIndex:    0,
		ReservesTokenIn: sdk.NewInt(0),
	})
	s.Equal(false, s.app.DexKeeper.TickHasToken1(s.ctx, &tick))
}

func (s *TickTestSuite) TestLiquidityHasToken1() {
	pairId := keeper.CreatePairId("TokenA", "TokenB")
	tick := keeper.NewTick(pairId, 0, 6)
	tick.TickData.Reserve0[0] = sdk.NewInt(100)
	tick.TickData.Reserve1[0] = sdk.NewInt(20)
	s.app.DexKeeper.SetLimitOrderTranche(s.ctx, types.LimitOrderTranche{
		PairId:          pairId,
		TokenIn:         "TokenB",
		TickIndex:       tick.TickIndex,
		TrancheIndex:    0,
		ReservesTokenIn: sdk.NewInt(0),
	})
	s.Equal(true, s.app.DexKeeper.TickHasToken1(s.ctx, &tick))
}

func (s *TickTestSuite) TestLiquidityHighFeeHasToken1() {
	pairId := keeper.CreatePairId("TokenA", "TokenB")
	tick := keeper.NewTick(pairId, 0, 6)
	tick.TickData.Reserve0[4] = sdk.NewInt(100)
	tick.TickData.Reserve1[4] = sdk.NewInt(20)
	s.app.DexKeeper.SetLimitOrderTranche(s.ctx, types.LimitOrderTranche{
		PairId:          pairId,
		TokenIn:         "TokenB",
		TickIndex:       tick.TickIndex,
		TrancheIndex:    0,
		ReservesTokenIn: sdk.NewInt(0),
	})
	s.Equal(true, s.app.DexKeeper.TickHasToken1(s.ctx, &tick))
}

func (s *TickTestSuite) TestNoLiquidityOnOneSideHasToken1() {
	pairId := keeper.CreatePairId("TokenA", "TokenB")
	tick := keeper.NewTick(pairId, 0, 6)
	tick.TickData.Reserve0[4] = sdk.NewInt(100)
	tick.TickData.Reserve1[4] = sdk.NewInt(0)
	s.app.DexKeeper.SetLimitOrderTranche(s.ctx, types.LimitOrderTranche{
		PairId:          pairId,
		TokenIn:         "TokenB",
		TickIndex:       tick.TickIndex,
		TrancheIndex:    0,
		ReservesTokenIn: sdk.NewInt(0),
	})
	s.Equal(false, s.app.DexKeeper.TickHasToken1(s.ctx, &tick))
}

// Prevent strconv unused error
var _ = strconv.IntSize

func createNTick(keeper *keeper.Keeper, ctx sdk.Context, pairId *types.PairId, n int) []types.Tick {
	items := make([]types.Tick, n)
	for i := range items {

		items[i].TickData = &types.TickDataType{
			Reserve0: []sdk.Int{sdk.OneInt()},
			Reserve1: []sdk.Int{sdk.ZeroInt()},
		}

		items[i].TickIndex = int64(i)
		items[i].PairId = pairId

		// testTick :=  &types.Tick{0, &types.TickDataType{Reserve0AndShares: []*types.Reserve0AndSharesType{
		// 	{Reserve0: sdk.OneDec(),
		// 	TotalShares: sdk.ZeroDec(),
		// }},
		// Reserve1: []sdk.Dec{sdk.ZeroDec()},
		// }}

		keeper.SetTick(ctx, pairId, items[i])
	}

	return items
}

func TestTickGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTick(keeper, ctx, defaultPairId, 10)
	for _, item := range items {
		rst, found := keeper.GetTick(ctx, defaultPairId,
			item.TickIndex,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTickRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTick(keeper, ctx, defaultPairId, 10)
	for _, item := range items {
		keeper.RemoveTick(ctx,
			defaultPairId,
			item.TickIndex,
		)
		_, found := keeper.GetTick(ctx,
			defaultPairId,
			item.TickIndex,
		)
		require.False(t, found)
	}
}

func TestTickGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTick(keeper, ctx, defaultPairId, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTick(ctx)),
	)
}
