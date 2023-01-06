package keeper_test

import (
	"math"

	"github.com/NicholasDotSol/duality/x/dex/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	feeCount = 4
)

// PairToTokens ///////////////////////////////////////////////////////////////

func (s *MsgServerTestSuite) TestPairToTokens() {

	token0, token1 := keeper.PairToTokens("TokenA<>TokenB")

	s.Assert().Equal("TokenA", token0)
	s.Assert().Equal("TokenB", token1)

}

func (s *MsgServerTestSuite) TestPairToTokensIBCis0() {

	token0, token1 := keeper.PairToTokens("ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2<>TokenB")

	s.Assert().Equal("ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2", token0)
	s.Assert().Equal("TokenB", token1)
}

func (s *MsgServerTestSuite) TestPairToTokensIBCis1() {

	token0, token1 := keeper.PairToTokens("TokenA<>ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2")

	s.Assert().Equal("TokenA", token0)
	s.Assert().Equal("ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2", token1)

}
func (s *MsgServerTestSuite) TestPairToTokensIBCisBoth() {

	token0, token1 :=
		keeper.PairToTokens("ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2<>ibc/94644FB092D9ACDA56123C74F36E4234926001AA44A9CA97EA622B25F41E5223")

	s.Assert().Equal("ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2", token0)
	s.Assert().Equal("ibc/94644FB092D9ACDA56123C74F36E4234926001AA44A9CA97EA622B25F41E5223", token1)
}

// TokenInit //////////////////////////////////////////////////////////////////

func (s *MsgServerTestSuite) TestTokenInitNew() {

	s.app.DexKeeper.TokenInit(s.ctx, "TokenA")

	tokenMap, found := s.app.DexKeeper.GetTokenMap(s.ctx, "TokenA")

	s.Assert().True(found)
	s.Assert().Equal("TokenA", tokenMap.Address)
	s.Assert().Equal(uint64(1), s.app.DexKeeper.GetTokensCount(s.ctx))
}

func (s *MsgServerTestSuite) TestTokenInitExisting() {

	s.app.DexKeeper.TokenInit(s.ctx, "TokenA")
	s.Require().Equal(uint64(1), s.app.DexKeeper.GetTokensCount(s.ctx))

	s.app.DexKeeper.TokenInit(s.ctx, "TokenA")
}

// GetOrInitPair //////////////////////////////////////////////////////////////

func (s *MsgServerTestSuite) TestGetOrInitPairNew() {
	// GIVEN we initialize a new pair
	s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")

	// THEN the component tokens are also initialized...
	_, found0 := s.app.DexKeeper.GetTokenMap(s.ctx, "TokenA")
	_, found1 := s.app.DexKeeper.GetTokenMap(s.ctx, "TokenA")

	s.Assert().True(found0)
	s.Assert().True(found1)

	// AND 1 pair is initialized with the correct default values
	pairCount := len(s.app.DexKeeper.GetAllTradingPair(s.ctx))
	s.Assert().Equal(1, pairCount)

	pair, foundPair := s.app.DexKeeper.GetTradingPair(s.ctx, "TokenA<>TokenB")

	s.Require().True(foundPair)

	s.Assert().Equal(pair.PairId, "TokenA<>TokenB")
	s.Assert().Equal(int64(math.MaxInt64), pair.CurrentTick0To1)
	s.Assert().Equal(int64(math.MinInt64), pair.CurrentTick1To0)
	s.Assert().Equal(int64(math.MaxInt64), pair.MinTick)
	s.Assert().Equal(int64(math.MinInt64), pair.MaxTick)
}

func (s *MsgServerTestSuite) TestGetOrInitPairExisting() {

	// GIVEN we initialize a pair TokenA/TokenB
	s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")

	// WHEN we update values on that pair
	pair, _ := s.app.DexKeeper.GetTradingPair(s.ctx, "TokenA<>TokenB")
	pair.MinTick = 20
	s.app.DexKeeper.SetTradingPair(s.ctx, pair)

	// AND try to initialize the same pair again
	newPair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")

	// THEN there is still only 1 pair and it retains the values we set
	pairCount := len(s.app.DexKeeper.GetAllTradingPair(s.ctx))
	s.Assert().Equal(1, pairCount)
	s.Assert().Equal(int64(20), newPair.MinTick)
}

