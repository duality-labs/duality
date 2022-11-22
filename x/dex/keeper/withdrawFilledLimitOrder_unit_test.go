package keeper_test

import (
	"math"
	//. "github.com/NicholasDotSol/duality/x/dex/keeper/internal/testutils"
)

func (s *MsgServerTestSuite) TestWithdrawFilledLOSingleFullyFilled() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 50)
	//Case
	// Alice places a limit order of A for B
	// Bob swaps from B to A
	// Alice withdraws the limit order

	// GIVEN
	s.aliceLimitSells("TokenA", 0, 10)
	s.assertAliceBalances(40, 50)
	s.assertBobBalances(50, 50)
	s.assertDexBalances(10, 0)
	s.assertLimitLiquidityAtTick("TokenA", 0, 10)
	s.assertCurr1To0(0)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(0)

	s.bobMarketSells("TokenB", 10, 10)
	s.assertAliceBalances(40, 50)
	s.assertBobBalances(60, 40)
	s.assertDexBalances(0, 10)

	// THEN
	// Alice withdraws her partialy filled LO (10 B)
	s.aliceWithdrawsLimitSell("TokenA", 0, 0)

	s.assertAliceBalances(40, 60)
	s.assertBobBalances(60, 40)
	s.assertDexBalances(0, 0)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(math.MaxInt64)
}

func (s *MsgServerTestSuite) TestWithdrawFilledLOSinglePartiallyFilled() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 50)
	//Case
	// Alice places a limit order of 10 A for 10 B
	// Bob swaps from 5 B to 5 A
	// Alice withdraws the limit order (partial)

	// GIVEN
	s.aliceLimitSells("TokenA", 0, 10)
	s.assertAliceBalances(40, 50)
	s.assertBobBalances(50, 50)
	s.assertDexBalances(10, 0)
	s.assertLimitLiquidityAtTick("TokenA", 0, 10)
	s.assertCurr1To0(0)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(0)

	s.bobMarketSells("TokenB", 5, 5)
	s.assertAliceBalances(40, 50)
	s.assertBobBalances(55, 45)
	s.assertDexBalances(5, 5)

	// THEN
	// Alice withdraws her partialy filled LO (5 B)
	s.aliceWithdrawsLimitSell("TokenA", 0, 0)

	s.assertAliceBalances(40, 55)
	s.assertBobBalances(55, 45)
	s.assertDexBalances(5, 0)
	s.assertCurr1To0(0)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(0)
}
func (s *MsgServerTestSuite) TestWithdrawFilledLOMultiSameTranchePartiallyFilled() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 50)
	s.fundCarolBalances(50, 50)
	//Case
	// Alice places a limit order of A for B
	// Carol places a limit order of A for B
	// Bob swaps from B to A
	// Alice withdraws part of the limit order

	// Given
	s.aliceLimitSells("TokenA", 0, 10)
	s.assertAliceBalances(40, 50)
	s.assertDexBalances(10, 0)
	s.assertCurr1To0(0)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(0)

	s.carolLimitSells("TokenA", 0, 10)
	s.assertCarolBalances(40, 50)
	s.assertDexBalances(20, 0)
	s.assertCurr1To0(0)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(0)

	s.assertBobBalances(50, 50)

	s.assertLimitLiquidityAtTick("TokenA", 0, 20)

	s.bobMarketSells("TokenB", 10, 10)
	s.assertBobBalances(60, 40)
	s.assertDexBalances(10, 10)

	//THEN
	// Alice & Carol can remove their partial position from the LO
	s.aliceWithdrawsLimitSell("TokenA", 0, 0)

	s.assertAliceBalances(40, 55)
	s.assertBobBalances(60, 40)
	s.assertDexBalances(10, 5)
	s.assertCurr1To0(0)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(0)

	s.carolWithdrawsLimitSell("TokenA", 0, 0)
	s.assertCarolBalances(40, 55)
	s.assertDexBalances(10, 0)
	s.assertCurr1To0(0)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(0)
}

func (s *MsgServerTestSuite) TestWithdrawFilledLOMultiSameTrancheFullyFilled() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 50)
	s.fundCarolBalances(50, 50)

	//Case
	// Alice places a limit order of A for B
	// Carol places a limit order of A for B
	// Bob swaps from B to A
	// Alice withdraws all of the limit order

	// GIVEN
	s.aliceLimitSells("TokenA", 0, 10)
	s.assertAliceBalances(40, 50)
	s.assertDexBalances(10, 0)

	s.carolLimitSells("TokenA", 0, 10)
	s.assertCarolBalances(40, 50)
	s.assertDexBalances(20, 0)

	s.assertBobBalances(50, 50)

	s.assertLimitLiquidityAtTick("TokenA", 0, 20)

	s.bobMarketSells("TokenB", 20, 20)
	s.assertBobBalances(70, 30)

	s.assertDexBalances(0, 20)

	//THEN
	// Alice & Carol can remove their full position from the LO
	s.aliceWithdrawsLimitSell("TokenA", 0, 0)

	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(math.MaxInt64)

	s.carolWithdrawsLimitSell("TokenA", 0, 0)
	s.assertCarolBalances(40, 60)
	s.assertDexBalances(0, 0)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(math.MaxInt64)

}

