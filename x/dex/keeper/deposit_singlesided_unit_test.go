package keeper_test

import (
	"math"

	"github.com/NicholasDotSol/duality/x/dex/types"
)

func (s *MsgServerTestSuite) TestSingleSidedDepositInSpread1To0() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// create spread around -5, 5
	s.aliceDeposits(NewDeposit(10, 10, 0, 2))
	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertCurr1To0(-5)
	s.assertCurr0To1(5)

	// WHEN
	// deposit in spread (10 of B at tick 1)
	s.aliceDeposits(NewDeposit(10, 0, 0, 0))
	s.assertAliceBalances(30, 40)
	s.assertDexBalances(20, 10)

	// THEN
	// assert currentTick1To0 moved
	s.assertCurr1To0(-1)
}
func (s *MsgServerTestSuite) TestSingleSidedDepositInSpread0To1() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// create spread around -5, 5
	s.aliceDeposits(NewDeposit(10, 10, 0, 2))
	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertCurr1To0(-5)
	s.assertCurr0To1(5)

	// WHEN
	// deposit in spread (10 of B at tick 1)
	s.aliceDeposits(NewDeposit(0, 10, 0, 0))
	s.assertAliceBalances(40, 30)
	s.assertDexBalances(10, 20)

	// THEN
	// assert currentTick0To1 moved
	s.assertCurr0To1(1)
}

func (s *MsgServerTestSuite) TestSingleSidedDepositInSpreadMinMaxNotAdjusted() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// create spread around -5, 5
	s.aliceDeposits(NewDeposit(10, 10, 0, 2))
	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertMinTick(-5)
	s.assertMaxTick(5)

	// WHEN
	// deposit in spread (10 of B at tick 1)
	s.aliceDeposits(NewDeposit(10, 0, 0, 0))
	s.assertAliceBalances(30, 40)
	s.assertDexBalances(20, 10)

	// THEN
	// assert min, max not moved
	s.assertMinTick(-5)
	s.assertMaxTick(5)
}

func (s *MsgServerTestSuite) TestSingleSidedDepositOutOfSpread0To1NotAdjusted() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// create spread around -1, 1
	s.aliceDeposits(NewDeposit(10, 10, 0, 0))
	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(1)

	// WHEN
	// deposit in spread (10 of A at tick -3)
	s.aliceDeposits(NewDeposit(0, 10, 0, 1))
	s.assertAliceBalances(40, 30)
	s.assertDexBalances(10, 20)

	// THEN
	// assert currentTick0To1 not moved
	s.assertCurr0To1(1)
}

func (s *MsgServerTestSuite) TestSingleSidedDepositOutOfSpread1To0NotAdjusted() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// create spread around -1, 1
	s.aliceDeposits(NewDeposit(10, 10, 0, 0))
	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(1)

	// WHEN
	// deposit in spread (10 of B at tick 3)
	s.aliceDeposits(NewDeposit(10, 0, 0, 1))
	s.assertAliceBalances(30, 40)
	s.assertDexBalances(20, 10)

	// THEN
	// assert currentTick1To0 not moved
	s.assertCurr1To0(-1)
}

func (s *MsgServerTestSuite) TestSingleSidedDepositOutOfSpreadMinAdjusted() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// create spread around -1, 1
	s.aliceDeposits(NewDeposit(10, 10, 0, 0))
	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(1)

	// WHEN
	// deposit in spread (10 of A at tick 3)
	s.aliceDeposits(NewDeposit(10, 0, 0, 1))
	s.assertAliceBalances(30, 40)
	s.assertDexBalances(20, 10)

	// THEN
	// assert min moved
	s.assertMinTick(-3)
}

func (s *MsgServerTestSuite) TestSingleSidedDepositOutOfSpreadMaxAdjusted() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// create spread around -1, 1
	s.aliceDeposits(NewDeposit(10, 10, 0, 0))
	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(1)

	// WHEN
	// deposit in spread (10 of B at tick 3)
	s.aliceDeposits(NewDeposit(0, 10, 0, 1))
	s.assertAliceBalances(40, 30)
	s.assertDexBalances(10, 20)

	// THEN
	// assert max moved
	s.assertMaxTick(3)
}

