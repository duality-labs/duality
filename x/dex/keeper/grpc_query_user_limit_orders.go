package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) UserLimitOrdersAll(goCtx context.Context, req *types.QueryAllUserLimitOrdersRequest) (*types.QueryAllUserLimitOrdersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}
	profile := NewUserProfile(addr)

	return &types.QueryAllUserLimitOrdersResponse{
		LimitOrders: profile.GetAllLimitOrders(goCtx, k),
	}, nil
}
