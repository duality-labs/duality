package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/duality-labs/duality/x/dex/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TokensAll(c context.Context, req *types.QueryAllTokensRequest) (*types.QueryAllTokensResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var tokenss []types.Tokens
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	tokensStore := prefix.NewStore(store, types.KeyPrefix(types.TokensKey))

	pageRes, err := query.Paginate(tokensStore, req.Pagination, func(key []byte, value []byte) error {
		var tokens types.Tokens
		if err := k.cdc.Unmarshal(value, &tokens); err != nil {
			return err
		}

		tokenss = append(tokenss, tokens)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTokensResponse{Tokens: tokenss, Pagination: pageRes}, nil
}

func (k Keeper) Tokens(c context.Context, req *types.QueryGetTokensRequest) (*types.QueryGetTokensResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	tokens, found := k.GetTokens(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetTokensResponse{Tokens: tokens}, nil
}
