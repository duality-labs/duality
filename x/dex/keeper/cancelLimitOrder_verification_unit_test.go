package keeper_test

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (s *MsgServerTestSuite) TestCancelLOVerificationSortTokenSameToken() {
	s.fundAliceBalances(50, 50)

	// Case
	// cancel limit order for 10 of token A at tick 0
	// Errors as TokenA and TokenB are the same token
	_, err := s.msgServer.CancelLimitOrder(s.goCtx, &types.MsgCancelLimitOrder{
		Creator:   s.alice.String(),
		Receiver:  s.alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenA",
		Key:       0,
		KeyToken:  "TokenA",
		SharesOut: sdk.NewDec(10),
		TickIndex: 0,
	})
	ExpectedErr := types.ErrInvalidTokenPair
	s.Assert().ErrorIs(ExpectedErr, err)

}

func (s *MsgServerTestSuite) TestCancelLOVerificationInvalidCreatorAddress() {
	s.fundAliceBalances(50, 50)

	// Case
	// cancel limit order for 10 of token A at tick 0
	// Errors as invalid creatorAddress
	_, err := s.msgServer.CancelLimitOrder(s.goCtx, &types.MsgCancelLimitOrder{
		Creator:   "",
		Receiver:  s.alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		Key:       0,
		KeyToken:  "TokenA",
		SharesOut: sdk.NewDec(10),
		TickIndex: 0,
	})
	ExpectedErr := sdkerrors.ErrInvalidAddress
	s.Assert().ErrorIs(ExpectedErr, err)

}

func (s *MsgServerTestSuite) TestCancelLOVerificationInvalidReceiverAddress() {
	s.fundAliceBalances(50, 50)

	// Case
	// cancel limit order for 10 of token A at tick 0
	// Errors as invalid receiverAddress
	_, err := s.msgServer.CancelLimitOrder(s.goCtx, &types.MsgCancelLimitOrder{
		Creator:   s.alice.String(),
		Receiver:  "",
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		Key:       0,
		KeyToken:  "TokenA",
		SharesOut: sdk.NewDec(10),
		TickIndex: 0,
	})
	ExpectedErr := sdkerrors.ErrInvalidAddress
	s.Assert().ErrorIs(ExpectedErr, err)
}

func (s *MsgServerTestSuite) TestCancelLOVerificationInvalidKeyToken() {
	s.fundAliceBalances(50, 50)

	// Case
	// place limit order for 10 of token A at tick 0
	// Errors as KeyToken must be TokenA or TokenB
	_, err := s.msgServer.CancelLimitOrder(s.goCtx, &types.MsgCancelLimitOrder{
		Creator:   s.alice.String(),
		Receiver:  s.alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		Key:       0,
		KeyToken:  "TokenC",
		SharesOut: sdk.NewDec(10),
		TickIndex: 0,
	})
	ExpectedErr := types.ErrInvalidTokenPair
	s.Assert().ErrorIs(ExpectedErr, err)

}

func (s *MsgServerTestSuite) TestCancelLOVerificationNoSharesFound() {
	s.fundAliceBalances(50, 50)

	// Case
	// cancel limit order for 10 of token A at tick 0
	// Errors as SharesOut is not found

	_, err := s.msgServer.CancelLimitOrder(s.goCtx, &types.MsgCancelLimitOrder{
		Creator:   s.alice.String(),
		Receiver:  s.alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		Key:       0,
		KeyToken:  "TokenA",
		SharesOut: sdk.NewDec(10),
		TickIndex: 0,
	})
	ExpectedErr := types.ErrNotEnoughShares
	s.Assert().ErrorIs(ExpectedErr, err)
}

func (s *MsgServerTestSuite) TestCancelLOVerificationNotEnoughShares() {
	s.fundAliceBalances(50, 50)
	s.aliceLimitSells("TokenA", 0, 5)
	// Case
	// cancel limit order for 10 of token A at tick 0
	// Errors as SharesOut is not found

	_, err := s.msgServer.CancelLimitOrder(s.goCtx, &types.MsgCancelLimitOrder{
		Creator:   s.alice.String(),
		Receiver:  s.alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		Key:       0,
		KeyToken:  "TokenA",
		SharesOut: sdk.NewDec(10),
		TickIndex: 0,
	})

	ExpectedErr := types.ErrNotEnoughShares
	s.Assert().ErrorIs(ExpectedErr, err)
}
