package keeper_test

import (
	"github.com/stretchr/testify/suite"
)

var _ = suite.TestingSuite(nil)

// // TestGRPCGetGaugeByID tests querying gauges via gRPC returns the correct response.
// func (suite *KeeperTestSuite) TestGRPCGetGaugeByID() {
// 	suite.SetupTest()

// 	// create a gauge
// 	gaugeID, _, coins, startTime := suite.SetupGauge(false, sdk.Coins{sdk.NewInt64Coin("stake", 10)})

// 	// ensure that querying for a gauge with an ID that doesn't exist returns an error.
// 	res, err := suite.QueryServer.GetGaugeByID(sdk.WrapSDKContext(suite.Ctx), &types.GetGaugeByIDRequest{Id: 1000})
// 	suite.Require().Error(err)
// 	suite.Require().Equal(res, (*types.GetGaugeByIDResponse)(nil))

// 	// check that querying a gauge with an ID that exists returns the gauge.
// 	res, err = suite.QueryServer.GetGaugeByID(sdk.WrapSDKContext(suite.Ctx), &types.GetGaugeByIDRequest{Id: gaugeID})
// 	suite.Require().NoError(err)
// 	suite.Require().NotEqual(res.Gauge, nil)
// 	expectedGauge := types.Gauge{
// 		Id:           gaugeID,
// 		IsPerpetual:  false,
// 		DistributeTo: types.QueryCondition{
// 			// LockQueryType: types.ByDuration,
// 			// Denom:         "lptoken",
// 			// Duration:      time.Second,
// 		},
// 		Coins:             coins,
// 		NumEpochsPaidOver: 2,
// 		FilledEpochs:      0,
// 		DistributedCoins:  sdk.Coins{},
// 		StartTime:         startTime,
// 	}
// 	suite.Require().Equal(res.Gauge.String(), expectedGauge.String())
// }

// // TestGRPCGauges tests querying upcoming and active gauges via gRPC returns the correct response.
// func (suite *KeeperTestSuite) TestGRPCGauges() {
// 	suite.SetupTest()

// 	// ensure initially querying gauges returns no gauges
// 	res, err := suite.QueryServer.types.Gauges(sdk.WrapSDKContext(suite.Ctx), &types.GetGaugesActiveUpcomingRequest{})
// 	suite.Require().NoError(err)
// 	suite.Require().Len(res.Data, 0)

// 	// create a gauge
// 	gaugeID, _, coins, startTime := suite.SetupGauge(false, sdk.Coins{sdk.NewInt64Coin("stake", 10)})

// 	// query gauges again, but this time expect the gauge created earlier in the response
// 	res, err = suite.QueryServer.types.Gauges(sdk.WrapSDKContext(suite.Ctx), &types.GetGaugesActiveUpcomingRequest{})
// 	suite.Require().NoError(err)
// 	suite.Require().Len(res.Data, 1)
// 	expectedGauge := types.Gauge{
// 		Id:           gaugeID,
// 		IsPerpetual:  false,
// 		DistributeTo: types.QueryCondition{
// 			// LockQueryType: types.ByDuration,
// 			// Denom:         "lptoken",
// 			// Duration:      time.Second,
// 		},
// 		Coins:             coins,
// 		NumEpochsPaidOver: 2,
// 		FilledEpochs:      0,
// 		DistributedCoins:  sdk.Coins{},
// 		StartTime:         startTime,
// 	}
// 	suite.Require().Equal(res.Data[0].String(), expectedGauge.String())

// 	// create 10 more gauges
// 	for i := 0; i < 10; i++ {
// 		suite.SetupGauge(false, sdk.Coins{sdk.NewInt64Coin("stake", 3)})
// 		suite.Ctx = suite.Ctx.WithBlockTime(startTime.Add(time.Second))
// 	}

// 	// check that setting page request limit to 10 will only return 10 out of the 11 gauges
// 	filter := query.PageRequest{Limit: 10}
// 	res, err = suite.QueryServer.types.Gauges(sdk.WrapSDKContext(suite.Ctx), &types.GetGaugesActiveUpcomingRequest{Pagination: &filter})
// 	suite.Require().Len(res.Data, 10)
// }

// // TestGRPCActiveGauges tests querying active gauges via gRPC returns the correct response.
// func (suite *KeeperTestSuite) TestGRPCActiveGauges() {
// 	suite.SetupTest()

// 	// ensure initially querying active gauges returns no gauges
// 	res, err := suite.QueryServer.ActiveGauges(sdk.WrapSDKContext(suite.Ctx), &types.ActiveGetGaugesActiveUpcomingRequest{})
// 	suite.Require().NoError(err)
// 	suite.Require().Len(res.Data, 0)

// 	// create a gauge and move it from upcoming to active
// 	gaugeID, gauge, coins, startTime := suite.SetupGauge(false, sdk.Coins{sdk.NewInt64Coin("stake", 10)})
// 	suite.Ctx = suite.Ctx.WithBlockTime(startTime.Add(time.Second))
// 	err = suite.QueryServer.MoveUpcomingGaugeToActiveGauge(suite.Ctx, *gauge)
// 	suite.Require().NoError(err)

// 	// query active gauges again, but this time expect the gauge created earlier in the response
// 	res, err = suite.QueryServer.ActiveGauges(sdk.WrapSDKContext(suite.Ctx), &types.ActiveGetGaugesActiveUpcomingRequest{})
// 	suite.Require().NoError(err)
// 	suite.Require().Len(res.Data, 1)
// 	expectedGauge := types.Gauge{
// 		Id:           gaugeID,
// 		IsPerpetual:  false,
// 		DistributeTo: types.QueryCondition{
// 			// LockQueryType: types.ByDuration,
// 			// Denom:         "lptoken",
// 			// Duration:      time.Second,
// 		},
// 		Coins:             coins,
// 		NumEpochsPaidOver: 2,
// 		FilledEpochs:      0,
// 		DistributedCoins:  sdk.Coins{},
// 		StartTime:         startTime,
// 	}
// 	suite.Require().Equal(res.Data[0].String(), expectedGauge.String())

// 	// create 20 more gauges
// 	for i := 0; i < 20; i++ {
// 		_, gauge, _, _ := suite.SetupGauge(false, sdk.Coins{sdk.NewInt64Coin("stake", 3)})
// 		suite.Ctx = suite.Ctx.WithBlockTime(startTime.Add(time.Second))

// 		// move the first 9 gauges from upcoming to active (now 10 active gauges, 30 total gauges)
// 		if i < 9 {
// 			suite.QueryServer.MoveUpcomingGaugeToActiveGauge(suite.Ctx, *gauge)
// 		}
// 	}

// 	// set page request limit to 5, expect only 5 active gauge responses
// 	res, err = suite.QueryServer.ActiveGauges(sdk.WrapSDKContext(suite.Ctx), &types.ActiveGetGaugesActiveUpcomingRequest{Pagination: &query.PageRequest{Limit: 5}})
// 	suite.Require().Len(res.Data, 5)

// 	// set page request limit to 15, expect only 10 active gauge responses
// 	res, err = suite.QueryServer.ActiveGauges(sdk.WrapSDKContext(suite.Ctx), &types.ActiveGetGaugesActiveUpcomingRequest{Pagination: &query.PageRequest{Limit: 15}})
// 	suite.Require().Len(res.Data, 10)
// }

// // TestGRPCActiveGaugesPerDenom tests querying active gauges by denom via gRPC returns the correct response.
// func (suite *KeeperTestSuite) TestGRPCActiveGaugesPerDenom() {
// 	suite.SetupTest()

