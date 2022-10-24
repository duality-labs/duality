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

func (k Keeper) PairInit(goCtx context.Context, token0 string, token1 string, tick_index int64, fee int64) (string, error) {

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

		// addEdges(goCtx, token0Index.Index, token1Index.Index)

		// Initializes a new pair object in mapping.
		// Note this is only one when nno ticks currently exists, and thus we set currentTick0to1 and currentTick1to0 to be tick_index +- fee
		k.SetPairMap(ctx, types.PairMap{
			PairId: pairId,
			TokenPair: &types.TokenPairType{
				CurrentTick0To1: tick_index - fee,
				CurrentTick1To0: tick_index + fee,
			},
			PairCount: 0,
		})

	}
	return pairId, nil

}

func (k Keeper) DepositHelper(goCtx context.Context, pairId string, pair types.PairMap, tickIndex int64, amount0 sdk.Dec, amount1 sdk.Dec, fee int64, feeIndex uint64) (sdk.Dec, sdk.Dec, sdk.Dec, sdk.Dec, sdk.Dec, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Getter functions for the tick that corresponds to reserve0 and shares (upper tick), and reserve1 (lower tick)
	lowerTick, lowerTickFound := k.GetTickMap(ctx, pairId, tickIndex-int64(fee))
	upperTick, upperTickFound := k.GetTickMap(ctx, pairId, tickIndex+int64(fee))

	// Default sets trueAmounts0/1 to amount0/1
	trueAmount0 := amount0
	trueAmount1 := amount1
	var sharesMinted sdk.Dec

	price, err := k.Calc_price(tickIndex, false)

	// Error check that price was calculated without error
	if err != nil {
		return sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec(), err
	}

	// In the case that the lower, upper tick is not found, or that the specified fee tier of the tick is empty, we default calculate shares (no reblancing, and setting initial share amounts)
	if !lowerTickFound || !upperTickFound || upperTick.TickData.Reserve0AndShares[feeIndex].TotalShares.Equal(sdk.ZeroDec()) {

		// a0 + a1 * price
		sharesMinted = amount0.Add(amount1.Mul(price))

		// Gets feeSize for feelist
		feeSize := k.GetFeeListCount(ctx)
		// if either of the lower or upper tick is not found we enumerate pairCount
		if !lowerTickFound || !upperTickFound {
			pair.PairCount = pair.PairCount + 1
		}
		// initialize lowerTick if not found
		if !lowerTickFound {

			// Creates an tick object of the speciied size and then iterates over each sub struct filling it with 0 values.

			lowerTick = types.TickMap{
				TickIndex: tickIndex - fee,
				TickData: &types.TickDataType{
					Reserve0AndShares: make([]*types.Reserve0AndSharesType, feeSize),
					Reserve1:          make([]sdk.Dec, feeSize),
				},
				LimitOrderPool0To1: &types.LimitOrderPool{0, 0},
				LimitOrderPool1To0: &types.LimitOrderPool{0, 0},
			}

			for i, _ := range lowerTick.TickData.Reserve0AndShares {
				lowerTick.TickData.Reserve0AndShares[i] = &types.Reserve0AndSharesType{sdk.ZeroDec(), sdk.ZeroDec()}

			}
			for i, _ := range lowerTick.TickData.Reserve1 {
				lowerTick.TickData.Reserve1[i] = sdk.ZeroDec()
			}

		}

		// intialize uppertick
		if !upperTickFound {

			// Creates an tick object of the specied size and then iterates over each sub struct filling it with 0 values.

			upperTick = types.TickMap{
				TickIndex: tickIndex + fee,
				TickData: &types.TickDataType{
					Reserve0AndShares: make([]*types.Reserve0AndSharesType, feeSize),
					Reserve1:          make([]sdk.Dec, feeSize),
				},
				LimitOrderPool0To1: &types.LimitOrderPool{0, 0},
				LimitOrderPool1To0: &types.LimitOrderPool{0, 0},
			}

			for i, _ := range upperTick.TickData.Reserve0AndShares {
				upperTick.TickData.Reserve0AndShares[i] = &types.Reserve0AndSharesType{sdk.ZeroDec(), sdk.ZeroDec()}

			}
			for i, _ := range upperTick.TickData.Reserve1 {
				upperTick.TickData.Reserve1[i] = sdk.ZeroDec()
			}

		}

		// No rebalancing is needed set trueamount0/1 to amount0/1
		trueAmount0 = amount0
		trueAmount1 = amount1
		// Sets the specifed tick/fee index in the array with the calculated value
		NewReserve0andShares := &types.Reserve0AndSharesType{
			Reserve0:    trueAmount0,
			TotalShares: sharesMinted,
		}

		fmt.Println(NewReserve0andShares)
		upperTick.TickData.Reserve0AndShares[feeIndex] = NewReserve0andShares

		lowerTick.TickData.Reserve1[feeIndex] = trueAmount1

	} else {

		// If feeList has been updated via a governance proposal updates future ticks to support this fee tier.
		if uint64(len(upperTick.TickData.Reserve1)) < k.GetFeeListCount(ctx) {
			upperTick.TickData.Reserve1 = append(upperTick.TickData.Reserve1, sdk.ZeroDec())
			upperTick.TickData.Reserve0AndShares = append(upperTick.TickData.Reserve0AndShares, &types.Reserve0AndSharesType{})
		}

		// If feeList has been updated via a governance proposal updates future ticks to support this fee tier.
		if uint64(len(lowerTick.TickData.Reserve1)) < k.GetFeeListCount(ctx) {
			lowerTick.TickData.Reserve1 = append(lowerTick.TickData.Reserve1, sdk.ZeroDec())
			lowerTick.TickData.Reserve0AndShares = append(lowerTick.TickData.Reserve0AndShares, &types.Reserve0AndSharesType{})
		}

		// Balance trueAmount1 to the pool ratio
		if upperTick.TickData.Reserve0AndShares[feeIndex].Reserve0.GT(sdk.ZeroDec()) {
			trueAmount1 = k.Min(amount1, lowerTick.TickData.Reserve1[feeIndex].Mul(amount0).Quo(upperTick.TickData.Reserve0AndShares[feeIndex].Reserve0))
			// trueAmount1 = min(amt1 , (reserve1 * amt0)/reserve0 )
		}

		// Balance trueAmount0 to the pool ratio
		if lowerTick.TickData.Reserve1[feeIndex].GT(sdk.ZeroDec()) {
			trueAmount0 = k.Min(amount0, upperTick.TickData.Reserve0AndShares[feeIndex].Reserve0.Mul(amount1).Quo(lowerTick.TickData.Reserve1[feeIndex]))
		}

		// if amount0 is 0 amt1/reserve1 * totalShares = sharesMinted
		// else if amt0/reserve0 * totalShares
		if trueAmount0.GT(sdk.ZeroDec()) {
			sharesMinted = (trueAmount0.Quo(upperTick.TickData.Reserve0AndShares[feeIndex].Reserve0).Mul(upperTick.TickData.Reserve0AndShares[feeIndex].TotalShares))
		} else {
			sharesMinted = (trueAmount1.Quo(lowerTick.TickData.Reserve1[feeIndex]).Mul(upperTick.TickData.Reserve0AndShares[feeIndex].TotalShares))
		}

		// Adds trueamount0 and sharesMinted to upperTick
		upperTick.TickData.Reserve0AndShares[feeIndex].Reserve0 = upperTick.TickData.Reserve0AndShares[feeIndex].Reserve0.Add(trueAmount0)
		upperTick.TickData.Reserve0AndShares[feeIndex].TotalShares = upperTick.TickData.Reserve0AndShares[feeIndex].TotalShares.Add(sharesMinted)

		// Adds trueamount1 to lowerTick
		lowerTick.TickData.Reserve1[feeIndex] = lowerTick.TickData.Reserve1[feeIndex].Add(trueAmount1)

	}

	// If a new tick has been placed that tigtens the range between currentTick0to1 and currentTick0to1 update CurrentTicks to the tighest ticks
	if trueAmount0.GT(sdk.ZeroDec()) && ((tickIndex+fee > pair.TokenPair.CurrentTick0To1) && (tickIndex+fee < pair.TokenPair.CurrentTick1To0)) {
		pair.TokenPair.CurrentTick1To0 = tickIndex + fee
	}

	if trueAmount1.GT(sdk.ZeroDec()) && ((tickIndex-fee > pair.TokenPair.CurrentTick0To1) && (tickIndex-fee < pair.TokenPair.CurrentTick1To0)) {
		pair.TokenPair.CurrentTick0To1 = tickIndex - fee
	}

	// Set pair, lower and upperTick KVStores
	k.SetPairMap(ctx, pair)
	k.SetTickMap(ctx, pairId, lowerTick)
	k.SetTickMap(ctx, pairId, upperTick)

	return trueAmount0, trueAmount1, sharesMinted, upperTick.TickData.Reserve0AndShares[feeIndex].Reserve0, lowerTick.TickData.Reserve1[feeIndex], nil

}

