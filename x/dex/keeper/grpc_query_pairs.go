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

func (k Keeper) PairsAll(c context.Context, req *types.QueryAllPairsRequest) (*types.QueryAllPairsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var pairss []types.Pairs
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	pairsStore := prefix.NewStore(store, types.PairsPrefix())

	pageRes, err := query.Paginate(pairsStore, req.Pagination, func(key []byte, value []byte) error {
		var pairs types.Pairs
		if err := k.cdc.Unmarshal(value, &pairs); err != nil {
			return err
		}

		pairss = append(pairss, pairs)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPairsResponse{Pairs: pairss, Pagination: pageRes}, nil
}

func (k Keeper) Pairs(c context.Context, req *types.QueryGetPairsRequest) (*types.QueryGetPairsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetPairs(
		ctx,
		req.Token0,
		req.Token1,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetPairsResponse{Pairs: val}, nil
}
