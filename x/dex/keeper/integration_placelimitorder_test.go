package keeper_test

import (
	"math"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/duality-labs/duality/x/dex/types"
)

// Core tests w/ GTC limitOrders //////////////////////////////////////////////
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

	// WHEN
	// place limit order for B at tick -1
	s.aliceLimitSells("TokenA", -1, 10)
	s.assertAliceBalances(30, 40)
	s.assertDexBalances(20, 10)

	// THEN
	// assert min, max not moved
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

	// WHEN
	// place limit order in spread (for B at tick -3)
	s.aliceLimitSells("TokenA", -3, 10)
	s.assertAliceBalances(20, 40)
	s.assertDexBalances(30, 10)

	// THEN
	// assert min not moved
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

	// WHEN
	// place limit order in spread (for A at tick 3)
	s.aliceLimitSells("TokenB", 3, 10)
	s.assertAliceBalances(40, 20)
	s.assertDexBalances(10, 30)

	// THEN
	// assert max not moved
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
	s.assertCurr1To0(-1)
	s.assertCurr0To1(math.MaxInt64)

	// WHEN
	// place limit order on same tick (for B at tick -1)
	s.aliceLimitSells("TokenA", -1, 10)

	// THEN
	// assert 20 of token A deposited at tick 0 fee 0 and ticks unchanged
	s.assertLimitLiquidityAtTick("TokenA", -1, 20)
	s.assertAliceLimitLiquidityAtTick("TokenA", 20, -1)
	s.assertAliceBalances(30, 50)
	s.assertDexBalances(20, 0)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(math.MaxInt64)
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
	s.assertCurr0To1(1)

	// WHEN
	// place limit order on same tick (for A at tick 1)
	s.aliceLimitSells("TokenB", 1, 10)

	// THEN
	// assert 20 of token B deposited at tick 0 fee 0 and ticks unchanged
	s.assertLimitLiquidityAtTick("TokenB", 1, 20)
	s.assertAliceLimitLiquidityAtTick("TokenB", 20, 1)
	s.assertAliceBalances(50, 30)
	s.assertDexBalances(0, 20)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(1)
}

// TODO: JCP delete me?
// func (s *MsgServerTestSuite) TestPlaceLimitOrderAboveEnemyLines() {
// 	s.fundAliceBalances(50, 50)

// 	// GIVEN
// 	// deposit 10 of token B at tick 0 fee 1
// 	s.aliceDeposits(NewDeposit(0, 10, 0, 0))
// 	s.assertAliceBalances(50, 40)
// 	s.assertDexBalances(0, 10)
// 	s.assertPoolLiquidity(0, 10, 0, 0)

// 	// WHEN
// 	// place limit order for token A above enemy lines at tick 5
// 	// THEN
// 	// deposit should fail with BEL error

// 	err := types.ErrPlaceLimitOrderBehindPairLiquidity
// 	s.assertAliceLimitSellFails(err, "TokenA", 5, 10)
// }

// func (s *MsgServerTestSuite) TestPlaceLimitOrderBelowEnemyLines() {
// 	s.fundAliceBalances(50, 50)

// 	// GIVEN
// 	// deposit 10 of token A at tick 0 fee 1
// 	s.aliceDeposits(NewDeposit(10, 0, 0, 0))
// 	s.assertAliceBalances(40, 50)
// 	s.assertDexBalances(10, 0)
// 	s.assertPoolLiquidity(10, 0, 0, 0)

// 	// WHEN
// 	// place limit order for token B below enemy lines at tick -5
// 	// THEN
// 	// deposit should fail with BEL error

// 	err := types.ErrPlaceLimitOrderBehindPairLiquidity
// 	s.assertAliceLimitSellFails(err, "TokenB", -5, 10)
// }

func (s *MsgServerTestSuite) TestPlaceLimitOrderNoLOPlaceLODoesntIncrementPlaceTrancheKey() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// no previous LO on existing tick
	s.aliceDeposits(NewDeposit(10, 0, 0, 0))
	s.assertPoolLiquidity(10, 0, 0, 0)
	s.assertFillAndPlaceTrancheKeys("TokenA", -1, "", "")

	// WHEN
	// placing order on same tick
	trancheKey := s.aliceLimitSells("TokenA", -1, 10)

	// THEN
	// fill and place tranche keys don't change
	s.assertFillAndPlaceTrancheKeys("TokenA", -1, trancheKey, trancheKey)
}