// TODO
// While this works this s a rough get around to doing float based math with sdk.Decs
// tickIndex refers to the index of a specified tick for a given pool
// StartingToken determines the ratio of our price, price when false, 1/price when true.
func (k Keeper) Calc_price(tick_Index int64, startingToken bool) (sdk.Dec, error) {
	floatPrice := math.Pow(1.0001, float64(tick_Index))
	sPrice := fmt.Sprintf("%f", floatPrice)

	price, err := sdk.NewDecFromStr(sPrice)

	if err != nil {
		return sdk.ZeroDec(), err
	} else {
		if startingToken {
			price = sdk.OneDec().Quo(price)
			return price, nil
		} else {
			return price, nil
		}

	}

}

// Returns the smaller of two sdk.Decs
func (k Keeper) Min(a, b sdk.Dec) sdk.Dec {
	if a.LT(b) {
		return a
	}
	return b
}

func (k Keeper) DepositCore(goCtx context.Context, msg *types.MsgDeposit, token0 string, token1 string, callerAddr sdk.AccAddress, amounts0 []sdk.Dec, amounts1 []sdk.Dec) error {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Note feeIndex is the corresponding index to a fee Value not the feeValue itself
	// Fee Initialized to initialize tick if pair does not exists

	feelist := k.GetAllFeeList(ctx)

	//Checks to see if given pair has been initialied, if not intializes, and returns pairId and pairMap
	pairId, err := k.PairInit(goCtx, token0, token1, msg.TickIndexes[0], feelist[msg.FeeIndexes[0]].Fee)

	if err != nil {
		return err
	}

	pair, _ := k.GetPairMap(ctx, pairId)
	totalAmountReserve0 := sdk.ZeroDec()
	totalAmountReserve1 := sdk.ZeroDec()

	for i, _ := range amounts0 {
		// Errors if depositing amount0 at a tick less than CurrentTick1to0
		if amounts0[i].GT(sdk.ZeroDec()) && ((msg.TickIndexes[i] + feelist[msg.FeeIndexes[i]].Fee) < pair.TokenPair.CurrentTick0To1) {
			return sdkerrors.Wrapf(types.ErrValidPairNotFound, "Cannot depsoit amount 0 at a tick less than the CurrentTick1to0")
		}
		// Errors if depositing amount1 at a tick less than CurrentTick0to1
		if amounts1[i].GT(sdk.ZeroDec()) && ((msg.TickIndexes[i] - feelist[msg.FeeIndexes[i]].Fee) > pair.TokenPair.CurrentTick1To0) {
			return sdkerrors.Wrapf(types.ErrValidPairNotFound, "Cannot depsoit amount 0 at a tick greater than the CurrentTick0to1")
		}

		// Calls k.DepositHelper which calculates the true amounts of token0, token1, and sets the corresponding pair and tick maps
		trueAmount0, trueAmount1, sharesMinted, newReserve0, newReserve1, err := k.DepositHelper(goCtx, pairId, pair, msg.TickIndexes[i], amounts0[i], amounts1[i], feelist[msg.FeeIndexes[i]].Fee, msg.FeeIndexes[i])

		if err != nil {
			return nil
		}

		// Calculate Shares
		shares, sharesFound := k.GetShares(ctx, msg.Receiver, pairId, msg.TickIndexes[i], msg.FeeIndexes[i])

		// Initializes a new tick if sharesFound does not exists
		if !sharesFound {
			shares = types.Shares{
				Address:     msg.Receiver,
				PairId:      pairId,
				TickIndex:   msg.TickIndexes[i],
				FeeIndex:    msg.FeeIndexes[i],
				SharesOwned: sharesMinted,
			}
		} else {
			//Updates shares.SharesOwned
			shares.SharesOwned = shares.SharesOwned.Add(sharesMinted)
		}

		// Update share logic to KVStore
		k.SetShares(ctx, shares)

		totalAmountReserve0 = totalAmountReserve0.Add(trueAmount0)
		totalAmountReserve1 = totalAmountReserve1.Add(trueAmount1)

		// Event is defined in types/Events.go
		ctx.EventManager().EmitEvent(types.CreateDepositEvent(msg.Creator, msg.Receiver,
			token0, token1, fmt.Sprint(msg.TickIndexes[i]), fmt.Sprint(msg.FeeIndexes[i]),
			newReserve0.Sub(trueAmount0).String(), newReserve1.Sub(trueAmount1).String(), newReserve0.String(), newReserve1.String(),
			sharesMinted.String()),
		)
	}

	// Send TrueAmount0 to Module
	/// @Dev Due to a sdk.send constraint this only sends if trueAmount0 is greater than 0
	if totalAmountReserve0.GT(sdk.ZeroDec()) {
		coin0 := sdk.NewCoin(token0, sdk.NewIntFromBigInt(totalAmountReserve0.BigInt()))
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin0}); err != nil {
			return err
		}
	}

	// Send TrueAmount1 to Module
	/// @Dev Due to a sdk.send constraint this only sends if trueAmount1 is greater than 0
	if totalAmountReserve1.GT(sdk.ZeroDec()) {
		coin1 := sdk.NewCoin(token1, sdk.NewIntFromBigInt(totalAmountReserve1.BigInt()))
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin1}); err != nil {
			return err
		}
	}

	_ = goCtx
	return nil
}

