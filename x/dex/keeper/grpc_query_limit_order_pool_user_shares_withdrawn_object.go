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

func (k Keeper) LimitOrderPoolUserSharesWithdrawnObjectAll(c context.Context, req *types.QueryAllLimitOrderPoolUserSharesWithdrawnObjectRequest) (*types.QueryAllLimitOrderPoolUserSharesWithdrawnObjectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var limitOrderPoolUserSharesWithdrawnObjects []types.LimitOrderPoolUserSharesWithdrawnObject
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	limitOrderPoolUserSharesWithdrawnObjectStore := prefix.NewStore(store, types.KeyPrefix(types.LimitOrderPoolUserSharesWithdrawnObjectKeyPrefix))

	pageRes, err := query.Paginate(limitOrderPoolUserSharesWithdrawnObjectStore, req.Pagination, func(key []byte, value []byte) error {
		var limitOrderPoolUserSharesWithdrawnObject types.LimitOrderPoolUserSharesWithdrawnObject
		if err := k.cdc.Unmarshal(value, &limitOrderPoolUserSharesWithdrawnObject); err != nil {
			return err
		}

		limitOrderPoolUserSharesWithdrawnObjects = append(limitOrderPoolUserSharesWithdrawnObjects, limitOrderPoolUserSharesWithdrawnObject)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLimitOrderPoolUserSharesWithdrawnObjectResponse{LimitOrderPoolUserSharesWithdrawnObject: limitOrderPoolUserSharesWithdrawnObjects, Pagination: pageRes}, nil
}

func (k Keeper) LimitOrderPoolUserSharesWithdrawnObject(c context.Context, req *types.QueryGetLimitOrderPoolUserSharesWithdrawnObjectRequest) (*types.QueryGetLimitOrderPoolUserSharesWithdrawnObjectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetLimitOrderPoolUserSharesWithdrawnObject(
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

	return &types.QueryGetLimitOrderPoolUserSharesWithdrawnObjectResponse{LimitOrderPoolUserSharesWithdrawnObject: val}, nil
}
