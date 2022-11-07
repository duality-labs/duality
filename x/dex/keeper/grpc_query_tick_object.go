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

func (k Keeper) TickObjectAll(c context.Context, req *types.QueryAllTickObjectRequest) (*types.QueryAllTickObjectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var tickObjects []types.TickObject
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	tickObjectStore := prefix.NewStore(store, types.KeyPrefix(types.BaseTickObjectKeyPrefix))

	pageRes, err := query.Paginate(tickObjectStore, req.Pagination, func(key []byte, value []byte) error {
		var tickObject types.TickObject
		if err := k.cdc.Unmarshal(value, &tickObject); err != nil {
			return err
		}

		tickObjects = append(tickObjects, tickObject)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTickObjectResponse{TickObject: tickObjects, Pagination: pageRes}, nil
}

func (k Keeper) TickObject(c context.Context, req *types.QueryGetTickObjectRequest) (*types.QueryGetTickObjectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetTickObject(
		ctx,
		req.PairId,
		req.TickIndex,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTickObjectResponse{TickObject: val}, nil
}
