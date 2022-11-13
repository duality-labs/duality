package keeper_test

import (
	. "github.com/NicholasDotSol/duality/x/dex/keeper/internal/testutils"
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (s *MsgServerTestSuite) TestSingle() {
	s.fundAliceBalances(100, 500)
	s.fundBobBalances(100, 200)

	s.alicePlacesLimitOrder("TokenA", 0, 50)

	s.assertAliceBalances(100, 450)
	s.assertBobBalances(100, 200)
	s.assertDexBalances(0, 50)
}

func (s *MsgServerTestSuite) TestMultiple() {
	s.fundAliceBalances(100, 500)
	s.fundBobBalances(100, 200)

	s.alicePlacesLimitOrder("TokenA", 0, 50)

	s.assertAliceBalances(100, 450)
	s.assertBobBalances(100, 200)
	s.assertDexBalances(0, 50)

	s.alicePlacesLimitOrder("TokenA", 0, 50)

	s.assertAliceBalances(100, 400)
	s.assertBobBalances(100, 200)
	s.assertDexBalances(0, 100)

	_, err := s.msgServer.PlaceLimitOrder(s.goCtx, &types.MsgPlaceLimitOrder{
		Creator:   s.bob.String(),
		Receiver:  s.bob.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: 0,
		TokenIn:   "TokenB",
		AmountIn:  NewDec(100),
	})
	s.Assert().Nil(err)

	s.assertAliceBalances(100, 400)
	s.assertBobBalances(100, 100)
	s.assertDexBalances(0, 200)
}

func (s *MsgServerTestSuite) TestDifferentReceiverAndCreator() {
	s.fundAliceBalances(100, 500)
	s.fundBobBalances(100, 200)

	_, err := s.msgServer.PlaceLimitOrder(s.goCtx, &types.MsgPlaceLimitOrder{
		Creator:   s.bob.String(),
		Receiver:  s.alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: 0,
		TokenIn:   "TokenB",
		AmountIn:  NewDec(100),
	})
	s.Assert().Nil(err)

	s.assertAliceBalances(100, 500)
	s.assertBobBalances(100, 100)
	s.assertDexBalances(0, 100)
}

func (s *MsgServerTestSuite) TestFailUnrecognizedToken() {
	s.fundAliceBalances(100, 500)
	s.fundBobBalances(100, 200)

	_, err := s.msgServer.PlaceLimitOrder(s.goCtx, &types.MsgPlaceLimitOrder{
		Creator:   s.bob.String(),
		Receiver:  s.bob.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: 0,
		TokenIn:   "TokenC",
		AmountIn:  NewDec(100),
	})
	s.Assert().Error(err)
}

func (s *MsgServerTestSuite) TestFailInsufficientBalance() {
	s.fundAliceBalances(100, 500)
	s.fundBobBalances(100, 200)

	_, err := s.msgServer.PlaceLimitOrder(s.goCtx, &types.MsgPlaceLimitOrder{
		Creator:   s.bob.String(),
		Receiver:  s.bob.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: 0,
		TokenIn:   "TokenB",
		AmountIn:  NewDec(1000),
	})
	s.Assert().Error(err)
}

func (s *MsgServerTestSuite) TestMultiTickLimitOrder1to0WithWithdraw() {
	s.fundAliceBalances(100, 500)
	s.fundBobBalances(100, 200)

	s.alicePlacesLimitOrder("TokenA", 0, 25)
	s.alicePlacesLimitOrder("TokenA", -1, 25)
	s.bobPlacesSwapOrder("TokenB", 40, 30)

	s.assertAliceBalances(100, 450)
	s.assertBobBalancesDec(NewDec(60), sdk.MustNewDecFromStr("239.9985"))

	s.aliceWithdrawsFilledLimitOrder("TokenB", 0)

	s.assertAliceBalances(125, 450)
	s.assertBobBalancesDec(NewDec(60), sdk.MustNewDecFromStr("239.9985"))

	s.aliceWithdrawsFilledLimitOrder("TokenB", -1)

	s.assertAliceBalancesDec(sdk.MustNewDecFromStr("140.0015"), NewDec(450))
	s.assertBobBalancesDec(NewDec(60), sdk.MustNewDecFromStr("239.9985"))
}

func (s *MsgServerTestSuite) TestMultiTickLimitOrder0to1WithWithdraw() {
	s.fundAliceBalances(100000, 500)
	s.fundBobBalances(100, 200)

	s.alicePlacesLimitOrder("TokenB", 0, 25)
	s.alicePlacesLimitOrder("TokenB", 1, 25)

	s.assertAliceBalances(99950, 500)
	s.assertBobBalances(100, 200)
	s.assertDexBalances(50, 0)

	s.bobPlacesSwapOrder("TokenA", 40, 30)

	s.assertAliceBalances(99950, 500)
	s.assertBobBalancesDec(sdk.MustNewDecFromStr("139.99850015"), NewDec(160))
	s.assertDexBalancesDec(sdk.MustNewDecFromStr("10.00149985"), NewDec(40))

	s.aliceWithdrawsFilledLimitOrder("TokenA", 0)

	s.assertAliceBalances(99950, 525)
	s.assertBobBalancesDec(sdk.MustNewDecFromStr("139.99850015"), NewDec(160))
	s.assertDexBalancesDec(sdk.MustNewDecFromStr("10.00149985"), NewDec(15))

	s.aliceWithdrawsFilledLimitOrder("TokenA", 1)

	s.assertAliceBalances(99950, 540)
	s.assertBobBalancesDec(sdk.MustNewDecFromStr("139.99850015"), NewDec(160))
	s.assertDexBalancesDec(sdk.MustNewDecFromStr("10.00149985"), NewDec(0))
}

func (s *MsgServerTestSuite) TestWithdrawFailsWhenNothingToWithdraw() {
	s.fundAliceBalances(100000, 500)
	s.fundBobBalances(100, 200)

	_, err := s.msgServer.WithdrawFilledLimitOrder(s.goCtx, &types.MsgWithdrawFilledLimitOrder{
		Creator:   s.alice.String(),
		Receiver:  s.alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: 0,
		KeyToken:  "TokenB",
		Key:       0,
	})
	s.Assert().Error(err)
}

func (s *MsgServerTestSuite) TestFailsWhenWithdrawNotCalledByOwner() {
	s.fundAliceBalances(100000, 500)
	s.fundBobBalances(100, 200)

	s.alicePlacesLimitOrder("TokenB", 0, 25)

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

	s.alicePlacesLimitOrder("TokenB", 0, 25)

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

	s.alicePlacesLimitOrder("TokenB", 0, 25)

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

func (s *MsgServerTestSuite) TestCancelSingle() {
	s.fundAliceBalances(100, 500)

	s.assertDexBalances(0, 0)

	s.alicePlacesLimitOrder("TokenA", 0, 50)

	s.assertAliceBalances(100, 450)
	s.assertDexBalances(0, 50)

	s.aliceCancelsLimitOrder("TokenB", 0, 0, 50)

	s.assertAliceBalances(100, 500)
	s.assertDexBalances(0, 0)
}

func (s *MsgServerTestSuite) TestCancelPartial() {
	s.fundAliceBalances(100, 500)

	s.assertDexBalances(0, 0)

	s.alicePlacesLimitOrder("TokenA", 0, 50)

	s.assertAliceBalances(100, 450)
	s.assertDexBalances(0, 50)

	s.aliceCancelsLimitOrder("TokenB", 0, 0, 25)

	s.assertAliceBalances(100, 475)
	s.assertDexBalances(0, 25)

	s.aliceCancelsLimitOrder("TokenB", 0, 0, 25)

	s.assertAliceBalances(100, 500)
	s.assertDexBalances(0, 0)
}

func (s *MsgServerTestSuite) TestProgressiveLimitOrderFill() {
	s.fundAliceBalances(100, 500)
	s.fundBobBalances(100, 200)

	s.aliceDeposits(NewDeposit(0, 10, 0, 0))
	s.alicePlacesLimitOrder("TokenA", 0, 50)

	s.assertAliceBalances(100, 440)
	s.assertBobBalances(100, 200)
	s.assertDexBalances(0, 60)

	s.bobPlacesSwapOrder("TokenB", 10, 10)

	s.assertAliceBalances(100, 440)
	s.assertBobBalances(90, 210)
	s.assertDexBalances(10, 50)

	s.aliceWithdrawsFilledLimitOrder("TokenB", 0)

	// Limit order is filled progressively
	s.assertAliceBalances(102, 440)
	s.assertBobBalances(90, 210)
	s.assertDexBalances(8, 50)

	// TODO: How to verify current tick?
}
