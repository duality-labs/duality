package keeper_test

import (
	"math"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
	err := types.ErrNotEnoughLiquidity
	s.bobMarketSellFails(err, "TokenB", 5, 0)
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
	//
	// WHEN
	// swap 20 of tokenA at
	amountIn := 20
	amountInDec := sdk.NewDec(20)
	s.bobMarketSells("TokenA", amountIn, 5)

	// THEN
	// swap should have in 10 out
	expectedAmountInLeft, expectedAmountOut := s.calculateSingleSwapNoLOAToB(1, sdk.NewDec(10), amountInDec)
	expectedAmountIn := amountInDec.Sub(expectedAmountInLeft)
	s.assertBobBalancesDec(sdk.NewDec(50).Sub(expectedAmountIn), expectedAmountOut)
	s.assertDexBalancesDec(expectedAmountIn, sdk.NewDec(10).Sub(expectedAmountOut))
	// TODO: this test case is acceptable but succeptible to DOSing by dusting many ticks with large distances between them
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
	amountIn, amountInDec := 20, sdk.NewDec(20)
	s.bobMarketSells("TokenB", amountIn, 5)

	// THEN
	// swap should have in out
	expectedAmountInLeft, expectedAmountOut := s.calculateSingleSwapNoLOBToA(-1, sdk.NewDec(10), amountInDec)
	expectedAmountIn := amountInDec.Sub(expectedAmountInLeft)
	s.assertBobBalancesDec(expectedAmountOut, sdk.NewDec(50).Sub(expectedAmountIn))
	s.assertDexBalancesDec(sdk.NewDec(10).Sub(expectedAmountOut), expectedAmountIn)
	// TODO: this test case is acceptable but succeptible to DOSing by dusting many ticks with large distances between them
}

func (s *MsgServerTestSuite) TestSwapOnlyLOSlippageToleranceReached() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// place LO selling 10 of token B at tick 1 and 10 of token B at tick 100,000
	s.aliceLimitSells("TokenB", 1, 10)
	s.aliceLimitSells("TokenB", 100000, 10)
	s.assertLimitLiquidityAtTick("TokenB", 1, 10)
	s.assertLimitLiquidityAtTick("TokenB", 100000, 10)
	s.assertAliceBalances(50, 30)
	s.assertDexBalances(0, 20)

	// WHEN
	// swap 20 of token A for B with minOut 15
	// THEN
	// swap should fail with ErrNotEnoughCoins error
	err := types.ErrNotEnoughLiquidity
	s.bobMarketSellFails(err, "TokenA", 20, 19)
}

func (s *MsgServerTestSuite) TestSwapOnlyLOPartiallyFilledSlippageToleranceReachedMinReached() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// place LO selling 10 of token A at tick -1
	s.aliceLimitSells("TokenA", -1, 10)
	s.assertLimitLiquidityAtTick("TokenA", -1, 10)
	s.assertAliceBalances(40, 50)
	s.assertDexBalances(10, 0)

	// WHEN
	// swap 20 of token A for B with minOut 15
	// THEN
	// swap should fail with ErrNotEnoughCoins error
	err := types.ErrNotEnoughCoins
	s.bobMarketSellFails(err, "TokenB", 20, 15)
}

func (s *MsgServerTestSuite) TestSwapOnlyLOPartiallyFilledSlippageToleranceReachedMaxReached() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// place LO selling 10 of token B at tick 1
	s.aliceLimitSells("TokenB", 1, 10)
	s.assertLimitLiquidityAtTick("TokenB", 1, 10)
	s.assertAliceBalances(50, 40)
	s.assertDexBalances(0, 10)

	// WHEN
	// swap 20 of token A for B with minOut 15
	// THEN
	// swap should fail with ErrNotEnoughCoins error
	err := types.ErrNotEnoughLiquidity
	s.bobMarketSellFails(err, "TokenA", 20, 15)
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
	// swap 5 of token B for A with minOut 4
	s.bobMarketSells("TokenB", 5, 4)

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
	// swap 15 of token B for A with minOut 14
	s.bobMarketSells("TokenB", 15, 14)

	// THEN
	// current 1to0 moves to -3
	s.assertCurr1To0(-3)
}

// TODO: 1to0 doesn't move curr0to1

