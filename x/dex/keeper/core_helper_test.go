package keeper_test

import (
	"math"

	"github.com/NicholasDotSol/duality/x/dex/keeper"
	"github.com/NicholasDotSol/duality/x/dex/types"
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
	pairCount := len(s.app.DexKeeper.GetAllPairMap(s.ctx))
	s.Assert().Equal(1, pairCount)

	pair, foundPair := s.app.DexKeeper.GetPairMap(s.ctx, "TokenA<>TokenB")

	s.Require().True(foundPair)

	s.Assert().Equal(pair.PairId, "TokenA<>TokenB")
	s.Assert().Equal(
		&types.TokenPairType{
			CurrentTick0To1: math.MaxInt64,
			CurrentTick1To0: math.MinInt64,
		},
		pair.TokenPair,
	)
	s.Assert().Equal(math.MaxInt64, int(pair.MinTick))
	s.Assert().Equal(math.MinInt64, int(pair.MaxTick))
}

func (s *MsgServerTestSuite) TestGetOrInitPairExisting() {

	// GIVEN we initialize a pair TokenA/TokenB
	s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")

	// WHEN we update values on that pair
	pair, _ := s.app.DexKeeper.GetPairMap(s.ctx, "TokenA<>TokenB")
	pair.MinTick = 20
	s.app.DexKeeper.SetPairMap(s.ctx, pair)

	// AND try to initialize the same pair again
	newPair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")

	// THEN there is still only 1 pair and it retains the values we set
	pairCount := len(s.app.DexKeeper.GetAllPairMap(s.ctx))
	s.Assert().Equal(1, pairCount)
	s.Assert().Equal(int64(20), newPair.MinTick)
}

// GetOrInitTick //////////////////////////////////////////////////////////////

