package keeper

import (
	"context"

	"github.com/duality-labs/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) FilledLimitOrderTrancheAll(c context.Context, req *types.QueryAllFilledLimitOrderTrancheRequest) (*types.QueryAllFilledLimitOrderTrancheResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var filledLimitOrderTranches []types.FilledLimitOrderTranche
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	filledLimitOrderTrancheStore := prefix.NewStore(store, types.KeyPrefix(types.FilledLimitOrderTrancheKeyPrefix))

	pageRes, err := query.Paginate(filledLimitOrderTrancheStore, req.Pagination, func(key []byte, value []byte) error {
		var filledLimitOrderTranche types.FilledLimitOrderTranche
		if err := k.cdc.Unmarshal(value, &filledLimitOrderTranche); err != nil {
			return err
		}

		filledLimitOrderTranches = append(filledLimitOrderTranches, filledLimitOrderTranche)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFilledLimitOrderTrancheResponse{FilledLimitOrderTranche: filledLimitOrderTranches, Pagination: pageRes}, nil
}

func (k Keeper) FilledLimitOrderTranche(c context.Context, req *types.QueryGetFilledLimitOrderTrancheRequest) (*types.QueryGetFilledLimitOrderTrancheResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	pairId, err := StringToPairId(req.PairId)
	if err != nil {
		return nil, err
	}
	val, found := k.GetFilledLimitOrderTranche(
		ctx,
		pairId,
		req.TokenIn,
		req.TickIndex,
		req.TrancheIndex,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetFilledLimitOrderTrancheResponse{FilledLimitOrderTranche: val}, nil
}
