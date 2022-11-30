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

func TestTickQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNTick(keeper, ctx, "TokenA<>TokenB", 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetTickRequest
		response *types.QueryGetTickResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetTickRequest{
				TickIndex: msgs[0].TickIndex,
				PairId:    "TokenA<>TokenB",
			},
			response: &types.QueryGetTickResponse{Tick: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetTickRequest{
				TickIndex: msgs[1].TickIndex,
				PairId:    "TokenA<>TokenB",
			},
			response: &types.QueryGetTickResponse{Tick: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetTickRequest{
				TickIndex: 100000,
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Tick(wctx, tc.request)
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

func TestTickQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNTick(keeper, ctx, "TokenB/TokenA", 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllTickRequest {
		return &types.QueryAllTickRequest{
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
			resp, err := keeper.TickAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Tick), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Tick),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.TickAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Tick), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Tick),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.TickAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Tick),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.TickAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
