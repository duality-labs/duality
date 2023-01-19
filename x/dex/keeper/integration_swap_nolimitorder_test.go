package keeper_test

import (
	"github.com/NicholasDotSol/duality/x/dex/keeper"
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (s *MsgServerTestSuite) TestSwapNoLONoLiquidity() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 50)
	// GIVEN
	// no liqudity of token A (deposit only token B at tick 0 fee 1)
	s.aliceDeposits(NewDeposit(0, 10, 0, 0))
	s.assertAliceBalances(50, 40)

	// WHEN
	// swap 5 of tokenB
	// THEN
	// swap should fail with Error Not enough coins
	err := types.ErrSlippageLimitReached
	s.bobMarketSellFails(err, "TokenB", 5, 0)
}

func (s *MsgServerTestSuite) TestSwapNoLOPartiallyFilledSlippageToleranceNotReachedMaxReached() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// deposit 10 of tokenB
	s.aliceDeposits(NewDeposit(0, 10, 0, 0))
	s.assertAliceBalances(50, 40)
	s.assertDexBalances(0, 10)
	s.assertPoolLiquidity(0, 10, 0, 0)
	//
	// WHEN
	// swap 20 of tokenA at
	amountIn := 20
	s.bobMarketSells("TokenA", amountIn, 5)

	// THEN
	// swap should have in 10 out 10
	s.assertBobBalances(40, 10)
	s.assertDexBalances(10, 0)
	// TODO: this test case is acceptable but succeptible to DOSing by dusting many ticks with large distances between them
}

func (s *MsgServerTestSuite) TestSwapNoLOPartiallyFilledSlippageToleranceNotReachedMinReached() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(0, 50)
	// GIVEN
	// deposit 10 of token A at tick 0 fee 1
	s.aliceDeposits(NewDeposit(10, 0, 0, 0))
	s.assertAliceBalances(40, 50)
	s.assertDexBalances(10, 0)
	s.assertPoolLiquidity(10, 0, 0, 0)
	//
	// WHEN
	// swap 20 of token A for B
	amountIn := 20
	s.bobMarketSells("TokenB", amountIn, 5)

	// THEN
	// swap should have in 10 out 10
	s.assertBobBalances(10, 40)
	s.assertDexBalances(0, 10)
	// TODO: this test case is acceptable but succeptible to DOSing by dusting many ticks with large distances between them
}

func (s *MsgServerTestSuite) TestSwapNoLOSlippageToleranceReached() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// deposit 10 of tokenB
	s.aliceDeposits(
		NewDeposit(0, 10, 0, 0),
		NewDeposit(0, 10, 100000, 1),
	)
	s.assertAliceBalances(50, 30)
	s.assertDexBalances(0, 20)

	// WHEN
	// swap 20 of token A for B with minOut 15
	// THEN
	// swap should fail with ErrNotEnoughCoins error
	err := types.ErrSlippageLimitReached
	s.bobMarketSellFails(err, "TokenA", 20, 19)
}

func (s *MsgServerTestSuite) TestSwapNoLOPartiallyFilledSlippageToleranceReachedMinReached() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// deposit 10 of token B at tick 0 fee 1
	s.aliceDeposits(NewDeposit(10, 0, 0, 0))
	s.assertAliceBalances(40, 50)
	s.assertDexBalances(10, 0)

	// WHEN
	// swap 20 of token A for B with minOut 15
	// THEN
	// swap should fail with ErrNotEnoughCoins error
	err := types.ErrSlippageLimitReached
	s.bobMarketSellFails(err, "TokenB", 20, 15)
}

func (s *MsgServerTestSuite) TestSwapNoLOPartiallyFilledSlippageToleranceReachedMaxReached() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// deposit 10 of token B at tick 0 fee 1
	s.aliceDeposits(NewDeposit(0, 10, 0, 0))
	s.assertAliceBalances(50, 40)
	s.assertDexBalances(0, 10)

	// WHEN
	// swap 20 of token A for B with minOut 15
	// THEN
	// swap should fail with ErrNotEnoughCoins error
	err := types.ErrSlippageLimitReached
	s.bobMarketSellFails(err, "TokenA", 20, 15)
}

