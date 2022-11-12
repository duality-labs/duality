package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/NicholasDotSol/duality/x/dex/types"
)

// ** CORE TESTS **
func (s *MsgServerTestSuite) TestSingleWithdrawlFull() {
	s.fundAliceBalances(100, 50)

	// IF Alice deposits 100TokenA & 100TokenB into tick 0 @ feeTier 0
	s.aliceDeposits(NewDeposit(100, 50, 0, 0))

	// WHEN Alice withdraws her entire position (150 Shares)
	err := s.aliceWithdraws(newWithdrawl(0, 0, 150))
	s.Assert().NoError(err)

	// THEN assert alice gets 100TokenA & 150TokenB; Dex retains no balance; No active Ticks
	s.assertAliceShares(0, 0, sdk.NewDec(0))
	s.assertAliceBalances(100, 50)
	s.assertDexBalances(0, 0)
	s.assertTickCount(0)
}

func (s *MsgServerTestSuite) TestSingleWithdrawlPartial() {
	s.fundAliceBalances(100, 50)

	// IF Alice deposits 100 TokenA into tick 0 @ feeTier 0
	s.aliceDeposits(NewDeposit(100, 50, 0, 0))

	// WHEN Alice withdraws her half her position (75 Shares)
	err := s.aliceWithdraws(newWithdrawl(0, 0, 75))
	s.Assert().NoError(err)


	// THEN Alice gets half her deposit back; Dex retains half
	s.assertAliceShares(0, 0, sdk.NewDec(75))
	s.assertAliceBalances(50, 25)
	s.assertDexBalances(50, 25)

	// CurrentTick values remain unchanged
	s.assertTickCount(1)
	s.assertCurrentTicks(-1, 1)

}

func (s *MsgServerTestSuite) TestSingleWithdrawlSingleSide() {
	s.fundAliceBalances(100, 0)

	// IF Alice deposits 100 TokenA into tick 0 @ feeTier 0
	s.aliceDeposits(NewDeposit(100, 0, 0, 0))

	// WHEN Alice withdraws her entire position
	s.aliceWithdraws(newWithdrawl(0, 0, 100))

	// THEN Alice gets 100TokenA back; Dex retains no balance; No Active Ticks
	s.assertAliceShares(0, 0, sdk.NewDec(0))
	s.assertAliceBalances(100, 0)
	s.assertDexBalances(0, 0)
	s.assertTickCount(0)

}

func (s *MsgServerTestSuite) TestSingleWithdrawlShiftsTickRight() {
	s.fundAliceBalances(200, 0)

	// IF Alice deposits 100TokenA into ticks 1 & 2
	s.aliceDeposits(
		NewDeposit(100, 0, 2, 0),
		NewDeposit(100, 0, 1, 0),
	)
	// currentTick1To0 = 2
	s.assertCurrentTicks(1, 2)

	// WHEN Alice withdraws her shares from Tick 1
	s.aliceWithdraws(newWithdrawl(1, 0, 100))

	// THEN currentTick1To0 = 3
	//TODO: this is currently failling because of TickCount bug
	s.assertCurrentTicks(1, 3)


}

func (s *MsgServerTestSuite) TestSingleWithdrawlShiftsTickLeft() {
	s.fundAliceBalances(0, 200)

	// IF Alice deposits 100TokenA into ticks 1 & 2
	s.aliceDeposits(
		NewDeposit(0, 100, 1, 0),
		NewDeposit(0, 100, 2, 0),
	)
	// currentTick0To1 = 1
	s.assertCurrentTicks(1, 2)

	// WHEN Alice withdraws her shares from Tick 2
	s.aliceWithdraws(newWithdrawl(2, 0, 100))

	// THEN currentTick1To0 = 1
	//TODO: this is currently failling because of TickCount bug
	s.assertCurrentTicks(0, 2)
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
	s.assertCurrentTicks(1, 2)

	// WHEN Alice withdraws all her shares from Tick 1 & 2
	s.aliceWithdraws(
		newWithdrawl(1, 0, 100),
		newWithdrawl(2, 0, 100))

	// THEN currentTick1To0 = 4
	//TODO: this is currently failling because of TickCount bug
	s.assertCurrentTicks(0, 4)
}




// d 100 t1 d 100 d 100 t3 t2 [w100 t3 w 50 t2] 100 t2 => currentTick t2

// d 100 t1 d 100 d 100 t3 t2 [w100 t2 w100 t3] 100 t2 => currentTick t1


func (s *MsgServerTestSuite) TestSingleWithdrawlMaxFee() {
	s.fundAliceBalances(100, 0)

	// IF Alice deposits 100 TokenA into tick 0 @ feeTier 3
	s.aliceDeposits(NewDeposit(100, 0, 0, 3))

	// WHEN Alice withdraws her half her position
	s.aliceWithdraws(newWithdrawl(0, 3, 50))

	// THEN Alice gets 50 TokenA back and Dex retains balance of 50
	s.assertAliceShares(0, 3, sdk.NewDec(50))
	s.assertAliceBalances(50, 0)
	s.assertDexBalances(50, 0)
}

