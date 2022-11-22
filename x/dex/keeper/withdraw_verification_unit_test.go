package keeper_test

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (s *MsgServerTestSuite) TestWithdrawVerificationSortTokenSameToken() {
	s.fundAliceBalances(50, 50)

	// Case
	// withdrawl 10 of token A at tick 0 fee 1
	// Errors as TokenA and TokenB are the same token
	_, err := s.msgServer.Withdrawl(s.goCtx, &types.MsgWithdrawl{
		Creator:        s.alice.String(),
		Receiver:       s.alice.String(),
		TokenA:         "TokenA",
		TokenB:         "TokenA",
		SharesToRemove: []sdk.Dec{sdk.NewDec(10)},
		TickIndexes:    []int64{0},
		FeeIndexes:     []uint64{0},
	})
	ExpectedErr := types.ErrInvalidTokenPair
	s.Assert().ErrorIs(ExpectedErr, err)

}

func (s *MsgServerTestSuite) TestWithdrawVerificationInvalidCreatorAddress() {
	s.fundAliceBalances(50, 50)

	// Case
	// withdraw 10 of token A at tick 0 fee 1
	// Errors as creator address is an invalid address
	_, err := s.msgServer.Withdrawl(s.goCtx, &types.MsgWithdrawl{
		Creator:        "",
		Receiver:       s.alice.String(),
		TokenA:         "TokenA",
		TokenB:         "TokenB",
		SharesToRemove: []sdk.Dec{sdk.NewDec(10)},
		TickIndexes:    []int64{0},
		FeeIndexes:     []uint64{0},
	})
	ExpectedErr := sdkerrors.ErrInvalidAddress
	s.Assert().ErrorIs(ExpectedErr, err)
}

func (s *MsgServerTestSuite) TestWithdrawVerificationInvalidReceiverAddress() {
	s.fundAliceBalances(50, 50)

	// Case
	// withdraw 10 of token A at tick 0 fee 1
	// Errors as receiver address is an invalid address
	_, err := s.msgServer.Withdrawl(s.goCtx, &types.MsgWithdrawl{
		Creator:        s.alice.String(),
		Receiver:       "",
		TokenA:         "TokenA",
		TokenB:         "TokenB",
		SharesToRemove: []sdk.Dec{sdk.NewDec(10)},
		TickIndexes:    []int64{0},
		FeeIndexes:     []uint64{0},
	})
	ExpectedErr := sdkerrors.ErrInvalidAddress
	s.Assert().ErrorIs(ExpectedErr, err)
}

func (s *MsgServerTestSuite) TestWithdrawVerificationInvalidSingleFeeTier() {
	s.fundAliceBalances(50, 50)

	// Case
	// withdraw 10 of token A at tick 0 fee 1
	// Errors as length of feeTier is out of range of valid feeIndices

	_, err := s.msgServer.Withdrawl(s.goCtx, &types.MsgWithdrawl{
		Creator:        s.alice.String(),
		Receiver:       s.alice.String(),
		TokenA:         "TokenA",
		TokenB:         "TokenB",
		SharesToRemove: []sdk.Dec{sdk.NewDec(10)},
		TickIndexes:    []int64{0},
		FeeIndexes:     []uint64{uint64(len(s.feeTiers))},
	})

	ExpectedErr := types.ErrValidFeeIndexNotFound
	s.Assert().ErrorIs(ExpectedErr, err)
}

func (s *MsgServerTestSuite) TestWithdrawVerificationInvalidMultiFeeTier() {
	s.fundAliceBalances(50, 50)

	// Case
	// withdraw 10 of token A at tick 0 fee 1
	// Errors as second fee index is the length of feeTier which is out of range of valid feeIndices
	_, err := s.msgServer.Withdrawl(s.goCtx, &types.MsgWithdrawl{
		Creator:        s.alice.String(),
		Receiver:       s.alice.String(),
		TokenA:         "TokenA",
		TokenB:         "TokenB",
		SharesToRemove: []sdk.Dec{sdk.NewDec(10), sdk.NewDec(10)},
		TickIndexes:    []int64{0, 0},
		FeeIndexes:     []uint64{0, uint64(len(s.feeTiers))},
	})
	ExpectedErr := types.ErrValidFeeIndexNotFound
	s.Assert().ErrorIs(ExpectedErr, err)
}

