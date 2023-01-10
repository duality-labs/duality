package keeper_test

import (
	"math"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ** CORE TESTS **

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

// ** EDGE CASE FAILURE TESTS **

func (s *MsgServerTestSuite) TestWithdrawalFailsWhenNotEnoughShares() {
	s.fundAliceBalances(100, 0)

	// IF  Alice deposits 100
	s.aliceDeposits(NewDeposit(100, 0, 0, 0))

	// WHEN Alice tries to withdraw 200
	err := s.aliceWithdraws(NewWithdrawl(200, 0, 0))

	// THEN ensure error is thrown and Alice and Dex balances remain unchanged
	s.Assert().ErrorIs(err, types.ErrNotEnoughShares)
	s.assertAliceShares(0, 0, 100)
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
		SharesToRemove: []sdk.Int{sdk.NewInt(10)},
		TickIndexes:    []int64{0},
		FeeIndexes:     []uint64{0},
	})

	// NOTE: As code is currently written we hit not enough shares check
	// before validating pair existence. This is correct from a
	// UX perspective --users should not care whether tick is initialized
	s.Assert().ErrorIs(err, types.ErrValidPairNotFound)
}

func (s *MsgServerTestSuite) TestWithdrawalFailsWithInvalidTick() {
	s.fundAliceBalances(100, 0)

	// IF Alice Deposists 100
	s.aliceDeposits(NewDeposit(100, 0, 0, 0))

	// WHEN Alice tries to withdraw from an invalid tick
	err := s.aliceWithdraws(NewWithdrawl(50, 10, 0))

	// NOTE: See above NOTE on error condition from TestFailsWithNonExistentPair
	s.Assert().ErrorIs(err, types.ErrNotEnoughShares)
	s.assertAliceShares(0, 0, 100)
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
	s.assertAliceShares(0, 0, 100)
	s.assertDexBalances(100, 0)
}

func (s *MsgServerTestSuite) TestWithdrawalFailsWithInvalidFee() {
	s.fundAliceBalances(100, 0)

	// IF Alice Deposists 100
	s.aliceDeposits(NewDeposit(100, 0, 0, 0))

	// WHEN Alice tries to withdraw from an invalid tick
	err := s.aliceWithdraws(NewWithdrawl(100, 0, 99))

	s.Assert().ErrorIs(err, types.ErrValidFeeIndexNotFound)
	s.assertAliceShares(0, 0, 100)
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
	s.assertAliceShares(0, 0, 100)
	s.assertDexBalances(100, 0)
}