func (s *MsgServerTestSuite) TestSwapNoLOCorrectExecutionMinFeeTier() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// deposit 10 of token B at tick 0 fee 1
	s.aliceDeposits(NewDeposit(0, 10, 0, 0))
	s.assertAliceBalances(50, 40)
	s.assertDexBalances(0, 10)

	// WHEN
	// swap 5 of token A for B with minOut 4
	amountIn := 5
	s.bobMarketSells("TokenA", amountIn, 4)

	// THEN
	// swap should have in 5 out 4
	s.assertBobBalances(45, 4)
	s.assertDexBalances(5, 6)
}

func (s *MsgServerTestSuite) TestSwapNoLOCorrectExecutionMaxFeeTier() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// deposit 10 of token B at tick 0 fee 10
	s.aliceDeposits(NewDeposit(0, 10, 0, len(s.feeTiers)-1))
	s.assertAliceBalances(50, 40)
	s.assertDexBalances(0, 10)

	// WHEN
	// swap 5 of token A for B with minOut 4
	amountIn := 5
	s.bobMarketSells("TokenA", amountIn, 4)

	// THEN
	// swap should have in 5 out 4
	s.assertBobBalances(45, 4)
	s.assertDexBalances(5, 6)
}

func (s *MsgServerTestSuite) TestSwapNoLOCorrectExecutionSomeFeeTiers() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// deposit 10 of token B at tick 0 fee 1 and 10 of token B at tick 0 fee 3
	s.aliceDeposits(
		NewDeposit(0, 10, 0, 0),
		NewDeposit(0, 10, 0, 1),
	)
	s.assertAliceBalances(50, 30)
	s.assertDexBalances(0, 20)

	// WHEN
	// swap 5 of token A for B with minOut 4
	amountIn := 15
	s.bobMarketSells("TokenA", amountIn, 14)

	// THEN
	// swap should have in 15 out 14
	s.assertBobBalances(35, 14)
	s.assertDexBalances(15, 6)
}

func (s *MsgServerTestSuite) TestSwapNoLO1to0DoesntMoveCurr1to0() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(0, 50)
	// GIVEN
	// deposit 10 of token A at tick 0 fee 1
	s.aliceDeposits(NewDeposit(10, 0, 0, 0))
	s.assertCurr1To0(-1)

	// WHEN
	// swap 5 of token B for A with minOut 4
	s.bobMarketSells("TokenB", 5, 4)

	// THEN
	// current1To0 unchanged
	s.assertCurr1To0(-1)
}

func (s *MsgServerTestSuite) TestSwapNoLO1to0MovesCurr1to0() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(0, 50)
	// GIVEN
	// deposit 10 of token A at tick 0 fee 1, 10 of token A at tick 0 fee 3
	s.aliceDeposits(
		NewDeposit(10, 0, 0, 0),
		NewDeposit(10, 0, 0, 1),
	)
	s.assertCurr1To0(-1)

	// WHEN
	// swap 15 of token B for A with minOut 14
	s.bobMarketSells("TokenB", 15, 14)

	// THEN
	// current 1to0 moves to -3
	s.assertCurr1To0(-3)
}

func (s *MsgServerTestSuite) TestSwapNoLO1to0DoesntMoveCurr0to1() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(0, 50)
	// GIVEN
	// deposit 10 of both token A and B at tick 0 fee 1
	s.aliceDeposits(NewDeposit(10, 10, 0, 0))
	s.assertCurr0To1(1)

	// WHEN
	// swap 5 of token B for A with minOut 4
	s.bobMarketSells("TokenB", 5, 4)

	// THEN
	// current 0to1 unchanged
	s.assertCurr0To1(1)
}

