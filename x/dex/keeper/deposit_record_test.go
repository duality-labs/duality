package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
)

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
	s.aliceWithdraws(&Withdrawal{
		TickIndex: -50,
		Fee:       1,
		Shares:    sdk.NewInt(1),
	},
	)

	// THEN GetAllDeposits returns the two remaining LP positions
	depositList := s.app.DexKeeper.GetAllDepositsForAddress(s.ctx, s.alice)
	s.Assert().Equal(2, len(depositList))
	s.Assert().Equal(&types.DepositRecord{
		PairID:          defaultPairID,
		SharesOwned:     sdk.NewInt(10),
		CenterTickIndex: 0,
		LowerTickIndex:  -1,
		UpperTickIndex:  1,
		Fee:             1,
	},
		depositList[0],
	)
	s.Assert().Equal(&types.DepositRecord{
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
