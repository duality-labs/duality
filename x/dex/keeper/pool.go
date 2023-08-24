package keeper

import (
	"encoding/binary"
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/duality-labs/duality/x/dex/utils"
)

func (k Keeper) GetOrInitPool(
	ctx sdk.Context,
	pairID *types.PairID,
	centerTickIndexNormalized int64,
	fee uint64,
) (*types.Pool, error) {
	pool, found := k.GetPool(ctx, pairID, centerTickIndexNormalized, fee)
	if found {
		return pool, nil
	}

	return k.InitPool(ctx, pairID, centerTickIndexNormalized, fee)
}

func (k Keeper) InitPool(
	ctx sdk.Context,
	pairID *types.PairID,
	centerTickIndexNormalized int64,
	fee uint64,
) (*types.Pool, error) {
	poolID := k.InitPoolKeys(ctx, pairID, centerTickIndexNormalized, fee)
	return types.NewPool(pairID, centerTickIndexNormalized, fee, poolID)
}

func (k Keeper) InitPoolKeys(ctx sdk.Context,
	pairID *types.PairID,
	centerTickIndexNormalized int64,
	fee uint64,
) string {
	poolKeyBz := types.PoolKey(*pairID, centerTickIndexNormalized, fee)

	poolID := k.GetNextPoolID(ctx)
	poolIDBz := []byte(poolID)

	store := ctx.KVStore(k.storeKey)

	poolIDStore := prefix.NewStore(store, types.KeyPrefix(types.PoolIDKeyPrefix))
	poolIDStore.Set(poolKeyBz, poolIDBz)

	poolRefStore := prefix.NewStore(store, types.KeyPrefix(types.PoolRefKeyPrefix))
	poolRefStore.Set(poolIDBz, poolKeyBz)

	return poolID
}

// GetNextPoolId get ID for the next pool to be created
func (k Keeper) GetNextPoolID(ctx sdk.Context) string {
	poolCount := k.GetPoolCount(ctx)
	// JCP TODO: this should happen in init keys once we switch poolID to uint64
	k.SetPoolCount(ctx, poolCount+1)
	return types.NewPoolDenom(poolCount)
}

func (k Keeper) GetPool(
	ctx sdk.Context,
	pairID *types.PairID,
	centerTickIndexNormalized int64,
	fee uint64,
) (*types.Pool, bool) {
	feeInt64 := utils.MustSafeUint64(fee)

	id0To1 := &types.PoolReservesKey{
		TradePairID:           types.NewTradePairIDFromMaker(pairID, pairID.Token1),
		TickIndexTakerToMaker: centerTickIndexNormalized + feeInt64,
		Fee:                   fee,
	}

	poolID, found := k.GetPoolIDByParams(ctx, pairID, centerTickIndexNormalized, fee)
	if !found {
		return nil, false
	}

	upperTick, upperTickFound := k.GetPoolReserves(ctx, id0To1)
	lowerTick, lowerTickFound := k.GetPoolReserves(ctx, id0To1.Counterpart())

	if !lowerTickFound && upperTickFound {
		lowerTick = types.NewPoolReservesFromCounterpart(upperTick)
	} else if lowerTickFound && !upperTickFound {
		upperTick = types.NewPoolReservesFromCounterpart(lowerTick)
	} else if !lowerTickFound && !upperTickFound {
		return nil, false
	}

	return &types.Pool{
		ID:         poolID,
		LowerTick0: lowerTick,
		UpperTick1: upperTick,
	}, true
}

// GetPoolCount get the total number of pool
func (k Keeper) GetPoolCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PoolCountKeyPrefix)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) GetPoolIDByParams(
	ctx sdk.Context,
	pairID *types.PairID,
	centerTickIndexNormalized int64,
	fee uint64,
) (id string, found bool) {
	poolRefKey := types.PoolKey(*pairID, centerTickIndexNormalized, fee)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolIDKeyPrefix))
	b := store.Get(poolRefKey)
	if b == nil {
		return "", false
	}

	return string(b), true
}

func (k Keeper) GetPoolParamsByID(
	ctx sdk.Context,
	id string,
) (types.PoolParams, error) {
	// JCP TODO: should this be found vs error?
	ref, found := k.GetPoolRefByID(ctx, id)
	if !found {
		return types.PoolParams{}, types.ErrInvalidDepositDenom
	}
	fmt.Printf("REf: %v (%v)\n", ref, string(ref))

	poolParams, err := types.ParsePoolRefToParams(ref)
	if err != nil {
		panic("Cannot parse pool ref")
	}

	return poolParams, nil
}

func (k Keeper) GetPoolRefByID(
	ctx sdk.Context,
	id string,
) (ref []byte, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolRefKeyPrefix))
	b := store.Get([]byte(id))
	if b == nil {
		return []byte{}, false
	}

	return b, true
}

// SetPoolCount set the total number of pool
func (k Keeper) SetPoolCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PoolCountKeyPrefix)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) SetPool(ctx sdk.Context, pool *types.Pool) {
	if pool.LowerTick0.HasToken() {
		k.SetPoolReserves(ctx, pool.LowerTick0)
	} else {
		k.RemovePoolReserves(ctx, pool.LowerTick0.Key)
	}
	if pool.UpperTick1.HasToken() {
		k.SetPoolReserves(ctx, pool.UpperTick1)
	} else {
		k.RemovePoolReserves(ctx, pool.UpperTick1.Key)
	}

	// TODO: this will create a bit of extra noise since not every Save is updating both ticks
	// This should be solved upstream by better tracking of dirty ticks
	ctx.EventManager().EmitEvent(types.CreateTickUpdatePoolReserves(*pool.LowerTick0))
	ctx.EventManager().EmitEvent(types.CreateTickUpdatePoolReserves(*pool.UpperTick1))
}

// Useful for testing
func MustNewPool(pairID *types.PairID, normalizedCenterTickIndex int64, fee uint64) *types.Pool {
	feeInt64 := utils.MustSafeUint64(fee)

	id0To1 := &types.PoolReservesKey{
		TradePairID:           types.NewTradePairIDFromMaker(pairID, pairID.Token1),
		TickIndexTakerToMaker: normalizedCenterTickIndex + feeInt64,
		Fee:                   fee,
	}

	upperTick, err := types.NewPoolReserves(id0To1)
	if err != nil {
		panic(err)
	}

	lowerTick := types.NewPoolReservesFromCounterpart(upperTick)

	return &types.Pool{
		LowerTick0: lowerTick,
		UpperTick1: upperTick,
	}
}
