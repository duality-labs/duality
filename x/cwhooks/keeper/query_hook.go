package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/duality-labs/duality/x/cwhooks/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) HookAll(goCtx context.Context, req *types.QueryAllHookRequest) (*types.QueryAllHookResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var hooks []types.Hook
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	hookStore := prefix.NewStore(store, types.KeyPrefix(types.HookKeyPrefix))

	pageRes, err := query.Paginate(hookStore, req.Pagination, func(key []byte, value []byte) error {
		var hook types.Hook
		if err := k.cdc.Unmarshal(value, &hook); err != nil {
			return err
		}

		hooks = append(hooks, hook)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllHookResponse{Hook: hooks, Pagination: pageRes}, nil
}

func (k Keeper) Hook(goCtx context.Context, req *types.QueryGetHookRequest) (*types.QueryGetHookResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	hook, found := k.GetHook(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetHookResponse{Hook: hook}, nil
}
