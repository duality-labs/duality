package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/cwhooks/types"
)

// GetHookCount get the total number of hook
func (k Keeper) GetHookCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.HookCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetHookCount set the total number of hook
func (k Keeper) SetHookCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.HookCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendHook appends a hook in the store with a new id and update the count
func (k Keeper) AppendHook(
	ctx sdk.Context,
	hook types.Hook,
) uint64 {
	// Create the hook
	count := k.GetHookCount(ctx)

	// Set the ID of the appended value
	hook.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HookKeyPrefix))
	appendedValue := k.cdc.MustMarshal(&hook)
	store.Set(types.HookKey(hook.TriggerKey, hook.TriggerValue, hook.Id), appendedValue)

	// Update hook count
	k.SetHookCount(ctx, count+1)

	return count
}

// SetHook set a specific hook in the store
func (k Keeper) SetHook(ctx sdk.Context, hook types.Hook) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HookKeyPrefix))
	b := k.cdc.MustMarshal(&hook)
	store.Set(types.HookKey(hook.TriggerKey, hook.TriggerValue, hook.Id), b)
}

// GetHook returns a hook from its id
func (k Keeper) GetHook(ctx sdk.Context, triggerKey string, triggerValue string, ID uint64) (val types.Hook, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HookKeyPrefix))
	b := store.Get(types.HookKey(triggerKey, triggerValue, ID))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveHook removes a hook from the store
func (k Keeper) RemoveHook(ctx sdk.Context, triggerKey string, triggerValue string, ID uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HookKeyPrefix))
	store.Delete(types.HookKey(triggerKey, triggerValue, ID))
}

// GetAllHook returns all hook
func (k Keeper) GetAllHook(ctx sdk.Context) (list []types.Hook) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HookKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Hook
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
