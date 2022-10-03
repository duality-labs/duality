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

func (k Keeper) LimitOrderPoolTotalSharesMapAll(c context.Context, req *types.QueryAllLimitOrderPoolTotalSharesMapRequest) (*types.QueryAllLimitOrderPoolTotalSharesMapResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var limitOrderPoolTotalSharesMaps []types.LimitOrderPoolTotalSharesMap
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	limitOrderPoolTotalSharesMapStore := prefix.NewStore(store, types.KeyPrefix(types.LimitOrderPoolTotalSharesMapKeyPrefix))

	pageRes, err := query.Paginate(limitOrderPoolTotalSharesMapStore, req.Pagination, func(key []byte, value []byte) error {
		var limitOrderPoolTotalSharesMap types.LimitOrderPoolTotalSharesMap
		if err := k.cdc.Unmarshal(value, &limitOrderPoolTotalSharesMap); err != nil {
			return err
		}

		limitOrderPoolTotalSharesMaps = append(limitOrderPoolTotalSharesMaps, limitOrderPoolTotalSharesMap)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLimitOrderPoolTotalSharesMapResponse{LimitOrderPoolTotalSharesMap: limitOrderPoolTotalSharesMaps, Pagination: pageRes}, nil
}

func (k Keeper) LimitOrderPoolTotalSharesMap(c context.Context, req *types.QueryGetLimitOrderPoolTotalSharesMapRequest) (*types.QueryGetLimitOrderPoolTotalSharesMapResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetLimitOrderPoolTotalSharesMap(
		ctx,
		req.Count,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetLimitOrderPoolTotalSharesMapResponse{LimitOrderPoolTotalSharesMap: val}, nil
}
