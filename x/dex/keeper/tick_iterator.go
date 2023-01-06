package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type TickIterator struct {
	start    int64
	end      int64
	pairId   string
	keeper   Keeper
	ctx      sdk.Context
	hasToken func(sdk.Context, *types.Tick) bool
	nextTick func(tick int64) int64
	stop     func(int64) bool
}

func (k Keeper) NewTickIterator(ctx context.Context,
	start int64,
	end int64,
	pairId string,
	findToken0 bool,
	scanLeft bool) types.TickIteratorI {
	var hasToken func(sdk.Context, *types.Tick) bool
	var nextTick func(tick int64) int64
	var stop func(int64) bool

	if findToken0 {
		hasToken = k.TickHasToken0
	} else {
		hasToken = k.TickHasToken1
	}

	if scanLeft {
		nextTick = func(tick int64) int64 { tick--; return tick }
		stop = func(idx int64) bool { return idx < end }
	} else {
		nextTick = func(tick int64) int64 { tick++; return tick }
		stop = func(idx int64) bool { return idx > end }
	}

	return TickIterator{
		start:    start,
		end:      end,
		pairId:   pairId,
		keeper:   k,
		hasToken: hasToken,
		nextTick: nextTick,
		stop:     stop,
		ctx:      sdk.UnwrapSDKContext(ctx),
	}
}

func (ti TickIterator) Next() (idx int64, found bool) {

	curIdx := ti.start

	for !ti.stop(curIdx) {
		// Checks for the next value tick containing liquidity
		tick, tickFound := ti.keeper.GetTick(ti.ctx, ti.pairId, curIdx)

		if tickFound && ti.hasToken(ti.ctx, &tick) {
			return curIdx, true
		}

		curIdx = ti.nextTick(curIdx)
	}
	return 0, false
}
