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

func (k Keeper) LimitOrderPoolTotalSharesObjectAll(c context.Context, req *types.QueryAllLimitOrderPoolTotalSharesObjectRequest) (*types.QueryAllLimitOrderPoolTotalSharesObjectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var limitOrderPoolTotalSharesObjects []types.LimitOrderPoolTotalSharesObject
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	limitOrderPoolTotalSharesObjectStore := prefix.NewStore(store, types.KeyPrefix(types.LimitOrderPoolTotalSharesObjectKeyPrefix))

	pageRes, err := query.Paginate(limitOrderPoolTotalSharesObjectStore, req.Pagination, func(key []byte, value []byte) error {
		var limitOrderPoolTotalSharesObject types.LimitOrderPoolTotalSharesObject
		if err := k.cdc.Unmarshal(value, &limitOrderPoolTotalSharesObject); err != nil {
			return err
		}

		limitOrderPoolTotalSharesObjects = append(limitOrderPoolTotalSharesObjects, limitOrderPoolTotalSharesObject)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLimitOrderPoolTotalSharesObjectResponse{LimitOrderPoolTotalSharesObject: limitOrderPoolTotalSharesObjects, Pagination: pageRes}, nil
}

func (k Keeper) LimitOrderPoolTotalSharesObject(c context.Context, req *types.QueryGetLimitOrderPoolTotalSharesObjectRequest) (*types.QueryGetLimitOrderPoolTotalSharesObjectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetLimitOrderPoolTotalSharesObject(
		ctx,
		req.PairId,
		req.TickIndex,
		req.Token,
		req.Count,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetLimitOrderPoolTotalSharesObjectResponse{LimitOrderPoolTotalSharesObject: val}, nil
}
