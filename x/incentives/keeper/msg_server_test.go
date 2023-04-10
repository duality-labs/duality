package keeper_test

import (
	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ = suite.TestingSuite(nil)

var seventyTokens = sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(70000000)))
var tenTokens = sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10000000)))

// func (suite *KeeperTestSuite) TestCreateGauge_Fee() {
// 	tests := []struct {
// 		name                 string
// 		accountBalanceToFund sdk.Coins
// 		gaugeAddition        sdk.Coins
// 		expectedEndBalance   sdk.Coins
// 		isPerpetual          bool
// 		expectErr            bool
// 	}{
// 		{
// 			name:                 "user creates a non-perpetual gauge and fills gauge with all remaining tokens",
// 			accountBalanceToFund: sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(60000000))),
// 			gaugeAddition:        tenTokens,
// 		},
// 		{
// 			name:                 "user creates a non-perpetual gauge and fills gauge with some remaining tokens",
// 			accountBalanceToFund: sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(70000000))),
// 			gaugeAddition:        tenTokens,
// 		},
// 		{
// 			name:                 "user with multiple denoms creates a non-perpetual gauge and fills gauge with some remaining tokens",
// 			accountBalanceToFund: sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(70000000)), sdk.NewCoin("foo", sdk.NewInt(70000000))),
// 			gaugeAddition:        tenTokens,
// 		},
// 		{
// 			name:                 "module account creates a perpetual gauge and fills gauge with some remaining tokens",
// 			accountBalanceToFund: sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(70000000)), sdk.NewCoin("foo", sdk.NewInt(70000000))),
// 			gaugeAddition:        tenTokens,
// 			isPerpetual:          true,
// 		},
// 		{
// 			name:                 "user with multiple denoms creates a perpetual gauge and fills gauge with some remaining tokens",
// 			accountBalanceToFund: sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(70000000)), sdk.NewCoin("foo", sdk.NewInt(70000000))),
// 			gaugeAddition:        tenTokens,
// 			isPerpetual:          true,
// 		},
// 		{
// 			name:                 "user tries to create a non-perpetual gauge but does not have enough funds to pay for the create gauge fee",
// 			accountBalanceToFund: sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(40000000))),
// 			gaugeAddition:        tenTokens,
// 			expectErr:            true,
// 		},
// 		{
// 			name:                 "user tries to create a non-perpetual gauge but does not have the correct fee denom",
// 			accountBalanceToFund: sdk.NewCoins(sdk.NewCoin("foo", sdk.NewInt(60000000))),
// 			gaugeAddition:        sdk.NewCoins(sdk.NewCoin("foo", sdk.NewInt(10000000))),
// 			expectErr:            true,
// 		},
// 		{
// 			name:                 "one user tries to create a gauge, has enough funds to pay for the create gauge fee but not enough to fill the gauge",
// 			accountBalanceToFund: sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(60000000))),
// 			gaugeAddition:        sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(30000000))),
// 			expectErr:            true,
// 		},
// 	}

// 	for _, tc := range tests {
// 		suite.T().Run(tc.name, func(t *testing.T) {
// 			suite.SetupTest()

// 			ctx := suite.Ctx
// 			msgServer := keeper.NewMsgServerImpl(&suite.App.IncentivesKeeper)

// 			suite.SetupLock(suite.SetupAddr(0), , defaultLPTokens, defaultLockDuration)
// 			distrTo := types.QueryCondition{
// 				// TODO
// 				// LockQueryType: types.ByDuration,
// 				// Denom:         defaultLPDenom,
// 				// Duration:      defaultLockDuration,
// 			}

// 			msg := &types.MsgCreateGauge{
// 				IsPerpetual:       tc.isPerpetual,
// 				Owner:             testAccountAddress.String(),
// 				DistributeTo:      distrTo,
// 				Coins:             tc.gaugeAddition,
// 				StartTime:         time.Now(),
// 				NumEpochsPaidOver: 1,
// 			}
// 			// System under test.
// 			_, err := msgServer.CreateGauge(sdk.WrapSDKContext(ctx), msg)

