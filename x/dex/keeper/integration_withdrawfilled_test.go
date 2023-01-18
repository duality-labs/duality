package keeper_test

import (
	"math"

	"github.com/NicholasDotSol/duality/x/dex/types"
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

	s.bobMarketSells("TokenB", 10, 10)

	s.assertAliceBalances(40, 50)
	s.assertBobBalances(60, 40)
	s.assertDexBalances(0, 10)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(math.MaxInt64)

	s.aliceWithdrawsLimitSell("TokenA", 0, 0)

	s.assertAliceBalances(40, 60)
	s.assertBobBalances(60, 40)
	s.assertDexBalances(0, 0)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(math.MaxInt64)
}

func (s *MsgServerTestSuite) TestWithdrawFilledPartial() {
	s.fundAliceBalances(100, 100)
	s.fundBobBalances(100, 100)

	// GIVEN
	// alice limit sells 50 B at tick 0
	s.aliceLimitSells("TokenB", 0, 50)
	s.assertAliceLimitLiquidityAtTick("TokenB", 50, 0)
	// bob market sells 10 A
	s.bobMarketSells("TokenA", 10, 10)
	// alice has 10 A filled
	s.assertAliceLimitFilledAtTickAtKey("TokenB", 10, 0, 0)
	// balances are 50, 100 for alice and 90, 100 for bob
	s.assertAliceBalances(100, 50)
	s.assertBobBalances(90, 110)

	// WHEN
	// alice withdraws filled limit order proceeds from tick 0 tranche 0
	s.aliceWithdrawsLimitSell("TokenB", 0, 0)

	// THEN
	// limit order has been partially filled
	s.assertAliceLimitLiquidityAtTick("TokenB", 40, 0)
	// the filled reserved have been withdrawn from
	s.assertAliceLimitFilledAtTickAtKey("TokenB", 0, 0, 0)
	// balances are 110, 100 for alice and 90, 100 for bob
	s.assertAliceBalances(110, 50)
	s.assertBobBalances(90, 110)
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

	s.bobMarketSells("TokenB", 10, 10)

	s.assertAliceBalances(40, 50)
	s.assertBobBalances(60, 40)
	s.assertDexBalances(0, 10)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(math.MaxInt64)

	s.aliceWithdrawsLimitSell("TokenA", 0, 0)
	s.aliceLimitSells("TokenA", 0, 10)

	s.assertAliceBalances(30, 60)
	s.assertBobBalances(60, 40)
	s.assertDexBalances(10, 0)
	s.assertCurr1To0(0)
	s.assertCurr0To1(math.MaxInt64)

	s.bobMarketSells("TokenB", 10, 10)

	s.assertAliceBalances(30, 60)
	s.assertBobBalances(70, 30)
	s.assertDexBalances(0, 10)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(math.MaxInt64)

	s.aliceWithdrawsLimitSell("TokenA", 0, 1)

	s.assertAliceBalances(30, 70)
	s.assertBobBalances(70, 30)
	s.assertDexBalances(0, 0)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(math.MaxInt64)
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

	s.bobMarketSells("TokenB", 10, 10)

	s.assertAliceBalances(40, 50)
	s.assertBobBalances(60, 40)
	s.assertDexBalances(0, 10)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(math.MaxInt64)

	s.aliceWithdrawsLimitSell("TokenA", 0, 0)
	s.aliceLimitSells("TokenB", 0, 10)

	s.assertAliceBalances(40, 50)
	s.assertBobBalances(60, 40)
	s.assertDexBalances(0, 10)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(0)

	s.bobMarketSells("TokenA", 10, 10)

	s.assertAliceBalances(40, 50)
	s.assertBobBalances(50, 50)
	s.assertDexBalances(10, 0)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(math.MaxInt64)

	s.aliceWithdrawsLimitSell("TokenB", 0, 0)

	s.assertAliceBalances(50, 50)
	s.assertBobBalances(50, 50)
	s.assertDexBalances(0, 0)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(math.MaxInt64)
}

func (s *MsgServerTestSuite) TestWithdrawFilledInvalidKeyToken() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 50)

	// GIVEN
	// existing limit order selling TokenA for TokenB that was filled
	s.aliceLimitSells("TokenA", 0, 10)
	s.bobMarketSells("TokenB", 10, 0)

	// WHEN
	// withdrawing the sold tokens
	// THEN
	// withdraw should fail with ErrInvalidTradingPair

	err := types.ErrInvalidKeyToken
	s.aliceWithdrawLimitSellFails(err, "TokenC", 0, 0)
}

func (s *MsgServerTestSuite) TestWithdrawFilledEmptyFilled() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// alice places limit order selling A for B at tick 0
	s.aliceLimitSells("TokenA", 0, 10)

	// WHEN
	// order is unfilled, i.e. trachne.filled = 0
	// THEN

	err := types.ErrWithdrawEmptyLimitOrder
	s.aliceWithdrawLimitSellFails(err, "TokenA", 0, 0)
}

func (s *MsgServerTestSuite) TestWithdrawFilledNoExistingOrderByUser() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 50)

	// GIVEN
	// only alice has an existing order placed
	s.aliceLimitSells("TokenA", 0, 10)

	// WHEN
	// bob tries to withdraw filled from tick 0 tranche 0
	// THEN

	err := types.ErrValidLimitOrderTrancheUserNotFound
	s.bobWithdrawLimitSellFails(err, "TokenA", 0, 0)
}

func (s *MsgServerTestSuite) TestWithdrawFilledTrancheKeyDoesntExist() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 50)

	// GIVEN
	// only alice has a single existing order placed
	s.aliceLimitSells("TokenA", 0, 10)

	// WHEN
	// bob tries to withdraw filled from tick 0 tranche 5
	// THEN

	err := types.ErrValidLimitOrderTrancheUserNotFound
	s.bobWithdrawLimitSellFails(err, "TokenA", 0, 5)
}
