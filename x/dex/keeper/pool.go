package keeper

import (
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

	pool, err := types.NewPool(pairID, centerTickIndexNormalized, fee)
	if err != nil {
		return nil, err
	}

	// This is important because it sets the pool's id (pool.Metadata.Id)
	k.SetPool(ctx, pool)

	return pool, err
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
		Metadata: &types.PoolMetadata{
			Id:     upperTick.PoolId,
			PairId: upperTick.Key.TradePairID.MustPairID(),
			NormalizedCenterTickIndex: upperTick.Key.TickIndexTakerToMaker - int64(
				upperTick.Key.Fee,
			),
			Fee: upperTick.Key.Fee,
		},
		LowerTick0: lowerTick,
		UpperTick1: upperTick,
	}, true
}

func (k Keeper) SetPool(ctx sdk.Context, pool *types.Pool) {
	poolID := &pool.Metadata.Id
	if *poolID == 0 {
		*poolID = k.getNextPoolIdAndIncrement(ctx)
		pool.LowerTick0.PoolId = *poolID
		pool.UpperTick1.PoolId = *poolID
		k.SetPoolMetadata(ctx, pool.Metadata)
	}

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
func MustNewPool(
	poolID uint64,
	pairID *types.PairID,
	normalizedCenterTickIndex int64,
	fee uint64,
) *types.Pool {
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
		Metadata: &types.PoolMetadata{
			Id:                        poolID,
			PairId:                    pairID,
			NormalizedCenterTickIndex: normalizedCenterTickIndex,
			Fee:                       fee,
		},
		LowerTick0: lowerTick,
		UpperTick1: upperTick,
	}
}