// 			if tc.expectErr {
// 				suite.Require().Error(err)
// 			} else {
// 				suite.Require().NoError(err)
// 			}

// 			balanceAmount := bankKeeper.GetAllBalances(ctx, testAccountAddress)

// 			if tc.expectErr {
// 				suite.Require().Equal(tc.accountBalanceToFund.String(), balanceAmount.String(), "test: %v", tc.name)
// 			} else {
// 				fee := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, types.CreateGaugeFee))
// 				accountBalance := tc.accountBalanceToFund.Sub(tc.gaugeAddition)
// 				finalAccountBalance := accountBalance.Sub(fee)
// 				suite.Require().Equal(finalAccountBalance.String(), balanceAmount.String(), "test: %v", tc.name)
// 			}
// 		})
// 	}
// }

// func (suite *KeeperTestSuite) TestAddToGauge_Fee() {
// 	tests := []struct {
// 		name                 string
// 		accountBalanceToFund sdk.Coins
// 		gaugeAddition        sdk.Coins
// 		nonexistentGauge     bool
// 		isPerpetual          bool
// 		isModuleAccount      bool
// 		isGaugeComplete      bool
// 		expectErr            bool
// 	}{
// 		{
// 			name:                 "user creates a non-perpetual gauge and fills gauge with all remaining tokens",
// 			accountBalanceToFund: sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(35000000))),
// 			gaugeAddition:        tenTokens,
// 		},
// 		{
// 			name:                 "user creates a non-perpetual gauge and fills gauge with some remaining tokens",
// 			accountBalanceToFund: seventyTokens,
// 			gaugeAddition:        tenTokens,
// 		},
// 		{
// 			name:                 "user with multiple denoms creates a non-perpetual gauge and fills gauge with some remaining tokens",
// 			accountBalanceToFund: sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(70000000)), sdk.NewCoin("foo", sdk.NewInt(70000000))),
// 			gaugeAddition:        tenTokens,
// 		},
// 		{
// 			name:                 "module account creates a perpetual gauge and fills gauge with some remaining tokens",
// 			accountBalanceToFund: sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(70000000)), sdk.NewCoin("foo", sdk.NewInt(70000000))),
// 			gaugeAddition:        tenTokens,
// 			isPerpetual:          true,
// 			isModuleAccount:      true,
// 		},
// 		{
// 			name:                 "user with multiple denoms creates a perpetual gauge and fills gauge with some remaining tokens",
// 			accountBalanceToFund: sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(70000000)), sdk.NewCoin("foo", sdk.NewInt(70000000))),
// 			gaugeAddition:        tenTokens,
// 			isPerpetual:          true,
// 		},
// 		{
// 			name:                 "user tries to create a non-perpetual gauge but does not have enough funds to pay for the create gauge fee",
// 			accountBalanceToFund: sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(20000000))),
// 			gaugeAddition:        tenTokens,
// 			expectErr:            true,
// 		},
// 		{
// 			name:                 "user tries to add to a non-perpetual gauge but does not have the correct fee denom",
// 			accountBalanceToFund: sdk.NewCoins(sdk.NewCoin("foo", sdk.NewInt(60000000))),
// 			gaugeAddition:        sdk.NewCoins(sdk.NewCoin("foo", sdk.NewInt(10000000))),
// 			expectErr:            true,
// 		},
// 		{
// 			name:                 "user tries to add to a finished gauge",
// 			accountBalanceToFund: seventyTokens,
// 			gaugeAddition:        tenTokens,
// 			isGaugeComplete:      true,
// 			expectErr:            true,
// 		},
// 	}

// 	for _, tc := range tests {
// 		suite.SetupTest()

// 		testAccountPubkey := secp256k1.GenPrivKeyFromSecret([]byte("acc")).PubKey()
// 		testAccountAddress := sdk.AccAddress(testAccountPubkey.Address())

