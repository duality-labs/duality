package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetTicks set a specific ticks in the store from its index
func (k Keeper) SetTicks(ctx sdk.Context, token0 string, token1 string, ticks types.Ticks) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TicksKeyPrefix))
	b := k.cdc.MustMarshal(&ticks)
	store.Set(types.TicksKey(
		token0,
		token1,
		ticks.Price,
		ticks.Fee,
		ticks.Direction,
		ticks.OrderType,
	), b)
}

// GetTicks returns a ticks from its index
func (k Keeper) GetTicks(
	ctx sdk.Context,
	token0 string,
	token1 string,
	price string,
	fee string,
	direction string,
	orderType string,

) (val types.Ticks, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TicksKeyPrefix))

	b := store.Get(types.TicksKey(
		token0,
		token1,
		price,
		fee,
		direction,
		orderType,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTicks removes a ticks from the store
func (k Keeper) RemoveTicks(
	ctx sdk.Context,
	token0 string,
	token1 string,
	price string,
	fee string,
	direction string,
	orderType string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TicksKeyPrefix))
	store.Delete(types.TicksKey(
		token0,
		token1,
		price,
		fee,
		direction,
		orderType,
	))
}

// GetAllTicks returns all ticks
func (k Keeper) GetAllTicks(ctx sdk.Context) (list []types.Ticks) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TicksKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Ticks
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
