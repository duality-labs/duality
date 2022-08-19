package keeper

import (
	"encoding/binary"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) enqueue(ctx sdk.Context, queue []*types.VirtualPriceQueueType, newQueueItem types.VirtualPriceQueueType) []*types.VirtualPriceQueueType {

	queue = append(queue, &newQueueItem)
	return queue
}

func (k Keeper) dequeue(ctx sdk.Context, queue []*types.VirtualPriceQueueType) (types.VirtualPriceQueueType, []*types.VirtualPriceQueueType) {
	if len(queue) == 0 {
		return types.VirtualPriceQueueType{sdk.ZeroDec(), sdk.ZeroDec(), &types.OrderParams{"", "", sdk.ZeroDec()}}, nil
	}
	element := queue[0]
	queue = queue[1:]
	return *element, queue
}

// GetVirtualPriceTickQueueCount get the total number of virtualPriceTickQueue
func (k Keeper) GetVirtualPriceTickQueueCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.VirtualPriceTickQueueCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetVirtualPriceTickQueueCount set the total number of virtualPriceTickQueue
func (k Keeper) SetVirtualPriceTickQueueCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.VirtualPriceTickQueueCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendVirtualPriceTickQueue appends a virtualPriceTickQueue in the store with a new id and update the count
func (k Keeper) AppendVirtualPriceTickQueue(
	ctx sdk.Context,
	virtualPriceTickQueue types.VirtualPriceTickQueue,
) uint64 {
	// Create the virtualPriceTickQueue
	count := k.GetVirtualPriceTickQueueCount(ctx)

	// Set the ID of the appended value
	virtualPriceTickQueue.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VirtualPriceTickQueueKey))
	appendedValue := k.cdc.MustMarshal(&virtualPriceTickQueue)
	store.Set(GetVirtualPriceTickQueueIDBytes(virtualPriceTickQueue.Id), appendedValue)

	// Update virtualPriceTickQueue count
	k.SetVirtualPriceTickQueueCount(ctx, count+1)

	return count
}

// SetVirtualPriceTickQueue set a specific virtualPriceTickQueue in the store
func (k Keeper) SetVirtualPriceTickQueue(ctx sdk.Context, virtualPriceTickQueue types.VirtualPriceTickQueue) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VirtualPriceTickQueueKey))
	b := k.cdc.MustMarshal(&virtualPriceTickQueue)
	store.Set(GetVirtualPriceTickQueueIDBytes(virtualPriceTickQueue.Id), b)
}

// GetVirtualPriceTickQueue returns a virtualPriceTickQueue from its id
func (k Keeper) GetVirtualPriceTickQueue(ctx sdk.Context, id uint64) (val types.VirtualPriceTickQueue, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VirtualPriceTickQueueKey))
	b := store.Get(GetVirtualPriceTickQueueIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveVirtualPriceTickQueue removes a virtualPriceTickQueue from the store
func (k Keeper) RemoveVirtualPriceTickQueue(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VirtualPriceTickQueueKey))
	store.Delete(GetVirtualPriceTickQueueIDBytes(id))
}

// GetAllVirtualPriceTickQueue returns all virtualPriceTickQueue
func (k Keeper) GetAllVirtualPriceTickQueue(ctx sdk.Context) (list []types.VirtualPriceTickQueue) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VirtualPriceTickQueueKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.VirtualPriceTickQueue
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetVirtualPriceTickQueueIDBytes returns the byte representation of the ID
func GetVirtualPriceTickQueueIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetVirtualPriceTickQueueIDFromBytes returns ID in uint64 format from a byte array
func GetVirtualPriceTickQueueIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