// 		bankKeeper := suite.App.BankKeeper
// 		incentivesKeeper := suite.App.IncentivesKeeper
// 		accountKeeper := suite.App.AccountKeeper
// 		msgServer := keeper.NewMsgServerImpl(&incentivesKeeper)

// 		suite.FundAcc(testAccountAddress, tc.accountBalanceToFund)

// 		if tc.isModuleAccount {
// 			modAcc := authtypes.NewModuleAccount(authtypes.NewBaseAccount(testAccountAddress, testAccountPubkey, 1, 0),
// 				"module",
// 				"permission",
// 			)
// 			accountKeeper.SetModuleAccount(suite.Ctx, modAcc)
// 		}

// 		// System under test.
// 		coins := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(500000000)))
// 		gaugeID, gauge, _, _ := suite.SetupGauge(tc.isPerpetual, coins)
// 		if tc.nonexistentGauge {
// 			gaugeID = incentivesKeeper.GetLastGaugeID(suite.Ctx) + 1
// 		}
// 		// simulate times to complete the gauge.
// 		if tc.isGaugeComplete {
// 			suite.completeGauge(gauge, sdk.AccAddress([]byte("a___________________")))
// 		}
// 		msg := &types.MsgAddToGauge{
// 			Owner:   testAccountAddress.String(),
// 			GaugeId: gaugeID,
// 			Rewards: tc.gaugeAddition,
// 		}

// 		_, err := msgServer.AddToGauge(sdk.WrapSDKContext(suite.Ctx), msg)

// 		if tc.expectErr {
// 			suite.Require().Error(err, tc.name)
// 		} else {
// 			suite.Require().NoError(err, tc.name)
// 		}

// 		bal := bankKeeper.GetAllBalances(suite.Ctx, testAccountAddress)

// 		if !tc.expectErr {
// 			fee := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, types.AddToGaugeFee))
// 			accountBalance := tc.accountBalanceToFund.Sub(tc.gaugeAddition)
// 			finalAccountBalance := accountBalance.Sub(fee)
// 			suite.Require().Equal(finalAccountBalance.String(), bal.String(), "test: %v", tc.name)
// 		} else if tc.expectErr && !tc.isGaugeComplete {
// 			suite.Require().Equal(tc.accountBalanceToFund.String(), bal.String(), "test: %v", tc.name)
// 		}
// 	}
// }

// func (suite *KeeperTestSuite) completeGauge(gauge *types.Gauge, sendingAddress sdk.AccAddress) {
// 	// TODO
// 	lockCoins := sdk.NewCoin(gauge.DistributeTo.Denom, sdk.NewInt(1000))
// 	suite.FundAcc(sendingAddress, sdk.NewCoins(lockCoins))
// 	suite.SetupLock(sendingAddress, sdk.NewCoins(lockCoins), gauge.DistributeTo.Duration)
// 	epochId := suite.App.IncentivesKeeper.GetEpochInfo(suite.Ctx).Identifier
// 	if suite.Ctx.BlockTime().Before(gauge.StartTime) {
// 		suite.Ctx = suite.Ctx.WithBlockTime(gauge.StartTime.Add(time.Hour))
// 	}
// 	suite.BeginNewBlock(false)
// 	for i := 0; i < int(gauge.NumEpochsPaidOver); i++ {
// 		suite.App.IncentivesKeeper.BeforeEpochStart(suite.Ctx, epochId, int64(i))
// 		suite.App.IncentivesKeeper.AfterEpochEnd(suite.Ctx, epochId, int64(i))
// 	}
// 	suite.BeginNewBlock(false)
// 	gauge2, err := suite.App.IncentivesKeeper.GetGaugeByID(suite.Ctx, gauge.Id)
// 	suite.Require().NoError(err)
// 	suite.Require().True(gauge2.IsFinishedGauge(suite.Ctx.BlockTime()))
// }

