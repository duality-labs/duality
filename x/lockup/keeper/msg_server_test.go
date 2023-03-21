package keeper_test

import (
	"time"

	dualityapp "github.com/duality-labs/duality/app"
	"github.com/duality-labs/duality/x/lockup/keeper"
	"github.com/duality-labs/duality/x/lockup/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestMsgLockTokens() {
	type param struct {
		coinsToLock         sdk.Coins
		lockOwner           sdk.AccAddress
		duration            time.Duration
		coinsInOwnerAddress sdk.Coins
	}

	tests := []struct {
		name       string
		param      param
		expectPass bool
	}{
		{
			name: "creation of lock via lockTokens",
			param: param{
				coinsToLock:         sdk.Coins{sdk.NewInt64Coin("stake", 10)},       // setup wallet
				lockOwner:           sdk.AccAddress([]byte("addr1---------------")), // setup wallet
				duration:            time.Second,
				coinsInOwnerAddress: sdk.Coins{sdk.NewInt64Coin("stake", 10)},
			},
			expectPass: true,
		},
		{
			name: "locking more coins than are in the address",
			param: param{
				coinsToLock:         sdk.Coins{sdk.NewInt64Coin("stake", 20)},       // setup wallet
				lockOwner:           sdk.AccAddress([]byte("addr1---------------")), // setup wallet
				duration:            time.Second,
				coinsInOwnerAddress: sdk.Coins{sdk.NewInt64Coin("stake", 10)},
			},
			expectPass: false,
		},
	}

	for _, test := range tests {
		suite.SetupTest()

		suite.FundAcc(test.param.lockOwner, test.param.coinsInOwnerAddress)

		msgServer := keeper.NewMsgServerImpl(&suite.App.LockupKeeper)
		c := sdk.WrapSDKContext(suite.Ctx)
		_, err := msgServer.LockTokens(c, types.NewMsgLockTokens(test.param.lockOwner, test.param.duration, test.param.coinsToLock))

		if test.expectPass {
			// creation of lock via LockTokens
			msgServer := keeper.NewMsgServerImpl(&suite.App.LockupKeeper)
			_, err = msgServer.LockTokens(sdk.WrapSDKContext(suite.Ctx), types.NewMsgLockTokens(test.param.lockOwner, test.param.duration, test.param.coinsToLock))

			// Check Locks
			locks, err := suite.App.LockupKeeper.GetPeriodLocks(suite.Ctx)
			suite.Require().NoError(err)
			suite.Require().Len(locks, 1)
			suite.Require().Equal(locks[0].Coins, test.param.coinsToLock)

			// check accumulation store is correctly updated
			accum := suite.App.LockupKeeper.GetPeriodLocksAccumulation(suite.Ctx, types.QueryCondition{
				LockQueryType: types.ByDuration,
				Denom:         "stake",
				Duration:      test.param.duration,
			})
			suite.Require().Equal(accum.String(), "10")

			// add more tokens to lock via LockTokens
			suite.FundAcc(test.param.lockOwner, test.param.coinsInOwnerAddress)

			_, err = msgServer.LockTokens(sdk.WrapSDKContext(suite.Ctx), types.NewMsgLockTokens(test.param.lockOwner, locks[0].Duration, test.param.coinsToLock))
			suite.Require().NoError(err)

			// check locks after adding tokens to lock
			locks, err = suite.App.LockupKeeper.GetPeriodLocks(suite.Ctx)
			suite.Require().NoError(err)
			suite.Require().Len(locks, 1)
			suite.Require().Equal(locks[0].Coins, test.param.coinsToLock.Add(test.param.coinsToLock...))

			// check accumulation store is correctly updated
			accum = suite.App.LockupKeeper.GetPeriodLocksAccumulation(suite.Ctx, types.QueryCondition{
				LockQueryType: types.ByDuration,
				Denom:         "stake",
				Duration:      test.param.duration,
			})
			suite.Require().Equal(accum.String(), "20")

		} else {
			// Fail simple lock token
			suite.Require().Error(err)
		}
	}
}