// GetOrInitTick //////////////////////////////////////////////////////////////

func (s *MsgServerTestSuite) TestGetOrInitTickNew() {
	// GIVEN we initialize a new tick
	s.app.DexKeeper.GetOrInitTick(s.goCtx, "TokenA<>TokenB", 0)

	// THEN 1 tick is initialized with the correct values

	tickCount := len(s.app.DexKeeper.GetAllTick(s.ctx))
	s.Assert().Equal(1, tickCount)

	tick, found := s.app.DexKeeper.GetTick(s.ctx, "TokenA<>TokenB", 0)

	s.Require().True(found)

	s.Assert().Equal(tick.PairId, "TokenA<>TokenB")
	s.Assert().Equal(tick.TickIndex, int64(0))
	s.Assert().Equal(feeCount, len(tick.TickData.Reserve0))
	s.Assert().Equal(
		sdk.ZeroInt(),
		tick.TickData.Reserve0[0],
	)

	s.Assert().Equal(
		sdk.ZeroInt(),
		tick.TickData.Reserve0[feeCount-1],
	)

	s.Assert().Equal(
		sdk.ZeroInt(),
		tick.TickData.Reserve1[0],
	)

	s.Assert().Equal(
		sdk.ZeroInt(),
		tick.TickData.Reserve1[feeCount-1],
	)

	//AND tranche fill maps are initialized
	_, fill0Found := s.app.DexKeeper.GetLimitOrderTranche(s.ctx, "TokenA<>TokenB", 0, "TokenA", 0)
	_, fill1Found := s.app.DexKeeper.GetLimitOrderTranche(s.ctx, "TokenA<>TokenB", 0, "TokenB", 0)

	s.Assert().True(fill0Found)
	s.Assert().True(fill1Found)
}

func (s *MsgServerTestSuite) TestGetOrInitTickExisting() {

	// GIVEN we initialize a tick
	s.app.DexKeeper.GetOrInitTick(s.goCtx, "TokenA<>TokenB", 0)

	// WHEN we update values on that tick
	tick, _ := s.app.DexKeeper.GetTick(s.ctx, "TokenA<>TokenB", 0)
	tick.TickData.Reserve0[0] = sdk.NewInt(10)
	s.app.DexKeeper.SetTick(s.ctx, "TokenA<>TokenB", tick)

	// AND try to initialize the same tick again
	newTick, err := s.app.DexKeeper.GetOrInitTick(s.goCtx, "TokenA<>TokenB", 0)
	s.Assert().NoError(err)

	// THEN there is still only 1 tick and it retains the values we set
	tickCount := len(s.app.DexKeeper.GetAllTick(s.ctx))
	s.Assert().Equal(1, tickCount)
	s.Assert().Equal(sdk.NewInt(10), newTick.TickData.Reserve0[0])
}

// GetOrInitTickTrancheFillMap ////////////////////////////////////////////////

// TODO: WRITE ME

// GetOrInitReserveData ///////////////////////////////////////////////////////

// TODO: WRITE ME

// GetOrInitUserShareData /////////////////////////////////////////////////////

// TODO: WRITE ME

// GetOrInitOrderPoolTotalShares //////////////////////////////////////////////

// TODO: WRITE ME

// GetOrInitLimitOrderMaps ////////////////////////////////////////////////////

// TODO: WRITE ME

// FindNextTick1To0 ///////////////////////////////////////////////////////////

func (s *MsgServerTestSuite) TestFindNextTick1To0NoLiq() {
	// GIVEN there is no ticks with token0 in the pool
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	// NOTE: Actually adding/funding a tick isn't really neccesary but should make the test less
	// dependent upon current implementation details
	s.setLPAtFee0Pool(1, 0, 10)

	// THEN FindNextTick1To0 doesn't find a tick

	_, found := s.app.DexKeeper.FindNextTick1To0(s.goCtx, pair)
	s.Assert().False(found)

}

