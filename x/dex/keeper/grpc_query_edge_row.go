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

func (k Keeper) EdgeRowAll(c context.Context, req *types.QueryAllEdgeRowRequest) (*types.QueryAllEdgeRowResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var edgeRows []types.EdgeRow
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	edgeRowStore := prefix.NewStore(store, types.KeyPrefix(types.EdgeRowKey))

	pageRes, err := query.Paginate(edgeRowStore, req.Pagination, func(key []byte, value []byte) error {
		var edgeRow types.EdgeRow
		if err := k.cdc.Unmarshal(value, &edgeRow); err != nil {
			return err
		}

		edgeRows = append(edgeRows, edgeRow)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllEdgeRowResponse{EdgeRow: edgeRows, Pagination: pageRes}, nil
}

func (k Keeper) EdgeRow(c context.Context, req *types.QueryGetEdgeRowRequest) (*types.QueryGetEdgeRowResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	edgeRow, found := k.GetEdgeRow(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetEdgeRowResponse{EdgeRow: edgeRow}, nil
}