// 	// ensure initially querying gauges by denom returns no gauges
// 	res, err := suite.QueryServer.ActiveGaugesPerDenom(sdk.WrapSDKContext(suite.Ctx), &types.ActiveGaugesPerDenomRequest{})
// 	suite.Require().NoError(err)
// 	suite.Require().Len(res.Data, 0)

// 	// create a gauge
// 	gaugeID, gauge, coins, startTime := suite.SetupGauge(false, sdk.Coins{sdk.NewInt64Coin("stake", 10)})
// 	suite.Ctx = suite.Ctx.WithBlockTime(startTime.Add(time.Second))
// 	err = suite.App.IncentivesKeeper.MoveUpcomingGaugeToActiveGauge(suite.Ctx, *gauge)

// 	// query gauges by denom again, but this time expect the gauge created earlier in the response
// 	res, err = suite.QueryServer.ActiveGaugesPerDenom(sdk.WrapSDKContext(suite.Ctx), &types.ActiveGaugesPerDenomRequest{Denom: "lptoken", Pagination: nil})
// 	suite.Require().NoError(err)
// 	suite.Require().Len(res.Data, 1)
// 	expectedGauge := types.Gauge{
// 		Id:           gaugeID,
// 		IsPerpetual:  false,
// 		DistributeTo: types.QueryCondition{
// 			// LockQueryType: types.ByDuration,
// 			// Denom:         "lptoken",
// 			// Duration:      time.Second,
// 		},
// 		Coins:             coins,
// 		NumEpochsPaidOver: 2,
// 		FilledEpochs:      0,
// 		DistributedCoins:  sdk.Coins{},
// 		StartTime:         startTime,
// 	}
// 	suite.Require().Equal(res.Data[0].String(), expectedGauge.String())

// 	// setup 20 more gauges with the pool denom
// 	for i := 0; i < 20; i++ {
// 		_, gauge, _, _ := suite.SetupNewGaugeWithDenom(false, sdk.Coins{sdk.NewInt64Coin("stake", 3)}, "pool")
// 		suite.Ctx = suite.Ctx.WithBlockTime(startTime.Add(time.Second))

// 		// move the first 10 of 20 gauges to an active status
// 		if i < 10 {
// 			suite.QueryServer.MoveUpcomingGaugeToActiveGauge(suite.Ctx, *gauge)
// 		}
// 	}

// 	// query active gauges by lptoken denom with a page request of 5 should only return one gauge
// 	res, err = suite.QueryServer.ActiveGaugesPerDenom(sdk.WrapSDKContext(suite.Ctx), &types.ActiveGaugesPerDenomRequest{Denom: "lptoken", Pagination: &query.PageRequest{Limit: 5}})
// 	suite.Require().Len(res.Data, 1)

// 	// query active gauges by pool denom with a page request of 5 should return 5 gauges
// 	res, err = suite.QueryServer.ActiveGaugesPerDenom(sdk.WrapSDKContext(suite.Ctx), &types.ActiveGaugesPerDenomRequest{Denom: "pool", Pagination: &query.PageRequest{Limit: 5}})
// 	suite.Require().Len(res.Data, 5)

// 	// query active gauges by pool denom with a page request of 15 should return 10 gauges
// 	res, err = suite.QueryServer.ActiveGaugesPerDenom(sdk.WrapSDKContext(suite.Ctx), &types.ActiveGaugesPerDenomRequest{Denom: "pool", Pagination: &query.PageRequest{Limit: 15}})
// 	suite.Require().Len(res.Data, 10)
// }

// // TestGRPCUpcomingGauges tests querying upcoming gauges via gRPC returns the correct response.
// func (suite *KeeperTestSuite) TestGRPCUpcomingGauges() {
// 	suite.SetupTest()

// 	// ensure initially querying upcoming gauges returns no gauges
// 	res, err := suite.QueryServer.UpcomingGauges(sdk.WrapSDKContext(suite.Ctx), &types.UpcomingGetGaugesActiveUpcomingRequest{})
// 	suite.Require().NoError(err)
// 	suite.Require().Len(res.Data, 0)

// 	// create a gauge
// 	gaugeID, _, coins, startTime := suite.SetupGauge(false, sdk.Coins{sdk.NewInt64Coin("stake", 10)})

// 	// query upcoming gauges again, but this time expect the gauge created earlier in the response
// 	res, err = suite.QueryServer.UpcomingGauges(sdk.WrapSDKContext(suite.Ctx), &types.UpcomingGetGaugesActiveUpcomingRequest{})
// 	suite.Require().NoError(err)
// 	suite.Require().Len(res.Data, 1)
// 	expectedGauge := types.Gauge{
// 		Id:           gaugeID,
// 		IsPerpetual:  false,
// 		DistributeTo: types.QueryCondition{
// 			// LockQueryType: types.ByDuration,
// 			// Denom:         "lptoken",
// 			// Duration:      time.Second,
// 		},
// 		Coins:             coins,
// 		NumEpochsPaidOver: 2,
// 		FilledEpochs:      0,
// 		DistributedCoins:  sdk.Coins{},
// 		StartTime:         startTime,
// 	}
// 	suite.Require().Equal(res.Data[0].String(), expectedGauge.String())

// 	// setup 20 more upcoming gauges
// 	for i := 0; i < 20; i++ {
// 		_, gauge, _, _ := suite.SetupGauge(false, sdk.Coins{sdk.NewInt64Coin("stake", 3)})
// 		suite.Ctx = suite.Ctx.WithBlockTime(startTime.Add(time.Second))

// 		// move the first 9 created gauges to an active status
// 		// 1 + (20 -9) = 12 upcoming gauges
// 		if i < 9 {
// 			suite.QueryServer.MoveUpcomingGaugeToActiveGauge(suite.Ctx, *gauge)
// 		}
// 	}

// 	// query upcoming gauges with a page request of 5 should return 5 gauges
// 	res, err = suite.QueryServer.UpcomingGauges(sdk.WrapSDKContext(suite.Ctx), &types.UpcomingGetGaugesActiveUpcomingRequest{Pagination: &query.PageRequest{Limit: 5}})
// 	suite.Require().Len(res.Data, 5)

// 	// query upcoming gauges with a page request of 15 should return 12 gauges
// 	res, err = suite.QueryServer.UpcomingGauges(sdk.WrapSDKContext(suite.Ctx), &types.UpcomingGetGaugesActiveUpcomingRequest{Pagination: &query.PageRequest{Limit: 15}})
// 	suite.Require().Len(res.Data, 12)
// }

// // TestGRPCUpcomingGaugesPerDenom tests querying upcoming gauges by denom via gRPC returns the correct response.
// func (suite *KeeperTestSuite) TestGRPCUpcomingGaugesPerDenom() {
// 	suite.SetupTest()

// 	// ensure initially querying upcoming gauges by denom returns no gauges
// 	upcomingGaugeRequest := types.UpcomingGaugesPerDenomRequest{Denom: "lptoken", Pagination: nil}
// 	res, err := suite.QueryServer.UpcomingGaugesPerDenom(sdk.WrapSDKContext(suite.Ctx), &upcomingGaugeRequest)
// 	suite.Require().NoError(err)
// 	suite.Require().Len(res.UpcomingGauges, 0)

// 	// create a gauge, and check upcoming gauge is working
// 	gaugeID, gauge, coins, startTime := suite.SetupGauge(false, sdk.Coins{sdk.NewInt64Coin("stake", 10)})

