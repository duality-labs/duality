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

func (k Keeper) TicksAll(c context.Context, req *types.QueryAllTicksRequest) (*types.QueryAllTicksResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var tickss []types.Ticks
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	ticksStore := prefix.NewStore(store, types.KeyPrefix(types.TicksKeyPrefix))

	pageRes, err := query.Paginate(ticksStore, req.Pagination, func(key []byte, value []byte) error {
		var ticks types.Ticks
		if err := k.cdc.Unmarshal(value, &ticks); err != nil {
			return err
		}

		tickss = append(tickss, ticks)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTicksResponse{Ticks: tickss, Pagination: pageRes}, nil
}

func (k Keeper) Ticks(c context.Context, req *types.QueryGetTicksRequest) (*types.QueryGetTicksResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetTicks(
		ctx,
		req.Token0,
		req.Token1,
		req.Price,
		req.Fee,
		req.Direction,
		req.OrderType,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTicksResponse{Ticks: val}, nil
}
