package keeper_test

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (s *MsgServerTestSuite) TestPlaceLOVerificationSortTokenSameToken() {
	s.fundAliceBalances(50, 50)

	// Case
	// place limit order for 10 of token A at tick 0
	// Errors as TokenA and TokenB are the same token
	_, err := s.msgServer.PlaceLimitOrder(s.goCtx, &types.MsgPlaceLimitOrder{
		Creator:   s.alice.String(),
		Receiver:  s.alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenA",
		AmountIn:  sdk.NewDec(10),
		TokenIn:   "TokenA",
		TickIndex: 0,
	})
	ExpectedErr := types.ErrInvalidTokenPair
	s.Assert().ErrorIs(ExpectedErr, err)

}

func (s *MsgServerTestSuite) TestPlaceLOVerificationInvalidCreatorAddress() {
	s.fundAliceBalances(50, 50)

	// Case
	// place limit order for 10 of token A at tick 0
	// Errors as invalid creatorAddress
	_, err := s.msgServer.PlaceLimitOrder(s.goCtx, &types.MsgPlaceLimitOrder{
		Creator:   "",
		Receiver:  s.alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		AmountIn:  sdk.NewDec(10),
		TokenIn:   "TokenA",
		TickIndex: 0,
	})
	ExpectedErr := sdkerrors.ErrInvalidAddress
	s.Assert().ErrorIs(ExpectedErr, err)
}

func (s *MsgServerTestSuite) TestPlaceLOVerificationInvalidReceiverAddress() {
	s.fundAliceBalances(50, 50)
	// Case
	// place limit order for 10 of token A at tick 0
	// Errors as invalid receiverAddress
	_, err := s.msgServer.PlaceLimitOrder(s.goCtx, &types.MsgPlaceLimitOrder{
		Creator:   s.alice.String(),
		Receiver:  "",
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		AmountIn:  sdk.NewDec(10),
		TokenIn:   "TokenA",
		TickIndex: 0,
	})
	ExpectedErr := sdkerrors.ErrInvalidAddress
	s.Assert().ErrorIs(ExpectedErr, err)
}

func (s *MsgServerTestSuite) TestPlaceLOVerificationInvalidTokenIn() {
	s.fundAliceBalances(50, 50)

	// Case
	// place limit order for 10 of token A at tick 0
	// Errors as TokenIn must be TokenA or TokenB
	_, err := s.msgServer.PlaceLimitOrder(s.goCtx, &types.MsgPlaceLimitOrder{
		Creator:   s.alice.String(),
		Receiver:  s.alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		AmountIn:  sdk.NewDec(10),
		TokenIn:   "TokenC",
		TickIndex: 0,
	})
	ExpectedErr := types.ErrInvalidTokenPair
	s.Assert().ErrorIs(ExpectedErr, err)

}

func (s *MsgServerTestSuite) TestPlaceLOVerificationInvalidAmountIn() {
	s.fundAliceBalances(50, 50)

	// Case
	// place limit order for 10 of token A at tick 0
	// Errors as AmountIn > AmountOwned by alice
	fundsOwned := s.aliceGetBalanceDec("TokenA")

	_, err := s.msgServer.PlaceLimitOrder(s.goCtx, &types.MsgPlaceLimitOrder{
		Creator:   s.alice.String(),
		Receiver:  s.alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		AmountIn:  fundsOwned.Mul(sdk.NewDec(2)),
		TokenIn:   "TokenA",
		TickIndex: 0,
	})
	ExpectedErr := types.ErrNotEnoughCoins
	s.Assert().ErrorIs(ExpectedErr, err)
}
