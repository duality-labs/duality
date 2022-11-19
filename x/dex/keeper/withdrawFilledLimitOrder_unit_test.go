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
	// Alice adds a limit order of A for B
	// Bob swaps through

	s.aliceLimitSells("TokenA", 0, 10)

	s.assertAliceBalances(40, 50)
	s.assertAliceBalances(40, 50)
	s.assertDexBalances(10, 0)
	s.assertCurr1To0(0)
	s.assertCurr0To1(math.MinInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(0)

	s.bobMarketSells("TokenB", 0, 0, 10)

	s.assertAliceBalances(50, 50)
	s.assertDexBalances(0, 0)
	s.assertCurr1To0(math.MaxInt64)
	s.assertCurr0To1(math.MinInt64)
	s.assertMaxTick(math.MinInt64)
	s.assertMinTick(math.MaxInt64)
}
