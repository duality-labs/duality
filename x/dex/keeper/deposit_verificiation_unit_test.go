package keeper_test

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (s *MsgServerTestSuite) TestDepositVerificationSortTokenSameToken() {
	s.fundAliceBalances(50, 50)

	// Case
	// deposit 10 of token A at tick 0 fee 1
	// Errors as TokenA and TokenB are the same token
	_, err := s.msgServer.Deposit(s.goCtx, &types.MsgDeposit{
		Creator:     s.alice.String(),
		Receiver:    s.alice.String(),
		TokenA:      "TokenA",
		TokenB:      "TokenA",
		AmountsA:    []sdk.Dec{sdk.NewDec(10)},
		AmountsB:    []sdk.Dec{sdk.NewDec(0)},
		TickIndexes: []int64{0},
		FeeIndexes:  []uint64{0},
	})
	ExpectedErr := types.ErrInvalidTokenPair
	s.Assert().ErrorIs(ExpectedErr, err)

}

func (s *MsgServerTestSuite) TestDepositVerificationInvalidCreatorAddress() {
	s.fundAliceBalances(50, 50)

	// Case
	// deposit 10 of token A at tick 0 fee 1
	// Errors as creator is an invalid address
	_, err := s.msgServer.Deposit(s.goCtx, &types.MsgDeposit{
		Creator:     "",
		Receiver:    s.alice.String(),
		TokenA:      "TokenA",
		TokenB:      "TokenB",
		AmountsA:    []sdk.Dec{sdk.NewDec(10)},
		AmountsB:    []sdk.Dec{sdk.NewDec(0)},
		TickIndexes: []int64{0},
		FeeIndexes:  []uint64{0},
	})
	ExpectedErr := sdkerrors.ErrInvalidAddress
	s.Assert().ErrorIs(ExpectedErr, err)
}

func (s *MsgServerTestSuite) TestDepositVerificationInvalidReceiverAddress() {
	s.fundAliceBalances(50, 50)

	// Case
	// deposit 10 of token A at tick 0 fee 1
	// Errors as receiver is an invalid Address
	_, err := s.msgServer.Deposit(s.goCtx, &types.MsgDeposit{
		Creator:     s.alice.String(),
		Receiver:    "",
		TokenA:      "TokenA",
		TokenB:      "TokenB",
		AmountsA:    []sdk.Dec{sdk.NewDec(10)},
		AmountsB:    []sdk.Dec{sdk.NewDec(0)},
		TickIndexes: []int64{0},
		FeeIndexes:  []uint64{0},
	})
	ExpectedErr := sdkerrors.ErrInvalidAddress
	s.Assert().ErrorIs(ExpectedErr, err)
}

func (s *MsgServerTestSuite) TestDepositVerificationInvalidSingleFeeTier() {
	s.fundAliceBalances(50, 50)

	// Case
	// deposit 10 of token A at tick 0 fee 1
	// Errors as length of feeTier is out of range of valid feeIndices
	_, err := s.msgServer.Deposit(s.goCtx, &types.MsgDeposit{
		Creator:     s.alice.String(),
		Receiver:    s.alice.String(),
		TokenA:      "TokenA",
		TokenB:      "TokenB",
		AmountsA:    []sdk.Dec{sdk.NewDec(10)},
		AmountsB:    []sdk.Dec{sdk.NewDec(0)},
		TickIndexes: []int64{0},
		FeeIndexes:  []uint64{uint64(len(s.feeTiers))},
	})
	ExpectedErr := types.ErrValidFeeIndexNotFound
	s.Assert().ErrorIs(ExpectedErr, err)
}

func (s *MsgServerTestSuite) TestDepositVerificationInvalidMultiFeeTier() {
	s.fundAliceBalances(50, 50)

	// Case
	// deposit 10 of token A at tick 0 fee 1
	// Errors as second fee index is the length of feeTier which is out of range of valid feeIndices
	_, err := s.msgServer.Deposit(s.goCtx, &types.MsgDeposit{
		Creator:     s.alice.String(),
		Receiver:    s.alice.String(),
		TokenA:      "TokenA",
		TokenB:      "TokenB",
		AmountsA:    []sdk.Dec{sdk.NewDec(10), sdk.NewDec(10)},
		AmountsB:    []sdk.Dec{sdk.NewDec(0), sdk.NewDec(0)},
		TickIndexes: []int64{0, 0},
		FeeIndexes:  []uint64{0, uint64(len(s.feeTiers))},
	})
	ExpectedErr := types.ErrValidFeeIndexNotFound
	s.Assert().ErrorIs(ExpectedErr, err)
}

func (s *MsgServerTestSuite) TestDepositVerificationInvalidAmountsAArray() {
	s.fundAliceBalances(50, 50)

	// Case
	// deposit 10 of token A at tick 0 fee 1
	// Errors as AmountsA array is not the same size as AmountsB, TickIndexes, FeeIndexes
	_, err := s.msgServer.Deposit(s.goCtx, &types.MsgDeposit{
		Creator:     s.alice.String(),
		Receiver:    s.alice.String(),
		TokenA:      "TokenA",
		TokenB:      "TokenB",
		AmountsA:    []sdk.Dec{sdk.NewDec(10)},
		AmountsB:    []sdk.Dec{sdk.NewDec(0), sdk.NewDec(0)},
		TickIndexes: []int64{0, 0},
		FeeIndexes:  []uint64{0, 0},
	})
	ExpectedErr := types.ErrUnbalancedTxArray
	s.Assert().ErrorIs(ExpectedErr, err)
}

