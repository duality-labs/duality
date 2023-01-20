package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
)

type Liquidity interface {
	Swap(maxAmount sdk.Int) (inAmount sdk.Int, outAmount sdk.Int, initedTick *types.Tick, deinitedTick *types.Tick)
	Save(ctx context.Context, keeper Keeper)
	Price() sdk.Dec
}

type LiquidityIterator interface {
	HasNext() bool
	Next() Liquidity
}

func NewLiquidityIterator(
	keeper Keeper,
	ctx context.Context,
	tradingPair types.DirectionalTradingPair,
	feeTiers []types.FeeTier,
) LiquidityIterator {

	if tradingPair.IsTokenInToken0() {
		return NewLiquidityIterator0To1(keeper, ctx, tradingPair.TradingPair, feeTiers)
	} else {
		return NewLiquidityIterator1To0(keeper, ctx, tradingPair.TradingPair, feeTiers)
	}
}

type LiquidityIterator0To1 struct {
	curTickIndex int64
	curFeeIndex  uint64
	maxTick      int64
	keeper       Keeper
	tradingPair  types.TradingPair
	ctx          context.Context
	nextLiq      Liquidity
	feeTiers     []types.FeeTier
}

func (s *LiquidityIterator0To1) HasNext() bool {
	return s.nextLiq != nil
}

func (s *LiquidityIterator0To1) Next() Liquidity {
	if s.nextLiq == nil {
		panic("should not call Next() if hasNext() returns false")
	}
	liq := s.nextLiq
	s.nextLiq = s.getNext()
	return liq
}

func (s *LiquidityIterator0To1) getNext() Liquidity {
	iter := s.keeper.NewTickIterator(s.ctx, s.curTickIndex, s.maxTick, s.tradingPair.PairId, false)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		upperTick := iter.Value()
		for int(s.curFeeIndex) < len(upperTick.TickData.Reserve1) {
			if upperTick.TickData.Reserve1[s.curFeeIndex].Equal(sdk.ZeroInt()) {
				s.curFeeIndex++
				continue
			}
			fee := s.feeTiers[s.curFeeIndex].Fee
			lowerTick, err := s.keeper.GetOrInitTick(s.ctx, s.tradingPair.PairId, s.curTickIndex-2*fee)
			if err != nil {
				return nil
			}

			pool := NewPool(
				&s.tradingPair,
				s.curTickIndex,
				s.curFeeIndex,
				&lowerTick,
				&upperTick,
			)
			s.curFeeIndex++
			return NewLiquidityFromPool0To1(&pool)
		}

		s.curFeeIndex = 0
		s.curTickIndex = upperTick.TickIndex + 1

		orderBook := s.keeper.GetLimitOrderBook1To0(
			s.ctx,
			&s.tradingPair,
			&upperTick,
		)

		if orderBook.HasLiquidity() {
			return orderBook
		}
	}
	return nil
}

func NewLiquidityIterator0To1(
	keeper Keeper,
	ctx context.Context,
	tradingPair types.TradingPair,
	feeTiers []types.FeeTier,
) *LiquidityIterator0To1 {
	iter := &LiquidityIterator0To1{
		curTickIndex: tradingPair.CurrentTick0To1,
		curFeeIndex:  0,
		keeper:       keeper,
		ctx:          ctx,
		tradingPair:  tradingPair,
		maxTick:      tradingPair.MaxTick,
		nextLiq:      nil,
		feeTiers:     feeTiers,
	}
	iter.nextLiq = iter.getNext()
	return iter
}

type LiquidityIterator1To0 struct {
	curTickIndex int64
	curFeeIndex  uint64
	minTick      int64
	keeper       Keeper
	tradingPair  types.TradingPair
	ctx          context.Context
	nextLiq      Liquidity
	feeTiers     []types.FeeTier
}

func (s *LiquidityIterator1To0) HasNext() bool {
	return s.nextLiq != nil
}

func (s *LiquidityIterator1To0) Next() Liquidity {
	if s.nextLiq == nil {
		panic("should not call next if hasNext() returns false")
	}
	liq := s.nextLiq
	s.nextLiq = s.getNext()
	return liq
}

func (s *LiquidityIterator1To0) getNext() Liquidity {
	iter := s.keeper.NewTickIterator(s.ctx, s.curTickIndex, s.minTick, s.tradingPair.PairId, true)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		lowerTick := iter.Value()

		for int(s.curFeeIndex) < len(lowerTick.TickData.Reserve0) {
			if lowerTick.TickData.Reserve0[s.curFeeIndex].Equal(sdk.ZeroInt()) {
				s.curFeeIndex++
				continue
			}
			fee := s.feeTiers[s.curFeeIndex].Fee
			upperTick, err := s.keeper.GetOrInitTick(s.ctx, s.tradingPair.PairId, s.curTickIndex+2*fee)
			if err != nil {
				return nil
			}

			pool := NewPool(
				&s.tradingPair,
				s.curTickIndex,
				s.curFeeIndex,
				&lowerTick,
				&upperTick,
			)
			s.curFeeIndex++
			return NewLiquidityFromPool1To0(&pool)
		}

		s.curFeeIndex = 0
		s.curTickIndex = lowerTick.TickIndex - 1

		orderBook := s.keeper.GetLimitOrderBook0To1(
			s.ctx,
			&s.tradingPair,
			&lowerTick,
		)

		if orderBook.HasLiquidity() {
			return orderBook
		}
	}

	return nil
}

func NewLiquidityIterator1To0(keeper Keeper, ctx context.Context, tradingPair types.TradingPair, feeTiers []types.FeeTier) *LiquidityIterator1To0 {
	iter := &LiquidityIterator1To0{
		curTickIndex: tradingPair.CurrentTick1To0,
		curFeeIndex:  0,
		keeper:       keeper,
		ctx:          ctx,
		tradingPair:  tradingPair,
		minTick:      tradingPair.MinTick,
		feeTiers:     feeTiers,
		nextLiq:      nil,
	}
	iter.nextLiq = iter.getNext()
	return iter
}
