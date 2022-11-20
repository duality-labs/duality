package keeper_test

import (
	//"fmt"
	"math"
	//. "github.com/NicholasDotSol/duality/x/dex/keeper/internal/testutils"
	//"github.com/NicholasDotSol/duality/x/dex/types"
)


func (s *MsgServerTestSuite) TestWithdrawFilledSimpleFull() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 50)
	// CASE
	// Alice places a limit order of A for B
	// Bob swaps through
	// Alice withdraws the limit order

	s.aliceLimitSells("TokenA", 0, 10)

	s.assertAliceBalances(40, 50)
	s.assertBobBalances(50, 50)
	s.assertDexBalances(10, 0)
	s.assertCurr1To0(0)
	s.assertCurr0To1(math.MinInt64)
	s.assertMaxTick(math.MaxInt64)
	s.assertMinTick(0)

	s.bobMarketSells("TokenB", 0, 10)

	s.assertAliceBalances(40, 50)
	s.assertBobBalances(60, 40)
	s.assertDexBalances(0, 10)
	s.assertCurr1To0(math.MaxInt64)
	s.assertCurr0To1(math.MinInt64)
	s.assertMaxTick(math.MaxInt64)
	s.assertMinTick(math.MinInt64)

	s.aliceWithdrawsLimitSell("TokenA", 0, 0)

	s.assertAliceBalances(40, 60)
	s.assertBobBalances(60, 40)
	s.assertDexBalances(0, 0)
	s.assertCurr1To0(math.MaxInt64)
	s.assertCurr0To1(math.MinInt64)
	s.assertMaxTick(math.MaxInt64)
	s.assertMinTick(math.MinInt64)
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
	s.assertCurr0To1(math.MinInt64)
	s.assertMaxTick(math.MaxInt64)
	s.assertMinTick(0)

	s.bobMarketSells("TokenB", 0, 10)

	s.assertAliceBalances(40, 50)
	s.assertBobBalances(60, 40)
	s.assertDexBalances(0, 10)
	s.assertCurr1To0(math.MaxInt64)
	s.assertCurr0To1(math.MinInt64)
	s.assertMaxTick(math.MaxInt64)
	s.assertMinTick(math.MinInt64)

	s.aliceWithdrawsLimitSell("TokenA", 0, 0)
	s.aliceLimitSells("TokenA", 0, 10)

	s.assertAliceBalances(30, 60)
	s.assertBobBalances(60, 40)
	s.assertDexBalances(10, 0)
	s.assertCurr1To0(0)
	s.assertCurr0To1(math.MinInt64)
	s.assertMaxTick(math.MaxInt64)
	s.assertMinTick(0)

	s.bobMarketSells("TokenB", 0, 10)

	s.assertAliceBalances(30, 60)
	s.assertBobBalances(70, 30)
	s.assertDexBalances(0, 0)
	s.assertCurr1To0(math.MaxInt64)
	s.assertCurr0To1(math.MinInt64)
	s.assertMaxTick(math.MaxInt64)
	s.assertMinTick(math.MinInt64)

	s.aliceWithdrawsLimitSell("TokenA", 0, 0)

	s.assertAliceBalances(30, 70)
	s.assertBobBalances(70, 30)
	s.assertDexBalances(0, 0)
	s.assertCurr1To0(math.MaxInt64)
	s.assertCurr0To1(math.MinInt64)
	s.assertMaxTick(math.MaxInt64)
	s.assertMinTick(math.MinInt64)
}