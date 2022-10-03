package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetLimitOrderPoolUserSharesFilled set a specific limitOrderPoolUserSharesFilled in the store from its index
func (k Keeper) SetLimitOrderPoolUserSharesFilled(ctx sdk.Context, pairId string, tickIndex int64, token string, limitOrderPoolUserSharesFilled types.LimitOrderPoolUserSharesFilled) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LimitOrderUserSharesFilledPrefix(pairId, tickIndex, token))
	b := k.cdc.MustMarshal(&limitOrderPoolUserSharesFilled)
	store.Set(types.LimitOrderPoolUserSharesFilledKey(
		limitOrderPoolUserSharesFilled.Count,
		limitOrderPoolUserSharesFilled.Address,
	), b)
}

// GetLimitOrderPoolUserSharesFilled returns a limitOrderPoolUserSharesFilled from its index
func (k Keeper) GetLimitOrderPoolUserSharesFilled(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,
	token string,
	count uint64,
	address string,

) (val types.LimitOrderPoolUserSharesFilled, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LimitOrderUserSharesFilledPrefix(pairId, tickIndex, token))

	b := store.Get(types.LimitOrderPoolUserSharesFilledKey(
		count,
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLimitOrderPoolUserSharesFilled removes a limitOrderPoolUserSharesFilled from the store
func (k Keeper) RemoveLimitOrderPoolUserSharesFilled(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,
	token string,
	count uint64,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LimitOrderUserSharesFilledPrefix(pairId, tickIndex, token))
	store.Delete(types.LimitOrderPoolUserSharesFilledKey(
		count,
		address,
	))
}
