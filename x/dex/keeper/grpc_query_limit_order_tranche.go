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

func (k Keeper) LimitOrderTrancheAll(c context.Context, req *types.QueryAllLimitOrderTrancheRequest) (*types.QueryAllLimitOrderTrancheResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var LimitOrderTranches []types.LimitOrderTranche
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	LimitOrderTrancheStore := prefix.NewStore(store, types.KeyPrefix(types.LimitOrderTrancheKeyPrefix))

	pageRes, err := query.Paginate(LimitOrderTrancheStore, req.Pagination, func(key []byte, value []byte) error {
		var LimitOrderTranche types.LimitOrderTranche
		if err := k.cdc.Unmarshal(value, &LimitOrderTranche); err != nil {
			return err
		}

		LimitOrderTranches = append(LimitOrderTranches, LimitOrderTranche)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLimitOrderTrancheResponse{LimitOrderTranche: LimitOrderTranches, Pagination: pageRes}, nil
}

func (k Keeper) LimitOrderTranche(c context.Context, req *types.QueryGetLimitOrderTrancheRequest) (*types.QueryGetLimitOrderTrancheResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetLimitOrderTranche(
		ctx,
		req.PairId,
		req.TickIndex,
		req.Token,
		req.TrancheIndex,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetLimitOrderTrancheResponse{LimitOrderTranche: val}, nil
}
