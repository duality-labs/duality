package keeper_test

// var (
// 	now         = time.Now().UTC()
// 	acc1        = sdk.AccAddress([]byte("addr1---------------"))
// 	acc2        = sdk.AccAddress([]byte("addr2---------------"))
// 	testGenesis = types.GenesisState{
// 		LastLockId: 10,
// 		Locks: types.Locks{
// 			{
// 				ID:      1,
// 				Owner:   acc1.String(),
// 				EndTime: time.Time{},
// 				Coins:   sdk.Coins{sdk.NewInt64Coin("foo", 10000000)},
// 			},
// 			{
// 				ID:      2,
// 				Owner:   acc1.String(),
// 				EndTime: time.Time{},
// 				Coins:   sdk.Coins{sdk.NewInt64Coin("foo", 15000000)},
// 			},
// 			{
// 				ID:      3,
// 				Owner:   acc2.String(),
// 				EndTime: time.Time{},
// 				Coins:   sdk.Coins{sdk.NewInt64Coin("foo", 5000000)},
// 			},
// 		},
// 	}
// )

// // TestIncentivesExportGenesis tests export genesis command for the incentives module.
// func TestIncentivesExportGenesis(t *testing.T) {
// 	// export genesis using default configurations
// 	// ensure resulting genesis params match default params
// 	app := dualityapp.Setup(false)
// 	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
// 	genesis := app.IncentivesKeeper.ExportGenesis(ctx)
// 	require.Equal(t, genesis.Params.DistrEpochIdentifier, "week")
// 	require.Len(t, genesis.types.Gauges, 0)

// 	// create an address and fund with coins
// 	addr := sdk.AccAddress([]byte("addr1---------------"))
// 	coins := sdk.Coins{sdk.NewInt64Coin("stake", 10000)}
// 	err := simapp.FundAccount(app.BankKeeper, ctx, addr, coins)
// 	require.NoError(t, err)

// 	// mints LP tokens and send to address created earlier
// 	// this ensures the supply exists on chain
// 	distrTo := types.QueryCondition{
// 		// TODO
// 		// Denom:         "lptoken",
// 		// Duration:      time.Second,
// 	}
// 	mintLPtokens := sdk.Coins{sdk.NewInt64Coin(distrTo.Denom, 200)}
// 	err = simapp.FundAccount(app.BankKeeper, ctx, addr, mintLPtokens)
// 	require.NoError(t, err)

// 	// create a gauge that distributes coins to earlier created LP token and duration
// 	startTime := time.Now()
// 	gaugeID, err := app.IncentivesKeeper.CreateGauge(ctx, true, addr, coins, distrTo, startTime, 1)
// 	require.NoError(t, err)

// 	// export genesis using default configurations
// 	// ensure resulting genesis params match default params
// 	genesis = app.IncentivesKeeper.ExportGenesis(ctx)
// 	require.Equal(t, genesis.Params.DistrEpochIdentifier, "week")
// 	require.Len(t, genesis.types.Gauges, 1)

// 	// ensure the first gauge listed in the exported genesis explicitly matches expectation
// 	require.Equal(t, genesis.types.Gauges[0], types.Gauge{
// 		Id:                gaugeID,
// 		IsPerpetual:       true,
// 		DistributeTo:      distrTo,
// 		Coins:             coins,
// 		NumEpochsPaidOver: 1,
// 		FilledEpochs:      0,
// 		DistributedCoins:  sdk.Coins(nil),
// 		StartTime:         startTime.UTC(),
// 	})
// }

// // TestIncentivesInitGenesis takes a genesis state and tests initializing that genesis for the incentives module.
// func TestIncentivesInitGenesis(t *testing.T) {
// 	app := dualityapp.Setup(false)
// 	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

// 	// checks that the default genesis parameters pass validation
// 	validateGenesis := types.DefaultGenesis().Params.Validate()
// 	require.NoError(t, validateGenesis)

// 	// create coins, lp tokens with lockup durations, and a gauge for this lockup
// 	coins := sdk.Coins{sdk.NewInt64Coin("stake", 10000)}
// 	startTime := time.Now()
// 	distrTo := types.QueryCondition{
// 		// TODO
// 		// Denom:         "lptoken",
// 		// Duration:      time.Second,
// 	}
// 	gauge := types.Gauge{
// 		Id:                1,
// 		IsPerpetual:       false,
// 		DistributeTo:      distrTo,
// 		Coins:             coins,
// 		NumEpochsPaidOver: 2,
// 		FilledEpochs:      0,
// 		DistributedCoins:  sdk.Coins(nil),
// 		StartTime:         startTime.UTC(),
// 	}

// 	// initialize genesis with specified parameter, the gauge created earlier, and lockable durations
// 	app.IncentivesKeeper.InitGenesis(ctx, types.GenesisState{
// 		Params: types.Params{
// 			DistrEpochIdentifier: "week",
// 		},
// 		types.Gauges: types.Gauges{gauge},
// 	})