func (k Keeper) WithdrawCore(goCtx context.Context, msg *types.MsgWithdrawl, token0 string, token1 string, callerAddr sdk.AccAddress, receiverAddr sdk.AccAddress) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pairId := k.CreatePairId(token0, token1)

	pair, pairFound := k.GetPairMap(ctx, pairId)

	if !pairFound {
		return sdkerrors.Wrapf(types.ErrValidPairNotFound, "Pair not found")
	}

	totalReserve0ToRemove := sdk.ZeroDec()
	totalReserve1ToRemove := sdk.ZeroDec()

	for i, _ := range msg.SharesToRemove {
		feeValue, _ := k.GetFeeList(ctx, uint64(msg.FeeIndexes[i]))
		fee := feeValue.Fee

		upperTick, upperTickFound := k.GetTickMap(ctx, pairId, msg.TickIndexes[i]+int64(fee))
		lowerTick, lowerTickFound := k.GetTickMap(ctx, pairId, msg.TickIndexes[i]-int64(fee))

		var OldReserve0 sdk.Dec
		var OldReserve1 sdk.Dec
		if !upperTickFound || !lowerTickFound {
			return sdkerrors.Wrapf(types.ErrValidTickNotFound, "No tick found at the requested index")
		}

		if upperTick.TickData.Reserve0AndShares[msg.FeeIndexes[0]].TotalShares.Equal(sdk.ZeroDec()) {
			return sdkerrors.Wrapf(types.ErrValidTickNotFound, "No tick found at the requested index")
		}

		reserve0ToRemove := (msg.SharesToRemove[i].Quo(upperTick.TickData.Reserve0AndShares[msg.FeeIndexes[i]].TotalShares)).Mul(upperTick.TickData.Reserve0AndShares[msg.FeeIndexes[i]].Reserve0)
		reserve1ToRemove := (msg.SharesToRemove[i].Quo(upperTick.TickData.Reserve0AndShares[msg.FeeIndexes[i]].TotalShares)).Mul(lowerTick.TickData.Reserve1[msg.FeeIndexes[i]])

		OldReserve0 = upperTick.TickData.Reserve0AndShares[msg.FeeIndexes[i]].Reserve0
		OldReserve1 = lowerTick.TickData.Reserve1[msg.FeeIndexes[i]]

		upperTick.TickData.Reserve0AndShares[msg.FeeIndexes[i]].Reserve0 = upperTick.TickData.Reserve0AndShares[msg.FeeIndexes[i]].Reserve0.Sub(reserve0ToRemove)
		lowerTick.TickData.Reserve1[msg.FeeIndexes[i]] = lowerTick.TickData.Reserve1[msg.FeeIndexes[i]].Sub(reserve1ToRemove)
		upperTick.TickData.Reserve0AndShares[msg.FeeIndexes[i]].TotalShares = upperTick.TickData.Reserve0AndShares[msg.FeeIndexes[i]].TotalShares.Sub(msg.SharesToRemove[i])

		shareOwner, shareOwnerFound := k.GetShares(ctx, msg.Creator, pairId, msg.TickIndexes[i], msg.FeeIndexes[i])

		if !shareOwnerFound {
			return sdkerrors.Wrapf(types.ErrValidShareNotFound, "No valid share owner fonnd")
		}

		shareOwner.SharesOwned = shareOwner.SharesOwned.Sub(msg.SharesToRemove[i])

		isTickEmpty := true
		if upperTick.TickData.Reserve0AndShares[msg.FeeIndexes[i]].TotalShares.Equal(sdk.ZeroDec()) {
			for _, s := range upperTick.TickData.Reserve0AndShares {
				if s.TotalShares.GT(sdk.ZeroDec()) {
					isTickEmpty = false
				}
			}
		}

		if isTickEmpty && upperTick.LimitOrderPool0To1.CurrentLimitOrderKey == 0 && upperTick.LimitOrderPool1To0.CurrentLimitOrderKey == 0 {

			// handles pairCount management given that MOST of the time we can not decrement pairCount.
			Pool0to1, Pool0to1Found := k.GetLimitOrderPoolTotalSharesMap(ctx, pairId, upperTick.TickIndex, token0, 0)
			Pool1to0, Pool1to0Found := k.GetLimitOrderPoolTotalSharesMap(ctx, pairId, upperTick.TickIndex, token1, 0)

			if (!Pool0to1Found || Pool0to1.TotalShares.Equal(sdk.ZeroDec())) && (!Pool1to0Found || Pool1to0.TotalShares.Equal(sdk.ZeroDec())) {
				pair.PairCount = pair.PairCount - 1
			}

		}

		if isTickEmpty && (msg.TickIndexes[i]+int64(fee) == pair.TokenPair.CurrentTick1To0) {

			tickFound := false
			c := 0
			for tickFound != true && pair.PairCount < int64(c) {
				c++
				_, tickFound = k.GetTickMap(ctx, pairId, (msg.TickIndexes[i] + fee + int64(c)))

			}

			pair.TokenPair.CurrentTick1To0 = (msg.TickIndexes[i] + fee + int64(c))
		}

		if isTickEmpty && (msg.TickIndexes[i]-int64(fee) == pair.TokenPair.CurrentTick0To1) {

			tickFound := false
			c := 0
			for tickFound != true && pair.PairCount < int64(c) {
				c++
				_, tickFound = k.GetTickMap(ctx, pairId, (msg.TickIndexes[i] - fee - int64(c)))

			}

			pair.TokenPair.CurrentTick1To0 = (msg.TickIndexes[i] - fee - int64(c))
		}

		totalReserve0ToRemove = totalReserve0ToRemove.Add(reserve0ToRemove)
		totalReserve1ToRemove = totalReserve1ToRemove.Add(reserve1ToRemove)

		k.SetShares(ctx, shareOwner)
		k.SetTickMap(ctx, pairId, upperTick)
		k.SetTickMap(ctx, pairId, lowerTick)

		ctx.EventManager().EmitEvent(types.CreateWithdrawEvent(msg.Creator, msg.Receiver,
			token0, token1, fmt.Sprint(msg.TickIndexes), fmt.Sprint(msg.FeeIndexes), OldReserve0.String(), OldReserve1.String(),
			upperTick.TickData.Reserve0AndShares[msg.FeeIndexes[i]].Reserve0.String(), lowerTick.TickData.Reserve1[msg.FeeIndexes[i]].String(),
		))
	}

	k.SetPairMap(ctx, pair)

	if totalReserve0ToRemove.GT(sdk.ZeroDec()) {
		coin0 := sdk.NewCoin(token0, sdk.NewIntFromBigInt(totalReserve0ToRemove.BigInt()))
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, sdk.Coins{coin0}); err != nil {
			return err
		}
	}

	if totalReserve1ToRemove.GT(sdk.ZeroDec()) {
		coin1 := sdk.NewCoin(token1, sdk.NewIntFromBigInt(totalReserve1ToRemove.BigInt()))
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, sdk.Coins{coin1}); err != nil {
			return err
		}
	}

	_ = ctx
	return nil
}