// func (suite *KeeperTestSuite) TestMsgSetupLock() {
// 	type param struct {
// 		coinsToLock         sdk.Coins
// 		lockOwner           sdk.AccAddress
// 		duration            time.Duration
// 		coinsInOwnerAddress sdk.Coins
// 	}

// 	tests := []struct {
// 		name       string
// 		param      param
// 		expectPass bool
// 	}{
// 		{
// 			name: "creation of lock via lockTokens",
// 			param: param{
// 				coinsToLock:         sdk.Coins{sdk.NewInt64Coin("stake", 10)},       // setup wallet
// 				lockOwner:           sdk.AccAddress([]byte("addr1---------------")), // setup wallet
// 				duration:            time.Second,
// 				coinsInOwnerAddress: sdk.Coins{sdk.NewInt64Coin("stake", 10)},
// 			},
// 			expectPass: true,
// 		},
// 		{
// 			name: "locking more coins than are in the address",
// 			param: param{
// 				coinsToLock:         sdk.Coins{sdk.NewInt64Coin("stake", 20)},       // setup wallet
// 				lockOwner:           sdk.AccAddress([]byte("addr1---------------")), // setup wallet
// 				duration:            time.Second,
// 				coinsInOwnerAddress: sdk.Coins{sdk.NewInt64Coin("stake", 10)},
// 			},
// 			expectPass: false,
// 		},
// 	}

// 	for _, test := range tests {
// 		suite.SetupTest()

// 		suite.FundAcc(test.param.lockOwner, test.param.coinsInOwnerAddress)

// 		msgServer := keeper.NewMsgServerImpl(&suite.App.IncentivesKeeper)
// 		c := sdk.WrapSDKContext(suite.Ctx)
// 		_, err := msgServer.SetupLock(c, types.NewMsgSetupLock(test.param.lockOwner, test.param.duration, test.param.coinsToLock))

// 		if test.expectPass {
// 			// creation of lock via LockTokens
// 			msgServer := keeper.NewMsgServerImpl(&suite.App.IncentivesKeeper)
// 			_, err = msgServer.SetupLock(sdk.WrapSDKContext(suite.Ctx), types.NewMsgSetupLock(test.param.lockOwner, test.param.duration, test.param.coinsToLock))

// 			// Check Locks
// 			locks, err := suite.App.IncentivesKeeper.GetLocks(suite.Ctx)
// 			suite.Require().NoError(err)
// 			suite.Require().Len(locks, 1)
// 			suite.Require().Equal(locks[0].Coins, test.param.coinsToLock)

// 			// check accumulation store is correctly updated
// 			// accum := suite.App.IncentivesKeeper.GetLocksAccumulation(suite.Ctx, types.QueryCondition{
// 			// TODO
// 			// LockQueryType: types.ByDuration,
// 			// Denom:         "stake",
// 			// Duration:      test.param.duration,
// 			// })
// 			// suite.Require().Equal(accum.String(), "10")

// 			// add more tokens to lock via LockTokens
// 			suite.FundAcc(test.param.lockOwner, test.param.coinsInOwnerAddress)

// 			_, err = msgServer.SetupLock(sdk.WrapSDKContext(suite.Ctx), types.NewMsgSetupLock(test.param.lockOwner, locks[0].Duration, test.param.coinsToLock))
// 			suite.Require().NoError(err)

// 			// check locks after adding tokens to lock
// 			locks, err = suite.App.IncentivesKeeper.GetLocks(suite.Ctx)
// 			suite.Require().NoError(err)
// 			suite.Require().Len(locks, 1)
// 			suite.Require().Equal(locks[0].Coins, test.param.coinsToLock.Add(test.param.coinsToLock...))

// 			// check accumulation store is correctly updated
// 			// accum = suite.App.IncentivesKeeper.GetLocksAccumulation(suite.Ctx, types.QueryCondition{
// 			// TODO
// 			// LockQueryType: types.ByDuration,
// 			// Denom:         "stake",
// 			// Duration:      test.param.duration,
// 			// })
// 			// suite.Require().Equal(accum.String(), "20")

