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
	"github.com/duality-labs/duality/x/dex/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestTickLiquidityQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := CreateNTickLiquidityLimitOrder(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetTickLiquidityRequest
		response *types.QueryGetTickLiquidityResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetTickLiquidityRequest{
				PairId:         msgs[0].PairId().Stringify(),
				TokenIn:        msgs[0].TokenIn(),
				TickIndex:      msgs[0].TickIndex(),
				LiquidityType:  msgs[0].LiquidityType(),
				LiquidityIndex: msgs[0].LiquidityIndex(),
			},
			response: &types.QueryGetTickLiquidityResponse{TickLiquidity: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetTickLiquidityRequest{
				PairId:         msgs[1].PairId().Stringify(),
				TokenIn:        msgs[1].TokenIn(),
				TickIndex:      msgs[1].TickIndex(),
				LiquidityType:  msgs[1].LiquidityType(),
				LiquidityIndex: msgs[1].LiquidityIndex(),
			},
			response: &types.QueryGetTickLiquidityResponse{TickLiquidity: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetTickLiquidityRequest{
				PairId:         "TokenZ<>TokenQ",
				TokenIn:        strconv.Itoa(100000),
				TickIndex:      100000,
				LiquidityType:  strconv.Itoa(100000),
				LiquidityIndex: 100000,
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.TickLiquidity(wctx, tc.request)
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

func TestTickLiquidityQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := CreateNTickLiquidityLimitOrder(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllTickLiquidityRequest {
		return &types.QueryAllTickLiquidityRequest{
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
			resp, err := keeper.TickLiquidityAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.TickLiquidity), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.TickLiquidity),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.TickLiquidityAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.TickLiquidity), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.TickLiquidity),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.TickLiquidityAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.TickLiquidity),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.TickLiquidityAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
