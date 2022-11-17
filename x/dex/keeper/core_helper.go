package keeper

import (
	"context"
	"math"
	"strings"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

///////////////////////////////////////////////////////////////////////////////
//                                   UTILS                                   //
///////////////////////////////////////////////////////////////////////////////

func PairToTokens(pairId string) (token0 string, token1 string) {
	tokens := strings.Split(pairId, "/")

	return tokens[0], tokens[1]
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
func (k Keeper) GetOrInitPair(goCtx context.Context, token0 string, token1 string) types.PairMap {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.TokenInit(ctx, token0)
	k.TokenInit(ctx, token1)
	pairId := k.CreatePairId(token0, token1)
	pair, found := k.GetPairMap(ctx, pairId)
	if !found {
		pair = types.PairMap{
			PairId: pairId,
			TokenPair: &types.TokenPairType{
				CurrentTick0To1: math.MinInt64,
				CurrentTick1To0: math.MaxInt64,
			},
			MinTick: math.MaxInt64,
			MaxTick: math.MinInt64,
		}
		k.SetPairMap(ctx, pair)
	}
	return pair
}

func (k Keeper) GetOrInitTickTrancheFillMap(goCtx context.Context, pairId string, tickIndex int64, trancheIndex uint64, token string) types.LimitOrderPoolFillMap {
	ctx := sdk.UnwrapSDKContext(goCtx)
	tickTranchFillMap, found := k.GetLimitOrderPoolFillMap(ctx, pairId, tickIndex, token, trancheIndex)
	if !found {
		tickTranchFillMap = types.LimitOrderPoolFillMap{
			PairId:         pairId,
			TickIndex:      tickIndex,
			Token:          token,
			Count:          trancheIndex,
			FilledReserves: sdk.ZeroDec(),
		}
		k.SetLimitOrderPoolFillMap(ctx, tickTranchFillMap)
	}
	return tickTranchFillMap
}

func (k Keeper) GetOrInitTick(goCtx context.Context, pairId string, tickIndex int64) types.TickMap {
	ctx := sdk.UnwrapSDKContext(goCtx)

	tick, tickFound := k.GetTickMap(ctx, pairId, tickIndex)
	if !tickFound {
		numFees := k.GetFeeListCount(ctx)
		tick = types.TickMap{
			PairId:    pairId,
			TickIndex: tickIndex,
			TickData: &types.TickDataType{
				Reserve0AndShares: make([]*types.Reserve0AndSharesType, numFees),
				Reserve1:          make([]sdk.Dec, numFees),
			},
			LimitOrderPool0To1: &types.LimitOrderPool{0, 0},
			LimitOrderPool1To0: &types.LimitOrderPool{0, 0},
		}
		for i := 0; i < int(numFees); i++ {
			tick.TickData.Reserve0AndShares[i] = &types.Reserve0AndSharesType{sdk.ZeroDec(), sdk.ZeroDec()}
			tick.TickData.Reserve1[i] = sdk.ZeroDec()
		}
		k.SetTickMap(ctx, pairId, tick)

		token0, token1 := PairToTokens(pairId)
		k.GetOrInitTickTrancheFillMap(goCtx, pairId, tickIndex, 0, token0)
		k.GetOrInitTickTrancheFillMap(goCtx, pairId, tickIndex, 0, token1)
	}
	return tick
}

func (k Keeper) GetOrInitReserveData(
	goCtx context.Context,
	pairId string,
	tickIndex int64,
	tokenIn string,
	currentLimitOrderKey uint64,
) types.LimitOrderPoolReserveMap {
	ctx := sdk.UnwrapSDKContext(goCtx)

	ReserveData, ReserveDataFound := k.GetLimitOrderPoolReserveMap(ctx, pairId, tickIndex, tokenIn, currentLimitOrderKey)

	if !ReserveDataFound {
		ReserveData.Count = currentLimitOrderKey
		ReserveData.Reserves = sdk.ZeroDec()
		ReserveData.TickIndex = tickIndex
		ReserveData.Token = tokenIn
		ReserveData.PairId = pairId
	}

	return ReserveData
}

func (k Keeper) GetOrInitUserShareData(
	goCtx context.Context,
	pairId string,
	tickIndex int64,
	tokenIn string,
	currentLimitOrderKey uint64,
	receiver string,
) types.LimitOrderPoolUserShareMap {
	ctx := sdk.UnwrapSDKContext(goCtx)

	UserShareData, UserShareDataFound := k.GetLimitOrderPoolUserShareMap(ctx, pairId, tickIndex, tokenIn, currentLimitOrderKey, receiver)

	// If UserShareData object is not found initialize it accordingly
	if !UserShareDataFound {
		UserShareData.Count = currentLimitOrderKey
		UserShareData.Address = receiver
		UserShareData.SharesOwned = sdk.ZeroDec()
		UserShareData.TickIndex = tickIndex
		UserShareData.Token = tokenIn
		UserShareData.PairId = pairId
	}

	return UserShareData
}

func (k Keeper) GetOrInitOrderPoolTotalShares(
	goCtx context.Context,
	pairId string,
	tickIndex int64,
	tokenIn string,
	currentLimitOrderKey uint64,
) types.LimitOrderPoolTotalSharesMap {
	ctx := sdk.UnwrapSDKContext(goCtx)

	TotalSharesData, TotalSharesDataFound := k.GetLimitOrderPoolTotalSharesMap(ctx, pairId, tickIndex, tokenIn, currentLimitOrderKey)

	// If TotalSharesData object is nout found initialize it accordingly
	if !TotalSharesDataFound {
		TotalSharesData.Count = currentLimitOrderKey
		TotalSharesData.TotalShares = sdk.ZeroDec()
		TotalSharesData.TickIndex = tickIndex
		TotalSharesData.Token = tokenIn
		TotalSharesData.PairId = pairId
	}

	return TotalSharesData
}

// Helper function for Place Limit order retrieving and or initializing mapppings used for keeping track of limit orders
// Note: FillMap initialization is handled seperately in placeLimitOrder as it needed prior to this function being called.
func (k Keeper) GetOrInitLimitOrderMaps(
	goCtx context.Context,
	pairId string,
	tickIndex int64,
	tokenIn string,
	currentLimitOrderKey uint64,
	receiver string,
) (types.LimitOrderPoolReserveMap, types.LimitOrderPoolUserShareMap, types.LimitOrderPoolTotalSharesMap) {
	ReserveData := k.GetOrInitReserveData(goCtx, pairId, tickIndex, tokenIn, currentLimitOrderKey)
	UserShareData := k.GetOrInitUserShareData(goCtx, pairId, tickIndex, tokenIn, currentLimitOrderKey, receiver)
	TotalSharesData := k.GetOrInitOrderPoolTotalShares(goCtx, pairId, tickIndex, tokenIn, currentLimitOrderKey)

	return ReserveData, UserShareData, TotalSharesData
}

///////////////////////////////////////////////////////////////////////////////
//                          STATE CALCULATIONS                               //
///////////////////////////////////////////////////////////////////////////////

func (k Keeper) FindNextTick1To0(goCtx context.Context, pairMap types.PairMap) (tickIdx int64, found bool) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// If MinTick == MaxInt64 it is unset
	// ie. There is no Token0 in the pool
	if pairMap.MinTick == math.MaxInt64 {
		return math.MaxInt64, false
	}
	// Start scanning from CurrentTick1To0 - 1
	tickIdx = pairMap.TokenPair.CurrentTick1To0 - 1

	// Scan through all tick to the left until we hit MinTick
	for tickIdx >= pairMap.MinTick {
		// Checks for the next value tick containing amount0
		tick, tickFound := k.GetTickMap(ctx, pairMap.PairId, tickIdx)
		if tickFound && k.HasToken0(ctx, &tick) {
			//Return the new tickIdx
			return tickIdx, true
		}

		tickIdx--
	}

	// If no tick found return false
	return math.MaxInt64, false
}

func (k Keeper) FindNextTick0To1(goCtx context.Context, pairMap types.PairMap) (tickIdx int64, found bool) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// If MaxTick == MinInt64 it is unset
	// There is no Token1 in the pool
	if pairMap.MaxTick == math.MinInt64 {
		return math.MinInt64, false
	}
	// Start scanning from CurrentTick0To1 + 1
	tickIdx = pairMap.TokenPair.CurrentTick0To1 + 1

	// Scan through all tick to the right until we hit MaxTick
	for int64(tickIdx) <= pairMap.MaxTick {
		// Checks for the next value tick containing amount1
		tick, tickFound := k.GetTickMap(ctx, pairMap.PairId, tickIdx)
		if tickFound && k.HasToken1(ctx, &tick) {
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
	lowerReserve0 sdk.Dec,
	upperReserve1 sdk.Dec,
	amount0 sdk.Dec,
	amount1 sdk.Dec,
	totalShares sdk.Dec,
) (trueAmount0 sdk.Dec, trueAmount1 sdk.Dec, sharesMinted sdk.Dec) {
	if lowerReserve0.GT(sdk.ZeroDec()) && upperReserve1.GT(sdk.ZeroDec()) {
		ratio0 := amount0.Quo(lowerReserve0)
		ratio1 := amount1.Quo(upperReserve1)
		if ratio0.LT(ratio1) {
			trueAmount0 = amount0
			trueAmount1 = ratio0.Mul(upperReserve1)
		} else {
			trueAmount0 = ratio1.Mul(lowerReserve0)
			trueAmount1 = amount1
		}
	} else if lowerReserve0.GT(sdk.ZeroDec()) { // && upperReserve1 == 0
		trueAmount0 = amount0
		trueAmount1 = sdk.ZeroDec()
	} else if upperReserve1.GT(sdk.ZeroDec()) { // && lowerReserve0 == 0
		trueAmount0 = sdk.ZeroDec()
		trueAmount1 = amount1
	} else {
		trueAmount0 = amount0
		trueAmount1 = amount1
	}
	valueMintedToken0 := trueAmount1.Mul(centerTickPrice1To0).Add(trueAmount0)
	valueExistingToken0 := upperReserve1.Mul(centerTickPrice1To0).Add(lowerReserve0)
	if valueExistingToken0.GT(sdk.ZeroDec()) {
		sharesMinted = valueMintedToken0.Quo(valueExistingToken0).Mul(totalShares)
	} else {
		sharesMinted = valueMintedToken0
	}
	return
}

// Calculates the price for a swap from token 0 to token 1 given a tick
// tickIndex refers to the index of a specified tick
func (k Keeper) CalcPrice0To1(tickIndex int64) sdk.Dec {
	if 0 <= tickIndex {
		return sdk.OneDec().Quo(Pow(BasePrice(), uint64(tickIndex)))
	} else {
		return Pow(BasePrice(), uint64(-1*tickIndex))
	}
}

// Calculates the price for a swap from token 1 to token 0 given a tick
// tickIndex refers to the index of a specified tick
func (k Keeper) CalcPrice1To0(tickIndex int64) sdk.Dec {
	if 0 <= tickIndex {
		return Pow(BasePrice(), uint64(tickIndex))
	} else {
		return sdk.OneDec().Quo(Pow(BasePrice(), uint64(-1*tickIndex)))
	}
}

// Checks if a tick has reserves0 at any fee tier
func (k Keeper) HasToken0(ctx sdk.Context, tick *types.TickMap) bool {
	for _, s := range tick.TickData.Reserve0AndShares {
		if s.Reserve0.GT(sdk.ZeroDec()) {
			return true
		}
	}

	token0, _ := PairToTokens(tick.PairId)
	reserve, reserveFound := k.GetLimitOrderPoolReserveMap(
		ctx,
		tick.PairId,
		tick.TickIndex,
		token0,
		tick.LimitOrderPool0To1.CurrentLimitOrderKey,
	)
	return reserveFound && reserve.Reserves.GT(sdk.ZeroDec())
}

// Checks if a tick has reserve1 at any fee tier
func (k Keeper) HasToken1(ctx sdk.Context, tick *types.TickMap) bool {
	// check LP tokens
	for _, s := range tick.TickData.Reserve1 {
		if s.GT(sdk.ZeroDec()) {
			return true
		}
	}

	_, token1 := PairToTokens(tick.PairId)
	reserve, reserveFound := k.GetLimitOrderPoolReserveMap(
		ctx,
		tick.PairId,
		tick.TickIndex,
		token1,
		tick.LimitOrderPool1To0.CurrentLimitOrderKey,
	)
	return reserveFound && reserve.Reserves.GT(sdk.ZeroDec())
}

// Currently Unused
// caclulates totalReserves for token0 and token1 for all fee tiers of a given tick.
func (k Keeper) GetTotalReservesAtTick(goCtx context.Context, pairId string, tick_index_ int64, swap0to1 bool) (sdk.Dec, sdk.Dec, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	feelist := k.GetAllFeeList(ctx)

	// inits totalReserve of 0 and 1 for all feeTiers
	var totalReserve0 = sdk.ZeroDec()
	var totalReserve1 = sdk.ZeroDec()

	// retrivies tick from tickMaping
	tick, tickFound := k.GetTickMap(ctx, pairId, tick_index_)

	// verifies that tick at the given tick index exists
	if !tickFound {
		return sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrValidTickNotFound, "No tick found at index %d", tick_index_)
	}

	// When we init a pair we init reserve0, reserve1 to 0 for all feetiers and thus we can iterate over the fee tiers without worrying about nil values.
	for i, _ := range feelist {

		if swap0to1 {
			// Given a tickIndex of reserve0 calculate the totalReserves for the tick composted of reserve0 and the related reserve1
			totalReserve0 = totalReserve0.Add(tick.TickData.Reserve0AndShares[i].Reserve0)
			feeValue := feelist[i].Fee
			totalReserve1 = totalReserve1.Add(tick.TickData.Reserve1[i-int(feeValue)])
		} else {
			// Given a tickIndex of reserve1 calculate the totalReserves for the tick composted of reserve0 and the related reserve0
			totalReserve1 = totalReserve1.Add(tick.TickData.Reserve1[i])
			feeValue := feelist[i].Fee
			totalReserve0 = totalReserve0.Add(tick.TickData.Reserve0AndShares[i+int(feeValue)].Reserve0)

		}

	}

	return totalReserve0, totalReserve1, nil
}

///////////////////////////////////////////////////////////////////////////////
//                                TICK UPDATES                               //
///////////////////////////////////////////////////////////////////////////////

// should be called for every pair, tick for which token1 is added
func (k Keeper) CalcTickPointersPostAddToken0(goCtx context.Context, pair *types.PairMap, tick *types.TickMap) *types.PairMap {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if !k.HasToken0(ctx, tick) {
		return nil
	}

	tickIndex := tick.TickIndex
	minTick := &pair.MinTick
	cur1To0 := &pair.TokenPair.CurrentTick1To0
	if *minTick == math.MaxInt64 {
		*minTick = tickIndex
		*cur1To0 = tickIndex
	} else {
		*cur1To0 = MaxInt64(*cur1To0, tickIndex)
		*minTick = MinInt64(*minTick, tickIndex)
	}

	return pair
}

func (k Keeper) UpdateTickPointersPostAddToken0(goCtx context.Context, pair *types.PairMap, tick *types.TickMap) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	newPair := k.CalcTickPointersPostAddToken0(goCtx, pair, tick)
	if newPair != nil {
		k.SetPairMap(ctx, *newPair)
	}
}

