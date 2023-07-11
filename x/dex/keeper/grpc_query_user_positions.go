package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetUserPositions(
	goCtx context.Context,
	req *types.QueryGetUserPositionsRequest,
) (*types.QueryGetUserPositionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return &types.QueryGetUserPositionsResponse{}, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	deposits := k.GetAllDepositsForAddress(ctx, addr)
	limitOrders := k.GetAllLimitOrderTrancheUserForAddress(ctx, addr)

	return &types.QueryGetUserPositionsResponse{
		UserPositions: &types.UserPositions{
			PoolDeposits: deposits,
			LimitOrders:  limitOrders,
		},
	}, nil
}
