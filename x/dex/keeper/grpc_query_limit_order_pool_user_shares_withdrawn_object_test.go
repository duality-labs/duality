package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/NicholasDotSol/duality/testutil/keeper"
	"github.com/NicholasDotSol/duality/testutil/nullify"
	"github.com/NicholasDotSol/duality/x/dex/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestLimitOrderPoolUserSharesWithdrawnObjectQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNLimitOrderPoolUserSharesWithdrawnObject(keeper, ctx, "TokenA/TokenB", 0, "TokenA", 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetLimitOrderPoolUserSharesWithdrawnObjectRequest
		response *types.QueryGetLimitOrderPoolUserSharesWithdrawnObjectResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetLimitOrderPoolUserSharesWithdrawnObjectRequest{
				PairId:    "TokenA/TokenB",
				TickIndex: 0,
				Token:     "TokenA",
				Count:     msgs[0].Count,
				Address:   msgs[0].Address,
			},
			response: &types.QueryGetLimitOrderPoolUserSharesWithdrawnObjectResponse{LimitOrderPoolUserSharesWithdrawnObject: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetLimitOrderPoolUserSharesWithdrawnObjectRequest{
				PairId:    "TokenA/TokenB",
				TickIndex: 0,
				Token:     "TokenA",
				Count:     msgs[1].Count,
				Address:   msgs[1].Address,
			},
			response: &types.QueryGetLimitOrderPoolUserSharesWithdrawnObjectResponse{LimitOrderPoolUserSharesWithdrawnObject: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetLimitOrderPoolUserSharesWithdrawnObjectRequest{
				PairId:    "TokenA/TokenB",
				TickIndex: 0,
				Token:     "TokenA",
				Count:     100000,
				Address:   strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.LimitOrderPoolUserSharesWithdrawnObject(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestLimitOrderPoolUserSharesWithdrawnObjectQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNLimitOrderPoolUserSharesWithdrawnObject(keeper, ctx, "TokenA/TokenB", 0, "TokenA", 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllLimitOrderPoolUserSharesWithdrawnObjectRequest {
		return &types.QueryAllLimitOrderPoolUserSharesWithdrawnObjectRequest{
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
			resp, err := keeper.LimitOrderPoolUserSharesWithdrawnObjectAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.LimitOrderPoolUserSharesWithdrawnObject), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.LimitOrderPoolUserSharesWithdrawnObject),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.LimitOrderPoolUserSharesWithdrawnObjectAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.LimitOrderPoolUserSharesWithdrawnObject), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.LimitOrderPoolUserSharesWithdrawnObject),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.LimitOrderPoolUserSharesWithdrawnObjectAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.LimitOrderPoolUserSharesWithdrawnObject),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.LimitOrderPoolUserSharesWithdrawnObjectAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
