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

func (k Keeper) VirtualPriceQueueAll(c context.Context, req *types.QueryAllVirtualPriceQueueRequest) (*types.QueryAllVirtualPriceQueueResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var virtualPriceQueues []types.VirtualPriceQueue
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	virtualPriceQueueStore := prefix.NewStore(store, types.KeyPrefix(types.VirtualPriceQueueKeyPrefix))

	pageRes, err := query.Paginate(virtualPriceQueueStore, req.Pagination, func(key []byte, value []byte) error {
		var virtualPriceQueue types.VirtualPriceQueue
		if err := k.cdc.Unmarshal(value, &virtualPriceQueue); err != nil {
			return err
		}

		virtualPriceQueues = append(virtualPriceQueues, virtualPriceQueue)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVirtualPriceQueueResponse{VirtualPriceQueue: virtualPriceQueues, Pagination: pageRes}, nil
}

func (k Keeper) VirtualPriceQueue(c context.Context, req *types.QueryGetVirtualPriceQueueRequest) (*types.QueryGetVirtualPriceQueueResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetVirtualPriceQueue(
		ctx,
		req.VPrice,
		req.Direction,
		req.OrderType,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetVirtualPriceQueueResponse{VirtualPriceQueue: val}, nil
}
