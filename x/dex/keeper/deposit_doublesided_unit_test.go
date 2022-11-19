package keeper_test

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
)

func (s *MsgServerTestSuite) TestDepositDoubleSidedInSpreadCurrTickAdjusted() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// create spread around -5, 5
	s.aliceDeposits(NewDeposit(10, 10, 0, 2))
	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertCurr1To0(-5)
	s.assertCurr0To1(5)

	// WHEN
	// deposit in spread (10 of A,B at tick 0 fee 1)
	s.aliceDeposits(NewDeposit(10, 10, 0, 0))
	s.assertAliceBalances(30, 30)
	s.assertDexBalances(20, 20)

	// THEN
	// assert currentTick1To0, currTick0to1 moved
	s.assertCurr1To0(-1)
	s.assertCurr0To1(1)
}

func (s *MsgServerTestSuite) TestDepositDoubleSidedInSpreadMinMax() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// create spread around -5, 5
	s.aliceDeposits(NewDeposit(10, 10, 0, 2))
	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertMinTick(-5)
	s.assertMaxTick(5)

	// WHEN
	// deposit in spread (10 of A,B at tick 0 fee 1)
	s.aliceDeposits(NewDeposit(10, 10, 0, 0))
	s.assertAliceBalances(30, 30)
	s.assertDexBalances(20, 20)

	// THEN
	// assert Min/Max unchanged
	s.assertMinTick(-5)
	s.assertMaxTick(5)
}

func (s *MsgServerTestSuite) TestDepositDoubleSidedAroundSpreadCurrTickNotAdjusted() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// create spread around -1, 1
	s.aliceDeposits(NewDeposit(10, 10, 0, 0))
	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(1)

	// WHEN
	// deposit around spread (10 of A,B at tick 0 fee 3)
	s.aliceDeposits(NewDeposit(10, 10, 0, 1))
	s.assertAliceBalances(30, 30)
	s.assertDexBalances(20, 20)

	// THEN
	// assert CurrentTick0To1, CurrentTick1To0 unchanged
	s.assertCurr1To0(-1)
	s.assertCurr0To1(1)
}

func (s *MsgServerTestSuite) TestDepositDoubleSidedAroundSpreadMinMaxNotAdjusted() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// create spread around -1, 1
	s.aliceDeposits(NewDeposit(10, 10, 0, 0))
	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertMinTick(-1)
	s.assertMaxTick(1)

	// WHEN
	// deposit around spread (10 of A,B at tick 0 fee 5)
	s.aliceDeposits(NewDeposit(10, 10, 0, 2))
	s.assertAliceBalances(30, 30)
	s.assertDexBalances(20, 20)
	s.assertMinTick(-5)
	s.assertMaxTick(5)

	// deposit in new spread (10 of A,B at tick 0 fee 3)
	s.aliceDeposits(NewDeposit(10, 10, 0, 1))
	s.assertAliceBalances(20, 20)
	s.assertDexBalances(30, 30)

	// THEN
	// assert Min/Max unchanged
	s.assertMinTick(-5)
	s.assertMaxTick(5)
}

func (s *MsgServerTestSuite) TestDepositDoubleSidedAroundSpreadMinMaxAdjusted() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// create spread around -1, 1
	s.aliceDeposits(NewDeposit(10, 10, 0, 0))
	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertMinTick(-1)
	s.assertMaxTick(1)

	// WHEN
	// deposit around spread (10 of A,B at tick 0 fee 5)
	s.aliceDeposits(NewDeposit(10, 10, 0, 2))
	s.assertAliceBalances(30, 30)
	s.assertDexBalances(20, 20)
	s.assertMinTick(-5)
	s.assertMaxTick(5)

	// deposit in new spread (10 of A,B at tick 0 fee 10)
	s.aliceDeposits(NewDeposit(10, 10, 0, 3))
	s.assertAliceBalances(20, 20)
	s.assertDexBalances(30, 30)

	// THEN
	// assert Min/Max adjusted
	s.assertMinTick(-10)
	s.assertMaxTick(10)
}

