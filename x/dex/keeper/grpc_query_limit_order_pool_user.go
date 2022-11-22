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

func (k Keeper) LimitOrderPoolUserAll(c context.Context, req *types.QueryAllLimitOrderPoolUserRequest) (*types.QueryAllLimitOrderPoolUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var LimitOrderPoolUsers []types.LimitOrderPoolUser
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	LimitOrderPoolUserStore := prefix.NewStore(store, types.KeyPrefix(types.LimitOrderPoolUserKeyPrefix))

	pageRes, err := query.Paginate(LimitOrderPoolUserStore, req.Pagination, func(key []byte, value []byte) error {
		var LimitOrderPoolUser types.LimitOrderPoolUser
		if err := k.cdc.Unmarshal(value, &LimitOrderPoolUser); err != nil {
			return err
		}

		LimitOrderPoolUsers = append(LimitOrderPoolUsers, LimitOrderPoolUser)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLimitOrderPoolUserResponse{LimitOrderPoolUser: LimitOrderPoolUsers, Pagination: pageRes}, nil
}

func (k Keeper) LimitOrderPoolUser(c context.Context, req *types.QueryGetLimitOrderPoolUserRequest) (*types.QueryGetLimitOrderPoolUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetLimitOrderPoolUser(
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

	return &types.QueryGetLimitOrderPoolUserResponse{LimitOrderPoolUser: val}, nil
}
