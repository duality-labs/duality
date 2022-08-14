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

func (k Keeper) VirtualPriceTickListAll(c context.Context, req *types.QueryAllVirtualPriceTickListRequest) (*types.QueryAllVirtualPriceTickListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var virtualPriceTickLists []types.VirtualPriceTickList
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	virtualPriceTickListStore := prefix.NewStore(store, types.KeyPrefix(types.VirtualPriceTickListKeyPrefix))

	pageRes, err := query.Paginate(virtualPriceTickListStore, req.Pagination, func(key []byte, value []byte) error {
		var virtualPriceTickList types.VirtualPriceTickList
		if err := k.cdc.Unmarshal(value, &virtualPriceTickList); err != nil {
			return err
		}

		virtualPriceTickLists = append(virtualPriceTickLists, virtualPriceTickList)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVirtualPriceTickListResponse{VirtualPriceTickList: virtualPriceTickLists, Pagination: pageRes}, nil
}

func (k Keeper) VirtualPriceTickList(c context.Context, req *types.QueryGetVirtualPriceTickListRequest) (*types.QueryGetVirtualPriceTickListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetVirtualPriceTickList(
		ctx,
		req.VPrice,
		req.Direction,
		req.OrderType,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetVirtualPriceTickListResponse{VirtualPriceTickList: val}, nil
}