// should be called for every pair, tick for which token1 is added
func (k Keeper) CalcTickPointersPostAddToken1(goCtx context.Context, pair *types.PairMap, tick *types.TickMap) *types.PairMap {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if !k.HasToken1(ctx, tick) {
		return nil
	}

	tickIndex := tick.TickIndex
	cur0To1 := &pair.TokenPair.CurrentTick0To1
	maxTick := &pair.MaxTick
	if *maxTick == math.MinInt64 {
		*maxTick = tickIndex
		*cur0To1 = tickIndex
	} else {
		*cur0To1 = MinInt64(*cur0To1, tickIndex)
		*maxTick = MaxInt64(*maxTick, tickIndex)
	}

	return pair
}

func (k Keeper) UpdateTickPointersPostAddToken1(goCtx context.Context, pair *types.PairMap, tick *types.TickMap) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	newPair := k.CalcTickPointersPostAddToken1(goCtx, pair, tick)
	if newPair != nil {
		k.SetPairMap(ctx, *newPair)
	}
}

// Should be called for every pair, tick for which token0 liquidity is removed
func (k Keeper) CalcTickPointersPostRemoveToken0(goCtx context.Context, pair *types.PairMap, tick *types.TickMap) *types.PairMap {
	ctx := sdk.UnwrapSDKContext(goCtx)
	tickIndex := tick.TickIndex
	minTick := &pair.MinTick
	cur1To0 := &pair.TokenPair.CurrentTick1To0

	// return when we're removing liquidity between the bounds
	// Or liquidity is not drained
	if *minTick < tickIndex && tickIndex < *cur1To0 || k.HasToken0(ctx, tick) {
		//Do Nothing
		return nil
	}

	// only need to act when the token is exhausted at one of the bounds

	// We have removed all of Token0 from the pool
	if tickIndex == *minTick && tickIndex == *cur1To0 {
		*minTick = math.MaxInt64
		// we leave cur1To0 where it is because otherwise we lose the last traded price
	} else if tickIndex == *minTick {
		// TODO: We should really search for the next minTick but this introduces a
		// vulnerability unless we have a dedicated data structure for avoiding.
		*minTick++
		// we are removing liquidity below the current1To0, no need to update that
	} else if tickIndex == *cur1To0 {
		next1To0, found := k.FindNextTick1To0(goCtx, *pair)
		if !found {
			// This case should be impossible if MinTick is tracked correctly
			*minTick = math.MaxInt64
		} else {
			*cur1To0 = next1To0
		}
	}
	return pair
}

