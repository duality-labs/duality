package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetOrInitLimitOrderTranche(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,
	tokenIn string,
	trancheIndex uint64,
) types.LimitOrderTranche {
	tranche, found := k.GetLimitOrderTranche(ctx, pairId, tickIndex, tokenIn, trancheIndex)
	if !found {
		tranche = types.LimitOrderTranche{
			TrancheIndex:     trancheIndex,
			TickIndex:        tickIndex,
			TokenIn:          tokenIn,
			PairId:           pairId,
			ReservesTokenIn:  sdk.ZeroInt(),
			ReservesTokenOut: sdk.ZeroInt(),
			TotalTokenIn:     sdk.ZeroInt(),
			TotalTokenOut:    sdk.ZeroInt(),
		}
		k.SetLimitOrderTranche(ctx, tranche)
	}

	return tranche
}

// SetLimitOrderTranche set a specific LimitOrderTranche in the store from its index
func (k Keeper) SetLimitOrderTranche(ctx sdk.Context, LimitOrderTranche types.LimitOrderTranche) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderTrancheKeyPrefix))
	b := k.cdc.MustMarshal(&LimitOrderTranche)
	store.Set(types.LimitOrderTrancheKey(
		LimitOrderTranche.PairId,
		LimitOrderTranche.TickIndex,
		LimitOrderTranche.TokenIn,
		LimitOrderTranche.TrancheIndex,
	), b)
}

// GetLimitOrderTranche returns a LimitOrderTranche from its index
func (k Keeper) GetLimitOrderTranche(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,
	token string,
	tranchIndex uint64,

) (val types.LimitOrderTranche, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderTrancheKeyPrefix))

	b := store.Get(types.LimitOrderTrancheKey(
		pairId,
		tickIndex,
		token,
		tranchIndex,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLimitOrderTranche removes a LimitOrderTranche from the store
func (k Keeper) RemoveLimitOrderTranche(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,
	token string,
	trancheIndex uint64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderTrancheKeyPrefix))
	store.Delete(types.LimitOrderTrancheKey(
		pairId,
		tickIndex,
		token,
		trancheIndex,
	))
}

// GetAllLimitOrderTranche returns all LimitOrderTrancheUser
func (k Keeper) GetAllLimitOrderTranche(ctx sdk.Context) (list []types.LimitOrderTranche) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderTrancheKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LimitOrderTranche
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
