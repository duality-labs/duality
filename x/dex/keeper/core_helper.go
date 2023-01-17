package keeper

import (
	"context"
	"math"

	. "github.com/NicholasDotSol/duality/utils"
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NOTE: -352,437 is the lowest possible tick at which price can be calculated with a < 1% error
// when using 18 digit decimal precision (via sdk.Dec)
const MaxTickExp uint64 = 352437

///////////////////////////////////////////////////////////////////////////////
//                           GETTERS & INITIALIZERS                          //
///////////////////////////////////////////////////////////////////////////////

func (k Keeper) TokenInit(ctx sdk.Context, address string) {
	_, found := k.GetTokenMap(ctx, address)
	if !found {
		tokenIndex := k.GetTokensCount(ctx)
		newTokenCount := tokenIndex + 1
		// TODO: Consolidate TokenMap and Tokens into one type
		k.SetTokenMap(ctx, types.TokenMap{Address: address, Index: int64(tokenIndex)})
		k.AppendTokens(ctx, types.Tokens{Address: address, Id: tokenIndex})
		k.SetTokensCount(ctx, newTokenCount)
	}
}

// Handles initializing a new pair (token0/token1) if not found, adds token0, token1 to global list of tokens active on the dex
func (k Keeper) GetOrInitPair(goCtx context.Context, token0 string, token1 string) types.TradingPair {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.TokenInit(ctx, token0)
	k.TokenInit(ctx, token1)
	pairId := CreatePairId(token0, token1)
	pair, found := k.GetTradingPair(ctx, pairId)
	if !found {
		pair = types.TradingPair{
			PairId:          &types.PairId{Token0: token0, Token1: token1},
			CurrentTick0To1: math.MaxInt64,
			CurrentTick1To0: math.MinInt64,
			MinTick:         math.MaxInt64,
			MaxTick:         math.MinInt64,
		}
		k.SetTradingPair(ctx, pair)
	}
	return pair
}

func (k Keeper) InitTick(ctx sdk.Context, pairId *types.PairId, tickIndex int64) (types.Tick, error) {
	price0To1, err := CalcPrice0To1(tickIndex)
	if err != nil {
		return types.Tick{}, err
	}
	numFees := k.GetFeeTierCount(ctx)
	tick := types.Tick{
		PairId:    pairId,
		TickIndex: tickIndex,
		Price0To1: &price0To1,
		TickData: &types.TickDataType{
			// TODO: clean up tickdata proto
			Reserve0: make([]sdk.Int, numFees),
			Reserve1: make([]sdk.Int, numFees),
		},
		LimitOrderTranche0To1: &types.LimitTrancheIndexes{0, 0},
		LimitOrderTranche1To0: &types.LimitTrancheIndexes{0, 0},
	}
	for i := 0; i < int(numFees); i++ {
		// TODO: clean up tickdata proto
		tick.TickData.Reserve0[i] = sdk.ZeroInt()
		tick.TickData.Reserve1[i] = sdk.ZeroInt()
	}
	k.SetTick(ctx, pairId, tick)

	token0, token1 := types.PairIdToTokens(pairId)
	k.GetOrInitLimitOrderTranche(ctx, pairId, tickIndex, token0, 0)
	k.GetOrInitLimitOrderTranche(ctx, pairId, tickIndex, token1, 0)

	return tick, nil
}

func (k Keeper) GetOrInitTick(goCtx context.Context, pairId *types.PairId, tickIndex int64) (types.Tick, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	tick, tickFound := k.GetTick(ctx, pairId, tickIndex)
	if tickFound {
		return tick, nil
	} else {
		return k.InitTick(ctx, pairId, tickIndex)
	}
}

func (k Keeper) GetOrInitLimitOrderTrancheUser(
	goCtx context.Context,
	pairId *types.PairId,
	tickIndex int64,
	tokenIn string,
	currentLimitOrderKey uint64,
	receiver string,
) types.LimitOrderTrancheUser {
	ctx := sdk.UnwrapSDKContext(goCtx)

	UserShareData, UserShareDataFound := k.GetLimitOrderTrancheUser(ctx, pairId, tickIndex, tokenIn, currentLimitOrderKey, receiver)

	if !UserShareDataFound {
		return types.LimitOrderTrancheUser{
			Count:           currentLimitOrderKey,
			Address:         receiver,
			SharesOwned:     sdk.ZeroInt(),
			SharesWithdrawn: sdk.ZeroInt(),
			SharesCancelled: sdk.ZeroInt(),
			TickIndex:       tickIndex,
			Token:           tokenIn,
			PairId:          pairId,
		}
	}

	return UserShareData
}

///////////////////////////////////////////////////////////////////////////////
//                          STATE CALCULATIONS                               //
///////////////////////////////////////////////////////////////////////////////

func (k Keeper) InitLiquidity(dp *types.DirectionalTradingPair, tickIndex int64) {
	if dp.IsTokenInToken0() {
		InitLiquidityToken0(&dp.TradingPair, tickIndex)
	} else {
		InitLiquidityToken1(&dp.TradingPair, tickIndex)
	}
}

func (k Keeper) DeinitLiquidity(ctx context.Context, dp *types.DirectionalTradingPair, tickIndex int64) {
	if dp.IsTokenOutToken0() {
		k.DeinitLiquidityToken0(ctx, &dp.TradingPair, tickIndex)
	} else {
		k.DeinitLiquidityToken1(ctx, &dp.TradingPair, tickIndex)
	}
}

// Assumes that the token0 liquidity is non-empty at this tick
func InitLiquidityToken0(pair *types.TradingPair, tickIndex int64) {
	minTick := &pair.MinTick
	curTick1To0 := &pair.CurrentTick1To0
	*minTick = MinInt64(*minTick, tickIndex)
	*curTick1To0 = MaxInt64(*curTick1To0, tickIndex)
}

// Assumes that the token1 liquidity is non-empty at this tick
func InitLiquidityToken1(pair *types.TradingPair, tickIndex int64) {
	maxTick := &pair.MaxTick
	curTick0To1 := &pair.CurrentTick0To1
	*maxTick = MaxInt64(*maxTick, tickIndex)
	*curTick0To1 = MinInt64(*curTick0To1, tickIndex)
}

// Assumes that the token0 liquidity is empty at this tick
func (k Keeper) DeinitLiquidityToken0(ctx context.Context, pair *types.TradingPair, tickIndex int64) {
	minTick := &pair.MinTick
	cur1To0 := &pair.CurrentTick1To0

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
		nexMinTick := k.FindNewMinTick(ctx, *pair)
		*minTick = nexMinTick

		// we are removing liquidity below the current1To0, no need to update that
	} else if tickIndex == *cur1To0 {
		next1To0, found := k.FindNextTick1To0(ctx, *pair)
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
func (k Keeper) DeinitLiquidityToken1(ctx context.Context, pair *types.TradingPair, tickIndex int64) {
	maxTick := &pair.MaxTick
	cur0To1 := &pair.CurrentTick0To1

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
		nextMaxTick := k.FindNewMaxTick(ctx, *pair)
		*maxTick = nextMaxTick

		// we are removing liquidity below the current1To0, no need to update that
	} else if tickIndex == *cur0To1 {
		next0To1, found := k.FindNextTick0To1(ctx, *pair)
		if !found {
			// This case should be impossible if MinTick is tracked correctly
			*maxTick = math.MinInt64
			*cur0To1 = math.MaxInt64
		} else {
			*cur0To1 = next0To1
		}
	}
}

func (k Keeper) UpdateTickPointersPostAddToken0(goCtx context.Context, pair *types.TradingPair, tick *types.Tick) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// TODO: Get rid of this, ideally should know exactly when to (de)init
	// 	and should not have to have the check before.
	if k.TickHasToken0(ctx, tick) {
		InitLiquidityToken0(pair, tick.TickIndex)
	}

}