func (s *MsgServerTestSuite) TestSwapNoLO1to0MovesCurr0to1() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(0, 50)
	// GIVEN
	// deposit 10 of token A at tick 0 fee 1 and 10 of both token A and B at tick 0 fee 3
	s.aliceDeposits(
		NewDeposit(10, 0, 0, 0),
		NewDeposit(10, 10, 0, 1),
	)
	s.assertCurr0To1(3)

	// WHEN
	// swap 5 of token B for A with minOut 4
	s.bobMarketSells("TokenB", 5, 4)

	// THEN
	// current 0to1 moves down to 1
	s.assertCurr0To1(1)
}

func (s *MsgServerTestSuite) TestSwapNoLO1to0DoesntMoveMin() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(0, 50)
	// GIVEN
	// deposit 10 of token A at tick 0 fee 1
	s.aliceDeposits(NewDeposit(10, 0, 0, 0))

	// WHEN
	// swap 5 of token B for A with minOut 4
	s.bobMarketSells("TokenB", 5, 4)

	// THEN
	// current1To0 unchanged
}

func (s *MsgServerTestSuite) TestSwapNoLO1to0ExhaustMin() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(0, 50)
	// GIVEN
	// deposit 10 of token A at tick 0 fee 1
	s.aliceDeposits(NewDeposit(10, 0, 0, 0))

	// WHEN
	// swap 5 of token B for A with minOut 4
	s.bobMarketSells("TokenB", 15, 10)

	// THEN
	// current1To0 unchanged
}

func (s *MsgServerTestSuite) TestSwapNoLO1to0MovesMaxUp() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(0, 50)
	// GIVEN
	// deposit 10 of token B at tick 0 fee 1 and 10 of token A at tick 0 fee 3
	s.aliceDeposits(
		NewDeposit(0, 10, 0, 0),
		NewDeposit(10, 0, 0, 1),
	)

	// WHEN
	// swap 5 of token B for A with minOut 4
	s.bobMarketSells("TokenB", 5, 4)

	// THEN
	// max tick moved up to 3
}

func (s *MsgServerTestSuite) TestSwapNoLO1to0DoesntMoveMaxUp() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(0, 50)
	// GIVEN
	// deposit 10 of token B at tick 0 fee 1
	// 		   10 of token A at tick 0 fee 3
	//         10 of token A and B at tick 0 fee 5
	s.aliceDeposits(
		NewDeposit(0, 10, 0, 0),
		NewDeposit(10, 0, 0, 1),
		NewDeposit(10, 10, 0, 2),
	)

	// WHEN
	// swap 5 of token B for A with minOut 4
	s.bobMarketSells("TokenB", 5, 4)

	// THEN
	// max unchanged
}

func (s *MsgServerTestSuite) TestSwapNoLO0to1DoesntMoveCurr0to1() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// deposit 10 of token B at tick 0 fee 1
	s.aliceDeposits(
		NewDeposit(0, 10, 0, 0),
	)
	s.assertCurr0To1(1)

	// WHEN
	// swap 5 of token A for B with minOut 4
	s.bobMarketSells("TokenA", 5, 4)

	// THEN
	// current0To1 unchanged
	s.assertCurr0To1(1)
}

func (s *MsgServerTestSuite) TestSwapNoLO0to1MovesCurr0to1() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// deposit 10 of token B at tick 0 fee 1, 10 of token B at tick 0 fee 3
	s.aliceDeposits(
		NewDeposit(0, 10, 0, 0),
		NewDeposit(0, 10, 0, 1),
	)
	s.assertCurr0To1(1)

	// WHEN
	// swap 15 of token A for B with minOut 14
	s.bobMarketSells("TokenA", 15, 14)

	// THEN
	// current 0to1 moves to 3
	s.assertCurr0To1(3)
}

