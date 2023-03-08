package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/incentives/types"
)

// SetIncentivePlan set a specific incentivePlan in the store from its index
func (k Keeper) SetIncentivePlan(ctx sdk.Context, incentivePlan types.IncentivePlan) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IncentivePlanKeyPrefix))
	b := k.cdc.MustMarshal(&incentivePlan)
	store.Set(types.IncentivePlanKey(
		incentivePlan.Index,
	), b)
}

// GetIncentivePlan returns a incentivePlan from its index
func (k Keeper) GetIncentivePlan(
	ctx sdk.Context,
	index string,

) (val types.IncentivePlan, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IncentivePlanKeyPrefix))

	b := store.Get(types.IncentivePlanKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveIncentivePlan removes a incentivePlan from the store
func (k Keeper) RemoveIncentivePlan(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IncentivePlanKeyPrefix))
	store.Delete(types.IncentivePlanKey(
		index,
	))
}

// GetAllIncentivePlan returns all incentivePlan
func (k Keeper) GetAllIncentivePlan(ctx sdk.Context) (list []types.IncentivePlan) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IncentivePlanKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.IncentivePlan
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
