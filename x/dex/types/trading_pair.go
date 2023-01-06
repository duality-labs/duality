package types

import (
	context "context"
	math "math"

	. "github.com/NicholasDotSol/duality/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper interface {
	FindNextTick1To0(goCtx context.Context, tradingPair TradingPair) (tickIdx int64, found bool)
	FindNextTick0To1(goCtx context.Context, tradingPair TradingPair) (tickIdx int64, found bool)
	FindNewMinTick(goCtx context.Context, tradingPair TradingPair) (minTickIdx int64)
	FindNewMaxTick(goCtx context.Context, tradingPair TradingPair) (minTickIdx int64)
	TickHasToken0(ctx sdk.Context, tick *Tick) bool
	TickHasToken1(ctx sdk.Context, tick *Tick) bool
	SetTradingPair(ctx sdk.Context, TradingPair TradingPair)
}

func (p *TradingPair) InitLiquidity(tickIndex int64, addingToken0 bool) {
	if addingToken0 {
		p.InitLiquidityToken0(tickIndex)
	} else {
		p.InitLiquidityToken1(tickIndex)
	}
}

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

func (p *TradingPair) DeinitLiquidity(ctx context.Context, k Keeper, tickIndex int64, removingToken0 bool) {
	if removingToken0 {
		p.DeinitLiquidityToken0(ctx, k, tickIndex)
	} else {
		p.DeinitLiquidityToken1(ctx, k, tickIndex)
	}
}

// Assumes that the token0 liquidity is empty at this tick
func (p *TradingPair) DeinitLiquidityToken0(ctx context.Context, keeper Keeper, tickIndex int64) {
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

		nexMinTick := keeper.FindNewMinTick(ctx, *p)
		*minTick = nexMinTick

		// we are removing liquidity below the current1To0, no need to update that
	} else if tickIndex == *cur1To0 {
		next1To0, found := keeper.FindNextTick1To0(ctx, *p)
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
		nextMaxTick := k.FindNewMaxTick(ctx, *p)
		*maxTick = nextMaxTick

		// we are removing liquidity below the current1To0, no need to update that
	} else if tickIndex == *cur0To1 {
		next0To1, found := k.FindNextTick0To1(ctx, *p)
		if !found {
			// This case should be impossible if MinTick is tracked correctly
			*maxTick = math.MinInt64
			*cur0To1 = math.MaxInt64
		} else {
			*cur0To1 = next0To1
		}
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
