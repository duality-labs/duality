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

	// Add hook refs

	k.AddHookRefs(ctx, hook)

	return count
}

func (k Keeper) AddHookRefs(ctx sdk.Context, hook types.Hook) {
	hookPrimaryKey := types.HookKey(hook.TriggerKey, hook.TriggerValue, hook.Id)
	hookIDBz := sdk.Uint64ToBigEndian(hook.Id)
	hookIDStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HookIDKeyPrefix))
	hookIDStore.Set(hookIDBz, hookPrimaryKey)

	// JCP TODO: add ref for owner
}

func (k Keeper) RemoveHookRefs(ctx sdk.Context, hook types.Hook) {
	hookIDBz := sdk.Uint64ToBigEndian(hook.Id)
	hookIDStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HookIDKeyPrefix))
	hookIDStore.Delete(hookIDBz)

	// JCP TODO: remove ref for owner
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

func (k Keeper) GetHookByID(ctx sdk.Context, id uint64) (val types.Hook, found bool) {
	hookRef, found := k.GetHookRefByID(ctx, id)

	if !found {
		return val, false
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HookKeyPrefix))
	b := store.Get(hookRef)
	if b == nil {
		panic("Hook ref exists, but hook not found")
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

func (k Keeper) GetHookRefByID(ctx sdk.Context, id uint64) ([]byte, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HookIDKeyPrefix))
	hookRef := store.Get(sdk.Uint64ToBigEndian(id))
	if hookRef == nil {
		return []byte{}, false
	}

	return hookRef, true
}

// RemoveHook removes a hook from the store
func (k Keeper) RemoveHook(ctx sdk.Context, hook types.Hook) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HookKeyPrefix))
	store.Delete(types.HookKey(hook.TriggerKey, hook.TriggerValue, hook.Id))

	k.RemoveHookRefs(ctx, hook)
}

// RemoveHook removes a hook from the store
func (k Keeper) RemoveHookByID(ctx sdk.Context, ID uint64) {
	// JCP TODO: we could save an extra lookup by only using the ref here since GetHookByID requires two lookups
	hook, found := k.GetHookByID(ctx, ID)

	if !found {
		return
	}

	k.RemoveHook(ctx, hook)
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

func (k Keeper) GetAllHooksForKeyValue(ctx sdk.Context, triggerKey, triggerValue string) (list []types.Hook) {
	kvPrefix := types.HookKVPrefix(triggerKey, triggerValue)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), kvPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Hook
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) ExecuteHook(ctx sdk.Context, hook types.Hook) ([]byte, error) {
	contractAddr := sdk.MustAccAddressFromBech32(hook.ContractAddress)
	callerAddr := sdk.MustAccAddressFromBech32(hook.Creator)

	return k.wasmKeeper.Execute(ctx, contractAddr, callerAddr, []byte(hook.Args), sdk.Coins{})

	// JCP TODO: emit event with result
}