func (s *MsgServerTestSuite) TestPlaceLimitOrderUnfilledLOPlaceLODoesntIncrementPlaceTrancheKey() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// unfilled limit order exists on tick -1
	trancheKey := s.aliceLimitSells("TokenA", -1, 10)
	s.assertLimitLiquidityAtTick("TokenA", -1, 10)
	s.assertFillAndPlaceTrancheKeys("TokenA", -1, trancheKey, trancheKey)

	// WHEN
	// placing order on same tick
	s.aliceLimitSells("TokenA", -1, 10)

	// THEN
	// fill and place tranche keys don't change
	s.assertFillAndPlaceTrancheKeys("TokenA", -1, trancheKey, trancheKey)
}

func (s *MsgServerTestSuite) TestPlaceLimitOrderPartiallyFilledLOPlaceLOIncrementsPlaceTrancheKey() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 50)

	// GIVEN
	// partially filled limit order exists on tick -1
	trancheKey0 := s.aliceLimitSells("TokenA", -1, 10)
	s.bobMarketSells("TokenB", 5, 0)
	trancheKey1 := s.aliceLimitSells("TokenA", -1, 10)
	s.assertFillAndPlaceTrancheKeys("TokenA", -1, trancheKey0, trancheKey1)

	// WHEN
	// placing order on same tick
	s.aliceLimitSells("TokenA", -1, 10)

	// THEN
	// place tranche key changes
	s.assertFillAndPlaceTrancheKeys("TokenA", -1, trancheKey0, trancheKey1)
}

func (s *MsgServerTestSuite) TestPlaceLimitOrderFilledLOPlaceLODoesntIncrementsPlaceTrancheKey() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 50)

	// GIVEN
	// filled LO with partially filled place tranche
	trancheKey0 := s.aliceLimitSells("TokenA", -1, 10)
	s.bobMarketSells("TokenB", 10, 0)
	trancheKey1 := s.aliceLimitSells("TokenA", -1, 10)
	s.assertFillAndPlaceTrancheKeys("TokenA", -1, trancheKey0, trancheKey1)

	// WHEN
	// placing order on same tick
	s.aliceLimitSells("TokenA", -1, 5)

	// THEN
	// fill and place tranche keys don't change
	s.assertFillAndPlaceTrancheKeys("TokenA", -1, trancheKey0, trancheKey1)
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

func (s *MsgServerTestSuite) TestLimitOrderPartialFillDepositCancel() {
	s.fundAliceBalances(100, 100)
	s.fundBobBalances(100, 100)
	s.assertDexBalances(0, 0)

	trancheKey0 := s.aliceLimitSells("TokenB", 0, 50)

	s.assertAliceBalances(100, 50)
	s.assertBobBalances(100, 100)
	s.assertDexBalances(0, 50)
	s.assertCurrentTicks(math.MinInt64, 0)

	s.bobMarketSells("TokenA", 10, 10)

	s.assertAliceBalances(100, 50)
	s.assertBobBalances(90, 110)
	s.assertDexBalances(10, 40)
	s.assertCurrentTicks(math.MinInt64, 0)

	trancheKey1 := s.aliceLimitSells("TokenB", 0, 50)

	s.assertAliceBalances(100, 0)
	s.assertBobBalances(90, 110)
	s.assertDexBalances(10, 90)
	s.assertCurrentTicks(math.MinInt64, 0)

	s.aliceCancelsLimitSell("TokenB", 0, trancheKey0)

	s.assertAliceBalances(100, 40)
	s.assertBobBalances(90, 110)
	s.assertDexBalances(10, 50)
	s.assertCurrentTicks(math.MinInt64, 0)

	s.bobMarketSells("TokenA", 10, 10)

	s.assertAliceBalances(100, 40)
	s.assertBobBalances(80, 120)
	s.assertDexBalances(20, 40)

	s.aliceCancelsLimitSell("TokenB", 0, trancheKey1)

	s.assertAliceBalances(100, 80)
	s.assertBobBalances(80, 120)
	s.assertDexBalances(20, 0)

	s.aliceWithdrawsLimitSell("TokenB", 0, trancheKey0)

	s.assertAliceBalances(110, 80)
	s.assertBobBalances(80, 120)
	s.assertDexBalances(10, 0)

	s.aliceWithdrawsLimitSell("TokenB", 0, trancheKey1)

	s.assertAliceBalances(120, 80)
	s.assertBobBalances(80, 120)
	s.assertDexBalances(0, 0)
}

