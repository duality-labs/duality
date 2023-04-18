package keeper_test

import (
	"github.com/duality-labs/duality/x/dex/types"
)

func (s *MsgServerTestSuite) TestSwapNoLONoLiquidity() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 50)
	// GIVEN
	// no liqudity of token A (deposit only token B at tick 0 fee 1)
	s.aliceDeposits(NewDeposit(0, 10, 0, 1))
	s.assertAliceBalances(50, 40)

	// WHEN
	// swap 5 of tokenB
	// THEN
	// swap should fail with Error Not enough coins
	err := types.ErrInsufficientLiquidity
	s.bobMarketSellFails(err, "TokenB", 5)
}

func (s *MsgServerTestSuite) TestSwapNoLOPartiallyFilledSlippageToleranceNotReachedMaxReached() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// deposit 10 of tokenB
	s.aliceDeposits(NewDeposit(0, 10, 0, 1))
	s.assertAliceBalances(50, 40)
	s.assertDexBalances(0, 10)
	s.assertPoolLiquidity(0, 10, 0, 1)
	//
	// WHEN
	// swap 20 of tokenA at
	amountIn := 20
	s.bobMarketSells("TokenA", amountIn)

	// THEN
	// swap should have in 11 out 10
	s.assertBobBalances(39, 10)
	s.assertDexBalances(11, 0)
}

func (s *MsgServerTestSuite) TestSwapNoLOPartiallyFilledSlippageToleranceNotReachedMinReached() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(0, 50)
	// GIVEN
	// deposit 10 of token A at tick 0 fee 1
	s.aliceDeposits(NewDeposit(10, 0, 0, 1))
	s.assertAliceBalances(40, 50)
	s.assertDexBalances(10, 0)
	s.assertPoolLiquidity(10, 0, 0, 1)
	//
	// WHEN
	// swap 20 of token A for B
	amountIn := 20
	s.bobMarketSells("TokenB", amountIn)

	// THEN
	// swap should have in 11 out 10
	s.assertBobBalances(10, 39)
	s.assertDexBalances(0, 11)
}

func (s *MsgServerTestSuite) TestSwapNoLOCorrectExecutionMinFee() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// deposit 10 of token B at tick 0 fee 0
	s.aliceDeposits(NewDeposit(0, 10, 0, 1))
	s.assertAliceBalances(50, 40)
	s.assertDexBalances(0, 10)

	// WHEN
	// swap 5 of token A for B
	s.bobMarketSells("TokenA", 5)

	// THEN
	// swap should have in 5 out 4
	s.assertBobBalances(45, 4)
	s.assertDexBalances(5, 6)
}

func (s *MsgServerTestSuite) TestSwapNoLOCorrectExecutionHighFee() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// deposit 10 of token B at tick 0 fee 10
	s.aliceDeposits(NewDeposit(0, 10, 0, 10))
	s.assertAliceBalances(50, 40)
	s.assertDexBalances(0, 10)

	// WHEN
	// swap 5 of token A for B
	s.bobMarketSells("TokenA", 5)

	// THEN
	// swap should have in 5 out 4
	s.assertBobBalances(45, 4)
	s.assertDexBalances(5, 6)
}

func (s *MsgServerTestSuite) TestSwapNoLOCorrectExecutionSomeFeeTiers() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// deposit 10 of token B at tick 0 fee 1 and 10 of token B at tick 0 fee 3
	s.aliceDeposits(
		NewDeposit(0, 10, 0, 1),
		NewDeposit(0, 10, 0, 3),
	)
	s.assertAliceBalances(50, 30)
	s.assertDexBalances(0, 20)

	// WHEN
	// swap 15 of token A for B
	s.bobMarketSells("TokenA", 15)

	// THEN
	// swap should have in 15 out 13
	s.assertBobBalances(35, 13)
	s.assertDexBalances(15, 7)
}

func (s *MsgServerTestSuite) TestSwapNoLO1to0DoesntMoveCurr1to0() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(0, 50)
	// GIVEN
	// deposit 10 of token A at tick 0 fee 1
	s.aliceDeposits(NewDeposit(10, 0, 0, 1))
	s.assertCurr1To0(-1)

	// WHEN
	// swap 5 of token B for A
	s.bobMarketSells("TokenB", 5)

	// THEN
	// current1To0 unchanged
	s.assertCurr1To0(-1)
}

func (s *MsgServerTestSuite) TestSwapNoLO1to0MovesCurr1to0() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(0, 50)
	// GIVEN
	// deposit 10 of token A at tick 0 fee 1, 10 of token A at tick 0 fee 3
	s.aliceDeposits(
		NewDeposit(10, 0, 0, 1),
		NewDeposit(10, 0, 0, 3),
	)
	s.assertCurr1To0(-1)

	// WHEN
	// swap 15 of token B for A
	s.bobMarketSells("TokenB", 15)

	// THEN
	// current 1to0 moves to -3
	s.assertCurr1To0(-3)
}

