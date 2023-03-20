package keeper_test

import (
	"math"

	"github.com/duality-labs/duality/x/dex/types"
)

func (s *MsgServerTestSuite) TestSwapOnlyLONoLiquidity() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 50)
	// GIVEN
	// no liqudity of token A (place LO only for token B at tick 0 fee 1)
	s.aliceLimitSells("TokenB", 1, 10)

	s.assertLimitLiquidityAtTick("TokenB", 1, 10)
	s.assertAliceBalances(50, 40)
	s.assertDexBalances(0, 10)

	// WHEN
	// swap 5 of tokenB
	// THEN
	// swap should fail with Error Not enough coins
	err := types.ErrInsufficientLiquidity
	s.bobMarketSellFails(err, "TokenB", 5)
}

func (s *MsgServerTestSuite) TestSwapOnlyLOPartiallyFilledSlippageToleranceNotReachedMaxReached() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// place LO selling 10 of token B at tick 1
	s.aliceLimitSells("TokenB", 1, 10)
	s.assertLimitLiquidityAtTick("TokenB", 1, 10)
	s.assertAliceBalances(50, 40)
	s.assertDexBalances(0, 10)

	// WHEN
	// swap 20 of tokenA
	s.bobMarketSells("TokenA", 20)

	// THEN
	// swap should have in 10 out 10
	s.assertBobBalances(39, 10)
	s.assertDexBalances(11, 0)
}

func (s *MsgServerTestSuite) TestSwapOnlyLOPartiallyFilledSlippageToleranceNotReachedMinReached() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(0, 50)
	// GIVEN
	// place LO selling 10 of token A at tick -1
	s.aliceLimitSells("TokenA", -1, 10)
	s.assertLimitLiquidityAtTick("TokenA", -1, 10)
	s.assertAliceBalances(40, 50)
	s.assertDexBalances(10, 0)
	//
	// WHEN
	// swap 20 of token A for B
	s.bobMarketSells("TokenB", 20)

	// THEN
	// swap should have in 10 out 10
	s.assertBobBalances(10, 39)
	s.assertDexBalances(0, 11)
}

func (s *MsgServerTestSuite) TestSwapOnlyLO1to0DoesntMoveCurr1to0() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(0, 50)
	// GIVEN
	// place LO selling 10 of token A at tick -1
	s.aliceLimitSells("TokenA", -1, 10)
	s.assertLimitLiquidityAtTick("TokenA", -1, 10)
	s.assertCurr1To0(-1)

	// WHEN
	// swap 5 of token B for A
	s.bobMarketSells("TokenB", 5)

	// THEN
	// current1To0 unchanged
	s.assertCurr1To0(-1)
}

func (s *MsgServerTestSuite) TestSwapOnlyLO1to0MovesCurr1To0() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(0, 50)
	// GIVEN
	// place LO selling 10 of token A at tick -1 and 10 of token A at tick -3
	s.aliceLimitSells("TokenA", -1, 10)
	s.aliceLimitSells("TokenA", -3, 10)
	s.assertLimitLiquidityAtTick("TokenA", -1, 10)
	s.assertLimitLiquidityAtTick("TokenA", -3, 10)
	s.assertCurr1To0(-1)

	// WHEN
	// swap 15 of token B for A
	s.bobMarketSells("TokenB", 15)

	// THEN
	// current 1to0 moves to -3
	s.assertCurr1To0(-3)
}

// TODO: 1to0 doesn't move curr0to1

func (s *MsgServerTestSuite) TestSwapOnlyLO0to1DoesntMoveCurr0to1() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// place LO selling 10 of token B at tick 1
	s.aliceLimitSells("TokenB", 1, 10)
	s.assertLimitLiquidityAtTick("TokenB", 1, 10)
	s.assertCurr0To1(1)

	// WHEN
	// swap 5 of token A for B
	s.bobMarketSells("TokenA", 5)

	// THEN
	// current0To1 unchanged
	s.assertCurr0To1(1)
}

func (s *MsgServerTestSuite) TestSwapOnlyLO0to1MovesCurr0to1() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// place LO selling 10 of token B at tick 1 and 10 of token B at tick 3
	s.aliceLimitSells("TokenB", 1, 10)
	s.aliceLimitSells("TokenB", 3, 10)
	s.assertLimitLiquidityAtTick("TokenB", 1, 10)
	s.assertLimitLiquidityAtTick("TokenB", 3, 10)
	s.assertCurr0To1(1)

	// WHEN
	// swap 15 of token A for B
	s.bobMarketSells("TokenA", 15)

	// THEN
	// current 0to1 moves to 3
	s.assertCurr0To1(3)
}

