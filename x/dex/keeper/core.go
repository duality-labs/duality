package keeper

import (
	"context"
	"fmt"
	"math"
	"strconv"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

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

// Handles initializing a new pair (token0/token1) if not found, adds token0, token1 to global list of tokens active on the dex
func (k Keeper) PairInit(goCtx context.Context, token0 string, token1 string, tick_index int64, fee int64, is_limit_order bool, is_0to1 bool) (string, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks to see if the token0 is contained in the tokenLsit
	token0Index, token0Found := k.GetTokenMap(ctx, token0)
	tokenLength := k.GetTokensCount(ctx)

	// If token0 is not found, add it to the list
	if !token0Found {
		k.SetTokenMap(ctx, types.TokenMap{Address: token0, Index: int64(tokenLength)})
		newTokenLength := tokenLength + 1
		token0Index.Index = int64(tokenLength)
		k.SetTokensCount(ctx, newTokenLength)
		k.AppendTokens(ctx, types.Tokens{Id: tokenLength, Address: token0})
	}

	// Checks to see if the token1 is contained in the tokenLsit
	token1Index, token1Found := k.GetTokenMap(ctx, token1)

	// If token1 is not found, add it to the list
	if !token1Found {
		k.SetTokenMap(ctx, types.TokenMap{Address: token1, Index: int64(tokenLength)})
		newTokenLength := tokenLength + 1
		token1Index.Index = int64(tokenLength)
		k.SetTokensCount(ctx, newTokenLength)
		k.AppendTokens(ctx, types.Tokens{Id: tokenLength, Address: token1})
	}

	pairId := k.CreatePairId(token0, token1)
	// Check for pair existance, if it does not exist, initialize it.
	_, PairFound := k.GetPairMap(ctx, pairId)

	if !PairFound {

		if is_limit_order {
			if is_0to1 {
				k.SetPairMap(ctx, types.PairMap{
					PairId: pairId,
					TokenPair: &types.TokenPairType{
						CurrentTick0To1: math.MaxInt64,
						CurrentTick1To0: tick_index,
					},
					MinTick: tick_index,
					MaxTick: math.MinInt64,
				})
			} else {
				k.SetPairMap(ctx, types.PairMap{
					PairId: pairId,
					TokenPair: &types.TokenPairType{
						CurrentTick0To1: tick_index,
						CurrentTick1To0: math.MinInt64,
					},
					MinTick: math.MaxInt64,
					MaxTick: tick_index,
				})
			}
		} else {
			// Initializes a new pair object in mapping.
			// Note this is only one when nno ticks currently exists, and thus we set currentTick0to1 and currentTick1to0 to be tick_index +- fee
			k.SetPairMap(ctx, types.PairMap{
				PairId: pairId,
				TokenPair: &types.TokenPairType{
					CurrentTick0To1: tick_index + fee,
					CurrentTick1To0: tick_index - fee,
				},
				MinTick: tick_index - fee,
				MaxTick: tick_index + fee,
			})
		}

	}

	// returns pairId (token0/token1)
	return pairId, nil

}

func (k Keeper) FindNextTick1To0(goCtx context.Context, pairMap types.PairMap) (found bool, tickIdx int64) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// If MinTick == MaxInt64 it is unset
	// ie. There is no Token0 in the pool
	if pairMap.MinTick == math.MaxInt64 {
		return false, math.MaxInt64
	}
	// Start scanning from CurrentTick1To0 - 1
	tickIdx = pairMap.TokenPair.CurrentTick1To0 - 1

	// Scan through all tick to the left until we hit MinTick
	for tickIdx >= pairMap.MinTick {
		// Checks for the next value tick containing amount0
		tick, tickFound := k.GetTickMap(ctx, pairMap.PairId, tickIdx)
		if tickFound && tick.TickData.HasToken0() {
			//Return the new tickIdx
			return true, tickIdx
		}

		tickIdx--
	}

	// If no tick found return false
	return false, math.MaxInt64
}

func (k Keeper) FindNextTick0To1(goCtx context.Context, pairMap types.PairMap) (found bool, tickIdx int64) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// If MaxTick == MinInt64 it is unset
	// There is no Token1 in the pool
	if pairMap.MaxTick == math.MinInt64 {
		return false, math.MinInt64
	}
	// Start scanning from CurrentTick0To1 + 1
	tickIdx = pairMap.TokenPair.CurrentTick0To1 + 1

	// Scan through all tick to the right until we hit MaxTick
	for int64(tickIdx) <= pairMap.MaxTick {
		// Checks for the next value tick containing amount1
		tick, tickFound := k.GetTickMap(ctx, pairMap.PairId, tickIdx)
		if tickFound && tick.TickData.HasToken1() {
			// Returns the new tickIdx
			return true, tickIdx
		}

		tickIdx++
	}

	// If no tick found return false
	return false, math.MinInt64
}

func (k Keeper) DepositHelper(
	goCtx context.Context,
	pairId string,
	pair types.PairMap,
	tickIndex int64,
	amount0 sdk.Dec,
	amount1 sdk.Dec,
	fee int64,
	feeIndex uint64,
) (trueAmount0 sdk.Dec, trueAmount1 sdk.Dec, sharesMinted sdk.Dec, reserve0 sdk.Dec, reserve1 sdk.Dec, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	lowerTickIndex := tickIndex - fee
	upperTickIndex := tickIndex + fee
	lowerTick, lowerTickFound := k.GetTickMap(ctx, pairId, lowerTickIndex)
	upperTick, upperTickFound := k.GetTickMap(ctx, pairId, upperTickIndex)

	// if we move the reference in the third subcondition out into a variable we lose the short-circuiting logic
	// and we get a segfault
	if !lowerTickFound || !upperTickFound || lowerTick.TickData.Reserve0AndShares[feeIndex].TotalShares.Equal(sdk.ZeroDec()) {

		// Default sets trueAmounts0/1 to amount0/1
		trueAmount0 = amount0
		trueAmount1 = amount1

		price_1to0, err := k.Calc_price_1to0(tickIndex)
		if err != nil {
			return sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec(), err
		}

		// a0 + a1 * price; this gets the shares in units of token0
		sharesMinted = amount0.Add(amount1.Mul(price_1to0))
		feeSize := k.GetFeeListCount(ctx)

		if !lowerTickFound {
			lowerTick = NewTick(pairId, lowerTickIndex, feeSize)
		}

		if !upperTickFound {
			upperTick = NewTick(pairId, upperTickIndex, feeSize)
		}

		upperTick.TickData.Reserve1[feeIndex] = trueAmount1
		lowerTick.TickData.Reserve0AndShares[feeIndex] = &types.Reserve0AndSharesType{
			Reserve0:    trueAmount0,
			TotalShares: sharesMinted,
		}
	} else {
		lowerReserve0 := lowerTick.TickData.Reserve0AndShares[feeIndex].Reserve0
		upperReserve1 := upperTick.TickData.Reserve1[feeIndex]
		lowerTotalShares := lowerTick.TickData.Reserve0AndShares[feeIndex].TotalShares

		// TODO(julian/teddy): This logic seems like it should be encapsulated
		// If feeList has been updated via a governance proposal updates future ticks to support this fee tier.
		if uint64(len(upperTick.TickData.Reserve1)) < k.GetFeeListCount(ctx) {
			upperTick.TickData.Reserve1 = append(upperTick.TickData.Reserve1, sdk.ZeroDec())
			upperTick.TickData.Reserve0AndShares = append(upperTick.TickData.Reserve0AndShares, &types.Reserve0AndSharesType{})
		}

		// TODO(julian/teddy): This logic seems like it should be encapsulated
		// If feeList has been updated via a governance proposal updates future ticks to support this fee tier.
		if uint64(len(lowerTick.TickData.Reserve1)) < k.GetFeeListCount(ctx) {
			lowerTick.TickData.Reserve1 = append(lowerTick.TickData.Reserve1, sdk.ZeroDec())
			lowerTick.TickData.Reserve0AndShares = append(lowerTick.TickData.Reserve0AndShares, &types.Reserve0AndSharesType{})
		}

		trueAmount0, trueAmount1, sharesMinted, err = CalcTrueAmounts(lowerReserve0, upperReserve1, lowerTotalShares, amount0, amount1)
		if err != nil {
			return sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec(), lowerReserve0, upperReserve1, err
		}

		// Add liquidity
		lowerTick.TickData.Reserve0AndShares[feeIndex].Reserve0 = lowerReserve0.Add(trueAmount0)
		lowerTick.TickData.Reserve0AndShares[feeIndex].TotalShares = lowerTotalShares.Add(sharesMinted)
		upperTick.TickData.Reserve1[feeIndex] = upperReserve1.Add(trueAmount1)
	}

	// If a new tick has been placed that tightens the range between currentTick1to0 and currentTick0to1 then update CurrentTicks to the tighest ticks
	if trueAmount0.GT(sdk.ZeroDec()) {
		// pair.MinTick is set to the max value to indicate that there is no liquidity for token0
		// at any tick. When this is the case, we leave CurrentTick1To0 where it is to mark the last
		// traded price, but this makes it so that when liquidity ever returns for token0, we must
		// force CurrentTick1To0 to reset wherever that liquidity is deposited, even if it is below
		// where it was before.
		if pair.MinTick == math.MaxInt64 {
			pair.TokenPair.CurrentTick1To0 = lowerTickIndex
		} else {
			pair.TokenPair.CurrentTick1To0 = MaxInt64(pair.TokenPair.CurrentTick1To0, lowerTickIndex)
		}
		pair.MinTick = MinInt64(pair.MinTick, lowerTickIndex)
	}

	if trueAmount1.GT(sdk.ZeroDec()) {
		if pair.MaxTick == math.MinInt64 {
			pair.TokenPair.CurrentTick0To1 = upperTickIndex
		} else {
			pair.TokenPair.CurrentTick0To1 = MinInt64(pair.TokenPair.CurrentTick0To1, upperTickIndex)
		}
		pair.MaxTick = MaxInt64(pair.MaxTick, upperTickIndex)
	}

	// Set pair, lower and upperTick KVStores
	k.SetPairMap(ctx, pair)
	k.SetTickMap(ctx, pairId, lowerTick)
	k.SetTickMap(ctx, pairId, upperTick)

	return trueAmount0, trueAmount1, sharesMinted, lowerTick.TickData.Reserve0AndShares[feeIndex].Reserve0, upperTick.TickData.Reserve1[feeIndex], nil

}

