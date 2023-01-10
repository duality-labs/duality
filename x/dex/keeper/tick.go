package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewTick(pairId string, tickIndex int64, numFees uint64) types.Tick {
	tick := types.Tick{
		PairId:    pairId,
		TickIndex: tickIndex,
		TickData: &types.TickDataType{
			Reserve0: make([]sdk.Int, numFees),
			Reserve1: make([]sdk.Int, numFees),
		},
		LimitOrderTranche0To1: &types.LimitTrancheIndexes{0, 0},
		LimitOrderTranche1To0: &types.LimitTrancheIndexes{0, 0},
	}
	for i := 0; i < int(numFees); i++ {
		tick.TickData.Reserve0[i] = sdk.ZeroInt()
		tick.TickData.Reserve1[i] = sdk.ZeroInt()
	}
	return tick
}

// SetTick set a specific Tick in the store from its index
func (k Keeper) SetTick(ctx sdk.Context, pairId string, Tick types.Tick) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BaseTickKeyPrefix))
	b := k.cdc.MustMarshal(&Tick)
	store.Set(types.TickKey(
		pairId,
		Tick.TickIndex,
	), b)
}

// GetTick returns a Tick from its index
func (k Keeper) GetTick(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,
) (val types.Tick, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BaseTickKeyPrefix))

	b := store.Get(types.TickKey(
		pairId,
		tickIndex,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTick removes a Tick from the store
func (k Keeper) RemoveTick(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BaseTickKeyPrefix))
	store.Delete(types.TickKey(
		pairId,
		tickIndex,
	))
}

// GetAllTick returns all Tick
func (k Keeper) GetAllTick(ctx sdk.Context) (list []types.Tick) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BaseTickKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Tick
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAllTick returns all Tick
func (k Keeper) GetAllTickByPair(ctx sdk.Context, pairId string) (list []types.Tick) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.TickPrefix(pairId))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Tick
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
