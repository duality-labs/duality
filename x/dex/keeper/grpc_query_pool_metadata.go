package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/duality-labs/duality/x/dex/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PoolMetadataAll(goCtx context.Context, req *types.QueryAllPoolMetadataRequest) (*types.QueryAllPoolMetadataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var poolMetadatas []types.PoolMetadata
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	poolMetadataStore := prefix.NewStore(store, types.KeyPrefix(types.PoolMetadataKeyPrefix))

	pageRes, err := query.Paginate(poolMetadataStore, req.Pagination, func(key []byte, value []byte) error {
		var poolMetadata types.PoolMetadata
		if err := k.cdc.Unmarshal(value, &poolMetadata); err != nil {
			return err
		}

		poolMetadatas = append(poolMetadatas, poolMetadata)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPoolMetadataResponse{PoolMetadata: poolMetadatas, Pagination: pageRes}, nil
}

func (k Keeper) PoolMetadata(goCtx context.Context, req *types.QueryGetPoolMetadataRequest) (*types.QueryGetPoolMetadataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	poolMetadata, found := k.GetPoolMetadata(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetPoolMetadataResponse{PoolMetadata: poolMetadata}, nil
}
