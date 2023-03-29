package keeper_test

import (
	"math"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/duality-labs/duality/x/dex/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// Core tests w/ GTC limitOrders //////////////////////////////////////////////
func (s *MsgServerTestSuite) TestPlaceLimitOrderInSpread1To0() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// create spread around -5, 5
	s.aliceDeposits(NewDeposit(10, 10, 0, 5))
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
	s.aliceDeposits(NewDeposit(10, 10, 0, 5))
	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertCurr1To0(-5)
	s.assertCurr0To1(5)

	// WHEN
	// place limit order for A at tick 1
	s.aliceLimitSells("TokenB", -1, 10)
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
	s.aliceDeposits(NewDeposit(10, 10, 0, 5))
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
	s.aliceDeposits(NewDeposit(10, 10, 0, 1))
	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(1)

	// WHEN
	// place limit order out of spread (for A at tick 3)
	s.aliceLimitSells("TokenB", -3, 10)
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
	s.aliceDeposits(NewDeposit(10, 10, 0, 1))
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
	s.aliceDeposits(NewDeposit(10, 10, 0, 1))
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
	s.aliceDeposits(NewDeposit(10, 10, 0, 1))
	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(1)

	// WHEN
	// place limit order out of spread (for A at tick 3)
	s.aliceLimitSells("TokenB", -3, 10)
	s.assertAliceBalances(40, 30)
	s.assertDexBalances(10, 20)

	// THEN
	// assert max moved
}

func (s *MsgServerTestSuite) TestPlaceLimitOrderOutOfSpreadMinNotAdjusted() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// create spread around -1, 1
	s.aliceDeposits(NewDeposit(10, 10, 0, 1))
	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(1)
	// deposit new min at -5
	s.aliceDeposits(NewDeposit(10, 0, 0, 5))
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
	s.aliceDeposits(NewDeposit(10, 10, 0, 1))
	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(1)
	// deposit new max at 5
	s.aliceDeposits(NewDeposit(0, 10, 0, 5))
	s.assertAliceBalances(40, 30)
	s.assertDexBalances(10, 20)

	// WHEN
	// place limit order in spread (for A at tick 3)
	s.aliceLimitSells("TokenB", -3, 10)
	s.assertAliceBalances(40, 20)
	s.assertDexBalances(10, 30)

	// THEN
	// assert max not moved
}

func (s *MsgServerTestSuite) TestPlaceLimitOrderExistingLiquidityA() {
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
	s.fundAliceBalances(50, 50)

	// GIVEN
	// deposit 10 of token B at tick 1 fee 0
	s.aliceLimitSells("TokenB", -1, 10)
	s.assertAliceBalances(50, 40)
	s.assertDexBalances(0, 10)
	s.assertLimitLiquidityAtTick("TokenB", 1, 10)
	s.assertAliceLimitLiquidityAtTick("TokenB", 10, 1)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(1)

	// WHEN
	// place limit order on same tick (for A at tick 1)
	s.aliceLimitSells("TokenB", -1, 10)

	// THEN
	// assert 20 of token B deposited at tick 0 fee 0 and ticks unchanged
	s.assertLimitLiquidityAtTick("TokenB", 1, 20)
	s.assertAliceLimitLiquidityAtTick("TokenB", 20, 1)
	s.assertAliceBalances(50, 30)
	s.assertDexBalances(0, 20)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(1)
}