func MaxInt64(a, b int64) int64 {
	if a < b {
		return b
	} else {
		return a
	}
}

func MinInt64(a, b int64) int64 {
	if a < b {
		return a
	} else {
		return b
	}
}

func CalcTrueAmounts(
	lowerReserve0 sdk.Dec,
	upperReserve1 sdk.Dec,
	lowerTotalShares sdk.Dec,
	amount0 sdk.Dec,
	amount1 sdk.Dec,
) (trueAmount0 sdk.Dec, trueAmount1 sdk.Dec, sharesMinted sdk.Dec, err error) {
	// Balance trueAmount1 to the pool ratio
	if lowerReserve0.GT(sdk.ZeroDec()) {
		// trueAmount1 = min(amt1 , (upperReserve1 / lowerReserve0) * amt0 )
		trueAmount1 = Min(amount1, upperReserve1.Quo(lowerReserve0).Mul(amount0))
	} else {
		trueAmount1 = sdk.ZeroDec()
	}

	// Balance trueAmount0 to the pool ratio
	if upperReserve1.GT(sdk.ZeroDec()) {
		// trueAmount0 = min(amt0, (lowerReserve0 / upperReserve1) * amt1)
		trueAmount0 = Min(amount0, lowerReserve0.Quo(upperReserve1).Mul(amount1))
	} else {
		trueAmount0 = sdk.ZeroDec()
	}

	if trueAmount0.GT(sdk.ZeroDec()) {
		// if trueAmount0 > 0 then sharesMinted = (trueAmount0/lowerReserve0) * totalShares
		sharesMinted = (trueAmount0.Quo(lowerReserve0).Mul(lowerTotalShares))
	} else if trueAmount1.GT(sdk.ZeroDec()) {
		// else if trueAmount1 > 0 then sharesMinted = (trueAmount1/upperReserve1) * totalShares
		sharesMinted = (trueAmount1.Quo(upperReserve1).Mul(lowerTotalShares))
	} else {
		// else (both = 0) then sharesMinted = 0
		sharesMinted = sdk.ZeroDec()
	}

	if !lowerReserve0.Equal(sdk.ZeroDec()) && trueAmount0.Equal(sdk.ZeroDec()) {
		err = sdkerrors.Wrapf(types.ErrNoSpendableCoins, "Need a non-zero amount of token0 to make this deposit")
	}

	if !upperReserve1.Equal(sdk.ZeroDec()) && trueAmount1.Equal(sdk.ZeroDec()) {
		err = sdkerrors.Wrapf(types.ErrNoSpendableCoins, "Need a non-zero amount of token1 to make this deposit")
	}

	return
}

func NewTick(pairId string, tickIndex int64, numFees uint64) types.TickMap {
	tick := types.TickMap{
		PairId:    pairId,
		TickIndex: tickIndex,
		TickData: &types.TickDataType{
			Reserve0AndShares: make([]*types.Reserve0AndSharesType, numFees),
			Reserve1:          make([]sdk.Dec, numFees),
		},
		LimitOrderPool0To1: &types.LimitOrderPool{0, 0},
		LimitOrderPool1To0: &types.LimitOrderPool{0, 0},
	}
	for i, _ := range tick.TickData.Reserve0AndShares {
		tick.TickData.Reserve0AndShares[i] = &types.Reserve0AndSharesType{sdk.ZeroDec(), sdk.ZeroDec()}

	}
	for i, _ := range tick.TickData.Reserve1 {
		tick.TickData.Reserve1[i] = sdk.ZeroDec()
	}
	return tick
}

// Calculates the price for a swap from token 0 to token 1 given a tick
// tickIndex refers to the index of a specified tick
func (k Keeper) Calc_price_0to1(tick_Index int64) (sdk.Dec, error) {
	floatPrice := math.Pow(1.0001, float64(tick_Index))
	sPrice := fmt.Sprintf("%f", floatPrice)

	price, err := sdk.NewDecFromStr(sPrice)

	if err != nil {
		return sdk.ZeroDec(), err
	} else {
		//if startingToken {
		price = sdk.OneDec().Quo(price)
		//return price, nil
		return price, nil
	}
}

// Calculates the price for a swap from token 1 to token 0 given a tick
// tickIndex refers to the index of a specified tick
func (k Keeper) Calc_price_1to0(tick_Index int64) (sdk.Dec, error) {
	floatPrice := math.Pow(1.0001, float64(tick_Index))
	sPrice := fmt.Sprintf("%f", floatPrice)

	price, err := sdk.NewDecFromStr(sPrice)

	if err != nil {
		return sdk.ZeroDec(), err
	} else {
		return price, nil
	}
}

// Returns the smaller of two sdk.Decs
func Min(a, b sdk.Dec) sdk.Dec {
	if a.LT(b) {
		return a
	}
	return b
}

// Handles core logic for MsgDeposit, checking and initializing data structures (tick, pair), calculating shares based on amount deposited, and sending funds to moduleAddress
func (k Keeper) DepositCore(
	goCtx context.Context,
	msg *types.MsgDeposit,
	token0 string,
	token1 string,
	callerAddr sdk.AccAddress,
	amounts0 []sdk.Dec,
	amounts1 []sdk.Dec,
) (amounts0Deposit []sdk.Dec, amounts1Deposit []sdk.Dec, err error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// returns list of all valid fee tiers (global variable for all pairs)
	feelist := k.GetAllFeeList(ctx)

	//Checks to see if given pair has been initialied, if not intializes, and returns pairId and pairMap
	pairId, err := k.PairInit(
		goCtx,
		token0,
		token1,
		msg.TickIndexes[0],
		feelist[msg.FeeIndexes[0]].Fee,
		false,
		false,
	)

	if err != nil {
		return nil, nil, err
	}

	// returns pair object from pairObject given pairId
	pair, _ := k.GetPairMap(ctx, pairId)

	// default to 0
	totalAmountReserve0 := sdk.ZeroDec()
	totalAmountReserve1 := sdk.ZeroDec()

	passedDeposit := 0

	Amounts0Deposited := make([]sdk.Dec, len(amounts0))
	Amounts1Deposited := make([]sdk.Dec, len(amounts1))

	for i, _ := range amounts0 {
		tickIndex := msg.TickIndexes[i]
		feeIndex := msg.FeeIndexes[i]
		fee := feelist[feeIndex].Fee
		curTick0to1 := pair.TokenPair.CurrentTick0To1
		curTick1to0 := pair.TokenPair.CurrentTick1To0
		lowerTickIndex := tickIndex - fee
		upperTickIndex := tickIndex + fee
		// TODO: Allow user to deposit "behind enemy lines"
		if amounts0[i].GT(sdk.ZeroDec()) && curTick0to1 <= lowerTickIndex {
			return nil, nil, sdkerrors.Wrapf(types.ErrValidPairNotFound, "Cannot depsosit amount_0 at a tick less than the CurrentTick0to1")
		}

		// TODO: Allow user to deposit "behind enemy lines"
		if amounts1[i].GT(sdk.ZeroDec()) && upperTickIndex <= curTick1to0 {
			return nil, nil, sdkerrors.Wrapf(types.ErrValidPairNotFound, "Cannot deposit amount_1 at a tick greater than the CurrentTick1to0")
		}

		// Calls k.DepositHelper which calculates the true amounts of token0, token1, and sets the corresponding pair and tick maps
		// returns trueAmounts0, trueAmount1 the actual amount able to be deposited at the specific tick, and the corresponding sharesMinted
		// Note: newReserve0, newReserve1 are returned for use in an event
		trueAmount0, trueAmount1, sharesMinted, newReserve0, newReserve1, err := k.DepositHelper(goCtx, pairId, pair, msg.TickIndexes[i], amounts0[i], amounts1[i], feelist[msg.FeeIndexes[i]].Fee, msg.FeeIndexes[i])

		if err != nil {
			return nil, nil, err
		}

		Amounts0Deposited[i] = trueAmount0
		Amounts1Deposited[i] = trueAmount1

		// if trueAmount0 == 0 && trueAmount1 == 0 then newReserves are unchanged, So emit partial fail event.
		if trueAmount0.Equal(sdk.ZeroDec()) && trueAmount1.Equal(sdk.ZeroDec()) {
			ctx.EventManager().EmitEvent(types.CreateDepositFailedEvent(msg.Creator, msg.Receiver,
				token0, token1, fmt.Sprint(msg.TickIndexes[i]), fmt.Sprint(msg.FeeIndexes[i]),
				newReserve0.String(), newReserve1.String(), amounts0[i].String(), amounts1[i].String()),
			)
			continue
		}

		// else, count as successful deposit
		passedDeposit++

		// Retreives receiver's share object for the specified tick fee pair
		shares, sharesFound := k.GetShares(ctx, msg.Receiver, pairId, msg.TickIndexes[i], msg.FeeIndexes[i])
		// Initializes a new share object if sharesFound does not exists
		if !sharesFound {
			shares = types.Shares{
				Address:     msg.Receiver,
				PairId:      pairId,
				TickIndex:   msg.TickIndexes[i],
				FeeIndex:    msg.FeeIndexes[i],
				SharesOwned: sharesMinted,
			}
		} else {
			shares.SharesOwned = shares.SharesOwned.Add(sharesMinted)
		}

		// Update share logic to KVStore
		k.SetShares(ctx, shares)

		// adds trueAmounts0 and trueAmounts1 to the total amount of token0/1 deposited so far
		totalAmountReserve0 = totalAmountReserve0.Add(trueAmount0)
		totalAmountReserve1 = totalAmountReserve1.Add(trueAmount1)

		// emit successful deposit event
		ctx.EventManager().EmitEvent(types.CreateDepositEvent(msg.Creator, msg.Receiver,
			token0, token1, fmt.Sprint(msg.TickIndexes[i]), fmt.Sprint(msg.FeeIndexes[i]),
			newReserve0.Sub(trueAmount0).String(), newReserve1.Sub(trueAmount1).String(), newReserve0.String(), newReserve1.String(),
			sharesMinted.String()),
		)
	}

	if passedDeposit == 0 {
		return nil, nil, sdkerrors.Wrapf(types.ErrAllDepositsFailed, "All deposits failed")
	}

	// Send TrueAmount0 to Module
	/// @Dev Due to a sdk.send constraint this only sends if trueAmount0 is greater than 0
	if totalAmountReserve0.GT(sdk.ZeroDec()) {
		coin0 := sdk.NewCoin(token0, sdk.NewIntFromBigInt(totalAmountReserve0.BigInt()))
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin0}); err != nil {
			return nil, nil, err
		}

	}

	// Send TrueAmount1 to Module
	/// @Dev Due to a sdk.send constraint this only sends if trueAmount1 is greater than 0
	if totalAmountReserve1.GT(sdk.ZeroDec()) {
		coin1 := sdk.NewCoin(token1, sdk.NewIntFromBigInt(totalAmountReserve1.BigInt()))
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin1}); err != nil {
			return nil, nil, err
		}
	}

	_ = goCtx
	return Amounts0Deposited, Amounts1Deposited, nil
}

