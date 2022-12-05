package keeper_test

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (s *MsgServerTestSuite) TestSingleDepositDepositAmountZero() {
	s.fundAliceBalances(25, 25)
	s.fundBobBalances(25, 25)

	_, err := s.msgServer.Deposit(s.goCtx, &types.MsgDeposit{
		Creator:     s.alice.String(),
		Receiver:    s.alice.String(),
		TokenA:      "TokenA",
		TokenB:      "TokenB",
		AmountsA:    []sdk.Int{sdk.NewInt(0)},
		AmountsB:    []sdk.Int{sdk.NewInt(0)},
		TickIndexes: []int64{0},
		FeeIndexes:  []uint64{1},
	})
	s.Assert().ErrorIs(err, sdkerrors.ErrInvalidType)

	s.assertAliceBalances(25, 25)
	s.assertDexBalances(0, 0)

}

func (s *MsgServerTestSuite) TestSingleDepositDepositFail() {
	s.fundAliceBalances(25, 25)
	s.fundBobBalances(25, 25)

	s.aliceDeposits(&Deposit{
		AmountA:   sdk.NewInt(5),
		AmountB:   sdk.NewInt(0),
		TickIndex: 0,
		FeeIndex:  0,
	})

	s.assertAliceBalances(20, 25)
	s.assertDexBalances(5, 0)

	_, err := s.msgServer.Deposit(s.goCtx, &types.MsgDeposit{
		Creator:     s.alice.String(),
		Receiver:    s.alice.String(),
		TokenA:      "TokenA",
		TokenB:      "TokenB",
		AmountsA:    []sdk.Int{sdk.NewInt(0)},
		AmountsB:    []sdk.Int{sdk.NewInt(5)},
		TickIndexes: []int64{0},
		FeeIndexes:  []uint64{0},
	})
	s.Assert().ErrorIs(err, types.ErrAllDepositsFailed)

	s.assertAliceBalances(20, 25)
	s.assertDexBalances(5, 0)

}

func (s *MsgServerTestSuite) TestMultiDepositDepositFail() {
	s.fundAliceBalances(25, 25)
	s.fundBobBalances(25, 25)

	s.aliceDeposits(&Deposit{
		AmountA:   sdk.NewInt(5),
		AmountB:   sdk.NewInt(0),
		TickIndex: 0,
		FeeIndex:  0,
	}, &Deposit{
		AmountA:   sdk.NewInt(5),
		AmountB:   sdk.NewInt(0),
		TickIndex: 1,
		FeeIndex:  0,
	})

	s.assertAliceBalances(15, 25)
	s.assertDexBalances(10, 0)

	_, err := s.msgServer.Deposit(s.goCtx, &types.MsgDeposit{
		Creator:     s.alice.String(),
		Receiver:    s.alice.String(),
		TokenA:      "TokenA",
		TokenB:      "TokenB",
		AmountsA:    []sdk.Int{sdk.NewInt(0), sdk.NewInt(0)},
		AmountsB:    []sdk.Int{sdk.NewInt(5), sdk.NewInt(5)},
		TickIndexes: []int64{0, 1},
		FeeIndexes:  []uint64{0, 0},
	})
	s.Assert().ErrorIs(err, types.ErrAllDepositsFailed)
	s.assertAliceBalances(15, 25)
	s.assertDexBalances(10, 0)

}

func (s *MsgServerTestSuite) TestMultiDepositSingleFail() {
	s.fundAliceBalances(25, 25)
	s.fundBobBalances(25, 25)

	s.aliceDeposits(&Deposit{
		AmountA:   sdk.NewInt(5),
		AmountB:   sdk.NewInt(0),
		TickIndex: 0,
		FeeIndex:  0,
	}, &Deposit{
		AmountA:   sdk.NewInt(5),
		AmountB:   sdk.NewInt(0),
		TickIndex: 1,
		FeeIndex:  0,
	})

	s.assertAliceBalances(15, 25)
	s.assertDexBalances(10, 0)

	depositResponse, err := s.msgServer.Deposit(s.goCtx, &types.MsgDeposit{
		Creator:     s.alice.String(),
		Receiver:    s.alice.String(),
		TokenA:      "TokenA",
		TokenB:      "TokenB",
		AmountsA:    []sdk.Int{sdk.NewInt(0), sdk.NewInt(5)},
		AmountsB:    []sdk.Int{sdk.NewInt(5), sdk.NewInt(0)},
		TickIndexes: []int64{0, 1},
		FeeIndexes:  []uint64{0, 0},
	})
	s.Require().Nil(err)

	s.assertDepositReponse(DepositReponse{
		amountsA: depositResponse.Reserve0Deposited,
		amountsB: depositResponse.Reserve1Deposited,
	}, DepositReponse{
		amountsA: []sdk.Int{sdk.ZeroInt(), sdk.NewInt(5)},
		amountsB: []sdk.Int{sdk.ZeroInt(), sdk.ZeroInt()},
	})

	s.assertAliceBalances(10, 25)
	s.assertDexBalances(15, 0)

}

func (s *MsgServerTestSuite) TestMultiDepositAllSucceed() {
	s.fundAliceBalances(25, 25)
	s.fundBobBalances(25, 25)

	s.aliceDeposits(&Deposit{
		AmountA:   sdk.NewInt(5),
		AmountB:   sdk.NewInt(0),
		TickIndex: 0,
		FeeIndex:  0,
	}, &Deposit{
		AmountA:   sdk.NewInt(5),
		AmountB:   sdk.NewInt(0),
		TickIndex: 1,
		FeeIndex:  0,
	})

	s.assertAliceBalances(15, 25)
	s.assertDexBalances(10, 0)

	depositResponse, err := s.msgServer.Deposit(s.goCtx, &types.MsgDeposit{
		Creator:     s.alice.String(),
		Receiver:    s.alice.String(),
		TokenA:      "TokenA",
		TokenB:      "TokenB",
		AmountsA:    []sdk.Int{sdk.NewInt(5), sdk.NewInt(5)},
		AmountsB:    []sdk.Int{sdk.NewInt(0), sdk.NewInt(0)},
		TickIndexes: []int64{0, 1},
		FeeIndexes:  []uint64{0, 0},
	})

	s.Require().Nil(err)

	s.assertDepositReponse(DepositReponse{
		amountsA: depositResponse.Reserve0Deposited,
		amountsB: depositResponse.Reserve1Deposited,
	}, DepositReponse{
		amountsA: []sdk.Int{sdk.NewInt(5), sdk.NewInt(5)},
		amountsB: []sdk.Int{sdk.ZeroInt(), sdk.ZeroInt()},
	})

	s.assertAliceBalances(5, 25)
	s.assertDexBalances(20, 0)
}