func (s *MsgServerTestSuite) TestPlaceLimitOrderNoLOPlaceLODoesntIncrementPlaceTrancheKey() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// no previous LO on existing tick
	s.aliceDeposits(NewDeposit(10, 0, 0, 1))
	s.assertPoolLiquidity(10, 0, 0, 1)
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
	s.bobMarketSells("TokenB", 5)
	s.assertFillAndPlaceTrancheKeys("TokenA", -1, trancheKey0, trancheKey0)

	// WHEN
	// placing order on same tick
	trancheKey1 := s.aliceLimitSells("TokenA", -1, 10)

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
	s.bobMarketSells("TokenB", 10)
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

	s.bobMarketSells("TokenA", 10)

	s.assertAliceBalances(100, 50)
	s.assertBobBalances(90, 110)
	s.assertDexBalances(10, 40)
	s.assertCurrentTicks(math.MinInt64, 0)

	trancheKey1 := s.aliceLimitSells("TokenB", 0, 50)

	s.assertAliceBalances(100, 0)
	s.assertBobBalances(90, 110)
	s.assertDexBalances(10, 90)
	s.assertCurrentTicks(math.MinInt64, 0)

	s.aliceCancelsLimitSell(trancheKey0)

	s.assertAliceBalances(100, 40)
	s.assertBobBalances(90, 110)
	s.assertDexBalances(10, 50)
	s.assertCurrentTicks(math.MinInt64, 0)

	s.bobMarketSells("TokenA", 10)

	s.assertAliceBalances(100, 40)
	s.assertBobBalances(80, 120)
	s.assertDexBalances(20, 40)

	s.aliceCancelsLimitSell(trancheKey1)

	s.assertAliceBalances(100, 80)
	s.assertBobBalances(80, 120)
	s.assertDexBalances(20, 0)

	s.aliceWithdrawsLimitSell(trancheKey0)

	s.assertAliceBalances(110, 80)
	s.assertBobBalances(80, 120)
	s.assertDexBalances(10, 0)

	s.aliceWithdrawsLimitSell(trancheKey1)

	s.assertAliceBalances(120, 80)
	s.assertBobBalances(80, 120)
	s.assertDexBalances(0, 0)
}

// Fill Or Kill limit orders ///////////////////////////////////////////////////////////
func (s *MsgServerTestSuite) TestPlaceLimitOrderFoKWithLPFills() {
	s.fundAliceBalances(10, 0)
	s.fundBobBalances(0, 20)
	// GIVEN LP liq at tick -1
	s.bobDeposits(NewDeposit(0, 20, -1, 1))
	//WHEN alice submits FoK limitOrder
	trancheKey := s.aliceLimitSells("TokenA", 0, 10, types.LimitOrderType_FILL_OR_KILL)
	s.assertAliceBalances(0, 0)
	// THEN alice's LimitOrder fills via swap

	// No maker LO is placed
	s.assertFillAndPlaceTrancheKeys("TokenA", 1, "", "")
	s.assertLimitLiquidityAtTick("TokenA", 1, 0)
	s.assertDexBalances(10, 20)

	// Alice can withdraw immediately
	s.aliceWithdrawsLimitSell(trancheKey)
	s.assertDexBalances(10, 10)
	s.assertAliceBalances(0, 10)
}

func (s *MsgServerTestSuite) TestPlaceLimitOrderFoKFailsWithInsufficientLiq() {
	s.fundAliceBalances(10, 0)
	s.fundBobBalances(0, 20)
	// GIVEN LP liq at tick -1 of 9 tokenB
	s.bobDeposits(NewDeposit(0, 9, -1, 1))
	//WHEN alice submits FoK limitOrder for 10 at tick 0 it fails
	s.assertAliceLimitSellFails(types.ErrFoKLimitOrderNotFilled, "TokenA", 0, 10, types.LimitOrderType_FILL_OR_KILL)

}

func (s *MsgServerTestSuite) TestPlaceLimitOrder0FoKFailsWithLowLimit() {
	s.fundAliceBalances(10, 0)
	s.fundBobBalances(0, 20)
	// GIVEN LP liq at tick -1 of 20 tokenB
	s.bobDeposits(NewDeposit(0, 20, -1, 1))
	//WHEN alice submits FoK limitOrder for 10 at tick -1 it fails
	s.assertAliceLimitSellFails(types.ErrFoKLimitOrderNotFilled, "TokenA", -1, 10, types.LimitOrderType_FILL_OR_KILL)

}

