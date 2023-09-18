package keeper_test

import (
	"testing"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	dualityapp "github.com/duality-labs/duality/app"
	keepertest "github.com/duality-labs/duality/x/dex/keeper/internal/testutils"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func simulateDeposit(ctx sdk.Context, app *dualityapp.App, addr sdk.AccAddress, deposit *types.DepositRecord) {
	// NOTE: For simplicyt sake, we are not actually doing a deposit, we are just initializing
	// the pool and adding the poolDenom to the users account
	pool, err := app.DexKeeper.InitPool(ctx, deposit.PairID, deposit.CenterTickIndex, deposit.Fee)
	if err != nil {
		panic("Cannot init pool")
	}
	coin := sdk.NewCoin(pool.GetPoolDenom(), deposit.SharesOwned)
	keepertest.FundAccount(app.BankKeeper, ctx, addr, sdk.Coins{coin})
}

func TestUserDepositsAllQueryPaginated(t *testing.T) {
	app := dualityapp.Setup(t, false)
	keeper := app.DexKeeper
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	wctx := sdk.WrapSDKContext(ctx)
	addr := sdk.AccAddress([]byte("test_addr"))
	msgs := []*types.DepositRecord{
		{
			PairID:          defaultPairID,
			SharesOwned:     sdk.NewInt(10),
			CenterTickIndex: 2,
			LowerTickIndex:  1,
			UpperTickIndex:  3,
			Fee:             1,
		},
		{
			PairID:          defaultPairID,
			SharesOwned:     sdk.NewInt(10),
			CenterTickIndex: 3,
			LowerTickIndex:  2,
			UpperTickIndex:  4,
			Fee:             1,
		},
		{
			PairID:          defaultPairID,
			SharesOwned:     sdk.NewInt(10),
			CenterTickIndex: 4,
			LowerTickIndex:  3,
			UpperTickIndex:  5,
			Fee:             1,
		},
		{
			PairID:          defaultPairID,
			SharesOwned:     sdk.NewInt(10),
			CenterTickIndex: 5,
			LowerTickIndex:  4,
			UpperTickIndex:  6,
			Fee:             1,
		},
		{
			PairID:          defaultPairID,
			SharesOwned:     sdk.NewInt(10),
			CenterTickIndex: 6,
			LowerTickIndex:  5,
			UpperTickIndex:  7,
			Fee:             1,
		},
	}

	for _, d := range msgs {
		simulateDeposit(ctx, app, addr, d)
	}

	// Add random noise to make sure only pool denoms are picked up
	randomCoins := sdk.Coins{sdk.NewInt64Coin("TokenA", 10), sdk.NewInt64Coin("TokenZ", 10)}
	keepertest.FundAccount(app.BankKeeper, ctx, addr, randomCoins)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllUserDepositsRequest {
		return &types.QueryAllUserDepositsRequest{
			Address: addr.String(),
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.UserDepositsAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Deposits), step)
			require.Subset(t,
				msgs,
				resp.Deposits,
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		var allRecords []*types.DepositRecord
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.UserDepositsAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Deposits), step)
			require.Subset(t,
				msgs,
				resp.Deposits,
			)

			allRecords = append(allRecords, resp.Deposits...)
			next = resp.Pagination.NextKey
		}
		require.ElementsMatch(t,
			msgs,
			allRecords,
		)
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.UserDepositsAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			msgs,
			resp.Deposits,
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.UserDepositsAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
