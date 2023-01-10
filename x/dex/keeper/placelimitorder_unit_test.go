package keeper_test

import (
	"math"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (s *MsgServerTestSuite) TestPlaceLimitOrderInSpread1To0() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// create spread around -5, 5
	s.aliceDeposits(NewDeposit(10, 10, 0, 2))
	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertCurr1To0(-5)
	s.assertCurr0To1(5)

	// WHEN
	// place limit order for B at tick -1
	s.aliceLimitSells("TokenA", -1, 10)
	s.assertAliceBalances(30, 40)
	s.assertDexBalances(20, 10)

	// THEN
	// assert currentTick1To0 moved
	s.assertCurr1To0(-1)
}
func (s *MsgServerTestSuite) TestPlaceLimitOrderInSpread0To1() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// create spread around -5, 5
	s.aliceDeposits(NewDeposit(10, 10, 0, 2))
	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertCurr1To0(-5)
	s.assertCurr0To1(5)

	// WHEN
	// place limit order for A at tick 1
	s.aliceLimitSells("TokenB", 1, 10)
	s.assertAliceBalances(40, 30)
	s.assertDexBalances(10, 20)

	// THEN
	// assert currentTick0To1 moved
	s.assertCurr0To1(1)
}

func (s *MsgServerTestSuite) TestPlaceLimitOrderInSpreadMinMaxNotAdjusted() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// create spread around -5, 5
	s.aliceDeposits(NewDeposit(10, 10, 0, 2))
	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertMinTick(-5)
	s.assertMaxTick(5)

	// WHEN
	// place limit order for B at tick -1
	s.aliceLimitSells("TokenA", -1, 10)
	s.assertAliceBalances(30, 40)
	s.assertDexBalances(20, 10)

	// THEN
	// assert min, max not moved
	s.assertMinTick(-5)
	s.assertMaxTick(5)
}

func (s *MsgServerTestSuite) TestPlaceLimitOrderOutOfSpread0To1NotAdjusted() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// create spread around -1, 1
	s.aliceDeposits(NewDeposit(10, 10, 0, 0))
	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(1)

	// WHEN
	// place limit order out of spread (for A at tick 3)
	s.aliceLimitSells("TokenB", 3, 10)
	s.assertAliceBalances(40, 30)
	s.assertDexBalances(10, 20)

	// THEN
	// assert currentTick0To1 not moved
	s.assertCurr0To1(1)
}

func (s *MsgServerTestSuite) TestPlaceLimitOrderOutOfSpread1To0NotAdjusted() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// create spread around -1, 1
	s.aliceDeposits(NewDeposit(10, 10, 0, 0))
	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(1)

	// WHEN
	// place limit order out of spread (for B at tick -3)
	s.aliceLimitSells("TokenA", -3, 10)
	s.assertAliceBalances(30, 40)
	s.assertDexBalances(20, 10)

	// THEN
	// assert currentTick1To0 not moved
	s.assertCurr1To0(-1)
}

func (s *MsgServerTestSuite) TestPlaceLimitOrderOutOfSpreadMinAdjusted() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// create spread around -1, 1
	s.aliceDeposits(NewDeposit(10, 10, 0, 0))
	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(1)

	// WHEN
	// place limit order out of spread (for B at tick -3)
	s.aliceLimitSells("TokenA", -3, 10)
	s.assertAliceBalances(30, 40)
	s.assertDexBalances(20, 10)

	// THEN
	// assert min moved
	s.assertMinTick(-3)
}

func (s *MsgServerTestSuite) TestPlaceLimitOrderOutOfSpreadMaxAdjusted() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// create spread around -1, 1
	s.aliceDeposits(NewDeposit(10, 10, 0, 0))
	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(1)

	// WHEN
	// place limit order out of spread (for A at tick 3)
	s.aliceLimitSells("TokenB", 3, 10)
	s.assertAliceBalances(40, 30)
	s.assertDexBalances(10, 20)

	// THEN
	// assert max moved
	s.assertMaxTick(3)
}

