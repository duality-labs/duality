package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
)

type Liquidity interface {
	Swap(maxAmount sdk.Int) (inAmount sdk.Int, outAmount sdk.Int)
	Save(sdkCtx sdk.Context, keeper Keeper)
	Price() sdk.Dec
}

func NewLiquidityIterator(
	keeper Keeper,
	ctx sdk.Context,
	tradingPair types.DirectionalTradingPair,
) *LiquidityIterator {

	return &LiquidityIterator{
		iter:   keeper.NewTickIterator(ctx, tradingPair.PairId, tradingPair.TokenOut),
		keeper: keeper,
		ctx:    ctx,
		pairId: tradingPair.PairId,
		is0To1: tradingPair.IsTokenInToken0(),
	}

}

type LiquidityIterator struct {
	keeper Keeper
	pairId *types.PairId
	ctx    sdk.Context
	iter   TickIterator
	is0To1 bool
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
		switch tick.LiquidityType {
		case types.LiquidityTypeLP:
			var err error
			var pool Liquidity
			if s.is0To1 {
				pool, err = s.createPool0To1(tick)
			} else {
				pool, err = s.createPool1To0(tick)
			}
			// TODO: we are not actually handling the error here we're just stopping iteration
			// Should be a very rare edge case where the opposing tick is initialized above/below the Min/Max tick limit
			if err != nil {
				return nil
			}
			return pool

		case types.LiquidityTypeLO:
			return NewLimitOrderTranche(&tick)

		default:
			panic("Tick does not have liquidity")

		}
	}
	return nil
}

func (s *LiquidityIterator) createPool0To1(upperTick types.TickLiquidity) (Liquidity, error) {
	tickIndex := upperTick.TickIndex
	lowerTickIndex := tickIndex - 2*int64(upperTick.LiquidityIndex)
	lowerTick, err := s.keeper.GetOrInitTickLP(s.ctx, s.pairId, s.pairId.Token0, lowerTickIndex, upperTick.LiquidityIndex)
	if err != nil {
		return nil, err
	}
	pool := NewPool(
		tickIndex,
		lowerTick,
		&upperTick,
	)
	return NewLiquidityFromPool0To1(&pool), nil
}

func (s *LiquidityIterator) createPool1To0(lowerTick types.TickLiquidity) (Liquidity, error) {
	tickIndex := lowerTick.TickIndex
	upperTickIndex := tickIndex + 2*int64(lowerTick.LiquidityIndex)
	upperTick, err := s.keeper.GetOrInitTickLP(s.ctx, s.pairId, s.pairId.Token1, upperTickIndex, lowerTick.LiquidityIndex)
	if err != nil {
		return nil, err
	}

	pool := NewPool(
		lowerTick.TickIndex,
		&lowerTick,
		upperTick,
	)
	return NewLiquidityFromPool1To0(&pool), nil
}

func (s *LiquidityIterator) Close() {
	s.iter.Close()
}
