package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type BranchableCache struct {
	Ctx   sdk.Context
	Write func()
}

func (bc BranchableCache) Branch() BranchableCache {
	cacheCtx, writeCache := bc.Ctx.CacheContext()
	newWriteFn := func() {
		writeCache()
		bc.Write()
	}

	return BranchableCache{Ctx: cacheCtx, Write: newWriteFn}
}

func NewBranchableCache(ctx sdk.Context) BranchableCache {
	// cacheCtx, write := ctx.CacheContext()
	return BranchableCache{Ctx: ctx, Write: func() {}}
}