func (s *MsgServerTestSuite) TestPlaceLimitOrder1FoKFailsWithHighLimit() {
	s.fundAliceBalances(0, 10)
	s.fundBobBalances(20, 0)
	// GIVEN LP liq at tick 20 of 20 tokenA
	s.bobDeposits(NewDeposit(20, 0, 20, 1))
	//WHEN alice submits FoK limitOrder for 10 at tick -1 it fails
	s.assertAliceLimitSellFails(types.ErrFoKLimitOrderNotFilled, "TokenB", -21, 10, types.LimitOrderType_FILL_OR_KILL)
}

// Immediate Or Cancel LimitOrders ////////////////////////////////////////////////////////////////////
func (s *MsgServerTestSuite) TestPlaceLimitOrderIoCWithLPFills() {
	s.fundAliceBalances(10, 0)
	s.fundBobBalances(0, 20)
	// GIVEN LP liq at tick -1
	s.bobDeposits(NewDeposit(0, 20, -1, 1))
	//WHEN alice submits IoC limitOrder
	trancheKey := s.aliceLimitSells("TokenA", 0, 10, types.LimitOrderType_IMMEDIATE_OR_CANCEL)
	s.assertAliceBalances(0, 0)
	// THEN alice's LimitOrder fills via swap
	s.assertLimitLiquidityAtTick("TokenA", 1, 0)
	s.assertDexBalances(10, 20)
	// No maker LO is placed
	s.assertFillAndPlaceTrancheKeys("TokenA", 1, "", "")
	// Alice can withdraw immediately
	s.aliceWithdrawsLimitSell(trancheKey)
	s.assertDexBalances(10, 10)
	s.assertAliceBalances(0, 10)
}

func (s *MsgServerTestSuite) TestPlaceLimitOrderIoCWithLPPartialFill() {
	s.fundAliceBalances(10, 0)
	s.fundBobBalances(0, 20)
	// GIVEN LP of 5 tokenB at tick -1
	s.bobDeposits(NewDeposit(0, 5, -1, 1))
	//WHEN alice submits IoC limitOrder for 10 tokenA
	trancheKey := s.aliceLimitSells("TokenA", 0, 10, types.LimitOrderType_IMMEDIATE_OR_CANCEL)
	s.assertAliceBalances(5, 0)
	// THEN alice's LimitOrder swap 5 TokenA

	// No maker LO is placed
	s.assertFillAndPlaceTrancheKeys("TokenA", 1, "", "")
	s.assertLimitLiquidityAtTick("TokenA", 1, 0)
	s.assertDexBalances(5, 5)

	// Alice can withdraw her partial fill immediately
	s.aliceWithdrawsLimitSell(trancheKey)
	s.assertDexBalances(5, 0)
	s.assertAliceBalances(5, 5)

}

func (s *MsgServerTestSuite) TestPlaceLimitOrderIoCWithLPNoFill() {
	s.fundAliceBalances(10, 0)
	s.fundBobBalances(0, 20)
	// GIVEN LP of 5 tokenB at tick -1
	s.bobDeposits(NewDeposit(0, 5, -1, 1))
	//WHEN alice submits IoC limitOrder for 10 tokenA below current 0To1 price
	s.aliceLimitSells("TokenA", -1, 10, types.LimitOrderType_IMMEDIATE_OR_CANCEL)

	// THEN alice's LimitOrder doesn't fill and is canceled
	s.assertDexBalances(0, 5)
	s.assertAliceBalances(10, 0)
	// No maker LO is placed
	s.assertFillAndPlaceTrancheKeys("TokenA", 1, "", "")
	s.assertLimitLiquidityAtTick("TokenA", 1, 0)

}

// Just In Time Limit Orders //////////////////////////////////////////////////