// 	// query upcoming gauges by denom again, but this time expect the gauge created earlier in the response
// 	res, err = suite.QueryServer.UpcomingGaugesPerDenom(sdk.WrapSDKContext(suite.Ctx), &upcomingGaugeRequest)
// 	suite.Require().NoError(err)
// 	suite.Require().Len(res.UpcomingGauges, 1)
// 	expectedGauge := types.Gauge{
// 		Id:           gaugeID,
// 		IsPerpetual:  false,
// 		DistributeTo: types.QueryCondition{
// 			// LockQueryType: types.ByDuration,
// 			// Denom:         "lptoken",
// 			// Duration:      time.Second,
// 		},
// 		Coins:             coins,
// 		NumEpochsPaidOver: 2,
// 		FilledEpochs:      0,
// 		DistributedCoins:  sdk.Coins{},
// 		StartTime:         startTime,
// 	}
// 	suite.Require().Equal(res.UpcomingGauges[0].String(), expectedGauge.String())

// 	// move gauge from upcoming to active
// 	// ensure the query no longer returns a response
// 	suite.Ctx = suite.Ctx.WithBlockTime(startTime.Add(time.Second))
// 	err = suite.App.IncentivesKeeper.MoveUpcomingGaugeToActiveGauge(suite.Ctx, *gauge)
// 	res, err = suite.QueryServer.UpcomingGaugesPerDenom(sdk.WrapSDKContext(suite.Ctx), &upcomingGaugeRequest)
// 	suite.Require().NoError(err)
// 	suite.Require().Len(res.UpcomingGauges, 0)

// 	// setup 20 more upcoming gauges with pool denom
// 	for i := 0; i < 20; i++ {
// 		_, gauge, _, _ := suite.SetupNewGaugeWithDenom(false, sdk.Coins{sdk.NewInt64Coin("stake", 3)}, "pool")
// 		suite.Ctx = suite.Ctx.WithBlockTime(startTime.Add(time.Second))

// 		// move the first 10 created gauges from upcoming to active
// 		// this leaves 10 upcoming gauges
// 		if i < 10 {
// 			suite.QueryServer.MoveUpcomingGaugeToActiveGauge(suite.Ctx, *gauge)
// 		}
// 	}

// 	// query upcoming gauges by lptoken denom with a page request of 5 should return 0 gauges
// 	res, err = suite.QueryServer.UpcomingGaugesPerDenom(sdk.WrapSDKContext(suite.Ctx), &types.UpcomingGaugesPerDenomRequest{Denom: "lptoken", Pagination: &query.PageRequest{Limit: 5}})
// 	suite.Require().Len(res.UpcomingGauges, 0)

// 	// query upcoming gauges by pool denom with a page request of 5 should return 5 gauges
// 	res, err = suite.QueryServer.UpcomingGaugesPerDenom(sdk.WrapSDKContext(suite.Ctx), &types.UpcomingGaugesPerDenomRequest{Denom: "pool", Pagination: &query.PageRequest{Limit: 5}})
// 	suite.Require().Len(res.UpcomingGauges, 5)

// 	// query upcoming gauges by pool denom with a page request of 15 should return 10 gauges
// 	res, err = suite.QueryServer.UpcomingGaugesPerDenom(sdk.WrapSDKContext(suite.Ctx), &types.UpcomingGaugesPerDenomRequest{Denom: "pool", Pagination: &query.PageRequest{Limit: 15}})
// 	suite.Require().Len(res.UpcomingGauges, 10)
// }

// // // TestGRPCRewardsEstimate tests querying rewards estimation at a future specific time (by epoch) via gRPC returns the correct response.
// // func (suite *KeeperTestSuite) TestGRPCRewardsEstimate() {
// // 	suite.SetupTest()

// // 	// create an address with no locks
// // 	// ensure rewards estimation returns a nil coins struct
// // 	lockOwner := sdk.AccAddress([]byte("addr1---------------"))
// // 	res, err := suite.QueryServer.RewardsEstimate(sdk.WrapSDKContext(suite.Ctx), &types.RewardsEstimateRequest{
// // 		Owner: lockOwner.String(),
// // 	})
// // 	suite.Require().NoError(err)
// // 	suite.Require().Equal(res.Coins, sdk.Coins{})

// // 	// setup a lock and gauge for a new address
// // 	lockOwner, _, coins, _ := suite.SetupLockAndGauge(false)

// // 	// query the rewards of the new address after 100 epochs
// // 	// since it is the only address the gauge is paying out to, the future rewards should equal the entirety of the gauge
// // 	res, err = suite.QueryServer.RewardsEstimate(sdk.WrapSDKContext(suite.Ctx), &types.RewardsEstimateRequest{
// // 		Owner:    lockOwner.String(),
// // 		EndEpoch: 100,
// // 	})
// // 	suite.Require().NoError(err)
// // 	suite.Require().Equal(res.Coins, coins)
// // }

// // // TestRewardsEstimateWithPoolIncentives tests querying rewards estimation at a future specific time (by epoch) via gRPC returns the correct response.
// // // Also changes distribution records for the pool incentives to distribute to the respective lock owner.
// // func (suite *KeeperTestSuite) TestRewardsEstimateWithPoolIncentives() {
// // 	suite.SetupTest()

// // 	// create an address with no locks
// // 	// ensure rewards estimation returns a nil coins struct
// // 	lockOwner := sdk.AccAddress([]byte("addr1---------------"))
// // 	res, err := suite.QueryServer.RewardsEstimate(sdk.WrapSDKContext(suite.Ctx), &types.RewardsEstimateRequest{
// // 		Owner: lockOwner.String(),
// // 	})
// // 	suite.Require().NoError(err)
// // 	suite.Require().Equal(res.Coins, sdk.Coins{})

// // 	// setup a lock and gauge for a new address
// // 	lockOwner, gaugeID, coins, _ := suite.SetupLockAndGauge(true)

// // 	// take newly created gauge and modify its pool incentives distribution weight to 100
// // 	distrRecord := pooltypes.DistrRecord{
// // 		GaugeId: gaugeID,
// // 		Weight:  sdk.NewInt(100),
// // 	}
// // 	err = suite.App.IncentivesKeeper.ReplaceDistrRecords(suite.Ctx, distrRecord)
// // 	suite.Require().NoError(err)

// // 	// query the rewards of the new address after the 10th epoch
// // 	// since it is the only address the gauge is paying out to, the future rewards should equal the entirety of the gauge
// // 	res, err = suite.QueryServer.RewardsEstimate(sdk.WrapSDKContext(suite.Ctx), &types.RewardsEstimateRequest{
// // 		Owner:    lockOwner.String(),
// // 		EndEpoch: 10,
// // 	})
// // 	suite.Require().NoError(err)
// // 	suite.Require().Equal(res.Coins, coins)

// // 	// after the current epoch ends, mint more coins that matches the lock coin demon created earlier
// // 	epochIdentifier := suite.App.MintKeeper.GetParams(suite.Ctx).EpochIdentifier
// // 	curEpochNumber := suite.App.EpochsKeeper.GetEpochInfo(suite.Ctx, epochIdentifier).CurrentEpoch
// // 	suite.App.EpochsKeeper.AfterEpochEnd(suite.Ctx, epochIdentifier, curEpochNumber)
// // 	// TODO: Figure out what this number should be
// // 	// TODO: Respond to this
// // 	mintCoins := sdk.NewCoin(coins[0].Denom, sdk.NewInt(1500000))

