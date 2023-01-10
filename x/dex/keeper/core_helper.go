package keeper

import (
	"context"
	"fmt"
	"math"
	"strings"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const MaxTickExp uint64 = 1048575

///////////////////////////////////////////////////////////////////////////////
//                                   UTILS                                   //
///////////////////////////////////////////////////////////////////////////////

func CreateSharesId(token0 string, token1 string, tickIndex int64, feeIndex uint64) (denom string) {
	t0 := strings.ReplaceAll(token0, "-", "")
	t1 := strings.ReplaceAll(token1, "-", "")
	return fmt.Sprintf("%s-%s-%s-t%d-f%d", types.DepositSharesPrefix, t0, t1, tickIndex, feeIndex)
}

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
			PairId:          pairId,
			CurrentTick0To1: math.MaxInt64,
			CurrentTick1To0: math.MinInt64,
			MinTick:         math.MaxInt64,
			MaxTick:         math.MinInt64,
		}
		k.SetTradingPair(ctx, pair)
	}
	return pair
}

func (k Keeper) InitTick(ctx sdk.Context, pairId string, tickIndex int64) (types.Tick, error) {
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

	token0, token1 := PairToTokens(pairId)
	k.GetOrInitLimitOrderTranche(ctx, pairId, tickIndex, token0, 0)
	k.GetOrInitLimitOrderTranche(ctx, pairId, tickIndex, token1, 0)

	return tick, nil
}

func (k Keeper) GetOrInitTick(goCtx context.Context, pairId string, tickIndex int64) (types.Tick, error) {
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
	pairId string,
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

func (k Keeper) FindNextTick1To0(goCtx context.Context, TradingPair types.TradingPair) (tickIdx int64, found bool) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// If MinTick == MaxInt64 it is unset
	// ie. There is no Token0 in the pool
	if TradingPair.MinTick == math.MaxInt64 {
		return math.MaxInt64, false
	}
	// Start scanning from CurrentTick1To0 - 1
	tickIdx = TradingPair.CurrentTick1To0 - 1

	// Scan through all tick to the left until we hit MinTick
	for tickIdx >= TradingPair.MinTick {
		// Checks for the next value tick containing amount0
		tick, tickFound := k.GetTick(ctx, TradingPair.PairId, tickIdx)
		if tickFound && k.TickHasToken0(ctx, &tick) {
			//Return the new tickIdx
			return tickIdx, true
		}

		tickIdx--
	}

	// If no tick found return false
	return math.MaxInt64, false
}

func (k Keeper) FindNewMinTick(goCtx context.Context, TradingPair types.TradingPair) (minTickIdx int64) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Start scanning from TradingPair.MinTick
	minTickIdx = TradingPair.MinTick

	// Scan through all tick to the left until we hit CurrentTick1To0
	for minTickIdx <= TradingPair.CurrentTick1To0 {
		// Checks for the next value tick containing amount0
		tick, tickFound := k.GetTick(ctx, TradingPair.PairId, minTickIdx)
		if tickFound && k.TickHasToken0(ctx, &tick) {
			//Return the new MinTickIdx
			return minTickIdx
		}

		minTickIdx++
	}

	// If no tick found return false
	return math.MaxInt64
}

func (k Keeper) FindNewMaxTick(goCtx context.Context, TradingPair types.TradingPair) (maxTickIdx int64) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Start scanning from TradingPair.MaxTick
	maxTickIdx = TradingPair.MaxTick

	// Scan through all tick to the left until we hit CurrentTick0To1
	for maxTickIdx >= TradingPair.CurrentTick0To1 {
		// Checks for the next value tick containing amount1
		tick, tickFound := k.GetTick(ctx, TradingPair.PairId, maxTickIdx)
		if tickFound && k.TickHasToken1(ctx, &tick) {
			//Return the new tickIdx
			return maxTickIdx
		}

		maxTickIdx--
	}
	// If no tick found return false
	return math.MinInt64
}

func (k Keeper) FindNextTick0To1(goCtx context.Context, TradingPair types.TradingPair) (tickIdx int64, found bool) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// If MaxTick == MinInt64 it is unset
	// There is no Token1 in the pool
	if TradingPair.MaxTick == math.MinInt64 {
		return math.MinInt64, false
	}
	// Start scanning from CurrentTick0To1 + 1
	tickIdx = TradingPair.CurrentTick0To1 + 1

	// Scan through all tick to the right until we hit MaxTick
	for int64(tickIdx) <= TradingPair.MaxTick {
		// Checks for the next value tick containing amount1
		tick, tickFound := k.GetTick(ctx, TradingPair.PairId, tickIdx)
		if tickFound && k.TickHasToken1(ctx, &tick) {
			// Returns the new tickIdx
			return tickIdx, true
		}

		tickIdx++
	}

	// If no tick found return false
	return math.MinInt64, false
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
	token0, _ := PairToTokens(tick.PairId)
	tranche, found := k.GetLimitOrderTranche(
		ctx,
		tick.PairId,
		tick.TickIndex,
		token0,
		trancheIndex,
	)
	return found && tranche.ReservesTokenIn.GT(sdk.ZeroInt())
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
	_, token1 := PairToTokens(tick.PairId)
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
//                                TICK UPDATES                               //
///////////////////////////////////////////////////////////////////////////////