func (s *MsgServerTestSuite) TestSwapNoLO0to1DoesntMoveCurr1to0() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// deposit 10 of both token A and B at tick 0 fee 1
	s.aliceDeposits(NewDeposit(10, 10, 0, 0))
	s.assertCurr1To0(-1)

	// WHEN
	// swap 5 of token A for B with minOut 4
	s.bobMarketSells("TokenA", 5, 4)

	// THEN
	// current 1to0 unchanged
	s.assertCurr1To0(-1)
}

func (s *MsgServerTestSuite) TestSwapNoLO0to1MovesCurr1to0() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// deposit 10 of token B at tick 0 fee 1 and 10 of both token A and B at tick 0 fee 3
	// to create spread of -3, 1
	s.aliceDeposits(
		NewDeposit(0, 10, 0, 0),
		NewDeposit(10, 10, 0, 1),
	)
	s.assertCurr1To0(-3)

	// WHEN
	// swap 5 of token B for A with minOut 4
	s.bobMarketSells("TokenA", 5, 4)

	// THEN
	// current 0to1 moves down to 1
	s.assertCurr1To0(-1)
}

func (s *MsgServerTestSuite) TestSwapNoLO0to1DoesntMoveMax() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// deposit 10 of token B at tick 0 fee 1 and 10 of both token A and B at tick 0 fee 3
	// to create spread of -3, 1
	s.aliceDeposits(
		NewDeposit(0, 10, 0, 0),
		NewDeposit(10, 10, 0, 1),
	)
	s.assertCurr1To0(-3)

	// WHEN
	// swap 5 of token B for A with minOut 4
	s.bobMarketSells("TokenA", 5, 4)

	// THEN
	// current 0to1 moves down to 1
	s.assertCurr1To0(-1)
}

func (s *MsgServerTestSuite) TestSwapNoLO0to1ExhaustMax() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// deposit 10 of token B at tick 0 fee 1
	s.aliceDeposits(NewDeposit(0, 10, 0, 0))

	// WHEN
	// swap 5 of token A for B with minOut 4
	s.bobMarketSells("TokenA", 15, 10)

	// THEN
	// current0To1 unchanged
}

func (s *MsgServerTestSuite) TestSwapNoLO0to1MovedMinDown() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// deposit 10 of token A at tick 0 fee 1 and 10 of token B at tick 0 fee 3
	s.aliceDeposits(
		NewDeposit(10, 0, 0, 0),
		NewDeposit(0, 10, 0, 1),
	)

	// WHEN
	// swap 5 of token A for B with minOut 4
	s.bobMarketSells("TokenA", 5, 4)

	// THEN
	// max tick moved up to 3
}

func (s *MsgServerTestSuite) TestSwapNoLO0to1DoesntMoveMinDown() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// deposit 10 of token A at tick 0 fee 1
	// 		   10 of token B at tick 0 fee 3
	//         10 of token A and B at tick 0 fee 5
	s.aliceDeposits(
		NewDeposit(10, 0, 0, 0),
		NewDeposit(0, 10, 0, 1),
		NewDeposit(10, 10, 0, 2),
	)

	// WHEN
	// swap 5 of token A for B with minOut 4
	s.bobMarketSells("TokenA", 5, 4)

	// THEN
	// min unchanged
}

// TODO: 0to1 moves min up
// TODO: 0to1 doesn't move min up

