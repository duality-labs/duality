package keeper_test

import (
	"time"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ = suite.TestingSuite(nil)

func (suite *KeeperTestSuite) TestLockLifecycle() {
	addr0 := suite.SetupAddr(0)

	// setup dex deposit and lock of those shares
	lock := suite.SetupDepositAndLock(depositSpec{
		addr:   addr0,
		token0: sdk.NewInt64Coin("TokenA", 10),
		token1: sdk.NewInt64Coin("TokenB", 10),
		tick:   0,
		fee:    1,
	})

	retrievedLock, err := suite.App.IncentivesKeeper.GetLockByID(suite.Ctx, lock.ID)
	suite.Require().NoError(err)
	suite.Require().NotNil(retrievedLock)

	// unlock the full amount
	suite.App.IncentivesKeeper.BeginUnlock(suite.Ctx, lock.ID, sdk.Coins{})

	// immediately run an end-blocker withdrawal; should do nothing
	suite.App.IncentivesKeeper.WithdrawAllMaturedLocks(suite.Ctx)

	// advance time to epoch at or after the lock end time
	suite.Ctx = suite.Ctx.WithBlockTime(suite.Ctx.BlockTime().Add(24 * time.Hour))

	// run again an end-blocker withdrawal; this time should work
	suite.App.IncentivesKeeper.WithdrawAllMaturedLocks(suite.Ctx)
	balances := suite.App.BankKeeper.GetAllBalances(suite.Ctx, addr0)
	suite.Require().Equal(sdk.NewCoins(sdk.NewInt64Coin(suite.LPDenom, 20)), balances)
	_, err = suite.App.IncentivesKeeper.GetLockByID(suite.Ctx, lock.ID)
	// should be deleted
	suite.Require().Error(err)

	// advance time again to assert that nothing changes
	suite.Ctx = suite.Ctx.WithBlockTime(suite.Ctx.BlockTime().Add(24 * time.Hour))
	suite.App.IncentivesKeeper.WithdrawAllMaturedLocks(suite.Ctx)
	suite.Require().Equal(sdk.NewCoins(sdk.NewInt64Coin(suite.LPDenom, 20)), balances)
	// fin.
}

func (suite *KeeperTestSuite) TestLockUnlockPartial() {
	addr0 := suite.SetupAddr(0)

	// setup dex deposit and lock of those shares
	lock := suite.SetupDepositAndLock(depositSpec{
		addr:   addr0,
		token0: sdk.NewInt64Coin("TokenA", 10),
		token1: sdk.NewInt64Coin("TokenB", 10),
		tick:   0,
		fee:    1,
	})

	retrievedLock, err := suite.App.IncentivesKeeper.GetLockByID(suite.Ctx, lock.ID)
	suite.Require().NoError(err)
	suite.Require().NotNil(retrievedLock)

	// unlock the full amount
	unlockId, err := suite.App.IncentivesKeeper.BeginUnlock(suite.Ctx, lock.ID, sdk.Coins{sdk.NewInt64Coin(suite.LPDenom, 10)})
	suite.Require().NoError(err)

	// immediately run an end-blocker withdrawal; should do nothing
	suite.App.IncentivesKeeper.WithdrawAllMaturedLocks(suite.Ctx)

	// advance time to epoch at or after the lock end time
	suite.Ctx = suite.Ctx.WithBlockTime(suite.Ctx.BlockTime().Add(24 * time.Hour))

	// run again an end-blocker withdrawal; this time should work
	suite.App.IncentivesKeeper.WithdrawAllMaturedLocks(suite.Ctx)
	balances := suite.App.BankKeeper.GetAllBalances(suite.Ctx, addr0)
	suite.Require().Equal(sdk.NewCoins(sdk.NewInt64Coin(suite.LPDenom, 10)), balances)
	// should be deleted
	_, err = suite.App.IncentivesKeeper.GetLockByID(suite.Ctx, unlockId)
	suite.Require().Error(err)
	// should still be accessible
	retrievedLock, err = suite.App.IncentivesKeeper.GetLockByID(suite.Ctx, lock.ID)
	suite.Require().NoError(err)
	suite.Require().NotNil(retrievedLock)

	// advance time again to assert that nothing changes
	suite.Ctx = suite.Ctx.WithBlockTime(suite.Ctx.BlockTime().Add(24 * time.Hour))
	suite.App.IncentivesKeeper.WithdrawAllMaturedLocks(suite.Ctx)
	// should be deleted
	_, err = suite.App.IncentivesKeeper.GetLockByID(suite.Ctx, unlockId)
	suite.Require().Error(err)
	// should still be accessible
	retrievedLock, err = suite.App.IncentivesKeeper.GetLockByID(suite.Ctx, lock.ID)
	suite.Require().NoError(err)
	suite.Require().NotNil(retrievedLock)
	// fin.
}