func (s *MsgServerTestSuite) TestPlaceLimitOrderJITFills() {
	s.fundAliceBalances(10, 0)
	s.fundBobBalances(0, 20)

	//GIVEN Alice submits JIT limitOrder for 10 tokenA at tick 0
	trancheKey := s.aliceLimitSells("TokenA", 0, 10, types.LimitOrderType_JUST_IN_TIME)
	s.assertLimitLiquidityAtTick("TokenA", 0, 10)
	s.assertAliceBalances(0, 0)

	// WHEN bob swaps through all the liquidity
	s.bobMarketSells("TokenB", 10)

	// THEN all liquidity is depleted
	s.assertLimitLiquidityAtTick("TokenA", 0, 0)
	//Alice can withdraw 10 TokenB
	s.aliceWithdrawsLimitSell(trancheKey)
	s.assertAliceBalances(0, 10)

}

func (s *MsgServerTestSuite) TestPlaceLimitOrderJITBehindEnemyLines() {
	s.fundAliceBalances(10, 0)
	s.fundBobBalances(0, 20)

	//GIVEN 10 LP liquidity for token exists at tick 0
	s.bobDeposits(NewDeposit(0, 10, 0, 1))

	// WHEN alice places a JIT limit order for TokenA at tick 1 (above enemy lines)
	trancheKey := s.aliceLimitSells("TokenA", 1, 10, types.LimitOrderType_JUST_IN_TIME)
	s.assertLimitLiquidityAtTick("TokenA", 1, 10)
	s.assertAliceBalances(0, 0)
	// AND bob swaps through all the liquidity
	s.bobMarketSells("TokenB", 10)

	// THEN all liquidity is depleted
	s.assertLimitLiquidityAtTick("TokenA", 1, 0)
	//Alice can withdraw 9 TokenB
	s.aliceWithdrawsLimitSell(trancheKey)
	s.assertAliceBalances(0, 9)

}

func (s *MsgServerTestSuite) TestPlaceLimitOrderJITNextBlock() {
	s.fundAliceBalances(10, 0)
	s.fundBobBalances(0, 20)

	//GIVEN Alice submits JIT limitOrder for 10 tokenA at tick 0 for block N
	trancheKey := s.aliceLimitSells("TokenA", 0, 10, types.LimitOrderType_JUST_IN_TIME)
	s.assertLimitLiquidityAtTick("TokenA", 0, 10)
	s.assertAliceBalances(0, 0)

	// WHEN we move to block N+1
	s.nextBlockWithTime(time.Now())
	s.app.EndBlock(abci.RequestEndBlock{Height: 0})

	// THEN there is no liquidity available
	s.bobMarketSellFails(types.ErrInsufficientLiquidity, "TokenB", 10)
	s.assertLimitLiquidityAtTick("TokenA", 0, 0)
	//Alice can withdraw the entirety of the unfilled limitOrder
	s.aliceWithdrawsLimitSell(trancheKey)
	s.assertAliceBalances(10, 0)

}

// GoodTilLimitOrders //////////////////////////////////////////////////

func (s *MsgServerTestSuite) TestPlaceLimitOrderGoodTilFills() {
	s.fundAliceBalances(10, 0)
	s.fundBobBalances(0, 20)
	tomorrow := time.Now().AddDate(0, 0, 1)
	//GIVEN Alice submits JIT limitOrder for 10 tokenA expiring tomorrow
	trancheKey := s.aliceLimitSellsGoodTil("TokenA", 0, 10, tomorrow)
	s.assertLimitLiquidityAtTick("TokenA", 0, 10)
	s.assertAliceBalances(0, 0)

	// WHEN bob swaps through all the liquidity
	s.bobMarketSells("TokenB", 10)

	// THEN all liquidity is depleted
	s.assertLimitLiquidityAtTick("TokenA", 0, 0)
	//Alice can withdraw 10 TokenB
	s.aliceWithdrawsLimitSell(trancheKey)
	s.assertAliceBalances(0, 10)

}