// 		} else {
// 			// Fail simple lock token
// 			suite.Require().Error(err)
// 		}
// 	}
// }

// func (suite *KeeperTestSuite) TestMsgBeginUnlocking() {
// 	type param struct {
// 		coinsToLock         sdk.Coins
// 		coinsToUnlock       sdk.Coins
// 		lockOwner           sdk.AccAddress
// 		duration            time.Duration
// 		coinsInOwnerAddress sdk.Coins
// 	}

// 	tests := []struct {
// 		name       string
// 		param      param
// 		expectPass bool
// 		isPartial  bool
// 	}{
// 		{
// 			name: "unlock full amount of tokens via begin unlock",
// 			param: param{
// 				coinsToLock:         sdk.Coins{sdk.NewInt64Coin("stake", 10)},       // setup wallet
// 				coinsToUnlock:       sdk.Coins{sdk.NewInt64Coin("stake", 10)},       // setup wallet
// 				lockOwner:           sdk.AccAddress([]byte("addr1---------------")), // setup wallet
// 				duration:            time.Second,
// 				coinsInOwnerAddress: sdk.Coins{sdk.NewInt64Coin("stake", 10)},
// 			},
// 			expectPass: true,
// 		},
// 		{
// 			name: "unlock partial amount of tokens via begin unlock",
// 			param: param{
// 				coinsToLock:         sdk.Coins{sdk.NewInt64Coin("stake", 10)},       // setup wallet
// 				coinsToUnlock:       sdk.Coins{sdk.NewInt64Coin("stake", 5)},        // setup wallet
// 				lockOwner:           sdk.AccAddress([]byte("addr1---------------")), // setup wallet
// 				duration:            time.Second,
// 				coinsInOwnerAddress: sdk.Coins{sdk.NewInt64Coin("stake", 10)},
// 			},
// 			isPartial:  true,
// 			expectPass: true,
// 		},
// 		{
// 			name: "unlock zero amount of tokens via begin unlock",
// 			param: param{
// 				coinsToLock:         sdk.Coins{sdk.NewInt64Coin("stake", 10)},       // setup wallet
// 				coinsToUnlock:       sdk.Coins{},                                    // setup wallet
// 				lockOwner:           sdk.AccAddress([]byte("addr1---------------")), // setup wallet
// 				duration:            time.Second,
// 				coinsInOwnerAddress: sdk.Coins{sdk.NewInt64Coin("stake", 10)},
// 			},
// 			expectPass: true,
// 		},
// 	}

// 	for _, test := range tests {
// 		suite.SetupTest()

// 		suite.FundAcc(test.param.lockOwner, test.param.coinsInOwnerAddress)

// 		msgServer := keeper.NewMsgServerImpl(&suite.App.IncentivesKeeper)
// 		goCtx := sdk.WrapSDKContext(suite.Ctx)
// 		resp, err := msgServer.SetupLock(goCtx, types.NewMsgSetupLock(test.param.lockOwner, test.param.duration, test.param.coinsToLock))
// 		suite.Require().NoError(err)

// 		unlockingResponse, err := msgServer.BeginUnlocking(goCtx, types.NewMsgBeginUnlocking(test.param.lockOwner, resp.ID, test.param.coinsToUnlock))

// 		if test.expectPass {
// 			suite.Require().NoError(err)
// 			suite.AssertEventEmitted(suite.Ctx, types.TypeEvtBeginUnlock, 1)
// 			suite.Require().True(unlockingResponse.Success)
// 			if test.isPartial {
// 				suite.Require().Equal(unlockingResponse.UnlockingLockID, resp.ID+1)
// 			} else {
// 				suite.Require().Equal(unlockingResponse.UnlockingLockID, resp.ID)
// 			}
// 		} else {
// 			suite.Require().Error(err)
// 			suite.AssertEventEmitted(suite.Ctx, types.TypeEvtBeginUnlock, 0)
// 		}
// 	}
// }

