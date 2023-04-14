package keeper_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/tendermint/tendermint/crypto/secp256k1"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/duality-labs/duality/app"
	dualityapp "github.com/duality-labs/duality/app"
	"github.com/duality-labs/duality/x/incentives/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// genRewardCoins takes coins and returns a randomized coin struct used as rewards for the distribution benchmark.
func genRewardCoins(r *rand.Rand, coins sdk.Coins) (res sdk.Coins) {
	numCoins := 1 + r.Intn(Min(coins.Len(), 1))
	denomIndices := r.Perm(numCoins)
	for i := 0; i < numCoins; i++ {
		denom := coins[denomIndices[i]].Denom
		amt, _ := simtypes.RandPositiveInt(r, coins[i].Amount)
		res = append(res, sdk.Coin{Denom: denom, Amount: amt})
	}

	return
}

// genQueryCondition takes coins and returns a QueryConditon struct.
func genQueryCondition(
	r *rand.Rand,
	bstaketime time.Time,
	coins sdk.Coins,
) types.QueryCondition {
	return types.QueryCondition{}
}

// // benchmarkDistributionLogic creates gauges with stakeups that get distributed to. Benchmarks the performance of the distribution process.
// func benchmarkDistributionLogic(numAccts, numDenoms, numGauges, numStakeups, numDistrs int, b *testing.B) {
// 	b.StopTimer()

// 	bstakeStartTime := time.Now().UTC()
// 	app, cleanupFn := app.SetupTestingAppWithLevelDb(false)
// 	defer cleanupFn()
// 	ctx := app.BaseApp.NewContext(false, tmproto.Header{Height: 1, ChainID: "duality-1", Time: bstakeStartTime})

// 	r := rand.New(rand.NewSource(10))

// 	// setup accounts with balances
// 	addrs := []sdk.AccAddress{}
// 	for i := 0; i < numAccts; i++ {
// 		addr := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
// 		coins := sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 100000000)}
// 		for j := 0; j < numDenoms; j++ {
// 			coins = coins.Add(sdk.NewInt64Coin(fmt.Sprintf("token%d", j), r.Int63n(100000000)))
// 		}
// 		_ = simapp.FundAccount(app.BankKeeper, ctx, addr, coins)
// 		app.AccountKeeper.SetAccount(ctx, authtypes.NewBaseAccount(addr, nil, 0, 0))
// 		addrs = append(addrs, addr)
// 	}

// 	distrEpoch := app.EpochsKeeper.GetEpochInfo(ctx, app.IncentivesKeeper.GetParams(ctx).DistrEpochIdentifier)
// 	// setup gauges
// 	gaugeIds := []uint64{}
// 	for i := 0; i < numGauges; i++ {
// 		addr := addrs[r.Int()%numAccts]
// 		simCoins := app.BankKeeper.SpendableCoins(ctx, addr)

// 		// isPerpetual := r.Int()%2 == 0
// 		isPerpetual := true
// 		distributeTo := genQueryCondition(r, ctx.BlockTime(), simCoins)
// 		rewards := genRewardCoins(r, simCoins)
// 		startTime := ctx.BlockTime().Add(time.Duration(-1) * time.Second)
// 		durationMillisecs := distributeTo.Duration.Milliseconds()
// 		numEpochsPaidOver := uint64(1)
// 		if !isPerpetual {
// 			millisecsPerEpoch := distrEpoch.Duration.Milliseconds()
// 			numEpochsPaidOver = uint64(r.Int63n(durationMillisecs/millisecsPerEpoch)) + 1
// 		}

// 		gaugeId, err := app.IncentivesKeeper.CreateGauge(ctx, isPerpetual, addr, rewards, distributeTo, startTime, numEpochsPaidOver)
// 		if err != nil {
// 			fmt.Printf("Create Gauge, %v\n", err)
// 			b.FailNow()
// 		} else {
// 			gaugeIds = append(gaugeIds, gaugeId)
// 		}
// 	}

// 	// jump time to the future
// 	futureSecs := r.Intn(1 * 60 * 60 * 24 * 7)
// 	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(time.Duration(futureSecs) * time.Second))

// 	stakeSecs := r.Intn(1 * 60 * 60 * 8)
// 	// setup stakeups
// 	for i := 0; i < numStakeups; i++ {
// 		addr := addrs[i%numAccts]
// 		simCoins := app.BankKeeper.SpendableCoins(ctx, addr)

// 		if i%10 == 0 {
// 			stakeSecs = r.Intn(1 * 60 * 60 * 8)
// 		}
// 		duration := time.Duration(stakeSecs) * time.Second
// 		_, err := app.IncentivesKeeper.CreateStake(ctx, addr, simCoins, duration)
// 		if err != nil {
// 			fmt.Printf("Stake tokens, %v\n", err)
// 			b.FailNow()
// 		}
// 	}
// 	fmt.Println("created all stakeups")

