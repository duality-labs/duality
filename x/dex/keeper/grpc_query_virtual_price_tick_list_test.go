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

func TestVirtualPriceTickListQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNVirtualPriceTickList(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetVirtualPriceTickListRequest
		response *types.QueryGetVirtualPriceTickListResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetVirtualPriceTickListRequest{
				VPrice:    msgs[0].VPrice,
				Direction: msgs[0].Direction,
				OrderType: msgs[0].OrderType,
			},
			response: &types.QueryGetVirtualPriceTickListResponse{VirtualPriceTickList: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetVirtualPriceTickListRequest{
				VPrice:    msgs[1].VPrice,
				Direction: msgs[1].Direction,
				OrderType: msgs[1].OrderType,
			},
			response: &types.QueryGetVirtualPriceTickListResponse{VirtualPriceTickList: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetVirtualPriceTickListRequest{
				VPrice:    strconv.Itoa(100000),
				Direction: strconv.Itoa(100000),
				OrderType: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.VirtualPriceTickList(wctx, tc.request)
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

func TestVirtualPriceTickListQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNVirtualPriceTickList(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllVirtualPriceTickListRequest {
		return &types.QueryAllVirtualPriceTickListRequest{
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
			resp, err := keeper.VirtualPriceTickListAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.VirtualPriceTickList), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.VirtualPriceTickList),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.VirtualPriceTickListAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.VirtualPriceTickList), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.VirtualPriceTickList),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.VirtualPriceTickListAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.VirtualPriceTickList),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.VirtualPriceTickListAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