func (suite *KeeperTestSuite) TestMsgBeginUnlocking() {
	type param struct {
		coinsToLock         sdk.Coins
		coinsToUnlock       sdk.Coins
		lockOwner           sdk.AccAddress
		duration            time.Duration
		coinsInOwnerAddress sdk.Coins
	}

	tests := []struct {
		name       string
		param      param
		expectPass bool
		isPartial  bool
	}{
		{
			name: "unlock full amount of tokens via begin unlock",
			param: param{
				coinsToLock:         sdk.Coins{sdk.NewInt64Coin("stake", 10)},       // setup wallet
				coinsToUnlock:       sdk.Coins{sdk.NewInt64Coin("stake", 10)},       // setup wallet
				lockOwner:           sdk.AccAddress([]byte("addr1---------------")), // setup wallet
				duration:            time.Second,
				coinsInOwnerAddress: sdk.Coins{sdk.NewInt64Coin("stake", 10)},
			},
			expectPass: true,
		},
		{
			name: "unlock partial amount of tokens via begin unlock",
			param: param{
				coinsToLock:         sdk.Coins{sdk.NewInt64Coin("stake", 10)},       // setup wallet
				coinsToUnlock:       sdk.Coins{sdk.NewInt64Coin("stake", 5)},        // setup wallet
				lockOwner:           sdk.AccAddress([]byte("addr1---------------")), // setup wallet
				duration:            time.Second,
				coinsInOwnerAddress: sdk.Coins{sdk.NewInt64Coin("stake", 10)},
			},
			isPartial:  true,
			expectPass: true,
		},
		{
			name: "unlock zero amount of tokens via begin unlock",
			param: param{
				coinsToLock:         sdk.Coins{sdk.NewInt64Coin("stake", 10)},       // setup wallet
				coinsToUnlock:       sdk.Coins{},                                    // setup wallet
				lockOwner:           sdk.AccAddress([]byte("addr1---------------")), // setup wallet
				duration:            time.Second,
				coinsInOwnerAddress: sdk.Coins{sdk.NewInt64Coin("stake", 10)},
			},
			expectPass: true,
		},
	}

	for _, test := range tests {
		suite.SetupTest()

		suite.FundAcc(test.param.lockOwner, test.param.coinsInOwnerAddress)

		msgServer := keeper.NewMsgServerImpl(&suite.App.LockupKeeper)
		goCtx := sdk.WrapSDKContext(suite.Ctx)
		resp, err := msgServer.LockTokens(goCtx, types.NewMsgLockTokens(test.param.lockOwner, test.param.duration, test.param.coinsToLock))
		suite.Require().NoError(err)

		unlockingResponse, err := msgServer.BeginUnlocking(goCtx, types.NewMsgBeginUnlocking(test.param.lockOwner, resp.ID, test.param.coinsToUnlock))

		if test.expectPass {
			suite.Require().NoError(err)
			suite.AssertEventEmitted(suite.Ctx, types.TypeEvtBeginUnlock, 1)
			suite.Require().True(unlockingResponse.Success)
			if test.isPartial {
				suite.Require().Equal(unlockingResponse.UnlockingLockID, resp.ID+1)
			} else {
				suite.Require().Equal(unlockingResponse.UnlockingLockID, resp.ID)
			}
		} else {
			suite.Require().Error(err)
			suite.AssertEventEmitted(suite.Ctx, types.TypeEvtBeginUnlock, 0)
		}
	}
}

func (suite *KeeperTestSuite) TestMsgBeginUnlockingAll() {
	type param struct {
		coinsToLock         sdk.Coins
		lockOwner           sdk.AccAddress
		duration            time.Duration
		coinsInOwnerAddress sdk.Coins
	}

	tests := []struct {
		name       string
		param      param
		expectPass bool
	}{
		{
			name: "unlock all lockups",
			param: param{
				coinsToLock:         sdk.Coins{sdk.NewInt64Coin("stake", 10)},       // setup wallet
				lockOwner:           sdk.AccAddress([]byte("addr1---------------")), // setup wallet
				duration:            time.Second,
				coinsInOwnerAddress: sdk.Coins{sdk.NewInt64Coin("stake", 10)},
			},
			expectPass: true,
		},
	}

	for _, test := range tests {
		suite.SetupTest()

		suite.FundAcc(test.param.lockOwner, test.param.coinsInOwnerAddress)

		msgServer := keeper.NewMsgServerImpl(&suite.App.LockupKeeper)
		c := sdk.WrapSDKContext(suite.Ctx)
		_, err := msgServer.LockTokens(c, types.NewMsgLockTokens(test.param.lockOwner, test.param.duration, test.param.coinsToLock))
		suite.Require().NoError(err)

		_, err = msgServer.BeginUnlockingAll(c, types.NewMsgBeginUnlockingAll(test.param.lockOwner))

		if test.expectPass {
			suite.Require().NoError(err)
		} else {
			suite.Require().Error(err)
		}
	}
}

func (suite *KeeperTestSuite) TestMsgEditLockup() {
	type param struct {
		coinsToLock sdk.Coins
		lockOwner   sdk.AccAddress
		duration    time.Duration
		newDuration time.Duration
	}

	tests := []struct {
		name       string
		param      param
		expectPass bool
	}{
		{
			name: "edit lockups by duration",
			param: param{
				coinsToLock: sdk.Coins{sdk.NewInt64Coin("stake", 10)},       // setup wallet
				lockOwner:   sdk.AccAddress([]byte("addr1---------------")), // setup wallet
				duration:    time.Second,
				newDuration: time.Second * 2,
			},
			expectPass: true,
		},
		{
			name: "edit lockups by lesser duration",
			param: param{
				coinsToLock: sdk.Coins{sdk.NewInt64Coin("stake", 10)},       // setup wallet
				lockOwner:   sdk.AccAddress([]byte("addr1---------------")), // setup wallet
				duration:    time.Second,
				newDuration: time.Second / 2,
			},
			expectPass: false,
		},
	}

	for _, test := range tests {
		suite.SetupTest()

		err := dualityapp.FundAccount(suite.App.BankKeeper, suite.Ctx, test.param.lockOwner, test.param.coinsToLock)
		suite.Require().NoError(err)

		msgServer := keeper.NewMsgServerImpl(&suite.App.LockupKeeper)
		c := sdk.WrapSDKContext(suite.Ctx)
		resp, err := msgServer.LockTokens(c, types.NewMsgLockTokens(test.param.lockOwner, test.param.duration, test.param.coinsToLock))
		suite.Require().NoError(err)

		_, err = msgServer.ExtendLockup(c, types.NewMsgExtendLockup(test.param.lockOwner, resp.ID, test.param.newDuration))

		if test.expectPass {
			suite.Require().NoError(err, test.name)
		} else {
			suite.Require().Error(err, test.name)
		}
	}
}