// // 	// query the rewards of the new address after the 10th epoch
// // 	// since it is the only address the gauge is paying out to, the future rewards should equal the entirety of the gauge plus the newly minted coins
// // 	res, err = suite.QueryServer.RewardsEstimate(sdk.WrapSDKContext(suite.Ctx), &types.RewardsEstimateRequest{
// // 		Owner:    lockOwner.String(),
// // 		EndEpoch: 10,
// // 	})
// // 	suite.Require().NoError(err)
// // 	suite.Require().Equal(res.Coins, coins.Add(mintCoins))
// // }

// // TestGRPCToDistributeCoins tests querying coins that are going to be distributed via gRPC returns the correct response.
// func (suite *KeeperTestSuite) TestGRPCToDistributeCoins() {
// 	suite.SetupTest()

// 	// ensure initially querying to distribute coins returns no coins
// 	res, err := suite.QueryServer.GetModuleCoinsToBeDistributed(sdk.WrapSDKContext(suite.Ctx), &types.GetModuleCoinsToBeDistributedRequest{})
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(res.Coins, sdk.Coins(nil))

// 	// create two locks with different durations
// 	addr1 := sdk.AccAddress([]byte("addr1---------------"))
// 	addr2 := sdk.AccAddress([]byte("addr2---------------"))
// 	suite.SetupLock(addr1, sdk.Coins{sdk.NewInt64Coin("lptoken", 10)}, time.Second)
// 	suite.SetupLock(addr2, sdk.Coins{sdk.NewInt64Coin("lptoken", 10)}, 2*time.Second)

// 	// setup a non perpetual gauge
// 	gaugeID, _, coins, startTime := suite.SetupGauge(false, sdk.Coins{sdk.NewInt64Coin("stake", 10)})
// 	gauge, err := suite.QueryServer.GetGaugeByID(suite.Ctx, gaugeID)
// 	suite.Require().NoError(err)
// 	suite.Require().NotNil(gauge)
// 	gauges := types.Gauges{*gauge}

// 	// check to distribute coins after gauge creation
// 	// ensure this equals the coins within the previously created non perpetual gauge
// 	res, err = suite.QueryServer.GetModuleCoinsToBeDistributed(sdk.WrapSDKContext(suite.Ctx), &types.GetModuleCoinsToBeDistributedRequest{})
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(res.Coins, coins)

// 	// distribute coins to stakers
// 	distCoins, err := suite.QueryServer.Distribute(suite.Ctx, gauges)
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(distCoins, sdk.Coins{sdk.NewInt64Coin("stake", 4)})

// 	// check gauge changes after distribution
// 	// ensure the gauge's filled epochs have been increased by 1
// 	// ensure we have distributed 4 out of the 10 stake tokens
// 	gauge, err = suite.QueryServer.GetGaugeByID(suite.Ctx, gaugeID)
// 	suite.Require().NoError(err)
// 	suite.Require().NotNil(gauge)
// 	suite.Require().Equal(gauge.FilledEpochs, uint64(1))
// 	suite.Require().Equal(gauge.DistributedCoins, sdk.Coins{sdk.NewInt64Coin("stake", 4)})
// 	gauges = types.Gauges{*gauge}

// 	// move gauge from an upcoming to an active status
// 	suite.Ctx = suite.Ctx.WithBlockTime(startTime)
// 	err = suite.QueryServer.MoveUpcomingGaugeToActiveGauge(suite.Ctx, *gauge)
// 	suite.Require().NoError(err)

// 	// check that the to distribute coins is equal to the initial gauge coin balance minus what has been distributed already (10-4=6)
// 	res, err = suite.QueryServer.GetModuleCoinsToBeDistributed(sdk.WrapSDKContext(suite.Ctx), &types.GetModuleCoinsToBeDistributedRequest{})
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(res.Coins, coins.Sub(distCoins))

// 	// distribute second round to stakers
// 	distCoins, err = suite.QueryServer.Distribute(suite.Ctx, gauges)
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(sdk.Coins{sdk.NewInt64Coin("stake", 6)}, distCoins)

// 	// now that all coins have been distributed (4 in first found 6 in the second round)
// 	// to distribute coins should be null
// 	res, err = suite.QueryServer.GetModuleCoinsToBeDistributed(sdk.WrapSDKContext(suite.Ctx), &types.GetModuleCoinsToBeDistributedRequest{})
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(res.Coins, sdk.Coins(nil))
// }

// // TestGRPCDistributedCoins tests querying coins that have been distributed via gRPC returns the correct response.
// func (suite *KeeperTestSuite) TestGRPCDistributedCoins() {
// 	suite.SetupTest()

// 	// create two locks with different durations
// 	addr1 := sdk.AccAddress([]byte("addr1---------------"))
// 	addr2 := sdk.AccAddress([]byte("addr2---------------"))
// 	suite.SetupLock(addr1, sdk.Coins{sdk.NewInt64Coin("lptoken", 10)}, time.Second)
// 	suite.SetupLock(addr2, sdk.Coins{sdk.NewInt64Coin("lptoken", 10)}, 2*time.Second)

// 	// setup a non perpetual gauge
// 	gaugeID, _, _, startTime := suite.SetupGauge(false, sdk.Coins{sdk.NewInt64Coin("stake", 10)})
// 	gauge, err := suite.QueryServer.GetGaugeByID(suite.Ctx, gaugeID)
// 	suite.Require().NoError(err)
// 	suite.Require().NotNil(gauge)
// 	gauges := types.Gauges{*gauge}

// 	// move gauge from upcoming to active
// 	suite.Ctx = suite.Ctx.WithBlockTime(startTime)
// 	err = suite.QueryServer.MoveUpcomingGaugeToActiveGauge(suite.Ctx, *gauge)
// 	suite.Require().NoError(err)

// 	// distribute coins to stakers
// 	distCoins, err := suite.QueryServer.Distribute(suite.Ctx, gauges)
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(distCoins, sdk.Coins{sdk.NewInt64Coin("stake", 4)})

// 	// check gauge changes after distribution
// 	// ensure the gauge's filled epochs have been increased by 1
// 	// ensure we have distributed 4 out of the 10 stake tokens
// 	gauge, err = suite.QueryServer.GetGaugeByID(suite.Ctx, gaugeID)
// 	suite.Require().NoError(err)
// 	suite.Require().NotNil(gauge)
// 	suite.Require().Equal(gauge.FilledEpochs, uint64(1))
// 	suite.Require().Equal(gauge.DistributedCoins, sdk.Coins{sdk.NewInt64Coin("stake", 4)})
// 	gauges = types.Gauges{*gauge}

// 	// distribute second round to stakers
// 	distCoins, err = suite.QueryServer.Distribute(suite.Ctx, gauges)
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(sdk.Coins{sdk.NewInt64Coin("stake", 6)}, distCoins)
// }

// func (suite *KeeperTestSuite) BeginUnlocking(addr sdk.AccAddress) {
// 	_, err := suite.QueryServer.BeginUnlockAllNotUnlockings(suite.Ctx, addr)
// 	suite.Require().NoError(err)
// }

// func (suite *KeeperTestSuite) WithdrawAllMaturedLocks() {
// 	suite.QueryServer.WithdrawAllMaturedLocks(suite.Ctx)
// }

// func (suite *KeeperTestSuite) TestModuleBalance() {
// 	suite.SetupTest()

// 	// initial check
// 	res, err := suite.QueryServer.ModuleBalance(sdk.WrapSDKContext(suite.Ctx), &types.ModuleBalanceRequest{})
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(res.Coins, sdk.Coins{})

