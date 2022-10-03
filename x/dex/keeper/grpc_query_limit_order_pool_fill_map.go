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

func (k Keeper) LimitOrderPoolFillMapAll(c context.Context, req *types.QueryAllLimitOrderPoolFillMapRequest) (*types.QueryAllLimitOrderPoolFillMapResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var limitOrderPoolFillMaps []types.LimitOrderPoolFillMap
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	limitOrderPoolFillMapStore := prefix.NewStore(store, types.KeyPrefix(types.LimitOrderPoolFillMapKeyPrefix))

	pageRes, err := query.Paginate(limitOrderPoolFillMapStore, req.Pagination, func(key []byte, value []byte) error {
		var limitOrderPoolFillMap types.LimitOrderPoolFillMap
		if err := k.cdc.Unmarshal(value, &limitOrderPoolFillMap); err != nil {
			return err
		}

		limitOrderPoolFillMaps = append(limitOrderPoolFillMaps, limitOrderPoolFillMap)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLimitOrderPoolFillMapResponse{LimitOrderPoolFillMap: limitOrderPoolFillMaps, Pagination: pageRes}, nil
}

func (k Keeper) LimitOrderPoolFillMap(c context.Context, req *types.QueryGetLimitOrderPoolFillMapRequest) (*types.QueryGetLimitOrderPoolFillMapResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetLimitOrderPoolFillMap(
		ctx,
		req.Count,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetLimitOrderPoolFillMapResponse{LimitOrderPoolFillMap: val}, nil
}