// Handles core logic for MsgWithdrawl; calculating and withdrawing reserve0,reserve1 from a specified tick given a specfied number of shares to remove.
// Calculates the amount of reserve0, reserve1 to withdraw based on the percetange of the desired number of shares to remove compared to the total number of shares at the given tick
func (k Keeper) WithdrawCore(goCtx context.Context, msg *types.MsgWithdrawl, token0 string, token1 string, callerAddr sdk.AccAddress, receiverAddr sdk.AccAddress) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//loads pairId for the gven token0/token1 pair
	pairId := k.CreatePairId(token0, token1)

	// loads pair from mapping
	pair, pairFound := k.GetPairMap(ctx, pairId)

	// if a pair does not exist, neither do any ticks, and thus a withdraw can not be made
	if !pairFound {
		return sdkerrors.Wrapf(types.ErrValidPairNotFound, "Pair not found")
	}

	//init totalamountToWIthdraw of both assets to zero
	totalReserve0ToRemove := sdk.ZeroDec()
	totalReserve1ToRemove := sdk.ZeroDec()

	//Iterate through the list of all sharesToRemove, calculating the amounttoWithdraw for specified each tick
	for i, feeIndex := range msg.FeeIndexes {
		sharesToRemove := msg.SharesToRemove[i]
		tickIndex := msg.TickIndexes[i]

		// loads UserShares type from mapping
		shareOwner, shareOwnerFound := k.GetShares(
			ctx,
			msg.Creator,
			pairId,
			tickIndex,
			feeIndex,
		)
		if !shareOwnerFound {
			return sdkerrors.Wrapf(types.ErrValidShareNotFound, "No valid share owner fonnd")
		}

		feeValue, _ := k.GetFeeList(ctx, feeIndex)
		fee := feeValue.Fee

		lowerTickIndex := tickIndex - fee
		upperTickIndex := tickIndex + fee
		lowerTick, lowerTickFound := k.GetTickMap(ctx, pairId, upperTickIndex)
		upperTick, upperTickFound := k.GetTickMap(ctx, pairId, lowerTickIndex)
		if !lowerTickFound || !upperTickFound {
			return sdkerrors.Wrapf(types.ErrValidTickNotFound, "No tick found at the requested index")
		}

		lowerTickFeeTotalShares := &lowerTick.TickData.Reserve0AndShares[feeIndex].TotalShares
		lowerTickFeeReserve0 := &lowerTick.TickData.Reserve0AndShares[feeIndex].Reserve0
		upperTickFeeReserve1 := &upperTick.TickData.Reserve1[feeIndex]

		// Checks to see if there are some totalShares to withdraw
		// In keeper/verification.go we check this condition for the msg.Creator, thus we know that they also has a valid position in the tick.
		if lowerTick.TickData.Reserve0AndShares[msg.FeeIndexes[i]].TotalShares.Equal(sdk.ZeroDec()) {
			return sdkerrors.Wrapf(types.ErrValidTickNotFound, "No tick found at the requested index")
		}

		// calculates the amount to withdraw of each token based on a ratio of the amountToRemove to totalShares multiplied by the amount of the respective asset
		reserve0ToRemove := lowerTickFeeReserve0.Mul(sharesToRemove.Quo(*lowerTickFeeTotalShares))
		reserve1ToRemove := (sharesToRemove.Quo(*lowerTickFeeTotalShares)).Mul(*upperTickFeeReserve1)

		//Updates upper/lowerTick based on subtracting the calculated amount from the previous reserve0 and reserve1
		*lowerTickFeeReserve0 = lowerTickFeeReserve0.Sub(reserve0ToRemove)
		*upperTickFeeReserve1 = upperTickFeeReserve1.Sub(reserve1ToRemove)
		*lowerTickFeeTotalShares = lowerTickFeeTotalShares.Sub(msg.SharesToRemove[i])

		// subtracts sahresToRemove from the User's personl number of sharesOwned.
		shareOwner.SharesOwned = shareOwner.SharesOwned.Sub(msg.SharesToRemove[i])

		if !lowerTick.TickData.HasToken0() {
			if pair.TokenPair.CurrentTick1To0 == lowerTickIndex {
				found, tickIdx := k.FindNextTick1To0(goCtx, pair)
				if found {
					pair.TokenPair.CurrentTick1To0 = tickIdx
				} else {
					// We leave the tick where it is so the state reflects the last price
					// before liquidity was exhausted.
				}
			}

			if pair.MinTick == lowerTickIndex {
				// This opens up a vulnerability where someone forces the
				// chain to iterate down to a low minTick if they exhaust all
				// other liquidity, causing the chain to hang.
				pair.MinTick++
			}
		}

		if !upperTick.TickData.HasToken1() {
			if pair.TokenPair.CurrentTick0To1 == upperTickIndex {
				found, tickIdx := k.FindNextTick0To1(goCtx, pair)
				if found {
					pair.TokenPair.CurrentTick0To1 = tickIdx
				} else {
					// We leave the tick where it is so the state reflects the last price
					// before liquidity was exhausted.
				}
			}

			if upperTickIndex == pair.MaxTick {
				// This opens up a vulnerability where someone forces the
				// chain to iterate up to a high maxTick if they exhaust all
				// other liquidity, causing the chain to hang.
				pair.MaxTick--
			}
		}

		// adds reserve0ToRemove/reserve1ToRemove to totals
		totalReserve0ToRemove = totalReserve0ToRemove.Add(reserve0ToRemove)
		totalReserve1ToRemove = totalReserve1ToRemove.Add(reserve1ToRemove)

		// sets changes to tick mappings, and share mappings
		k.SetShares(ctx, shareOwner)
		k.SetTickMap(ctx, pairId, upperTick)
		k.SetTickMap(ctx, pairId, lowerTick)

		// emits event for individiual withdrawl
		ctx.EventManager().EmitEvent(types.CreateWithdrawEvent(
			msg.Creator,
			msg.Receiver,
			token0,
			token1,
			fmt.Sprint(msg.TickIndexes),
			fmt.Sprint(msg.FeeIndexes),
			lowerTickFeeReserve0.Add(reserve0ToRemove).String(),
			upperTickFeeReserve1.Add(reserve1ToRemove).String(),
			lowerTickFeeReserve0.String(),
			upperTickFeeReserve1.String(),
		))
	}

	//Sets changes to pair mapping
	k.SetPairMap(ctx, pair)

	// sends totalReserve0ToRemove to msg.Receiver
	if totalReserve0ToRemove.GT(sdk.ZeroDec()) {
		coin0 := sdk.NewCoin(token0, sdk.NewIntFromBigInt(totalReserve0ToRemove.BigInt()))
		err := k.bankKeeper.SendCoinsFromModuleToAccount(
			ctx,
			types.ModuleName,
			receiverAddr,
			sdk.Coins{coin0},
		)
		if err != nil {
			return err
		}
	}

	// sends totalReserve1ToRemove to msg.Receiver
	if totalReserve1ToRemove.GT(sdk.ZeroDec()) {
		coin1 := sdk.NewCoin(token1, sdk.NewIntFromBigInt(totalReserve1ToRemove.BigInt()))
		err := k.bankKeeper.SendCoinsFromModuleToAccount(
			ctx,
			types.ModuleName,
			receiverAddr,
			sdk.Coins{coin1},
		)
		if err != nil {
			return err
		}
	}

	_ = ctx
	return nil
}

