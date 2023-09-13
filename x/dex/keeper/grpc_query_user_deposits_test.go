package keeper_test

import (
	"testing"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	dualityapp "github.com/duality-labs/duality/app"
	keepertest "github.com/duality-labs/duality/x/dex/keeper/internal/testutils"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/duality-labs/duality/x/dex/utils"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func poolTokenToDepositRecord(coin sdk.Coin) *types.DepositRecord {
	depositDenom, err := types.NewDepositDenomFromString(coin.Denom)
	if err != nil {
		panic("failed to parse deposit denom")
	}

	return &types.DepositRecord{
		PairID:          depositDenom.PairID,
		SharesOwned:     coin.Amount,
		CenterTickIndex: depositDenom.Tick,
		LowerTickIndex:  depositDenom.Tick - utils.MustSafeUint64(depositDenom.Fee),
		UpperTickIndex:  depositDenom.Tick + utils.MustSafeUint64(depositDenom.Fee),
		Fee:             depositDenom.Fee,
	}
}

func TestUserDepositsAllQueryPaginated(t *testing.T) {
	app := dualityapp.Setup(t, false)
	keeper := app.DexKeeper
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	wctx := sdk.WrapSDKContext(ctx)
	addr := sdk.AccAddress([]byte("test_addr"))

	depositDenoms := sdk.Coins{
		sdk.NewInt64Coin(types.NewDepositDenom(defaultPairID, 2, 1).String(), 10),
		sdk.NewInt64Coin(types.NewDepositDenom(defaultPairID, 3, 1).String(), 10),
		sdk.NewInt64Coin(types.NewDepositDenom(defaultPairID, 4, 1).String(), 10),
		sdk.NewInt64Coin(types.NewDepositDenom(defaultPairID, 5, 1).String(), 10),
		sdk.NewInt64Coin(types.NewDepositDenom(defaultPairID, 6, 1).String(), 10),
	}
	var msgs []*types.DepositRecord

	for _, d := range depositDenoms {
		msgs = append(msgs, poolTokenToDepositRecord(d))
	}
	randomCoins := sdk.Coins{sdk.NewInt64Coin("TokenA", 10), sdk.NewInt64Coin("TokenZ", 10)}
	allCoins := randomCoins.Add(depositDenoms...)
	keepertest.FundAccount(app.BankKeeper, ctx, addr, allCoins)
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
