package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type BranchableCache struct {
	Ctx   sdk.Context
	Write func()
}

func (bc BranchableCache) Branch() BranchableCache {
	cacheCtx, writeCache := bc.Ctx.CacheContext()
	newWriteFn := func() {
		// To write a branch back the root KVstore we have to recursively call
		// the write fn for all parent branches
		writeCache()
		// // NOTE: in current version events from the cache are never passed to the
		// // parent context. This is fixed in cosmos v0.46.4
		// // Once we update, the below code can be removed
		// bc.Ctx.EventManager().EmitEvents(cacheCtx.EventManager().Events())
		// bc.Write()
	}

	return BranchableCache{Ctx: cacheCtx, Write: newWriteFn}
}

func NewBranchableCache(ctx sdk.Context) BranchableCache {
	return BranchableCache{Ctx: ctx, Write: func() {}}
}
