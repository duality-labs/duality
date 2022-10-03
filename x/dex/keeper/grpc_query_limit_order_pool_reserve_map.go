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

func (k Keeper) LimitOrderPoolReserveMapAll(c context.Context, req *types.QueryAllLimitOrderPoolReserveMapRequest) (*types.QueryAllLimitOrderPoolReserveMapResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var limitOrderPoolReserveMaps []types.LimitOrderPoolReserveMap
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	limitOrderPoolReserveMapStore := prefix.NewStore(store, types.KeyPrefix(types.LimitOrderPoolReserveMapKeyPrefix))

	pageRes, err := query.Paginate(limitOrderPoolReserveMapStore, req.Pagination, func(key []byte, value []byte) error {
		var limitOrderPoolReserveMap types.LimitOrderPoolReserveMap
		if err := k.cdc.Unmarshal(value, &limitOrderPoolReserveMap); err != nil {
			return err
		}

		limitOrderPoolReserveMaps = append(limitOrderPoolReserveMaps, limitOrderPoolReserveMap)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLimitOrderPoolReserveMapResponse{LimitOrderPoolReserveMap: limitOrderPoolReserveMaps, Pagination: pageRes}, nil
}

func (k Keeper) LimitOrderPoolReserveMap(c context.Context, req *types.QueryGetLimitOrderPoolReserveMapRequest) (*types.QueryGetLimitOrderPoolReserveMapResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetLimitOrderPoolReserveMap(
		ctx,
		req.Count,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetLimitOrderPoolReserveMapResponse{LimitOrderPoolReserveMap: val}, nil
}
