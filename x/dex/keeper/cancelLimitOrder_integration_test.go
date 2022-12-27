package keeper_test

import (
	"math"

	"github.com/NicholasDotSol/duality/x/dex/types"
)

func (s *MsgServerTestSuite) TestCancelEntireLimitOrderAOneExists() {
	s.fundAliceBalances(50, 50)
	// CASE
	// Alice adds a limit order of A for B and cancels it right away

	s.aliceLimitSells("TokenA", 0, 10)

	s.assertAliceBalances(40, 50)
	s.assertDexBalances(10, 0)
	s.assertCurr1To0(0)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(0)

	s.aliceCancelsLimitSell("TokenA", 0, 0)

	s.assertAliceBalances(50, 50)
	s.assertDexBalances(0, 0)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(math.MaxInt64)
}

func (s *MsgServerTestSuite) TestCancelEntireLimitOrderBOneExists() {
	s.fundAliceBalances(50, 50)
	// CASE
	// Alice adds a limit order of B for A and cancels it right away

	s.aliceLimitSells("TokenB", 0, 10)

	s.assertAliceBalances(50, 40)
	s.assertDexBalances(0, 10)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(0)
	s.assertMaxTick(0)
	s.assertMinTick(math.MaxInt64)

	s.aliceCancelsLimitSell("TokenB", 0, 0)

	s.assertAliceBalances(50, 50)
	s.assertDexBalances(0, 0)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(math.MaxInt64)
}

func (s *MsgServerTestSuite) TestCancelHigherEntireLimitOrderATwoExistDiffTicksSameDirection() {
	s.fundAliceBalances(50, 50)
	// CASE
	// Alice adds two limit orders from A to B and removes the one at the higher tick (0)

	s.aliceLimitSells("TokenA", 0, 10)
	s.aliceLimitSells("TokenA", -1, 10)

	s.assertAliceBalances(30, 50)
	s.assertDexBalances(20, 0)
	s.assertCurr1To0(0)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(-1)

	s.aliceCancelsLimitSell("TokenA", 0, 0)

	s.assertAliceBalances(40, 50)
	s.assertDexBalances(10, 0)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(-1)
}

func (s *MsgServerTestSuite) TestCancelLowerEntireLimitOrderATwoExistDiffTicksSameDirection() {
	s.fundAliceBalances(50, 50)
	// CASE
	// Alice adds two limit orders from A to B and removes the one at the lower tick (-1)

	s.aliceLimitSells("TokenA", 0, 10)
	s.aliceLimitSells("TokenA", -1, 10)

	s.assertAliceBalances(30, 50)
	s.assertDexBalances(20, 0)
	s.assertCurr1To0(0)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(-1)

	s.aliceCancelsLimitSell("TokenA", -1, 0)

	s.assertAliceBalances(40, 50)
	s.assertDexBalances(10, 0)
	s.assertCurr1To0(0)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(0)
}

func (s *MsgServerTestSuite) TestCancelLowerEntireLimitOrderATwoExistDiffTicksDiffDirection() {
	s.fundAliceBalances(50, 50)
	// CASE
	// Alice adds one limit orders from A to B and one from B to A and removes the one from A to B

	s.aliceLimitSells("TokenA", 0, 10)
	s.aliceLimitSells("TokenB", 1, 10)

	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertCurr1To0(0)
	s.assertCurr0To1(1)
	s.assertMaxTick(1)
	s.assertMinTick(0)

	s.aliceCancelsLimitSell("TokenA", 0, 0)

	s.assertAliceBalances(50, 40)
	s.assertDexBalances(0, 10)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(1)
	s.assertMaxTick(1)
	s.assertMinTick(math.MaxInt64)
}

func (s *MsgServerTestSuite) TestCancelHigherEntireLimitOrderBTwoExistDiffTicksSameDirection() {
	s.fundAliceBalances(50, 50)
	// CASE
	// Alice adds two limit orders from B to A and removes the one at tick 0

	s.aliceLimitSells("TokenB", 0, 10)
	s.aliceLimitSells("TokenB", -1, 10)

	s.assertAliceBalances(50, 30)
	s.assertDexBalances(0, 20)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(-1)
	s.assertMaxTick(0)
	s.assertMinTick(math.MaxInt64)

	s.aliceCancelsLimitSell("TokenB", 0, 0)

	s.assertAliceBalances(50, 40)
	s.assertDexBalances(0, 10)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(-1)
	s.assertMaxTick(-1)
	s.assertMinTick(math.MaxInt64)
}

func (s *MsgServerTestSuite) TestCancelLowerEntireLimitOrderBTwoExistDiffTicksSameDirection() {
	s.fundAliceBalances(50, 50)
	// CASE
	// Alice adds two limit orders from B to A and removes the one at tick 0

	s.aliceLimitSells("TokenB", 0, 10)
	s.aliceLimitSells("TokenB", -1, 10)

	s.assertAliceBalances(50, 30)
	s.assertDexBalances(0, 20)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(-1)
	s.assertMaxTick(0)
	s.assertMinTick(math.MaxInt64)

	s.aliceCancelsLimitSell("TokenB", -1, 0)

	s.assertAliceBalances(50, 40)
	s.assertDexBalances(0, 10)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(0)
	s.assertMaxTick(0)
	s.assertMinTick(math.MaxInt64)
}

func (s *MsgServerTestSuite) TestCancelTwiceFails() {
	s.fundAliceBalances(50, 50)
	// CASE
	// Alice tries to cancel the same limit order twice

	s.aliceLimitSells("TokenB", 0, 10)

	s.assertAliceBalances(50, 40)
	s.assertDexBalances(0, 10)

	s.aliceCancelsLimitSell("TokenB", 0, 0)

	s.assertAliceBalances(50, 50)
	s.assertDexBalances(0, 0)

	s.aliceCancelsLimitSellFails("TokenB", -1, 0, types.ErrValidTickNotFound)

}
