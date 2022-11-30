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

func (k Keeper) TradingPairAll(c context.Context, req *types.QueryAllTradingPairRequest) (*types.QueryAllTradingPairResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var TradingPairs []types.TradingPair
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	TradingPairStore := prefix.NewStore(store, types.KeyPrefix(types.TradingPairKeyPrefix))

	pageRes, err := query.Paginate(TradingPairStore, req.Pagination, func(key []byte, value []byte) error {
		var TradingPair types.TradingPair
		if err := k.cdc.Unmarshal(value, &TradingPair); err != nil {
			return err
		}

		TradingPairs = append(TradingPairs, TradingPair)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTradingPairResponse{TradingPair: TradingPairs, Pagination: pageRes}, nil
}

func (k Keeper) TradingPair(c context.Context, req *types.QueryGetTradingPairRequest) (*types.QueryGetTradingPairResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetTradingPair(
		ctx,
		req.PairId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTradingPairResponse{TradingPair: val}, nil
}