func (s *MsgServerTestSuite) TestWithdrawVerificationInvalidSharesToRemoveArray() {
	s.fundAliceBalances(50, 50)

	// Case
	// withdraw 10 of token A at tick 0 fee 1
	// Errors as SharesToRemove array is not the same size as d TickIndexes, FeeIndexes
	_, err := s.msgServer.Withdrawl(s.goCtx, &types.MsgWithdrawl{
		Creator:        s.alice.String(),
		Receiver:       s.alice.String(),
		TokenA:         "TokenA",
		TokenB:         "TokenB",
		SharesToRemove: []sdk.Dec{sdk.NewDec(10)},
		TickIndexes:    []int64{0, 0},
		FeeIndexes:     []uint64{0, 0},
	})
	ExpectedErr := types.ErrUnbalancedTxArray
	s.Assert().ErrorIs(ExpectedErr, err)
}

func (s *MsgServerTestSuite) TestWithdrawVerificationInvalidTickIndexesArray() {
	s.fundAliceBalances(50, 50)

	// Case
	// withdraw 10 of token A at tick 0 fee 1
	// Errors as TickIndexes array is not the same size as SharesToRemove, FeeIndexes
	_, err := s.msgServer.Withdrawl(s.goCtx, &types.MsgWithdrawl{
		Creator:        s.alice.String(),
		Receiver:       s.alice.String(),
		TokenA:         "TokenA",
		TokenB:         "TokenB",
		SharesToRemove: []sdk.Dec{sdk.NewDec(10), sdk.NewDec(10)},
		TickIndexes:    []int64{0},
		FeeIndexes:     []uint64{0, 0},
	})
	ExpectedErr := types.ErrUnbalancedTxArray
	s.Assert().ErrorIs(ExpectedErr, err)
}

func (s *MsgServerTestSuite) TestWithdrawVerificationInvalidFeeIndexesArray() {
	s.fundAliceBalances(50, 50)

	// Case
	// withdraw 10 of token A at tick 0 fee 1
	// Errors as FeeIndexes array is not the same size as SharesToRemove, TickIndexes
	_, err := s.msgServer.Withdrawl(s.goCtx, &types.MsgWithdrawl{
		Creator:        s.alice.String(),
		Receiver:       s.alice.String(),
		TokenA:         "TokenA",
		TokenB:         "TokenB",
		SharesToRemove: []sdk.Dec{sdk.NewDec(10), sdk.NewDec(10)},
		TickIndexes:    []int64{0, 0},
		FeeIndexes:     []uint64{0},
	})
	ExpectedErr := types.ErrUnbalancedTxArray
	s.Assert().ErrorIs(ExpectedErr, err)
}

func (s *MsgServerTestSuite) TestWithdrawVerificationNoSharesOwned() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 50)
	// Case
	// withdraw 10 of token A at tick 0 fee 1
	// Errors as the valid shares exist
	s.bobDeposits(NewDeposit(0, 10, 0, 0))

	_, err := s.msgServer.Withdrawl(s.goCtx, &types.MsgWithdrawl{
		Creator:        s.alice.String(),
		Receiver:       s.alice.String(),
		TokenA:         "TokenA",
		TokenB:         "TokenB",
		SharesToRemove: []sdk.Dec{sdk.NewDec(10), sdk.NewDec(10)},
		TickIndexes:    []int64{0, 0},
		FeeIndexes:     []uint64{0, 0},
	})
	ExpectedErr := types.ErrNotEnoughShares
	s.Assert().ErrorIs(ExpectedErr, err)
}

func (s *MsgServerTestSuite) TestWithdrawVerificationNotEnoughShares() {
	s.fundAliceBalances(50, 50)
	// Case
	// withdraw 10 of token A at tick 0 fee 1
	// Errors as the valid shares exist
	s.aliceDeposits(NewDeposit(0, 10, 0, 0))

	_, err := s.msgServer.Withdrawl(s.goCtx, &types.MsgWithdrawl{
		Creator:        s.alice.String(),
		Receiver:       s.alice.String(),
		TokenA:         "TokenA",
		TokenB:         "TokenB",
		SharesToRemove: []sdk.Dec{sdk.NewDec(15)},
		TickIndexes:    []int64{0},
		FeeIndexes:     []uint64{0},
	})
	ExpectedErr := types.ErrNotEnoughShares
	s.Assert().ErrorIs(ExpectedErr, err)
}
