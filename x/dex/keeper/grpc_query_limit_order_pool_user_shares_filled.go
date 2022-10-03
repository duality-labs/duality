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

func (k Keeper) LimitOrderPoolUserSharesFilledAll(c context.Context, req *types.QueryAllLimitOrderPoolUserSharesFilledRequest) (*types.QueryAllLimitOrderPoolUserSharesFilledResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var limitOrderPoolUserSharesFilleds []types.LimitOrderPoolUserSharesFilled
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	limitOrderPoolUserSharesFilledStore := prefix.NewStore(store, types.KeyPrefix(types.LimitOrderPoolUserSharesFilledKeyPrefix))

	pageRes, err := query.Paginate(limitOrderPoolUserSharesFilledStore, req.Pagination, func(key []byte, value []byte) error {
		var limitOrderPoolUserSharesFilled types.LimitOrderPoolUserSharesFilled
		if err := k.cdc.Unmarshal(value, &limitOrderPoolUserSharesFilled); err != nil {
			return err
		}

		limitOrderPoolUserSharesFilleds = append(limitOrderPoolUserSharesFilleds, limitOrderPoolUserSharesFilled)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLimitOrderPoolUserSharesFilledResponse{LimitOrderPoolUserSharesFilled: limitOrderPoolUserSharesFilleds, Pagination: pageRes}, nil
}

func (k Keeper) LimitOrderPoolUserSharesFilled(c context.Context, req *types.QueryGetLimitOrderPoolUserSharesFilledRequest) (*types.QueryGetLimitOrderPoolUserSharesFilledResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetLimitOrderPoolUserSharesFilled(
		ctx,
		req.PairId,
		req.TickIndex,
		req.Token,
		req.Count,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetLimitOrderPoolUserSharesFilledResponse{LimitOrderPoolUserSharesFilled: val}, nil
}