// 	// lock coins
// 	addr1 := sdk.AccAddress([]byte("addr1---------------"))
// 	coins := sdk.Coins{sdk.NewInt64Coin("stake", 10)}
// 	suite.SetupLock(addr1, coins, time.Second)

// 	// final check
// 	res, err = suite.QueryServer.ModuleBalance(sdk.WrapSDKContext(suite.Ctx), &types.ModuleBalanceRequest{})
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(res.Coins, coins)
// }

// func (suite *KeeperTestSuite) TestModuleLockedAmount() {
// 	// test for module locked balance check
// 	suite.SetupTest()

// 	// initial check
// 	res, err := suite.QueryServer.ModuleLockedAmount(sdk.WrapSDKContext(suite.Ctx), &types.ModuleLockedAmountRequest{})
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(res.Coins, sdk.Coins(nil))

// 	// lock coins
// 	addr1 := sdk.AccAddress([]byte("addr1---------------"))
// 	coins := sdk.Coins{sdk.NewInt64Coin("stake", 10)}
// 	suite.SetupLock(addr1, coins, time.Second)
// 	suite.BeginUnlocking(addr1)

// 	// current module locked balance check = unlockTime - 1s
// 	res, err = suite.QueryServer.ModuleLockedAmount(sdk.WrapSDKContext(suite.Ctx), &types.ModuleLockedAmountRequest{})
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(res.Coins, coins)

// 	// module locked balance after 1 second = unlockTime
// 	now := suite.Ctx.BlockTime()
// 	res, err = suite.QueryServer.ModuleLockedAmount(sdk.WrapSDKContext(suite.Ctx.WithBlockTime(now.Add(time.Second))), &types.ModuleLockedAmountRequest{})
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(res.Coins, sdk.Coins(nil))

// 	// module locked balance after 2 second = unlockTime + 1s
// 	res, err = suite.QueryServer.ModuleLockedAmount(sdk.WrapSDKContext(suite.Ctx.WithBlockTime(now.Add(2*time.Second))), &types.ModuleLockedAmountRequest{})
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(res.Coins, sdk.Coins(nil))
// }

// func (suite *KeeperTestSuite) TestAccountUnlockableCoins() {
// 	suite.SetupTest()
// 	addr1 := sdk.AccAddress([]byte("addr1---------------"))

// 	// empty address unlockable coins check
// 	_, err := suite.QueryServer.AccountUnlockableCoins(sdk.WrapSDKContext(suite.Ctx), &types.AccountUnlockableCoinsRequest{Owner: ""})
// 	suite.Require().Error(err)

// 	// initial check
// 	res, err := suite.QueryServer.AccountUnlockableCoins(sdk.WrapSDKContext(suite.Ctx), &types.AccountUnlockableCoinsRequest{Owner: addr1.String()})
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(res.Coins, sdk.Coins{})

// 	// lock coins
// 	coins := sdk.Coins{sdk.NewInt64Coin("stake", 10)}
// 	suite.SetupLock(addr1, coins, time.Second)

// 	// check before start unlocking
// 	res, err = suite.QueryServer.AccountUnlockableCoins(sdk.WrapSDKContext(suite.Ctx), &types.AccountUnlockableCoinsRequest{Owner: addr1.String()})
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(res.Coins, sdk.Coins{})

// 	suite.BeginUnlocking(addr1)

// 	// check = unlockTime - 1s
// 	res, err = suite.QueryServer.AccountUnlockableCoins(sdk.WrapSDKContext(suite.Ctx), &types.AccountUnlockableCoinsRequest{Owner: addr1.String()})
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(res.Coins, sdk.Coins{})

// 	// check after 1 second = unlockTime
// 	now := suite.Ctx.BlockTime()
// 	res, err = suite.QueryServer.AccountUnlockableCoins(sdk.WrapSDKContext(suite.Ctx.WithBlockTime(now.Add(time.Second))), &types.AccountUnlockableCoinsRequest{Owner: addr1.String()})
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(res.Coins, coins)

// 	// check after 2 second = unlockTime + 1s
// 	res, err = suite.QueryServer.AccountUnlockableCoins(sdk.WrapSDKContext(suite.Ctx.WithBlockTime(now.Add(2*time.Second))), &types.AccountUnlockableCoinsRequest{Owner: addr1.String()})
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(res.Coins, coins)
// }

// func (suite *KeeperTestSuite) TestAccountUnlockingCoins() {
// 	suite.SetupTest()
// 	addr1 := sdk.AccAddress([]byte("addr1---------------"))

// 	// empty address unlockable coins check
// 	_, err := suite.QueryServer.AccountUnlockingCoins(sdk.WrapSDKContext(suite.Ctx), &types.AccountUnlockingCoinsRequest{Owner: ""})
// 	suite.Require().Error(err)

// 	// initial check
// 	res, err := suite.QueryServer.AccountUnlockingCoins(sdk.WrapSDKContext(suite.Ctx), &types.AccountUnlockingCoinsRequest{Owner: addr1.String()})
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(res.Coins, sdk.Coins{})

// 	// lock coins
// 	coins := sdk.Coins{sdk.NewInt64Coin("stake", 10)}
// 	suite.SetupLock(addr1, coins, time.Second)

// 	// check before start unlocking
// 	res, err = suite.QueryServer.AccountUnlockingCoins(sdk.WrapSDKContext(suite.Ctx), &types.AccountUnlockingCoinsRequest{Owner: addr1.String()})
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(res.Coins, sdk.Coins{})

// 	suite.BeginUnlocking(addr1)

// 	// check at unlockTime - 1s
// 	res, err = suite.QueryServer.AccountUnlockingCoins(sdk.WrapSDKContext(suite.Ctx), &types.AccountUnlockingCoinsRequest{Owner: addr1.String()})
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(res.Coins, sdk.Coins{sdk.NewInt64Coin("stake", 10)})

// 	// check after 1 second = unlockTime
// 	now := suite.Ctx.BlockTime()
// 	res, err = suite.QueryServer.AccountUnlockingCoins(sdk.WrapSDKContext(suite.Ctx.WithBlockTime(now.Add(time.Second))), &types.AccountUnlockingCoinsRequest{Owner: addr1.String()})
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(res.Coins, sdk.Coins{})

// 	// check after 2 second = unlockTime + 1s
// 	res, err = suite.QueryServer.AccountUnlockingCoins(sdk.WrapSDKContext(suite.Ctx.WithBlockTime(now.Add(2*time.Second))), &types.AccountUnlockingCoinsRequest{Owner: addr1.String()})
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(res.Coins, sdk.Coins{})
// }

// func (suite *KeeperTestSuite) TestAccountLockedCoins() {
// 	suite.SetupTest()
// 	addr1 := sdk.AccAddress([]byte("addr1---------------"))

// 	// empty address locked coins check
// 	_, err := suite.QueryServer.AccountLockedCoins(sdk.WrapSDKContext(suite.Ctx), &types.AccountLockedCoinsRequest{})
// 	suite.Require().Error(err)

// 	// initial check
// 	res, err := suite.QueryServer.AccountLockedCoins(sdk.WrapSDKContext(suite.Ctx), &types.AccountLockedCoinsRequest{Owner: addr1.String()})
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(res.Coins, sdk.Coins(nil))

// 	// lock coins
// 	coins := sdk.Coins{sdk.NewInt64Coin("stake", 10)}
// 	suite.SetupLock(addr1, coins, time.Second)
// 	suite.BeginUnlocking(addr1)

