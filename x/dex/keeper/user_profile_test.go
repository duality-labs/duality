package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/duality-labs/duality/x/dex/keeper"
	"github.com/duality-labs/duality/x/dex/types"
)

func (s *MsgServerTestSuite) TestGetAllLimitOrders() {
	// WHEN Alice places 2 limit orders
	s.fundAliceBalances(20, 20)
	s.fundBobBalances(20, 20)
	trancheKeyA := s.aliceLimitSells("TokenA", -1, 10)
	trancheKeyB := s.aliceLimitSells("TokenB", 0, 10)
	s.bobLimitSells("TokenA", -1, 10)
	profile := NewUserProfile(s.alice)

	// THEN GetAllLimitOrders returns alice's same two orders
	LOList := profile.GetAllLimitOrders(s.ctx, s.app.DexKeeper)
	s.Assert().Equal(2, len(LOList))
	s.Assert().Equal(types.LimitOrderTrancheUser{
		PairID:          defaultPairID,
		Token:           "TokenA",
		TickIndex:       -1,
		TrancheKey:      trancheKeyA,
		Address:         s.alice.String(),
		SharesOwned:     sdk.NewInt(10),
		SharesWithdrawn: sdk.NewInt(0),
		SharesCancelled: sdk.NewInt(0),
		TakerReserves:   sdk.ZeroInt(),
	},
		LOList[0],
	)
	s.Assert().Equal(types.LimitOrderTrancheUser{
		PairID:          defaultPairID,
		Token:           "TokenB",
		TickIndex:       0,
		TrancheKey:      trancheKeyB,
		Address:         s.alice.String(),
		SharesOwned:     sdk.NewInt(10),
		SharesWithdrawn: sdk.NewInt(0),
		SharesCancelled: sdk.NewInt(0),
		TakerReserves:   sdk.ZeroInt(),
	},
		LOList[1],
	)
}

func (s *MsgServerTestSuite) TestGetAllDeposits() {
	s.fundAliceBalances(20, 20)
	// GIVEN Alice Deposits 3 positions and withdraws the first
	s.aliceDeposits(
		&Deposit{
			AmountA:   sdk.NewInt(1),
			AmountB:   sdk.NewInt(0),
			TickIndex: -50,
			Fee:       1,
		},
		&Deposit{
			AmountA:   sdk.NewInt(5),
			AmountB:   sdk.NewInt(5),
			TickIndex: 0,
			Fee:       1,
		},
		&Deposit{
			AmountA:   sdk.NewInt(0),
			AmountB:   sdk.NewInt(10),
			TickIndex: 2,
			Fee:       1,
		},
	)
	s.aliceWithdraws(&Withdrawl{
		TickIndex: -50,
		Fee:       1,
		Shares:    sdk.NewInt(1),
	},
	)
	profile := NewUserProfile(s.alice)

	// THEN GetAllDeposits returns the two remaining LP positions
	depositList := profile.GetAllDeposits(s.ctx, s.app.DexKeeper)
	s.Assert().Equal(2, len(depositList))
	s.Assert().Equal(types.DepositRecord{
		PairID:          defaultPairID,
		SharesOwned:     sdk.NewInt(10),
		CenterTickIndex: 0,
		LowerTickIndex:  -1,
		UpperTickIndex:  1,
		Fee:             1,
	},
		depositList[0],
	)
	s.Assert().Equal(types.DepositRecord{
		PairID:          defaultPairID,
		SharesOwned:     sdk.NewInt(10),
		CenterTickIndex: 2,
		LowerTickIndex:  1,
		UpperTickIndex:  3,
		Fee:             1,
	},
		depositList[1],
	)
}

func (s *MsgServerTestSuite) TestGetAllPositions() {
	s.fundAliceBalances(50, 50)
	s.aliceDeposits(
		&Deposit{
			AmountA:   sdk.NewInt(1),
			AmountB:   sdk.NewInt(0),
			TickIndex: 1,
			Fee:       1,
		},
		&Deposit{
			AmountA:   sdk.NewInt(5),
			AmountB:   sdk.NewInt(5),
			TickIndex: 0,
			Fee:       1,
		},
	)

	s.aliceLimitSells("TokenA", 0, 10)
	profile := NewUserProfile(s.alice)
	positions := profile.GetAllPositions(s.ctx, s.app.DexKeeper)

	s.Assert().Equal(2, len(positions.PoolDeposits))
	s.Assert().Equal(1, len(positions.LimitOrders))
}
