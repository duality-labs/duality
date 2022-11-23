package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewTick(pairId string, tickIndex int64, numFees uint64) types.TickMap {
	tick := types.TickMap{
		PairId:    pairId,
		TickIndex: tickIndex,
		TickData: &types.TickDataType{
			Reserve0AndShares: make([]*types.Reserve0AndSharesType, numFees),
			Reserve1:          make([]sdk.Dec, numFees),
		},
		LimitOrderTranche0To1: &types.LimitOrderTrancheTrancheIndexes{0, 0},
		LimitOrderTranche1To0: &types.LimitOrderTrancheTrancheIndexes{0, 0},
	}
	for i := 0; i < int(numFees); i++ {
		tick.TickData.Reserve0AndShares[i] = &types.Reserve0AndSharesType{sdk.ZeroDec(), sdk.ZeroDec()}
		tick.TickData.Reserve1[i] = sdk.ZeroDec()
	}
	return tick
}

// SetTickMap set a specific tickMap in the store from its index
func (k Keeper) SetTickMap(ctx sdk.Context, pairId string, tickMap types.TickMap) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BaseTickMapKeyPrefix))
	b := k.cdc.MustMarshal(&tickMap)
	store.Set(types.TickMapKey(
		pairId,
		tickMap.TickIndex,
	), b)
}

// GetTickMap returns a tickMap from its index
func (k Keeper) GetTickMap(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,
) (val types.TickMap, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BaseTickMapKeyPrefix))

	b := store.Get(types.TickMapKey(
		pairId,
		tickIndex,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTickMap removes a tickMap from the store
func (k Keeper) RemoveTickMap(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BaseTickMapKeyPrefix))
	store.Delete(types.TickMapKey(
		pairId,
		tickIndex,
	))
}

// GetAllTickMap returns all tickMap
func (k Keeper) GetAllTickMap(ctx sdk.Context) (list []types.TickMap) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BaseTickMapKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TickMap
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAllTickMap returns all tickMap
func (k Keeper) GetAllTickMapByPair(ctx sdk.Context, pairId string) (list []types.TickMap) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.TickPrefix(pairId))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TickMap
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