// should be called for every pair, tick for which token1 is added
func (k Keeper) CalcTickPointersPostAddToken0(goCtx context.Context, pair *types.TradingPair, tick *types.Tick) *types.TradingPair {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if !k.TickHasToken0(ctx, tick) {
		return nil
	}

	tickIndex := tick.TickIndex
	minTick := &pair.MinTick
	cur1To0 := &pair.CurrentTick1To0
	*minTick = MinInt64(*minTick, tickIndex)
	*cur1To0 = MaxInt64(*cur1To0, tickIndex)
	return pair
}

func (k Keeper) UpdateTickPointersPostAddToken0(goCtx context.Context, pair *types.TradingPair, tick *types.Tick) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	newPair := k.CalcTickPointersPostAddToken0(goCtx, pair, tick)
	if newPair != nil {
		k.SetTradingPair(ctx, *newPair)
	}
}

// should be called for every pair, tick for which token1 is added
func (k Keeper) CalcTickPointersPostAddToken1(goCtx context.Context, pair *types.TradingPair, tick *types.Tick) *types.TradingPair {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if !k.TickHasToken1(ctx, tick) {
		return nil
	}

	tickIndex := tick.TickIndex
	cur0To1 := &pair.CurrentTick0To1
	maxTick := &pair.MaxTick
	*cur0To1 = MinInt64(*cur0To1, tickIndex)
	*maxTick = MaxInt64(*maxTick, tickIndex)
	return pair
}

func (k Keeper) UpdateTickPointersPostAddToken1(goCtx context.Context, pair *types.TradingPair, tick *types.Tick) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	newPair := k.CalcTickPointersPostAddToken1(goCtx, pair, tick)
	if newPair != nil {
		k.SetTradingPair(ctx, *newPair)
	}
}

// Should be called for every pair, tick for which token0 liquidity is removed
func (k Keeper) CalcTickPointersPostRemoveToken0(goCtx context.Context, pair *types.TradingPair, tick *types.Tick) *types.TradingPair {
	ctx := sdk.UnwrapSDKContext(goCtx)
	tickIndex := tick.TickIndex
	minTick := &pair.MinTick
	cur1To0 := &pair.CurrentTick1To0

	// return when we're removing liquidity between the bounds
	// Or liquidity is not drained
	if *minTick < tickIndex && tickIndex < *cur1To0 || k.TickHasToken0(ctx, tick) {
		//Do Nothing
		return nil
	}

	// only need to act when the token is exhausted at one of the bounds

	// We have removed all of Token0 from the pool
	if tickIndex == *minTick && tickIndex == *cur1To0 {
		*minTick = math.MaxInt64
		*cur1To0 = math.MinInt64
		// we leave cur1To0 where it is because otherwise we lose the last traded price
	} else if tickIndex == *minTick {
		// Finds the new minTick
		nexMinTick := k.FindNewMinTick(goCtx, *pair)
		*minTick = nexMinTick

		// we are removing liquidity below the current1To0, no need to update that
	} else if tickIndex == *cur1To0 {
		next1To0, found := k.FindNextTick1To0(goCtx, *pair)
		if !found {
			// This case should be impossible if MinTick is tracked correctly
			*minTick = math.MaxInt64
			*cur1To0 = math.MinInt64
		} else {
			*cur1To0 = next1To0
		}
	}

	return pair
}

func (k Keeper) UpdateTickPointersPostRemoveToken0(goCtx context.Context, pair *types.TradingPair, tick *types.Tick) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	newPair := k.CalcTickPointersPostRemoveToken0(goCtx, pair, tick)
	if newPair != nil {
		k.SetTradingPair(ctx, *newPair)
	}
}

// Should be called for every pair, tick for which token1 liquidity is removed
func (k Keeper) CalcTickPointersPostRemoveToken1(goCtx context.Context, pair *types.TradingPair, tick *types.Tick) *types.TradingPair {
	ctx := sdk.UnwrapSDKContext(goCtx)
	tickIndex := tick.TickIndex
	maxTick := &pair.MaxTick
	cur0To1 := &pair.CurrentTick0To1

	// return when we're removing liquidity between the bounds
	// OR liquidity is not drained
	if *cur0To1 < tickIndex && tickIndex < *maxTick || k.TickHasToken1(ctx, tick) {
		// Do nothing
		return nil
	}

	// only need to act when the token is exhausted at one of the bounds
	if tickIndex == *maxTick && tickIndex == *cur0To1 {
		*maxTick = math.MinInt64
		*cur0To1 = math.MaxInt64
		// we leave cur0To1 where it is because otherwise we lose the last traded price
	} else if tickIndex == *maxTick {
		// Finds the new max tick
		nexMaxTick := k.FindNewMaxTick(goCtx, *pair)
		*maxTick = nexMaxTick
		// we are removing liquidity above the current0to1, no need to update that
	} else if tickIndex == *cur0To1 {
		next0To1, found := k.FindNextTick0To1(goCtx, *pair)
		if !found {
			*maxTick = math.MinInt64
			*cur0To1 = math.MaxInt64
			// This case should be impossible if MinTick is tracked correctly
		} else {
			*cur0To1 = next0To1
		}
	}

	return pair
}

func (k Keeper) UpdateTickPointersPostRemoveToken1(goCtx context.Context, pair *types.TradingPair, tick *types.Tick) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	newPair := k.CalcTickPointersPostRemoveToken1(goCtx, pair, tick)
	if newPair != nil {
		k.SetTradingPair(ctx, *newPair)
	}
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
