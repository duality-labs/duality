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

func (k Keeper) IndexQueueAll(c context.Context, req *types.QueryAllIndexQueueRequest) (*types.QueryAllIndexQueueResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var IndexQueues []types.IndexQueue
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	IndexQueueStore := prefix.NewStore(store, types.IndexQueuePrefix(req.Token0, req.Token1))

	pageRes, err := query.Paginate(IndexQueueStore, req.Pagination, func(key []byte, value []byte) error {
		var IndexQueue types.IndexQueue
		if err := k.cdc.Unmarshal(value, &IndexQueue); err != nil {
			return err
		}

		IndexQueues = append(IndexQueues, IndexQueue)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllIndexQueueResponse{IndexQueue: IndexQueues, Pagination: pageRes}, nil
}

func (k Keeper) IndexQueue(c context.Context, req *types.QueryGetIndexQueueRequest) (*types.QueryGetIndexQueueResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetIndexQueue(
		ctx,
		req.Token0,
		req.Token1,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetIndexQueueResponse{IndexQueue: val}, nil
}