// 	// check = unlockTime - 1s
// 	res, err = suite.QueryServer.AccountLockedCoins(sdk.WrapSDKContext(suite.Ctx), &types.AccountLockedCoinsRequest{Owner: addr1.String()})
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(coins, res.Coins)

// 	// check after 1 second = unlockTime
// 	now := suite.Ctx.BlockTime()
// 	res, err = suite.QueryServer.AccountLockedCoins(sdk.WrapSDKContext(suite.Ctx.WithBlockTime(now.Add(time.Second))), &types.AccountLockedCoinsRequest{Owner: addr1.String()})
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(res.Coins, sdk.Coins(nil))

// 	// check after 2 second = unlockTime + 1s
// 	res, err = suite.QueryServer.AccountLockedCoins(sdk.WrapSDKContext(suite.Ctx.WithBlockTime(now.Add(2*time.Second))), &types.AccountLockedCoinsRequest{Owner: addr1.String()})
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(res.Coins, sdk.Coins(nil))
// }

// func (suite *KeeperTestSuite) TestAccountLockedPastTime() {
// 	suite.SetupTest()
// 	addr1 := sdk.AccAddress([]byte("addr1---------------"))
// 	now := suite.Ctx.BlockTime()

// 	// empty address locks check
// 	_, err := suite.QueryServer.AccountLockedPastTime(sdk.WrapSDKContext(suite.Ctx), &types.AccountLockedPastTimeRequest{Owner: "", Timestamp: now})
// 	suite.Require().Error(err)

// 	// initial check
// 	res, err := suite.QueryServer.AccountLockedPastTime(sdk.WrapSDKContext(suite.Ctx), &types.AccountLockedPastTimeRequest{Owner: addr1.String(), Timestamp: now})
// 	suite.Require().NoError(err)
// 	suite.Require().Len(res.Locks, 0)

// 	// lock coins
// 	coins := sdk.Coins{sdk.NewInt64Coin("stake", 10)}
// 	suite.SetupLock(addr1, coins, time.Second)
// 	suite.BeginUnlocking(addr1)

// 	// check = unlockTime - 1s
// 	res, err = suite.QueryServer.AccountLockedPastTime(sdk.WrapSDKContext(suite.Ctx), &types.AccountLockedPastTimeRequest{Owner: addr1.String(), Timestamp: now})
// 	suite.Require().NoError(err)
// 	suite.Require().Len(res.Locks, 1)

// 	// check after 1 second = unlockTime
// 	res, err = suite.QueryServer.AccountLockedPastTime(sdk.WrapSDKContext(suite.Ctx), &types.AccountLockedPastTimeRequest{Owner: addr1.String(), Timestamp: now.Add(time.Second)})
// 	suite.Require().NoError(err)
// 	suite.Require().Len(res.Locks, 0)

// 	// check after 2 second = unlockTime + 1s
// 	res, err = suite.QueryServer.AccountLockedPastTime(sdk.WrapSDKContext(suite.Ctx), &types.AccountLockedPastTimeRequest{Owner: addr1.String(), Timestamp: now.Add(2 * time.Second)})
// 	suite.Require().NoError(err)
// 	suite.Require().Len(res.Locks, 0)
// }

// func (suite *KeeperTestSuite) TestAccountLockedPastTimeNotUnlockingOnly() {
// 	suite.SetupTest()
// 	addr1 := sdk.AccAddress([]byte("addr1---------------"))
// 	now := suite.Ctx.BlockTime()

// 	// empty address locks check
// 	_, err := suite.QueryServer.AccountLockedPastTimeNotUnlockingOnly(sdk.WrapSDKContext(suite.Ctx), &types.AccountLockedPastTimeNotUnlockingOnlyRequest{Owner: "", Timestamp: now})
// 	suite.Require().Error(err)

// 	// initial check
// 	res, err := suite.QueryServer.AccountLockedPastTimeNotUnlockingOnly(sdk.WrapSDKContext(suite.Ctx), &types.AccountLockedPastTimeNotUnlockingOnlyRequest{Owner: addr1.String(), Timestamp: now})
// 	suite.Require().NoError(err)
// 	suite.Require().Len(res.Locks, 0)

// 	// lock coins
// 	coins := sdk.Coins{sdk.NewInt64Coin("stake", 10)}
// 	suite.SetupLock(addr1, coins, time.Second)

// 	// check when not start unlocking
// 	res, err = suite.QueryServer.AccountLockedPastTimeNotUnlockingOnly(sdk.WrapSDKContext(suite.Ctx), &types.AccountLockedPastTimeNotUnlockingOnlyRequest{Owner: addr1.String(), Timestamp: now})
// 	suite.Require().NoError(err)
// 	suite.Require().Len(res.Locks, 1)

// 	// begin unlocking
// 	suite.BeginUnlocking(addr1)

// 	// check after start unlocking
// 	res, err = suite.QueryServer.AccountLockedPastTimeNotUnlockingOnly(sdk.WrapSDKContext(suite.Ctx), &types.AccountLockedPastTimeNotUnlockingOnlyRequest{Owner: addr1.String(), Timestamp: now})
// 	suite.Require().NoError(err)
// 	suite.Require().Len(res.Locks, 0)
// }

// func (suite *KeeperTestSuite) TestAccountUnlockedBeforeTime() {
// 	suite.SetupTest()
// 	addr1 := sdk.AccAddress([]byte("addr1---------------"))
// 	now := suite.Ctx.BlockTime()

// 	// empty address unlockables check
// 	_, err := suite.QueryServer.AccountUnlockedBeforeTime(sdk.WrapSDKContext(suite.Ctx), &types.AccountUnlockedBeforeTimeRequest{Owner: "", Timestamp: now})
// 	suite.Require().Error(err)

// 	// initial check
// 	res, err := suite.QueryServer.AccountUnlockedBeforeTime(sdk.WrapSDKContext(suite.Ctx), &types.AccountUnlockedBeforeTimeRequest{Owner: addr1.String(), Timestamp: now})
// 	suite.Require().NoError(err)
// 	suite.Require().Len(res.Locks, 0)

// 	// lock coins
// 	coins := sdk.Coins{sdk.NewInt64Coin("stake", 10)}
// 	suite.SetupLock(addr1, coins, time.Second)
// 	suite.BeginUnlocking(addr1)

// 	// check = unlockTime - 1s
// 	res, err = suite.QueryServer.AccountUnlockedBeforeTime(sdk.WrapSDKContext(suite.Ctx), &types.AccountUnlockedBeforeTimeRequest{Owner: addr1.String(), Timestamp: now})
// 	suite.Require().NoError(err)
// 	suite.Require().Len(res.Locks, 0)

// 	// check after 1 second = unlockTime
// 	res, err = suite.QueryServer.AccountUnlockedBeforeTime(sdk.WrapSDKContext(suite.Ctx), &types.AccountUnlockedBeforeTimeRequest{Owner: addr1.String(), Timestamp: now.Add(time.Second)})
// 	suite.Require().NoError(err)
// 	suite.Require().Len(res.Locks, 1)

// 	// check after 2 second = unlockTime + 1s
// 	res, err = suite.QueryServer.AccountUnlockedBeforeTime(sdk.WrapSDKContext(suite.Ctx), &types.AccountUnlockedBeforeTimeRequest{Owner: addr1.String(), Timestamp: now.Add(2 * time.Second)})
// 	suite.Require().NoError(err)
// 	suite.Require().Len(res.Locks, 1)
// }

