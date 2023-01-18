package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TickLiquidityAll(c context.Context, req *types.QueryAllTickLiquidityRequest) (*types.QueryAllTickLiquidityResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var tickLiquiditys []types.TickLiquidity
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	tickLiquidityStore := prefix.NewStore(store, types.KeyPrefix(types.TickLiquidityKeyPrefix))

	pageRes, err := query.Paginate(tickLiquidityStore, req.Pagination, func(key []byte, value []byte) error {
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

func (k Keeper) TickLiquidity(c context.Context, req *types.QueryGetTickLiquidityRequest) (*types.QueryGetTickLiquidityResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	pairId, err := StringToPairId(req.PairId)
	if err != nil {
		return nil, err
	}

	val, found := k.GetTickLiquidity(
		ctx,
		pairId,
		req.TokenIn,
		req.TickIndex,
		req.LiquidityType,
		req.LiquidityIndex,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTickLiquidityResponse{TickLiquidity: val}, nil
}