// func (suite *KeeperTestSuite) TestMsgBeginUnlockingAll() {
// 	type param struct {
// 		coinsToLock         sdk.Coins
// 		lockOwner           sdk.AccAddress
// 		duration            time.Duration
// 		coinsInOwnerAddress sdk.Coins
// 	}

// 	tests := []struct {
// 		name       string
// 		param      param
// 		expectPass bool
// 	}{
// 		{
// 			name: "unlock all lockups",
// 			param: param{
// 				coinsToLock:         sdk.Coins{sdk.NewInt64Coin("stake", 10)},       // setup wallet
// 				lockOwner:           sdk.AccAddress([]byte("addr1---------------")), // setup wallet
// 				duration:            time.Second,
// 				coinsInOwnerAddress: sdk.Coins{sdk.NewInt64Coin("stake", 10)},
// 			},
// 			expectPass: true,
// 		},
// 	}

// 	for _, test := range tests {
// 		suite.SetupTest()

// 		suite.FundAcc(test.param.lockOwner, test.param.coinsInOwnerAddress)

// 		msgServer := keeper.NewMsgServerImpl(&suite.App.IncentivesKeeper)
// 		c := sdk.WrapSDKContext(suite.Ctx)
// 		_, err := msgServer.SetupLock(c, types.NewMsgSetupLock(test.param.lockOwner, test.param.duration, test.param.coinsToLock))
// 		suite.Require().NoError(err)

// 		_, err = msgServer.BeginUnlockingAll(c, types.NewMsgBeginUnlockingAll(test.param.lockOwner))

// 		if test.expectPass {
// 			suite.Require().NoError(err)
// 		} else {
// 			suite.Require().Error(err)
// 		}
// 	}
// }

// func (suite *KeeperTestSuite) TestMsgEditLockup() {
// 	type param struct {
// 		coinsToLock sdk.Coins
// 		lockOwner   sdk.AccAddress
// 		duration    time.Duration
// 		newDuration time.Duration
// 	}

// 	tests := []struct {
// 		name       string
// 		param      param
// 		expectPass bool
// 	}{
// 		{
// 			name: "edit lockups by duration",
// 			param: param{
// 				coinsToLock: sdk.Coins{sdk.NewInt64Coin("stake", 10)},       // setup wallet
// 				lockOwner:   sdk.AccAddress([]byte("addr1---------------")), // setup wallet
// 				duration:    time.Second,
// 				newDuration: time.Second * 2,
// 			},
// 			expectPass: true,
// 		},
// 		{
// 			name: "edit lockups by lesser duration",
// 			param: param{
// 				coinsToLock: sdk.Coins{sdk.NewInt64Coin("stake", 10)},       // setup wallet
// 				lockOwner:   sdk.AccAddress([]byte("addr1---------------")), // setup wallet
// 				duration:    time.Second,
// 				newDuration: time.Second / 2,
// 			},
// 			expectPass: false,
// 		},
// 	}

// 	for _, test := range tests {
// 		suite.SetupTest()

// 		err := dualityapp.FundAccount(suite.App.BankKeeper, suite.Ctx, test.param.lockOwner, test.param.coinsToLock)
// 		suite.Require().NoError(err)

// 		msgServer := keeper.NewMsgServerImpl(&suite.App.IncentivesKeeper)
// 		c := sdk.WrapSDKContext(suite.Ctx)
// 		resp, err := msgServer.SetupLock(c, types.NewMsgSetupLock(test.param.lockOwner, test.param.duration, test.param.coinsToLock))
// 		suite.Require().NoError(err)

// 		_, err = msgServer.ExtendLockup(c, types.NewMsgExtendLockup(test.param.lockOwner, resp.ID, test.param.newDuration))

// 		if test.expectPass {
// 			suite.Require().NoError(err, test.name)
// 		} else {
// 			suite.Require().Error(err, test.name)
// 		}
// 	}
// }