// Fill Or Kill limit orders ///////////////////////////////////////////////////////////
func (s *MsgServerTestSuite) TestPlaceLimitOrderFoKWithLPFills() {
	s.fundAliceBalances(10, 0)
	s.fundBobBalances(0, 20)
	// GIVEN LP liq at tick -1
	s.bobDeposits(NewDeposit(0, 20, -1, 0))
	//WHEN alice submits FoK limitOrder
	trancheKey := s.aliceLimitSells("TokenA", 0, 10, types.LimitOrderType_FILL_OR_KILL)
	s.assertAliceBalances(0, 0)
	// THEN alice's LimitOrder fills via swap

	// No maker LO is placed
	s.assertFillAndPlaceTrancheKeys("TokenA", 1, "", "")
	s.assertLimitLiquidityAtTick("TokenA", 1, 0)
	s.assertDexBalances(10, 20)

	// Alice can withdraw immediately
	s.aliceWithdrawsLimitSell("TokenA", 0, trancheKey)
	s.assertDexBalances(10, 10)
	s.assertAliceBalances(0, 10)
}

func (s *MsgServerTestSuite) TestPlaceLimitOrderFoKFailsWithInsufficientLiq() {
	s.fundAliceBalances(10, 0)
	s.fundBobBalances(0, 20)
	// GIVEN LP liq at tick -1 of 9 tokenB
	s.bobDeposits(NewDeposit(0, 9, -1, 0))
	//WHEN alice submits FoK limitOrder for 10 at tick 0 it fails
	s.assertAliceLimitSellFails(types.ErrFOKLimitOrderNotFilled, "TokenA", 0, 10, types.LimitOrderType_FILL_OR_KILL)

}

func (s *MsgServerTestSuite) TestPlaceLimitOrder0FoKFailsWithLowLimit() {
	s.fundAliceBalances(10, 0)
	s.fundBobBalances(0, 20)
	// GIVEN LP liq at tick -1 of 20 tokenB
	s.bobDeposits(NewDeposit(0, 20, -1, 0))
	//WHEN alice submits FoK limitOrder for 10 at tick -1 it fails
	s.assertAliceLimitSellFails(types.ErrFOKLimitOrderNotFilled, "TokenA", -1, 10, types.LimitOrderType_FILL_OR_KILL)

}

func (s *MsgServerTestSuite) TestPlaceLimitOrder1FoKFailsWithHighLimit() {
	s.fundAliceBalances(0, 10)
	s.fundBobBalances(20, 0)
	// GIVEN LP liq at tick 20 of 20 tokenA
	s.bobDeposits(NewDeposit(20, 0, 20, 0))
	//WHEN alice submits FoK limitOrder for 10 at tick -1 it fails
	s.assertAliceLimitSellFails(types.ErrFOKLimitOrderNotFilled, "TokenB", 21, 10, types.LimitOrderType_FILL_OR_KILL)

}

// Immediate Or Cancel LimitOrders ////////////////////////////////////////////////////////////////////
func (s *MsgServerTestSuite) TestPlaceLimitOrderIoCWithLPFills() {
	s.fundAliceBalances(10, 0)
	s.fundBobBalances(0, 20)
	// GIVEN LP liq at tick -1
	s.bobDeposits(NewDeposit(0, 20, -1, 0))
	//WHEN alice submits IoC limitOrder
	trancheKey := s.aliceLimitSells("TokenA", 0, 10, types.LimitOrderType_IMMEDIATE_OR_CANCEL)
	s.assertAliceBalances(0, 0)
	// THEN alice's LimitOrder fills via swap
	s.assertLimitLiquidityAtTick("TokenA", 1, 0)
	s.assertDexBalances(10, 20)
	// No maker LO is placed
	s.assertFillAndPlaceTrancheKeys("TokenA", 1, "", "")
	// Alice can withdraw immediately
	s.aliceWithdrawsLimitSell("TokenA", 0, trancheKey)
	s.assertDexBalances(10, 10)
	s.assertAliceBalances(0, 10)
}

func (s *MsgServerTestSuite) TestPlaceLimitOrderIoCWithLPPartialFill() {
	s.fundAliceBalances(10, 0)
	s.fundBobBalances(0, 20)
	// GIVEN LP of 5 tokenB at tick -1
	s.bobDeposits(NewDeposit(0, 5, -1, 0))
	//WHEN alice submits IoC limitOrder for 10 tokenA
	trancheKey := s.aliceLimitSells("TokenA", 0, 10, types.LimitOrderType_IMMEDIATE_OR_CANCEL)
	s.assertAliceBalances(5, 0)
	// THEN alice's LimitOrder swap 5 TokenA

	// No maker LO is placed
	s.assertFillAndPlaceTrancheKeys("TokenA", 1, "", "")
	s.assertLimitLiquidityAtTick("TokenA", 1, 0)
	s.assertDexBalances(5, 5)

	// Alice can withdraw her partial fill immediately
	s.aliceWithdrawsLimitSell("TokenA", 0, trancheKey)
	s.assertDexBalances(5, 0)
	s.assertAliceBalances(5, 5)

}
