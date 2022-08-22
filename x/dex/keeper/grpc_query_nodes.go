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

func (k Keeper) NodesAll(c context.Context, req *types.QueryAllNodesRequest) (*types.QueryAllNodesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var nodess []types.Nodes
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	nodesStore := prefix.NewStore(store, types.KeyPrefix(types.NodesKeyPrefix))

	pageRes, err := query.Paginate(nodesStore, req.Pagination, func(key []byte, value []byte) error {
		var nodes types.Nodes
		if err := k.cdc.Unmarshal(value, &nodes); err != nil {
			return err
		}

		nodess = append(nodess, nodes)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNodesResponse{Nodes: nodess, Pagination: pageRes}, nil
}

func (k Keeper) Nodes(c context.Context, req *types.QueryGetNodesRequest) (*types.QueryGetNodesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetNodes(
		ctx,
		req.Node,
		req.OutgoingEdges,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetNodesResponse{Nodes: val}, nil
}
