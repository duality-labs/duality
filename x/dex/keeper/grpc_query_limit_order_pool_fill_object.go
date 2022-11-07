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

func (k Keeper) LimitOrderPoolFillObjectAll(c context.Context, req *types.QueryAllLimitOrderPoolFillObjectRequest) (*types.QueryAllLimitOrderPoolFillObjectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var limitOrderPoolFillObjects []types.LimitOrderPoolFillObject
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	limitOrderPoolFillObjectStore := prefix.NewStore(store, types.KeyPrefix(types.LimitOrderPoolFillObjectKeyPrefix))

	pageRes, err := query.Paginate(limitOrderPoolFillObjectStore, req.Pagination, func(key []byte, value []byte) error {
		var limitOrderPoolFillObject types.LimitOrderPoolFillObject
		if err := k.cdc.Unmarshal(value, &limitOrderPoolFillObject); err != nil {
			return err
		}

		limitOrderPoolFillObjects = append(limitOrderPoolFillObjects, limitOrderPoolFillObject)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLimitOrderPoolFillObjectResponse{LimitOrderPoolFillObject: limitOrderPoolFillObjects, Pagination: pageRes}, nil
}

func (k Keeper) LimitOrderPoolFillObject(c context.Context, req *types.QueryGetLimitOrderPoolFillObjectRequest) (*types.QueryGetLimitOrderPoolFillObjectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetLimitOrderPoolFillObject(
		ctx,
		req.PairId,
		req.TickIndex,
		req.Token,
		req.Count,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetLimitOrderPoolFillObjectResponse{LimitOrderPoolFillObject: val}, nil
}
