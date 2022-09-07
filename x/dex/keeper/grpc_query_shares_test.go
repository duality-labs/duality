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

func TestSharesQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNShares(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetSharesRequest
		response *types.QueryGetSharesResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetSharesRequest{
				Address:    msgs[0].Address,
				PairId:     msgs[0].PairId,
				PriceIndex: msgs[0].PriceIndex,
				Fee:        msgs[0].Fee,
			},
			response: &types.QueryGetSharesResponse{Shares: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetSharesRequest{
				Address:    msgs[1].Address,
				PairId:     msgs[1].PairId,
				PriceIndex: msgs[1].PriceIndex,
				Fee:        msgs[1].Fee,
			},
			response: &types.QueryGetSharesResponse{Shares: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetSharesRequest{
				Address:    strconv.Itoa(100000),
				PairId:     strconv.Itoa(100000),
				PriceIndex: strconv.Itoa(100000),
				Fee:        strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Shares(wctx, tc.request)
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

func TestSharesQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNShares(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllSharesRequest {
		return &types.QueryAllSharesRequest{
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
			resp, err := keeper.SharesAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Shares), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Shares),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.SharesAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Shares), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Shares),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.SharesAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Shares),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.SharesAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
