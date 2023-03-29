package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/duality-labs/duality/x/dex/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NOTE: For single queries of tick liquidity use explicty typed queries
// (ie. the k.LimitOrderTranche & k.PoolReserves)

func (k Keeper) TickLiquidityAll(c context.Context, req *types.QueryAllTickLiquidityRequest) (*types.QueryAllTickLiquidityResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var tickLiquiditys []types.TickLiquidity
	ctx := sdk.UnwrapSDKContext(c)

	pairId, err := StringToPairId(req.PairId)
	if err != nil {
		return nil, err
	}

	store := ctx.KVStore(k.storeKey)
	tickLiquidityStore := prefix.NewStore(store, types.TickLiquidityPrefix(pairId, req.TokenIn))

	pageRes, err := query.Paginate(tickLiquidityStore, req.Pagination, func(key, value []byte) error {
		var tickLiquidity types.TickLiquidity
		if err := k.cdc.Unmarshal(value, &tickLiquidity); err != nil {
			return err
		}

		tickLiquiditys = append(tickLiquiditys, tickLiquidity)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTickLiquidityResponse{TickLiquidity: tickLiquiditys, Pagination: pageRes}, nil
}
