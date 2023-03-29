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

// Returns all ACTIVE limit order tranches for a given pairId/tokenIn combination
// Does NOT return inactiveLimitOrderTranches
func (k Keeper) LimitOrderTrancheAll(c context.Context, req *types.QueryAllLimitOrderTrancheRequest) (*types.QueryAllLimitOrderTrancheResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var LimitOrderTranches []types.LimitOrderTranche
	ctx := sdk.UnwrapSDKContext(c)

	pairId, err := StringToPairId(req.PairId)
	if err != nil {
		return nil, err
	}
	store := ctx.KVStore(k.storeKey)
	LimitOrderTrancheStore := prefix.NewStore(store, types.TickLiquidityPrefix(pairId, req.TokenIn))

	pageRes, err := query.FilteredPaginate(LimitOrderTrancheStore, req.Pagination, func(key, value []byte, accum bool) (hit bool, err error) {
		var tick types.TickLiquidity

		if err := k.cdc.Unmarshal(value, &tick); err != nil {
			return false, err
		}
		tranche := tick.GetLimitOrderTranche()
		// Check if this is a LimitOrderTranche and not PoolReserves
		if tranche != nil {
			if accum {
				LimitOrderTranches = append(LimitOrderTranches, *tranche)
			}
			return true, nil
		} else {
			return false, nil
		}
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLimitOrderTrancheResponse{LimitOrderTranche: LimitOrderTranches, Pagination: pageRes}, nil
}

// Returns a specific limit order tranche either from the tickLiquidity index or from the FillLimitOrderTranche index
func (k Keeper) LimitOrderTranche(c context.Context, req *types.QueryGetLimitOrderTrancheRequest) (*types.QueryGetLimitOrderTrancheResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	pairId, err := StringToPairId(req.PairId)
	if err != nil {
		return nil, err
	}
	val, _, found := k.FindLimitOrderTranche(ctx, pairId, req.TickIndex, req.TokenIn, req.TrancheKey)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetLimitOrderTrancheResponse{LimitOrderTranche: val}, nil
}
