package keeper_test

import (
	"testing"
	time "time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/app"
	dextypes "github.com/duality-labs/duality/x/dex/types"
	. "github.com/duality-labs/duality/x/incentives/keeper"
	"github.com/duality-labs/duality/x/incentives/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	tmtypes "github.com/tendermint/tendermint/proto/tendermint/types"
)

type MockKeeper struct {
	locks types.Locks
}

func NewMockKeeper(locks types.Locks) MockKeeper {
	return MockKeeper{
		locks: locks,
	}
}

func (k MockKeeper) ValueForShares(ctx sdk.Context, coin sdk.Coin, tick int64) (sdk.Int, error) {
	return coin.Amount.Mul(sdk.NewInt(2)), nil
}

func (k MockKeeper) GetLocksByQueryCondition(ctx sdk.Context, distrTo *types.QueryCondition) types.Locks {
	return k.locks
}

func TestDistributor(t *testing.T) {
	app := app.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmtypes.Header{Height: 1, ChainID: "duality-1", Time: time.Now().UTC()})

	coins := sdk.Coins{sdk.NewCoin("coin1", sdk.NewInt(100))}
	gauge := types.NewGauge(
		1,
		false,
		types.QueryCondition{
			PairID: &dextypes.PairID{
				Token0: "TokenA",
				Token1: "TokenB",
			},
			StartTick: -10,
			EndTick:   10},
		coins,
		ctx.BlockTime(),
		10,
		0,
		sdk.Coins{},
		0,
	)
	rewardedDenom := dextypes.NewDepositDenom(&dextypes.PairID{Token0: "TokenA", Token1: "TokenB"}, 5, 1).String()
	nonRewardedDenom := dextypes.NewDepositDenom(&dextypes.PairID{Token0: "TokenA", Token1: "TokenB"}, 12, 1).String()
	allLocks := types.Locks{
		{1, "addr1", time.Minute, time.Time{}, sdk.Coins{sdk.NewCoin(rewardedDenom, sdk.NewInt(50))}},
		{2, "addr2", time.Minute, time.Time{}, sdk.Coins{sdk.NewCoin(rewardedDenom, sdk.NewInt(25))}},
		{3, "addr2", time.Minute, time.Time{}, sdk.Coins{sdk.NewCoin(rewardedDenom, sdk.NewInt(25))}},
		{4, "addr3", time.Minute, time.Time{}, sdk.Coins{sdk.NewCoin(nonRewardedDenom, sdk.NewInt(50))}},
	}

	distributor := NewDistributor(NewMockKeeper(allLocks))

	testCases := []struct {
		name        string
		timeOffset  time.Duration
		filterLocks types.Locks
		expected    types.DistributionSpec
		expectedErr error
	}{
		{
			name:        "Error case: gauge not active",
			timeOffset:  -1 * time.Minute,
			filterLocks: allLocks,
			expected:    nil,
			expectedErr: types.ErrGaugeNotActive,
		},
		{
			name:        "Successful case: distribute to all locks",
			timeOffset:  0,
			filterLocks: allLocks,
			expected: types.DistributionSpec{
				"addr1": sdk.Coins{sdk.NewCoin("coin1", sdk.NewInt(5))},
				"addr2": sdk.Coins{sdk.NewCoin("coin1", sdk.NewInt(4))},
			},
			expectedErr: nil,
		},
		{
			name:        "Successful case: distribute to one lock",
			timeOffset:  0,
			filterLocks: types.Locks{allLocks[0]},
			expected: types.DistributionSpec{
				"addr1": sdk.Coins{sdk.NewCoin("coin1", sdk.NewInt(5))},
			},
			expectedErr: nil,
		},
		{
			name:        "No distribution: empty filterLocks",
			timeOffset:  0,
			filterLocks: types.Locks{},
			expected:    types.DistributionSpec{},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			distSpec, err := distributor.Distribute(
				ctx.WithBlockTime(ctx.BlockTime().Add(tc.timeOffset)),
				&gauge,
				tc.filterLocks,
			)
			if tc.expectedErr != nil {
				assert.Equal(t, tc.expectedErr, err)
			} else {
				require.NoError(t, err)
			}
			assert.Equal(t, tc.expected, distSpec)
		})
	}
}
