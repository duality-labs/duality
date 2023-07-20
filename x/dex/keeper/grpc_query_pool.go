package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Pool(
	goCtx context.Context,
	req *types.QueryPoolRequest,
) (*types.QueryPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	pairID, err := types.NewPairIDFromCanonicalString(req.PairID)
	if err != nil {
		return nil, err
	}

	pool, found := k.GetPool(ctx, pairID, req.TickIndex, req.Fee)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryPoolResponse{Pool: pool}, nil
}