func (s *MsgServerTestSuite) TestSwapNoLOMinLimitTickNotMet() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(0, 50)
	// GIVEN
	// deposit 10 of tokenA
	s.aliceDeposits(NewDeposit(10, 0, 0, 0))
	s.assertAliceBalances(40, 50)
	s.assertDexBalances(10, 0)
	s.assertPoolLiquidity(10, 0, 0, 0)
	//
	// WHEN
	// swap 20 of tokenB at
	amountIn := 10
	amountInInt := sdk.NewInt(10)

	limitPrice, err := keeper.CalcPrice1To0(-10)
	s.Assert().Nil(err)

	s.bobMarketSellsWithLimitPrice("TokenB", amountIn, 5, limitPrice)

	// THEN
	// swap should have in out
	expectedAmountInLeft, expectedAmountOut := s.calculateSingleSwapNoLOBToA(-1, 10, int64(amountIn))
	expectedAmountIn := amountInInt.Sub(expectedAmountInLeft)
	s.assertBobBalancesInt(expectedAmountOut, sdk.NewInt(50).Sub(expectedAmountIn))
	s.assertDexBalancesInt(sdk.NewInt(10).Sub(expectedAmountOut), expectedAmountIn)
	// TODO: this test case is acceptable but succeptible to DOSing by dusting many ticks with large distances between them
}

func (s *MsgServerTestSuite) TestSwapNoLOMaxLimitTickNotMet() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// deposit 10 of tokenB
	s.aliceDeposits(NewDeposit(0, 10, 0, 0))
	s.assertAliceBalances(50, 40)
	s.assertDexBalances(0, 10)
	s.assertPoolLiquidity(0, 10, 0, 0)
	//
	// WHEN
	// swap 20 of tokenA at
	amountIn := 10
	amountInInt := sdk.NewInt(10)

	limitPrice, err := keeper.CalcPrice0To1(10)
	s.Assert().Nil(err)

	s.bobMarketSellsWithLimitPrice("TokenA", amountIn, 5, limitPrice)

	// THEN
	// swap should have in out
	expectedAmountInLeft, expectedAmountOut := s.calculateSingleSwapNoLOAToB(1, 10, int64(amountIn))
	expectedAmountIn := amountInInt.Sub(expectedAmountInLeft)
	s.assertBobBalancesInt(sdk.NewInt(50).Sub(expectedAmountIn), expectedAmountOut)
	s.assertDexBalancesInt(expectedAmountIn, sdk.NewInt(10).Sub(expectedAmountOut))
	// TODO: this test case is acceptable but succeptible to DOSing by dusting many ticks with large distances between them
}

func (s *MsgServerTestSuite) TestSwapNoLOMaxLimitTickMet() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// deposit 10 of tokenB
	s.aliceDeposits(NewDeposit(0, 10, 0, 0), NewDeposit(0, 10, 1, 1))
	s.assertAliceBalances(50, 30)
	s.assertDexBalances(0, 20)
	s.assertPoolLiquidity(0, 10, 0, 0)
	s.assertPoolLiquidity(0, 10, 1, 1)
	//
	// WHEN
	// swap 20 of tokenA at
	amountIn := 10
	amountInInt := sdk.NewInt(10)

	limitPrice, err := keeper.CalcPrice0To1(1)
	s.Assert().Nil(err)

	s.bobMarketSellsWithLimitPrice("TokenA", amountIn, 5, limitPrice)

	// THEN
	// swap should have in out
	expectedAmountInLeft, expectedAmountOut := s.calculateSingleSwapNoLOAToB(1, 10, int64(amountIn))
	expectedAmountIn := amountInInt.Sub(expectedAmountInLeft)
	s.assertBobBalancesInt(sdk.NewInt(50).Sub(expectedAmountIn), expectedAmountOut)
	s.assertDexBalancesInt(expectedAmountIn, sdk.NewInt(20).Sub(expectedAmountOut))
	// TODO: this test case is acceptable but succeptible to DOSing by dusting many ticks with large distances between them
}