func (s *MsgServerTestSuite) TestSwapOnlyLO1to0DoesntMoveMin() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(0, 50)
	// GIVEN
	// place LO selling 10 of token A at tick -1
	s.aliceLimitSells("TokenA", -1, 10)
	s.assertLimitLiquidityAtTick("TokenA", -1, 10)
	s.assertMinTick(-1)

	// WHEN
	// swap 5 of token B for A with minOut 4
	s.bobMarketSells("TokenB", 5, 4)

	// THEN
	// min unchanged
	s.assertMinTick(-1)
}

func (s *MsgServerTestSuite) TestSwapOnlyLO1to0ExhaustMin() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(0, 50)
	// GIVEN
	// place LO selling 10 of token A at tick -1
	s.aliceLimitSells("TokenA", -1, 10)
	s.assertLimitLiquidityAtTick("TokenA", -1, 10)
	s.assertMinTick(-1)

	// WHEN
	// swap 15 of token B for A with minOut 14
	s.bobMarketSells("TokenB", 15, 10)

	// THEN
	// min set to null value
	s.assertMinTick(math.MaxInt64)
}

// TODO: 1to0 doesn't move max

func (s *MsgServerTestSuite) TestSwapOnlyLO0to1DoesntMoveCurr0to1() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// place LO selling 10 of token B at tick 1
	s.aliceLimitSells("TokenB", 1, 10)
	s.assertLimitLiquidityAtTick("TokenB", 1, 10)
	s.assertCurr0To1(1)

	// WHEN
	// swap 5 of token A for B with minOut 4
	s.bobMarketSells("TokenA", 5, 4)

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
	// swap 15 of token A for B with minOut 14
	s.bobMarketSells("TokenA", 15, 14)

	// THEN
	// current 0to1 moves to 3
	s.assertCurr0To1(3)
}

// TODO: 0to1 doesn't move curr1to0

func (s *MsgServerTestSuite) TestSwapOnlyLO0to1DoesntMoveMax() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// place LO selling 10 of token B at tick 1
	s.aliceLimitSells("TokenB", 1, 10)
	s.assertLimitLiquidityAtTick("TokenB", 1, 10)
	s.assertMaxTick(1)

	// WHEN
	// swap 5 of token A for B with minOut 4
	s.bobMarketSells("TokenA", 5, 4)

	// THEN
	// max unchanged
	s.assertMaxTick(1)
}

func (s *MsgServerTestSuite) TestSwapOnlyLO0to1ExhaustMax() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// place LO selling 10 of token B at tick 1
	s.aliceLimitSells("TokenB", 1, 10)
	s.assertLimitLiquidityAtTick("TokenB", 1, 10)
	s.assertMaxTick(1)

	// WHEN
	// swap 15 of token A for B with minOut 14
	s.bobMarketSells("TokenA", 15, 10)

	// THEN
	// max set to null value
	s.assertMaxTick(math.MinInt64)
}

func (s *MsgServerTestSuite) TestSwapOnlyLOCorrectExecution1to0() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(0, 50)
	// GIVEN
	// place LO selling 10 of token A at tick 1
	s.aliceLimitSells("TokenA", 1, 10)
	s.assertAliceBalances(40, 50)
	s.assertDexBalances(10, 0)

	// WHEN
	// swap 5 of token B for A with minOut 4
	amountIn, amountInDec := 5, sdk.NewDec(5)
	s.bobMarketSells("TokenB", amountIn, 4)

	// THEN
	// swap should have in out
	expectedAmountLeft, expectedAmountOut := s.calculateSingleSwapOnlyLOBToA(1, sdk.NewDec(10), amountInDec)
	expectedAmountIn := amountInDec.Sub(expectedAmountLeft)
	s.assertBobBalancesDec(expectedAmountOut, sdk.NewDec(50).Sub(expectedAmountIn))
	s.assertDexBalancesDec(sdk.NewDec(10).Sub(expectedAmountOut), expectedAmountIn)
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
	// swap 5 of token A for B with minOut 4
	amountIn, amountInDec := 5, sdk.NewDec(5)
	s.bobMarketSells("TokenA", amountIn, 4)

	// THEN
	// swap should have in out
	expectedAmountLeft, expectedAmountOut := s.calculateSingleSwapOnlyLOAToB(1, sdk.NewDec(10), amountInDec)
	expectedAmountIn := amountInDec.Sub(expectedAmountLeft)
	s.assertBobBalancesDec(sdk.NewDec(50).Sub(expectedAmountIn), expectedAmountOut)
	s.assertDexBalancesDec(expectedAmountIn, sdk.NewDec(10).Sub(expectedAmountOut))
}