func (s *MsgServerTestSuite) TestWithdrawFilledLOMultiDifferentTranchesPartiallyFilled() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 50)
	s.fundCarolBalances(50, 50)

	//Case
	// Alice places a limit order of A for B (10A for 10B)
	// Carol places a limit order of A for B (10A for 10B)
	// Bob partially swaps from B to A (5)
	// Alice places a limit order of A for B (10A for 10b)
	// Bob swaps from B to A (10)
	// Aliec withdraws partially

	// GIVEN
	s.aliceLimitSells("TokenA", 0, 10)
	s.assertAliceBalances(40, 50)
	s.assertDexBalances(10, 0)

	s.carolLimitSells("TokenA", 0, 10)
	s.assertCarolBalances(40, 50)
	s.assertDexBalances(20, 0)
	s.assertCurr1To0(0)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(0)

	s.assertBobBalances(50, 50)
	s.assertLimitLiquidityAtTick("TokenA", 0, 20)

	s.bobMarketSells("TokenB", 10, 10)
	s.assertBobBalances(60, 40)
	s.assertDexBalances(10, 10)

	s.aliceLimitSells("TokenA", 0, 10)
	s.assertLimitLiquidityAtTick("TokenA", 0, 20)

	s.assertAliceBalances(30, 50)
	s.assertDexBalances(20, 10)
	s.assertCurr1To0(0)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(0)

	s.bobMarketSells("TokenB", 15, 15)
	s.assertLimitLiquidityAtTick("TokenA", 0, 5)
	s.assertBobBalances(75, 25)
	s.assertDexBalances(5, 25)

	// THEN
	// Alice and Carol withdraws their partial liqudity
	s.aliceWithdrawsLimitSell("TokenA", 0, 0)

	s.assertAliceBalances(30, 60)
	s.assertBobBalances(75, 25)
	s.assertDexBalances(5, 15)

	s.assertCurr1To0(0)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(0)

	s.carolWithdrawsLimitSell("TokenA", 0, 0)
	s.assertCarolBalances(40, 60)
	s.assertBobBalances(75, 25)
	s.assertDexBalances(5, 5)

	s.assertCurr1To0(0)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(0)
}

func (s *MsgServerTestSuite) TestWithdrawFilledLOMultiDifferentTranchesFullyFilled() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 50)
	s.fundCarolBalances(50, 50)

	//Case
	// Alice places a limit order of A for B (10A for 10B)
	// Carol places a limit order of A for B (10A for 10B)
	// Bob partially swaps from B to A (5)
	// Alice places a limit order of A for B (10A for 10b)
	// Bob swaps from B to A (10)
	// Alice and Carol withdraws fully

	// GIVEN
	s.aliceLimitSells("TokenA", 0, 10)
	s.assertAliceBalances(40, 50)
	s.assertDexBalances(10, 0)

	s.carolLimitSells("TokenA", 0, 10)
	s.assertCarolBalances(40, 50)
	s.assertDexBalances(20, 0)
	s.assertCurr1To0(0)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(0)

	s.assertBobBalances(50, 50)
	s.assertLimitLiquidityAtTick("TokenA", 0, 20)

	s.bobMarketSells("TokenB", 10, 10)
	s.assertBobBalances(60, 40)
	s.assertDexBalances(10, 10)

	s.aliceLimitSells("TokenA", 0, 10)
	s.assertLimitLiquidityAtTick("TokenA", 0, 20)

	s.assertAliceBalances(30, 50)
	s.assertDexBalances(20, 10)
	s.assertCurr1To0(0)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(0)

	s.bobMarketSells("TokenB", 20, 20)
	s.assertBobBalances(80, 20)
	s.assertDexBalances(0, 30)

	// THEN
	// Alice and Carol fully withdraw their liqudity
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(math.MaxInt64)

	s.aliceWithdrawsLimitSell("TokenA", 0, 0)
	s.assertAliceBalances(30, 60)
	s.assertBobBalances(80, 20)
	s.assertDexBalances(0, 20)

	s.carolWithdrawsLimitSell("TokenA", 0, 0)
	s.assertCarolBalances(40, 60)
	s.assertBobBalances(80, 20)
	s.assertDexBalances(0, 10)

	s.aliceWithdrawsLimitSell("TokenA", 0, 0)
	s.assertAliceBalances(30, 70)
	s.assertBobBalances(80, 20)
	s.assertDexBalances(0, 0)
}