func (s *MsgServerTestSuite) TestPlaceLimitOrderOutOfSpreadMinNotAdjusted() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// create spread around -1, 1
	s.aliceDeposits(NewDeposit(10, 10, 0, 0))
	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(1)
	// deposit new min at -5
	s.aliceDeposits(NewDeposit(10, 0, 0, 2))
	s.assertAliceBalances(30, 40)
	s.assertDexBalances(20, 10)
	s.assertMinTick(-5)
	s.assertMaxTick(1)

	// WHEN
	// place limit order in spread (for B at tick -3)
	s.aliceLimitSells("TokenA", -3, 10)
	s.assertAliceBalances(20, 40)
	s.assertDexBalances(30, 10)

	// THEN
	// assert min not moved
	s.assertMinTick(-5)
}

func (s *MsgServerTestSuite) TestPlaceLimitOrderOutOfSpreadMaxNotAdjusted() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// create spread around -1, 1
	s.aliceDeposits(NewDeposit(10, 10, 0, 0))
	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(1)
	// deposit new max at 5
	s.aliceDeposits(NewDeposit(0, 10, 0, 2))
	s.assertAliceBalances(40, 30)
	s.assertDexBalances(10, 20)
	s.assertMinTick(-1)
	s.assertMaxTick(5)

	// WHEN
	// place limit order in spread (for A at tick 3)
	s.aliceLimitSells("TokenB", 3, 10)
	s.assertAliceBalances(40, 20)
	s.assertDexBalances(10, 30)

	// THEN
	// assert max not moved
	s.assertMaxTick(5)
}

func (s *MsgServerTestSuite) TestPlaceLimitOrderExistingLiquidityA() {
	//TODO: this fails because PairInit doesn't account for single sided liquidity
	s.fundAliceBalances(50, 50)

	// GIVEN
	// deposit 10 of token A at tick 0 fee 1
	s.aliceLimitSells("TokenA", -1, 10)
	s.assertAliceBalances(40, 50)
	s.assertDexBalances(10, 0)
	s.assertLimitLiquidityAtTick("TokenA", -1, 10)
	s.assertAliceLimitLiquidityAtTick("TokenA", 10, -1)
	s.assertMinTick(-1)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)

	// WHEN
	// place limit order on same tick (for B at tick -1)
	s.aliceLimitSells("TokenA", -1, 10)

	// THEN
	// assert 20 of token A deposited at tick 0 fee 0 and ticks unchanged
	s.assertLimitLiquidityAtTick("TokenA", -1, 20)
	s.assertAliceLimitLiquidityAtTick("TokenA", 20, -1)
	s.assertAliceBalances(30, 50)
	s.assertDexBalances(20, 0)
	s.assertMinTick(-1)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
}

func (s *MsgServerTestSuite) TestPlaceLimitOrderExistingLiquidityB() {
	// TODO: this fails because PairInit doesn't account for single sided liquidity
	s.fundAliceBalances(50, 50)

	// GIVEN
	// deposit 10 of token B at tick 1 fee 0
	s.aliceLimitSells("TokenB", 1, 10)
	s.assertAliceBalances(50, 40)
	s.assertDexBalances(0, 10)
	s.assertLimitLiquidityAtTick("TokenB", 1, 10)
	s.assertAliceLimitLiquidityAtTick("TokenB", 10, 1)
	s.assertCurr1To0(math.MinInt64)
	s.assertMinTick(math.MaxInt64)
	s.assertCurr0To1(1)
	s.assertMaxTick(1)

	// WHEN
	// place limit order on same tick (for A at tick 1)
	s.aliceLimitSells("TokenB", 1, 10)

	// THEN
	// assert 20 of token B deposited at tick 0 fee 0 and ticks unchanged
	s.assertLimitLiquidityAtTick("TokenB", 1, 20)
	s.assertAliceLimitLiquidityAtTick("TokenB", 20, 1)
	s.assertAliceBalances(50, 30)
	s.assertDexBalances(0, 20)
	s.assertMinTick(math.MaxInt64)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(1)
	s.assertMaxTick(1)
}

func (s *MsgServerTestSuite) TestPlaceLimitOrderBelowEnemyLines() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// deposit 10 of token B at tick 0 fee 1
	s.aliceDeposits(NewDeposit(0, 10, 0, 0))
	s.assertAliceBalances(50, 40)
	s.assertDexBalances(0, 10)
	s.assertPoolLiquidity(0, 10, 0, 0)

	// WHEN
	// place limit order for token A below enemy lines at tick -5
	// THEN
	// deposit should fail with BEL error

	err := types.ErrPlaceLimitOrderBehindPairLiquidity // TODO: this needs to be changed to a more specific error type
	s.assertAliceLimitSellFails(err, "TokenA", 5, 10)
}