// TODO: 0to1 doesn't move curr1to0

func (s *MsgServerTestSuite) TestSwapOnlyLOCorrectExecution1to0() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(0, 50)
	// GIVEN
	// place LO selling 10 of token A at tick 1
	s.aliceLimitSells("TokenA", 1, 10)
	s.assertAliceBalances(40, 50)
	s.assertDexBalances(10, 0)

	// WHEN
	// swap 5 of token B for A
	amountIn := 5
	s.bobMarketSells("TokenB", amountIn)

	// THEN
	// swap should have in 5 out 5
	s.assertBobBalances(5, 45)
	s.assertDexBalances(5, 5)
}

func (s *MsgServerTestSuite) TestSwapOnlyLOCorrectExecution0to1() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// place LO selling 10 of token B at tick 1
	s.aliceLimitSells("TokenB", 1, 10)
	s.assertAliceBalances(50, 40)
	s.assertDexBalances(0, 10)

	// WHEN
	// swap 5 of token A for B
	amountIn := 5
	s.bobMarketSells("TokenA", amountIn)

	// THEN
	// swap should have in 5 out 4
	s.assertBobBalances(45, 4)
	s.assertDexBalances(5, 6)
}

func (s *MsgServerTestSuite) TestSwapOnlyLOPartiallyFilledCorrectExecution() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// place LO selling 10 of token B at tick 1
	s.aliceLimitSells("TokenB", 1, 10)

	// Partially fill the LO, will have some token B remaining to fill
	s.bobMarketSells("TokenA", 5)
	// in 5 out 4
	s.assertLimitLiquidityAtTick("TokenB", 1, 6)

	// place another LO selling 10 of token B at tick 1
	s.aliceLimitSells("TokenB", 1, 10)
	s.assertLimitLiquidityAtTick("TokenB", 1, 16)
	s.assertBobBalances(45, 4)

	// WHEN
	// swap 5 of token A for B
	s.bobMarketSells("TokenA", 5)

	// THEN
	// swap should have in 5 out 4
	s.assertBobBalances(40, 8)
	s.assertDexBalances(10, 12)
	s.assertLimitLiquidityAtTick("TokenB", 1, 12)
}

func (s *MsgServerTestSuite) TestSwapOnlyLOExhaustLOCorrectExecution() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// place LO selling 10 of token B at tick 1
	s.aliceLimitSells("TokenB", 1, 10)

	// Partially fill the LO, will have some token B remaining to fill
	s.bobMarketSells("TokenA", 5)
	// in 5 out 4
	s.assertLimitLiquidityAtTick("TokenB", 1, 6)

	// place another LO selling 10 of token B at tick 1
	s.aliceLimitSells("TokenB", 1, 10)
	s.assertLimitLiquidityAtTick("TokenB", 1, 16)
	s.assertBobBalances(45, 4)

	// WHEN
	// swap 20 of token A for B
	s.bobMarketSells("TokenA", 20)

	// THEN
	// swap should have in 16 out 16
	s.assertBobBalances(27, 20)
	s.assertDexBalances(23, 0)
	s.assertLimitLiquidityAtTick("TokenB", 1, 0)
}

func (s *MsgServerTestSuite) TestSwapOnlyLOPartiallyFilled0to1DoesntMove0to1() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// place LO selling 10 of token B at tick 1
	s.aliceLimitSells("TokenB", 1, 10)

	// Partially fill the LO, will have some token B remaining to fill
	s.bobMarketSells("TokenA", 5)
	// place another LO selling 10 of token B at tick 1
	s.aliceLimitSells("TokenB", 1, 10)
	s.assertCurr0To1(1)

	// WHEN
	// swap 5 of token A for B
	s.bobMarketSells("TokenA", 5)

	// THEN
	// curr0to1 unchanged
	s.assertCurr0To1(1)
}