////// Swap Functions

func (k Keeper) Swap0to1(goCtx context.Context, msg *types.MsgSwap, token0 string, token1 string, callerAddr sdk.AccAddress) (sdk.Dec, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// pair idea: "token0/token1"
	pairId := k.CreatePairId(token0, token1)

	// size of the feeList
	feeSize := k.GetFeeListCount(ctx)
	feelist := k.GetAllFeeList(ctx)
	// geets the PairMap from the KVstore given pairId
	pair, pairFound := k.GetPairMap(ctx, pairId)

	// If toknePair does not exists a swap cannot be made through it, error
	if !pairFound {
		return sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrValidPairNotFound, "Pair not found")
	}

	// Counts how many ticks we have iterated through, compare to initialized ticks in the pair
	// @Note Heuristic to remove unecessary looping
	count := 0

	//amount_left is the amount left to deposit
	amount_left := msg.AmountIn

	fmt.Println(amount_left)
	// amount to return to receiver
	amount_out := sdk.ZeroDec()

	// verify that amount left is not zero and that there are additional valid ticks to check
	for !amount_left.Equal(sdk.ZeroDec()) && (count < int(pair.PairCount)) {

		// Tick data for tick that holds information about reserve1
		Current1Data, Current1Found := k.GetTickMap(ctx, pairId, pair.TokenPair.CurrentTick0To1)

		fmt.Println(count)
		if !Current1Found {
			pair.TokenPair.CurrentTick0To1 = pair.TokenPair.CurrentTick0To1 - 1
			continue
		}

		// iterate count
		count++

		var i uint64

		// iterator for feeList
		i = 0
		for i < feeSize && !amount_left.Equal(sdk.ZeroDec()) {
			// gets fee for given feeIndex
			fee := feelist[i].Fee
			Current0Data, Current0Found := k.GetTickMap(ctx, pairId, pair.TokenPair.CurrentTick0To1+2*fee)
			//Current0Datam := Current0Data.TickData.Reserve1[i]

			// If tick/feeIndex pair is not found continue
			if !Current0Found {
				i++
				continue
			}
			// calculate currentPrice
			price, err := k.Calc_price(pair.TokenPair.CurrentTick0To1, false)

			if err != nil {
				return sdk.ZeroDec(), err
			}

			// price * amout_left + amount_out < minOut, error we cannot meet minOut threshold
			if price.Mul(amount_left).Add(amount_out).LT(msg.MinOut) {
				return sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Amount Out is less than minium amount out specified: swap failed")
			}

			// price * r1 < amount_left
			if price.Mul(Current1Data.TickData.Reserve1[i]).LT(amount_left) {
				// amount_out += r1 (adds as all of reserve1 to amount_out)
				amount_out = amount_out.Add(Current1Data.TickData.Reserve1[i])
				// decrement amount_left by price * r1
				amount_left = amount_left.Sub(price.Mul(Current1Data.TickData.Reserve1[i]))
				//updates reserve0 with the new amountIn
				Current0Data.TickData.Reserve0AndShares[i].Reserve0 = Current0Data.TickData.Reserve0AndShares[i].Reserve0.Add(price.Mul(Current1Data.TickData.Reserve1[i]))
				// sets reserve1 to 0
				Current1Data.TickData.Reserve1[i] = sdk.ZeroDec()

			} else {
				// amountOut += amount_left * price
				amount_out = amount_out.Add(amount_left.Mul(price))
				// increment reserve0 with amountLeft
				Current0Data.TickData.Reserve0AndShares[i].Reserve0 = Current0Data.TickData.Reserve0AndShares[i].Reserve0.Add(amount_left)
				// decrement reserve1 with amount_left * price
				Current1Data.TickData.Reserve1[i] = Current1Data.TickData.Reserve1[i].Sub(amount_left.Mul(price))
				// set amountLeft to 0
				amount_left = sdk.ZeroDec()
			}

			//updates feeIndex
			i++

			//Make updates to tickMap containing reserve0/1 data to the KVStore
			k.SetTickMap(ctx, pairId, Current0Data)
			k.SetTickMap(ctx, pairId, Current1Data)
		}

		// if feeIndex is equal to the largest index in feeList
		if i == feeSize {

			// assigns a new variable err to handle err not being initialized in this conditional
			var err error
			// runs swaps for any limitOrders at the specified tick, updating amount_left, amount_out accordingly
			fmt.Println("PreSwap", amount_left, amount_out)
			amount_left, amount_out, err = k.SwapLimitOrder1to0(goCtx, pairId, token1, amount_out, amount_left, pair.TokenPair.CurrentTick0To1)
			fmt.Println("Post Swap", amount_left, amount_out)

			if err != nil {
				return sdk.ZeroDec(), err
			}

			// iterates CurrentTick0to1
			pair.TokenPair.CurrentTick0To1 = pair.TokenPair.CurrentTick0To1 - 1
		}
	}

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

