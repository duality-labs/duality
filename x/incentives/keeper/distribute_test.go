package keeper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/app/apptesting"
	dextypes "github.com/duality-labs/duality/x/dex/types"
	"github.com/duality-labs/duality/x/incentives/types"
)

var _ = suite.TestingSuite(nil)

type balanceAssertion struct {
	addr     sdk.AccAddress
	balances sdk.Coins
}

func (suite *KeeperTestSuite) TestValueForShares() {
	addrs := apptesting.SetupAddrs(3)

	tests := []struct {
		name        string
		deposits    []depositSpec
		coin        sdk.Coin
		tick        int64
		expectation sdk.Int
		err         error
	}{
		// gauge 1 gives 3k coins. three stakes, all eligible. 1k coins per stake.
		// 1k should go to oneStakeupUser and 2k to twoStakeupUser.
		{
			name: "one deposit",
			deposits: []depositSpec{
				{
					addr:   addrs[0],
					token0: sdk.NewInt64Coin("TokenA", 10),
					token1: sdk.NewInt64Coin("TokenB", 10),
					tick:   0,
					fee:    1,
				},
			},
			coin: sdk.NewInt64Coin(
				dextypes.NewDepositDenom(&dextypes.PairID{Token0: "TokenA", Token1: "TokenB"}, 0, 1).String(),
				20,
			),
			tick:        1000,
			expectation: sdk.NewInt(21),
		},
		{
			name: "one deposit: no adjustment",
			deposits: []depositSpec{
				{
					addr:   addrs[0],
					token0: sdk.NewInt64Coin("TokenA", 10),
					token1: sdk.NewInt64Coin("TokenB", 10),
					tick:   0,
					fee:    1,
				},
			},
			coin: sdk.NewInt64Coin(
				dextypes.NewDepositDenom(&dextypes.PairID{Token0: "TokenA", Token1: "TokenB"}, 0, 1).String(),
				20,
			),
			tick:        0,
			expectation: sdk.NewInt(20),
		},
		{
			name: "two deposits: one extraneous",
			deposits: []depositSpec{
				{
					addr:   addrs[0],
					token0: sdk.NewInt64Coin("TokenA", 10),
					token1: sdk.NewInt64Coin("TokenB", 10),
					tick:   0,
					fee:    1,
				},
				{
					addr:   addrs[0],
					token0: sdk.NewInt64Coin("TokenA", 10),
					token1: sdk.NewInt64Coin("TokenB", 10),
					tick:   0,
					fee:    2,
				},
			},
			coin: sdk.NewInt64Coin(
				dextypes.NewDepositDenom(&dextypes.PairID{Token0: "TokenA", Token1: "TokenB"}, 0, 1).String(),
				20,
			),
			tick:        1000,
			expectation: sdk.NewInt(21),
		},
		{
			name: "two deposits: both relevant",
			deposits: []depositSpec{
				{
					addr:   addrs[0],
					token0: sdk.NewInt64Coin("TokenA", 10),
					token1: sdk.NewInt64Coin("TokenB", 10),
					tick:   0,
					fee:    1,
				},
				{
					addr:   addrs[0],
					token0: sdk.NewInt64Coin("TokenA", 10),
					token1: sdk.NewInt64Coin("TokenB", 10),
					tick:   0,
					fee:    1,
				},
			},
			coin: sdk.NewInt64Coin(
				dextypes.NewDepositDenom(&dextypes.PairID{Token0: "TokenA", Token1: "TokenB"}, 0, 1).String(),
				20,
			),
			tick:        1000,
			expectation: sdk.NewInt(21),
		},
	}
	for _, tc := range tests {
		suite.T().Run(tc.name, func(t *testing.T) {
			suite.SetupTest()
			for _, depositSpec := range tc.deposits {
				suite.SetupDeposit(depositSpec)
			}
			value, err := suite.App.IncentivesKeeper.ValueForShares(suite.Ctx, tc.coin, tc.tick)
			if tc.err == nil {
				require.NoError(t, err)
				require.Equal(t, tc.expectation, value)
			} else {
				require.Error(t, err)
			}
		})
	}
}