func (k Keeper) UpdateTickPointersPostAddToken1(goCtx context.Context, pair *types.TradingPair, tick *types.Tick) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// TODO: Get rid of this, ideally should know exactly when to (de)init
	//	and should not have to have the check before.
	if k.TickHasToken1(ctx, tick) {
		InitLiquidityToken1(pair, tick.TickIndex)
	}
}

func (k Keeper) UpdateTickPointersPostRemoveToken0(goCtx context.Context, pair *types.TradingPair, tick *types.Tick) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// TODO: Get rid of this, ideally should know exactly when to (de)init
	//	and should not have to have the check before.
	if !k.TickHasToken0(ctx, tick) {
		k.DeinitLiquidityToken0(goCtx, pair, tick.TickIndex)
	}
}

func (k Keeper) UpdateTickPointersPostRemoveToken1(goCtx context.Context, pair *types.TradingPair, tick *types.Tick) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// TODO: Get rid of this, ideally should know exactly when to (de)init
	//	and should not have to have the check before.
	if !k.TickHasToken1(ctx, tick) {
		k.DeinitLiquidityToken1(goCtx, pair, tick.TickIndex)
	}
}

func (k Keeper) FindNextTick1To0(goCtx context.Context, pair types.TradingPair) (tickIdx int64, found bool) {

	sdkCtx := sdk.UnwrapSDKContext(goCtx)
	// If MinTick == MaxInt64 it is unset
	// ie. There is no Token0 in the pool
	if pair.MinTick == math.MaxInt64 {
		return math.MaxInt64, false
	}
	// Start scanning from CurrentTick1To0 - 1
	tickIdx = pair.CurrentTick1To0 - 1

	ti := k.NewTickIterator(goCtx, tickIdx, pair.MinTick, pair.PairId, true)

	defer ti.Close()
	for ; ti.Valid(); ti.Next() {
		tick := ti.Value()
		if k.TickHasToken0(sdkCtx, &tick) {
			return tick.TickIndex, true
		}
	}
	return math.MinInt64, false

}

