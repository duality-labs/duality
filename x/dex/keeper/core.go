package keeper

import (
	"context"
	"fmt"
	"math"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) DepositPairInit(goCtx context.Context, token0 string, token1 string, tick_index int64, fee int64) (string, error) {

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

	price, err := k.Calc_price(tickIndex)

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

// FIX ME
// While this works this s a rough get around to doing float based math with sdk.Decs
func (k Keeper) Calc_price(price_Index int64) (sdk.Dec, error) {
	floatPrice := math.Pow(1.0001, float64(price_Index))
	sPrice := fmt.Sprintf("%f", floatPrice)

	price, err := sdk.NewDecFromStr(sPrice)

	if err != nil {
		return sdk.ZeroDec(), err
	} else {
		return price, nil
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
	pairId, err := k.DepositPairInit(goCtx, token0, token1, msg.TickIndexes[0], feelist[msg.FeeIndexes[0]].Fee)

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

func (k Keeper) WithdrawCore(goCtx context.Context, msg *types.MsgWithdrawl, token0 string, token1 string, callerAddr sdk.AccAddress) error {
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
		if isTickEmpty {
			pair.PairCount = pair.PairCount - 1
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
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, callerAddr, sdk.Coins{coin0}); err != nil {
			return err
		}
	}

	if totalReserve1ToRemove.GT(sdk.ZeroDec()) {
		coin1 := sdk.NewCoin(token1, sdk.NewIntFromBigInt(totalReserve1ToRemove.BigInt()))
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, callerAddr, sdk.Coins{coin1}); err != nil {
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
	// amount to return to receiver
	amount_out := sdk.ZeroDec()

	// verify that amount left is not zero and that there are additional valid ticks to check
	for !amount_left.Equal(sdk.ZeroDec()) && (count < int(pair.PairCount)) {

		// Tick data for tick that holds information about reserve1
		Current1Data, Current1Found := k.GetTickMap(ctx, pairId, pair.TokenPair.CurrentTick0To1)

		fmt.Println(count)
		if !Current1Found {
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
			price, err := k.Calc_price(pair.TokenPair.CurrentTick0To1)

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
		if i == feeSize-1 {
			// iterates CurrentTick0to1
			pair.TokenPair.CurrentTick0To1 = pair.TokenPair.CurrentTick0To1 + 1
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

		Current0Data, Current0Found := k.GetTickMap(ctx, pairId, pair.TokenPair.CurrentTick1To0)
		//Current0Datam := Current0Data.TickData.Reserve1[i]

		// If tick/feeIndex pair is not found continue

		if !Current0Found {

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
			price, err := k.Calc_price(pair.TokenPair.CurrentTick1To0)
			price = sdk.OneDec().Quo(price)

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
		if i == feeSize-1 {

			pair.TokenPair.CurrentTick0To1 = pair.TokenPair.CurrentTick0To1 - 1
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

///// Limit Order Functions

func (k Keeper) PlaceLimitOrderCore(goCtx context.Context, msg *types.MsgPlaceLimitOrder) error {

	return nil
}

func (k Keeper) CancelLimitOrderCore(goCtx context.Context, msg *types.MsgCancelLimitOrder) error {

	return nil
}

func (k Keeper) WithdrawFilledLimitOrderCore(goCtx context.Context, msg *types.MsgWithdrawlFilledLimitOrder) error {

	return nil
}