////// Swap Functions

// Handles core logic for the asset 0 to asset1 direction of MsgSwap; faciliates swapping amount0 for some amount of amount1, given a specified pair (token0, token1)

func (k Keeper) Swap0to1(goCtx context.Context, msg *types.MsgSwap, token0 string, token1 string, callerAddr sdk.AccAddress) (sdk.Dec, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// pair idea: "token0/token1"
	pairId := k.CreatePairId(token0, token1)

	// size of the feeList
	feeSize := k.GetFeeListCount(ctx)
	feelist := k.GetAllFeeList(ctx)
	// gets the PairMap from the KVstore given pairId
	pair, pairFound := k.GetPairMap(ctx, pairId)

	// If toknePair does not exists, a swap cannot be made through it; error
	if !pairFound {
		return sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrValidPairNotFound, "Pair not found")
	}

	//amount_left is the amount left to deposit
	amount_left := msg.AmountIn

	// amount to return to receiver
	amount_out := sdk.ZeroDec()

	// verify that amount left is not zero and that there are additional valid ticks to check
	for !amount_left.Equal(sdk.ZeroDec()) && pair.TokenPair.CurrentTick0To1 <= pair.MaxTick {

		// Tick data for tick that holds information about reserve1
		Current1Data, Current1Found := k.GetTickMap(ctx, pairId, pair.TokenPair.CurrentTick0To1)

		// If a tick at Current0to1 is not found, decrement CurrentTick0to1 (to the next tick that is supposed to contain reserve1) and check again
		if !Current1Found {
			pair.TokenPair.CurrentTick0To1 = pair.TokenPair.CurrentTick0To1 + 1
			continue
		}

		var i uint64

		// iterator for feeList
		i = 0
		for i < feeSize && !amount_left.Equal(sdk.ZeroDec()) {
			// gets fee for given feeIndex
			fee := feelist[i].Fee

			// @dev CurrentTick0to1 - 2 * fee finds the respective tickPair (containing totalShares, reserve0)
			Current0Data, Current0Found := k.GetTickMap(ctx, pairId, pair.TokenPair.CurrentTick0To1-2*fee)
			//Current0Datam := Current0Data.TickData.Reserve1[i]

			// If tick/feeIndex pair is not found continue
			if !Current0Found {
				i++
				continue
			}
			// calculate currentPrice
			price_0to1, err := k.Calc_price_0to1(pair.TokenPair.CurrentTick0To1)

			if err != nil {
				return sdk.ZeroDec(), err
			}

			// price * amout_left + amount_out < minOut, error we cannot meet minOut threshold
			if price_0to1.Mul(amount_left).Add(amount_out).LT(msg.MinOut) {
				return sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Amount Out is less than minium amount out specified: swap failed")
			}

			// If the amount of reserves is not enough to finish the swap
			// R1  < amount_left * p0to1
			if Current1Data.TickData.Reserve1[i].LT(amount_left.Mul(price_0to1)) {
				// amount_out += r1 (adds as all of reserve1 to amount_out)
				amount_out = amount_out.Add(Current1Data.TickData.Reserve1[i])

				// AmountOut = reserves1 = amountInTemp * price0to1
				// => amountInTemp = reserves1 / price0to1
				amountInTemp := Current1Data.TickData.Reserve1[i].Quo(price_0to1)
				// decrement amount_left by amountInTemp
				amount_left = amount_left.Sub(amountInTemp)
				//updates reserve0 with the new amountInTemp
				Current0Data.TickData.Reserve0AndShares[i].Reserve0 = Current0Data.TickData.Reserve0AndShares[i].Reserve0.Add(amountInTemp)
				// sets reserve1 to 0
				Current1Data.TickData.Reserve1[i] = sdk.ZeroDec()

			} else {
				amountOutTemp := amount_left.Mul(price_0to1)
				// amountOut += amount_left * price
				amount_out = amount_out.Add(amountOutTemp)
				// increment reserve0 with amountLeft
				Current0Data.TickData.Reserve0AndShares[i].Reserve0 = Current0Data.TickData.Reserve0AndShares[i].Reserve0.Add(amount_left)
				// decrement reserve1 with amount_left * price
				Current1Data.TickData.Reserve1[i] = Current1Data.TickData.Reserve1[i].Sub(amountOutTemp)
				// set amountLeft to 0
				amount_left = sdk.ZeroDec()
			}

			//updates feeIndex
			i++

			//Make updates to tickMap containing reserve0/1 data to the KVStore
			k.SetTickMap(ctx, pairId, Current0Data)

			pair.MinTick = MinInt64(Current0Data.TickIndex, pair.MinTick)
			pair.TokenPair.CurrentTick1To0 = MaxInt64(Current0Data.TickIndex, pair.TokenPair.CurrentTick1To0)
		}

		k.SetTickMap(ctx, pairId, Current1Data)

		// if feeIndex is equal to the largest index in feeList check for valid limit orders
		if i == feeSize {

			// assigns a new variable err to handle err not being initialized in this conditional
			var err error

			// runs swaps for any limitOrders at the specified tick, updating amount_left, amount_out accordingly
			// passes in the outToken (token1), as this is the direction of the limit order for which we check

			amount_left, amount_out, err = k.SwapLimitOrder1to0(goCtx, pairId, token1, amount_out, amount_left, pair.TokenPair.CurrentTick0To1)

			if err != nil {
				return sdk.ZeroDec(), err
			}

			// decrements CurrentTick0to1 to the next tick containing reserve1
			pair.TokenPair.CurrentTick0To1 = pair.TokenPair.CurrentTick0To1 + 1
		}

		// If we drain Current1 and it was MaxTick we have exhausted all Token1 in the pair
		// Therefore we need to unset MaxTick ie. (ie MaxTick = MinInt64)
		if !Current1Data.TickData.HasToken1() && Current1Data.TickIndex == pair.MaxTick {
			pair.MaxTick = math.MinInt64
		}
	}

	// Make updates to pairMap containing updates CurrentTick0to1 to the KVStore
	k.SetPairMap(ctx, pair)

	// Check to see if amount_out meets the threshold of minOut
	if amount_out.LT(msg.MinOut) {
		return sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Amount Out is less than minium amount out specified: swap failed")
	}

	ctx.EventManager().EmitEvent(types.CreateSwapEvent(msg.Creator, msg.Receiver,
		token0, token1, msg.TokenIn, msg.AmountIn.String(), amount_out.String(), msg.MinOut.String(),
	))

	// Returns amount_out to keeper/msg.server: Swap
	// @Dev token transfers happen in keeper/msg.server: Swap
	return amount_out, nil
}

