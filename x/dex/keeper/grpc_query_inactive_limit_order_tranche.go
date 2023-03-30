package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/duality-labs/duality/x/dex/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) InactiveLimitOrderTrancheAll(
	c context.Context,
	req *types.QueryAllInactiveLimitOrderTrancheRequest,
) (*types.QueryAllInactiveLimitOrderTrancheResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var inactiveLimitOrderTranches []types.LimitOrderTranche
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	inactiveLimitOrderTrancheStore := prefix.NewStore(store, types.KeyPrefix(types.InactiveLimitOrderTrancheKeyPrefix))

	pageRes, err := query.Paginate(inactiveLimitOrderTrancheStore, req.Pagination, func(key, value []byte) error {
		var inactiveLimitOrderTranche types.LimitOrderTranche
		if err := k.cdc.Unmarshal(value, &inactiveLimitOrderTranche); err != nil {
			return err
		}

		inactiveLimitOrderTranches = append(inactiveLimitOrderTranches, inactiveLimitOrderTranche)

		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllInactiveLimitOrderTrancheResponse{
		InactiveLimitOrderTranche: inactiveLimitOrderTranches,
		Pagination:                pageRes,
	}, nil
}

func (k Keeper) InactiveLimitOrderTranche(
	c context.Context,
	req *types.QueryGetInactiveLimitOrderTrancheRequest,
) (*types.QueryGetInactiveLimitOrderTrancheResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	pairID, err := StringToPairID(req.PairID)
	if err != nil {
		return nil, err
	}
	val, found := k.GetInactiveLimitOrderTranche(
		ctx,
		pairID,
		req.TokenIn,
		req.TickIndex,
		req.TrancheKey,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetInactiveLimitOrderTrancheResponse{InactiveLimitOrderTranche: val}, nil
}