func (s *MsgServerTestSuite) TestPlaceLimitOrderGoodTilExpires() {
	s.fundAliceBalances(10, 0)
	s.fundBobBalances(0, 20)
	tomorrow := time.Now().AddDate(0, 0, 1)
	//GIVEN Alice submits JIT limitOrder for 10 tokenA expiring tomorrow
	trancheKey := s.aliceLimitSellsGoodTil("TokenA", 0, 10, tomorrow)
	s.assertLimitLiquidityAtTick("TokenA", 0, 10)
	s.assertAliceBalances(0, 0)

	// When two days go by and multiple blocks are created (ie. purge is run)
	s.nextBlockWithTime(time.Now().AddDate(0, 0, 2))
	s.app.EndBlock(abci.RequestEndBlock{Height: 0})
	// THEN there is no liquidity available
	s.bobMarketSellFails(types.ErrInsufficientLiquidity, "TokenB", 10)
	s.assertLimitLiquidityAtTick("TokenA", 0, 0)
	//Alice can withdraw the entirety of the unfilled limitOrder
	s.aliceWithdrawsLimitSell(trancheKey)
	s.assertAliceBalances(10, 0)

}

func (s *MsgServerTestSuite) TestPlaceLimitOrderGoodTilExpiresNotPurged() {
	// This is testing the case where the limitOrder has expired but has not yet been purged
	s.fundAliceBalances(10, 0)
	s.fundBobBalances(0, 20)
	tomorrow := time.Now().AddDate(0, 0, 1)
	//GIVEN Alice submits JIT limitOrder for 10 tokenA expiring tomorrow
	trancheKey := s.aliceLimitSellsGoodTil("TokenA", 0, 10, tomorrow)
	s.assertLimitLiquidityAtTick("TokenA", 0, 10)
	s.assertAliceBalances(0, 0)

	// When two days go by
	// for simplicity sake we never run endBlock, it reality it would be run, but gas limit would be hit
	s.nextBlockWithTime(time.Now().AddDate(0, 0, 2))

	// THEN there is no liquidity available
	s.bobMarketSellFails(types.ErrInsufficientLiquidity, "TokenB", 10)
	s.assertLimitLiquidityAtTick("TokenA", 0, 0)
	//Alice can cancel the entirety of the unfilled limitOrder
	s.aliceCancelsLimitSell(trancheKey)
	s.assertAliceBalances(10, 0)

}

func (s *MsgServerTestSuite) TestPlaceLimitOrderGoodTilHandlesTimezoneCorrectly() {
	s.fundAliceBalances(10, 0)
	timeInPST, _ := time.Parse(time.RFC3339, "2050-01-02T15:04:05-08:00")
	trancheKey := s.aliceLimitSellsGoodTil("TokenA", 0, 10, timeInPST)
	tranche, _ := s.app.DexKeeper.GetLimitOrderTranche(s.ctx, defaultPairId, "TokenA", 0, trancheKey)

	s.Assert().Equal(tranche.ExpirationTime.Unix(), timeInPST.Unix())

}

func (s *MsgServerTestSuite) TestPlaceLimitOrderGoodTilAlreadyExpiredFails() {
	s.fundAliceBalances(10, 0)

	now := time.Now()
	yesterday := time.Now().AddDate(0, 0, -1)
	s.nextBlockWithTime(now)

	_, err := s.msgServer.PlaceLimitOrder(s.goCtx, &types.MsgPlaceLimitOrder{
		Creator:        s.alice.String(),
		Receiver:       s.alice.String(),
		TokenIn:        "TokenA",
		TokenOut:       "TokenB",
		TickIndex:      0,
		AmountIn:       sdk.NewInt(50),
		OrderType:      types.LimitOrderType_GOOD_TIL_TIME,
		ExpirationTime: &yesterday,
	})
	s.Assert().ErrorIs(err, types.ErrExpirationTimeInPast)
}