func (s *MsgServerTestSuite) TestSwapNoLO1to0DoesntMoveCurr0to1() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(0, 50)
	// GIVEN
	// deposit 10 of both token A and B at tick 0 fee 1
	s.aliceDeposits(NewDeposit(10, 10, 0, 1))
	s.assertCurr0To1(1)

	// WHEN
	// swap 5 of token B for A
	s.bobMarketSells("TokenB", 5)

	// THEN
	// current 0to1 unchanged
	s.assertCurr0To1(1)
}

func (s *MsgServerTestSuite) TestSwapNoLO1to0MovesCurr0to1() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(0, 50)
	// GIVEN
	// deposit 10 of token A at tick 0 fee 1 and 10 of both token A and B at tick 0 fee 3
	s.aliceDeposits(
		NewDeposit(10, 0, 0, 1),
		NewDeposit(10, 10, 0, 3),
	)
	s.assertCurr0To1(3)

	// WHEN
	// swap 5 of token B for A
	s.bobMarketSells("TokenB", 5)

	// THEN
	// current 0to1 moves down to 1
	s.assertCurr0To1(1)
}

func (s *MsgServerTestSuite) TestSwapNoLO0to1DoesntMoveCurr0to1() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// deposit 10 of token B at tick 0 fee 1
	s.aliceDeposits(
		NewDeposit(0, 10, 0, 1),
	)
	s.assertCurr0To1(1)

	// WHEN
	// swap 5 of token A for B
	s.bobMarketSells("TokenA", 5)

	// THEN
	// current0To1 unchanged
	s.assertCurr0To1(1)
}

func (s *MsgServerTestSuite) TestSwapNoLO0to1MovesCurr0to1() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// deposit 10 of token B at tick 0 fee 1, 10 of token B at tick 0 fee 3
	s.aliceDeposits(
		NewDeposit(0, 10, 0, 1),
		NewDeposit(0, 10, 0, 3),
	)
	s.assertCurr0To1(1)

	// WHEN
	// swap 15 of token A for B
	s.bobMarketSells("TokenA", 15)

	// THEN
	// current 0to1 moves to 3
	s.assertCurr0To1(3)
}

func (s *MsgServerTestSuite) TestSwapNoLO0to1DoesntMoveCurr1to0() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// deposit 10 of both token A and B at tick 0 fee 1
	s.aliceDeposits(NewDeposit(10, 10, 0, 1))
	s.assertCurr1To0(-1)

	// WHEN
	// swap 5 of token A for B
	s.bobMarketSells("TokenA", 5)

	// THEN
	// current 1to0 unchanged
	s.assertCurr1To0(-1)
}

func (s *MsgServerTestSuite) TestSwapNoLO0to1MovesCurr1to0() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// deposit 10 of token B at tick 0 fee 1 and 10 of both token A and B at tick 0 fee 3
	// to create spread of -3, 1
	s.aliceDeposits(
		NewDeposit(0, 10, 0, 1),
		NewDeposit(10, 10, 0, 3),
	)
	s.assertCurr1To0(-3)

	// WHEN
	// swap 5 of token B for A
	s.bobMarketSells("TokenA", 5)

	// THEN
	// current 0to1 moves down to 1
	s.assertCurr1To0(-1)
}

func (s *MsgServerTestSuite) TestSwapNoLO0to1DoesntMoveMax() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// deposit 10 of token B at tick 0 fee 1 and 10 of both token A and B at tick 0 fee 3
	// to create spread of -3, 1
	s.aliceDeposits(
		NewDeposit(0, 10, 0, 1),
		NewDeposit(10, 10, 0, 3),
	)
	s.assertCurr1To0(-3)

	// WHEN
	// swap 5 of token B for A
	s.bobMarketSells("TokenA", 5)

	// THEN
	// current 0to1 moves down to 1
	s.assertCurr1To0(-1)
}

func (s *MsgServerTestSuite) TestSwapNoLOMaxAmountOutUsed() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// 10 TokenB available
	s.aliceDeposits(
		NewDeposit(0, 10, 0, 1),
	)

	// WHEN
	// swap 50 with maxOut of 5
	s.bobMarketSellsWithMaxOut("TokenA", 50, 5)

	// THEN
	// bob gets 5 out
	s.assertBobBalances(44, 5)
}

func (s *MsgServerTestSuite) TestSwapNoLOMaxAmountNotOutUsed() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// 10 TokenB available
	s.aliceDeposits(
		NewDeposit(0, 10, 0, 1),
	)

	// WHEN
	// swap 8 with maxOut of 15
	s.bobMarketSellsWithMaxOut("TokenA", 8, 15)

	// THEN
	// bob gets 7 out
	s.assertBobBalances(42, 7)
}
