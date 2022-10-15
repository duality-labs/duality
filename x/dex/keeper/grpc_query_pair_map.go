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

func (k Keeper) PairMapAll(c context.Context, req *types.QueryAllPairMapRequest) (*types.QueryAllPairMapResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var pairMaps []types.PairMap
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	pairMapStore := prefix.NewStore(store, types.KeyPrefix(types.PairMapKeyPrefix))

	pageRes, err := query.Paginate(pairMapStore, req.Pagination, func(key []byte, value []byte) error {
		var pairMap types.PairMap
		if err := k.cdc.Unmarshal(value, &pairMap); err != nil {
			return err
		}

		pairMaps = append(pairMaps, pairMap)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPairMapResponse{PairMap: pairMaps, Pagination: pageRes}, nil
}

func (k Keeper) PairMap(c context.Context, req *types.QueryGetPairMapRequest) (*types.QueryGetPairMapResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetPairMap(
		ctx,
		req.PairId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetPairMapResponse{PairMap: val}, nil
}