func (s *MsgServerTestSuite) TestDepositVerificationInvalidAmountBArray() {
	s.fundAliceBalances(50, 50)

	// Case
	// deposit 10 of token A at tick 0 fee 1
	// Errors as AmountsB array is not the same size as AmountsA, TickIndexes, FeeIndexes
	_, err := s.msgServer.Deposit(s.goCtx, &types.MsgDeposit{
		Creator:     s.alice.String(),
		Receiver:    s.alice.String(),
		TokenA:      "TokenA",
		TokenB:      "TokenB",
		AmountsA:    []sdk.Dec{sdk.NewDec(10), sdk.NewDec(10)},
		AmountsB:    []sdk.Dec{sdk.NewDec(0)},
		TickIndexes: []int64{0, 0},
		FeeIndexes:  []uint64{0, 0},
	})
	ExpectedErr := types.ErrUnbalancedTxArray
	s.Assert().ErrorIs(ExpectedErr, err)
}

func (s *MsgServerTestSuite) TestDepositVerificationInvalidTickIndexesArray() {
	s.fundAliceBalances(50, 50)

	// Case
	// deposit 10 of token A at tick 0 fee 1
	// Errors as TickIndexes array is not the same size as AmountsA, AmountsB, FeeIndexes
	_, err := s.msgServer.Deposit(s.goCtx, &types.MsgDeposit{
		Creator:     s.alice.String(),
		Receiver:    s.alice.String(),
		TokenA:      "TokenA",
		TokenB:      "TokenB",
		AmountsA:    []sdk.Dec{sdk.NewDec(10), sdk.NewDec(10)},
		AmountsB:    []sdk.Dec{sdk.NewDec(0), sdk.NewDec(0)},
		TickIndexes: []int64{0},
		FeeIndexes:  []uint64{0, 0},
	})
	ExpectedErr := types.ErrUnbalancedTxArray
	s.Assert().ErrorIs(ExpectedErr, err)
}

func (s *MsgServerTestSuite) TestDepositVerificationInvalidFeeIndexesArray() {
	s.fundAliceBalances(50, 50)

	// Case
	// deposit 10 of token A at tick 0 fee 1
	// Errors as FeeIndexes array is not the same size as AmountsA, AmountsB, TickIndexes
	_, err := s.msgServer.Deposit(s.goCtx, &types.MsgDeposit{
		Creator:     s.alice.String(),
		Receiver:    s.alice.String(),
		TokenA:      "TokenA",
		TokenB:      "TokenB",
		AmountsA:    []sdk.Dec{sdk.NewDec(10), sdk.NewDec(10)},
		AmountsB:    []sdk.Dec{sdk.NewDec(0), sdk.NewDec(0)},
		TickIndexes: []int64{0, 0},
		FeeIndexes:  []uint64{0},
	})
	ExpectedErr := types.ErrUnbalancedTxArray
	s.Assert().ErrorIs(ExpectedErr, err)
}

func (s *MsgServerTestSuite) TestDepositVerificationAllZeroAmountFail() {
	s.fundAliceBalances(50, 50)

	// Case
	// deposit 10 of token A at tick 0 fee 1
	// Errors as the first entry in AmountsA/AmountsB is zero for both entries
	_, err := s.msgServer.Deposit(s.goCtx, &types.MsgDeposit{
		Creator:     s.alice.String(),
		Receiver:    s.alice.String(),
		TokenA:      "TokenA",
		TokenB:      "TokenB",
		AmountsA:    []sdk.Dec{sdk.NewDec(0), sdk.NewDec(10)},
		AmountsB:    []sdk.Dec{sdk.NewDec(0), sdk.NewDec(0)},
		TickIndexes: []int64{0, 0},
		FeeIndexes:  []uint64{0, 0},
	})
	ExpectedErr := sdkerrors.ErrInvalidType
	s.Assert().ErrorIs(ExpectedErr, err)
}

func (s *MsgServerTestSuite) TestDepositVerificationNotEnoughTokenA() {
	s.fundAliceBalances(50, 50)

	// Case
	// deposit 10 of token A at tick 0 fee 1
	// Errors as the first entry in AmountsA/AmountsB is zero for both entries
	fundsOwned := s.aliceGetBalanceDec("TokenA")
	_, err := s.msgServer.Deposit(s.goCtx, &types.MsgDeposit{
		Creator:     s.alice.String(),
		Receiver:    s.alice.String(),
		TokenA:      "TokenA",
		TokenB:      "TokenB",
		AmountsA:    []sdk.Dec{fundsOwned, fundsOwned},
		AmountsB:    []sdk.Dec{sdk.NewDec(0), sdk.NewDec(0)},
		TickIndexes: []int64{0, 0},
		FeeIndexes:  []uint64{0, 0},
	})
	ExpectedErr := types.ErrNotEnoughCoins
	s.Assert().ErrorIs(ExpectedErr, err)
}

func (s *MsgServerTestSuite) TestDepositVerificationNotEnoughTokenB() {
	s.fundAliceBalances(50, 50)

	// Case
	// deposit 10 of token A at tick 0 fee 1
	// Errors as the first entry in AmountsA/AmountsB is zero for both entries

	fundsOwned := s.aliceGetBalanceDec("TokenB")
	_, err := s.msgServer.Deposit(s.goCtx, &types.MsgDeposit{
		Creator:     s.alice.String(),
		Receiver:    s.alice.String(),
		TokenA:      "TokenA",
		TokenB:      "TokenB",
		AmountsA:    []sdk.Dec{sdk.NewDec(0), sdk.NewDec(0)},
		AmountsB:    []sdk.Dec{fundsOwned, fundsOwned},
		TickIndexes: []int64{0, 0},
		FeeIndexes:  []uint64{0, 0},
	})
	ExpectedErr := types.ErrNotEnoughCoins
	s.Assert().ErrorIs(ExpectedErr, err)
}
