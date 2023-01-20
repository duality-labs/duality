package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
)

func (s *MsgServerTestSuite) TestWithdrawMultiFailure() {
	s.fundAliceBalances(50, 50)
	// GIVEN
	// alice deposits 5 A, 5 B in tick 0 fee 0
	s.aliceDeposits(NewDeposit(5, 5, 0, 0))
	s.assertAliceShares(0, 0, 10)
	s.assertLiquidityAtTickInt(sdk.NewInt(5), sdk.NewInt(5), 0, 0)
	s.assertAliceBalances(45, 45)
	s.assertDexBalances(5, 5)

	// WHEN
	// alice withdraws 6 shares, then 10 shares
	// THEN
	// failure on second withdraw (insufficient shares) will trigger ErrNotEnoughShares
	err := types.ErrInsufficientShares
	s.aliceWithdrawFails(err,
		NewWithdrawl(6, 0, 0),
		NewWithdrawl(10, 0, 0),
	)
}

func (s *MsgServerTestSuite) TestWithdrawMultiSuccess() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// alice deposits 5 A, 5 B in tick 0 fee 0
	s.aliceDeposits(NewDeposit(5, 5, 0, 0))
	s.assertAliceShares(0, 0, 10)
	s.assertLiquidityAtTickInt(sdk.NewInt(5), sdk.NewInt(5), 0, 0)
	s.assertAliceBalances(45, 45)
	s.assertDexBalances(5, 5)

	// WHEN
	// alice withdraws 6 shares, then 4 shares
	s.aliceWithdraws(
		NewWithdrawl(6, 0, 0),
		NewWithdrawl(4, 0, 0),
	)

	// THEN
	// both withdraws should work
	// i.e. no shares remaining and entire balance transferred out
	s.assertAliceShares(0, 0, 0)
	s.assertLiquidityAtTickInt(sdk.ZeroInt(), sdk.ZeroInt(), 0, 0)
	s.assertAliceBalances(50, 50)
	s.assertDexBalances(0, 0)
}
