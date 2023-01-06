package types

import (
	context "context"
	math "math"

	. "github.com/NicholasDotSol/duality/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper interface {
	TickHasToken0(sdk.Context, *Tick) bool
	TickHasToken1(sdk.Context, *Tick) bool
	SetTradingPair(sdk.Context, TradingPair)
	NewTickIterator(context.Context, int64, int64, string, bool, bool) TickIteratorI
}

// Assumes that the token0 liquidity is non-empty at this tick
func (p *TradingPair) InitLiquidityToken0(tickIndex int64) {
	minTick := &p.MinTick
	curTick1To0 := &p.CurrentTick1To0
	*minTick = MinInt64(*minTick, tickIndex)
	*curTick1To0 = MaxInt64(*curTick1To0, tickIndex)
}

// Assumes that the token1 liquidity is non-empty at this tick
func (p *TradingPair) InitLiquidityToken1(tickIndex int64) {
	maxTick := &p.MaxTick
	curTick0To1 := &p.CurrentTick0To1
	*maxTick = MaxInt64(*maxTick, tickIndex)
	*curTick0To1 = MinInt64(*curTick0To1, tickIndex)
}

func (p *TradingPair) InitLiquidity(tickIndex int64, addingToken0 bool) {
	if addingToken0 {
		p.InitLiquidityToken0(tickIndex)
	} else {
		p.InitLiquidityToken1(tickIndex)
	}
}

// Assumes that the token0 liquidity is empty at this tick
func (p *TradingPair) DeinitLiquidityToken0(ctx context.Context, k Keeper, tickIndex int64) {
	minTick := &p.MinTick
	cur1To0 := &p.CurrentTick1To0

	// Do nothing when liquidity is deinited between the current bounds.
	if *minTick < tickIndex && tickIndex < *cur1To0 {
		return
	}

	// We have removed all of Token0 from the pool
	if tickIndex == *minTick && tickIndex == *cur1To0 {
		*minTick = math.MaxInt64
		*cur1To0 = math.MinInt64
		// we leave cur1To0 where it is because otherwise we lose the last traded price
	} else if tickIndex == *minTick {

		nexMinTick := p.FindNewMinTick(ctx, k)
		*minTick = nexMinTick

		// we are removing liquidity below the current1To0, no need to update that
	} else if tickIndex == *cur1To0 {
		next1To0, found := p.FindNextTick1To0(ctx, k)
		if !found {
			// This case should be impossible if MinTick is tracked correctly
			*minTick = math.MaxInt64
			*cur1To0 = math.MinInt64
		} else {
			*cur1To0 = next1To0
		}
	}
}

// Assumes that the token1 liquidity is empty at this tick
func (p *TradingPair) DeinitLiquidityToken1(ctx context.Context, k Keeper, tickIndex int64) {
	maxTick := &p.MaxTick
	cur0To1 := &p.CurrentTick0To1

	// Do nothing when liquidity is deinited between the current bounds.
	if *cur0To1 < tickIndex && tickIndex < *maxTick {
		return
	}

	// We have removed all of Token0 from the pool
	if tickIndex == *cur0To1 && tickIndex == *maxTick {
		*maxTick = math.MinInt64
		*cur0To1 = math.MaxInt64
		// we leave cur1To0 where it is because otherwise we lose the last traded price
	} else if tickIndex == *maxTick {
		nextMaxTick := p.FindNewMaxTick(ctx, k)
		*maxTick = nextMaxTick

		// we are removing liquidity below the current1To0, no need to update that
	} else if tickIndex == *cur0To1 {
		next0To1, found := p.FindNextTick0To1(ctx, k)
		if !found {
			// This case should be impossible if MinTick is tracked correctly
			*maxTick = math.MinInt64
			*cur0To1 = math.MaxInt64
		} else {
			*cur0To1 = next0To1
		}
	}
}

func (p *TradingPair) DeinitLiquidity(ctx context.Context, k Keeper, tickIndex int64, removingToken0 bool) {
	if removingToken0 {
		p.DeinitLiquidityToken0(ctx, k, tickIndex)
	} else {
		p.DeinitLiquidityToken1(ctx, k, tickIndex)
	}
}

func (p *TradingPair) UpdateTickPointersPostAddToken0(goCtx context.Context, k Keeper, tick *Tick) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// TODO: Get rid of this, ideally should know exactly when to (de)init
	// 	and should not have to have the check before.
	if k.TickHasToken0(ctx, tick) {
		p.InitLiquidityToken0(tick.TickIndex)
	}

}

func (p *TradingPair) UpdateTickPointersPostAddToken1(goCtx context.Context, k Keeper, tick *Tick) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// TODO: Get rid of this, ideally should know exactly when to (de)init
	//	and should not have to have the check before.
	if k.TickHasToken1(ctx, tick) {
		p.InitLiquidityToken1(tick.TickIndex)
	}
}

func (p *TradingPair) UpdateTickPointersPostRemoveToken0(goCtx context.Context, k Keeper, tick *Tick) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// TODO: Get rid of this, ideally should know exactly when to (de)init
	//	and should not have to have the check before.
	if !k.TickHasToken0(ctx, tick) {
		p.DeinitLiquidityToken0(goCtx, k, tick.TickIndex)
	}
}

func (p *TradingPair) UpdateTickPointersPostRemoveToken1(goCtx context.Context, k Keeper, tick *Tick) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// TODO: Get rid of this, ideally should know exactly when to (de)init
	//	and should not have to have the check before.
	if !k.TickHasToken1(ctx, tick) {
		p.DeinitLiquidityToken1(goCtx, k, tick.TickIndex)
	}
}

func (p TradingPair) FindNextTick1To0(goCtx context.Context, k Keeper) (tickIdx int64, found bool) {

	// If MinTick == MaxInt64 it is unset
	// ie. There is no Token0 in the pool
	if p.MinTick == math.MaxInt64 {
		return math.MaxInt64, false
	}
	// Start scanning from CurrentTick1To0 - 1
	tickIdx = p.CurrentTick1To0 - 1

	ti := k.NewTickIterator(goCtx, tickIdx, p.MinTick, p.PairId, true, true)

	return ti.Next()
}

func (p TradingPair) FindNewMinTick(goCtx context.Context, k Keeper) (minTickIdx int64) {

	ti := k.NewTickIterator(goCtx, p.MinTick, p.CurrentTick1To0, p.PairId, true, false)
	idx, found := ti.Next()
	if found {
		return idx
	} else {
		return math.MaxInt64
	}
}

func (p TradingPair) FindNewMaxTick(goCtx context.Context, k Keeper) (maxTickIdx int64) {

	ti := k.NewTickIterator(goCtx, p.MaxTick, p.CurrentTick0To1, p.PairId, false, true)
	idx, found := ti.Next()
	if found {
		return idx
	} else {
		return math.MinInt64
	}
}

func (p TradingPair) FindNextTick0To1(goCtx context.Context, k Keeper) (tickIdx int64, found bool) {

	// If MaxTick == MinInt64 it is unset
	// There is no Token1 in the pool
	if p.MaxTick == math.MinInt64 {
		return math.MinInt64, false
	}

	// Start scanning from CurrentTick0To1 + 1
	tickIdx = p.CurrentTick0To1 + 1
	ti := k.NewTickIterator(goCtx, tickIdx, p.MaxTick, p.PairId, false, false)
	return ti.Next()
}