func (s *MsgServerTestSuite) TestSwapNoLOMinLimitTickMet() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(0, 50)
	// GIVEN
	// deposit 10 of tokenA
	s.aliceDeposits(NewDeposit(10, 0, 0, 0), NewDeposit(10, 0, -1, 1))
	s.assertAliceBalances(30, 50)
	s.assertDexBalances(20, 0)
	s.assertPoolLiquidity(10, 0, 0, 0)
	s.assertPoolLiquidity(10, 0, -1, 1)
	//
	// WHEN
	// swap 20 of tokenB at
	amountIn := 10
	amountInInt := sdk.NewInt(10)

	limitPrice, err := keeper.CalcPrice1To0(-1)
	s.Assert().Nil(err)

	s.bobMarketSellsWithLimitPrice("TokenB", amountIn, 5, limitPrice)

	// THEN
	// swap should have in out
	expectedAmountInLeft, expectedAmountOut := s.calculateSingleSwapNoLOBToA(-1, 10, int64(amountIn))
	expectedAmountIn := amountInInt.Sub(expectedAmountInLeft)
	s.assertBobBalancesInt(expectedAmountOut, sdk.NewInt(50).Sub(expectedAmountIn))
	s.assertDexBalancesInt(sdk.NewInt(20).Sub(expectedAmountOut), expectedAmountIn)
	// TODO: this test case is acceptable but succeptible to DOSing by dusting many ticks with large distances between them
}

func (s *MsgServerTestSuite) TestSwapNoLOMinLimitTickMetWithPrecisionPrice() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(0, 50)
	// GIVEN
	// deposit 10 of tokenA
	s.aliceDeposits(NewDeposit(10, 0, 0, 0), NewDeposit(10, 0, -1, 1))
	s.assertAliceBalances(30, 50)
	s.assertDexBalances(20, 0)
	s.assertPoolLiquidity(10, 0, 0, 0)
	s.assertPoolLiquidity(10, 0, -1, 1)
	//
	// WHEN
	// swap 20 of tokenB at
	amountIn := 10
	amountInInt := sdk.NewInt(10)

	limitPriceOutsideTickPrecision, err := sdk.NewDecFromStr("0.999900000999000100")
	s.Assert().Nil(err)

	s.bobMarketSellsWithLimitPrice("TokenB", amountIn, 5, limitPriceOutsideTickPrecision)

	// THEN
	// swap should have in out
	expectedAmountInLeft, expectedAmountOut := s.calculateSingleSwapNoLOBToA(-1, 10, int64(amountIn))
	expectedAmountIn := amountInInt.Sub(expectedAmountInLeft)
	s.assertBobBalancesInt(expectedAmountOut, sdk.NewInt(50).Sub(expectedAmountIn))
	s.assertDexBalancesInt(sdk.NewInt(20).Sub(expectedAmountOut), expectedAmountIn)
	// TODO: this test case is acceptable but succeptible to DOSing by dusting many ticks with large distances between them
}

func (s *MsgServerTestSuite) TestSwapNoLOMaxLimitTickMetWithPrecisionPrice() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// deposit 10 of tokenB
	s.aliceDeposits(NewDeposit(0, 10, 0, 0), NewDeposit(0, 10, 1, 1))
	s.assertAliceBalances(50, 30)
	s.assertDexBalances(0, 20)
	s.assertPoolLiquidity(0, 10, 0, 0)
	s.assertPoolLiquidity(0, 10, 1, 1)
	//
	// WHEN
	// swap 20 of tokenA at
	amountIn := 10
	amountInInt := sdk.NewInt(10)

	limitPriceOutsideTickPrecision, err := sdk.NewDecFromStr("0.999900000999000100")
	s.Assert().Nil(err)

	s.bobMarketSellsWithLimitPrice("TokenA", amountIn, 5, limitPriceOutsideTickPrecision)

	// THEN
	// swap should have in out
	expectedAmountInLeft, expectedAmountOut := s.calculateSingleSwapNoLOAToB(1, 10, int64(amountIn))
	expectedAmountIn := amountInInt.Sub(expectedAmountInLeft)
	s.assertBobBalancesInt(sdk.NewInt(50).Sub(expectedAmountIn), expectedAmountOut)
	s.assertDexBalancesInt(expectedAmountIn, sdk.NewInt(20).Sub(expectedAmountOut))
	// TODO: this test case is acceptable but succeptible to DOSing by dusting many ticks with large distances between them
}
