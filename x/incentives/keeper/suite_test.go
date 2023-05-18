package keeper_test

import (
	"time"

	dextypes "github.com/duality-labs/duality/x/dex/types"
	"github.com/duality-labs/duality/x/incentives/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type userStakes struct {
	stakeAmounts []sdk.Coins
}

type depositStakeSpec struct {
	depositSpec     depositSpec
	stakeTimeOffset time.Duration // used for simulating the time of staking
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
	startTime   time.Time
}

// AddToGauge adds coins to the specified gauge.
func (suite *KeeperTestSuite) AddToGauge(coins sdk.Coins, gaugeID uint64) uint64 {
	addr := sdk.AccAddress([]byte("addrx---------------"))
	suite.FundAcc(addr, coins)
	err := suite.App.IncentivesKeeper.AddToGaugeRewards(suite.Ctx, addr, coins, gaugeID)
	suite.Require().NoError(err)
	return gaugeID
}

func (suite *KeeperTestSuite) SetupDeposit(s depositSpec) sdk.Coins {
	suite.FundAcc(s.addr, sdk.Coins{s.token0, s.token1})
	_, _, shares, err := suite.App.DexKeeper.DepositCore(
		sdk.WrapSDKContext(suite.Ctx),
		dextypes.MustNewPairID(s.token0.Denom, s.token1.Denom),
		s.addr,
		s.addr,
		[]sdk.Int{s.token0.Amount},
		[]sdk.Int{s.token1.Amount},
		[]int64{s.tick},
		[]uint64{s.fee},
		[]*dextypes.DepositOptions{{}},
	)
	suite.Require().NoError(err)
	suite.Require().NotEmpty(shares)
	return shares
}

func (suite *KeeperTestSuite) SetupDepositAndStake(s depositStakeSpec) *types.Stake {
	shares := suite.SetupDeposit(s.depositSpec)
	return suite.SetupStake(s.depositSpec.addr, shares, s.stakeTimeOffset)
}

// StakeTokens stakes tokens for the specified duration
func (suite *KeeperTestSuite) SetupStake(
	addr sdk.AccAddress,
	shares sdk.Coins,
	timeOffset time.Duration,
) *types.Stake {
	stake, err := suite.App.IncentivesKeeper.CreateStake(suite.Ctx, addr, shares, suite.Ctx.BlockTime().Add(timeOffset))
	suite.Require().NoError(err)
	return stake
}

func GetQualifyingDenom(qc types.QueryCondition) *dextypes.DepositDenom {
	tick := qc.StartTick + (qc.EndTick-qc.StartTick)/2
	fee := (qc.EndTick - qc.StartTick) / 3
	return dextypes.NewDepositDenom(
		qc.PairID,
		tick,
		uint64(fee),
	)
}

// setupNewGauge creates a gauge with the specified duration.
func (suite *KeeperTestSuite) SetupGauge(s gaugeSpec) *types.Gauge {
	addr := sdk.AccAddress([]byte("Gauge_Creation_Addr_"))

	// fund reward tokens
	suite.FundAcc(addr, s.rewards)

	// create gauge
	gauge, err := suite.App.IncentivesKeeper.CreateGauge(
		suite.Ctx,
		s.isPerpetual,
		addr,
		s.rewards,
		types.QueryCondition{
			PairID: &dextypes.PairID{
				Token0: "TokenA",
				Token1: "TokenB",
			},
			StartTick: s.startTick,
			EndTick:   s.endTick,
		},
		s.startTime,
		s.paidOver,
		s.pricingTick,
	)
	suite.Require().NoError(err)
	return gauge
}

func (suite *KeeperTestSuite) SetupGauges(specs []gaugeSpec) {
	for _, s := range specs {
		suite.SetupGauge(s)
	}
}