func (s *MsgServerTestSuite) TestSwapOnlyLOPartiallyFilledCorrectExecution() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// place LO selling 10 of token B at tick 1
	s.aliceLimitSells("TokenB", 1, 10)

	// Partially fill the LO, will have some token B remaining to fill
	s.bobMarketSells("TokenA", 5, 4)
	// in 5.000499950004999500, out 4.999500049995000500
	expectedAmountLeftSetup, amountOutSetup := s.calculateSingleSwapOnlyLOAToB(1, sdk.NewDec(10), sdk.NewDec(5))
	amountInSetup := sdk.NewDec(5).Sub(expectedAmountLeftSetup)
	s.assertLimitLiquidityAtTickDec("TokenB", sdk.NewDec(10).Sub(amountOutSetup), 1)

	// place another LO selling 10 of token B at tick 1
	s.aliceLimitSells("TokenB", 1, 10)
	// TODO: uncomment
	// s.assertLimitLiquidityAtTickDec("TokenB", 1, sdk.NewDec(20).Sub(amountOutSetup))
	bobBalanceSetupB := sdk.NewDec(50).Sub(amountInSetup)
	s.assertBobBalancesDec(bobBalanceSetupB, amountOutSetup)

	// WHEN
	// swap 5 of token A for B with minOut 4
	amountIn, amountInDec := 5, sdk.NewDec(5)
	s.bobMarketSells("TokenA", amountIn, 4)

	// THEN
	// swap should have in out
	expectedAmountLeft, expectedAmountOut := s.calculateSingleSwapOnlyLOAToB(1, sdk.NewDec(10), amountInDec)
	expectedAmountIn := amountInDec.Sub(expectedAmountLeft)
	s.assertBobBalancesDec(bobBalanceSetupB.Sub(expectedAmountIn), amountOutSetup.Add(expectedAmountOut))
	s.assertDexBalancesDec(expectedAmountIn.Add(amountInSetup), sdk.NewDec(20).Sub(amountOutSetup).Sub(expectedAmountOut))
	// TODO: uncomment
	// s.assertLimitLiquidityAtTickDec("TokenB", 1, sdk.NewDec(20).Sub(amountOutSetup).Sub(expectedAmountOut))
}

func (s *MsgServerTestSuite) TestSwapOnlyLOExhaustLOCorrectExecution() {
	// TODO: this fails due to fill and place key bug
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// place LO selling 10 of token B at tick 1
	s.aliceLimitSells("TokenB", 1, 10)

	// Partially fill the LO, will have some token B remaining to fill
	s.bobMarketSells("TokenA", 5, 4)
	// in 5.000499950004999500, out 4.999500049995000500
	expectedAmountLeftSetup, amountOutSetup := s.calculateSingleSwapOnlyLOAToB(1, sdk.NewDec(10), sdk.NewDec(5))
	amountInSetup := sdk.NewDec(5).Sub(expectedAmountLeftSetup)
	s.assertLimitLiquidityAtTickDec("TokenB", sdk.NewDec(10).Sub(amountOutSetup), 1)

	// place another LO selling 10 of token B at tick 1
	s.aliceLimitSells("TokenB", 1, 10)
	limitLiquiditySetup := sdk.NewDec(20).Sub(amountOutSetup)
	s.assertLimitLiquidityAtTickDec("TokenB", limitLiquiditySetup, 1)
	bobBalanceSetupB := sdk.NewDec(50).Sub(amountInSetup)
	s.assertBobBalancesDec(bobBalanceSetupB, amountOutSetup)

	// WHEN
	// swap 5 of token A for B with minOut 4
	amountIn, amountInDec := 20, sdk.NewDec(20)
	s.bobMarketSells("TokenA", amountIn, 0)

	// THEN
	// swap should have in out
	expectedAmountLeft, expectedAmountOut := s.calculateSingleSwapOnlyLOAToB(1, limitLiquiditySetup, amountInDec)
	expectedAmountIn := amountInDec.Sub(expectedAmountLeft)
	s.assertBobBalancesDec(bobBalanceSetupB.Sub(expectedAmountIn), amountOutSetup.Add(expectedAmountOut))
	s.assertDexBalancesDec(expectedAmountIn.Add(amountInSetup), limitLiquiditySetup.Sub(expectedAmountOut))
	s.assertLimitLiquidityAtTickDec("TokenB", sdk.NewDec(20).Sub(amountOutSetup).Sub(expectedAmountOut), 1)
}

