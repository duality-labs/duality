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
	tmtypes "github.com/cometbft/cometbft/proto/tendermint/types"
)

var _ DistributorKeeper = MockKeeper{}

type MockKeeper struct {
	stakes types.Stakes
}

func NewMockKeeper(stakes types.Stakes) MockKeeper {
	return MockKeeper{
		stakes: stakes,
	}
}

func (k MockKeeper) ValueForShares(ctx sdk.Context, coin sdk.Coin, tick int64) (sdk.Int, error) {
	return coin.Amount.Mul(sdk.NewInt(2)), nil
}

func (k MockKeeper) GetStakesByQueryCondition(ctx sdk.Context, distrTo *types.QueryCondition) types.Stakes {
	return k.stakes
}

func TestDistributor(t *testing.T) {
	app := app.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmtypes.Header{Height: 1, ChainID: "duality-1", Time: time.Now().UTC()})

	gauge := types.NewGauge(
		1,
		false,
		types.QueryCondition{
			PairID: &dextypes.PairID{
				Token0: "TokenA",
				Token1: "TokenB",
			},
			StartTick: -10,
			EndTick:   10,
		},
		sdk.Coins{sdk.NewCoin("coin1", sdk.NewInt(100))},
		ctx.BlockTime(),
		10,
		0,
		sdk.Coins{},
		0,
	)
	rewardedDenom := dextypes.NewDepositDenom(&dextypes.PairID{Token0: "TokenA", Token1: "TokenB"}, 5, 1).String()
	nonRewardedDenom := dextypes.NewDepositDenom(&dextypes.PairID{Token0: "TokenA", Token1: "TokenB"}, 12, 1).String()
	allStakes := types.Stakes{
		{1, "addr1", ctx.BlockTime(), sdk.Coins{sdk.NewCoin(rewardedDenom, sdk.NewInt(50))}},
		{2, "addr2", ctx.BlockTime(), sdk.Coins{sdk.NewCoin(rewardedDenom, sdk.NewInt(25))}},
		{3, "addr2", ctx.BlockTime(), sdk.Coins{sdk.NewCoin(rewardedDenom, sdk.NewInt(25))}},
		{4, "addr3", ctx.BlockTime(), sdk.Coins{sdk.NewCoin(nonRewardedDenom, sdk.NewInt(50))}},
	}

	distributor := NewDistributor(NewMockKeeper(allStakes))

	testCases := []struct {
		name         string
		timeOffset   time.Duration
		filterStakes types.Stakes
		expected     types.DistributionSpec
		expectedErr  error
	}{
		{
			name:         "Error case: gauge not active",
			timeOffset:   -1 * time.Minute,
			filterStakes: allStakes,
			expected:     nil,
			expectedErr:  types.ErrGaugeNotActive,
		},
		{
			name:         "Successful case: distribute to all stakes",
			timeOffset:   0,
			filterStakes: allStakes,
			expected: types.DistributionSpec{
				"addr1": sdk.Coins{sdk.NewCoin("coin1", sdk.NewInt(5))},
				"addr2": sdk.Coins{sdk.NewCoin("coin1", sdk.NewInt(4))},
			},
			expectedErr: nil,
		},
		{
			name:         "Successful case: distribute to one stake",
			timeOffset:   0,
			filterStakes: types.Stakes{allStakes[0]},
			expected: types.DistributionSpec{
				"addr1": sdk.Coins{sdk.NewCoin("coin1", sdk.NewInt(5))},
			},
			expectedErr: nil,
		},
		{
			name:         "No distribution: empty filterStakes",
			filterStakes: types.Stakes{},
			expected:     types.DistributionSpec{},
			expectedErr:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			distSpec, err := distributor.Distribute(
				ctx.WithBlockTime(ctx.BlockTime().Add(tc.timeOffset)),
				&gauge,
				tc.filterStakes,
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
