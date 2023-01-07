package types_test

import (
	"context"
	math "math"
	"testing"

	dualityapp "github.com/NicholasDotSol/duality/app"
	. "github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

type TradingPairTestSuite struct {
	suite.Suite
	app   *dualityapp.App
	ctx   sdk.Context
	goCtx context.Context
}

func TestTradingPairTestSuite(t *testing.T) {
	suite.Run(t, new(TradingPairTestSuite))
}

func (s *TradingPairTestSuite) SetupTest() {
	s.app = dualityapp.Setup(false)
	s.ctx = s.app.BaseApp.NewContext(false, tmproto.Header{})
	s.goCtx = sdk.WrapSDKContext(s.ctx)
	feeTiers := []FeeTier{
		{Id: 0, Fee: 1},
		{Id: 1, Fee: 3},
		{Id: 2, Fee: 5},
		{Id: 3, Fee: 10},
	}

	// Set Fee List
	s.app.DexKeeper.AppendFeeTier(s.ctx, feeTiers[0])
	s.app.DexKeeper.AppendFeeTier(s.ctx, feeTiers[1])
	s.app.DexKeeper.AppendFeeTier(s.ctx, feeTiers[2])
	s.app.DexKeeper.AppendFeeTier(s.ctx, feeTiers[3])
}

func (s *TradingPairTestSuite) setLPAtFee0Pool(tickIndex int64, amountA int, amountB int) (lowerTick Tick, upperTick Tick) {
	pairId := "TokenA<>TokenB"
	// sharesId := fmt.Sprintf("%s%st%df%d", "TokenA", "TokenB", tickIndex, 1)
	// sharesId := keeper.CreateSharesId("TokenA", "TokenB", tickIndex, 0)
	lowerTick, err := s.app.DexKeeper.GetOrInitTick(s.goCtx, pairId, tickIndex-1)
	s.Assert().NoError(err)
	upperTick, err = s.app.DexKeeper.GetOrInitTick(s.goCtx, pairId, tickIndex+1)
	s.Assert().NoError(err)
	// priceCenter1To0, err := keeper.CalcPrice0To1(tickIndex)
	// if err != nil {
	// 	panic(err)
	// }

	amountAInt := sdk.NewInt(int64(amountA))
	amountBInt := sdk.NewInt(int64(amountB))
	lowerTick.TickData.Reserve0[0] = amountAInt
	// totalShares := keeper.CalcShares(amountAInt, amountBInt, priceCenter1To0).TruncateInt()
	// s.app.DexKeeper.MintShares(s.ctx, s.alice, totalShares, sharesId)
	upperTick.TickData.Reserve1[0] = amountBInt
	s.app.DexKeeper.SetTick(s.ctx, pairId, lowerTick)
	s.app.DexKeeper.SetTick(s.ctx, pairId, upperTick)
	return lowerTick, upperTick
}

// CalcTickPointersPostAddToken0 //////////////////////////////////////////////

func (s *TradingPairTestSuite) TestCalcTickPointersPostAddToken0NoToken() {
	// GIVEN current tick still has 0 Token0
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	originalPair := pair
	lower, _ := s.setLPAtFee0Pool(1, 0, 0)
	// THEN no changes are made
	pair.UpdateTickPointersPostAddToken0(s.goCtx, s.app.DexKeeper, &lower)
	s.Assert().Equal(pair, originalPair)
}

func (s *TradingPairTestSuite) TestCalcTickPointersPostAddToken0NoLiq() {
	// GIVEN minTick == MaxInt64 ie. no liquidity of Token0 in pool
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	lower, _ := s.setLPAtFee0Pool(2, 10, 0)
	// THEN MinTick and cur1To0 are set to CurrentTick's index
	pair.UpdateTickPointersPostAddToken0(s.goCtx, s.app.DexKeeper, &lower)
	s.Assert().Equal(int64(1), pair.CurrentTick1To0)
	s.Assert().Equal(int64(1), pair.MinTick)
}

func (s *TradingPairTestSuite) TestCalcTickPointersPostAddToken0New1To0() {
	// GIVEN current tick provides Token0 at a higher price (tick 2) ie. spread tightens
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MinTick = -10
	pair.CurrentTick1To0 = 1
	lower, _ := s.setLPAtFee0Pool(3, 10, 0)
	// THEN curTick1To0 is set to CurrentTick's index
	pair.UpdateTickPointersPostAddToken0(s.goCtx, s.app.DexKeeper, &lower)
	s.Assert().Equal(int64(2), pair.CurrentTick1To0)
	s.Assert().Equal(int64(-10), pair.MinTick)
}

func (s *TradingPairTestSuite) TestCalcTickPointersPostAddToken0NewMinTick() {
	// GIVEN current tick provides Token0 at a new lowest price (tick -12)
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MinTick = -10
	pair.CurrentTick1To0 = 2
	lower, _ := s.setLPAtFee0Pool(-11, 10, 0)
	// THEN MinTick is set to CurrentTick's index
	pair.UpdateTickPointersPostAddToken0(s.goCtx, s.app.DexKeeper, &lower)
	s.Assert().Equal(int64(2), pair.CurrentTick1To0)
	s.Assert().Equal(int64(-12), pair.MinTick)
}

func (s *TradingPairTestSuite) TestCalcTickPointersPostAddToken0NoChange() {
	// GIVEN current tick provides Token0 below 1To0 Price
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MinTick = -10
	pair.CurrentTick1To0 = 2
	lower, _ := s.setLPAtFee0Pool(1, 10, 0)
	// THEN no changes are made to MinTick & Cur1To0
	pair.UpdateTickPointersPostAddToken0(s.goCtx, s.app.DexKeeper, &lower)
	s.Assert().Equal(int64(2), pair.CurrentTick1To0)
	s.Assert().Equal(int64(-10), pair.MinTick)
}

// CalcTickPointersPostAddToken1 //////////////////////////////////////////////

func (s *TradingPairTestSuite) TestCalcTickPointersPostAddToken1NoToken() {
	// GIVEN current tick still has NO Token1
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	lower, _ := s.setLPAtFee0Pool(1, 0, 0)
	// THEN no changes are made
	pair.UpdateTickPointersPostAddToken1(s.goCtx, s.app.DexKeeper, &lower)
	s.Assert().Equal(pair, pair)
}

func (s *TradingPairTestSuite) TestCalcTickPointersPostAddToken1NoLiq() {
	// GIVEN maxTick == MinInt64 ie. no liquidity of Token1 in pool
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MaxTick = math.MinInt64
	_, upper := s.setLPAtFee0Pool(1, 0, 10)
	// THEN MinTick and cur0To1 are set to CurrentTick's index
	pair.UpdateTickPointersPostAddToken1(s.goCtx, s.app.DexKeeper, &upper)
	s.Assert().Equal(int64(2), pair.CurrentTick0To1)
	s.Assert().Equal(int64(2), pair.MaxTick)
}

func (s *TradingPairTestSuite) TestCalcTickPointersPostAddToken1New1To0() {
	// GIVEN current tick provides Token1 at a lower price (tick 2) ie. spread tightens
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MaxTick = 10
	pair.CurrentTick0To1 = 3
	_, upper := s.setLPAtFee0Pool(1, 0, 10)
	// THEN curTick1To0 is set to CurrentTick's index
	pair.UpdateTickPointersPostAddToken1(s.goCtx, s.app.DexKeeper, &upper)
	s.Assert().Equal(int64(2), pair.CurrentTick0To1)
	s.Assert().Equal(int64(10), pair.MaxTick)
}

func (s *TradingPairTestSuite) TestCalcTickPointersPostAddToken1NewMaxTick() {
	// GIVEN current tick provides Token1 at a new highest price
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MaxTick = 10
	pair.CurrentTick0To1 = 2
	_, upper := s.setLPAtFee0Pool(11, 0, 10)
	// THEN  MaxTick is set to CurrentTick's index
	pair.UpdateTickPointersPostAddToken1(s.goCtx, s.app.DexKeeper, &upper)
	s.Assert().Equal(int64(2), pair.CurrentTick0To1)
	s.Assert().Equal(int64(12), pair.MaxTick)
}

func (s *TradingPairTestSuite) TestCalcTickPointersPostAddToken1NoChange() {
	// GIVEN current tick provides Token1 above 0To1 Price
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MaxTick = 5
	pair.CurrentTick0To1 = 2
	_, upper := s.setLPAtFee0Pool(3, 0, 10)
	// THEN no changes are made to MinTick & Cur0To1
	pair.UpdateTickPointersPostAddToken1(s.goCtx, s.app.DexKeeper, &upper)
	s.Assert().Equal(int64(2), pair.CurrentTick0To1)
	s.Assert().Equal(int64(5), pair.MaxTick)
}

// CalcTickPointersPostRemoveToken0 ///////////////////////////////////////////

func (s *TradingPairTestSuite) TestCalcTickPointersPostRemoveToken0NoChange() {
	// GIVEN current tick removes liquidity between bounds
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MinTick = -5
	pair.CurrentTick1To0 = 1
	originalPair := pair
	lower, _ := s.setLPAtFee0Pool(0, 0, 0)
	// THEN no changes are made
	pair.UpdateTickPointersPostRemoveToken0(s.goCtx, s.app.DexKeeper, &lower)
	s.Assert().Equal(pair, originalPair)
}

func (s *TradingPairTestSuite) TestCalcTickPointersPostRemoveToken0NotDrained() {
	// GIVEN current tick still has Token0
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MinTick = -5
	pair.CurrentTick1To0 = -4
	originalPair := pair
	lower, _ := s.setLPAtFee0Pool(-10, 10, 0)
	// THEN no changes are made
	pair.UpdateTickPointersPostRemoveToken0(s.goCtx, s.app.DexKeeper, &lower)
	s.Assert().Equal(pair, originalPair)
}

func (s *TradingPairTestSuite) TestCalcTickPointersPostRemoveToken0DrainLiq() {
	// GIVEN current tick removes liquidity at MinTick && MinTick == Current1To0
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MinTick = -6
	pair.CurrentTick1To0 = -6
	lower, _ := s.setLPAtFee0Pool(-5, 0, 0)
	// THEN Current0to1 is set to MinInt && MaxTick tick is set to MaxInt64
	pair.UpdateTickPointersPostRemoveToken0(s.goCtx, s.app.DexKeeper, &lower)
	s.Assert().Equal(math.MinInt64, int(pair.CurrentTick1To0))
	s.Assert().Equal(math.MaxInt64, int(pair.MinTick))
}

func (s *TradingPairTestSuite) TestCalcTickPointersPostRemoveToken0MinTick() {
	// GIVEN current tick removes liquidity at MinTick
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MinTick = -5
	pair.CurrentTick1To0 = 0
	lower, _ := s.setLPAtFee0Pool(-4, 0, 0)
	// THEN Current0to1 is unchanged && MinTick tick is increased
	pair.UpdateTickPointersPostRemoveToken0(s.goCtx, s.app.DexKeeper, &lower)
	s.Assert().Equal(0, int(pair.CurrentTick1To0))
	s.Assert().Less(-5, int(pair.MinTick))
}

func (s *TradingPairTestSuite) TestCalcTickPointersPostRemoveToken0CurTick() {
	// GIVEN current tick removes liquidity at the CurrentTick
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	s.setLPAtFee0Pool(-1, 10, 0)
	s.setLPAtFee0Pool(-2, 10, 0)
	pair.MinTick = -3
	pair.CurrentTick1To0 = 0
	lower, _ := s.setLPAtFee0Pool(1, 0, 0)

	// THEN Current1to0 is changed to the next lowest tick (-2)
	pair.UpdateTickPointersPostRemoveToken0(s.goCtx, s.app.DexKeeper, &lower)
	s.Assert().Equal(-2, int(pair.CurrentTick1To0))
	s.Assert().Equal(-3, int(pair.MinTick))
}

// CalcTickPointersPostRemoveToken1 ///////////////////////////////////////////

func (s *TradingPairTestSuite) TestCalcTickPointersPostRemoveToken1NoChange() {
	// GIVEN current tick removes liquidity between bounds
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MaxTick = 5
	pair.CurrentTick0To1 = 1
	originalPair := pair
	_, upper := s.setLPAtFee0Pool(3, 0, 0)
	// THEN no changes are made
	pair.UpdateTickPointersPostRemoveToken1(s.goCtx, s.app.DexKeeper, &upper)
	s.Assert().Equal(pair, originalPair)
}

func (s *TradingPairTestSuite) TestCalcTickPointersPostRemoveToken1NotDrained() {
	// GIVEN current tick still has Token1
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MaxTick = 5
	pair.CurrentTick0To1 = 4
	originalPair := pair
	_, upper := s.setLPAtFee0Pool(5, 0, 10)
	// THEN no changes are made
	pair.UpdateTickPointersPostRemoveToken1(s.goCtx, s.app.DexKeeper, &upper)
	s.Assert().Equal(pair, originalPair)
}

func (s *TradingPairTestSuite) TestCalcTickPointersPostRemoveToken1DrainLiq() {
	// GIVEN current tick removes liquidity at MaxTick && MaxTick == Current0To1
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MaxTick = 5
	pair.CurrentTick0To1 = 5
	_, upper := s.setLPAtFee0Pool(4, 0, 0)
	// THEN Current0to1 is set to MaxInt && MaxTick tick is set to MinInt
	pair.UpdateTickPointersPostRemoveToken1(s.goCtx, s.app.DexKeeper, &upper)
	s.Assert().Equal(math.MaxInt64, int(pair.CurrentTick0To1))
	s.Assert().Equal(math.MinInt64, int(pair.MaxTick))
}

func (s *TradingPairTestSuite) TestCalcTickPointersPostRemoveToken1MaxTick() {
	// GIVEN current tick removes liquidity at MaxTick
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MaxTick = 5
	pair.CurrentTick0To1 = 0
	_, upper := s.setLPAtFee0Pool(4, 0, 0)
	// THEN Current0to1 is unchanged && MaxTick tick is decreased
	pair.UpdateTickPointersPostRemoveToken1(s.goCtx, s.app.DexKeeper, &upper)
	s.Assert().Equal(0, int(pair.CurrentTick0To1))
	s.Assert().Greater(5, int(pair.MaxTick))
}

func (s *TradingPairTestSuite) TestCalcTickPointersPostRemoveToken1CurTick() {
	// GIVEN current tick removes liquidity at the CurrentTick
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	s.setLPAtFee0Pool(1, 0, 10)
	s.setLPAtFee0Pool(2, 0, 10)
	s.setLPAtFee0Pool(3, 0, 10)
	pair.MaxTick = 4
	pair.CurrentTick0To1 = 2
	_, upper := s.setLPAtFee0Pool(1, 0, 0)

	// THEN Current0to1 is changed to the next lowest tick (3)
	pair.UpdateTickPointersPostRemoveToken1(s.goCtx, s.app.DexKeeper, &upper)
	s.Assert().Equal(3, int(pair.CurrentTick0To1))
	s.Assert().Equal(4, int(pair.MaxTick))
}

// FindNextTick1To0 ///////////////////////////////////////////////////////////

func (s *TradingPairTestSuite) TestFindNextTick1To0NoLiq() {
	// GIVEN there is no ticks with token0 in the pool
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	// NOTE: Actually adding/funding a tick isn't really neccesary but should make the test less
	// dependent upon current implementation details
	s.setLPAtFee0Pool(1, 0, 10)

	// THEN FindNextTick1To0 doesn't find a tick

	_, found := pair.FindNextTick1To0(s.goCtx, s.app.DexKeeper)
	s.Assert().False(found)

}

func (s *TradingPairTestSuite) TestFindNextTick1To0WithLiq() {
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

	// THEN FindNextTick1To0 finds the tick at -1

	tickIdx, found := pair.FindNextTick1To0(s.goCtx, s.app.DexKeeper)
	s.Require().True(found)
	s.Assert().Equal(int64(-1), tickIdx)

}

func (s *TradingPairTestSuite) TestFindNextTick1To0WithMinLiq() {
	// GIVEN tick with token0 @ index -1 & MinTick = -1
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	s.setLPAtFee0Pool(-1, 10, 0)
	s.setLPAtFee0Pool(1, 0, 0)

	// tick -2: (10, 0)
	pair.CurrentTick1To0 = 1
	pair.MinTick = -2
	s.app.DexKeeper.SetTradingPair(s.ctx, pair)

	// THEN FindNextTick1To0 finds the tick at -2

	tickIdx, found := pair.FindNextTick1To0(s.goCtx, s.app.DexKeeper)
	s.Require().True(found)
	s.Assert().Equal(int64(-2), tickIdx)

}

// FindNextTick0To1 ///////////////////////////////////////////////////////////

func (s *TradingPairTestSuite) TestFindNextTick0To1NoLiq() {
	// GIVEN there are no tick with Token1 in the pool
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	// NOTE: Actually adding/funding a tick isn't really neccesary but should make the test less
	// dependent upon current implementation details
	s.setLPAtFee0Pool(0, 10, 0)

	// THEN FindNextTick0To1 doesn't find a tick

	_, found := pair.FindNextTick0To1(s.goCtx, s.app.DexKeeper)
	s.Assert().False(found)

}

func (s *TradingPairTestSuite) TestFindNextTick0To1WithLiq() {
	// WHEN tick with token1 @ index 0 & currentTick0To1 is -1
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	s.setLPAtFee0Pool(-1, 10, 0)
	s.setLPAtFee0Pool(0, 0, 10)
	s.setLPAtFee0Pool(1, 0, 10)

	pair.CurrentTick0To1 = -1
	pair.MaxTick = 1
	s.app.DexKeeper.SetTradingPair(s.ctx, pair)

	// THEN FindNextTick0To1 finds the tick at 1

	tickIdx, found := pair.FindNextTick0To1(s.goCtx, s.app.DexKeeper)
	s.Require().True(found)
	s.Assert().Equal(int64(1), tickIdx)
}

func (s *TradingPairTestSuite) TestFindNextTick0To1WithMinLiq() {
	// WHEN tick with token1 @ index 1 & MaxTick = 1
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	s.setLPAtFee0Pool(-1, 0, 0)
	s.setLPAtFee0Pool(1, 0, 10)

	// tick 2: (0, 10)
	pair.CurrentTick0To1 = -1
	pair.MaxTick = 2
	s.app.DexKeeper.SetTradingPair(s.ctx, pair)

	// THEN FindNextTick0To1 finds the tick at 2

	tickIdx, found := pair.FindNextTick0To1(s.goCtx, s.app.DexKeeper)
	s.Require().True(found)
	s.Assert().Equal(int64(2), tickIdx)
}

// PairIdToTokens ///////////////////////////////////////////////////////////////

func (s *TradingPairTestSuite) TestPairToTokens() {

	token0, token1 := PairIdToTokens("TokenA<>TokenB")

	s.Assert().Equal("TokenA", token0)
	s.Assert().Equal("TokenB", token1)

}

func (s *TradingPairTestSuite) TestPairToTokensIBCis0() {

	token0, token1 := PairIdToTokens("ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2<>TokenB")

	s.Assert().Equal("ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2", token0)
	s.Assert().Equal("TokenB", token1)
}

func (s *TradingPairTestSuite) TestPairToTokensIBCis1() {

	token0, token1 := PairIdToTokens("TokenA<>ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2")

	s.Assert().Equal("TokenA", token0)
	s.Assert().Equal("ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2", token1)

}
func (s *TradingPairTestSuite) TestPairToTokensIBCisBoth() {

	token0, token1 :=
		PairIdToTokens("ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2<>ibc/94644FB092D9ACDA56123C74F36E4234926001AA44A9CA97EA622B25F41E5223")

	s.Assert().Equal("ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2", token0)
	s.Assert().Equal("ibc/94644FB092D9ACDA56123C74F36E4234926001AA44A9CA97EA622B25F41E5223", token1)
}
