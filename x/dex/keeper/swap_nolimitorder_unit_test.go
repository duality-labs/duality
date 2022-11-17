package keeper_test

import (
	. "github.com/NicholasDotSol/duality/x/dex/keeper/internal/testutils"
	"github.com/NicholasDotSol/duality/x/dex/types"
)

func (s *MsgServerTestSuite) TestSwapNoLONoLiqudityPairNotFound() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 50)
	// GIVEN
	// no liqudity

	// WHEN
	// swap 5 of tokenA
	// THEN
	// swap should fail with PairNotFound error
	err := types.ErrValidPairNotFound
	s.bobMarketSellFails(err, "TokenA", 5, 0)

}

func (s *MsgServerTestSuite) TestSwapNoLONoLiqudity() {
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
	err := types.ErrNotEnoughCoins
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
	s.assertLiquidityAtTick(0, 10, 0, 0)
	//
	// WHEN
	// swap 20 of tokenA at
	amountIn, amountInDec := 20, NewDec(20)
	s.bobMarketSells("TokenA", amountIn, 5)

	// THEN
	// swap should have in out
	expectedAmountInLeft, expectedAmountOut := s.calculateSingleSwapNoLOAToB(1, NewDec(10), amountInDec)
	expectedAmountIn := amountInDec.Sub(expectedAmountInLeft)
	s.assertBobBalancesDec(NewDec(50).Sub(expectedAmountIn), expectedAmountOut)
	s.assertDexBalancesDec(expectedAmountIn, NewDec(10).Sub(expectedAmountOut))
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
	s.assertLiquidityAtTick(10, 0, 0, 0)
	//
	// WHEN
	// swap 20 of token A for B
	amountIn, amountInDec := 20, NewDec(20)
	s.bobMarketSells("TokenB", amountIn, 5)

	// THEN
	// swap should have in 9.9990000000000000000 out 10.001000000000000000
	expectedAmountInLeft, expectedAmountOut := s.calculateSingleSwapNoLOBToA(-1, NewDec(10), amountInDec)
	expectedAmountIn := amountInDec.Sub(expectedAmountInLeft)
	s.assertBobBalancesDec(expectedAmountOut, NewDec(50).Sub(expectedAmountIn))
	s.assertDexBalancesDec(NewDec(10).Sub(expectedAmountOut), expectedAmountIn)
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
	err := types.ErrNotEnoughCoins
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
	err := types.ErrNotEnoughCoins
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
	err := types.ErrNotEnoughCoins
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
	amountIn, amountInDec := 5, NewDec(5)
	s.bobMarketSells("TokenA", amountIn, 4)

	// THEN
	// swap should have in 5.000000000000000000 out 4.999500049995000500
	expectedAmountInLeft, expectedAmountOut := s.calculateSingleSwapNoLOAToB(1, NewDec(10), amountInDec)
	expectedAmountIn := amountInDec.Sub(expectedAmountInLeft)
	s.assertBobBalancesDec(NewDec(50).Sub(expectedAmountIn), expectedAmountOut)
	s.assertDexBalancesDec(expectedAmountIn, NewDec(10).Sub(expectedAmountOut))
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
	amountIn, amountInDec := 5, NewDec(5)
	s.bobMarketSells("TokenA", amountIn, 4)

	// THEN
	// swap should have in out
	expectedAmountInLeft, expectedAmountOut := s.calculateSingleSwapNoLOAToB(10, NewDec(10), amountInDec)
	expectedAmountIn := amountInDec.Sub(expectedAmountInLeft)
	s.assertBobBalancesDec(NewDec(50).Sub(expectedAmountIn), expectedAmountOut)
	s.assertDexBalancesDec(expectedAmountIn, NewDec(10).Sub(expectedAmountOut))
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
	amountIn, amountInDec := 15, NewDec(15)
	s.bobMarketSells("TokenA", amountIn, 14)

	// THEN
	// swap should have in out
	amountInLeftTick1, expectedAmountOutTick1 := s.calculateSingleSwapNoLOAToB(1, NewDec(10), amountInDec)
	amountInLeftTick3, expectedAmountOutTick3 := s.calculateSingleSwapNoLOAToB(3, NewDec(10), amountInLeftTick1)
	expectedAmountIn := amountInDec.Sub(amountInLeftTick3)
	expectedAmountOut := expectedAmountOutTick1.Add(expectedAmountOutTick3)
	s.assertBobBalancesDec(NewDec(50).Sub(expectedAmountIn), expectedAmountOut)
	s.assertDexBalancesDec(expectedAmountIn, NewDec(20).Sub(expectedAmountOut))
}

func (s *MsgServerTestSuite) TestSwapNoLODoesntMove1to0() {
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

func (s *MsgServerTestSuite) TestSwapNoLODoesntMove0to1() {
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

func (s *MsgServerTestSuite) TestSwapNoLOMoves1To0() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(0, 50)
	// GIVEN
	// deposit 10 of token B at tick 0 fee 1
	s.aliceDeposits(
		NewDeposit(10, 0, 0, 0),
		NewDeposit(10, 0, 0, 1),
	)
	s.assertCurr1To0(-1)

	// WHEN
	// swap 15 of token B for A with minOut 14
	s.bobMarketSells("TokenB", 15, 14)

	// THEN
	// current 1to0 unchanged
	s.assertCurr1To0(-3)
}

func (s *MsgServerTestSuite) TestSwapNoLOMoves0to1() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(0, 50)
	// GIVEN
	// deposit 10 of token B at tick 0 fee 1
	s.aliceDeposits(
		NewDeposit(0, 10, 0, 0),
		NewDeposit(0, 10, 0, 1),
	)
	s.assertCurr0To1(1)

	// WHEN
	// swap 15 of token A for B with minOut 14
	s.bobMarketSells("TokenA", 15, 14)

	// THEN
	// current 1to0 unchanged
	s.assertCurr0To1(3)
}
