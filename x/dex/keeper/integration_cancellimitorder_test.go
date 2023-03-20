package keeper_test

import (
	"math"
	"time"

	"github.com/duality-labs/duality/x/dex/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func (s *MsgServerTestSuite) TestCancelEntireLimitOrderAOneExists() {
	s.fundAliceBalances(50, 50)
	// CASE
	// Alice adds a limit order of A for B and cancels it right away

	trancheKey := s.aliceLimitSells("TokenA", 0, 10)

	s.assertAliceBalances(40, 50)
	s.assertDexBalances(10, 0)
	s.assertCurr1To0(0)
	s.assertCurr0To1(math.MaxInt64)

	s.aliceCancelsLimitSell("TokenA", 0, trancheKey)

	s.assertAliceBalances(50, 50)
	s.assertDexBalances(0, 0)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(math.MaxInt64)

	// Assert that the LimitOrderTrancheUser has been deleted
	_, found := s.app.DexKeeper.GetLimitOrderTrancheUser(s.ctx, defaultPairId, 0, "TokenA", trancheKey, s.alice.String())
	s.Assert().False(found)
}

func (s *MsgServerTestSuite) TestCancelEntireLimitOrderBOneExists() {
	s.fundAliceBalances(50, 50)
	// CASE
	// Alice adds a limit order of B for A and cancels it right away

	trancheKey := s.aliceLimitSells("TokenB", 0, 10)

	s.assertAliceBalances(50, 40)
	s.assertDexBalances(0, 10)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(0)

	s.aliceCancelsLimitSell("TokenB", 0, trancheKey)

	s.assertAliceBalances(50, 50)
	s.assertDexBalances(0, 0)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(math.MaxInt64)
}

func (s *MsgServerTestSuite) TestCancelHigherEntireLimitOrderATwoExistDiffTicksSameDirection() {
	s.fundAliceBalances(50, 50)
	// CASE
	// Alice adds two limit orders from A to B and removes the one at the higher tick (0)

	trancheKey := s.aliceLimitSells("TokenA", 0, 10)
	s.aliceLimitSells("TokenA", -1, 10)

	s.assertAliceBalances(30, 50)
	s.assertDexBalances(20, 0)
	s.assertCurr1To0(0)
	s.assertCurr0To1(math.MaxInt64)

	s.aliceCancelsLimitSell("TokenA", 0, trancheKey)

	s.assertAliceBalances(40, 50)
	s.assertDexBalances(10, 0)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(math.MaxInt64)
}

func (s *MsgServerTestSuite) TestCancelLowerEntireLimitOrderATwoExistDiffTicksSameDirection() {
	s.fundAliceBalances(50, 50)
	// CASE
	// Alice adds two limit orders from A to B and removes the one at the lower tick (-1)

	s.aliceLimitSells("TokenA", 0, 10)
	trancheKey := s.aliceLimitSells("TokenA", -1, 10)

	s.assertAliceBalances(30, 50)
	s.assertDexBalances(20, 0)
	s.assertCurr1To0(0)
	s.assertCurr0To1(math.MaxInt64)

	s.aliceCancelsLimitSell("TokenA", -1, trancheKey)

	s.assertAliceBalances(40, 50)
	s.assertDexBalances(10, 0)
	s.assertCurr1To0(0)
	s.assertCurr0To1(math.MaxInt64)
}

func (s *MsgServerTestSuite) TestCancelLowerEntireLimitOrderATwoExistDiffTicksDiffDirection() {
	s.fundAliceBalances(50, 50)
	// CASE
	// Alice adds one limit orders from A to B and one from B to A and removes the one from A to B

	trancheKey := s.aliceLimitSells("TokenA", 0, 10)
	s.aliceLimitSells("TokenB", 1, 10)

	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertCurr1To0(0)
	s.assertCurr0To1(1)

	s.aliceCancelsLimitSell("TokenA", 0, trancheKey)

	s.assertAliceBalances(50, 40)
	s.assertDexBalances(0, 10)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(1)
}

func (s *MsgServerTestSuite) TestCancelHigherEntireLimitOrderBTwoExistDiffTicksSameDirection() {
	s.fundAliceBalances(50, 50)
	// CASE
	// Alice adds two limit orders from B to A and removes the one at tick 0

	trancheKey := s.aliceLimitSells("TokenB", 0, 10)
	s.aliceLimitSells("TokenB", -1, 10)

	s.assertAliceBalances(50, 30)
	s.assertDexBalances(0, 20)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(-1)

	s.aliceCancelsLimitSell("TokenB", 0, trancheKey)

	s.assertAliceBalances(50, 40)
	s.assertDexBalances(0, 10)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(-1)
}

func (s *MsgServerTestSuite) TestCancelLowerEntireLimitOrderBTwoExistDiffTicksSameDirection() {
	s.fundAliceBalances(50, 50)
	// CASE
	// Alice adds two limit orders from B to A and removes the one at tick 0

	s.aliceLimitSells("TokenB", 0, 10)
	trancheKey := s.aliceLimitSells("TokenB", -1, 10)

	s.assertAliceBalances(50, 30)
	s.assertDexBalances(0, 20)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(-1)

	s.aliceCancelsLimitSell("TokenB", -1, trancheKey)

	s.assertAliceBalances(50, 40)
	s.assertDexBalances(0, 10)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(0)
}

