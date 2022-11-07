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

func (k Keeper) PairObjectAll(c context.Context, req *types.QueryAllPairObjectRequest) (*types.QueryAllPairObjectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var pairObjects []types.PairObject
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	pairObjectStore := prefix.NewStore(store, types.KeyPrefix(types.PairObjectKeyPrefix))

	pageRes, err := query.Paginate(pairObjectStore, req.Pagination, func(key []byte, value []byte) error {
		var pairObject types.PairObject
		if err := k.cdc.Unmarshal(value, &pairObject); err != nil {
			return err
		}

		pairObjects = append(pairObjects, pairObject)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPairObjectResponse{PairObject: pairObjects, Pagination: pageRes}, nil
}

func (k Keeper) PairObject(c context.Context, req *types.QueryGetPairObjectRequest) (*types.QueryGetPairObjectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetPairObject(
		ctx,
		req.PairId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetPairObjectResponse{PairObject: val}, nil
}
