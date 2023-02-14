package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
)

// GetTokensCount get the total number of tokens
func (k Keeper) GetTokensCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TokensCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetTokensCount set the total number of tokens
func (k Keeper) SetTokensCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TokensCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendTokens appends a tokens in the store with a new id and update the count
func (k Keeper) AppendTokens(
	ctx sdk.Context,
	tokens types.Tokens,
) uint64 {
	// Create the tokens
	count := k.GetTokensCount(ctx)

	// Set the ID of the appended value
	tokens.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokensKey))
	appendedValue := k.cdc.MustMarshal(&tokens)
	store.Set(GetTokensIDBytes(tokens.Id), appendedValue)

	// Update tokens count
	k.SetTokensCount(ctx, count+1)

	return count
}

// SetTokens set a specific tokens in the store
func (k Keeper) SetTokens(ctx sdk.Context, tokens types.Tokens) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokensKey))
	b := k.cdc.MustMarshal(&tokens)
	store.Set(GetTokensIDBytes(tokens.Id), b)
}

// GetTokens returns a tokens from its id
func (k Keeper) GetTokens(ctx sdk.Context, id uint64) (val types.Tokens, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokensKey))
	b := store.Get(GetTokensIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTokens removes a tokens from the store
func (k Keeper) RemoveTokens(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokensKey))
	store.Delete(GetTokensIDBytes(id))
}

// GetAllTokens returns all tokens
func (k Keeper) GetAllTokens(ctx sdk.Context) (list []types.Tokens) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokensKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Tokens
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetTokensIDBytes returns the byte representation of the ID
func GetTokensIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetTokensIDFromBytes returns ID in uint64 format from a byte array
func GetTokensIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
