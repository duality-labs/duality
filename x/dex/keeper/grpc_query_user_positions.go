package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetUserPositions(goCtx context.Context, req *types.QueryGetUserPositionsRequest) (*types.QueryGetUserPositionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	address, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return &types.QueryGetUserPositionsResponse{}, err
	}

	userProfile := NewUserProfile(address)

	return &types.QueryGetUserPositionsResponse{
		UserPositions: userProfile.GetAllPositions(goCtx, k),
	}, nil
}