// Handles core logic for the asset 1 to asset 0 direction of MsgSwap; faciliates swapping amount1 for some amount of amount0, given a specified pair (token0, token1)
func (k Keeper) Swap1to0(goCtx context.Context, msg *types.MsgSwap, token0 string, token1 string, callerAddr sdk.AccAddress) (sdk.Dec, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// pair idea: "token0/token1"
	pairId := k.CreatePairId(token0, token1)

	// size of the feeList
	feeSize := k.GetFeeListCount(ctx)
	feelist := k.GetAllFeeList(ctx)
	// geets the PairMap from the KVstore given pairId
	pair, pairFound := k.GetPairMap(ctx, pairId)

	// If toknePair does not exists, a swap cannot be made through it; error
	if !pairFound {
		return sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrValidPairNotFound, "Pair not found")
	}

	//amount_left is the amount left to deposit
	amount_left := msg.AmountIn

	// amount to return to receiver
	amount_out := sdk.ZeroDec()

	// verify that amount left is not zero and that there are additional valid ticks to check

	for !amount_left.Equal(sdk.ZeroDec()) && pair.TokenPair.CurrentTick1To0 >= pair.MinTick {

		Current0Data, Current0Found := k.GetTickMap(ctx, pairId, pair.TokenPair.CurrentTick1To0)
		//Current0Datam := Current0Data.TickData.Reserve1[i]

		// If tick/feeIndex pair is not found continue

		// If a tick at Current1to0 is not found, incremenet CurrentTick1to0 (to the next tick that is supposed to contain reserve0) and check again
		if !Current0Found {
			pair.TokenPair.CurrentTick1To0 = pair.TokenPair.CurrentTick1To0 - 1
			continue
		}

		var i uint64

		// iterator for feeList
		i = 0
		for i < feeSize && !amount_left.Equal(sdk.ZeroDec()) {
			// gets fee for given feeIndex
			fee := feelist[i].Fee

			// @dev CurrentTick1to0 - 2 * fee finds the respective tickPair (reserve1)
			Current1Data, Current1Found := k.GetTickMap(ctx, pairId, pair.TokenPair.CurrentTick1To0+2*fee)

			if !Current1Found {
				i++
				continue
			}
			//Current0Datam := Current0Data.TickData.Reserve1[i]

			// calculate currentPrice and inverts
			price_1to0, err := k.Calc_price_1to0(pair.TokenPair.CurrentTick1To0)

			if err != nil {
				return sdk.ZeroDec(), err
			}

			// price * amout_left + amount_out < minOut, error we cannot meet minOut threshold
			if price_1to0.Mul(amount_left).Add(amount_out).LT(msg.MinOut) {
				return sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Amount Out is less than minium amount out specified: swap failed")
			}

			// If there is not enough to complete the trade
			if Current0Data.TickData.Reserve0AndShares[i].Reserve0.LT(amount_left.Mul(price_1to0)) {
				// Add the reserves to the amount out
				amount_out = amount_out.Add(Current0Data.TickData.Reserve0AndShares[i].Reserve0)

				amountInTemp := Current0Data.TickData.Reserve0AndShares[i].Reserve0.Quo(price_1to0)

				amount_left = amount_left.Sub(amountInTemp)

				Current1Data.TickData.Reserve1[i] = Current0Data.TickData.Reserve0AndShares[i].Reserve0.Add(amountInTemp)

				Current0Data.TickData.Reserve0AndShares[i].Reserve0 = sdk.ZeroDec()

			} else {
				amountOutTemp := amount_left.Mul(price_1to0)

				amount_out = amount_out.Add(amountOutTemp)

				Current1Data.TickData.Reserve1[i] = Current0Data.TickData.Reserve0AndShares[i].Reserve0.Add(amount_left)

				Current0Data.TickData.Reserve0AndShares[i].Reserve0 = Current1Data.TickData.Reserve1[i].Sub(amountOutTemp)

				amount_left = sdk.ZeroDec()
			}

			//updates feeIndex
			i++

			//Make updates to tickMap containing reserve0/1 data to the KVStore

			k.SetTickMap(ctx, pairId, Current1Data)

			pair.MaxTick = MaxInt64(Current1Data.TickIndex, pair.MaxTick)
			pair.TokenPair.CurrentTick0To1 = MinInt64(Current0Data.TickIndex, pair.TokenPair.CurrentTick0To1)

		}

		k.SetTickMap(ctx, pairId, Current0Data)
		// if feeIndex is equal to the largest index in feeList, check for valid limit orders at the specfied tick
		if i == feeSize {

			// assigns a new variable err to handle err not being initialized in this conditional
			var err error
			// runs swaps for any limitOrders at the specified tick, updating amount_left, amount_out accordingly

			// passes in the outToken (token0), as this is the direction of the limit order for which we check
			amount_left, amount_out, err = k.SwapLimitOrder1to0(goCtx, pairId, token0, amount_out, amount_left, pair.TokenPair.CurrentTick1To0)

			if err != nil {
				return sdk.ZeroDec(), err
			}

			pair.TokenPair.CurrentTick1To0 = pair.TokenPair.CurrentTick1To0 - 1
		}

		// If we drain Current0 and it was MinTick we have exhausted all Token0 in the pair
		// Therefore we need to unset MinTick ie. (ie MinTick = MaxInt64)
		if !Current0Data.TickData.HasToken0() && Current0Data.TickIndex == pair.MinTick {
			pair.MinTick = math.MaxInt64
		}
	}

	// Check to see if amount_out meets the threshold of minOut
	k.SetPairMap(ctx, pair)

	if amount_out.LT(msg.MinOut) {
		return sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Amount Out is less than minium amount out specified: swap failed")
	}

	ctx.EventManager().EmitEvent(types.CreateSwapEvent(msg.Creator, msg.Receiver,
		token0, token1, msg.TokenIn, msg.AmountIn.String(), amount_out.String(), msg.MinOut.String(),
	))

	// Returns amount_out to keeper/msg.server: Swap
	// @Dev token transfers happen in keeper/msg.server: Swap
	return amount_out, nil
}

// Swap Limit Orders

// Handles swapping asset 0 for asset 1 through any active limit orders at a specified tick
// Returns amount_out, amount_left, error
func (k Keeper) SwapLimitOrder0to1(goCtx context.Context, pairId string, tokenIn string, amount_out sdk.Dec, amount_left sdk.Dec, CurrentTick0to1 int64) (sdk.Dec, sdk.Dec, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// returns price for the given tick and specified direction (0 -> 1)
	price_0to1, err := k.Calc_price_0to1(CurrentTick0to1)

	if err != nil {
		return sdk.ZeroDec(), sdk.ZeroDec(), err
	}

	// Gets tick for specified tick at currentTick0to1
	tick, tickFound := k.GetTickMap(ctx, pairId, CurrentTick0to1)

	// Edge case: if tick specified by CurrentTick1to0 is not found, there does not exists any valid limit orders by construction, and thus return the same amount_left, amount_out as inputted
	if !tickFound {
		return amount_left, amount_out, nil
	}

	// Gets Reserve, FilledReservesmap for the specified CurrentLimitOrderKey
	ReserveData, ReserveDataFound := k.GetLimitOrderPoolReserveMap(ctx, pairId, CurrentTick0to1, tokenIn, tick.LimitOrderPool1To0.CurrentLimitOrderKey)
	FillData, _ := k.GetLimitOrderPoolFillMap(ctx, pairId, CurrentTick0to1, tokenIn, tick.LimitOrderPool1To0.CurrentLimitOrderKey)

	// errors if ReserveDataFound is not found
	if !ReserveDataFound {
		return amount_left, amount_out, nil
	}

	// If the amount of reserves is not enough to finish the swap

	// If there isn't enough liqudity to end trade handle updates this way
	// R1  < amount_left * p0to1
	if ReserveData.Reserves.LT(amount_left.Mul(price_0to1)) {
		// Adds remaining reserves to amount_out
		amount_out = amount_out.Add(ReserveData.Reserves)

		amountInTemp := ReserveData.Reserves.Quo(price_0to1)

		// Subtracts reserves from amount_left
		amount_left = amount_left.Sub(amountInTemp)
		// adds price * reserves to the filledMap
		FillData.FilledReserves = FillData.FilledReserves.Add(amountInTemp)
		// sets reserves to 0
		ReserveData.Reserves = sdk.ZeroDec()

		// increments the limitOrderkey as previous tick has been completely filled
		tick.LimitOrderPool1To0.CurrentLimitOrderKey++

		// checks the next currentLimitOrderKey
		ReserveDataNextKey, ReserveDataNextKeyFound := k.GetLimitOrderPoolReserveMap(ctx, pairId, CurrentTick0to1, tokenIn, tick.LimitOrderPool1To0.CurrentLimitOrderKey)
		FillDataNextKey, FillDataNextKeyFound := k.GetLimitOrderPoolFillMap(ctx, pairId, CurrentTick0to1, tokenIn, tick.LimitOrderPool1To0.CurrentLimitOrderKey)

		// if no tokens have been filled at this key value, initialize to 0
		if !FillDataNextKeyFound {
			FillDataNextKey.Count = tick.LimitOrderPool1To0.CurrentLimitOrderKey
			FillDataNextKey.TickIndex = CurrentTick0to1
			FillDataNextKey.PairId = pairId
			FillDataNextKey.FilledReserves = sdk.ZeroDec()
		}

		// If there is still not enough liquidity to end trade handle update this way
		if ReserveDataNextKeyFound && price_0to1.Mul(ReserveDataNextKey.Reserves).LT(amount_left) {
			// Adds remaining reserves to amount_out
			amount_out = amount_out.Add(ReserveDataNextKey.Reserves)

			amountInTemp := ReserveDataNextKey.Reserves.Quo(price_0to1)
			// Subtracts reserves from amount_left
			amount_left = amount_left.Sub(amountInTemp)
			// adds price * reserves to the filledMap
			FillDataNextKey.FilledReserves = FillDataNextKey.FilledReserves.Add(amountInTemp)
			// sets reserve to 0
			ReserveDataNextKey.Reserves = sdk.ZeroDec()

			// increments the limitOrderKey
			tick.LimitOrderPool1To0.CurrentLimitOrderKey++

			// If there IS enough liqudity to end trade handle update this way
		} else if ReserveDataNextKeyFound {
			amountOutTemp := amount_left.Mul(price_0to1)
			// calculate anmout to output (will be a portion of reserves)
			amount_out = amount_out.Add(amountOutTemp)
			// Add the amount_left to the amount flled in the filledReservesmapping
			FillDataNextKey.FilledReserves = FillDataNextKey.FilledReserves.Add(amount_left)
			// subtract amount_left * price to the ReserveMapping
			ReserveDataNextKey.Reserves = ReserveDataNextKey.Reserves.Sub(amountOutTemp)
			// set amount_left to 0
			amount_left = sdk.ZeroDec()
		}

		// Updates mapping for the original limit order key + 1 (next key)
		// @dev we set mappings within the conditionnal, as the tick mappings have not been initialized outside of it
		k.SetLimitOrderPoolFillMap(ctx, FillDataNextKey)
		k.SetLimitOrderPoolReserveMap(ctx, ReserveDataNextKey)

		// If there IS enough liqudity to end trade handle update this way
	} else {
		// calculate anmout to output (will be a portion of reserves)
		amount_out = amount_out.Add(amount_left.Mul(price_0to1))
		// Add the amount_left to the amount flled in the filledReservesmapping
		FillData.FilledReserves = FillData.FilledReserves.Add(amount_left)
		// subtract amount_left * price to the ReserveMapping
		ReserveData.Reserves = ReserveData.Reserves.Sub(amount_left.Mul(price_0to1))
		// set amount_left to 0
		amount_left = sdk.ZeroDec()
	}

	// Updates mappings of reserve and filledReserves based on the original limitOrderCurrentKey to the KVStore
	k.SetLimitOrderPoolReserveMap(ctx, ReserveData)
	k.SetLimitOrderPoolFillMap(ctx, FillData)

	//Updates limitOrderCurrentKey based on if any limitOrders were completely filled.
	k.SetTickMap(ctx, pairId, tick)

	_ = ctx

	return amount_left, amount_out, nil
}