func (s *MsgServerTestSuite) TestDepositDoubleSidedHalfInSpreadCurrTick0To1Adjusted() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// create spread around -5, 5
	s.aliceDeposits(NewDeposit(10, 10, 0, 2))
	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertCurr1To0(-5)
	s.assertCurr0To1(5)

	// WHEN
	// deposit half in spread (10 of A,B at tick 1 fee 5)
	s.aliceDeposits(NewDeposit(10, 10, 1, 2))
	s.assertAliceBalances(30, 30)
	s.assertDexBalances(20, 20)

	// THEN
	// assert CurrTick1to0 unchanged, CurrTick0to1 adjusted
	s.assertCurr1To0(-4)
	s.assertCurr0To1(5)
}

func (s *MsgServerTestSuite) TestDepositDoubleSidedHalfInSpreadCurrTick1To0Adjusted() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// create spread around -5, 5
	s.aliceDeposits(NewDeposit(10, 10, 0, 2))
	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertCurr1To0(-5)
	s.assertCurr0To1(5)

	// WHEN
	// deposit half in spread (10 of A,B at tick -1 fee 5)
	s.aliceDeposits(NewDeposit(10, 10, -1, 2))
	s.assertAliceBalances(30, 30)
	s.assertDexBalances(20, 20)

	// THEN
	// assert CurrTick0to1 unchanged, CurrTick1to0 adjusted
	s.assertCurr1To0(-5)
	s.assertCurr0To1(4)
}

func (s *MsgServerTestSuite) TestDepositDoubleSidedHalfInSpreadMinAdjusted() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// create spread around -5, 5
	s.aliceDeposits(NewDeposit(10, 10, 0, 2))
	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertMinTick(-5)
	s.assertMaxTick(5)

	// WHEN
	// deposit half in spread (10 of A,B at tick -1 fee 5)
	s.aliceDeposits(NewDeposit(10, 10, -1, 2))
	s.assertAliceBalances(30, 30)
	s.assertDexBalances(20, 20)

	// THEN
	// assert Min adjusted, Max unchanged
	s.assertMinTick(-6)
	s.assertMaxTick(5)
}

func (s *MsgServerTestSuite) TestDepositDoubleSidedHalfInSpreadMaxAdjusted() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// create spread around -5, 5
	s.aliceDeposits(NewDeposit(10, 10, 0, 2))
	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertMinTick(-5)
	s.assertMaxTick(5)

	// WHEN
	// deposit half in spread (10 of A,B at tick 1 fee 5)
	s.aliceDeposits(NewDeposit(10, 10, 1, 2))
	s.assertAliceBalances(30, 30)
	s.assertDexBalances(20, 20)

	// THEN
	// assert Max adjusted, Min unchanged
	s.assertMinTick(-5)
	s.assertMaxTick(6)
}

func (s *MsgServerTestSuite) TestDepositDoubleSidedBelowEnemyLines() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// deposit 10 of token A at tick 0 fee 1
	s.aliceDeposits(NewDeposit(10, 0, 0, 0))
	s.assertAliceBalances(40, 50)
	s.assertDexBalances(10, 0)
	s.assertLiquidityAtTick(10, 0, 0, 0)

	// WHEN
	// depositing below enemy lines at tick -5
	// THEN
	// deposit should fail with BEL error, balances and liquidity should not change at deposited tick

	err := types.ErrDepositBehindPairLiquidity
	s.assertAliceDepositFails(err, NewDeposit(10, 10, -5, 0))
}

func (s *MsgServerTestSuite) TestDepositDoubleSidedAboveEnemyLines() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// deposit 10 of token A at tick 0 fee 1
	s.aliceDeposits(NewDeposit(0, 10, 0, 0))
	s.assertAliceBalances(50, 40)
	s.assertDexBalances(0, 10)
	s.assertLiquidityAtTick(0, 10, 0, 0)

	// WHEN
	// depositing above enemy lines at tick 5
	// THEN
	// deposit should fail with BEL error, balances and liquidity should not change at deposited tick

	err := types.ErrDepositBehindPairLiquidity
	s.assertAliceDepositFails(err, NewDeposit(10, 10, 5, 0))
}