func (s *MsgServerTestSuite) TestFindNextTick1To0WithLiq() {
	// GIVEN tick with token0 @ index 0 & currentTick0To1 is 1
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	s.setLPAtFee0Pool(-1, 10, 0)
	s.setLPAtFee0Pool(0, 10, 0)
	s.setLPAtFee0Pool(1, 0, 0)

	// tick -2: (10, 0)
	// tick -1: (10, 0)
	pair.CurrentTick1To0 = 1
	pair.MinTick = -2
	s.app.DexKeeper.SetTradingPair(s.ctx, pair)

	// THEN FindNextTick1To0 finds the tick at 0

	tickIdx, found := s.app.DexKeeper.FindNextTick1To0(s.goCtx, pair)
	s.Require().True(found)
	s.Assert().Equal(int64(-1), tickIdx)

}

func (s *MsgServerTestSuite) TestFindNextTick1To0WithMinLiq() {
	// GIVEN tick with token0 @ index -1 & MinTick = -1
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	s.setLPAtFee0Pool(-1, 10, 0)
	s.setLPAtFee0Pool(1, 0, 0)

	// tick -2: (10, 0)
	pair.CurrentTick1To0 = 1
	pair.MinTick = -2
	s.app.DexKeeper.SetTradingPair(s.ctx, pair)

	// THEN FindNextTick1To0 finds the tick at -1

	tickIdx, found := s.app.DexKeeper.FindNextTick1To0(s.goCtx, pair)
	s.Require().True(found)
	s.Assert().Equal(int64(-2), tickIdx)

}

// FindNextTick0To1 ///////////////////////////////////////////////////////////

func (s *MsgServerTestSuite) TestFindNextTick0To1NoLiq() {
	// GIVEN there are no tick with Token1 in the pool
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	// NOTE: Actually adding/funding a tick isn't really neccesary but should make the test less
	// dependent upon current implementation details
	s.setLPAtFee0Pool(0, 10, 0)

	// THEN FindNextTick0To1 doesn't find a tick

	_, found := s.app.DexKeeper.FindNextTick0To1(s.goCtx, pair)
	s.Assert().False(found)

}

func (s *MsgServerTestSuite) TestFindNextTick0To1WithLiq() {
	// WHEN tick with token1 @ index 0 & currentTick0To1 is -1
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	s.setLPAtFee0Pool(-1, 10, 0)
	s.setLPAtFee0Pool(0, 0, 10)
	s.setLPAtFee0Pool(1, 0, 10)

	// tick -2: (10,  0)
	// tick -1: ( 0,  0)
	// tick  0: ( 0,  0)
	// tick  1: ( 0, 10)
	// tick  2: ( 0, 10)

	pair.CurrentTick0To1 = -1
	pair.MaxTick = 1
	s.app.DexKeeper.SetTradingPair(s.ctx, pair)

	// THEN FindNextTick0To1 finds the tick at 1

	tickIdx, found := s.app.DexKeeper.FindNextTick0To1(s.goCtx, pair)
	s.Require().True(found)
	s.Assert().Equal(int64(1), tickIdx)
}

func (s *MsgServerTestSuite) TestFindNextTick0To1WithMinLiq() {
	// WHEN tick with token1 @ index 1 & MaxTick = 1
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	s.setLPAtFee0Pool(-1, 0, 0)
	s.setLPAtFee0Pool(1, 0, 10)

	// tick 2: (0, 10)
	pair.CurrentTick0To1 = -1
	pair.MaxTick = 2
	s.app.DexKeeper.SetTradingPair(s.ctx, pair)

	// THEN FindNextTick0To1 finds the tick at 1

	tickIdx, found := s.app.DexKeeper.FindNextTick0To1(s.goCtx, pair)
	s.Require().True(found)
	s.Assert().Equal(int64(2), tickIdx)
}

// CalcTrueAmounts ////////////////////////////////////////////////////////////

// func (s *MsgServerTestSuite) TestCalcTrueAmountsEmptyPoolBothSides(){
// 	// WHEN deposit into an empty pool
// 	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(sdk.NewDec(1), sdk.NewDec(0), sdk.NewDec(0), sdk.NewDec(50), sdk.NewDec(20))

// 	// THEN both amounts are used fully