func (s *MsgServerTestSuite) TestCancelTwiceFails() {
	s.fundAliceBalances(50, 50)
	// CASE
	// Alice tries to cancel the same limit order twice

	trancheKey := s.aliceLimitSells("TokenB", 0, 10)

	s.assertAliceBalances(50, 40)
	s.assertDexBalances(0, 10)

	s.aliceCancelsLimitSell("TokenB", 0, trancheKey)

	s.assertAliceBalances(50, 50)
	s.assertDexBalances(0, 0)

	s.aliceCancelsLimitSellFails("TokenB", -1, trancheKey, types.ErrActiveLimitOrderNotFound)

}

func (s *MsgServerTestSuite) TestCancelPartiallyFilled() {
	s.fundAliceBalances(50, 0)
	s.fundBobBalances(0, 50)

	// GIVEN alice limit sells 50 TokenA
	trancheKey := s.aliceLimitSells("TokenA", 0, 50)
	// Bob swaps 25 TokenB for TokenA
	s.bobMarketSells("TokenB", 25)

	s.assertDexBalances(25, 25)
	s.assertAliceBalances(0, 0)

	//WHEN alice cancels her limit order
	s.aliceCancelsLimitSell("TokenA", 0, trancheKey)

	//Then alice gets back remaining 25 TokenA LO reserves
	s.assertAliceBalances(25, 0)
	s.assertDexBalances(0, 25)
}

func (s *MsgServerTestSuite) TestCancelPartiallyFilledMultiUser() {
	s.fundAliceBalances(50, 0)
	s.fundBobBalances(0, 50)
	s.fundCarolBalances(100, 0)

	// GIVEN alice limit sells 50 TokenA; carol sells 100 tokenA
	trancheKey := s.aliceLimitSells("TokenA", 0, 50)
	s.carolLimitSells("TokenA", 0, 100)
	// Bob swaps 25 TokenB for TokenA
	s.bobMarketSells("TokenB", 25)

	s.assertLimitLiquidityAtTick("TokenA", 0, 125)
	s.assertDexBalances(125, 25)
	s.assertAliceBalances(0, 0)

	//WHEN alice and carol cancel their limit orders
	s.aliceCancelsLimitSell("TokenA", 0, trancheKey)
	s.carolCancelsLimitSell("TokenA", 0, trancheKey)

	//THEN alice gets back 41 TokenA (125 * 1/3)
	s.assertAliceBalances(41, 0)

	//Carol gets back 83 TokenA (125 * 2/3)
	s.assertCarolBalances(83, 0)
	s.assertDexBalances(1, 25)
}

func (s *MsgServerTestSuite) TestCancelGoodTil() {
	s.fundAliceBalances(50, 0)
	tomorrow := time.Now().AddDate(0, 0, 1)
	// GIVEN alice limit sells 50 TokenA with goodTil date of tommrow
	trancheKey := s.aliceLimitSellsGoodTil("TokenA", 0, 50, tomorrow)
	s.assertLimitLiquidityAtTick("TokenA", 0, 50)
	s.assertNLimitOrderExpiration(1)

	// WHEN alice cancels the limit order
	s.aliceCancelsLimitSell("TokenA", 0, trancheKey)
	// THEN there is no liquidity left
	s.assertLimitLiquidityAtTick("TokenA", 0, 0)
	// and the LimitOrderExpiration has been removed
	s.assertNLimitOrderExpiration(0)
}

func (s *MsgServerTestSuite) TestCancelJITSameBlock() {
	s.fundAliceBalances(50, 0)
	// GIVEN alice limit sells 50 TokenA as JIT
	trancheKey := s.aliceLimitSells("TokenA", 0, 50, types.LimitOrderType_JUST_IN_TIME)
	s.assertLimitLiquidityAtTick("TokenA", 0, 50)
	s.assertNLimitOrderExpiration(1)

	// WHEN alice cancels the limit order
	s.aliceCancelsLimitSell("TokenA", 0, trancheKey)
	// THEN there is no liquidity left
	s.assertLimitLiquidityAtTick("TokenA", 0, 0)
	// and the LimitOrderExpiration has been removed
	s.assertNLimitOrderExpiration(0)
}

func (s *MsgServerTestSuite) TestCancelJITNextBlock() {
	s.fundAliceBalances(50, 0)
	// GIVEN alice limit sells 50 TokenA as JIT
	trancheKey := s.aliceLimitSells("TokenA", 0, 50, types.LimitOrderType_JUST_IN_TIME)
	s.assertLimitLiquidityAtTick("TokenA", 0, 50)
	s.assertNLimitOrderExpiration(1)

	// WHEN we move to block N+1
	s.nextBlockWithTime(time.Now())
	s.app.EndBlock(abci.RequestEndBlock{Height: 0})

	// THEN alice cancellation fails
	s.aliceCancelsLimitSellFails("TokenA", 0, trancheKey, types.ErrActiveLimitOrderNotFound)
	s.assertAliceBalances(0, 0)
}
