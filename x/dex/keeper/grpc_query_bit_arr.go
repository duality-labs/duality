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

func (k Keeper) BitArrAll(c context.Context, req *types.QueryAllBitArrRequest) (*types.QueryAllBitArrResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var bitArrs []types.BitArr
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	bitArrStore := prefix.NewStore(store, types.KeyPrefix(types.BitArrKey))

	pageRes, err := query.Paginate(bitArrStore, req.Pagination, func(key []byte, value []byte) error {
		var bitArr types.BitArr
		if err := k.cdc.Unmarshal(value, &bitArr); err != nil {
			return err
		}

		bitArrs = append(bitArrs, bitArr)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllBitArrResponse{BitArr: bitArrs, Pagination: pageRes}, nil
}

func (k Keeper) BitArr(c context.Context, req *types.QueryGetBitArrRequest) (*types.QueryGetBitArrResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	bitArr, found := k.GetBitArr(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetBitArrResponse{BitArr: bitArr}, nil
}