func (k Keeper) UpdateTickPointersPostRemoveToken0(goCtx context.Context, pair *types.PairMap, tick *types.TickMap) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	newPair := k.CalcTickPointersPostRemoveToken0(goCtx, pair, tick)
	if newPair != nil {
		k.SetPairMap(ctx, *newPair)
	}
}

// Should be called for every pair, tick for which token1 liquidity is removed
func (k Keeper) CalcTickPointersPostRemoveToken1(goCtx context.Context, pair *types.PairMap, tick *types.TickMap) *types.PairMap {
	ctx := sdk.UnwrapSDKContext(goCtx)
	tickIndex := tick.TickIndex
	maxTick := &pair.MaxTick
	cur0To1 := &pair.TokenPair.CurrentTick0To1

	// return when we're removing liquidity between the bounds
	// OR Liquididity is not drained
	if *cur0To1 < tickIndex && tickIndex < *maxTick || k.HasToken1(ctx, tick) {
		// Do nothing
		return nil
	}

	// only need to act when the token is exhausted at one of the bounds
	if tickIndex == *maxTick && tickIndex == *cur0To1 {
		*maxTick = math.MinInt64
		// we leave cur0To1 where it is because otherwise we lose the last traded price
	} else if tickIndex == *maxTick {
		// TODO: We should really search for the next maxTick but this introduces a
		// vulnerability unless we have a dedicated data structure for avoiding.
		*maxTick--
		// we are removing liquidity above the current0to1, no need to update that
	} else if tickIndex == *cur0To1 {
		next0To1, found := k.FindNextTick0To1(goCtx, *pair)
		if !found {
			*maxTick = math.MinInt64
			// This case should be impossible if MinTick is tracked correctly
		} else {
			*cur0To1 = next0To1
		}
	}
	return pair
}

func (k Keeper) UpdateTickPointersPostRemoveToken1(goCtx context.Context, pair *types.PairMap, tick *types.TickMap) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	newPair := k.CalcTickPointersPostRemoveToken1(goCtx, pair, tick)
	if newPair != nil {
		k.SetPairMap(ctx, *newPair)
	}
}
