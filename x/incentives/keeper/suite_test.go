package keeper_test

import (
	"time"

	dexkeeper "github.com/duality-labs/duality/x/dex/keeper"
	dextypes "github.com/duality-labs/duality/x/dex/types"
	"github.com/duality-labs/duality/x/incentives/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type userLocks struct {
	lockAmounts []sdk.Coins
}

type depositSpec struct {
	addr   sdk.AccAddress
	token0 sdk.Coin
	token1 sdk.Coin
	tick   int64
	fee    uint64
}

type gaugeSpec struct {
	isPerpetual bool
	rewards     sdk.Coins
	paidOver    uint64
	startTick   int64
	endTick     int64
	pricingTick int64
}

// AddToGauge adds coins to the specified gauge.
func (suite *KeeperTestSuite) AddToGauge(coins sdk.Coins, gaugeID uint64) uint64 {
	addr := sdk.AccAddress([]byte("addrx---------------"))
	suite.FundAcc(addr, coins)
	err := suite.App.IncentivesKeeper.AddToGaugeRewards(suite.Ctx, addr, coins, gaugeID)
	suite.Require().NoError(err)
	return gaugeID
}

func (suite *KeeperTestSuite) SetupDeposit(addr sdk.AccAddress, token0 sdk.Coin, token1 sdk.Coin, tick int64, fee uint64) sdk.Coins {
	suite.FundAcc(addr, sdk.Coins{token0, token1})
	_, _, shares, err := suite.App.DexKeeper.DepositCore(
		sdk.WrapSDKContext(suite.Ctx),
		token0.Denom,
		token1.Denom,
		addr,
		addr,
		[]sdk.Int{token0.Amount},
		[]sdk.Int{token1.Amount},
		[]int64{tick},
		[]uint64{fee},
		[]*dextypes.DepositOptions{{}},
	)
	suite.Require().NoError(err)
	suite.Require().NotEmpty(shares)
	return shares
}

// LockTokens locks tokens for the specified duration
func (suite *KeeperTestSuite) SetupLock(
	addr sdk.AccAddress,
	shares sdk.Coins,
) *types.Lock {
	lock, err := suite.App.IncentivesKeeper.CreateLock(suite.Ctx, addr, shares, 24*time.Hour)
	suite.Require().NoError(err)
	return lock
}

func GetQualifyingDenom(qc types.QueryCondition) *dexkeeper.DepositDenom {
	tick := qc.StartTick + (qc.EndTick-qc.StartTick)/2
	fee := (qc.EndTick - qc.StartTick) / 3
	return dexkeeper.NewDepositDenom(
		qc.PairID,
		tick,
		uint64(fee),
	)
}

// setupNewGauge creates a gauge with the specified duration.
func (suite *KeeperTestSuite) SetupGauge(
	isPerpetual bool,
	coins sdk.Coins,
	paidOver uint64,
	startTick int64,
	endTick int64,
	pricingTick int64,
) *types.Gauge {
	addr := sdk.AccAddress([]byte("Gauge_Creation_Addr_"))
	distrTo := types.QueryCondition{
		PairID: &dextypes.PairID{
			Token0: "TokenA",
			Token1: "TokenB",
		},
		StartTick: startTick,
		EndTick:   endTick,
	}

	if isPerpetual {
		paidOver = uint64(1)
	}

	// fund reward tokens
	suite.FundAcc(addr, coins)

	// create gauge
	gauge, err := suite.App.IncentivesKeeper.CreateGauge(
		suite.Ctx,
		isPerpetual,
		addr,
		coins,
		distrTo,
		suite.Ctx.BlockTime(),
		paidOver,
		pricingTick,
	)
	suite.Require().NoError(err)
	return gauge
}
