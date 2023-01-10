package keeper_test

import (
	"math"

	"github.com/NicholasDotSol/duality/x/dex/types"
)

func (s *MsgServerTestSuite) TestLimitOrderOverdraw() {
	s.fundAliceBalances(100, 100)
	s.fundBobBalances(100, 100)
	s.fundCarolBalances(100, 100)

	s.aliceLimitSells("TokenB", 0, 20)
	s.bobLimitSells("TokenB", 0, 20)

	s.assertAliceBalances(100, 80)
	s.assertBobBalances(100, 80)
	s.assertCarolBalances(100, 100)
	s.assertDexBalances(0, 40)

	s.carolMarketSells("TokenA", 20, 20)

	s.assertAliceBalances(100, 80)
	s.assertBobBalances(100, 80)
	s.assertCarolBalances(80, 120)
	s.assertDexBalances(20, 20)

	s.aliceWithdrawsLimitSell("TokenB", 0, 0)

	s.assertAliceBalances(110, 80)
	s.assertBobBalances(100, 80)
	s.assertCarolBalances(80, 120)
	s.assertDexBalances(10, 20)

	s.bobWithdrawsLimitSell("TokenB", 0, 0)

	s.assertAliceBalances(110, 80)
	s.assertBobBalances(110, 80)
	s.assertCarolBalances(80, 120)
	s.assertDexBalances(0, 20)

	_, err := s.msgServer.WithdrawFilledLimitOrder(s.goCtx, &types.MsgWithdrawFilledLimitOrder{
		Creator:   s.alice.String(),
		Receiver:  s.alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: int64(0),
		KeyToken:  "TokenB",
		Key:       uint64(0),
	})
	s.Assert().NotEmpty(err)

	s.assertAliceBalances(110, 80)
	s.assertBobBalances(110, 80)
	s.assertCarolBalances(80, 120)
	s.assertDexBalances(0, 20)
}

func (s *MsgServerTestSuite) TestFailsWhenWithdrawNotCalledByOwner() {
	s.fundAliceBalances(100000, 500)
	s.fundBobBalances(100, 200)

	s.aliceLimitSells("TokenA", 0, 25)

	_, err := s.msgServer.WithdrawFilledLimitOrder(s.goCtx, &types.MsgWithdrawFilledLimitOrder{
		Creator:   s.bob.String(),
		Receiver:  s.alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: 0,
		KeyToken:  "TokenB",
		Key:       0,
	})
	s.Assert().Error(err)
}

func (s *MsgServerTestSuite) TestFailsWhenWrongKeyToken() {
	s.fundAliceBalances(100000, 500)
	s.fundBobBalances(100, 200)

	s.aliceLimitSells("TokenA", 0, 25)

	// Errors because of wrong KeyToken
	_, err := s.msgServer.WithdrawFilledLimitOrder(s.goCtx, &types.MsgWithdrawFilledLimitOrder{
		Creator:   s.alice.String(),
		Receiver:  s.alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: 0,
		KeyToken:  "TokenA",
		Key:       0,
	})
	s.Assert().Error(err)
}

func (s *MsgServerTestSuite) TestFailsWhenWrongKey() {
	s.fundAliceBalances(100000, 500)
	s.fundBobBalances(100, 200)

	s.aliceLimitSells("TokenA", 0, 25)

	// errors because of wrong key
	_, err := s.msgServer.WithdrawFilledLimitOrder(s.goCtx, &types.MsgWithdrawFilledLimitOrder{
		Creator:   s.alice.String(),
		Receiver:  s.alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: 0,
		KeyToken:  "TokenB",
		Key:       1,
	})
	s.Assert().Error(err)
}

func (s *MsgServerTestSuite) TestProgressiveLimitOrderFill() {
	s.fundAliceBalances(100, 500)
	s.fundBobBalances(100, 200)

	s.aliceDeposits(NewDeposit(0, 10, 0, 0))
	s.aliceLimitSells("TokenB", 0, 50)

	s.assertAliceBalances(100, 440)
	s.assertBobBalances(100, 200)
	s.assertDexBalances(0, 60)

	s.bobMarketSells("TokenA", 10, 10)

	s.assertAliceBalances(100, 440)
	s.assertBobBalances(90, 210)
	s.assertDexBalances(10, 50)

	s.aliceWithdrawsLimitSell("TokenB", 0, 0)

	// Limit order is filled progressively
	s.assertAliceBalances(110, 440)
	s.assertBobBalances(90, 210)
	s.assertDexBalances(0, 50)

	// TODO: How to verify current tick?
}

func (s *MsgServerTestSuite) TestLimitOrderPartialFillDepositCancel() {
	s.fundAliceBalances(100, 100)
	s.fundBobBalances(100, 100)
	s.assertDexBalances(0, 0)

	s.aliceLimitSells("TokenB", 0, 50)

	s.assertAliceBalances(100, 50)
	s.assertBobBalances(100, 100)
	s.assertDexBalances(0, 50)
	s.assertCurrentTicks(math.MinInt64, 0)
	s.assertMaxTick(0)
	s.assertMinTick(math.MaxInt64)

	s.bobMarketSells("TokenA", 10, 10)

	s.assertAliceBalances(100, 50)
	s.assertBobBalances(90, 110)
	s.assertDexBalances(10, 40)
	s.assertCurrentTicks(math.MinInt64, 0)
	s.assertMaxTick(0)
	s.assertMinTick(math.MaxInt64)

	s.aliceLimitSells("TokenB", 0, 50)

	s.assertAliceBalances(100, 0)
	s.assertBobBalances(90, 110)
	s.assertDexBalances(10, 90)
	s.assertCurrentTicks(math.MinInt64, 0)
	s.assertMaxTick(0)
	s.assertMinTick(math.MaxInt64)

	s.aliceCancelsLimitSell("TokenB", 0, 0)

	s.assertAliceBalances(100, 40)
	s.assertBobBalances(90, 110)
	s.assertDexBalances(10, 50)
	s.assertCurrentTicks(math.MinInt64, 0)
	s.assertMaxTick(0)
	s.assertMinTick(math.MaxInt64)

	s.bobMarketSells("TokenA", 10, 10)

	s.assertAliceBalances(100, 40)
	s.assertBobBalances(80, 120)
	s.assertDexBalances(20, 40)

	s.aliceCancelsLimitSell("TokenB", 0, 1)

	s.assertAliceBalances(100, 80)
	s.assertBobBalances(80, 120)
	s.assertDexBalances(20, 0)

	s.aliceWithdrawsLimitSell("TokenB", 0, 0)

	s.assertAliceBalances(110, 80)
	s.assertBobBalances(80, 120)
	s.assertDexBalances(10, 0)

	s.aliceWithdrawsLimitSell("TokenB", 0, 1)

	s.assertAliceBalances(120, 80)
	s.assertBobBalances(80, 120)
	s.assertDexBalances(0, 0)
}