func (s *MsgServerTestSuite) TestSingleSidedDepositOutOfSpreadMinNotAdjusted() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// create spread around -1, 1
	s.aliceDeposits(NewDeposit(10, 10, 0, 0))
	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(1)
	// deposit new min
	s.aliceDeposits(NewDeposit(10, 0, 0, 2))
	s.assertAliceBalances(30, 40)
	s.assertDexBalances(20, 10)
	s.assertMinTick(-5)
	s.assertMaxTick(1)

	// WHEN
	// deposit in spread (10 of B at tick 3)
	s.aliceDeposits(NewDeposit(10, 0, 0, 1))
	s.assertAliceBalances(20, 40)
	s.assertDexBalances(30, 10)

	// THEN
	// assert min not moved
	s.assertMinTick(-5)
}

func (s *MsgServerTestSuite) TestSingleSidedDepositOutOfSpreadMaxNotAdjusted() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// create spread around -1, 1
	s.aliceDeposits(NewDeposit(10, 10, 0, 0))
	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(1)
	// deposit new max
	s.aliceDeposits(NewDeposit(0, 10, 0, 2))
	s.assertAliceBalances(40, 30)
	s.assertDexBalances(10, 20)
	s.assertMinTick(-1)
	s.assertMaxTick(5)

	// WHEN
	// deposit in spread (10 of B at tick 3)
	s.aliceDeposits(NewDeposit(0, 10, 0, 1))
	s.assertAliceBalances(40, 20)
	s.assertDexBalances(10, 30)

	// THEN
	// assert max not moved
	s.assertMaxTick(5)
}

func (s *MsgServerTestSuite) TestSingleSidedDepositExistingLiquidityA() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// deposit 10 of token B at tick 1 fee 0
	s.aliceDeposits(NewDeposit(10, 0, 0, 0))
	s.assertAliceBalances(40, 50)
	s.assertDexBalances(10, 0)
	s.assertLiquidityAtTick(10, 0, 0, 0)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(math.MinInt64)
	s.assertMinTick(-1)
	s.assertMaxTick(math.MinInt64)

	// WHEN
	// deposit in spread (10 of B at tick 3)
	s.aliceDeposits(NewDeposit(0, 10, 0, 0))

	// THEN
	// assert 20 of token B deposited at tick 1 fee 0 and ticks unchanged
	s.assertAliceBalances(40, 50)
	s.assertDexBalances(10, 0)
	s.assertLiquidityAtTick(10, 0, 0, 0)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(math.MinInt64)
	s.assertMinTick(-1)
	s.assertMaxTick(math.MinInt64)
	s.Require().Fail("TODO: this fails because PairInit doesn't account for single sided liquidity")
}

func (s *MsgServerTestSuite) TestSingleSidedDepositExistingLiquidityB() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// deposit 10 of token B at tick 1 fee 0
	s.aliceDeposits(NewDeposit(0, 10, 0, 0))
	s.assertAliceBalances(50, 40)
	s.assertDexBalances(0, 10)
	s.assertLiquidityAtTick(0, 10, 0, 0)
	s.assertCurr1To0(math.MaxInt64)
	s.assertCurr0To1(1)
	s.assertMinTick(math.MaxInt64)
	s.assertMaxTick(1)

	// WHEN
	// deposit in spread (10 of B at tick 3)
	s.aliceDeposits(NewDeposit(0, 10, 0, 0))

	// THEN
	// assert 20 of token B deposited at tick 1 fee 0 and ticks unchanged
	s.assertLiquidityAtTick(0, 20, 0, 0)
	s.assertAliceBalances(50, 30)
	s.assertDexBalances(0, 20)
	s.assertCurr1To0(math.MaxInt64)
	s.assertCurr0To1(1)
	s.assertMinTick(math.MaxInt64)
	s.assertMaxTick(1)
	s.Require().Fail("TODO: this fails because PairInit doesn't account for single sided liquidity")
}

