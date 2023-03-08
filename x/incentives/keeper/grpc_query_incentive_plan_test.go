package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/duality-labs/duality/testutil/keeper"
	"github.com/duality-labs/duality/testutil/nullify"
	"github.com/duality-labs/duality/x/incentives/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestIncentivePlanQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.IncentivesKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNIncentivePlan(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetIncentivePlanRequest
		response *types.QueryGetIncentivePlanResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetIncentivePlanRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetIncentivePlanResponse{IncentivePlan: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetIncentivePlanRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetIncentivePlanResponse{IncentivePlan: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetIncentivePlanRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.IncentivePlan(wctx, tc.request)
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

func TestIncentivePlanQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.IncentivesKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNIncentivePlan(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllIncentivePlanRequest {
		return &types.QueryAllIncentivePlanRequest{
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
			resp, err := keeper.IncentivePlanAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.IncentivePlan), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.IncentivePlan),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.IncentivePlanAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.IncentivePlan), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.IncentivePlan),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.IncentivePlanAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.IncentivePlan),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.IncentivePlanAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