// 	s.Assert().Equal(sdk.NewDec(50), amount0)
// 	s.Assert().Equal(sdk.NewDec(20), amount1)
// 	s.Assert().Equal(sdk.NewDec(70), sharesMinted)
// }

// func (s *MsgServerTestSuite) TestCalcTrueAmountsEmptyPoolToken0(){
// 	// WHEN deposit only Token0 into an empty pool
// 	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(sdk.NewDec(2), sdk.NewDec(0), sdk.NewDec(0), sdk.NewDec(50), sdk.NewDec(0))

// 	// THEN all of Token0 is used

// 	s.Assert().Equal(sdk.NewDec(50), amount0)
// 	s.Assert().Equal(sdk.NewDec(0), amount1)
// 	s.Assert().Equal(sdk.NewDec(50), sharesMinted)
// }

// func (s *MsgServerTestSuite) TestCalcTrueAmountsEmptyPoolToken1(){
// 	// WHEN deposit only Token1 into an empty pool
// 	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(sdk.NewDec(2), sdk.NewDec(0), sdk.NewDec(0), sdk.NewDec(0), sdk.NewDec(50))

// 	// THEN all of Token1 is used

// 	s.Assert().Equal(sdk.NewDec(0), amount0)
// 	s.Assert().Equal(sdk.NewDec(50), amount1)
// 	s.Assert().Equal(sdk.NewDec(100), sharesMinted)
// }

// func (s *MsgServerTestSuite) TestCalcTrueAmounts2SidedPoolBothSidesRightRatio(){
// 	// WHEN deposit into a pool with a ratio of 2:5 with the same ratio all of the tokens are used
// 	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(sdk.NewDec(2), sdk.NewDec(20), sdk.NewDec(50), sdk.NewDec(4), sdk.NewDec(10))

// 	// THEN both amounts are fully user

// 	s.Assert().Equal(sdk.NewDec(4), amount0)
// 	s.Assert().Equal(sdk.NewDec(10), amount1)
// 	s.Assert().Equal(sdk.NewDec(24), sharesMinted)
// }

// func (s *MsgServerTestSuite) TestCalcTrueAmounts2SidedPoolBothSidesWrongRatio(){
// 	// WHEN deposit into a pool with a ratio of 3:2 with a ratio of 2:1
// 	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(sdk.NewDec(3), sdk.NewDec(3), sdk.NewDec(2), sdk.NewDec(20), sdk.NewDec(10))

// 	// THEN all of Token1 is used and 3/4 of token0 is used

// 	s.Assert().Equal(sdk.NewDec(15), amount0)
// 	s.Assert().Equal(sdk.NewDec(10), amount1)
// 	s.Assert().Equal(sdk.NewDec(45), sharesMinted)
// }

// func (s *MsgServerTestSuite) TestCalcTrueAmounts2SidedPoolBothSidesWrongRatio2(){
// 	// IF deposit into a pool with a ratio of 2:3 with a ratio of 1:2
// 	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(sdk.NewDec(3), sdk.NewDec(2), sdk.NewDec(3), sdk.NewDec(10), sdk.NewDec(20))

// 	// THEN all of Token0 is used and 3/4 of token1 is used

// 	s.Assert().Equal(sdk.NewDec(10), amount0)
// 	s.Assert().Equal(sdk.NewDec(15), amount1)
// 	s.Assert().Equal(sdk.NewDec(55), sharesMinted)
// }

// func (s *MsgServerTestSuite) TestCalcTrueAmounts1SidedPoolBothSides(){
// 	// WHEN deposit Token0 and Token1 into a pool with only Token0
// 	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(sdk.NewDec(3), sdk.NewDec(10), sdk.NewDec(0), sdk.NewDec(10), sdk.NewDec(10))

// 	// THEN only Token0 is used

// 	s.Assert().Equal(sdk.NewDec(10), amount0)
// 	s.Assert().Equal(sdk.NewDec(0), amount1)
// 	s.Assert().Equal(sdk.NewDec(10), sharesMinted)
// }

// func (s *MsgServerTestSuite) TestCalcTrueAmounts1SidedPoolBothSides2(){
// 	// WHEN deposit Token0 and Token1 into a pool with only Token1
// 	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(sdk.NewDec(3), sdk.NewDec(0), sdk.NewDec(10), sdk.NewDec(10), sdk.NewDec(10))

