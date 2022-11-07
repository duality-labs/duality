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

func (k Keeper) LimitOrderPoolUserShareObjectAll(c context.Context, req *types.QueryAllLimitOrderPoolUserShareObjectRequest) (*types.QueryAllLimitOrderPoolUserShareObjectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var limitOrderPoolUserShareObjects []types.LimitOrderPoolUserShareObject
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	limitOrderPoolUserShareObjectStore := prefix.NewStore(store, types.KeyPrefix(types.LimitOrderPoolUserShareObjectKeyPrefix))

	pageRes, err := query.Paginate(limitOrderPoolUserShareObjectStore, req.Pagination, func(key []byte, value []byte) error {
		var limitOrderPoolUserShareObject types.LimitOrderPoolUserShareObject
		if err := k.cdc.Unmarshal(value, &limitOrderPoolUserShareObject); err != nil {
			return err
		}

		limitOrderPoolUserShareObjects = append(limitOrderPoolUserShareObjects, limitOrderPoolUserShareObject)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLimitOrderPoolUserShareObjectResponse{LimitOrderPoolUserShareObject: limitOrderPoolUserShareObjects, Pagination: pageRes}, nil
}

func (k Keeper) LimitOrderPoolUserShareObject(c context.Context, req *types.QueryGetLimitOrderPoolUserShareObjectRequest) (*types.QueryGetLimitOrderPoolUserShareObjectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetLimitOrderPoolUserShareObject(
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

	return &types.QueryGetLimitOrderPoolUserShareObjectResponse{LimitOrderPoolUserShareObject: val}, nil
}
