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

func (k Keeper) LimitOrderTrancheUserAll(c context.Context, req *types.QueryAllLimitOrderTrancheUserRequest) (*types.QueryAllLimitOrderTrancheUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var LimitOrderTrancheUsers []types.LimitOrderTrancheUser
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	LimitOrderTrancheUserStore := prefix.NewStore(store, types.KeyPrefix(types.LimitOrderTrancheUserKeyPrefix))

	pageRes, err := query.Paginate(LimitOrderTrancheUserStore, req.Pagination, func(key []byte, value []byte) error {
		var LimitOrderTrancheUser types.LimitOrderTrancheUser
		if err := k.cdc.Unmarshal(value, &LimitOrderTrancheUser); err != nil {
			return err
		}

		LimitOrderTrancheUsers = append(LimitOrderTrancheUsers, LimitOrderTrancheUser)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLimitOrderTrancheUserResponse{LimitOrderTrancheUser: LimitOrderTrancheUsers, Pagination: pageRes}, nil
}

func (k Keeper) LimitOrderTrancheUser(c context.Context, req *types.QueryGetLimitOrderTrancheUserRequest) (*types.QueryGetLimitOrderTrancheUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetLimitOrderTrancheUser(
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

	return &types.QueryGetLimitOrderTrancheUserResponse{LimitOrderTrancheUser: val}, nil
}
