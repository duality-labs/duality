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

func (k Keeper) TokenMapAll(c context.Context, req *types.QueryAllTokenMapRequest) (*types.QueryAllTokenMapResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var tokenMaps []types.TokenMap
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	tokenMapStore := prefix.NewStore(store, types.KeyPrefix(types.TokenMapKeyPrefix))

	pageRes, err := query.Paginate(tokenMapStore, req.Pagination, func(key []byte, value []byte) error {
		var tokenMap types.TokenMap
		if err := k.cdc.Unmarshal(value, &tokenMap); err != nil {
			return err
		}

		tokenMaps = append(tokenMaps, tokenMap)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTokenMapResponse{TokenMap: tokenMaps, Pagination: pageRes}, nil
}

func (k Keeper) TokenMap(c context.Context, req *types.QueryGetTokenMapRequest) (*types.QueryGetTokenMapResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetTokenMap(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTokenMapResponse{TokenMap: val}, nil
}