func (s *MsgServerTestSuite) TestSingleSidedDepositBelowEnemyLines() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// deposit 10 of token B at tick 0 fee 1
	s.aliceDeposits(NewDeposit(0, 10, 0, 0))
	s.assertAliceBalances(50, 40)
	s.assertDexBalances(0, 10)
	s.assertLiquidityAtTick(0, 10, 0, 0)

	// WHEN
	// depositing above enemy lines at tick -5
	// THEN
	// deposit should fail with BEL error, balances and liquidity should not change at deposited tick

	err := types.ErrValidPairNotFound // TODO: this needs to be changed to a more specific error type
	s.assertAliceDepositFails(err, NewDeposit(0, 10, -5, 0))

	s.assertAliceBalances(50, 40)
	s.assertDexBalances(0, 10)
	s.assertNoLiquidityAtTick(-5, 0)
}

func (s *MsgServerTestSuite) TestSingleSidedDepositAboveEnemyLines() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// deposit 10 of token A at tick 0 fee -1
	s.aliceDeposits(NewDeposit(10, 0, 0, 0))
	s.assertAliceBalances(40, 50)
	s.assertDexBalances(10, 0)
	s.assertLiquidityAtTick(10, 0, 0, 0)

	// WHEN
	// depositing above enemy lines at tick -5
	// THEN
	// deposit should fail with BEL error, balances and liquidity should not change at deposited tick

	err := types.ErrValidPairNotFound // TODO: this needs to be changed to a more specific error type
	s.assertAliceDepositFails(err, NewDeposit(10, 0, 5, 0))

	s.assertAliceBalances(40, 50)
	s.assertDexBalances(10, 0)
	s.assertNoLiquidityAtTick(-5, 0)
}

func (s *MsgServerTestSuite) TestSingleSidedDepositMultiA() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// deposit 10 of token B at tick 1 fee 0
	s.aliceDeposits(NewDeposit(10, 0, 0, 0))
	s.assertAliceBalances(40, 50)
	s.assertDexBalances(10, 0)
	s.assertLiquidityAtTick(10, 0, 0, 0)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(math.MinInt64)
	s.assertMinTick(-1)
	s.assertMaxTick(math.MinInt64)

	// WHEN
	// multi deposit at
	s.aliceDeposits(
		NewDeposit(10, 0, 0, 0),
		NewDeposit(10, 0, 0, 1),
	)

	// THEN
	// assert 20 of token B deposited at tick 1 fee 0 and ticks unchanged
	s.assertAliceBalances(30, 50)
	s.assertDexBalances(20, 0)
	s.assertLiquidityAtTick(20, 0, 0, 0)
	s.assertLiquidityAtTick(10, 0, 0, 1)
	s.assertCurr1To0(-3)
	s.assertCurr0To1(math.MinInt64)
	s.assertMinTick(-3)
	s.assertMaxTick(math.MinInt64)
	s.Require().Fail("TODO: this fails because PairInit doesn't account for single sided liquidity")
}

func (s *MsgServerTestSuite) TestSingleSidedDepositMultiB() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// deposit 10 of token B at tick 1 fee 0
	s.aliceDeposits(NewDeposit(0, 10, 0, 0))
	s.assertAliceBalances(50, 40)
	s.assertDexBalances(0, 10)
	s.assertLiquidityAtTick(0, 10, 0, 0)
	s.assertCurr1To0(math.MaxInt64)
	s.assertCurr0To1(1)
	s.assertMinTick(math.MaxInt64)
	s.assertMaxTick(1)

	// WHEN
	// multi deposit at
	s.aliceDeposits(
		NewDeposit(0, 10, 0, 0),
		NewDeposit(0, 10, 0, 1),
	)

	// THEN
	// assert 20 of token B deposited at tick 1 fee 0 and ticks unchanged
	s.assertAliceBalances(50, 30)
	s.assertDexBalances(0, 20)
	s.assertLiquidityAtTick(0, 20, 0, 0)
	s.assertLiquidityAtTick(0, 10, 0, 1)
	s.assertCurr1To0(math.MaxInt64)
	s.assertCurr0To1(3)
	s.assertMinTick(math.MaxInt64)
	s.assertMaxTick(3)
	s.Require().Fail("TODO: this fails because PairInit doesn't account for single sided liquidity")
}
