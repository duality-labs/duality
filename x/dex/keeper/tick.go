package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetTick set a specific tick in the store from its index
func (k Keeper) SetTick(ctx sdk.Context, tick types.Tick) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TickKeyPrefix))
	b := k.cdc.MustMarshal(&tick)
	store.Set(types.TickKey(
		tick.Token0,
		tick.Token1,
		tick.Price,
		tick.Fee,
	), b)
}

// GetTick returns a tick from its index
func (k Keeper) GetTick(
	ctx sdk.Context,
	token0 string,
	token1 string,
	price string,
	fee uint64,

) (val types.Tick, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TickKeyPrefix))

	b := store.Get(types.TickKey(
		token0,
		token1,
		price,
		fee,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTick removes a tick from the store
func (k Keeper) RemoveTick(
	ctx sdk.Context,
	token0 string,
	token1 string,
	price string,
	fee uint64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TickKeyPrefix))
	store.Delete(types.TickKey(
		token0,
		token1,
		price,
		fee,
	))
}

// GetAllTick returns all tick
func (k Keeper) GetAllTick(ctx sdk.Context) (list []types.Tick) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TickKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Tick
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
