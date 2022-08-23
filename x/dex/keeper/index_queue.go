package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) enqueue(ctx sdk.Context, queue []*types.IndexQueueType, newQueueItem types.IndexQueueType) []*types.IndexQueueType {

	queue = append(queue, &newQueueItem)
	return queue
}

func (k Keeper) dequeue(ctx sdk.Context, queue []*types.IndexQueueType) (types.IndexQueueType, []*types.IndexQueueType) {
	if len(queue) == 0 {
		return types.IndexQueueType{sdk.ZeroDec(), sdk.ZeroDec(), &types.OrderParams{"", "", sdk.ZeroDec()}}, nil
	}
	element := queue[0]
	queue = queue[1:]
	return *element, queue
}

// SetIndexQueue set a specific IndexQueue in the store from its index
func (k Keeper) SetIndexQueue(ctx sdk.Context, token0 string, token1 string, IndexQueue types.IndexQueue) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IndexQueueKeyPrefix))
	b := k.cdc.MustMarshal(&IndexQueue)
	//fmt.Println(token0)
	//fmt.Println(token1)
	//fmt.Println("set index", IndexQueue.Index)
	store.Set(types.IndexQueueKey(token0, token1,
		IndexQueue.Index,
	), b)
}

// GetIndexQueue returns a IndexQueue from its index
func (k Keeper) GetIndexQueue(
	ctx sdk.Context,
	token0 string,
	token1 string,
	index int32,

) (val types.IndexQueue, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IndexQueueKeyPrefix))
	//fmt.Println(token0)
	//fmt.Println(token1)
	//fmt.Println("Get Index:", index)
	b := store.Get(types.IndexQueueKey(token0, token1,
		index,
	))
	//fmt.Println(b)
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveIndexQueue removes a IndexQueue from the store
func (k Keeper) RemoveIndexQueue(
	ctx sdk.Context,
	token0 string,
	token1 string,
	index int32,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IndexQueueKeyPrefix))
	store.Delete(types.IndexQueueKey(token0, token1,
		index,
	))
}

// GetAllIndexQueue returns all IndexQueue
func (k Keeper) GetAllIndexQueue(ctx sdk.Context) (list []types.IndexQueue) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IndexQueueKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.IndexQueue
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