// Handles swapping asset 1 for asset 0 through any active limit orders at a specified tick
// Returns amount_out, amount_left, error

func (k Keeper) SwapLimitOrder1to0(goCtx context.Context, pairId string, tokenIn string, amount_out sdk.Dec, amount_left sdk.Dec, CurrentTick1to0 int64) (sdk.Dec, sdk.Dec, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// returns price for the given tick and specified direction (0 -> 1)
	price_1to0, err := k.Calc_price_1to0(CurrentTick1to0)

	if err != nil {
		return sdk.ZeroDec(), sdk.ZeroDec(), err
	}

	tick, tickFound := k.GetTickMap(ctx, pairId, CurrentTick1to0)

	// Edge case: if tick specified by CurrentTick1to0 is not found, there does not exists any valid limit orders by construction, and thus return the same amount_left, amount_out as inputted
	if !tickFound {
		return amount_left, amount_out, nil
	}

	ReserveData, ReserveDataFound := k.GetLimitOrderPoolReserveMap(ctx, pairId, CurrentTick1to0, tokenIn, tick.LimitOrderPool0To1.CurrentLimitOrderKey)
	FillData, _ := k.GetLimitOrderPoolFillMap(ctx, pairId, CurrentTick1to0, tokenIn, tick.LimitOrderPool0To1.CurrentLimitOrderKey)

	// errors if ReserveDataFound is not found
	if !ReserveDataFound {
		return amount_left, amount_out, nil
	}

	// If there isn't enough liqudity to end trade handle updates this way
	// R1  < amount_left * p0to1
	// amountOut = R1 = amountInTemp * p1to0
	// => amountInTemp = R1 / p1to0
	if ReserveData.Reserves.LT(amount_left.Mul(price_1to0)) {
		amountInTemp := ReserveData.Reserves.Quo(price_1to0)
		// Adds remaining reserves to amount_out
		amount_out = amount_out.Add(ReserveData.Reserves)
		// Subtracts reserves from amount_left
		amount_left = amount_left.Sub(amountInTemp)
		// adds price * reserves to the filledMap
		FillData.FilledReserves = FillData.FilledReserves.Add(amountInTemp)
		// sets reserves to 0
		ReserveData.Reserves = sdk.ZeroDec()

		// increments the limitOrderkey as previous tick has been completely filled
		tick.LimitOrderPool0To1.CurrentLimitOrderKey++

		// checks the next currentLimitOrderKey
		ReserveDataNextKey, ReserveDataNextKeyFound := k.GetLimitOrderPoolReserveMap(ctx, pairId, CurrentTick1to0, tokenIn, tick.LimitOrderPool0To1.CurrentLimitOrderKey)
		FillDataNextKey, FillMapNextKeyFound := k.GetLimitOrderPoolFillMap(ctx, pairId, CurrentTick1to0, tokenIn, tick.LimitOrderPool0To1.CurrentLimitOrderKey)

		// if no tokens have been filled at this key value, initialize to 0
		if !FillMapNextKeyFound {
			FillDataNextKey.Count = tick.LimitOrderPool0To1.CurrentLimitOrderKey
			FillDataNextKey.TickIndex = CurrentTick1to0
			FillDataNextKey.PairId = pairId
			FillDataNextKey.FilledReserves = sdk.ZeroDec()
		}

		if ReserveDataNextKeyFound && ReserveDataNextKey.Reserves.LT(amount_left.Mul(price_1to0)) {
			// Adds remaining reserves to amount_out
			amountInTemp := ReserveDataNextKey.Reserves.Quo(price_1to0)
			amount_out = amount_out.Add(ReserveDataNextKey.Reserves)
			// Subtracts reserves from amount_left
			amount_left = amount_left.Sub(amountInTemp)
			// adds price * reserves to the filledMap
			FillDataNextKey.FilledReserves = FillDataNextKey.FilledReserves.Add(amountInTemp)
			// sets reserve to 0
			ReserveDataNextKey.Reserves = sdk.ZeroDec()

			// increments the limitOrderKey
			tick.LimitOrderPool0To1.CurrentLimitOrderKey++

		} else if ReserveDataNextKeyFound {
			amountOutTemp := amount_left.Mul(price_1to0)
			// calculate anmout to output (will be a portion of reserves)
			amount_out = amount_out.Add(amountOutTemp)
			// Add the amount_left to the amount flled in the filledReservesmapping
			FillDataNextKey.FilledReserves = FillDataNextKey.FilledReserves.Add(amount_left)
			// subtract amount_left * price to the ReserveMapping
			ReserveDataNextKey.Reserves = ReserveDataNextKey.Reserves.Sub(amount_left.Mul(price_1to0))
			// set amount_left to 0
			amount_left = sdk.ZeroDec()
		}

		// Updates mapping for the original limit order key + 1 (next key)
		// @dev we set mappings within the conditionnal, as the tick mappings have not been initialized outside of it
		k.SetLimitOrderPoolFillMap(ctx, FillDataNextKey)
		k.SetLimitOrderPoolReserveMap(ctx, ReserveDataNextKey)

		// If there IS enough liqudity to end trade handle update this way
	} else {
		amountOutTemp := amount_left.Mul(price_1to0)
		// calculate anmout to output (will be a portion of reserves)
		amount_out = amount_out.Add(amountOutTemp)
		// Add the amount_left to the amount flled in the filledReservesmapping
		FillData.FilledReserves = FillData.FilledReserves.Add(amount_left)
		// subtract amount_left * price to the ReserveMapping
		ReserveData.Reserves = ReserveData.Reserves.Sub(amountOutTemp)
		// set amount_left to 0
		amount_left = sdk.ZeroDec()
	}

	// Updates mappings of reserve and filledReserves based on the original limitOrderCurrentKey to the KVStore
	k.SetLimitOrderPoolReserveMap(ctx, ReserveData)
	k.SetLimitOrderPoolFillMap(ctx, FillData)

	//Updates limitOrderCurrentKey based on if any limitOrders were completely filled.
	k.SetTickMap(ctx, pairId, tick)

	_ = ctx

	return amount_left, amount_out, nil
}

///// Limit Order Functions

// Helper Function for PlaceLimitOrder handling initializing a new tick
func (k Keeper) PlaceLimitOrderHelper(goCtx context.Context, pairId string, tickIndex int64) types.TickMap {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// returns tick object for tick specified by tickIndex
	tick, tickFound := k.GetTickMap(ctx, pairId, tickIndex)

	// size of the feeList
	feeSize := k.GetFeeListCount(ctx)

	if !tickFound {

		// If tick does not exists intialize it
		// @Dev initialize reserves struct to avoid having to check for individual subtype existance between deposit and placeLimitOrder
		tick = types.TickMap{
			PairId:    pairId,
			TickIndex: tickIndex,
			TickData: &types.TickDataType{
				Reserve0AndShares: make([]*types.Reserve0AndSharesType, feeSize),
				Reserve1:          make([]sdk.Dec, feeSize),
			},
			LimitOrderPool0To1: &types.LimitOrderPool{0, 0},
			LimitOrderPool1To0: &types.LimitOrderPool{0, 0},
		}

		// Sets Reserve0AShares to 0
		// Eliminates us having to check that each fee tier is nil or 0 when calling deposit
		for i, _ := range tick.TickData.Reserve0AndShares {
			tick.TickData.Reserve0AndShares[i] = &types.Reserve0AndSharesType{sdk.ZeroDec(), sdk.ZeroDec()}

		}
		// Sets Reserve1 to 0
		// Eliminates us having to check that each fee tier is nil or 0 when calling deposit
		for i, _ := range tick.TickData.Reserve1 {
			tick.TickData.Reserve1[i] = sdk.ZeroDec()
		}
	}

	return tick

}

