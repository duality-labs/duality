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

func (k Keeper) SharesAll(c context.Context, req *types.QueryAllSharesRequest) (*types.QueryAllSharesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var sharess []types.Shares
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	sharesStore := prefix.NewStore(store, types.KeyPrefix(types.SharesKeyPrefix))

	pageRes, err := query.Paginate(sharesStore, req.Pagination, func(key []byte, value []byte) error {
		var shares types.Shares
		if err := k.cdc.Unmarshal(value, &shares); err != nil {
			return err
		}

		sharess = append(sharess, shares)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSharesResponse{Shares: sharess, Pagination: pageRes}, nil
}

func (k Keeper) Shares(c context.Context, req *types.QueryGetSharesRequest) (*types.QueryGetSharesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetShares(
		ctx,
		req.Address,
		req.PairId,
		req.TickIndex,
		req.Fee,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetSharesResponse{Shares: val}, nil
}