func (k Keeper) FindNextTick0To1(goCtx context.Context, pair types.TradingPair) (tickIdx int64, found bool) {
	sdkCtx := sdk.UnwrapSDKContext(goCtx)
	// If MaxTick == MinInt64 it is unset
	// There is no Token1 in the pool
	if pair.MaxTick == math.MinInt64 {
		return math.MinInt64, false
	}

	// Start scanning from CurrentTick0To1 + 1
	tickIdx = pair.CurrentTick0To1 + 1
	ti := k.NewTickIterator(goCtx, tickIdx, pair.MaxTick, pair.PairId, false)
	defer ti.Close()
	for ; ti.Valid(); ti.Next() {
		tick := ti.Value()
		if k.TickHasToken1(sdkCtx, &tick) {
			return tick.TickIndex, true
		}
	}

	return math.MinInt64, false
}

func (k Keeper) FindNewMinTick(goCtx context.Context, pair types.TradingPair) (minTickIdx int64) {
	sdkCtx := sdk.UnwrapSDKContext(goCtx)

	ti := k.NewTickIterator(goCtx, pair.MinTick, pair.CurrentTick1To0, pair.PairId, false)
	defer ti.Close()
	for ; ti.Valid(); ti.Next() {
		tick := ti.Value()
		if k.TickHasToken0(sdkCtx, &tick) {
			return tick.TickIndex
		}
	}
	return math.MaxInt64

}

func (k Keeper) FindNewMaxTick(goCtx context.Context, pair types.TradingPair) (maxTickIdx int64) {
	sdkCtx := sdk.UnwrapSDKContext(goCtx)

	ti := k.NewTickIterator(goCtx, pair.MaxTick, pair.CurrentTick0To1, pair.PairId, true)
	defer ti.Close()
	for ; ti.Valid(); ti.Next() {
		tick := ti.Value()
		if k.TickHasToken1(sdkCtx, &tick) {
			return tick.TickIndex
		}
	}
	return math.MinInt64
}

// Balance trueAmount1 to the pool ratio
func CalcTrueAmounts(
	centerTickPrice1To0 sdk.Dec,
	lowerReserve0 sdk.Int,
	upperReserve1 sdk.Int,
	amount0 sdk.Int,
	amount1 sdk.Int,
	totalShares sdk.Int,
) (trueAmount0 sdk.Int, trueAmount1 sdk.Int, sharesMinted sdk.Int) {
	lowerReserve0Dec := lowerReserve0.ToDec()
	upperReserve1Dec := upperReserve1.ToDec()
	amount0Dec := amount0.ToDec()
	amount1Dec := amount1.ToDec()

	// See spec: https://www.notion.so/dualityxyz/Autoswap-Spec-e856fa7b2438403c95147010d479b98c
	if upperReserve1Dec.GT(sdk.ZeroDec()) {
		trueAmount0 = sdk.MinDec(
			amount0Dec,
			amount1Dec.Mul(lowerReserve0Dec).Quo(upperReserve1Dec)).TruncateInt()
	} else {
		trueAmount0 = amount0
	}

	if lowerReserve0Dec.GT(sdk.ZeroDec()) {
		trueAmount1 = sdk.MinDec(
			amount1Dec,
			amount0Dec.Mul(upperReserve1Dec).Quo(lowerReserve0Dec)).TruncateInt()
	} else {
		trueAmount1 = amount1
	}

	valueMintedToken0 := CalcShares(trueAmount0, trueAmount1, centerTickPrice1To0)
	valueExistingToken0 := CalcShares(lowerReserve0, upperReserve1, centerTickPrice1To0)
	if valueExistingToken0.GT(sdk.ZeroDec()) {
		sharesMinted = valueMintedToken0.Quo(valueExistingToken0).MulInt(totalShares).TruncateInt()
	} else {
		sharesMinted = valueMintedToken0.TruncateInt()
	}
	return
}
func IsTickOutOfRange(tickIndex int64) bool {
	absTickIndex := Abs(tickIndex)
	return absTickIndex > MaxTickExp
}

func MustCalcPrice0To1(tickIndex int64) sdk.Dec {
	price, err := CalcPrice0To1(tickIndex)
	if err != nil {
		panic(err)
	}
	return price
}