func (s *MsgServerTestSuite) TestGetOrInitTickNew() {
	// GIVEN we initialize a new tick
	s.app.DexKeeper.GetOrInitTick(s.goCtx, "TokenA<>TokenB", 0)

	// THEN 1 tick is initialized with the correct values

	tickCount := len(s.app.DexKeeper.GetAllTickMap(s.ctx))
	s.Assert().Equal(1, tickCount)

	tick, found := s.app.DexKeeper.GetTickMap(s.ctx, "TokenA<>TokenB", 0)

	s.Require().True(found)

	s.Assert().Equal(tick.PairId, "TokenA<>TokenB")
	s.Assert().Equal(tick.TickIndex, int64(0))
	s.Assert().Equal(feeCount, len(tick.TickData.Reserve0AndShares))
	s.Assert().Equal(
		&types.Reserve0AndSharesType{sdk.ZeroDec(), sdk.ZeroDec()},
		tick.TickData.Reserve0AndShares[0],
	)

	s.Assert().Equal(
		&types.Reserve0AndSharesType{sdk.ZeroDec(), sdk.ZeroDec()},
		tick.TickData.Reserve0AndShares[feeCount-1],
	)

	s.Assert().Equal(
		sdk.ZeroDec(),
		tick.TickData.Reserve1[0],
	)

	s.Assert().Equal(
		sdk.ZeroDec(),
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
	tick, _ := s.app.DexKeeper.GetTickMap(s.ctx, "TokenA<>TokenB", 0)
	tick.TickData.Reserve0AndShares[0] = &types.Reserve0AndSharesType{sdk.NewDec(10), sdk.NewDec(10)}
	s.app.DexKeeper.SetTickMap(s.ctx, "TokenA<>TokenB", tick)

	// AND try to initialize the same tick again
	newTick := s.app.DexKeeper.GetOrInitTick(s.goCtx, "TokenA<>TokenB", 0)

	// THEN there is still only 1 tick and it retains the values we set
	tickCount := len(s.app.DexKeeper.GetAllTickMap(s.ctx))
	s.Assert().Equal(1, tickCount)
	s.Assert().Equal(sdk.NewDec(10), newTick.TickData.Reserve0AndShares[0].Reserve0)
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
	pair.TokenPair.CurrentTick1To0 = 1
	pair.MinTick = -2
	s.app.DexKeeper.SetPairMap(s.ctx, pair)

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
	pair.TokenPair.CurrentTick1To0 = 1
	pair.MinTick = -2
	s.app.DexKeeper.SetPairMap(s.ctx, pair)

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

	pair.TokenPair.CurrentTick0To1 = -1
	pair.MaxTick = 1
	s.app.DexKeeper.SetPairMap(s.ctx, pair)

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
	pair.TokenPair.CurrentTick0To1 = -1
	pair.MaxTick = 2
	s.app.DexKeeper.SetPairMap(s.ctx, pair)

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
	tick := s.app.DexKeeper.GetOrInitTick(s.goCtx, "TokenA<>TokenB", 0)
	tick.TickData.Reserve1[0] = sdk.NewDec(10)

	tick.LimitOrderTranche0To1.FillTrancheIndex = 0
	tranche := s.app.DexKeeper.GetOrInitLimitOrderTranche(s.ctx, "TokenA<>TokenB", 0, "TokenB", 0)
	tranche.ReservesTokenIn = sdk.NewDec(100)
	s.app.DexKeeper.SetLimitOrderTranche(s.ctx, tranche)

	s.Assert().False(s.app.DexKeeper.TickHasToken0(s.ctx, &tick))
}

func (s *MsgServerTestSuite) TestHasToken0HasReserves() {

	// WHEN tick has Reserves0
	tick := s.app.DexKeeper.GetOrInitTick(s.goCtx, "TokenA<>TokenB", 0)
	tick.TickData.Reserve0AndShares[3].Reserve0 = sdk.NewDec(10)

	s.Assert().True(s.app.DexKeeper.TickHasToken0(s.ctx, &tick))
}

func (s *MsgServerTestSuite) TestHasToken0HasLimitOrders() {

	// WHEN there are limit orders at the tick
	tick := s.app.DexKeeper.GetOrInitTick(s.goCtx, "TokenA<>TokenB", 0)
	tick.LimitOrderTranche0To1.FillTrancheIndex = 0
	tranche := s.app.DexKeeper.GetOrInitLimitOrderTranche(s.ctx, "TokenA<>TokenB", 0, "TokenA", 0)
	tranche.ReservesTokenIn = sdk.NewDec(100)
	s.app.DexKeeper.SetLimitOrderTranche(s.ctx, tranche)

	s.Assert().True(s.app.DexKeeper.TickHasToken0(s.ctx, &tick))
}

// HasToken1 //////////////////////////////////////////////////////////////////

func (s *MsgServerTestSuite) TestHasToken1Empty() {

	// WHEN tick only has limit orders and reserves of Token0
	tick := s.app.DexKeeper.GetOrInitTick(s.goCtx, "TokenA<>TokenB", 0)
	tick.TickData.Reserve0AndShares[0].Reserve0 = sdk.NewDec(10)

	tick.LimitOrderTranche0To1.FillTrancheIndex = 0
	tranche := s.app.DexKeeper.GetOrInitLimitOrderTranche(s.ctx, "TokenA<>TokenB", 0, "TokenA", 0)
	tranche.ReservesTokenIn = sdk.NewDec(100)
	s.app.DexKeeper.SetLimitOrderTranche(s.ctx, tranche)

	s.Assert().False(s.app.DexKeeper.TickHasToken1(s.ctx, &tick))
}

func (s *MsgServerTestSuite) TestHasToken1HasReserves() {

	// WHEN tick has Reserves0
	tick := s.app.DexKeeper.GetOrInitTick(s.goCtx, "TokenA<>TokenB", 0)
	tick.TickData.Reserve1[0] = sdk.NewDec(10)

	// THEN HasToken1() = true
	s.Assert().True(s.app.DexKeeper.TickHasToken1(s.ctx, &tick))
}

func (s *MsgServerTestSuite) TestHasToken1HasLimitOrders() {

	// WHEN there are limit orders at the tick
	tick := s.app.DexKeeper.GetOrInitTick(s.goCtx, "TokenA<>TokenB", 0)
	tick.LimitOrderTranche0To1.FillTrancheIndex = 0
	tranche := s.app.DexKeeper.GetOrInitLimitOrderTranche(s.ctx, "TokenA<>TokenB", 0, "TokenB", 0)
	tranche.ReservesTokenIn = sdk.NewDec(100)
	s.app.DexKeeper.SetLimitOrderTranche(s.ctx, tranche)

	s.Assert().True(s.app.DexKeeper.TickHasToken1(s.ctx, &tick))
}

// CalcTickPointersPostAddToken0 //////////////////////////////////////////////

func (s *MsgServerTestSuite) TestCalcTickPointersPostAddToken0NoToken() {
	// GIVEN current tick still has 0 Token0
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	lower, _ := s.setLPAtFee0Pool(1, 0, 0)
	// THEN no changes are made
	updatedPair := s.app.DexKeeper.CalcTickPointersPostAddToken0(s.goCtx, &pair, &lower)
	s.Assert().Nil(updatedPair)
}

func (s *MsgServerTestSuite) TestCalcTickPointersPostAddToken0NoLiq() {
	// GIVEN minTick == MaxInt64 ie. no liquidity of Token0 in pool
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MinTick = math.MaxInt64
	lower, _ := s.setLPAtFee0Pool(2, 10, 0)
	// THEN MinTick and cur1To0 are set to CurrentTick's index
	updatedPair := s.app.DexKeeper.CalcTickPointersPostAddToken0(s.goCtx, &pair, &lower)
	s.Assert().Equal(int64(1), updatedPair.TokenPair.CurrentTick1To0)
	s.Assert().Equal(int64(1), updatedPair.MinTick)
}

func (s *MsgServerTestSuite) TestCalcTickPointersPostAddToken0New1To0() {
	// GIVEN current tick provides Token0 at a higher price (tick 2) ie. spread tightens
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MinTick = -10
	pair.TokenPair.CurrentTick1To0 = 1
	lower, _ := s.setLPAtFee0Pool(3, 10, 0)
	// THEN curTick1To0 is set to CurrentTick's index
	updatedPair := s.app.DexKeeper.CalcTickPointersPostAddToken0(s.goCtx, &pair, &lower)
	s.Assert().Equal(int64(2), updatedPair.TokenPair.CurrentTick1To0)
	s.Assert().Equal(int64(-10), updatedPair.MinTick)
}

func (s *MsgServerTestSuite) TestCalcTickPointersPostAddToken0NewMinTick() {
	// GIVEN current tick provides Token0 at a new lowest price (tick -12)
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MinTick = -10
	pair.TokenPair.CurrentTick1To0 = 2
	lower, _ := s.setLPAtFee0Pool(-11, 10, 0)
	// THEN MinTick is set to CurrentTick's index
	updatedPair := s.app.DexKeeper.CalcTickPointersPostAddToken0(s.goCtx, &pair, &lower)
	s.Assert().Equal(int64(2), updatedPair.TokenPair.CurrentTick1To0)
	s.Assert().Equal(int64(-12), updatedPair.MinTick)
}

func (s *MsgServerTestSuite) TestCalcTickPointersPostAddToken0NoChange() {
	// GIVEN current tick provides Token0 below 1To0 Price
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MinTick = -10
	pair.TokenPair.CurrentTick1To0 = 2
	lower, _ := s.setLPAtFee0Pool(1, 10, 0)
	// THEN no changes are made to MinTick & Cur1To0
	updatedPair := s.app.DexKeeper.CalcTickPointersPostAddToken0(s.goCtx, &pair, &lower)
	s.Assert().Equal(int64(2), updatedPair.TokenPair.CurrentTick1To0)
	s.Assert().Equal(int64(-10), updatedPair.MinTick)
}

// CalcTickPointersPostAddToken1 //////////////////////////////////////////////

func (s *MsgServerTestSuite) TestCalcTickPointersPostAddToken1NoToken() {
	// GIVEN current tick still has NO Token1
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	lower, _ := s.setLPAtFee0Pool(1, 0, 0)
	// THEN no changes are made
	updatedPair := s.app.DexKeeper.CalcTickPointersPostAddToken1(s.goCtx, &pair, &lower)
	s.Assert().Nil(updatedPair)
}

func (s *MsgServerTestSuite) TestCalcTickPointersPostAddToken1NoLiq() {
	// GIVEN maxTick == MinInt64 ie. no liquidity of Token1 in pool
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MaxTick = math.MinInt64
	_, upper := s.setLPAtFee0Pool(1, 0, 10)
	// THEN MinTick and cur0To1 are set to CurrentTick's index
	updatedPair := s.app.DexKeeper.CalcTickPointersPostAddToken1(s.goCtx, &pair, &upper)
	s.Assert().Equal(int64(2), updatedPair.TokenPair.CurrentTick0To1)
	s.Assert().Equal(int64(2), updatedPair.MaxTick)
}

func (s *MsgServerTestSuite) TestCalcTickPointersPostAddToken1New1To0() {
	// GIVEN current tick provides Token1 at a lower price (tick 2) ie. spread tightens
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MaxTick = 10
	pair.TokenPair.CurrentTick0To1 = 3
	_, upper := s.setLPAtFee0Pool(1, 0, 10)
	// THEN curTick1To0 is set to CurrentTick's index
	updatedPair := s.app.DexKeeper.CalcTickPointersPostAddToken1(s.goCtx, &pair, &upper)
	s.Assert().Equal(int64(2), updatedPair.TokenPair.CurrentTick0To1)
	s.Assert().Equal(int64(10), updatedPair.MaxTick)
}

func (s *MsgServerTestSuite) TestCalcTickPointersPostAddToken1NewMaxTick() {
	// GIVEN current tick provides Token1 at a new highest price
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MaxTick = 10
	pair.TokenPair.CurrentTick0To1 = 2
	_, upper := s.setLPAtFee0Pool(11, 0, 10)
	// THEN  MaxTick is set to CurrentTick's index
	updatedPair := s.app.DexKeeper.CalcTickPointersPostAddToken1(s.goCtx, &pair, &upper)
	s.Assert().Equal(int64(2), updatedPair.TokenPair.CurrentTick0To1)
	s.Assert().Equal(int64(12), updatedPair.MaxTick)
}

func (s *MsgServerTestSuite) TestCalcTickPointersPostAddToken1NoChange() {
	// GIVEN current tick provides Token1 above 0To1 Price
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MaxTick = 5
	pair.TokenPair.CurrentTick0To1 = 2
	_, upper := s.setLPAtFee0Pool(3, 0, 10)
	// THEN no changes are made to MinTick & Cur0To1
	updatedPair := s.app.DexKeeper.CalcTickPointersPostAddToken1(s.goCtx, &pair, &upper)
	s.Assert().Equal(int64(2), updatedPair.TokenPair.CurrentTick0To1)
	s.Assert().Equal(int64(5), updatedPair.MaxTick)
}

// CalcTickPointersPostRemoveToken0 ///////////////////////////////////////////

func (s *MsgServerTestSuite) TestCalcTickPointersPostRemoveToken0NoChange() {
	// GIVEN current tick removes liquidity between bounds
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MinTick = -5
	pair.TokenPair.CurrentTick1To0 = 1
	lower, _ := s.setLPAtFee0Pool(0, 0, 0)
	// THEN no changes are made
	updatedPair := s.app.DexKeeper.CalcTickPointersPostRemoveToken0(s.goCtx, &pair, &lower)
	s.Assert().Nil(updatedPair)
}

func (s *MsgServerTestSuite) TestCalcTickPointersPostRemoveToken0NotDrained() {
	// GIVEN current tick still has Token0
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MinTick = -5
	pair.TokenPair.CurrentTick1To0 = -4
	lower, _ := s.setLPAtFee0Pool(-10, 10, 0)
	// THEN no changes are made
	updatedPair := s.app.DexKeeper.CalcTickPointersPostRemoveToken0(s.goCtx, &pair, &lower)
	s.Assert().Nil(updatedPair)
}

func (s *MsgServerTestSuite) TestCalcTickPointersPostRemoveToken0DrainLiq() {
	// GIVEN current tick removes liquidity at MinTick && MinTick == Current1To0
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MinTick = -6
	pair.TokenPair.CurrentTick1To0 = -6
	lower, _ := s.setLPAtFee0Pool(-5, 0, 0)
	// THEN Current0to1 is set to MinInt && MaxTick tick is set to MaxInt64
	updatedPair := s.app.DexKeeper.CalcTickPointersPostRemoveToken0(s.goCtx, &pair, &lower)
	s.Assert().Equal(math.MinInt64, int(updatedPair.TokenPair.CurrentTick1To0))
	s.Assert().Equal(math.MaxInt64, int(updatedPair.MinTick))
}

func (s *MsgServerTestSuite) TestCalcTickPointersPostRemoveToken0MinTick() {
	// GIVEN current tick removes liquidity at MinTick
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MinTick = -5
	pair.TokenPair.CurrentTick1To0 = 0
	lower, _ := s.setLPAtFee0Pool(-4, 0, 0)
	// THEN Current0to1 is unchanged && MinTick tick is increased
	updatedPair := s.app.DexKeeper.CalcTickPointersPostRemoveToken0(s.goCtx, &pair, &lower)
	s.Assert().Equal(0, int(updatedPair.TokenPair.CurrentTick1To0))
	s.Assert().Less(-5, int(updatedPair.MinTick))
}

func (s *MsgServerTestSuite) TestCalcTickPointersPostRemoveToken0CurTick() {
	// GIVEN current tick removes liquidity at the CurrentTick
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	s.setLPAtFee0Pool(-1, 10, 0)
	s.setLPAtFee0Pool(-2, 10, 0)
	pair.MinTick = -3
	pair.TokenPair.CurrentTick1To0 = 0
	lower, _ := s.setLPAtFee0Pool(1, 0, 0)

	// THEN Current1to0 is changed to the next lowest tick (-2)
	updatedPair := s.app.DexKeeper.CalcTickPointersPostRemoveToken0(s.goCtx, &pair, &lower)
	s.Assert().Equal(-2, int(updatedPair.TokenPair.CurrentTick1To0))
	s.Assert().Equal(-3, int(updatedPair.MinTick))
}

// CalcTickPointersPostRemoveToken1 ///////////////////////////////////////////

func (s *MsgServerTestSuite) TestCalcTickPointersPostRemoveToken1NoChange() {
	// GIVEN current tick removes liquidity between bounds
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MaxTick = 5
	pair.TokenPair.CurrentTick0To1 = 1
	_, upper := s.setLPAtFee0Pool(3, 0, 0)
	// THEN no changes are made
	updatedPair := s.app.DexKeeper.CalcTickPointersPostRemoveToken1(s.goCtx, &pair, &upper)
	s.Assert().Nil(updatedPair)
}

func (s *MsgServerTestSuite) TestCalcTickPointersPostRemoveToken1NotDrained() {
	// GIVEN current tick still has Token1
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MaxTick = 5
	pair.TokenPair.CurrentTick0To1 = 4
	_, upper := s.setLPAtFee0Pool(5, 0, 10)
	// THEN no changes are made
	updatedPair := s.app.DexKeeper.CalcTickPointersPostRemoveToken1(s.goCtx, &pair, &upper)
	s.Assert().Nil(updatedPair)
}

func (s *MsgServerTestSuite) TestCalcTickPointersPostRemoveToken1DrainLiq() {
	// GIVEN current tick removes liquidity at MaxTick && MaxTick == Current0To1
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MaxTick = 5
	pair.TokenPair.CurrentTick0To1 = 5
	_, upper := s.setLPAtFee0Pool(4, 0, 0)
	// THEN Current0to1 is set to MaxInt && MaxTick tick is set to MinInt
	updatedPair := s.app.DexKeeper.CalcTickPointersPostRemoveToken1(s.goCtx, &pair, &upper)
	s.Assert().Equal(math.MaxInt64, int(updatedPair.TokenPair.CurrentTick0To1))
	s.Assert().Equal(math.MinInt64, int(updatedPair.MaxTick))
}

func (s *MsgServerTestSuite) TestCalcTickPointersPostRemoveToken1MaxTick() {
	// GIVEN current tick removes liquidity at MaxTick
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MaxTick = 5
	pair.TokenPair.CurrentTick0To1 = 0
	_, upper := s.setLPAtFee0Pool(4, 0, 0)
	// THEN Current0to1 is unchanged && MaxTick tick is decreased
	updatedPair := s.app.DexKeeper.CalcTickPointersPostRemoveToken1(s.goCtx, &pair, &upper)
	s.Assert().Equal(0, int(updatedPair.TokenPair.CurrentTick0To1))
	s.Assert().Greater(5, int(updatedPair.MaxTick))
}

func (s *MsgServerTestSuite) TestCalcTickPointersPostRemoveToken1CurTick() {
	// GIVEN current tick removes liquidity at the CurrentTick
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	s.setLPAtFee0Pool(1, 0, 10)
	s.setLPAtFee0Pool(2, 0, 10)
	s.setLPAtFee0Pool(3, 0, 10)
	pair.MaxTick = 4
	pair.TokenPair.CurrentTick0To1 = 2
	_, upper := s.setLPAtFee0Pool(1, 0, 0)

	// THEN Current0to1 is changed to the next lowest tick (3)
	updatedPair := s.app.DexKeeper.CalcTickPointersPostRemoveToken1(s.goCtx, &pair, &upper)
	s.Assert().Equal(3, int(updatedPair.TokenPair.CurrentTick0To1))
	s.Assert().Equal(4, int(updatedPair.MaxTick))
}
