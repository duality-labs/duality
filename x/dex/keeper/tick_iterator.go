package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) NewTickIterator(ctx context.Context,
	start int64,
	end int64,
	pairId *types.PairId,
	scanLeft bool) types.TickIteratorI {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	prefixStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.TickPrefix(pairId))
	startKey := types.TickIndexToBytes(startIndex)

	if scanLeft {
		return prefixStore.Iterator(startKey, nil)
	} else {
		return prefixStore.ReverseIterator(nil, startKey)
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
