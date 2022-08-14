package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) VirtualPriceTickQueueAll(c context.Context, req *types.QueryAllVirtualPriceTickQueueRequest) (*types.QueryAllVirtualPriceTickQueueResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var virtualPriceTickQueues []types.VirtualPriceTickQueue
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	virtualPriceTickQueueStore := prefix.NewStore(store, types.KeyPrefix(types.VirtualPriceTickQueueKey))

	pageRes, err := query.Paginate(virtualPriceTickQueueStore, req.Pagination, func(key []byte, value []byte) error {
		var virtualPriceTickQueue types.VirtualPriceTickQueue
		if err := k.cdc.Unmarshal(value, &virtualPriceTickQueue); err != nil {
			return err
		}

		virtualPriceTickQueues = append(virtualPriceTickQueues, virtualPriceTickQueue)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVirtualPriceTickQueueResponse{VirtualPriceTickQueue: virtualPriceTickQueues, Pagination: pageRes}, nil
}

func (k Keeper) VirtualPriceTickQueue(c context.Context, req *types.QueryGetVirtualPriceTickQueueRequest) (*types.QueryGetVirtualPriceTickQueueResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	virtualPriceTickQueue, found := k.GetVirtualPriceTickQueue(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetVirtualPriceTickQueueResponse{VirtualPriceTickQueue: virtualPriceTickQueue}, nil
}