// 	// THEN only Token1 is used

// 	s.Assert().Equal(sdk.NewDec(0), amount0)
// 	s.Assert().Equal(sdk.NewDec(10), amount1)
// 	s.Assert().Equal(sdk.NewDec(30), sharesMinted)
// }

// func (s *MsgServerTestSuite) TestCalcTrueAmounts1SidedPool1SidedToken0(){
// 	// WHEN deposit Token0 into a pool with only Token1
// 	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(sdk.NewDec(3), sdk.NewDec(0), sdk.NewDec(10), sdk.NewDec(10), sdk.NewDec(0))

// 	// THEN no amounts are used

// 	s.Assert().Equal(sdk.NewDec(0), amount0)
// 	s.Assert().Equal(sdk.NewDec(0), amount1)
// 	s.Assert().Equal(sdk.NewDec(0), sharesMinted)
// }

// func (s *MsgServerTestSuite) TestCalcTrueAmounts1SidedPool1SidedToken0B(){
// 	// WHEN deposit Token0 into a pool with only Token0
// 	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(sdk.NewDec(3), sdk.NewDec(10), sdk.NewDec(0), sdk.NewDec(10), sdk.NewDec(0))

// 	// THEN all of Token0 is used

// 	s.Assert().Equal(sdk.NewDec(10), amount0)
// 	s.Assert().Equal(sdk.NewDec(0), amount1)
// 	s.Assert().Equal(sdk.NewDec(10), sharesMinted)
// }

// func (s *MsgServerTestSuite) TestCalcTrueAmounts1SidedPool1SidedToken1(){
// 	// WHEN deposit Token1 into a pool with only Token0
// 	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(sdk.NewDec(3), sdk.NewDec(10), sdk.NewDec(0), sdk.NewDec(0), sdk.NewDec(1))

// 	// THEN no amounts are used

// 	s.Assert().Equal(sdk.NewDec(0), amount0)
// 	s.Assert().Equal(sdk.NewDec(0), amount1)
// 	s.Assert().Equal(sdk.NewDec(0), sharesMinted)
// }

// func (s *MsgServerTestSuite) TestCalcTrueAmounts1SidedPool1SidedToken1B(){
// 	// WHEN deposit Token1 into a pool with only Token1
// 	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(sdk.NewDec(4), sdk.NewDec(0), sdk.NewDec(10), sdk.NewDec(0), sdk.NewDec(10))

// 	// THEN all of Token1 is used

// 	s.Assert().Equal(sdk.NewDec(0), amount0)
// 	s.Assert().Equal(sdk.NewDec(10), amount1)
// 	s.Assert().Equal(sdk.NewDec(40), sharesMinted)
// }

// Calc_price_0to1 ////////////////////////////////////////////////////////////

// func (s *MsgServerTestSuite) TestCalc_price_1to0(){
// 	price := s.app.DexKeeper.Calc_price_1to0(0)
// 	expected, _ := sdk.NewDecFromStr("1.0")

// 	s.Assert().Equal(expected, price)

// 	price= s.app.DexKeeper.Calc_price_1to0(1)
// 	expected, _ = sdk.NewDecFromStr("1.0001")

// 	s.Assert().Equal(expected, price)

// 	// TODO: ADD MORE ITAMAR
// }

// Calc_price_1to0 ////////////////////////////////////////////////////////////

// func (s *MsgServerTestSuite) TestCalc_price_0to1(){
// 	price := s.app.DexKeeper.Calc_price_0to1(0)
// 	expected, _ := sdk.NewDecFromStr("1.0")

// 	s.Assert().Equal(expected, price)

// 	price = s.app.DexKeeper.Calc_price_0to1(1)
// 	expected, _ = sdk.NewDecFromStr("0.9999000099990001")

// 	s.Assert().Equal(expected, price)
// 	// TODO: ADD MORE ITAMAR
// }

// HasToken0 //////////////////////////////////////////////////////////////////

