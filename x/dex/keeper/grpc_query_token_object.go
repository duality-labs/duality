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

func (k Keeper) TokenObjectAll(c context.Context, req *types.QueryAllTokenObjectRequest) (*types.QueryAllTokenObjectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var tokenObjects []types.TokenObject
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	tokenObjectStore := prefix.NewStore(store, types.KeyPrefix(types.TokenObjectKeyPrefix))

	pageRes, err := query.Paginate(tokenObjectStore, req.Pagination, func(key []byte, value []byte) error {
		var tokenObject types.TokenObject
		if err := k.cdc.Unmarshal(value, &tokenObject); err != nil {
			return err
		}

		tokenObjects = append(tokenObjects, tokenObject)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTokenObjectResponse{TokenObject: tokenObjects, Pagination: pageRes}, nil
}

func (k Keeper) TokenObject(c context.Context, req *types.QueryGetTokenObjectRequest) (*types.QueryGetTokenObjectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetTokenObject(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTokenObjectResponse{TokenObject: val}, nil
}
