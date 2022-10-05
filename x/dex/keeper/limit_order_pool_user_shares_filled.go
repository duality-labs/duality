package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetLimitOrderPoolUserSharesWithdrawn set a specific limitOrderPoolUserSharesWithdrawn in the store from its index
func (k Keeper) SetLimitOrderPoolUserSharesWithdrawn(ctx sdk.Context, pairId string, limitOrderPoolUserSharesWithdrawn types.LimitOrderPoolUserSharesWithdrawn) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LimitOrderUserSharesWithdrawnPrefix(pairId))
	b := k.cdc.MustMarshal(&limitOrderPoolUserSharesWithdrawn)
	store.Set(types.LimitOrderPoolUserSharesWithdrawnKey(
		limitOrderPoolUserSharesWithdrawn.TickIndex,
		limitOrderPoolUserSharesWithdrawn.Token,
		limitOrderPoolUserSharesWithdrawn.Count,
		limitOrderPoolUserSharesWithdrawn.Address,
	), b)
}

// GetLimitOrderPoolUserSharesWithdrawn returns a limitOrderPoolUserSharesWithdrawn from its index
func (k Keeper) GetLimitOrderPoolUserSharesWithdrawn(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,
	token string,
	count uint64,
	address string,

) (val types.LimitOrderPoolUserSharesWithdrawn, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LimitOrderUserSharesWithdrawnPrefix(pairId))

	b := store.Get(types.LimitOrderPoolUserSharesWithdrawnKey(
		tickIndex,
		token,
		count,
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLimitOrderPoolUserSharesWithdrawn removes a limitOrderPoolUserSharesWithdrawn from the store
func (k Keeper) RemoveLimitOrderPoolUserSharesWithdrawn(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,
	token string,
	count uint64,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LimitOrderUserSharesWithdrawnPrefix(pairId))
	store.Delete(types.LimitOrderPoolUserSharesWithdrawnKey(
		tickIndex,
		token,
		count,
		address,
	))
}

// GetAllLimitOrderPoolUserSharesWithdrawn returns all limitOrderPoolUserSharesWithdrawn
func (k Keeper) GetAllLimitOrderPoolUserSharesWithdrawn(ctx sdk.Context) (list []types.LimitOrderPoolUserSharesWithdrawn) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderPoolUserSharesWithdrawnKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LimitOrderPoolUserSharesWithdrawn
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