// func (suite *KeeperTestSuite) TestAccountLockedPastTimeDenom() {
// 	suite.SetupTest()
// 	addr1 := sdk.AccAddress([]byte("addr1---------------"))
// 	now := suite.Ctx.BlockTime()

// 	// empty address locks by denom check
// 	_, err := suite.QueryServer.AccountLockedPastTimeDenom(sdk.WrapSDKContext(suite.Ctx), &types.AccountLockedPastTimeDenomRequest{Owner: "", Denom: "stake", Timestamp: now})
// 	suite.Require().Error(err)

// 	// initial check
// 	res, err := suite.QueryServer.AccountLockedPastTimeDenom(sdk.WrapSDKContext(suite.Ctx), &types.AccountLockedPastTimeDenomRequest{Owner: addr1.String(), Denom: "stake", Timestamp: now})
// 	suite.Require().NoError(err)
// 	suite.Require().Len(res.Locks, 0)

// 	// lock coins
// 	coins := sdk.Coins{sdk.NewInt64Coin("stake", 10)}
// 	suite.SetupLock(addr1, coins, time.Second)
// 	suite.BeginUnlocking(addr1)

// 	// check = unlockTime - 1s
// 	res, err = suite.QueryServer.AccountLockedPastTimeDenom(sdk.WrapSDKContext(suite.Ctx), &types.AccountLockedPastTimeDenomRequest{Owner: addr1.String(), Denom: "stake", Timestamp: now})
// 	suite.Require().NoError(err)
// 	suite.Require().Len(res.Locks, 1)

// 	// account locks by not available denom
// 	res, err = suite.QueryServer.AccountLockedPastTimeDenom(sdk.WrapSDKContext(suite.Ctx), &types.AccountLockedPastTimeDenomRequest{Owner: addr1.String(), Denom: "stake2", Timestamp: now})
// 	suite.Require().NoError(err)
// 	suite.Require().Len(res.Locks, 0)

// 	// account locks by denom after 1 second = unlockTime
// 	res, err = suite.QueryServer.AccountLockedPastTimeDenom(sdk.WrapSDKContext(suite.Ctx), &types.AccountLockedPastTimeDenomRequest{Owner: addr1.String(), Denom: "stake", Timestamp: now.Add(time.Second)})
// 	suite.Require().NoError(err)
// 	suite.Require().Len(res.Locks, 0)

// 	// account locks by denom after 2 second = unlockTime + 1s
// 	res, err = suite.QueryServer.AccountLockedPastTimeDenom(sdk.WrapSDKContext(suite.Ctx), &types.AccountLockedPastTimeDenomRequest{Owner: addr1.String(), Denom: "stake", Timestamp: now.Add(2 * time.Second)})
// 	suite.Require().NoError(err)
// 	suite.Require().Len(res.Locks, 0)

// 	// try querying with prefix coins like "stak" for potential attack
// 	res, err = suite.QueryServer.AccountLockedPastTimeDenom(sdk.WrapSDKContext(suite.Ctx), &types.AccountLockedPastTimeDenomRequest{Owner: addr1.String(), Denom: "stak", Timestamp: now})
// 	suite.Require().NoError(err)
// 	suite.Require().Len(res.Locks, 0)
// }

// func (suite *KeeperTestSuite) TestLockedByID() {
// 	suite.SetupTest()
// 	addr1 := sdk.AccAddress([]byte("addr1---------------"))

// 	// lock by not available id check
// 	res, err := suite.QueryServer.LockedByID(sdk.WrapSDKContext(suite.Ctx), &types.LockedRequest{LockId: 0})
// 	suite.Require().Error(err)

// 	// lock coins
// 	coins := sdk.Coins{sdk.NewInt64Coin("stake", 10)}
// 	suite.SetupLock(addr1, coins, time.Second)

// 	// lock by available available id check
// 	res, err = suite.QueryServer.LockedByID(sdk.WrapSDKContext(suite.Ctx), &types.LockedRequest{LockId: 1})
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(res.Lock.ID, uint64(1))
// 	suite.Require().Equal(res.Lock.Owner, addr1.String())
// 	suite.Require().Equal(res.Lock.Coins, coins)
// 	suite.Require().Equal(res.Lock.EndTime, time.Time{})
// 	suite.Require().Equal(res.Lock.IsUnlocking(), false)
// }

// func (suite *KeeperTestSuite) TestNextLockID() {
// 	suite.SetupTest()
// 	addr1 := sdk.AccAddress([]byte("addr1---------------"))

// 	// lock coins
// 	coins := sdk.Coins{sdk.NewInt64Coin("stake", 10)}
// 	suite.SetupLock(addr1, coins, time.Second)

// 	// lock by available available id check
// 	res, err := suite.QueryServer.NextLockID(sdk.WrapSDKContext(suite.Ctx), &types.NextLockIDRequest{})
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(res.LockId, uint64(2))

// 	// create 2 more locks
// 	coins = sdk.Coins{sdk.NewInt64Coin("stake", 10)}
// 	suite.SetupLock(addr1, coins, time.Second)
// 	coins = sdk.Coins{sdk.NewInt64Coin("stake", 10)}
// 	suite.SetupLock(addr1, coins, time.Second)
// 	res, err = suite.QueryServer.NextLockID(sdk.WrapSDKContext(suite.Ctx), &types.NextLockIDRequest{})
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(res.LockId, uint64(4))
// }

// // func (suite *KeeperTestSuite) TestAccountLockedLongerDuration() {
// // 	suite.SetupTest()
// // 	addr1 := sdk.AccAddress([]byte("addr1---------------"))

// // 	// empty address locks longer than duration check
// // 	res, err := suite.QueryServer.AccountLockedLongerDuration(sdk.WrapSDKContext(suite.Ctx), &types.AccountLockedLongerDurationRequest{Owner: "", Duration: time.Second})
// // 	suite.Require().Error(err)

// // 	// initial check
// // 	res, err = suite.QueryServer.AccountLockedLongerDuration(sdk.WrapSDKContext(suite.Ctx), &types.AccountLockedLongerDurationRequest{Owner: addr1.String(), Duration: time.Second})
// // 	suite.Require().NoError(err)
// // 	suite.Require().Len(res.Locks, 0)

// // 	// lock coins
// // 	coins := sdk.Coins{sdk.NewInt64Coin("stake", 10)}
// // 	suite.SetupLock(addr1, coins, time.Second)
// // 	suite.BeginUnlocking(addr1)

// // 	// account locks longer than duration check, duration = 0s
// // 	res, err = suite.QueryServer.AccountLockedLongerDuration(sdk.WrapSDKContext(suite.Ctx), &types.AccountLockedLongerDurationRequest{Owner: addr1.String(), Duration: 0})
// // 	suite.Require().NoError(err)
// // 	suite.Require().Len(res.Locks, 1)

// // 	// account locks longer than duration check, duration = 1s
// // 	res, err = suite.QueryServer.AccountLockedLongerDuration(sdk.WrapSDKContext(suite.Ctx), &types.AccountLockedLongerDurationRequest{Owner: addr1.String(), Duration: time.Second})
// // 	suite.Require().NoError(err)
// // 	suite.Require().Len(res.Locks, 1)

// // 	// account locks longer than duration check, duration = 2s
// // 	res, err = suite.QueryServer.AccountLockedLongerDuration(sdk.WrapSDKContext(suite.Ctx), &types.AccountLockedLongerDurationRequest{Owner: addr1.String(), Duration: 2 * time.Second})
// // 	suite.Require().NoError(err)
// // 	suite.Require().Len(res.Locks, 0)
// // }

