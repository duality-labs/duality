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

func (k Keeper) FeeTierAll(c context.Context, req *types.QueryAllFeeTierRequest) (*types.QueryAllFeeTierResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var FeeTiers []types.FeeTier
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	FeeTierStore := prefix.NewStore(store, types.KeyPrefix(types.FeeTierKey))

	pageRes, err := query.Paginate(FeeTierStore, req.Pagination, func(key []byte, value []byte) error {
		var FeeTier types.FeeTier
		if err := k.cdc.Unmarshal(value, &FeeTier); err != nil {
			return err
		}

		FeeTiers = append(FeeTiers, FeeTier)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFeeTierResponse{FeeTier: FeeTiers, Pagination: pageRes}, nil
}

func (k Keeper) FeeTier(c context.Context, req *types.QueryGetFeeTierRequest) (*types.QueryGetFeeTierResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	FeeTier, found := k.GetFeeTier(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetFeeTierResponse{FeeTier: FeeTier}, nil
}