// Helper function for Place Limit order retrieving and or initializing mapppings used for keeping track of limit orders
// Note: FillMap initialization is handled seperately in placeLimitOrder as it needed prior to this function being called.
func (k Keeper) PlaceLimitOrderMappingHelper(goCtx context.Context, pairId string, tickIndex int64, tokenIn string, currentLimitOrderKey uint64, receiver string) (types.LimitOrderPoolReserveMap, types.LimitOrderPoolUserShareMap, types.LimitOrderPoolTotalSharesMap) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Retrieves ReserveData Object from KVStore
	ReserveData, ReserveDataFound := k.GetLimitOrderPoolReserveMap(ctx, pairId, tickIndex, tokenIn, currentLimitOrderKey)
	// Retrieves UserShareMap object from KVStore
	UserShareData, UserShareDataFound := k.GetLimitOrderPoolUserShareMap(ctx, pairId, tickIndex, tokenIn, currentLimitOrderKey, receiver)
	// Retrives TotalSharesMap object from KVStore
	TotalSharesData, TotalSharesDataFound := k.GetLimitOrderPoolTotalSharesMap(ctx, pairId, tickIndex, tokenIn, currentLimitOrderKey)

	// If ReserveData object not found initialize it accordingly
	if !ReserveDataFound {
		ReserveData.Count = currentLimitOrderKey
		ReserveData.Reserves = sdk.ZeroDec()
		ReserveData.TickIndex = tickIndex
		ReserveData.Token = tokenIn
		ReserveData.PairId = pairId
	}

	// If UserShareData object is not found initialize it accordingly
	if !UserShareDataFound {
		UserShareData.Count = currentLimitOrderKey
		UserShareData.Address = receiver
		UserShareData.SharesOwned = sdk.ZeroDec()
		UserShareData.TickIndex = tickIndex
		UserShareData.Token = tokenIn
		UserShareData.PairId = pairId
	}

	// If TotalSharesData object is nout found initialize it accordingly
	if !TotalSharesDataFound {
		TotalSharesData.Count = currentLimitOrderKey
		TotalSharesData.TotalShares = sdk.ZeroDec()
		TotalSharesData.TickIndex = tickIndex
		TotalSharesData.Token = tokenIn
		TotalSharesData.PairId = pairId
	}

	_ = ctx

	return ReserveData, UserShareData, TotalSharesData
}

// Handles MsgPlaceLimitOrder, initializing (tick, pair) data structures if needed, calculating and storing information for a new limit order at a specific tick
func (k Keeper) PlaceLimitOrderCore(goCtx context.Context, msg *types.MsgPlaceLimitOrder, token0 string, token1 string, callerAddr sdk.AccAddress) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	is_0to1 := false
	if msg.TokenIn == token0 {
		is_0to1 = true
	}

	// checks if pair is initialized, if not intialize it and return pairId
	pairId, err := k.PairInit(goCtx, token0, token1, msg.TickIndex, 0, true, is_0to1)

	if err != nil {
		return err
	}

	//checks if tick is initialized and if not set count and currentLimitOrder key to 0
	/// tick is an object of the tickMapping for the specified tickIndex and mapping
	tick := k.PlaceLimitOrderHelper(goCtx, pairId, msg.TickIndex)

	// currentLimitOrder key for Pool0to1 or Pool0to1
	var currentLimitOrderKey uint64
	// currentLimitOrder count for Pool0to1 or Pool0to1
	var LimitOrderCount uint64
	// Shares Minted
	var newShares sdk.Dec

	// get pair map after checking if it is initialized in PairInit
	pair, _ := k.GetPairMap(ctx, pairId)

	// If we are placing limit order above current MaxTick update MaxTick
	if msg.TokenIn == token1 && msg.TickIndex > pair.MaxTick {
		pair.MaxTick = msg.TickIndex
	}

	// If we are placing limit order below current MinTick update MinTick
	if msg.TokenIn == token0 && msg.TickIndex < pair.MinTick {
		pair.MinTick = msg.TickIndex
	}

	// If TokenIn == token0 set count and key to values of LimitOrderPool0to1
	if msg.TokenIn == token0 {
		currentLimitOrderKey = tick.LimitOrderPool0To1.CurrentLimitOrderKey
		LimitOrderCount = tick.LimitOrderPool0To1.Count

		// Errors if place a limit order in amount0 at a tick greater than CurrentTick0to1
		if msg.TickIndex > pair.TokenPair.CurrentTick0To1 {
			return sdkerrors.Wrapf(types.ErrValidPairNotFound, "Cannot depsoit amount 0 at a tick greater than the CurrentTick0to1")
		}

		// If tokenIn == token1 set count and key to values of LimitOrderPool1to0
	} else {
		currentLimitOrderKey = tick.LimitOrderPool1To0.CurrentLimitOrderKey
		LimitOrderCount = tick.LimitOrderPool1To0.Count

		// Errors if placing a limit order in amoount1 at a tick less than Current1To0
		if msg.TickIndex < pair.TokenPair.CurrentTick1To0 {
			return sdkerrors.Wrapf(types.ErrValidPairNotFound, "Cannot depsoit amount 1 at a tick less than the CurrentTick0to1")
		}
	}

	// Retrives FillMap object from FillMapping KVSTore
	FillData, FillDataFound := k.GetLimitOrderPoolFillMap(ctx, pairId, msg.TickIndex, msg.TokenIn, currentLimitOrderKey)

	// inits FillMap object if not found
	if !FillDataFound {
		FillData.PairId = pairId
		FillData.TickIndex = msg.TickIndex
		FillData.Token = msg.TokenIn
		FillData.Count = currentLimitOrderKey
		FillData.FilledReserves = sdk.ZeroDec()

		// Handles creating a limit order given the current limit order is already partially filled
	} else if LimitOrderCount == currentLimitOrderKey && FillData.FilledReserves.GT(sdk.ZeroDec()) {
		// increments currentLimitOrderKey
		LimitOrderCount = LimitOrderCount + 1

		// Updates Pool0t01 / Pool1to0 accordingly
		if msg.TokenIn == token0 {
			tick.LimitOrderPool0To1.Count = LimitOrderCount
		} else {
			tick.LimitOrderPool1To0.Count = LimitOrderCount
		}
	}

	// Returns Map object for Reserve, UserShares, and TotalShares mapping
	ReserveData, UserShareData, TotalSharesData := k.PlaceLimitOrderMappingHelper(goCtx, pairId, msg.TickIndex, msg.TokenIn, currentLimitOrderKey, msg.Receiver)

	// Adds amountIn to ReserveData
	ReserveData.Reserves = ReserveData.Reserves.Add(msg.AmountIn)

	// newShares is equivalent to AmountIn
	newShares = msg.AmountIn

	// Adds newShares to User's shares owned
	UserShareData.SharesOwned = UserShareData.SharesOwned.Add(newShares)
	// Adds newShares to totalShares
	TotalSharesData.TotalShares = TotalSharesData.TotalShares.Add(newShares)

	// Set Fill, Reserve, UserShares, and TotalShares maps in KVStore
	k.SetLimitOrderPoolFillMap(ctx, FillData)
	k.SetLimitOrderPoolReserveMap(ctx, ReserveData)
	k.SetLimitOrderPoolUserShareMap(ctx, UserShareData)
	k.SetLimitOrderPoolTotalSharesMap(ctx, TotalSharesData)
	k.SetTickMap(ctx, pairId, tick)

	// If a new tick has been placed that tigtens the spread between currentTick0to1 and currentTick0to1 update CurrentTicks to the tighest ticks
	// @Dev assumes that msg.amountIn > 0
	if msg.TokenIn == token0 && ((msg.TickIndex < pair.TokenPair.CurrentTick0To1) && (msg.TickIndex > pair.TokenPair.CurrentTick1To0)) {
		pair.TokenPair.CurrentTick1To0 = msg.TickIndex
	} else if (msg.TickIndex < pair.TokenPair.CurrentTick0To1) && (msg.TickIndex > pair.TokenPair.CurrentTick1To0) {
		pair.TokenPair.CurrentTick0To1 = msg.TickIndex
	}

	// updates currentTick1to0/0To1 given the conditionals above

	k.SetPairMap(ctx, pair)

	// Sends AmountIn from Address to Module
	if msg.AmountIn.GT(sdk.ZeroDec()) {
		coin0 := sdk.NewCoin(msg.TokenIn, sdk.NewIntFromBigInt(msg.AmountIn.BigInt()))
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin0}); err != nil {
			return err
		}
	}

	ctx.EventManager().EmitEvent(types.CreatePlaceLimitOrderEvent(msg.Creator, msg.Receiver,
		token0, token1, msg.TokenIn, msg.AmountIn.String(), newShares.String(), strconv.Itoa(int(currentLimitOrderKey)),
	))

	return nil
}

