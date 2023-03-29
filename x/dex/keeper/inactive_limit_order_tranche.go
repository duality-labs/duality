package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
)

// SetInactiveLimitOrderTranche set a specific inactiveLimitOrderTranche in the store from its index
func (k Keeper) SetInactiveLimitOrderTranche(ctx sdk.Context, inactiveLimitOrderTranche types.LimitOrderTranche) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InactiveLimitOrderTrancheKeyPrefix))
	b := k.cdc.MustMarshal(&inactiveLimitOrderTranche)
	store.Set(types.InactiveLimitOrderTrancheKey(
		inactiveLimitOrderTranche.PairId,
		inactiveLimitOrderTranche.TokenIn,
		inactiveLimitOrderTranche.TickIndex,
		inactiveLimitOrderTranche.TrancheKey,
	), b)
}

// GetInactiveLimitOrderTranche returns a inactiveLimitOrderTranche from its index
func (k Keeper) GetInactiveLimitOrderTranche(
	ctx sdk.Context,
	pairId *types.PairId,
	tokenIn string,
	tickIndex int64,
	trancheKey string,
) (val types.LimitOrderTranche, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InactiveLimitOrderTrancheKeyPrefix))

	b := store.Get(types.InactiveLimitOrderTrancheKey(
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

// RemoveInactiveLimitOrderTranche removes a inactiveLimitOrderTranche from the store
func (k Keeper) RemoveInactiveLimitOrderTranche(
	ctx sdk.Context,
	pairId *types.PairId,
	tokenIn string,
	tickIndex int64,
	trancheKey string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InactiveLimitOrderTrancheKeyPrefix))
	store.Delete(types.InactiveLimitOrderTrancheKey(
		pairId,
		tokenIn,
		tickIndex,
		trancheKey,
	))
}

// GetAllInactiveLimitOrderTranche returns all inactiveLimitOrderTranche
func (k Keeper) GetAllInactiveLimitOrderTranche(ctx sdk.Context) (list []types.LimitOrderTranche) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InactiveLimitOrderTrancheKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LimitOrderTranche
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) SaveInactiveTranche(sdkCtx sdk.Context, tranche types.LimitOrderTranche) {
	if tranche.HasTokenIn() || tranche.HasTokenOut() {
		k.SetInactiveLimitOrderTranche(sdkCtx, tranche)
	} else {
		k.RemoveInactiveLimitOrderTranche(sdkCtx, tranche.PairId, tranche.TokenIn, tranche.TickIndex, tranche.TrancheKey)
	}
}
