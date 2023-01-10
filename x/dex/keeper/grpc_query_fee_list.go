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

func (k Keeper) FeeListAll(c context.Context, req *types.QueryAllFeeListRequest) (*types.QueryAllFeeListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var feeLists []types.FeeList
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	feeListStore := prefix.NewStore(store, types.KeyPrefix(types.FeeListKey))

	pageRes, err := query.Paginate(feeListStore, req.Pagination, func(key []byte, value []byte) error {
		var feeList types.FeeList
		if err := k.cdc.Unmarshal(value, &feeList); err != nil {
			return err
		}

		feeLists = append(feeLists, feeList)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFeeListResponse{FeeList: feeLists, Pagination: pageRes}, nil
}

func (k Keeper) FeeList(c context.Context, req *types.QueryGetFeeListRequest) (*types.QueryGetFeeListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	feeList, found := k.GetFeeList(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetFeeListResponse{FeeList: feeList}, nil
}
