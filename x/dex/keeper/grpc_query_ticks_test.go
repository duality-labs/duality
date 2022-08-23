package keeper_test

//NOTE: Not sure but believe as Tick is always a subtype of Pair that is should not have its own grpc calls (this file exists as I had to define the mapping (the subtype) as its own mapping during scaffolding)
// I think this should be delted and we should remove grpc-query for Ticks as well

// import (
// 	"strconv"
// 	"testing"

// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	"github.com/cosmos/cosmos-sdk/types/query"
// 	"github.com/stretchr/testify/require"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"

// 	keepertest "github.com/NicholasDotSol/duality/testutil/keeper"
// 	"github.com/NicholasDotSol/duality/testutil/nullify"
// 	"github.com/NicholasDotSol/duality/x/dex/types"
// )

// // Prevent strconv unused error
// var _ = strconv.IntSize

// func TestTicksQuerySingle(t *testing.T) {
// 	keeper, ctx := keepertest.DexKeeper(t)
// 	wctx := sdk.WrapSDKContext(ctx)
// 	msgs := createNTicks(keeper, ctx, 2)
// 	for _, tc := range []struct {
// 		desc     string
// 		request  *types.QueryGetTicksRequest
// 		response *types.QueryGetTicksResponse
// 		err      error
// 	}{
// 		{
// 			desc: "First",
// 			request: &types.QueryGetTicksRequest{
// 				Token0:    "TokenB",
// 				Token1:    "TokenA",
// 				Price:     msgs[0].Price,
// 				Fee:       msgs[0].Fee,
// 				OrderType: msgs[0].OrderType,
// 			},
// 			response: &types.QueryGetTicksResponse{Ticks: msgs[0]},
// 		},
// 		{
// 			desc: "Second",
// 			request: &types.QueryGetTicksRequest{
// 				Token0:    "TokenB",
// 				Token1:    "TokenA",
// 				Price:     msgs[1].Price,
// 				Fee:       msgs[1].Fee,
// 				OrderType: msgs[1].OrderType,
// 			},
// 			response: &types.QueryGetTicksResponse{Ticks: msgs[1]},
// 		},
// 		{
// 			desc: "KeyNotFound",
// 			request: &types.QueryGetTicksRequest{
// 				Token0:    "TokenB",
// 				Token1:    "TokenA",
// 				Price:     strconv.Itoa(100000),
// 				Fee:       strconv.Itoa(100000),
// 				OrderType: strconv.Itoa(100000),
// 			},
// 			err: status.Error(codes.NotFound, "not found"),
// 		},
// 		{
// 			desc: "InvalidRequest",
// 			err:  status.Error(codes.InvalidArgument, "invalid request"),
// 		},
// 	} {
// 		t.Run(tc.desc, func(t *testing.T) {
// 			response, err := keeper.Ticks(wctx, tc.request)
// 			if tc.err != nil {
// 				require.ErrorIs(t, err, tc.err)
// 			} else {
// 				require.NoError(t, err)
// 				require.Equal(t,
// 					nullify.Fill(tc.response),
// 					nullify.Fill(response),
// 				)
// 			}
// 		})
// 	}
// }

// func TestTicksQueryPaginated(t *testing.T) {
// 	keeper, ctx := keepertest.DexKeeper(t)
// 	wctx := sdk.WrapSDKContext(ctx)
// 	msgs := createNTicks(keeper, ctx, 5)

// 	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllTicksRequest {
// 		return &types.QueryAllTicksRequest{
// 			Pagination: &query.PageRequest{
// 				Key:        next,
// 				Offset:     offset,
// 				Limit:      limit,
// 				CountTotal: total,
// 			},
// 		}
// 	}
// 	t.Run("ByOffset", func(t *testing.T) {
// 		step := 2
// 		for i := 0; i < len(msgs); i += step {
// 			resp, err := keeper.TicksAll(wctx, request(nil, uint64(i), uint64(step), false))
// 			require.NoError(t, err)
// 			require.LessOrEqual(t, len(resp.Ticks), step)
// 			require.Subset(t,
// 				nullify.Fill(msgs),
// 				nullify.Fill(resp.Ticks),
// 			)
// 		}
// 	})
// 	t.Run("ByKey", func(t *testing.T) {
// 		step := 2
// 		var next []byte
// 		for i := 0; i < len(msgs); i += step {
// 			resp, err := keeper.TicksAll(wctx, request(next, 0, uint64(step), false))
// 			require.NoError(t, err)
// 			require.LessOrEqual(t, len(resp.Ticks), step)
// 			require.Subset(t,
// 				nullify.Fill(msgs),
// 				nullify.Fill(resp.Ticks),
// 			)
// 			next = resp.Pagination.NextKey
// 		}
// 	})
// 	t.Run("Total", func(t *testing.T) {
// 		resp, err := keeper.TicksAll(wctx, request(nil, 0, 0, true))
// 		require.NoError(t, err)
// 		require.Equal(t, len(msgs), int(resp.Pagination.Total))
// 		require.ElementsMatch(t,
// 			nullify.Fill(msgs),
// 			nullify.Fill(resp.Ticks),
// 		)
// 	})
// 	t.Run("InvalidRequest", func(t *testing.T) {
// 		_, err := keeper.TicksAll(wctx, nil)
// 		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
// 	})
// }
