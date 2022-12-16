package keeper_test

import (
	"math"

	testing_scripts "github.com/NicholasDotSol/duality/testing_scripts"
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (s *MsgServerTestSuite) TestSingle() {
	s.fundAliceBalances(100, 500)
	s.fundBobBalances(100, 200)

	s.aliceLimitSells("TokenB", 0, 50)

	s.assertAliceBalances(100, 450)
	s.assertBobBalances(100, 200)
	s.assertDexBalances(0, 50)
}

func (s *MsgServerTestSuite) TestMultiple() {
	s.fundAliceBalances(100, 500)
	s.fundBobBalances(100, 200)

	s.aliceLimitSells("TokenB", 0, 50)

	s.assertAliceBalances(100, 450)
	s.assertBobBalances(100, 200)
	s.assertDexBalances(0, 50)

	s.aliceLimitSells("TokenB", 0, 50)

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
		AmountIn:  sdk.NewInt(100),
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
		AmountIn:  sdk.NewInt(100),
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
		AmountIn:  sdk.NewInt(100),
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
		AmountIn:  sdk.NewInt(1000),
	})
	s.Assert().Error(err)
}

func (s *MsgServerTestSuite) TestMultiTickLimitOrder1to0WithWithdraw() {
	s.fundAliceBalances(100000, 500000)
	s.fundBobBalances(100000, 200000)

	s.aliceLimitSells("TokenB", 1, 25000)
	s.aliceLimitSells("TokenB", 0, 25000)
	s.aliceLimitSells("TokenB", -1, 25000)
	s.bobMarketSells("TokenA", 40000, 30000)

	// limit order at -1: (25000 * 1.0001^-1) A<=>B 25000
	// limit order at 0: (40000 - (25000 * 1.0001^-1)) A<=>B (40000 - (25000 * 1.0001^-1)) * 1.0001^0
	s.assertAliceBalances(100000, 425000)
	s.assertBobBalancesInt(sdk.NewInt(60000), sdk.NewInt(240002))

	s.aliceWithdrawsLimitSell("TokenB", 0, 0)

	s.assertAliceBalancesInt(sdk.NewInt(115002), sdk.NewInt(425000))
	s.assertBobBalancesInt(sdk.NewInt(60000), sdk.NewInt(240002))

	s.aliceWithdrawsLimitSell("TokenB", -1, 0)

	s.assertAliceBalancesEpsilon(sdk.NewInt(140000), sdk.NewInt(425000))
	s.assertBobBalancesInt(sdk.NewInt(60000), sdk.NewInt(240002))
}

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

func (s *MsgServerTestSuite) TestMultiTickLimitOrder0to1WithWithdraw() {
	s.fundAliceBalances(100000000, 500000)
	s.fundBobBalances(100000, 200000)

	//Alices balance for TokenA should be 100000 - 25 - 25 = 99950
	//Alices limit orders can be traded through at a price_1to0, 1 and 1.0001
	s.aliceLimitSells("TokenA", 0, 25000)
	s.aliceLimitSells("TokenA", 1, 25000)

	s.assertAliceBalances(99950000, 500000)
	s.assertBobBalances(100000, 200000)
	s.assertDexBalances(50000, 0)

	testing_scripts.MultipleLimitOrderFills([]sdk.Int{sdk.NewInt(25000), sdk.NewInt(25000)}, []sdk.Dec{sdk.MustNewDecFromStr("1.0001"), sdk.NewDec(1)}, sdk.NewInt(40000))

	//Bobs balance for TokenB should be 200 - 40 = 160
	//Tick 1 should be a swap of 25 / 1.0001 TokenB (1) for 25 of TokenA (0) exhausting all the liquidity
	// => This means the amount of LO filled at tick 1 is 25 / 1.0001 = 24.997500249975
	//Tick 0 should be a swap of the remaining ~15.002499750024999 of TokenB for ~15.002499750024999
	// This is because the price is 1
	//Bobs balance for TokenA should be (1 * 15.002499750024999) + (1.0001 * 24.997500249975) + 100 = 140.002499750024997500
	//DEX Balance should be 50 - (1 * 9.997500249975002500) - (1.0001 * 24.997500249975002500) = 9.997500249975002500
	s.bobMarketSells("TokenB", 40000, 30000)

	s.assertAliceBalances(99950000, 500000)
	s.assertBobBalancesInt(sdk.NewInt(140002), sdk.NewInt(160000))
	s.assertDexBalancesEpsilon(sdk.NewInt(9997), sdk.NewInt(40000))

	s.aliceWithdrawsLimitSell("TokenA", 1, 0)

	s.assertAliceBalancesEpsilon(sdk.NewInt(99950000), sdk.NewInt(524997))
	s.assertBobBalancesInt(sdk.NewInt(140002), sdk.NewInt(160000))
	//40 - 24.997500249975002500 = 15.0024997500249975
	s.assertDexBalancesEpsilon(sdk.NewInt(9997), sdk.NewInt(15002))

	s.aliceWithdrawsLimitSell("TokenA", 0, 0)

	s.assertAliceBalancesEpsilon(sdk.NewInt(99950000), sdk.NewInt(540000))
	s.assertBobBalancesInt(sdk.NewInt(140002), sdk.NewInt(160000))
	s.assertDexBalancesEpsilon(sdk.NewInt(9997), sdk.NewInt(0))
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
