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

func (k Keeper) LimitOrderPoolUserSharesWithdrawnAll(c context.Context, req *types.QueryAllLimitOrderPoolUserSharesWithdrawnRequest) (*types.QueryAllLimitOrderPoolUserSharesWithdrawnResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var limitOrderPoolUserSharesWithdrawns []types.LimitOrderPoolUserSharesWithdrawn
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	limitOrderPoolUserSharesWithdrawnStore := prefix.NewStore(store, types.KeyPrefix(types.LimitOrderPoolUserSharesWithdrawnKeyPrefix))

	pageRes, err := query.Paginate(limitOrderPoolUserSharesWithdrawnStore, req.Pagination, func(key []byte, value []byte) error {
		var limitOrderPoolUserSharesWithdrawn types.LimitOrderPoolUserSharesWithdrawn
		if err := k.cdc.Unmarshal(value, &limitOrderPoolUserSharesWithdrawn); err != nil {
			return err
		}

		limitOrderPoolUserSharesWithdrawns = append(limitOrderPoolUserSharesWithdrawns, limitOrderPoolUserSharesWithdrawn)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLimitOrderPoolUserSharesWithdrawnResponse{LimitOrderPoolUserSharesWithdrawn: limitOrderPoolUserSharesWithdrawns, Pagination: pageRes}, nil
}

func (k Keeper) LimitOrderPoolUserSharesWithdrawn(c context.Context, req *types.QueryGetLimitOrderPoolUserSharesWithdrawnRequest) (*types.QueryGetLimitOrderPoolUserSharesWithdrawnResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetLimitOrderPoolUserSharesWithdrawn(
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

	return &types.QueryGetLimitOrderPoolUserSharesWithdrawnResponse{LimitOrderPoolUserSharesWithdrawn: val}, nil
}