// // func (suite *KeeperTestSuite) TestAccountLockedLongerDurationNotUnlockingOnly() {
// // 	suite.SetupTest()
// // 	addr1 := sdk.AccAddress([]byte("addr1---------------"))

// // 	// empty address locks longer than duration check
// // 	res, err := suite.QueryServer.AccountLockedLongerDurationNotUnlockingOnly(sdk.WrapSDKContext(suite.Ctx), &types.AccountLockedLongerDurationNotUnlockingOnlyRequest{Owner: "", Duration: time.Second})
// // 	suite.Require().Error(err)

// // 	// initial check
// // 	res, err = suite.QueryServer.AccountLockedLongerDurationNotUnlockingOnly(sdk.WrapSDKContext(suite.Ctx), &types.AccountLockedLongerDurationNotUnlockingOnlyRequest{Owner: addr1.String(), Duration: time.Second})
// // 	suite.Require().NoError(err)
// // 	suite.Require().Len(res.Locks, 0)

// // 	// lock coins
// // 	coins := sdk.Coins{sdk.NewInt64Coin("stake", 10)}
// // 	suite.SetupLock(addr1, coins, time.Second)

// // 	// account locks longer than duration check before start unlocking, duration = 1s
// // 	res, err = suite.QueryServer.AccountLockedLongerDurationNotUnlockingOnly(sdk.WrapSDKContext(suite.Ctx), &types.AccountLockedLongerDurationNotUnlockingOnlyRequest{Owner: addr1.String(), Duration: time.Second})
// // 	suite.Require().NoError(err)
// // 	suite.Require().Len(res.Locks, 1)

// // 	suite.BeginUnlocking(addr1)

// // 	// account locks longer than duration check after start unlocking, duration = 1s
// // 	res, err = suite.QueryServer.AccountLockedLongerDurationNotUnlockingOnly(sdk.WrapSDKContext(suite.Ctx), &types.AccountLockedLongerDurationNotUnlockingOnlyRequest{Owner: addr1.String(), Duration: time.Second})
// // 	suite.Require().NoError(err)
// // 	suite.Require().Len(res.Locks, 0)
// // }

// // func (suite *KeeperTestSuite) TestAccountLockedLongerDurationDenom() {
// // 	suite.SetupTest()
// // 	addr1 := sdk.AccAddress([]byte("addr1---------------"))

// // 	// empty address locks longer than duration by denom check
// // 	res, err := suite.QueryServer.AccountLockedLongerDurationDenom(sdk.WrapSDKContext(suite.Ctx), &types.AccountLockedLongerDurationDenomRequest{Owner: "", Duration: time.Second, Denom: "stake"})
// // 	suite.Require().Error(err)

// // 	// initial check
// // 	res, err = suite.QueryServer.AccountLockedLongerDurationDenom(sdk.WrapSDKContext(suite.Ctx), &types.AccountLockedLongerDurationDenomRequest{Owner: addr1.String(), Duration: time.Second, Denom: "stake"})
// // 	suite.Require().NoError(err)
// // 	suite.Require().Len(res.Locks, 0)

// // 	// lock coins
// // 	coins := sdk.Coins{sdk.NewInt64Coin("stake", 10)}
// // 	suite.SetupLock(addr1, coins, time.Second)
// // 	suite.BeginUnlocking(addr1)

// // 	// account locks longer than duration check by denom, duration = 0s
// // 	res, err = suite.QueryServer.AccountLockedLongerDurationDenom(sdk.WrapSDKContext(suite.Ctx), &types.AccountLockedLongerDurationDenomRequest{Owner: addr1.String(), Duration: 0, Denom: "stake"})
// // 	suite.Require().NoError(err)
// // 	suite.Require().Len(res.Locks, 1)

// // 	// account locks longer than duration check by denom, duration = 1s
// // 	res, err = suite.QueryServer.AccountLockedLongerDurationDenom(sdk.WrapSDKContext(suite.Ctx), &types.AccountLockedLongerDurationDenomRequest{Owner: addr1.String(), Duration: time.Second, Denom: "stake"})
// // 	suite.Require().NoError(err)
// // 	suite.Require().Len(res.Locks, 1)

// // 	// account locks longer than duration check by not available denom, duration = 1s
// // 	res, err = suite.QueryServer.AccountLockedLongerDurationDenom(sdk.WrapSDKContext(suite.Ctx), &types.AccountLockedLongerDurationDenomRequest{Owner: addr1.String(), Duration: time.Second, Denom: "stake2"})
// // 	suite.Require().NoError(err)
// // 	suite.Require().Len(res.Locks, 0)

// // 	// account locks longer than duration check by denom, duration = 2s
// // 	res, err = suite.QueryServer.AccountLockedLongerDurationDenom(sdk.WrapSDKContext(suite.Ctx), &types.AccountLockedLongerDurationDenomRequest{Owner: addr1.String(), Duration: 2 * time.Second, Denom: "stake"})
// // 	suite.Require().NoError(err)
// // 	suite.Require().Len(res.Locks, 0)

// // 	// try querying with prefix coins like "stak" for potential attack
// // 	res, err = suite.QueryServer.AccountLockedLongerDurationDenom(sdk.WrapSDKContext(suite.Ctx), &types.AccountLockedLongerDurationDenomRequest{Owner: addr1.String(), Duration: 0, Denom: "sta"})
// // 	suite.Require().NoError(err)
// // 	suite.Require().Len(res.Locks, 0)
// // }

// // func (suite *KeeperTestSuite) TestLockedDenom() {
// // 	suite.SetupTest()
// // 	addr1 := sdk.AccAddress([]byte("addr1---------------"))

// // 	testTotalLockedDuration := func(durationStr string, expectedAmount int64) {
// // 		duration, _ := time.ParseDuration(durationStr)
// // 		res, err := suite.QueryServer.LockedDenom(
// // 			sdk.WrapSDKContext(suite.Ctx),
// // 			&types.LockedDenomRequest{Denom: "stake", Duration: duration})
// // 		suite.Require().NoError(err)
// // 		suite.Require().Equal(res.Amount, sdk.NewInt(expectedAmount))
// // 	}

// // 	// lock coins
// // 	coins := sdk.Coins{sdk.NewInt64Coin("stake", 10)}
// // 	suite.SetupLock(addr1, coins, time.Hour)

// // 	// test with single lockup
// // 	testTotalLockedDuration("0s", 10)
// // 	testTotalLockedDuration("30m", 10)
// // 	testTotalLockedDuration("1h", 10)
// // 	testTotalLockedDuration("2h", 0)

// // 	// adds different account and lockup for testing
// // 	addr2 := sdk.AccAddress([]byte("addr2---------------"))

// // 	coins = sdk.Coins{sdk.NewInt64Coin("stake", 20)}
// // 	suite.SetupLock(addr2, coins, time.Hour*2)

// // 	testTotalLockedDuration("30m", 30)
// // 	testTotalLockedDuration("1h", 30)
// // 	testTotalLockedDuration("2h", 20)

// // 	// test unlocking
// // 	suite.BeginUnlocking(addr2)
// // 	testTotalLockedDuration("2h", 20)

// // 	// finish unlocking
// // 	duration, _ := time.ParseDuration("2h")
// // 	suite.Ctx = suite.Ctx.WithBlockTime(suite.Ctx.BlockTime().Add(duration))
// // 	suite.WithdrawAllMaturedLocks()
// // 	testTotalLockedDuration("2h", 0)
// // 	testTotalLockedDuration("1h", 10)
// // }
