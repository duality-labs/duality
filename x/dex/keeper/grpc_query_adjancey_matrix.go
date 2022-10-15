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

func (k Keeper) AdjanceyMatrixAll(c context.Context, req *types.QueryAllAdjanceyMatrixRequest) (*types.QueryAllAdjanceyMatrixResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var adjanceyMatrixs []types.AdjanceyMatrix
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	adjanceyMatrixStore := prefix.NewStore(store, types.KeyPrefix(types.AdjanceyMatrixKey))

	pageRes, err := query.Paginate(adjanceyMatrixStore, req.Pagination, func(key []byte, value []byte) error {
		var adjanceyMatrix types.AdjanceyMatrix
		if err := k.cdc.Unmarshal(value, &adjanceyMatrix); err != nil {
			return err
		}

		adjanceyMatrixs = append(adjanceyMatrixs, adjanceyMatrix)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAdjanceyMatrixResponse{AdjanceyMatrix: adjanceyMatrixs, Pagination: pageRes}, nil
}

func (k Keeper) AdjanceyMatrix(c context.Context, req *types.QueryGetAdjanceyMatrixRequest) (*types.QueryGetAdjanceyMatrixResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	adjanceyMatrix, found := k.GetAdjanceyMatrix(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetAdjanceyMatrixResponse{AdjanceyMatrix: adjanceyMatrix}, nil
}