func (s *MsgServerTestSuite) TestSwapOnlyLOPartiallyFilled0to1DoesntMove0to1() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// place LO selling 10 of token B at tick 1
	s.aliceLimitSells("TokenB", 1, 10)

	// Partially fill the LO, will have some token B remaining to fill
	s.bobMarketSells("TokenA", 5, 4)
	// place another LO selling 10 of token B at tick 1
	s.aliceLimitSells("TokenB", 1, 10)
	s.assertCurr0To1(1)

	// WHEN
	// swap 5 of token A for B with minOut 4
	s.bobMarketSells("TokenA", 5, 4)

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
	s.bobMarketSells("TokenB", 5, 4)
	// place another LO selling 10 of token A at tick -1
	s.aliceLimitSells("TokenA", -1, 10)
	s.assertCurr1To0(-1)

	// WHEN
	// swap 5 of token B for A with minOut 4
	s.bobMarketSells("TokenB", 5, 4)

	// THEN
	// curr0to1 unchanged
	s.assertCurr1To0(-1)
}

func (s *MsgServerTestSuite) TestSwapOnlyLOExhaustFillAndPlace0to1Moves0to1() {
	// TODO: this fails due to fill and place key bug
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// place LO selling 10 of token B at tick -10
	s.aliceLimitSells("TokenB", 10, 10)
	// place LO selling 10 of token B at tick 1
	s.aliceLimitSells("TokenB", 1, 10)
	// Partially fill the LO, will have some token B remaining to fill
	s.bobMarketSells("TokenA", 5, 4)
	// place another LO selling 10 of token A at tick -1
	s.aliceLimitSells("TokenB", 1, 10)
	s.assertCurr0To1(1)

	// WHEN
	// swap 20 of token A for B with minOut 0
	s.bobMarketSells("TokenA", 20, 0)

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
	s.bobMarketSells("TokenA", 5, 4)
	// place another LO selling 10 of token A at tick -1
	s.aliceLimitSells("TokenB", 1, 10)
	s.assertCurr0To1(1)
	s.assertMaxTick(1)

	// WHEN
	// swap 20 of token A for B with minOut 0
	s.bobMarketSells("TokenA", 20, 0)

	// THEN
	// curr0to1 and max set to null values
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
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
	s.bobMarketSells("TokenB", 5, 4)
	// place another LO selling 10 of token A at tick -1
	s.aliceLimitSells("TokenA", -1, 10)
	s.assertCurr1To0(-1)

	// WHEN
	// swap 20 of token B for A with minOut 0
	s.bobMarketSells("TokenB", 20, 0)

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
	s.bobMarketSells("TokenB", 5, 4)
	// place another LO selling 10 of token A at tick -1
	s.aliceLimitSells("TokenA", -1, 10)
	s.assertCurr1To0(-1)
	s.assertMinTick(-1)

	// WHEN
	// swap 20 of token A for B with minOut 0
	s.bobMarketSells("TokenB", 20, 0)

	// THEN
	// curr0to1 and max set to null values
	s.assertCurr1To0(math.MinInt64)
	s.assertMinTick(math.MaxInt64)
}

func (s *MsgServerTestSuite) TestSwapOnlyLOUnfilledLOSwapIncrementsFillKey() {
	// TODO: this fails due to fill and place key bug
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(0, 50)
	// GIVEN
	// place LO selling 10 of token A at tick -1
	s.aliceLimitSells("TokenA", -1, 10)
	s.assertFillAndPlaceTrancheKeys("TokenA", -1, 0, 0)

	// WHEN
	// swap 20 of token A for B with minOut 0
	s.bobMarketSells("TokenB", 5, 0)

	// THEN
	// place increased
	s.assertFillAndPlaceTrancheKeys("TokenA", -1, 0, 0)
}
