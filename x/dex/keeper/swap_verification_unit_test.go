package keeper_test

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (s *MsgServerTestSuite) TestSwapVerificationSortTokenSameToken() {
	s.fundAliceBalances(50, 50)

	// Case
	// swap 10 of token A for token B
	// Errors as TokenA and TokenB are the same token
	_, err := s.msgServer.Swap(s.goCtx, &types.MsgSwap{
		Creator:  s.alice.String(),
		Receiver: s.alice.String(),
		TokenA:   "TokenA",
		TokenB:   "TokenA",
		AmountIn: sdk.NewDec(10),
		TokenIn:  "TokenA",
		MinOut:   sdk.NewDec(10),
	})
	ExpectedErr := types.ErrInvalidTokenPair
	s.Assert().ErrorIs(ExpectedErr, err)

}

func (s *MsgServerTestSuite) TestSwapVerificationInvalidCreatorAddress() {
	s.fundAliceBalances(50, 50)

	// Case
	// swap 10 of token A for token B
	// Errors as invalid creatorAddress
	_, err := s.msgServer.Swap(s.goCtx, &types.MsgSwap{
		Creator:  "",
		Receiver: s.alice.String(),
		TokenA:   "TokenA",
		TokenB:   "TokenB",
		AmountIn: sdk.NewDec(10),
		TokenIn:  "TokenA",
		MinOut:   sdk.NewDec(10),
	})
	ExpectedErr := sdkerrors.ErrInvalidAddress
	s.Assert().ErrorIs(ExpectedErr, err)
}

func (s *MsgServerTestSuite) TestSwapVerificationInvalidReceiverAddress() {
	s.fundAliceBalances(50, 50)
	// Case
	// swap 10 of token A for token B
	// Errors as invalid receiverAddress
	_, err := s.msgServer.Swap(s.goCtx, &types.MsgSwap{
		Creator:  s.alice.String(),
		Receiver: "",
		TokenA:   "TokenA",
		TokenB:   "TokenB",
		AmountIn: sdk.NewDec(10),
		TokenIn:  "TokenA",
		MinOut:   sdk.NewDec(10),
	})
	ExpectedErr := sdkerrors.ErrInvalidAddress
	s.Assert().ErrorIs(ExpectedErr, err)
}

func (s *MsgServerTestSuite) TestSwapVerificationInvalidTokenIn() {
	s.fundAliceBalances(50, 50)
	// Case
	// swap 10 of token A for token B
	// Errors as TokenIn must be TokenA or TokenB
	_, err := s.msgServer.Swap(s.goCtx, &types.MsgSwap{
		Creator:  s.alice.String(),
		Receiver: s.alice.String(),
		TokenA:   "TokenA",
		TokenB:   "TokenB",
		AmountIn: sdk.NewDec(10),
		TokenIn:  "TokenC",
		MinOut:   sdk.NewDec(10),
	})
	ExpectedErr := types.ErrInvalidTokenPair
	s.Assert().ErrorIs(ExpectedErr, err)
}

func (s *MsgServerTestSuite) TestSwapVerificationInvalidAmountIn() {
	s.fundAliceBalances(50, 50)

	// Case
	// swap 10 of token A for token B
	// Errors as AmountIn is greater than amountOwned
	fundsOwned := s.aliceGetBalanceDec("TokenA")
	_, err := s.msgServer.Swap(s.goCtx, &types.MsgSwap{
		Creator:  s.alice.String(),
		Receiver: s.alice.String(),
		TokenA:   "TokenA",
		TokenB:   "TokenB",
		AmountIn: fundsOwned.Mul(sdk.NewDec(2)),
		TokenIn:  "TokenA",
		MinOut:   sdk.NewDec(10),
	})
	ExpectedErr := types.ErrNotEnoughCoins
	s.Assert().ErrorIs(ExpectedErr, err)
}