// Handles MsgCancelLimitOrder, removing a specifed number of shares from a limit order and returning the respective amount in terms of the reserve to the user
func (k Keeper) CancelLimitOrderCore(goCtx context.Context, msg *types.MsgCancelLimitOrder, token0 string, token1 string, callerAddr sdk.AccAddress, receiverAddr sdk.AccAddress) error {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// PairId for token0, token1 ("token0/token1")
	pairId := k.CreatePairId(token0, token1)
	// Retrives TickMap object from KVStore
	_, tickFound := k.GetTickMap(ctx, pairId, msg.TickIndex)

	// If tick does not exist, then there is no liqudity to withdraw and thus error
	if !tickFound {
		return sdkerrors.Wrapf(types.ErrValidTickNotFound, "Valid tick not found ")
	}

	// Retrieves LimitOrderUserSharesMap object from KVStore for the specified key and keyToken
	UserSharesData, UserSharesDataFound := k.GetLimitOrderPoolUserShareMap(ctx, pairId, msg.TickIndex, msg.KeyToken, msg.Key, msg.Creator)
	// Retrieves LimitOrderReserevMap object from KVStore for the specified key and keyToken
	ReserveData, ReserveDataFound := k.GetLimitOrderPoolReserveMap(ctx, pairId, msg.TickIndex, msg.KeyToken, msg.Key)
	// Retrieves LimitOrderTotalSharesMap object from KVStore for the specified key and keyToken
	TotalSharesData, TotalShareDataFound := k.GetLimitOrderPoolTotalSharesMap(ctx, pairId, msg.TickIndex, msg.KeyToken, msg.Key)

	// If the UserShareMap does not exists, error (no shares exists for this user)
	// If ReserveDataFound or TotalSharesData is not found then this must not be a valid limit order to begin with
	if !UserSharesDataFound || !ReserveDataFound || !TotalShareDataFound {
		return sdkerrors.Wrapf(types.ErrValidLimitOrderMapsNotFound, "UserShareMap not found")
	}

	// Checks that sharesOUt is less than or equal to the number of shares owned by a specific users, error otherwise
	if msg.SharesOut.GT(UserSharesData.SharesOwned) {
		return sdkerrors.Wrapf(types.ErrCannotWithdrawLimitOrder, "sharesOut is larger than shares Owned at the specified tick")
	}

	// Calculate the value of the shares (in terms of the reserves) of the limit order to cancel
	amountOut := msg.SharesOut.Mul(ReserveData.Reserves).Quo(TotalSharesData.TotalShares)

	// Subtract shares canceled from the user mapping
	UserSharesData.SharesOwned = UserSharesData.SharesOwned.Sub(msg.SharesOut)
	// Subtract the value of shares (amountOut) from the overall reserveMap
	ReserveData.Reserves = ReserveData.Reserves.Sub(amountOut)
	// Subtract sharesCancled from the totalShares mapping
	TotalSharesData.TotalShares = TotalSharesData.TotalShares.Sub(msg.SharesOut)

	// Updates changes to mappings in the KVStore
	k.SetLimitOrderPoolUserShareMap(ctx, UserSharesData)
	k.SetLimitOrderPoolReserveMap(ctx, ReserveData)
	k.SetLimitOrderPoolTotalSharesMap(ctx, TotalSharesData)

	// Sends amountOut from module address to msg.Receiver account address
	if amountOut.GT(sdk.ZeroDec()) {
		coinOut := sdk.NewCoin(msg.KeyToken, sdk.NewIntFromBigInt(amountOut.BigInt()))
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, sdk.Coins{coinOut}); err != nil {
			return err
		}
	} else {
		return sdkerrors.Wrapf(types.ErrCannotWithdrawLimitOrder, "Cannot cancel liqudity from this limit order at this time")
	}

	// emit CancelLimitOrderEvent
	ctx.EventManager().EmitEvent(types.CancelLimitOrderEvent(msg.Creator, msg.Receiver,
		token0, token1, msg.KeyToken, strconv.Itoa(int(msg.Key)), amountOut.String(),
	))

	return nil
}

// Handles MsgWithdrawFilledLimitOrder, calculates and sends filled liqudity from module to user for a limit order based on amount wished to receive.
func (k Keeper) WithdrawFilledLimitOrderCore(goCtx context.Context, msg *types.MsgWithdrawFilledLimitOrder, token0 string, token1 string, callerAddr sdk.AccAddress, receiverAddr sdk.AccAddress) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// PairId for token0, token1 ("token0/token1")
	pairId := k.CreatePairId(token0, token1)
	// Retrives TickMap object from KVStore
	tick, tickFound := k.GetTickMap(ctx, pairId, msg.TickIndex)

	// If tick does not exist, then there is no liqudity to withdraw and thus error
	if !tickFound {
		return sdkerrors.Wrapf(types.ErrValidTickNotFound, "Valid tick not found ")
	}

	// Retrives LimitOrderFillMap object from KVStore for the specified key and keyToken
	FillData, FillDataFound := k.GetLimitOrderPoolFillMap(ctx, pairId, msg.TickIndex, msg.KeyToken, msg.Key)
	// Retrives LimitOrderReserevMap object from KVStore for the specified key and keyToken
	ReserveData, ReserveDataFound := k.GetLimitOrderPoolReserveMap(ctx, pairId, msg.TickIndex, msg.KeyToken, msg.Key)
	// Retrives LimitOrderUserSharesMap object from KVStore for the specified key and keyToken
	UserShareData, UserShareDataFound := k.GetLimitOrderPoolUserShareMap(ctx, pairId, msg.TickIndex, msg.KeyToken, msg.Key, msg.Creator)
	// Retrives LimitOrderUserSharesWithdrawnMap object from KVStore for the specified key and keyToken
	UserSharesWithdrawnData, UserSharesWithdrawnDataFound := k.GetLimitOrderPoolUserSharesWithdrawn(ctx, pairId, msg.TickIndex, msg.KeyToken, msg.Key, msg.Creator)
	// Retrives LimitOrderTotalSharesMap object from KVStore for the specified key and keyToken
	TotalSharesData, TotalSharesDataFound := k.GetLimitOrderPoolTotalSharesMap(ctx, pairId, msg.TickIndex, msg.KeyToken, msg.Key)

	// default initialize UserSharesWithdrawn (keeps track of liqudity withdrawn) if not initialized.
	if !UserSharesWithdrawnDataFound {
		UserSharesWithdrawnData = types.LimitOrderPoolUserSharesWithdrawn{
			PairId:          pairId,
			TickIndex:       msg.TickIndex,
			Token:           msg.KeyToken,
			Count:           msg.Key,
			Address:         msg.Creator,
			SharesWithdrawn: sdk.ZeroDec(),
		}
	}

	// If any of these map objects are not found, then a valid withdraw option will not exist, and thus error
	if !FillDataFound || !UserShareDataFound || !TotalSharesDataFound || !ReserveDataFound {
		return sdkerrors.Wrapf(types.ErrValidLimitOrderMapsNotFound, "Valid mappings for limit order withdraw not found")
	}

	if FillData.FilledReserves.Quo(FillData.FilledReserves.Add(ReserveData.Reserves)).LTE(UserSharesWithdrawnData.SharesWithdrawn.Quo(UserSharesWithdrawnData.SharesWithdrawn.Add(UserShareData.SharesOwned))) {
		return sdkerrors.Wrapf(types.ErrCannotWithdrawLimitOrder, "Cannot withdraw additional liqudity from this limit order at this time")
	}

	// If msg.KeyToken is token0 then the limit order was placed to exchange token 0 and receive token 1
	// So it is filled in a swap from token 1 to token 0 (since this adds token 1 into the limit order pool and removes 0)
	price := sdk.ZeroDec()
	if msg.KeyToken == token0 {
		p, err := k.Calc_price_1to0(tick.TickIndex)
		if err != nil {
			// TODO
		}
		price = p
	} else {
		p, err := k.Calc_price_0to1(tick.TickIndex)
		if err != nil {
			// TODO
		}
		price = p
	}

	sharesFilled := FillData.FilledReserves.Quo(price)

	// Calculates the sharesOut based on the UserShares withdrawn  compared to sharesLeft compared to remaining liquidity in reserves
	sharesOut := sharesFilled.Mul(UserShareData.SharesOwned.Add(UserSharesWithdrawnData.SharesWithdrawn)).Quo(sharesFilled.Add(ReserveData.Reserves)).Sub(UserSharesWithdrawnData.SharesWithdrawn)
	// calculate amountOut given sharesOut
	amountOut := sharesOut.Quo(price)
	// Subtracts amountOut from FilledReserves
	FillData.FilledReserves = FillData.FilledReserves.Sub(amountOut)

	// Updates useSharesWithdrawMap to include sharesOut
	UserSharesWithdrawnData.SharesWithdrawn = UserSharesWithdrawnData.SharesWithdrawn.Add(sharesOut)
	// Remove sharesOut from UserSharesMap
	UserShareData.SharesOwned = UserShareData.SharesOwned.Sub(sharesOut)
	// Removes sharesOut from TotalSharesMap

	// calculate amountOout given sharesOut
	TotalSharesData.TotalShares = TotalSharesData.TotalShares.Sub(sharesOut)

	// Updates changed LimitOrder Mappings in KVstore
	k.SetLimitOrderPoolFillMap(ctx, FillData)
	k.SetLimitOrderPoolUserShareMap(ctx, UserShareData)
	k.SetLimitOrderPoolUserSharesWithdrawn(ctx, UserSharesWithdrawnData)
	k.SetLimitOrderPoolTotalSharesMap(ctx, TotalSharesData)

	var tokenOut string

	// determines in which token to withdraw amountOut into
	if msg.KeyToken == token0 {
		tokenOut = token1
	} else {
		tokenOut = token0
	}

	// Sends amountOut from module address to msg.Receiver account address
	if amountOut.GT(sdk.ZeroDec()) {
		coinOut := sdk.NewCoin(tokenOut, sdk.NewIntFromBigInt(amountOut.BigInt()))
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, sdk.Coins{coinOut}); err != nil {
			return err
		}
	} else {
		return sdkerrors.Wrapf(types.ErrCannotWithdrawLimitOrder, "Cannot withdraw additional liqudity from this limit order at this time")
	}

	// emit WithdrawFilledLimitOrderEvent
	ctx.EventManager().EmitEvent(types.WithdrawFilledLimitOrderEvent(msg.Creator, msg.Receiver,
		token0, token1, msg.KeyToken, strconv.Itoa(int(msg.Key)), amountOut.String(),
	))

	return nil
}