// 	// check that the gauge created earlier was initialized through initGenesis and still exists on chain
// 	gauges := app.IncentivesKeeper.GetGauges(ctx)
// 	require.Len(t, gauges, 1)
// 	require.Equal(t, gauges[0], gauge)
// }

// func TestInitGenesis(t *testing.T) {
// 	app := dualityapp.Setup(false)
// 	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
// 	ctx = ctx.WithBlockTime(now.Add(time.Second))
// 	genesis := testGenesis
// 	app.IncentivesKeeper.InitGenesis(ctx, genesis)

// 	coins := app.IncentivesKeeper.GetAccountLockedCoins(ctx, acc1)
// 	require.Equal(t, coins.String(), sdk.NewInt64Coin("foo", 25000000).String())

// 	coins = app.IncentivesKeeper.GetAccountLockedCoins(ctx, acc2)
// 	require.Equal(t, coins.String(), sdk.NewInt64Coin("foo", 5000000).String())

// 	lastLockId := app.IncentivesKeeper.GetLastLockID(ctx)
// 	require.Equal(t, lastLockId, uint64(10))

// 	// acc := app.IncentivesKeeper.GetLocksAccumulation(ctx, types.QueryCondition{
// 	// 	// TODO
// 	// 	// Denom:    "foo",
// 	// 	// Duration: time.Second,
// 	// })
// 	// require.Equal(t, sdk.NewInt(30000000), acc)
// }

// func TestExportGenesis(t *testing.T) {
// 	app := dualityapp.Setup(false)
// 	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
// 	ctx = ctx.WithBlockTime(now.Add(time.Second))
// 	genesis := testGenesis
// 	app.IncentivesKeeper.InitGenesis(ctx, genesis)

// 	err := dualityapp.FundAccount(app.BankKeeper, ctx, acc2, sdk.Coins{sdk.NewInt64Coin("foo", 5000000)})
// 	require.NoError(t, err)
// 	_, err = app.IncentivesKeeper.CreateLock(ctx, acc2, sdk.Coins{sdk.NewInt64Coin("foo", 5000000)}, time.Second*5)
// 	require.NoError(t, err)

// 	coins := app.IncentivesKeeper.GetAccountLockedCoins(ctx, acc2)
// 	require.Equal(t, coins.String(), sdk.NewInt64Coin("foo", 10000000).String())

// 	genesisExported := app.IncentivesKeeper.ExportGenesis(ctx)
// 	require.Equal(t, genesisExported.LastLockId, uint64(11))
// 	require.Equal(t, genesisExported.Locks, types.Locks{
// 		{
// 			ID:      1,
// 			Owner:   acc1.String(),
// 			EndTime: time.Time{},
// 			Coins:   sdk.Coins{sdk.NewInt64Coin("foo", 10000000)},
// 		},
// 		{
// 			ID:      11,
// 			Owner:   acc2.String(),
// 			EndTime: time.Time{},
// 			Coins:   sdk.Coins{sdk.NewInt64Coin("foo", 5000000)},
// 		},
// 		{
// 			ID:      3,
// 			Owner:   acc2.String(),
// 			EndTime: time.Time{},
// 			Coins:   sdk.Coins{sdk.NewInt64Coin("foo", 5000000)},
// 		},
// 		{
// 			ID:      2,
// 			Owner:   acc1.String(),
// 			EndTime: time.Time{},
// 			Coins:   sdk.Coins{sdk.NewInt64Coin("foo", 15000000)},
// 		},
// 	})
// }

// func TestMarshalUnmarshalGenesis(t *testing.T) {
// 	app := dualityapp.Setup(false)
// 	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
// 	ctx = ctx.WithBlockTime(now.Add(time.Second))

// 	encodingConfig := dualityapp.MakeTestEncodingConfig()
// 	appCodec := encodingConfig.Marshaler
// 	am := incentives.NewAppModule(app.IncentivesKeeper, app.AccountKeeper, app.BankKeeper, app.EpochsKeeper)

// 	err := dualityapp.FundAccount(app.BankKeeper, ctx, acc2, sdk.Coins{sdk.NewInt64Coin("foo", 5000000)})
// 	require.NoError(t, err)
// 	_, err = app.IncentivesKeeper.CreateLock(ctx, acc2, sdk.Coins{sdk.NewInt64Coin("foo", 5000000)}, time.Second*5)
// 	require.NoError(t, err)

// 	genesisExported := am.ExportGenesis(ctx, appCodec)
// 	assert.NotPanics(t, func() {
// 		app := dualityapp.Setup(false)
// 		ctx := app.BaseApp.NewContext(false, tmproto.Header{})
// 		ctx = ctx.WithBlockTime(now.Add(time.Second))
// 		am := incentives.NewAppModule(app.IncentivesKeeper, app.AccountKeeper, app.BankKeeper, app.EpochsKeeper)
// 		am.InitGenesis(ctx, appCodec, genesisExported)
// 	})
// }