func (s *MsgServerTestSuite) TestPlaceLimitOrderAboveEnemyLines() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// deposit 10 of token A at tick 0 fee 1
	s.aliceDeposits(NewDeposit(10, 0, 0, 0))
	s.assertAliceBalances(40, 50)
	s.assertDexBalances(10, 0)
	s.assertPoolLiquidity(10, 0, 0, 0)

	// WHEN
	// place limit order for token B above enemy lines at tick 5
	// THEN
	// deposit should fail with BEL error

	err := types.ErrPlaceLimitOrderBehindPairLiquidity // TODO: this needs to be changed to a more specific error type
	s.assertAliceLimitSellFails(err, "TokenB", -5, 10)
}

func (s *MsgServerTestSuite) TestPlaceLimitOrderNoLOPlaceLODoesntIncrementPlaceTrancheKey() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// no previous LO on existing tick
	s.aliceDeposits(NewDeposit(10, 0, 0, 0))
	s.assertPoolLiquidity(10, 0, 0, 0)
	s.assertFillAndPlaceTrancheKeys("TokenA", -1, 0, 0)

	// WHEN
	// placing order on same tick
	s.aliceLimitSells("TokenA", -1, 10)

	// THEN
	// fill and place tranche keys don't change
	s.assertFillAndPlaceTrancheKeys("TokenA", -1, 0, 0)
}

func (s *MsgServerTestSuite) TestPlaceLimitOrderUnfilledLOPlaceLODoesntIncrementPlaceTrancheKey() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// unfilled limit order exists on tick -1
	s.aliceLimitSells("TokenA", -1, 10)
	s.assertLimitLiquidityAtTick("TokenA", -1, 10)
	s.assertFillAndPlaceTrancheKeys("TokenA", -1, 0, 0)

	// WHEN
	// placing order on same tick
	s.aliceLimitSells("TokenA", -1, 10)

	// THEN
	// fill and place tranche keys don't change
	s.assertFillAndPlaceTrancheKeys("TokenA", -1, 0, 0)
}

func (s *MsgServerTestSuite) TestPlaceLimitOrderPartiallyFilledLOPlaceLOIncrementsPlaceTrancheKey() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 50)

	// GIVEN
	// partially filled limit order exists on tick -1
	s.aliceLimitSells("TokenA", -1, 10)
	s.bobMarketSells("TokenB", 5, 0)
	s.assertFillAndPlaceTrancheKeys("TokenA", -1, 0, 0)

	// WHEN
	// placing order on same tick
	s.aliceLimitSells("TokenA", -1, 10)

	// THEN
	// place tranche key changes
	s.assertFillAndPlaceTrancheKeys("TokenA", -1, 0, 1)
}

func (s *MsgServerTestSuite) TestPlaceLimitOrderFilledLOPlaceLODoesntIncrementsPlaceTrancheKey() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 50)

	// GIVEN
	// filled LO with partially filled place tranche
	s.aliceLimitSells("TokenA", -1, 10)
	s.bobMarketSells("TokenB", 10, 0)
	s.aliceLimitSells("TokenA", -1, 10)
	s.assertFillAndPlaceTrancheKeys("TokenA", -1, 0, 1)

	// WHEN
	// placing order on same tick
	s.aliceLimitSells("TokenA", -1, 5)

	// THEN
	// fill and place tranche keys don't change
	s.assertFillAndPlaceTrancheKeys("TokenA", -1, 0, 1)
}

func (s *MsgServerTestSuite) TestPlaceLimitOrderInsufficientFunds() {
	// GIVEN
	// alice has no funds
	s.assertAliceBalances(0, 0)

	// WHEN
	// place limit order selling non zero amount of token A for token B
	// THEN
	// deposit should fail with InsufficientFunds error

	err := sdkerrors.ErrInsufficientFunds
	s.assertAliceLimitSellFails(err, "TokenA", 0, 10)
}
