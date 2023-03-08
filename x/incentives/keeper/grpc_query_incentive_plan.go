package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/duality-labs/duality/x/incentives/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) IncentivePlanAll(c context.Context, req *types.QueryAllIncentivePlanRequest) (*types.QueryAllIncentivePlanResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var incentivePlans []types.IncentivePlan
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	incentivePlanStore := prefix.NewStore(store, types.KeyPrefix(types.IncentivePlanKeyPrefix))

	pageRes, err := query.Paginate(incentivePlanStore, req.Pagination, func(key []byte, value []byte) error {
		var incentivePlan types.IncentivePlan
		if err := k.cdc.Unmarshal(value, &incentivePlan); err != nil {
			return err
		}

		incentivePlans = append(incentivePlans, incentivePlan)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllIncentivePlanResponse{IncentivePlan: incentivePlans, Pagination: pageRes}, nil
}

func (k Keeper) IncentivePlan(c context.Context, req *types.QueryGetIncentivePlanRequest) (*types.QueryGetIncentivePlanResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetIncentivePlan(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetIncentivePlanResponse{IncentivePlan: val}, nil
}
