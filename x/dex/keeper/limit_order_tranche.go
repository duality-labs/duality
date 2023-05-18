package keeper

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/duality-labs/duality/x/dex/utils"
)

func (k Keeper) FindLimitOrderTranche(
	ctx sdk.Context,
	tradePairID *types.TradePairID,
	tickIndex int64,
	trancheKey string,
) (val *types.LimitOrderTranche, fromFilled, found bool) {
	// Try to find the tranche in the active liq index
	tick, found := k.GetLimitOrderTranche(ctx, tradePairID, tickIndex, trancheKey)
	if found {
		return tick, false, true
	}
	// Look for filled limit orders
	tranche, found := k.GetInactiveLimitOrderTranche(ctx, tradePairID, tickIndex, trancheKey)
	if found {
		return tranche, true, true
	}

	return nil, false, false
}

func (k Keeper) SaveTranche(ctx sdk.Context, tranche *types.LimitOrderTranche) {
	if tranche.HasTokenIn() {
		k.SetLimitOrderTranche(ctx, tranche)
	} else {
		k.SetInactiveLimitOrderTranche(ctx, tranche)
		k.RemoveLimitOrderTranche(ctx, tranche)
	}

	ctx.EventManager().EmitEvent(types.CreateTickUpdateLimitOrderTranche(tranche))
}

func (k Keeper) SetLimitOrderTranche(ctx sdk.Context, tranche *types.LimitOrderTranche) {
	// Wrap tranche back into TickLiquidity
	tick := types.TickLiquidity{
		Liquidity: &types.TickLiquidity_LimitOrderTranche{
			LimitOrderTranche: tranche,
		},
	}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TickLiquidityKeyPrefix))
	b := k.cdc.MustMarshal(&tick)
	store.Set(types.TickLiquidityKey(
		tranche.TradePairID,
		tranche.TickIndex,
		types.LiquidityTypeLimitOrder,
		tranche.TrancheKey,
	), b)
}

func (k Keeper) GetLimitOrderTranche(
	ctx sdk.Context,
	tradePairID *types.TradePairID,
	tickIndex int64,
	trancheKey string,
) (tranche *types.LimitOrderTranche, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TickLiquidityKeyPrefix))
	b := store.Get(types.TickLiquidityKey(
		tradePairID,
		tickIndex,
		types.LiquidityTypeLimitOrder,
		trancheKey,
	))

	if b == nil {
		return nil, false
	}

	var tick types.TickLiquidity
	k.cdc.MustUnmarshal(b, &tick)

	return tick.GetLimitOrderTranche(), true
}

func (k Keeper) GetLimitOrderTrancheByKey(
	ctx sdk.Context,
	key []byte,
) (tranche *types.LimitOrderTranche, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TickLiquidityKeyPrefix))
	b := store.Get(key)

	if b == nil {
		return nil, false
	}

	var tick types.TickLiquidity
	k.cdc.MustUnmarshal(b, &tick)

	return tick.GetLimitOrderTranche(), true
}

func (k Keeper) RemoveLimitOrderTranche(ctx sdk.Context, tranche *types.LimitOrderTranche) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TickLiquidityKeyPrefix))
	store.Delete(types.TickLiquidityKey(
		tranche.TradePairID,
		tranche.TickIndex,
		types.LiquidityTypeLimitOrder,
		tranche.TrancheKey,
	))
}

func (k Keeper) GetPlaceTranche(
	sdkCtx sdk.Context,
	tradePairID *types.TradePairID,
	tickIndex int64,
) *types.LimitOrderTranche {
	prefixStore := prefix.NewStore(
		sdkCtx.KVStore(k.storeKey),
		types.TickLiquidityLimitOrderPrefix(tradePairID, tickIndex),
	)
	iter := prefixStore.Iterator(nil, nil)

	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var tick types.TickLiquidity
		k.cdc.MustUnmarshal(iter.Value(), &tick)
		tranche := tick.GetLimitOrderTranche()
		if tranche.IsPlaceTranche() {
			return tranche
		}
	}

	return nil
}

func (k Keeper) GetFillTranche(
	sdkCtx sdk.Context,
	tradePairID *types.TradePairID,
	tickIndex int64,
) (*types.LimitOrderTranche, bool) {
	prefixStore := prefix.NewStore(
		sdkCtx.KVStore(k.storeKey),
		types.TickLiquidityLimitOrderPrefix(tradePairID, tickIndex),
	)
	iter := sdk.KVStorePrefixIterator(prefixStore, []byte{})

	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var tick types.TickLiquidity
		k.cdc.MustUnmarshal(iter.Value(), &tick)

		return tick.GetLimitOrderTranche(), true
	}

	return &types.LimitOrderTranche{}, false
}

func (k Keeper) GetAllLimitOrderTrancheAtIndex(
	sdkCtx sdk.Context,
	tradePairID *types.TradePairID,
	tickIndex int64,
) (trancheList []types.LimitOrderTranche) {
	prefixStore := prefix.NewStore(
		sdkCtx.KVStore(k.storeKey),
		types.TickLiquidityLimitOrderPrefix(tradePairID, tickIndex),
	)
	iter := sdk.KVStorePrefixIterator(prefixStore, []byte{})

	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var tick types.TickLiquidity
		k.cdc.MustUnmarshal(iter.Value(), &tick)
		maybeTranche := tick.GetLimitOrderTranche()
		if maybeTranche != nil {
			trancheList = append(trancheList, *maybeTranche)
		}
	}

	return trancheList
}

func NewTrancheKey(sdkCtx sdk.Context) string {
	blockHeight := sdkCtx.BlockHeight()
	txGas := sdkCtx.GasMeter().GasConsumed()

	var blockGas uint64 = 0

	// TEMP: REMOVE ME
	if sdkCtx.BlockGasMeter() != nil {
		blockGas = sdkCtx.BlockGasMeter().GasConsumed()
	}

	totalGas := blockGas + txGas

	blockStr := utils.Uint64ToSortableString(uint64(blockHeight))
	gasStr := utils.Uint64ToSortableString(totalGas)

	return fmt.Sprintf("%s%s", blockStr, gasStr)
}

func (k Keeper) GetOrInitPlaceTranche(ctx sdk.Context,
	tradePairID *types.TradePairID,
	tickIndex int64,
	goodTil *time.Time,
	orderType types.LimitOrderType,
) (placeTranche *types.LimitOrderTranche, err error) {
	// NOTE: Right now we are not indexing by goodTil date so we can't easily check if there's already a tranche
	// with the same goodTil date so instead we create a new tranche for each goodTil order
	// if there is a large number of limitOrders with the same goodTilTime (most likely JIT)
	// aggregating might be more efficient particularly for deletion, but if they are relatively sparse
	// it will incur fewer lookups to just create a new limitOrderTranche
	// Also trying to cancel aggregated good_til orders will be a PITA
	JITGoodTilTime := types.JITGoodTilTime()
	switch orderType {
	case types.LimitOrderType_JUST_IN_TIME:
		placeTranche, err = NewLimitOrderTranche(ctx, tradePairID, tickIndex, &JITGoodTilTime)
	case types.LimitOrderType_GOOD_TIL_TIME:
		placeTranche, err = NewLimitOrderTranche(ctx, tradePairID, tickIndex, goodTil)
	default:
		placeTranche = k.GetPlaceTranche(ctx, tradePairID, tickIndex)
		if placeTranche == nil {
			placeTranche, err = NewLimitOrderTranche(ctx, tradePairID, tickIndex, nil)
			if err != nil {
				return nil, err
			}
		}
	}
	if err != nil {
		return nil, err
	}

	return placeTranche, nil
}