// 	// begin distribution for all gauges
// 	for _, gaugeId := range gaugeIds {
// 		gauge, _ := app.IncentivesKeeper.GetGaugeByID(ctx, gaugeId)
// 		err := app.IncentivesKeeper.MoveUpcomingGaugeToActiveGauge(ctx, *gauge)
// 		if err != nil {
// 			fmt.Printf("Begin distribution, %v\n", err)
// 			b.FailNow()
// 		}
// 	}

// 	b.StartTimer()
// 	// distribute coins from gauges to stakeup owners
// 	for i := 0; i < numDistrs; i++ {
// 		gauges := types.Gauges{}
// 		for _, gaugeId := range gaugeIds {
// 			gauge, _ := app.IncentivesKeeper.GetGaugeByID(ctx, gaugeId)
// 			gauges = append(gauges, *gauge)
// 		}
// 		_, err := app.IncentivesKeeper.Distribute(ctx, gauges)
// 		if err != nil {
// 			b.FailNow()
// 		}
// 	}
// }

// func BenchmarkDistributionLogicTiny(b *testing.B) {
// 	numAccts := 1
// 	numDenoms := 1
// 	numGauges := 1
// 	numStakeups := 1
// 	numDistrs := 1
// 	benchmarkDistributionLogic(numAccts, numDenoms, numGauges, numStakeups, numDistrs, b)
// }

// func BenchmarkDistributionLogicSmall(b *testing.B) {
// 	numAccts := 10
// 	numDenoms := 1
// 	numGauges := 10
// 	numStakeups := 1000
// 	numDistrs := 100
// 	benchmarkDistributionLogic(numAccts, numDenoms, numGauges, numStakeups, numDistrs, b)
// }

// func BenchmarkDistributionLogicMedium(b *testing.B) {
// 	numAccts := 1000
// 	numDenoms := 8
// 	numGauges := 30
// 	numStakeups := 20000
// 	numDistrs := 1

// 	benchmarkDistributionLogic(numAccts, numDenoms, numGauges, numStakeups, numDistrs, b)
// }

// func BenchmarkDistributionLogicLarge(b *testing.B) {
// 	numAccts := 50000
// 	numDenoms := 10
// 	numGauges := 60
// 	numStakeups := 100000
// 	numDistrs := 1

// 	benchmarkDistributionLogic(numAccts, numDenoms, numGauges, numStakeups, numDistrs, b)
// }

// func BenchmarkDistributionLogicHuge(b *testing.B) {
// 	numAccts := 1000
// 	numDenoms := 100
// 	numGauges := 1000
// 	numStakeups := 1000
// 	numDistrs := 30000
// 	benchmarkDistributionLogic(numAccts, numDenoms, numGauges, numStakeups, numDistrs, b)
// }

// from stakeup

func benchmarkResetLogic(numStakeups int, b *testing.B) {
	// b.ReportAllocs()
	b.StopTimer()

	bstakeStartTime := time.Now().UTC()
	app := app.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Height: 1, ChainID: "duality-1", Time: bstakeStartTime})

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	numAccts := 100
	numDenoms := 1

	denom := fmt.Sprintf("token%d", 0)

	// setup accounts with balances
	addrs := []sdk.AccAddress{}
	for i := 0; i < numAccts; i++ {
		addr := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
		coins := sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 100000000)}
		for j := 0; j < numDenoms; j++ {
			coins = coins.Add(sdk.NewInt64Coin(fmt.Sprintf("token%d", j), r.Int63n(100000000)))
		}
		_ = dualityapp.FundAccount(app.BankKeeper, ctx, addr, coins)
		app.AccountKeeper.SetAccount(ctx, authtypes.NewBaseAccount(addr, nil, 0, 0))
		addrs = append(addrs, addr)
	}

	// jump time to the future
	futureSecs := r.Intn(1 * 60 * 60 * 24 * 7)
	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(time.Duration(futureSecs) * time.Second))

	stakes := make(types.Stakes, numStakeups)
	// setup stakeups
	for i := 0; i < numStakeups; i++ {
		addr := addrs[r.Int()%numAccts]
		simCoins := sdk.NewCoins(sdk.NewCoin(denom, sdk.NewInt(r.Int63n(100))))
		// duration := time.Duration(r.Intn(1*60*60*24*7)) * time.Second
		stake := types.NewStake(uint64(i+1), addr, simCoins, ctx.BlockTime())
		stakes[i] = stake
	}

	b.StartTimer()
	b.ReportAllocs()
	// distribute coins from gauges to stakeup owners
	_ = app.IncentivesKeeper.InitializeAllStakes(ctx, stakes)
}

func BenchmarkResetLogicMedium(b *testing.B) {
	benchmarkResetLogic(50000, b)
}