// ** EDGE CASE FAILURE TESTS **

func (s *MsgServerTestSuite) TestFailsWhenNotEnoughShares() {
	s.fundAliceBalances(100, 0)

	// IF  Alice deposits 100
	s.aliceDeposits(NewDeposit(100, 0, 0, 0))

	// WHEN Alice tries to withdraw 200
	err := s.aliceWithdraws(newWithdrawl(0, 0, 200))

	// THEN ensure error is thrown and Alice and Dex balances remain unchanged
	s.Assert().Error(err)
	s.assertAliceShares(0, 0, sdk.NewDec(100))
	s.assertAliceBalances(0, 0)
}

func (s *MsgServerTestSuite) TestFailsWhenNotEnoughSharesMulti() {
	s.fundAliceBalances(100, 0)

	// IF Alice Deposists 100
	s.aliceDeposits(NewDeposit(100, 0, 0, 0))

	// WHEN Alice does multiple withdrawals > 100
	err := s.aliceWithdraws(
		newWithdrawl(0, 0, 50),
		newWithdrawl(0, 0, 50),
		newWithdrawl(0, 0, 50),
	)


	// THEN an error is thrown and Alice and Dex balances remain unchanged
	s.Assert().Error(err)

	s.assertAliceShares(0, 0, sdk.NewDec(100))
	s.assertDexBalances(100, 0)
}

func (s *MsgServerTestSuite) TestFailsWithNonExistentPair() {
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

	// THEN ensure error is throw
	// TODO: Not sure how to trigger this failure case because we hit the insufficient shares check first
	s.Assert().Error(err)
}

func (s *MsgServerTestSuite) TestFailsWithInvalidTick() {
	s.fundAliceBalances(100, 0)

	// IF Alice Deposists 100
	s.aliceDeposits(NewDeposit(100, 0, 0, 0))

	// WHEN Alice tries to withdraw from an invalid tick
	err := s.aliceWithdraws(newWithdrawl(10, 0, 50))

	// THEN ensure error is thrown and Alice and Dex balances remain unchanged
	// TODO: Not sure how to trigger this failure case because we hit the insufficient shares check first
	s.Assert().Error(err)
	s.assertAliceShares(0, 0, sdk.NewDec(100))
	s.assertDexBalances(100, 0)

}

func (s *MsgServerTestSuite) TestFailsWithInvalidTickMulti() {
	s.fundAliceBalances(100, 0)

	// IF Alice Deposists 100
	s.aliceDeposits(NewDeposit(100, 0, 0, 0))

	// WHEN Alice tries to withdraw from a mix of valid and invalid ticks
	err := s.aliceWithdraws(
		// INVALID
		newWithdrawl(10, 0, 50),
		// VALID
		newWithdrawl(0, 0, 50),
	)

	// THEN ensure error is thrown and Alice and Dex balances remain unchanged
	// TODO: Not sure how to trigger this failure case because we hit the insufficient shares check first
	s.Assert().Error(err)
	s.assertAliceShares(0, 0, sdk.NewDec(100))
	s.assertDexBalances(100, 0)

}

func (s *MsgServerTestSuite) TestFailsWithInvalidFeeTick() {
	s.fundAliceBalances(100, 0)

	// IF Alice Deposists 100
	s.aliceDeposits(NewDeposit(100, 0, 0, 0))

	// WHEN Alice tries to withdraw from an invalid tick
	err := s.aliceWithdraws(newWithdrawl(10, 0, 50))

	// THEN ensure error is thrown and Alice and Dex balances remain unchanged
	// TODO: Not sure how to trigger this failure case because we hit the insufficient shares check first
	s.Assert().Error(err)
	s.assertAliceShares(0, 0, sdk.NewDec(100))
	s.assertDexBalances(100, 0)

}

func (s *MsgServerTestSuite) TestFailsWithInvalidFeeMulti() {
	s.fundAliceBalances(100, 0)

	// IF Alice Deposists 100
	s.aliceDeposits(NewDeposit(100, 0, 0, 0))

	// WHEN Alice tries to withdraw from a mix of valid and invalid FeeTiers
	err := s.aliceWithdraws(
		// INVALID
		newWithdrawl(0, 10, 50),
		// VALID
		newWithdrawl(0, 0, 50),
	)

	// THEN ensure error is thrown and Alice and Dex balances remain unchanged
	// TODO: Not sure how to trigger this failure case because we hit the insufficient shares check first
	s.Assert().Error(err)
	s.assertAliceShares(0, 0, sdk.NewDec(100))
	s.assertDexBalances(100, 0)

}




