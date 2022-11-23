package keeper_test

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (s *MsgServerTestSuite) TestWithdrawFilledLOVerificationSortTokenSameToken() {
	s.fundAliceBalances(50, 50)

	// Case
	// withdraw Filled limit order for 10 of token A at tick 0, key 0
	// Errors as TokenA and TokenB are the same token
	_, err := s.msgServer.WithdrawFilledLimitOrder(s.goCtx, &types.MsgWithdrawFilledLimitOrder{
		Creator:   s.alice.String(),
		Receiver:  s.alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenA",
		Key:       0,
		KeyToken:  "TokenA",
		TickIndex: 0,
	})
	ExpectedErr := types.ErrInvalidTokenPair
	s.Assert().ErrorIs(ExpectedErr, err)

}

func (s *MsgServerTestSuite) TestWithdrawFilledLOVerificationInvalidCreatorAddress() {
	s.fundAliceBalances(50, 50)

	// Case
	// withdraw filled limit order for 10 of token A at tick 0
	// Errors as invalid creatorAddress
	_, err := s.msgServer.WithdrawFilledLimitOrder(s.goCtx, &types.MsgWithdrawFilledLimitOrder{
		Creator:   "",
		Receiver:  s.alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		Key:       0,
		KeyToken:  "TokenA",
		TickIndex: 0,
	})
	ExpectedErr := sdkerrors.ErrInvalidAddress
	s.Assert().ErrorIs(ExpectedErr, err)

}

func (s *MsgServerTestSuite) TestWithdrawFilledLOVerificationInvalidReceiverAddress() {
	s.fundAliceBalances(50, 50)

	// Case
	// withdraw filled limit order for 10 of token A at tick 0
	// Errors as invalid receiverAddress
	_, err := s.msgServer.WithdrawFilledLimitOrder(s.goCtx, &types.MsgWithdrawFilledLimitOrder{
		Creator:   s.alice.String(),
		Receiver:  "",
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		Key:       0,
		KeyToken:  "TokenA",
		TickIndex: 0,
	})
	ExpectedErr := sdkerrors.ErrInvalidAddress
	s.Assert().ErrorIs(ExpectedErr, err)
}

func (s *MsgServerTestSuite) TestWithdrawFilledLOVerificationInvalidKeyToken() {
	s.fundAliceBalances(50, 50)

	// Case
	// withdraw filled limit order for 10 of token A at tick 0
	// Errors as KeyToken must be TokenA or TokenB
	_, err := s.msgServer.WithdrawFilledLimitOrder(s.goCtx, &types.MsgWithdrawFilledLimitOrder{
		Creator:   s.alice.String(),
		Receiver:  s.alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		Key:       0,
		KeyToken:  "TokenC",
		TickIndex: 0,
	})
	ExpectedErr := types.ErrInvalidTokenPair
	s.Assert().ErrorIs(ExpectedErr, err)

}

func (s *MsgServerTestSuite) TestWithdrawFilledLOVerificationNoSharesFound() {
	s.fundAliceBalances(50, 50)

	// Case
	// withdraw filled limit order for 10 of token A at tick 0
	// Errors as SharesOut is not found

	_, err := s.msgServer.WithdrawFilledLimitOrder(s.goCtx, &types.MsgWithdrawFilledLimitOrder{
		Creator:   s.alice.String(),
		Receiver:  s.alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		Key:       0,
		KeyToken:  "TokenB",
		TickIndex: 0,
	})
	ExpectedErr := types.ErrNotEnoughShares
	s.Assert().ErrorIs(ExpectedErr, err)
}

func (s *MsgServerTestSuite) TestWithdrawFilledLOVerificationNotEnoughShares() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 50)
	s.aliceLimitSells("TokenA", 0, 5)
	s.bobMarketSells("TokenB", 5, 5)
	s.aliceWithdrawsLimitSell("TokenA", 0, 0)
	// Case
	// cancel limit order for 10 of token A at tick 0
	// Errors as SharesOut is already 0

	_, err := s.msgServer.WithdrawFilledLimitOrder(s.goCtx, &types.MsgWithdrawFilledLimitOrder{
		Creator:   s.alice.String(),
		Receiver:  s.alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		Key:       0,
		KeyToken:  "TokenB",
		TickIndex: 0,
	})

	ExpectedErr := types.ErrNotEnoughShares
	s.Assert().ErrorIs(ExpectedErr, err)
}
