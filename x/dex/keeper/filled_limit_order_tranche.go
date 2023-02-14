package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
)

// SetFilledLimitOrderTranche set a specific filledLimitOrderTranche in the store from its index
func (k Keeper) SetFilledLimitOrderTranche(ctx sdk.Context, filledLimitOrderTranche types.FilledLimitOrderTranche) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FilledLimitOrderTrancheKeyPrefix))
	b := k.cdc.MustMarshal(&filledLimitOrderTranche)
	store.Set(types.FilledLimitOrderTrancheKey(
		filledLimitOrderTranche.PairId,
		filledLimitOrderTranche.TokenIn,
		filledLimitOrderTranche.TickIndex,
		filledLimitOrderTranche.TrancheKey,
	), b)
}

// GetFilledLimitOrderTranche returns a filledLimitOrderTranche from its index
func (k Keeper) GetFilledLimitOrderTranche(
	ctx sdk.Context,
	pairId *types.PairId,
	tokenIn string,
	tickIndex int64,
	trancheKey string,

) (val types.FilledLimitOrderTranche, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FilledLimitOrderTrancheKeyPrefix))

	b := store.Get(types.FilledLimitOrderTrancheKey(
		pairId,
		tokenIn,
		tickIndex,
		trancheKey,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetNewestFilledLimitOrderTranche(sdkCtx sdk.Context, pairId *types.PairId, tokenIn string, tickIndex int64) (types.FilledLimitOrderTranche, bool) {
	prefixStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.FilledLimitOrderTranchePrefix(pairId, tokenIn, tickIndex))
	iter := sdk.KVStoreReversePrefixIterator(prefixStore, []byte{})

	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var tranche types.FilledLimitOrderTranche
		k.cdc.MustUnmarshal(iter.Value(), &tranche)
		return tranche, true
	}
	return types.FilledLimitOrderTranche{}, false
}

// RemoveFilledLimitOrderTranche removes a filledLimitOrderTranche from the store
func (k Keeper) RemoveFilledLimitOrderTranche(
	ctx sdk.Context,
	pairId *types.PairId,
	tokenIn string,
	tickIndex int64,
	trancheKey string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FilledLimitOrderTrancheKeyPrefix))
	store.Delete(types.FilledLimitOrderTrancheKey(
		pairId,
		tokenIn,
		tickIndex,
		trancheKey,
	))
}

// GetAllFilledLimitOrderTranche returns all filledLimitOrderTranche
func (k Keeper) GetAllFilledLimitOrderTranche(ctx sdk.Context) (list []types.FilledLimitOrderTranche) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FilledLimitOrderTrancheKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.FilledLimitOrderTranche
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
