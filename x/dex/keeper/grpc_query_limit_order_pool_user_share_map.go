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

func (k Keeper) LimitOrderPoolUserShareMapAll(c context.Context, req *types.QueryAllLimitOrderPoolUserShareMapRequest) (*types.QueryAllLimitOrderPoolUserShareMapResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var limitOrderPoolUserShareMaps []types.LimitOrderPoolUserShareMap
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	limitOrderPoolUserShareMapStore := prefix.NewStore(store, types.KeyPrefix(types.LimitOrderPoolUserShareMapKeyPrefix))

	pageRes, err := query.Paginate(limitOrderPoolUserShareMapStore, req.Pagination, func(key []byte, value []byte) error {
		var limitOrderPoolUserShareMap types.LimitOrderPoolUserShareMap
		if err := k.cdc.Unmarshal(value, &limitOrderPoolUserShareMap); err != nil {
			return err
		}

		limitOrderPoolUserShareMaps = append(limitOrderPoolUserShareMaps, limitOrderPoolUserShareMap)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLimitOrderPoolUserShareMapResponse{LimitOrderPoolUserShareMap: limitOrderPoolUserShareMaps, Pagination: pageRes}, nil
}

func (k Keeper) LimitOrderPoolUserShareMap(c context.Context, req *types.QueryGetLimitOrderPoolUserShareMapRequest) (*types.QueryGetLimitOrderPoolUserShareMapResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetLimitOrderPoolUserShareMap(
		ctx,
		req.PairId,
		req.TickIndex,
		req.Token,
		req.Count,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetLimitOrderPoolUserShareMapResponse{LimitOrderPoolUserShareMap: val}, nil
}
