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

func (k Keeper) LimitOrderPoolReserveObjectAll(c context.Context, req *types.QueryAllLimitOrderPoolReserveObjectRequest) (*types.QueryAllLimitOrderPoolReserveObjectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var limitOrderPoolReserveObjects []types.LimitOrderPoolReserveObject
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	limitOrderPoolReserveObjectStore := prefix.NewStore(store, types.KeyPrefix(types.LimitOrderPoolReserveObjectKeyPrefix))

	pageRes, err := query.Paginate(limitOrderPoolReserveObjectStore, req.Pagination, func(key []byte, value []byte) error {
		var limitOrderPoolReserveObject types.LimitOrderPoolReserveObject
		if err := k.cdc.Unmarshal(value, &limitOrderPoolReserveObject); err != nil {
			return err
		}

		limitOrderPoolReserveObjects = append(limitOrderPoolReserveObjects, limitOrderPoolReserveObject)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLimitOrderPoolReserveObjectResponse{LimitOrderPoolReserveObject: limitOrderPoolReserveObjects, Pagination: pageRes}, nil
}

func (k Keeper) LimitOrderPoolReserveObject(c context.Context, req *types.QueryGetLimitOrderPoolReserveObjectRequest) (*types.QueryGetLimitOrderPoolReserveObjectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetLimitOrderPoolReserveObject(
		ctx,
		req.PairId,
		req.TickIndex,
		req.Token,
		req.Count,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetLimitOrderPoolReserveObjectResponse{LimitOrderPoolReserveObject: val}, nil
}
