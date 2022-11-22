package keeper_test

import (
	"math"
	//. "github.com/NicholasDotSol/duality/x/dex/keeper/internal/testutils"
	//"github.com/NicholasDotSol/duality/x/dex/types"
)


func (s *MsgServerTestSuite) TestWithdrawFilledSimpleFull() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 50)
	// CASE
	// Alice places a limit order of A for B
	// Bob swaps from B to A
	// Alice withdraws the limit order

	s.aliceLimitSells("TokenA", 0, 10)

	s.assertAliceBalances(40, 50)
	s.assertBobBalances(50, 50)
	s.assertDexBalances(10, 0)
	s.assertCurr1To0(0)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(0)

	s.bobMarketSells("TokenB", 10, 10)

	s.assertAliceBalances(40, 50)
	s.assertBobBalances(60, 40)
	s.assertDexBalances(0, 10)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(math.MaxInt64)

	s.aliceWithdrawsLimitSell("TokenA", 0, 0)

	s.assertAliceBalances(40, 60)
	s.assertBobBalances(60, 40)
	s.assertDexBalances(0, 0)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(math.MaxInt64)
}


func (s *MsgServerTestSuite) TestWithdrawFilledTwiceFullSameDirection() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 50)
	// CASE
	// Alice places a limit order of A for B
	// Bob swaps through
	// Alice withdraws the limit order and places a new one
	// Bob swaps through again
	// Alice withdraws the limit order

	s.aliceLimitSells("TokenA", 0, 10)

	s.assertAliceBalances(40, 50)
	s.assertBobBalances(50, 50)
	s.assertDexBalances(10, 0)
	s.assertCurr1To0(0)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(0)

	s.bobMarketSells("TokenB", 10, 10)

	s.assertAliceBalances(40, 50)
	s.assertBobBalances(60, 40)
	s.assertDexBalances(0, 10)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(math.MaxInt64)

	s.aliceWithdrawsLimitSell("TokenA", 0, 0)
	s.aliceLimitSells("TokenA", 0, 10)

	s.assertAliceBalances(30, 60)
	s.assertBobBalances(60, 40)
	s.assertDexBalances(10, 0)
	s.assertCurr1To0(0)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(0)

	s.bobMarketSells("TokenB", 10, 10)

	s.assertAliceBalances(30, 60)
	s.assertBobBalances(70, 30)
	s.assertDexBalances(0, 10)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(math.MaxInt64)

	s.aliceWithdrawsLimitSell("TokenA", 0, 1)

	s.assertAliceBalances(30, 70)
	s.assertBobBalances(70, 30)
	s.assertDexBalances(0, 0)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(math.MaxInt64)
}

func (s *MsgServerTestSuite) TestWithdrawFilledTwiceFullDifferentDirection() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 50)
	// CASE
	// Alice places a limit order of A for B
	// Bob swaps through
	// Alice withdraws the limit order and places a new one
	// Bob swaps through again
	// Alice withdraws the limit order

	s.aliceLimitSells("TokenA", 0, 10)

	s.assertAliceBalances(40, 50)
	s.assertBobBalances(50, 50)
	s.assertDexBalances(10, 0)
	s.assertCurr1To0(0)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(0)

	s.bobMarketSells("TokenB", 10, 10)

	s.assertAliceBalances(40, 50)
	s.assertBobBalances(60, 40)
	s.assertDexBalances(0, 10)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(math.MaxInt64)

	s.aliceWithdrawsLimitSell("TokenA", 0, 0)
	s.aliceLimitSells("TokenB", 0, 10)

	s.assertAliceBalances(40, 50)
	s.assertBobBalances(60, 40)
	s.assertDexBalances(0, 10)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(0)
	s.assertMaxTick(0)
	s.assertMinTick(math.MaxInt64)

	s.bobMarketSells("TokenA", 10, 10)

	s.assertAliceBalances(40, 50)
	s.assertBobBalances(50, 50)
	s.assertDexBalances(10, 0)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(math.MaxInt64)

	s.aliceWithdrawsLimitSell("TokenB", 0, 0)

	s.assertAliceBalances(50, 50)
	s.assertBobBalances(50, 50)
	s.assertDexBalances(0, 0)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(math.MaxInt64)
}