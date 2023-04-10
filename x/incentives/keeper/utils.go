package keeper

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// findIndex takes an array of IDs. Then return the index of a specific ID.
func findIndex(IDs []uint64, ID uint64) int {
	for index, id := range IDs {
		if id == ID {
			return index
		}
	}
	return -1
}

// removeValue takes an array of IDs. Then finds the index of the IDs and remove those IDs from the array.
func removeValue(IDs []uint64, ID uint64) ([]uint64, int) {
	index := findIndex(IDs, ID)
	if index < 0 {
		return IDs, index
	}
	IDs[index] = IDs[len(IDs)-1] // set last element to index
	return IDs[:len(IDs)-1], index
}

// getRefs returns the IDs specified by the provided key.
func (k Keeper) getRefs(ctx sdk.Context, key []byte) []uint64 {
	store := ctx.KVStore(k.storeKey)
	ids := []uint64{}
	if store.Has(key) {
		bz := store.Get(key)
		err := json.Unmarshal(bz, &ids)
		if err != nil {
			panic(err)
		}
	}
	return ids
}

// addRefByKey appends the provided object ID into an array associated with the provided key.
func (k Keeper) addRefByKey(ctx sdk.Context, key []byte, id uint64) error {
	store := ctx.KVStore(k.storeKey)
	ids := k.getRefs(ctx, key)
	if findIndex(ids, id) > -1 {
		return fmt.Errorf("object with same ID exists: %d", id)
	}
	ids = append(ids, id)
	bz, err := json.Marshal(ids)
	if err != nil {
		return err
	}
	store.Set(key, bz)
	return nil
}

// deleteRefByKey removes the provided object ID from an array associated with the provided key.
func (k Keeper) deleteRefByKey(ctx sdk.Context, key []byte, id uint64) error {
	store := ctx.KVStore(k.storeKey)
	ids := k.getRefs(ctx, key)
	ids, index := removeValue(ids, id)
	if index < 0 {
		return fmt.Errorf("specific object with ID %d not found by reference %s", id, key)
	}
	if len(ids) == 0 {
		store.Delete(key)
	} else {
		bz, err := json.Marshal(ids)
		if err != nil {
			return err
		}
		store.Set(key, bz)
	}
	return nil
}
