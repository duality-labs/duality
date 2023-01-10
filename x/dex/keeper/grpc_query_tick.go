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

func (k Keeper) TickAll(c context.Context, req *types.QueryAllTickRequest) (*types.QueryAllTickResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var Ticks []types.Tick
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	TickStore := prefix.NewStore(store, types.KeyPrefix(types.BaseTickKeyPrefix))

	pageRes, err := query.Paginate(TickStore, req.Pagination, func(key []byte, value []byte) error {
		var Tick types.Tick
		if err := k.cdc.Unmarshal(value, &Tick); err != nil {
			return err
		}

		Ticks = append(Ticks, Tick)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTickResponse{Tick: Ticks, Pagination: pageRes}, nil
}

func (k Keeper) Tick(c context.Context, req *types.QueryGetTickRequest) (*types.QueryGetTickResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetTick(
		ctx,
		req.PairId,
		req.TickIndex,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTickResponse{Tick: val}, nil
}
