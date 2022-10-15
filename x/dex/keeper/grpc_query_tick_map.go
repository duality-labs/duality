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

func (k Keeper) TickMapAll(c context.Context, req *types.QueryAllTickMapRequest) (*types.QueryAllTickMapResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var tickMaps []types.TickMap
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	tickMapStore := prefix.NewStore(store, types.KeyPrefix(types.BaseTickMapKeyPrefix))

	pageRes, err := query.Paginate(tickMapStore, req.Pagination, func(key []byte, value []byte) error {
		var tickMap types.TickMap
		if err := k.cdc.Unmarshal(value, &tickMap); err != nil {
			return err
		}

		tickMaps = append(tickMaps, tickMap)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTickMapResponse{TickMap: tickMaps, Pagination: pageRes}, nil
}

func (k Keeper) TickMap(c context.Context, req *types.QueryGetTickMapRequest) (*types.QueryGetTickMapResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetTickMap(
		ctx,
		req.PairId,
		req.TickIndex,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTickMapResponse{TickMap: val}, nil
}
