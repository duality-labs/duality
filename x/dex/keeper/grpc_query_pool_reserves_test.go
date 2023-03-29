package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/duality-labs/duality/testutil/keeper"
	"github.com/duality-labs/duality/testutil/nullify"
	"github.com/duality-labs/duality/x/dex/types"
)

func TestPoolReservesQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNPoolReserves(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetPoolReservesRequest
		response *types.QueryGetPoolReservesResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetPoolReservesRequest{
				PairId:    "TokenA<>TokenB",
				TickIndex: msgs[0].TickIndex,
				TokenIn:   "TokenA",
				Fee:       msgs[0].Fee,
			},
			response: &types.QueryGetPoolReservesResponse{PoolReserves: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetPoolReservesRequest{
				PairId:    "TokenA<>TokenB",
				TickIndex: msgs[1].TickIndex,
				TokenIn:   "TokenA",
				Fee:       msgs[1].Fee,
			},
			response: &types.QueryGetPoolReservesResponse{PoolReserves: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetPoolReservesRequest{
				PairId:    "TokenA<>TokenB",
				TickIndex: 0,
				TokenIn:   "TokenA",
				Fee:       100000,
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.PoolReserves(wctx, tc.request)
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

func TestPoolReservesQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNPoolReserves(keeper, ctx, 2)
	// Add more data to make sure only LO tranches are returned
	createNLimitOrderTranches(keeper, ctx, 2)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllPoolReservesRequest {
		return &types.QueryAllPoolReservesRequest{
			PairId:  "TokenA<>TokenB",
			TokenIn: "TokenA",
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
			resp, err := keeper.PoolReservesAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.PoolReserves), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.PoolReserves),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.PoolReservesAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.PoolReserves), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.PoolReserves),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.PoolReservesAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.PoolReserves),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.PoolReservesAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