// TestDistribute tests that when the distribute command is executed on a provided gauge
// that the correct amount of rewards is sent to the correct stake owners.
func (suite *KeeperTestSuite) TestDistribute() {
	addrs := apptesting.SetupAddrs(3)
	tests := []struct {
		name              string
		addrs             []sdk.AccAddress
		depositStakeSpecs []depositStakeSpec
		gaugeSpecs        []gaugeSpec
		assertions        []balanceAssertion
	}{
		{
			name: "one gauge",
			depositStakeSpecs: []depositStakeSpec{
				{
					depositSpec: depositSpec{
						addr:   addrs[0],
						token0: sdk.NewInt64Coin("TokenA", 10),
						token1: sdk.NewInt64Coin("TokenB", 10),
						tick:   0,
						fee:    1,
					},
					stakeTimeOffset: -24 * time.Hour,
				},
				{
					depositSpec: depositSpec{
						addr:   addrs[1],
						token0: sdk.NewInt64Coin("TokenA", 10),
						token1: sdk.NewInt64Coin("TokenB", 10),
						tick:   0,
						fee:    1,
					},
					stakeTimeOffset: -24 * time.Hour,
				},
				{
					depositSpec: depositSpec{
						addr:   addrs[1],
						token0: sdk.NewInt64Coin("TokenA", 10),
						token1: sdk.NewInt64Coin("TokenB", 10),
						tick:   0,
						fee:    1,
					},
					stakeTimeOffset: -24 * time.Hour,
				},
			},
			gaugeSpecs: []gaugeSpec{
				{
					isPerpetual: false,
					rewards:     sdk.Coins{sdk.NewInt64Coin("reward", 3000)},
					startTick:   -10,
					endTick:     10,
					paidOver:    1,
					pricingTick: 0,
				},
			},
			assertions: []balanceAssertion{
				{addr: addrs[0], balances: sdk.Coins{sdk.NewInt64Coin("reward", 1000)}},
				{addr: addrs[1], balances: sdk.Coins{sdk.NewInt64Coin("reward", 2000)}},
			},
		},
		{
			name: "two gauges",
			depositStakeSpecs: []depositStakeSpec{
				{
					depositSpec: depositSpec{
						addr:   addrs[0],
						token0: sdk.NewInt64Coin("TokenA", 10),
						token1: sdk.NewInt64Coin("TokenB", 10),
						tick:   0,
						fee:    1,
					},
					stakeTimeOffset: -24 * time.Hour,
				},
				{
					depositSpec: depositSpec{
						addr:   addrs[1],
						token0: sdk.NewInt64Coin("TokenA", 10),
						token1: sdk.NewInt64Coin("TokenB", 10),
						tick:   0,
						fee:    1,
					},
					stakeTimeOffset: -24 * time.Hour,
				},
				{
					depositSpec: depositSpec{
						addr:   addrs[1],
						token0: sdk.NewInt64Coin("TokenA", 10),
						token1: sdk.NewInt64Coin("TokenB", 10),
						tick:   0,
						fee:    1,
					},
					stakeTimeOffset: -24 * time.Hour,
				},
				{
					depositSpec: depositSpec{
						addr:   addrs[0],
						token0: sdk.NewInt64Coin("TokenA", 10),
						token1: sdk.NewInt64Coin("TokenB", 10),
						tick:   0,
						fee:    1,
					},
					stakeTimeOffset: -12 * time.Hour,
				},
			},
			gaugeSpecs: []gaugeSpec{
				{
					isPerpetual: false,
					rewards:     sdk.Coins{sdk.NewInt64Coin("reward", 3000)},
					startTick:   -10,
					endTick:     10,
					paidOver:    1,
					pricingTick: 0,
				},
				{
					isPerpetual: false,
					rewards:     sdk.Coins{sdk.NewInt64Coin("reward", 3000)},
					startTick:   -10,
					endTick:     10,
					paidOver:    2,
					pricingTick: 0,
				},
			},
			assertions: []balanceAssertion{
				{addr: addrs[0], balances: sdk.Coins{sdk.NewInt64Coin("reward", 1500)}},
				{addr: addrs[1], balances: sdk.Coins{sdk.NewInt64Coin("reward", 3000)}},
			},
		},
		{
			name: "one stake with adjustment",
			depositStakeSpecs: []depositStakeSpec{
				{
					depositSpec: depositSpec{
						addr:   addrs[0],
						token0: sdk.NewInt64Coin("TokenA", 10),
						token1: sdk.NewInt64Coin("TokenB", 10),
						tick:   999,
						fee:    1,
					},
					stakeTimeOffset: -24 * time.Hour,
				},
				{
					depositSpec: depositSpec{
						addr:   addrs[1],
						token0: sdk.NewInt64Coin("TokenA", 10),
						token1: sdk.NewInt64Coin("TokenB", 10),
						tick:   999,
						fee:    1,
					},
					stakeTimeOffset: -24 * time.Hour,
				},
				{
					depositSpec: depositSpec{
						addr:   addrs[1],
						token0: sdk.NewInt64Coin("TokenA", 10),
						token1: sdk.NewInt64Coin("TokenB", 10),
						tick:   999,
						fee:    40,
					},
					stakeTimeOffset: -24 * time.Hour,
				},
				{
					depositSpec: depositSpec{
						addr:   addrs[1],
						token0: sdk.NewInt64Coin("TokenA", 10),
						token1: sdk.NewInt64Coin("TokenB", 10),
						tick:   999,
						fee:    40,
					},
					stakeTimeOffset: -12 * time.Hour,
				},
			},
			gaugeSpecs: []gaugeSpec{
				{
					isPerpetual: false,
					rewards:     sdk.Coins{sdk.NewInt64Coin("reward", 3000)},
					startTick:   -1000,
					endTick:     1000,
					paidOver:    1,
					pricingTick: 0,
				},
			},
			assertions: []balanceAssertion{
				{addr: addrs[0], balances: sdk.Coins{sdk.NewInt64Coin("reward", 1500)}},
				{addr: addrs[1], balances: sdk.Coins{sdk.NewInt64Coin("reward", 1500)}},
			},
		},
	}
	for _, tc := range tests {
		suite.T().Run(tc.name, func(t *testing.T) {
			suite.SetupTest()
			for _, depositSpec := range tc.depositStakeSpecs {
				suite.SetupDepositAndStake(depositSpec)
			}
			gauges := make(types.Gauges, len(tc.gaugeSpecs))
			for i, gaugeSpec := range tc.gaugeSpecs {
				gauge := suite.SetupGauge(gaugeSpec)
				gauges[i] = gauge
			}
			_, err := suite.App.IncentivesKeeper.Distribute(suite.Ctx, gauges)
			require.NoError(t, err)
			// check expected rewards against actual rewards received
			for i, assertion := range tc.assertions {
				bal := suite.App.BankKeeper.GetAllBalances(suite.Ctx, assertion.addr)
				assert.Equal(t, assertion.balances.String(), bal.String(), "test %v, person %d", tc.name, i)
			}
		})
	}
}
