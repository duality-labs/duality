package keeper_test

import (
	"math"

	"github.com/NicholasDotSol/duality/x/dex/keeper"
	. "github.com/NicholasDotSol/duality/x/dex/keeper/internal/testutils"
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ** CORE TESTS **
func (s *MsgServerTestSuite) TestSingleWithdrawlFull() {
	s.fundAliceBalances(100, 50)

	// IF Alice deposits 100TokenA & 100TokenB into tick 0 @ feeTier 0
	s.aliceDeposits(NewDeposit(100, 50, 0, 0))

	// WHEN Alice withdraws her entire position (150 Shares)
	err := s.aliceWithdraws(NewWithdrawl(150, 0, 0))
	s.Assert().NoError(err)

	// THEN assert alice gets 100TokenA & 150TokenB; Dex retains no balance; No active Ticks
	s.assertAliceShares(0, 0, sdk.NewDec(0))
	s.assertAliceBalances(100, 50)
	s.assertDexBalances(0, 0)
}

func (s *MsgServerTestSuite) TestSingleWithdrawlPartial() {
	s.fundAliceBalances(100, 50)

	// IF Alice deposits 100 TokenA into tick 0 @ feeTier 0
	s.aliceDeposits(NewDeposit(100, 50, 0, 0))

	// WHEN Alice withdraws her half her position (75 Shares)
	err := s.aliceWithdraws(NewWithdrawl(75, 0, 0))
	s.Assert().NoError(err)

	// THEN Alice gets half her deposit back; Dex retains half
	s.assertAliceShares(0, 0, sdk.NewDec(75))
	s.assertAliceBalances(50, 25)
	s.assertDexBalances(50, 25)
	// CurrentTick values remain unchanged
	s.assertCurrentTicks(-1, 1)
}

func (s *MsgServerTestSuite) TestSingleWithdrawlMaxFee() {
	s.fundAliceBalances(100, 0)

	// IF Alice deposits 100 TokenA into tick 0 @ feeTier 3
	s.aliceDeposits(NewDeposit(100, 0, 0, 3))

	// WHEN Alice withdraws her half her position
	s.aliceWithdraws(NewWithdrawl(50, 0, 3))

	// THEN Alice gets 50 TokenA back and Dex retains balance of 50
	s.assertAliceShares(0, 3, sdk.NewDec(50))
	s.assertAliceBalances(50, 0)
	s.assertDexBalances(50, 0)
}

func (s *MsgServerTestSuite) TestSingleWithdrawlSingleSide() {
	s.fundAliceBalances(100, 0)

	// IF Alice deposits 100 TokenA into tick 0 @ feeTier 0
	s.aliceDeposits(NewDeposit(100, 0, 0, 0))

	// WHEN Alice withdraws her entire position
	s.aliceWithdraws(NewWithdrawl(100, 0, 0))

	// THEN Alice gets 100TokenA back; Dex retains no balance; No Active Ticks
	s.assertAliceShares(0, 0, sdk.NewDec(0))
	s.assertAliceBalances(100, 0)
	s.assertDexBalances(0, 0)

}

func (s *MsgServerTestSuite) TestSingleWithdrawlShiftsTickRight() {
	s.fundAliceBalances(200, 0)

	// IF Alice deposits 100TokenA into ticks 1 & 2
	s.aliceDeposits(
		NewDeposit(100, 0, 2, 0),
		NewDeposit(100, 0, 1, 0),
	)
	// currentTick1To0 = 2
	s.assertCurrentTicks(1, math.MaxInt64)

	// WHEN Alice withdraws her shares from Tick 1
	s.aliceWithdraws(NewWithdrawl(100, 2, 0))

	// THEN currentTick1To0 = 3
	//TODO: this is currently failling because of TickCount bug
	s.assertCurrentTicks(0, math.MaxInt64)
}

func (s *MsgServerTestSuite) TestSingleWithdrawlShiftsTickLeft() {
	s.fundAliceBalances(0, 200)

	// IF Alice deposits 100TokenA into ticks 1 & 2
	s.aliceDeposits(
		NewDeposit(0, 100, 1, 0),
		NewDeposit(0, 100, 2, 0),
	)
	// currentTick0To1 = 1
	s.assertCurrentTicks(math.MinInt64, 2)

	// WHEN Alice withdraws her shares from Tick 1
	sharesToWithdraw := keeper.CalcShares(
		NewDec(0),
		NewDec(100),
		keeper.CalcPrice1To0(1),
	)
	s.aliceWithdraws(NewWithdrawlDec(sharesToWithdraw, 1, 0))

	// THEN currentTick0To1 = 0
	s.assertCurrentTicks(math.MinInt64, 3)
}

func (s *MsgServerTestSuite) TestMultiWithdrawlShiftsTickLeft() {
	s.fundAliceBalances(300, 0)

	// IF Alice deposits 100TokenA into ticks 1, 2 & 3
	s.aliceDeposits(
		NewDeposit(100, 0, 3, 0),
		NewDeposit(100, 0, 2, 0),
		NewDeposit(100, 0, 1, 0),
	)
	// currentTick1To0 = 2
	s.assertCurrentTicks(2, math.MaxInt64)

	// WHEN Alice withdraws all her shares from Tick 2 & 3
	s.aliceWithdraws(
		NewWithdrawl(100, 2, 0),
		NewWithdrawl(100, 3, 0))

	s.assertCurrentTicks(0, math.MaxInt64)
}

// TODO: more test to write

// TestMultiWithdrawlShiftsTickRight

//TestSingleWithdrawTiersShiftsTickRight
//TestSingleWithdrawTiersShiftsTickLeft
//TestMultiWithdrawTiersShiftsTickRight
//TestMultiWithdrawTiersShiftsTickLeft

// d 100 t1 d 100 d 100 t3 t2 [w100 t3 w 50 t2] 100 t2 => currentTick1To0 = 3

// d 100 t1 d 100 d 100 t3 t2 [w100 t2 w100 t3] 100 t2 => currentTick1To0 =

// ** EDGE CASE FAILURE TESTS **

func (s *MsgServerTestSuite) TestWithdrawalFailsWhenNotEnoughShares() {
	s.fundAliceBalances(100, 0)

	// IF  Alice deposits 100
	s.aliceDeposits(NewDeposit(100, 0, 0, 0))

	// WHEN Alice tries to withdraw 200
	err := s.aliceWithdraws(NewWithdrawl(200, 0, 0))

	// THEN ensure error is thrown and Alice and Dex balances remain unchanged
	s.Assert().ErrorIs(err, types.ErrNotEnoughShares)
	s.assertAliceShares(0, 0, sdk.NewDec(100))
	s.assertAliceBalances(0, 0)
}

func (s *MsgServerTestSuite) TestWithdrawalFailsWhenNotEnoughSharesMulti() {
	s.fundAliceBalances(100, 0)

	// IF Alice Deposists 100
	s.aliceDeposits(NewDeposit(100, 0, 0, 0))

	// WHEN Alice does multiple withdrawals > 100
	err := s.aliceWithdraws(
		NewWithdrawl(50, 0, 0),
		NewWithdrawl(50, 0, 0),
		NewWithdrawl(50, 0, 0),
	)

	// THEN an error is thrown and Alice and Dex balances remain unchanged
	s.Require().ErrorIs(err, types.ErrNotEnoughShares)
}

func (s *MsgServerTestSuite) TestWithdrawalFailsWithNonExistentPair() {
	s.fundAliceBalances(100, 0)

	// IF Alice Deposists 100
	s.aliceDeposits(NewDeposit(100, 0, 0, 0))

	// WHEN Alice tries to withdraw from a nonexistent tokenPair
	_, err := s.msgServer.Withdrawl(s.goCtx, &types.MsgWithdrawl{
		Creator:        s.alice.String(),
		Receiver:       s.alice.String(),
		TokenA:         "TokenX",
		TokenB:         "TokenZ",
		SharesToRemove: []sdk.Dec{sdk.NewDec(10)},
		TickIndexes:    []int64{0},
		FeeIndexes:     []uint64{0},
	})

	// NOTE: As code is currently written we hit not enough shares check
	// before validating pair existence. This is correct from a
	// UX perspective --users should not care whether tick is initialized
	s.Assert().ErrorIs(err, types.ErrNotEnoughShares)
}

func (s *MsgServerTestSuite) TestWithdrawalFailsWithInvalidTick() {
	s.fundAliceBalances(100, 0)

	// IF Alice Deposists 100
	s.aliceDeposits(NewDeposit(100, 0, 0, 0))

	// WHEN Alice tries to withdraw from an invalid tick
	err := s.aliceWithdraws(NewWithdrawl(50, 10, 0))

	// NOTE: See above NOTE on error condition from TestFailsWithNonExistentPair
	s.Assert().ErrorIs(err, types.ErrNotEnoughShares)
	s.assertAliceShares(0, 0, sdk.NewDec(100))
	s.assertDexBalances(100, 0)
}

func (s *MsgServerTestSuite) TestWithdrawalFailsWithInvalidTickMulti() {
	s.fundAliceBalances(100, 0)

	// IF Alice Deposists 100
	s.aliceDeposits(NewDeposit(100, 0, 0, 0))

	// WHEN Alice tries to withdraw from a mix of valid and invalid ticks
	err := s.aliceWithdraws(
		// INVALID
		NewWithdrawl(50, 10, 0),
		// VALID
		NewWithdrawl(50, 0, 0),
	)

	// NOTE: See above NOTE on error condition from TestFailsWithNonExistentPair
	s.Assert().ErrorIs(err, types.ErrNotEnoughShares)
	s.assertAliceShares(0, 0, sdk.NewDec(100))
	s.assertDexBalances(100, 0)
}

func (s *MsgServerTestSuite) TestWithdrawalFailsWithInvalidFee() {
	s.fundAliceBalances(100, 0)

	// IF Alice Deposists 100
	s.aliceDeposits(NewDeposit(100, 0, 0, 0))

	// WHEN Alice tries to withdraw from an invalid tick
	err := s.aliceWithdraws(NewWithdrawl(100, 0, 99))

	s.Assert().ErrorIs(err, types.ErrValidFeeIndexNotFound)
	s.assertAliceShares(0, 0, sdk.NewDec(100))
	s.assertDexBalances(100, 0)
}

func (s *MsgServerTestSuite) TestWithdrawalFailsWithInvalidFeeMulti() {
	s.fundAliceBalances(100, 0)

	// IF Alice Deposists 100
	s.aliceDeposits(NewDeposit(100, 0, 0, 0))

	// WHEN Alice tries to withdraw from a mix of valid and invalid FeeTiers
	err := s.aliceWithdraws(
		// INVALID
		NewWithdrawl(50, 0, 10),
		// VALID
		NewWithdrawl(50, 0, 0),
	)

	// THEN ensure error is thrown and Alice and Dex balances remain unchanged
	s.Assert().ErrorIs(err, types.ErrValidFeeIndexNotFound)
	s.assertAliceShares(0, 0, sdk.NewDec(100))
	s.assertDexBalances(100, 0)
}
