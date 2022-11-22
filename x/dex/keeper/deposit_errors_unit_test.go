package keeper_test

import "github.com/NicholasDotSol/duality/x/dex/types"

func (s *MsgServerTestSuite) TestDepositErrorBEL1to0() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// token A deposited at tick 0 fee 1
	s.aliceDeposits(NewDeposit(10, 0, 0, 0))

	// WHEN
	// token B deposited at tick -5 fee 1
	// THEN
	// types.ErrDepositBehindPairLiquidity should be returned
	err := types.ErrDepositBehindPairLiquidity
	s.assertAliceDepositFails(err, NewDeposit(0, 10, -5, 0))
}

func (s *MsgServerTestSuite) TestDepositErrorBEL0to1() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// token B deposited at tick 0 fee 1
	s.aliceDeposits(NewDeposit(0, 10, 0, 0))

	// WHEN
	// token A deposited at tick 5 fee 1
	// THEN
	// types.ErrDepositBehindPairLiquidity should be returned
	err := types.ErrDepositBehindPairLiquidity
	s.assertAliceDepositFails(err, NewDeposit(10, 0, 5, 0))
}

func (s *MsgServerTestSuite) TestDepositErrorAllDepositsFailed() {
	s.fundAliceBalances(50, 50)
	// GIVEN
	// a single sided pool at tick 0 fee 1 and double sided pool at tick 0 fee 3
	s.aliceDeposits(
		NewDeposit(10, 0, 0, 0),
		NewDeposit(10, 10, 0, 1),
	)

	// WHEN
	// double sided deposit in the single sided pool and single sided deposit at the double sided pool
	// THEN
	// types.ErrAllDepositsFailed should be returned
	err := types.ErrAllDepositsFailed
	s.assertAliceDepositFails(err,
		NewDeposit(10, 10, 0, 0),
		NewDeposit(0, 10, 0, 1),
	)
}

func (s *MsgServerTestSuite) TestDepositErrorBankTransferFailed0() {
	// GIVEN
	// alice has no A funds
	// WHEN
	// depositing a non-zero amount of A
	// THEN
	// types.ErrNotEnoughCoins should be returned
	err := types.ErrNotEnoughCoins
	s.assertAliceDepositFails(err, NewDeposit(10, 0, 0, 0))
}

func (s *MsgServerTestSuite) TestDepositErrorBankTransferFailed1() {
	// GIVEN
	// user has no B funds
	// WHEN
	// depositing a non zero amount of B
	// THEN
	// types.ErrNotEnoughCoins should be returned
	err := types.ErrNotEnoughCoins
	s.assertAliceDepositFails(err, NewDeposit(0, 10, 0, 0))
}