// Calculates the price for a swap from token 0 to token 1 given a tick
// tickIndex refers to the index of a specified tick
func CalcPrice0To1(tickIndex int64) (sdk.Dec, error) {
	if IsTickOutOfRange(tickIndex) {
		return sdk.ZeroDec(), types.ErrTickOutsideRange
	}

	if 0 <= tickIndex {
		return sdk.OneDec().Quo(Pow(BasePrice(), uint64(tickIndex))), nil
	} else {
		return Pow(BasePrice(), uint64(-1*tickIndex)), nil
	}
}

func MustCalcPrice1To0(tickIndex int64) sdk.Dec {
	price, err := CalcPrice1To0(tickIndex)
	if err != nil {
		panic(err)
	}
	return price
}

// Calculates the price for a swap from token 1 to token 0 given a tick
// tickIndex refers to the index of a specified tick
func CalcPrice1To0(tickIndex int64) (sdk.Dec, error) {

	if IsTickOutOfRange(tickIndex) {
		return sdk.ZeroDec(), types.ErrTickOutsideRange
	}

	if 0 <= tickIndex {
		return Pow(BasePrice(), uint64(tickIndex)), nil
	} else {
		return sdk.OneDec().Quo(Pow(BasePrice(), uint64(-1*tickIndex))), nil
	}
}

func (k Keeper) ShouldDeinit(
	ctx sdk.Context,
	deinitedTick *types.Tick,
	tradingPair types.DirectionalTradingPair) bool {

	if deinitedTick == nil {
		return false
	}
	if tradingPair.IsTokenOutToken0() {
		return !k.TickHasToken0(ctx, deinitedTick)
	} else {
		return !k.TickHasToken1(ctx, deinitedTick)
	}
}

// Checks if a tick has reserves0 at any fee tier
func (k Keeper) TickHasToken0(ctx sdk.Context, tick *types.Tick) bool {
	// TODO: clean up tickdata proto
	for _, r0 := range tick.TickData.Reserve0 {
		if r0.GT(sdk.ZeroInt()) {
			return true
		}
	}

	for i := tick.LimitOrderTranche0To1.FillTrancheIndex; i <= tick.LimitOrderTranche0To1.PlaceTrancheIndex; i++ {
		if k.TickTrancheHasToken0(ctx, tick, i) {
			return true
		}
	}

	return false
}

func (k Keeper) TickTrancheHasToken0(ctx sdk.Context, tick *types.Tick, trancheIndex uint64) bool {
	token0, _ := types.PairIdToTokens(tick.PairId)
	tranche, found := k.GetLimitOrderTranche(
		ctx,
		tick.PairId,
		tick.TickIndex,
		token0,
		trancheIndex,
	)
	if found && tranche.ReservesTokenIn.GT(sdk.ZeroInt()) {
		return true
	} else {
		return false
	}
}

// Checks if a tick has reserve1 at any fee tier
func (k Keeper) TickHasToken1(ctx sdk.Context, tick *types.Tick) bool {
	for _, s := range tick.TickData.Reserve1 {
		if s.GT(sdk.ZeroInt()) {
			return true
		}
	}

	for i := tick.LimitOrderTranche1To0.FillTrancheIndex; i <= tick.LimitOrderTranche1To0.PlaceTrancheIndex; i++ {
		if k.TickTrancheHasToken1(ctx, tick, i) {
			return true
		}
	}

	return false
}

func (k Keeper) TickTrancheHasToken1(ctx sdk.Context, tick *types.Tick, trancheIndex uint64) bool {
	_, token1 := types.PairIdToTokens(tick.PairId)
	tranche, found := k.GetLimitOrderTranche(
		ctx,
		tick.PairId,
		tick.TickIndex,
		token1,
		trancheIndex,
	)
	return found && tranche.ReservesTokenIn.GT(sdk.ZeroInt())
}

///////////////////////////////////////////////////////////////////////////////
//                            TOKENIZER UTILS                                //
///////////////////////////////////////////////////////////////////////////////

func (k Keeper) MintShares(ctx sdk.Context, addr sdk.AccAddress, amount sdk.Int, sharesId string) error {
	// mint share tokens
	sharesCoins := sdk.Coins{sdk.NewCoin(sharesId, amount)}
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sharesCoins); err != nil {
		return err
	}
	// transfer them to addr
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, sharesCoins); err != nil {
		return err
	}
	return nil
}

func (k Keeper) BurnShares(ctx sdk.Context, addr sdk.AccAddress, amount sdk.Int, sharesId string) error {
	sharesCoins := sdk.Coins{sdk.NewCoin(sharesId, amount)}
	// transfer tokens to module
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, sharesCoins); err != nil {
		return err
	}
	// burn tokens
	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sharesCoins); err != nil {
		return err
	}
	return nil
}