func (s *MsgServerTestSuite) TestHasToken0Empty() {

	// WHEN tick only has limit orders and reserves of Token1
	tick, err := s.app.DexKeeper.GetOrInitTick(s.goCtx, "TokenA<>TokenB", 0)
	s.Assert().NoError(err)
	tick.TickData.Reserve1[0] = sdk.NewInt(10)

	tick.LimitOrderTranche0To1.FillTrancheIndex = 0
	tranche := s.app.DexKeeper.GetOrInitLimitOrderTranche(s.ctx, "TokenA<>TokenB", 0, "TokenB", 0)
	tranche.ReservesTokenIn = sdk.NewInt(100)
	s.app.DexKeeper.SetLimitOrderTranche(s.ctx, tranche)

	s.Assert().False(s.app.DexKeeper.TickHasToken0(s.ctx, &tick))
}

func (s *MsgServerTestSuite) TestHasToken0HasReserves() {

	// WHEN tick has Reserves0
	tick, err := s.app.DexKeeper.GetOrInitTick(s.goCtx, "TokenA<>TokenB", 0)
	s.Assert().NoError(err)
	tick.TickData.Reserve0[3] = sdk.NewInt(10)

	s.Assert().True(s.app.DexKeeper.TickHasToken0(s.ctx, &tick))
}

func (s *MsgServerTestSuite) TestHasToken0HasLimitOrders() {

	// WHEN there are limit orders at the tick
	tick, err := s.app.DexKeeper.GetOrInitTick(s.goCtx, "TokenA<>TokenB", 0)
	s.Assert().NoError(err)
	tick.LimitOrderTranche0To1.FillTrancheIndex = 0
	tranche := s.app.DexKeeper.GetOrInitLimitOrderTranche(s.ctx, "TokenA<>TokenB", 0, "TokenA", 0)
	tranche.ReservesTokenIn = sdk.NewInt(100)
	s.app.DexKeeper.SetLimitOrderTranche(s.ctx, tranche)

	s.Assert().True(s.app.DexKeeper.TickHasToken0(s.ctx, &tick))
}

// HasToken1 //////////////////////////////////////////////////////////////////

func (s *MsgServerTestSuite) TestHasToken1Empty() {

	// WHEN tick only has limit orders and reserves of Token0
	tick, err := s.app.DexKeeper.GetOrInitTick(s.goCtx, "TokenA<>TokenB", 0)
	s.Assert().NoError(err)
	tick.TickData.Reserve0[0] = sdk.NewInt(10)

	tick.LimitOrderTranche0To1.FillTrancheIndex = 0
	tranche := s.app.DexKeeper.GetOrInitLimitOrderTranche(s.ctx, "TokenA<>TokenB", 0, "TokenA", 0)
	tranche.ReservesTokenIn = sdk.NewInt(100)
	s.app.DexKeeper.SetLimitOrderTranche(s.ctx, tranche)

	s.Assert().False(s.app.DexKeeper.TickHasToken1(s.ctx, &tick))
}

func (s *MsgServerTestSuite) TestHasToken1HasReserves() {

	// WHEN tick has Reserves0
	tick, err := s.app.DexKeeper.GetOrInitTick(s.goCtx, "TokenA<>TokenB", 0)
	s.Assert().NoError(err)
	tick.TickData.Reserve1[0] = sdk.NewInt(10)

	// THEN HasToken1() = true
	s.Assert().True(s.app.DexKeeper.TickHasToken1(s.ctx, &tick))
}

func (s *MsgServerTestSuite) TestHasToken1HasLimitOrders() {

	// WHEN there are limit orders at the tick
	tick, err := s.app.DexKeeper.GetOrInitTick(s.goCtx, "TokenA<>TokenB", 0)
	s.Assert().NoError(err)
	tick.LimitOrderTranche0To1.FillTrancheIndex = 0
	tranche := s.app.DexKeeper.GetOrInitLimitOrderTranche(s.ctx, "TokenA<>TokenB", 0, "TokenB", 0)
	tranche.ReservesTokenIn = sdk.NewInt(100)
	s.app.DexKeeper.SetLimitOrderTranche(s.ctx, tranche)

	s.Assert().True(s.app.DexKeeper.TickHasToken1(s.ctx, &tick))
}