func (s *MsgServerTestSuite) TestSwapOnlyLOPartiallyFilled1to0DoesntMove1to0() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(0, 50)
	// GIVEN
	// place LO selling 10 of token A at tick -1
	s.aliceLimitSells("TokenA", -1, 10)
	// Partially fill the LO, will have some token A remaining to fill
	s.bobMarketSells("TokenB", 5)
	// place another LO selling 10 of token A at tick -1
	s.aliceLimitSells("TokenA", -1, 10)
	s.assertCurr1To0(-1)

	// WHEN
	// swap 5 of token B for A
	s.bobMarketSells("TokenB", 5)

	// THEN
	// curr0to1 unchanged
	s.assertCurr1To0(-1)
}

func (s *MsgServerTestSuite) TestSwapOnlyLOExhaustFillAndPlace0to1Moves0to1() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// place LO selling 10 of token B at tick 10
	s.aliceLimitSells("TokenB", 10, 10)
	// place LO selling 10 of token B at tick 1
	s.aliceLimitSells("TokenB", 1, 10)
	// Partially fill the LO, will have some token B remaining to fill
	s.bobMarketSells("TokenA", 5)
	// place another LO selling 10 of token B at tick 1
	s.aliceLimitSells("TokenB", 1, 10)
	s.assertCurr0To1(1)

	// WHEN
	// swap 20 of token A for B
	s.bobMarketSells("TokenA", 20)

	// THEN
	// curr0to1 moved to 10
	s.assertCurr0To1(10)
}

func (s *MsgServerTestSuite) TestSwapOnlyLOExhaustFillAndPlace0to1ExhaustMax() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// place LO selling 10 of token B at tick 1
	s.aliceLimitSells("TokenB", 1, 10)
	// Partially fill the LO, will have some token B remaining to fill
	s.bobMarketSells("TokenA", 5)
	// place another LO selling 10 of token A at tick -1
	s.aliceLimitSells("TokenB", 1, 10)
	s.assertCurr0To1(1)

	// WHEN
	// swap 20 of token A for B
	s.bobMarketSells("TokenA", 20)

	// THEN
	// curr0to1 and max set to null values
	s.assertCurr0To1(math.MaxInt64)
}

func (s *MsgServerTestSuite) TestSwapOnlyLOExhaustFillAndPlace1to0Moves1to0() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(0, 50)
	// GIVEN
	// place LO selling 10 of token A at tick -10
	s.aliceLimitSells("TokenA", -10, 10)
	// place LO selling 10 of token B at tick 1
	s.aliceLimitSells("TokenA", -1, 10)
	// Partially fill the LO, will have some token B remaining to fill
	s.bobMarketSells("TokenB", 5)
	// place another LO selling 10 of token A at tick -1
	s.aliceLimitSells("TokenA", -1, 10)
	s.assertCurr1To0(-1)

	// WHEN
	// swap 20 of token B for A
	s.bobMarketSells("TokenB", 20)

	// THEN
	// curr0to1 unchanged
	s.assertCurr1To0(-10)
}

func (s *MsgServerTestSuite) TestSwapOnlyLOExhaustFillAndPlace1to0ExhaustMin() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(0, 50)
	// GIVEN
	// place LO selling 10 of token A at tick -1
	s.aliceLimitSells("TokenA", -1, 10)
	// Partially fill the LO, will have some token A remaining to fill
	s.bobMarketSells("TokenB", 5)
	// place another LO selling 10 of token A at tick -1
	s.aliceLimitSells("TokenA", -1, 10)
	s.assertCurr1To0(-1)

	// WHEN
	// swap 20 of token A for B
	s.bobMarketSells("TokenB", 20)

	// THEN
	// curr0to1 and max set to null values
	s.assertCurr1To0(math.MinInt64)
}

func (s *MsgServerTestSuite) TestSwapOnlyLOUnfilledLOSwapIncrementsFillKey() {
	// TODO: this fails due to fill and place key bug
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(0, 50)
	// GIVEN
	// place LO selling 10 of token A at tick -1
	trancheKey0 := s.aliceLimitSells("TokenA", -1, 10)
	s.assertFillAndPlaceTrancheKeys("TokenA", -1, trancheKey0, trancheKey0)

	// WHEN
	// swap 20 of token A for B and Place a new limitOrder
	s.bobMarketSells("TokenB", 5)
	trancheKey1 := s.aliceLimitSells("TokenA", -1, 10)

	// THEN
	// place increased
	s.assertFillAndPlaceTrancheKeys("TokenA", -1, trancheKey0, trancheKey1)
}
