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

func (k Keeper) AdjMatrixAll(c context.Context, req *types.QueryAllAdjMatrixRequest) (*types.QueryAllAdjMatrixResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var adjMatrixs []types.AdjMatrix
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	adjMatrixStore := prefix.NewStore(store, types.KeyPrefix(types.AdjMatrixKey))

	pageRes, err := query.Paginate(adjMatrixStore, req.Pagination, func(key []byte, value []byte) error {
		var adjMatrix types.AdjMatrix
		if err := k.cdc.Unmarshal(value, &adjMatrix); err != nil {
			return err
		}

		adjMatrixs = append(adjMatrixs, adjMatrix)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAdjMatrixResponse{AdjMatrix: adjMatrixs, Pagination: pageRes}, nil
}

func (k Keeper) AdjMatrix(c context.Context, req *types.QueryGetAdjMatrixRequest) (*types.QueryGetAdjMatrixResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	adjMatrix, found := k.GetAdjMatrix(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetAdjMatrixResponse{AdjMatrix: adjMatrix}, nil
}
