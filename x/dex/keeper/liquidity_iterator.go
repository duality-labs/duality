package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/duality-labs/duality/x/dex/utils"
)

type LiquidityIterator struct {
	keeper      *Keeper
	tradePairID *types.TradePairID
	ctx         sdk.Context
	iter        TickIterator
}

func (k Keeper) NewLiquidityIterator(
	ctx sdk.Context,
	tradePairID *types.TradePairID,
) *LiquidityIterator {
	return &LiquidityIterator{
		iter:        k.NewTickIterator(ctx, tradePairID),
		keeper:      &k,
		ctx:         ctx,
		tradePairID: tradePairID,
	}
}

func (s *LiquidityIterator) Next() Liquidity {
	// Move iterator to the next tick after each call
	// iter must be in valid state to call next
	defer func() {
		if s.iter.Valid() {
			s.iter.Next()
		}
	}()

	for ; s.iter.Valid(); s.iter.Next() {
		tick := s.iter.Value()
		// Don't bother to look up pool counter-liquidities if there's no liquidity here
		if !tick.HasToken() {
			continue
		}

		liq := s.WrapTickLiquidity(tick)
		if liq != nil {
			return liq
		}
	}

	return nil
}

func (s *LiquidityIterator) Close() {
	s.iter.Close()
}

func (s *LiquidityIterator) NewLiquidityFromUpperPoolReserves(upperTick types.PoolReserves) (Liquidity, error) {
	upperTickIndex := upperTick.TickIndex
	centerTickIndex := upperTickIndex - utils.MustSafeUint64(upperTick.Fee)
	lowerTickIndex := centerTickIndex - utils.MustSafeUint64(upperTick.Fee)
	lowerTick, found := s.keeper.GetPoolReserves(
		s.ctx,
		s.tradePairID.Reversed(),
		lowerTickIndex,
		upperTick.Fee,
	)
	if !found {
		return nil, errors.Wrapf(
			types.ErrCorruptPoolLiquidity, "pairID: %s, makerDenom: %s, tickIndex: %d, fee: %d",
			s.tradePairID.MustPairID().CanonicalString(),
			s.tradePairID.MakerDenom,
			lowerTickIndex,
			upperTick.Fee,
		)
	}

	pool := NewPool(
		centerTickIndex,
		lowerTick,
		&upperTick,
	)

	return NewPoolLiquidity(s.tradePairID, &pool), nil
}

func (s *LiquidityIterator) NewLiquidityFromLowerPoolReserves(lowerTick types.PoolReserves) (Liquidity, error) {
	lowerTickIndex := lowerTick.TickIndex
	centerTickIndex := lowerTickIndex + utils.MustSafeUint64(lowerTick.Fee)
	upperTickIndex := centerTickIndex + utils.MustSafeUint64(lowerTick.Fee)
	upperTick, found := s.keeper.GetPoolReserves(
		s.ctx,
		s.tradePairID.Reversed(),
		upperTickIndex,
		lowerTick.Fee,
	)
	if !found {
		return nil, errors.Wrapf(
			types.ErrCorruptPoolLiquidity,
			"pairID: %s, makerDenom: %s, tickIndex: %d, fee: %d",
			s.tradePairID.MustPairID().CanonicalString(),
			s.tradePairID.MakerDenom,
			upperTickIndex,
			lowerTick.Fee,
		)
	}

	pool := NewPool(
		centerTickIndex,
		&lowerTick,
		upperTick,
	)

	return NewPoolLiquidity(s.tradePairID, &pool), nil
}

func (s *LiquidityIterator) WrapTickLiquidity(tick types.TickLiquidity) Liquidity {
	switch liquidity := tick.Liquidity.(type) {
	case *types.TickLiquidity_PoolReserves:
		var err error
		var pool Liquidity
		poolReserves := *liquidity.PoolReserves
		if s.tradePairID.IsTakerDenomToken0() {
			// Pool Reserves is upperTick
			pool, err = s.NewLiquidityFromUpperPoolReserves(poolReserves)
		} else {
			// Pool Reserves is lowerTick
			pool, err = s.NewLiquidityFromLowerPoolReserves(poolReserves)
		}
		// TODO: we are not actually handling the error here we're just stopping iteration
		// Should be a very rare edge case where the opposing tick is initialized
		// above/below the Min/Max tick limit
		if err != nil {
			// TODO: Remove
			panic(err)
		}

		return pool

	case *types.TickLiquidity_LimitOrderTranche:
		tranche := liquidity.LimitOrderTranche
		// If we hit a tranche with an expired goodTil date keep iterating
		if tranche.IsExpired(s.ctx) {
			return nil
		}

		return tranche

	default:
		panic("Tick does not have liquidity")
	}
}