func (k Keeper) Swap1to0(goCtx context.Context, msg *types.MsgSwap, token0 string, token1 string, callerAddr sdk.AccAddress) (sdk.Dec, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// pair idea: "token0/token1"
	pairId := k.CreatePairId(token0, token1)

	// size of the feeList
	feeSize := k.GetFeeListCount(ctx)
	feelist := k.GetAllFeeList(ctx)
	// geets the PairMap from the KVstore given pairId
	pair, pairFound := k.GetPairMap(ctx, pairId)

	if !pairFound {
		return sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrValidPairNotFound, "Pair not found")
	}

	// Counts how many ticks we have iterated through, compare to initialized ticks in the pair
	// @Note Heuristic to remove unecessary looping
	count := 0

	//amount_left is the amount left to deposit
	amount_left := msg.AmountIn

	// amount to return to receiver
	amount_out := sdk.ZeroDec()

	// verify that amount left is not zero and that there are additional valid ticks to check

	for !amount_left.Equal(sdk.ZeroDec()) && (count < int(pair.PairCount)) {
		fmt.Println("Amount left", amount_left)
		fmt.Println("Count", count)

		Current0Data, Current0Found := k.GetTickMap(ctx, pairId, pair.TokenPair.CurrentTick1To0)
		//Current0Datam := Current0Data.TickData.Reserve1[i]

		// If tick/feeIndex pair is not found continue

		if !Current0Found {
			pair.TokenPair.CurrentTick1To0 = pair.TokenPair.CurrentTick1To0 + 1
			continue
		}

		// iterate count
		count++

		var i uint64

		// iterator for feeList
		i = 0
		for i < feeSize && !amount_left.Equal(sdk.ZeroDec()) {
			// gets fee for given feeIndex
			fee := feelist[i].Fee

			Current1Data, Current1Found := k.GetTickMap(ctx, pairId, pair.TokenPair.CurrentTick1To0-2*fee)

			if !Current1Found {
				i++
				continue
			}
			//Current0Datam := Current0Data.TickData.Reserve1[i]

			// calculate currentPrice and inverts
			price, err := k.Calc_price(pair.TokenPair.CurrentTick1To0, true)

			if err != nil {
				return sdk.ZeroDec(), err
			}

			// price * amout_left + amount_out < minOut, error we cannot meet minOut threshold
			if price.Mul(amount_left).Add(amount_out).LT(msg.MinOut) {
				return sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Amount Out is less than minium amount out specified: swap failed")
			}

			// price * r1 < amount_left
			if price.Mul(Current0Data.TickData.Reserve0AndShares[i].Reserve0).LT(amount_left) {
				// amountOut += amount_left * price
				amount_out = amount_out.Add(Current0Data.TickData.Reserve0AndShares[i].Reserve0)
				// decrement amount_left by price * reserve0
				amount_left = amount_left.Sub(price.Mul(Current0Data.TickData.Reserve0AndShares[i].Reserve0))
				//updates reserve1 with the new amountIn
				Current1Data.TickData.Reserve1[i] = Current1Data.TickData.Reserve1[i].Add(price.Mul(Current0Data.TickData.Reserve0AndShares[i].Reserve0))
				// sets reserve0 to 0
				Current0Data.TickData.Reserve0AndShares[i].Reserve0 = sdk.ZeroDec()

			} else {
				// amountOut += amount_left * price
				amount_out = amount_out.Add(amount_left.Mul(price))
				// increment reserve1 with amountLeft
				Current1Data.TickData.Reserve1[i] = Current1Data.TickData.Reserve1[i].Add(amount_left)
				// decrement reserve0 with amount_left * price
				Current0Data.TickData.Reserve0AndShares[i].Reserve0 = Current0Data.TickData.Reserve0AndShares[i].Reserve0.Sub(amount_left.Mul(price))
				// set amountLeft to 0
				amount_left = sdk.ZeroDec()
			}

			//updates feeIndex
			i++

			//Make updates to tickMap containing reserve0/1 data to the KVStore
			k.SetTickMap(ctx, pairId, Current0Data)
			k.SetTickMap(ctx, pairId, Current1Data)
		}

		// if feeIndex is equal to the largest index in feeList
		if i == feeSize {

			// assigns a new variable err to handle err not being initialized in this conditional
			var err error
			// runs swaps for any limitOrders at the specified tick, updating amount_left, amount_out accordingly
			amount_left, amount_out, err = k.SwapLimitOrder1to0(goCtx, pairId, token0, amount_out, amount_left, pair.TokenPair.CurrentTick1To0)

			if err != nil {
				return sdk.ZeroDec(), err
			}

			pair.TokenPair.CurrentTick1To0 = pair.TokenPair.CurrentTick1To0 + 1
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

// Swaps through Limit Orders
// Returns amount_out, amount_left, error
func (k Keeper) SwapLimitOrder0to1(goCtx context.Context, pairId string, tokenIn string, amount_out sdk.Dec, amount_left sdk.Dec, CurrentTick1to0 int64) (sdk.Dec, sdk.Dec, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// returns price for the given tick and specified direction (0 -> 1)
	price, err := k.Calc_price(CurrentTick1to0, false)

	if err != nil {
		return sdk.ZeroDec(), sdk.ZeroDec(), err
	}

	// Gets tick for specified tick at currentTick0to1
	tick, tickFound := k.GetTickMap(ctx, pairId, CurrentTick1to0)

	if !tickFound {
		return amount_left, amount_out, nil
	}

	// Gets Reserve, Fill map for the specified CurrentLimitOrderKey
	ReserveMap, ReserveMapFound := k.GetLimitOrderPoolReserveMap(ctx, pairId, CurrentTick1to0, tokenIn, tick.LimitOrderPool1To0.CurrentLimitOrderKey)
	FillMap, FillMapFound := k.GetLimitOrderPoolFillMap(ctx, pairId, CurrentTick1to0, tokenIn, tick.LimitOrderPool1To0.CurrentLimitOrderKey)

	// errors if ReserveMapFound is not found
	if !ReserveMapFound {
		return amount_left, amount_out, nil
	}

	// if no tokens have been filled at this key value, initialize to 0
	if !FillMapFound {
		FillMap.Count = tick.LimitOrderPool1To0.CurrentLimitOrderKey
		FillMap.TickIndex = CurrentTick1to0
		FillMap.PairId = pairId
		FillMap.Fill = sdk.ZeroDec()
	}

	// If there isn't enough liqudity to end trade handle updates this way
	if price.Mul(ReserveMap.Reserves).LT(amount_left) {
		// Adds remaining reserves to amount_out
		amount_out = amount_out.Add(ReserveMap.Reserves)
		// Subtracts reserves from amount_left
		amount_left = amount_left.Sub(ReserveMap.Reserves)
		// adds price * reserves to the filledMap
		FillMap.Fill = FillMap.Fill.Add(price.Mul(ReserveMap.Reserves))
		// sets reserves to 0
		ReserveMap.Reserves = sdk.ZeroDec()

		// increments the limitOrderkey as previous tick has been completely filled
		tick.LimitOrderPool1To0.CurrentLimitOrderKey++

		// checks the next currentLimitOrderKey
		ReserveMapNextKey, ReserveMapNextKeyFound := k.GetLimitOrderPoolReserveMap(ctx, pairId, CurrentTick1to0, tokenIn, tick.LimitOrderPool1To0.CurrentLimitOrderKey)
		FillMapNextKey, FillMapNextKeyFound := k.GetLimitOrderPoolFillMap(ctx, pairId, CurrentTick1to0, tokenIn, tick.LimitOrderPool1To0.CurrentLimitOrderKey)

		// if no tokens have been filled at this key value, initialize to 0
		if !FillMapNextKeyFound {
			FillMapNextKey.Count = tick.LimitOrderPool1To0.CurrentLimitOrderKey
			FillMapNextKey.TickIndex = CurrentTick1to0
			FillMapNextKey.PairId = pairId
			FillMapNextKey.Fill = sdk.ZeroDec()
		}

		// If there is still not enough liquidity to end trade handle update this way
		if ReserveMapNextKeyFound && price.Mul(ReserveMapNextKey.Reserves).LT(amount_left) {
			// Adds remaining reserves to amount_out
			amount_out = amount_out.Add(ReserveMapNextKey.Reserves)
			// Subtracts reserves from amount_left
			amount_left = amount_left.Sub(ReserveMapNextKey.Reserves)
			// adds price * reserves to the filledMap
			FillMapNextKey.Fill = FillMapNextKey.Fill.Add(price.Mul(ReserveMapNextKey.Reserves))
			// sets reserve to 0
			ReserveMapNextKey.Reserves = sdk.ZeroDec()

			// increments the limitOrderKey
			tick.LimitOrderPool1To0.CurrentLimitOrderKey++

			// If there IS enough liqudity to end trade handle update this way
		} else if ReserveMapNextKeyFound && price.Mul(ReserveMapNextKey.Reserves).GT(amount_left) {
			// calculate anmout to output (will be a portion of reserves)
			amount_out = amount_out.Add(amount_left.Mul(price))
			// Add the amount_left to the amount flled in the fill mapping
			FillMapNextKey.Fill = FillMapNextKey.Fill.Add(amount_left)
			// subtract amount_left * price to the ReserveMapping
			ReserveMapNextKey.Reserves = ReserveMapNextKey.Reserves.Sub(amount_left.Mul(price))
			// set amount_left to 0
			amount_left = sdk.ZeroDec()
		}

		// Updates mapping for the original limit order key + 1 (next key)
		// @dev we set mappings within the conditionnal, as the tick mappings have not been initialized outside of it
		k.SetLimitOrderPoolFillMap(ctx, FillMapNextKey)
		k.SetLimitOrderPoolReserveMap(ctx, ReserveMapNextKey)

		// If there IS enough liqudity to end trade handle update this way
	} else {
		// calculate anmout to output (will be a portion of reserves)
		amount_out = amount_out.Add(amount_left.Mul(price))
		// Add the amount_left to the amount flled in the fill mapping
		FillMap.Fill = FillMap.Fill.Add(amount_left)
		// subtract amount_left * price to the ReserveMapping
		ReserveMap.Reserves = ReserveMap.Reserves.Sub(amount_left.Mul(price))
		// set amount_left to 0
		amount_left = sdk.ZeroDec()
	}

	// Updates mappings of reserve and filledReserves based on the original limitOrderCurrentKey to the KVStore
	k.SetLimitOrderPoolReserveMap(ctx, ReserveMap)
	k.SetLimitOrderPoolFillMap(ctx, FillMap)

	//Updates limitOrderCurrentKey based on if any limitOrders were completely filled.
	k.SetTickMap(ctx, pairId, tick)

	_ = ctx

	return amount_left, amount_out, nil
}

func (k Keeper) SwapLimitOrder1to0(goCtx context.Context, pairId string, tokenIn string, amount_out sdk.Dec, amount_left sdk.Dec, CurrentTick0to1 int64) (sdk.Dec, sdk.Dec, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// returns price for the given tick and specified direction (0 -> 1)
	price, err := k.Calc_price(CurrentTick0to1, false)

	if err != nil {
		return sdk.ZeroDec(), sdk.ZeroDec(), err
	}

	tick, tickFound := k.GetTickMap(ctx, pairId, CurrentTick0to1)

	if !tickFound {
		return amount_left, amount_out, nil
	}

	fmt.Println("Mapping gettng paramters")
	fmt.Println(tick.LimitOrderPool0To1)
	fmt.Println(tokenIn)
	fmt.Println(CurrentTick0to1)
	fmt.Println(pairId)

	ReserveMap, ReserveMapFound := k.GetLimitOrderPoolReserveMap(ctx, pairId, CurrentTick0to1, tokenIn, tick.LimitOrderPool0To1.CurrentLimitOrderKey)
	FillMap, FillMapFound := k.GetLimitOrderPoolFillMap(ctx, pairId, CurrentTick0to1, tokenIn, tick.LimitOrderPool0To1.CurrentLimitOrderKey)

	// errors if ReserveMapFound is not found
	if !ReserveMapFound {
		return amount_left, amount_out, nil
	}

	// if no tokens have been filled at this key value, initialize to 0
	if !FillMapFound {
		FillMap.Count = tick.LimitOrderPool0To1.CurrentLimitOrderKey
		FillMap.TickIndex = CurrentTick0to1
		FillMap.PairId = pairId
		FillMap.Fill = sdk.ZeroDec()
	}

	// If there isn't enough liqudity to end trade handle updates this way
	if price.Mul(ReserveMap.Reserves).LT(amount_left) {
		// Adds remaining reserves to amount_out
		amount_out = amount_out.Add(ReserveMap.Reserves)
		// Subtracts reserves from amount_left
		amount_left = amount_left.Sub(ReserveMap.Reserves)
		// adds price * reserves to the filledMap
		FillMap.Fill = FillMap.Fill.Add(price.Mul(ReserveMap.Reserves))
		// sets reserves to 0
		ReserveMap.Reserves = sdk.ZeroDec()

		// increments the limitOrderkey as previous tick has been completely filled
		tick.LimitOrderPool0To1.CurrentLimitOrderKey++

		// checks the next currentLimitOrderKey
		ReserveMapNextKey, ReserveMapNextKeyFound := k.GetLimitOrderPoolReserveMap(ctx, pairId, CurrentTick0to1, tokenIn, tick.LimitOrderPool0To1.CurrentLimitOrderKey)
		FillMapNextKey, FillMapNextKeyFound := k.GetLimitOrderPoolFillMap(ctx, pairId, CurrentTick0to1, tokenIn, tick.LimitOrderPool0To1.CurrentLimitOrderKey)

		// if no tokens have been filled at this key value, initialize to 0
		if !FillMapNextKeyFound {
			FillMapNextKey.Count = tick.LimitOrderPool0To1.CurrentLimitOrderKey
			FillMapNextKey.TickIndex = CurrentTick0to1
			FillMapNextKey.PairId = pairId
			FillMapNextKey.Fill = sdk.ZeroDec()
		}

		if ReserveMapNextKeyFound && price.Mul(ReserveMapNextKey.Reserves).LT(amount_left) {
			// Adds remaining reserves to amount_out
			amount_out = amount_out.Add(ReserveMapNextKey.Reserves)
			// Subtracts reserves from amount_left
			amount_left = amount_left.Sub(ReserveMapNextKey.Reserves)
			// adds price * reserves to the filledMap
			FillMapNextKey.Fill = FillMapNextKey.Fill.Add(price.Mul(ReserveMapNextKey.Reserves))
			// sets reserve to 0
			ReserveMapNextKey.Reserves = sdk.ZeroDec()

			// increments the limitOrderKey
			tick.LimitOrderPool0To1.CurrentLimitOrderKey++

		} else if ReserveMapNextKeyFound && price.Mul(ReserveMapNextKey.Reserves).GT(amount_left) {
			// calculate anmout to output (will be a portion of reserves)
			amount_out = amount_out.Add(amount_left.Mul(price))
			// Add the amount_left to the amount flled in the fill mapping
			FillMapNextKey.Fill = FillMapNextKey.Fill.Add(amount_left)
			// subtract amount_left * price to the ReserveMapping
			ReserveMapNextKey.Reserves = ReserveMapNextKey.Reserves.Sub(amount_left.Mul(price))
			// set amount_left to 0
			amount_left = sdk.ZeroDec()
		}

		// Updates mapping for the original limit order key + 1 (next key)
		// @dev we set mappings within the conditionnal, as the tick mappings have not been initialized outside of it
		k.SetLimitOrderPoolFillMap(ctx, FillMapNextKey)
		k.SetLimitOrderPoolReserveMap(ctx, ReserveMapNextKey)

		// If there IS enough liqudity to end trade handle update this way
	} else {
		// calculate anmout to output (will be a portion of reserves)
		amount_out = amount_out.Add(amount_left.Mul(price))
		// Add the amount_left to the amount flled in the fill mapping
		FillMap.Fill = FillMap.Fill.Add(amount_left)
		// subtract amount_left * price to the ReserveMapping
		ReserveMap.Reserves = ReserveMap.Reserves.Sub(amount_left.Mul(price))
		// set amount_left to 0
		amount_left = sdk.ZeroDec()
	}

	// Updates mappings of reserve and filledReserves based on the original limitOrderCurrentKey to the KVStore
	k.SetLimitOrderPoolReserveMap(ctx, ReserveMap)
	k.SetLimitOrderPoolFillMap(ctx, FillMap)

	//Updates limitOrderCurrentKey based on if any limitOrders were completely filled.
	k.SetTickMap(ctx, pairId, tick)

	_ = ctx

	return amount_left, amount_out, nil
}

///// Limit Order Functions

func (k Keeper) PlaceLimitOrderHelper(goCtx context.Context, pairId string, tickIndex int64) types.TickMap {
	ctx := sdk.UnwrapSDKContext(goCtx)

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

func (k Keeper) PlaceLimitOrderMappingHelper(goCtx context.Context, pairId string, tickIndex int64, tokenIn string, currentLimitOrderKey uint64, receiver string) (types.LimitOrderPoolReserveMap, types.LimitOrderPoolUserShareMap, types.LimitOrderPoolTotalSharesMap) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Retrieves ReserveMap Object from KVStore
	ReserveMap, ReserveMapFound := k.GetLimitOrderPoolReserveMap(ctx, pairId, tickIndex, tokenIn, currentLimitOrderKey)
	// Retrieves UserShareMap object from KVStore
	UserShareMap, UserShareMapFound := k.GetLimitOrderPoolUserShareMap(ctx, pairId, tickIndex, tokenIn, currentLimitOrderKey, receiver)
	// Retrives TotalSharesMap object from KVStore
	TotalSharesMap, TotalShareMapFound := k.GetLimitOrderPoolTotalSharesMap(ctx, pairId, tickIndex, tokenIn, currentLimitOrderKey)

	// If ReserveMap object not found initialize it accordingly
	if !ReserveMapFound {
		ReserveMap.Count = currentLimitOrderKey
		ReserveMap.Reserves = sdk.ZeroDec()
		ReserveMap.TickIndex = tickIndex
		ReserveMap.Token = tokenIn
		ReserveMap.PairId = pairId
	}

	// If UserShareMap object is not found initialize it accordingly
	if !UserShareMapFound {
		UserShareMap.Count = currentLimitOrderKey
		UserShareMap.Address = receiver
		UserShareMap.SharesOwned = sdk.ZeroDec()
		UserShareMap.TickIndex = tickIndex
		UserShareMap.Token = tokenIn
		UserShareMap.PairId = pairId
	}

	// If TotalSharesMap object is nout found initialize it accordingly
	if !TotalShareMapFound {
		TotalSharesMap.Count = currentLimitOrderKey
		TotalSharesMap.TotalShares = sdk.ZeroDec()
		TotalSharesMap.TickIndex = tickIndex
		TotalSharesMap.Token = tokenIn
		TotalSharesMap.PairId = pairId
	}

	_ = ctx

	return ReserveMap, UserShareMap, TotalSharesMap
}

func (k Keeper) PlaceLimitOrderCore(goCtx context.Context, msg *types.MsgPlaceLimitOrder, token0 string, token1 string, callerAddr sdk.AccAddress) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// checks if pair is initialized, if not intialize it and return pairId
	pairId, err := k.PairInit(goCtx, token0, token1, msg.TickIndex, 0)

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

	if tick.LimitOrderPool0To1.CurrentLimitOrderKey == 0 && tick.LimitOrderPool1To0.CurrentLimitOrderKey == 0 {
		pair.PairCount = pair.PairCount + 1
	}

	// If TokenIn == token0 set count and key to values of LimitOrderPool0to1
	if msg.TokenIn == token0 {
		currentLimitOrderKey = tick.LimitOrderPool0To1.CurrentLimitOrderKey
		LimitOrderCount = tick.LimitOrderPool0To1.Count

		// Errors if place a limit order in amount0 at a tick less than CurrentTick0to1
		if msg.TickIndex < pair.TokenPair.CurrentTick0To1 {
			return sdkerrors.Wrapf(types.ErrValidPairNotFound, "Cannot depsoit amount 0 at a tick less than the CurrentTick1to0")
		}

		// If tokenIn == token1 set count and key to values of LimitOrderPool1to0
	} else {
		currentLimitOrderKey = tick.LimitOrderPool1To0.CurrentLimitOrderKey
		LimitOrderCount = tick.LimitOrderPool1To0.Count

		// Errors if placing a limit order in amoount1 at a tick greater than Current1To0
		if msg.TickIndex > pair.TokenPair.CurrentTick1To0 {
			return sdkerrors.Wrapf(types.ErrValidPairNotFound, "Cannot depsoit amount 0 at a tick greater than the CurrentTick0to1")
		}
	}

	// Retrives FillMap object from FillMapping KVSTore
	FillMap, FillMapFound := k.GetLimitOrderPoolFillMap(ctx, pairId, msg.TickIndex, msg.TokenIn, currentLimitOrderKey)

	// inits FillMap object if not found
	if !FillMapFound {
		FillMap.PairId = pairId
		FillMap.TickIndex = msg.TickIndex
		FillMap.Token = msg.TokenIn
		FillMap.Count = currentLimitOrderKey
		FillMap.Fill = sdk.ZeroDec()

		// Handles creating a limit order given the current limit order is already partially filled
	} else if LimitOrderCount == currentLimitOrderKey && FillMap.Fill.GT(sdk.ZeroDec()) {
		// increments currentLimitOrderKey
		currentLimitOrderKey = currentLimitOrderKey + 1

		// Updates Pool0t01 / Pool1to0 accordingly
		if msg.TokenIn == token0 {
			tick.LimitOrderPool0To1.CurrentLimitOrderKey = currentLimitOrderKey
		} else {
			tick.LimitOrderPool1To0.CurrentLimitOrderKey = currentLimitOrderKey
		}
	}

	// Returns Map object for Reserve, UserShares, and TotalShares mapping
	ReserveMap, UserShareMap, TotalSharesMap := k.PlaceLimitOrderMappingHelper(goCtx, pairId, msg.TickIndex, msg.TokenIn, currentLimitOrderKey, msg.Receiver)

	// Adds amountIn to ReserveMap
	ReserveMap.Reserves = ReserveMap.Reserves.Add(msg.AmountIn)

	if msg.TokenIn == token0 {
		// calculates NewShares given amountIn in terms of token0
		newShares = msg.AmountIn
	} else {
		//Calculates price as amt1/amt0
		price, err := k.Calc_price(msg.TickIndex, false)

		if err != nil {
			return err
		}
		//calculates NewShares given amountIn is in term of token1
		newShares = msg.AmountIn.Mul(price)

	}

	// Adds newShares to User's shares owned
	UserShareMap.SharesOwned = UserShareMap.SharesOwned.Add(newShares)
	// Adds newShares to totalShares
	TotalSharesMap.TotalShares = TotalSharesMap.TotalShares.Add(newShares)

	// Set Fill, Reserve, UserShares, and TotalShares maps in KVStore
	k.SetLimitOrderPoolFillMap(ctx, FillMap)
	k.SetLimitOrderPoolReserveMap(ctx, ReserveMap)
	k.SetLimitOrderPoolUserShareMap(ctx, UserShareMap)
	k.SetLimitOrderPoolTotalSharesMap(ctx, TotalSharesMap)
	k.SetTickMap(ctx, pairId, tick)

	// If a new tick has been placed that tigtens the range between currentTick0to1 and currentTick0to1 update CurrentTicks to the tighest ticks
	// @Dev assumes that msg.amountIn > 0
	if msg.TokenIn == token0 && ((msg.TickIndex > pair.TokenPair.CurrentTick0To1) && (msg.TickIndex < pair.TokenPair.CurrentTick1To0)) {
		pair.TokenPair.CurrentTick1To0 = msg.TickIndex
	} else if (msg.TickIndex > pair.TokenPair.CurrentTick0To1) && (msg.TickIndex < pair.TokenPair.CurrentTick1To0) {
		pair.TokenPair.CurrentTick0To1 = msg.TickIndex
	}

	// updates currentTick1to0/0To1 given the conditionals above

	k.SetPairMap(ctx, pair)

	// Sends AmoutIn from Address to Module
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

func (k Keeper) CancelLimitOrderCore(goCtx context.Context, msg *types.MsgCancelLimitOrder, token0 string, token1 string, callerAddr sdk.AccAddress, receiverAddr sdk.AccAddress) error {

	return nil
}

func (k Keeper) WithdrawWithdrawnLimitOrderCore(goCtx context.Context, msg *types.MsgWithdrawFilledLimitOrder, token0 string, token1 string, callerAddr sdk.AccAddress, receiverAddr sdk.AccAddress) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// PairId for token0, token1 ("token0/token1")
	pairId := k.CreatePairId(token0, token1)
	// Retrives TickMap object from KVStore
	_, tickFound := k.GetTickMap(ctx, pairId, msg.TickIndex)

	// If tick does not exist, then there is no liqudity to withdraw and thus error
	if !tickFound {
		return sdkerrors.Wrapf(types.ErrValidTickNotFound, "Valid tick not found ")
	}

	// Retrives LimitOrderFillMap object from KVStore for the specified key and keyToken
	FillMap, FillMapFound := k.GetLimitOrderPoolFillMap(ctx, pairId, msg.TickIndex, msg.KeyToken, msg.Key)
	// Retrives LimitOrderReserevMap object from KVStore for the specified key and keyToken
	ReserveMap, ReserveMapFound := k.GetLimitOrderPoolReserveMap(ctx, pairId, msg.TickIndex, msg.KeyToken, msg.Key)
	// Retrives LimitOrderUserSharesMap object from KVStore for the specified key and keyToken
	UserShareMap, UserShareMapFound := k.GetLimitOrderPoolUserShareMap(ctx, pairId, msg.TickIndex, msg.KeyToken, msg.Key, msg.Creator)
	// Retrives LimitOrderUserSharesWithdrawnMap object from KVStore for the specified key and keyToken
	UserSharesWithdrawnMap, UserSharesWithdrawnFound := k.GetLimitOrderPoolUserSharesWithdrawn(ctx, pairId, msg.TickIndex, msg.KeyToken, msg.Key, msg.Creator)
	// Retrives LimitOrderTotalSharesMap object from KVStore for the specified key and keyToken
	TotalSharesMap, TotalShareMapFound := k.GetLimitOrderPoolTotalSharesMap(ctx, pairId, msg.TickIndex, msg.KeyToken, msg.Key)

	if !UserSharesWithdrawnFound {
		UserSharesWithdrawnMap = types.LimitOrderPoolUserSharesWithdrawn{
			PairId:          pairId,
			TickIndex:       msg.TickIndex,
			Token:           msg.KeyToken,
			Count:           msg.Key,
			Address:         msg.Creator,
			SharesWithdrawn: sdk.ZeroDec(),
		}
	}

	// If any of these maps are not found, then a valid withdraw option will not exist, and thus error
	if !FillMapFound || !UserShareMapFound || !TotalShareMapFound || !ReserveMapFound {
		return sdkerrors.Wrapf(types.ErrValidLimitOrderMapsNotFound, "Valid mappings for limit order withdraw not found")
	}

	if FillMap.Fill.Quo(FillMap.Fill.Add(ReserveMap.Reserves)).LTE(UserSharesWithdrawnMap.SharesWithdrawn.Quo(UserSharesWithdrawnMap.SharesWithdrawn.Add(UserShareMap.SharesOwned))) {
		return sdkerrors.Wrapf(types.ErrCannotWithdrawLimitOrder, "Cannot withdraw additional liqudity from this limit order at this time")

	}
	// Calculates the sharesOut based on the UserShares withdrawn  compared to sharesLeft compared to remaining liquidity in reserves
	sharesOut := ((FillMap.Fill.Mul(UserSharesWithdrawnMap.SharesWithdrawn.Add(UserShareMap.SharesOwned))).Quo(FillMap.Fill.Add(ReserveMap.Reserves))).Sub(UserSharesWithdrawnMap.SharesWithdrawn)

	amountOut := (sharesOut.Mul(FillMap.Fill)).Quo(TotalSharesMap.TotalShares)
	// Calculates amount to subtract from fillMap object given sharesOut
	FillMap.Fill = FillMap.Fill.Sub(sharesOut.Mul(FillMap.Fill).Quo(TotalSharesMap.TotalShares))
	// Updates useSharesWithdrawMap to include sharesOut
	UserSharesWithdrawnMap.SharesWithdrawn = UserSharesWithdrawnMap.SharesWithdrawn.Add(sharesOut)
	// Remove sharesOut from UserSharesMap
	UserShareMap.SharesOwned = UserShareMap.SharesOwned.Sub(sharesOut)
	// Removes sharesOut from TotalSharesMap
	// TODO: this wasn't in the spec but I assumed this is needed?
	// calculate amountOout given sharesOut

	//TotalSharesMap.TotalShares = TotalSharesMap.TotalShares.Sub(sharesOut)

	// Updates changed LimitOrder Mappings in KVstore
	k.SetLimitOrderPoolFillMap(ctx, FillMap)
	k.SetLimitOrderPoolUserShareMap(ctx, UserShareMap)
	k.SetLimitOrderPoolUserSharesWithdrawn(ctx, UserSharesWithdrawnMap)
	k.SetLimitOrderPoolTotalSharesMap(ctx, TotalSharesMap)

	var tokenOut string

	// determines in which token to withdraw amountOut into
	if msg.KeyToken == token0 {
		tokenOut = token1
	} else {
		tokenOut = token0
	}

	fmt.Println("Amount out: ", amountOut)
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
