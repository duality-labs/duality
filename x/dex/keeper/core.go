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

// Deposit & Withdraw Liquidity ////////////////////////////////////////////////////////////////////

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
	pair := k.GetOrInitPair(
		goCtx,
		token0,
		token1,
	)
	pairId := pair.PairId
	totalAmountReserve0 := sdk.ZeroDec()
	totalAmountReserve1 := sdk.ZeroDec()
	passedDeposit := 0
	amounts0Deposited := make([]sdk.Dec, len(amounts0))
	amounts1Deposited := make([]sdk.Dec, len(amounts1))
	for i := 0; i < len(amounts0); i++ {
		amounts0Deposited[i] = sdk.ZeroDec()
		amounts1Deposited[i] = sdk.ZeroDec()
	}

	feelist := k.GetAllFeeList(ctx)

	for i, amount0 := range amounts0 {
		amount1 := amounts1[i]
		tickIndex := msg.TickIndexes[i]
		price1To0 := k.CalcPrice1To0(tickIndex)
		feeIndex := msg.FeeIndexes[i]
		fee := feelist[feeIndex].Fee
		curTick0to1 := pair.TokenPair.CurrentTick0To1
		curTick1to0 := pair.TokenPair.CurrentTick1To0
		minTick := pair.MinTick
		maxTick := pair.MaxTick
		lowerTickIndex := tickIndex - fee
		upperTickIndex := tickIndex + fee

		// TODO: Allow user to deposit "behind enemy lines"
		if amounts0[i].GT(sdk.ZeroDec()) && curTick0to1 <= lowerTickIndex && maxTick != math.MinInt64 {
			return nil, nil, sdkerrors.Wrapf(types.ErrValidPairNotFound, "Cannot depsosit amount_0 at tick greater than or equal to the CurrentTick0to1")
		}

		// TODO: Allow user to deposit "behind enemy lines"
		if amounts1[i].GT(sdk.ZeroDec()) && upperTickIndex <= curTick1to0 && minTick != math.MaxInt64 {
			return nil, nil, sdkerrors.Wrapf(types.ErrValidPairNotFound, "Cannot deposit amount_1 at tick less than or equal to the CurrentTick1to0")
		}

		lowerTick := k.GetOrInitTick(goCtx, pairId, lowerTickIndex)
		upperTick := k.GetOrInitTick(goCtx, pairId, upperTickIndex)

		lowerReserve0 := &lowerTick.TickData.Reserve0AndShares[feeIndex].Reserve0
		lowerTotalShares := &lowerTick.TickData.Reserve0AndShares[feeIndex].TotalShares
		upperReserve1 := &upperTick.TickData.Reserve1[feeIndex]

		trueAmount0, trueAmount1, sharesMinted := CalcTrueAmounts(
			price1To0,
			*lowerReserve0,
			*upperReserve1,
			amount0,
			amount1,
			*lowerTotalShares,
		)

		if trueAmount0.Equal(sdk.ZeroDec()) && trueAmount1.Equal(sdk.ZeroDec()) {
			ctx.EventManager().EmitEvent(types.CreateDepositFailedEvent(
				msg.Creator,
				msg.Receiver,
				token0,
				token1,
				fmt.Sprint(tickIndex),
				fmt.Sprint(tickIndex),
				lowerReserve0.String(),
				upperReserve1.String(),
				amount0.String(),
				amount1.String(),
			))
			continue
		}

		// Add liquidity
		*lowerReserve0 = lowerReserve0.Add(trueAmount0)
		*lowerTotalShares = lowerTotalShares.Add(sharesMinted)
		*upperReserve1 = upperReserve1.Add(trueAmount1)
		k.SetPairMap(ctx, pair)
		k.SetTickMap(ctx, pairId, lowerTick)
		k.SetTickMap(ctx, pairId, upperTick)

		k.UpdateTickPointersPostAddToken0(goCtx, &pair, &lowerTick)
		k.UpdateTickPointersPostAddToken1(goCtx, &pair, &upperTick)

		amounts0Deposited[i] = trueAmount0
		amounts1Deposited[i] = trueAmount1

		passedDeposit++

		shares, sharesFound := k.GetShares(ctx, msg.Receiver, pairId, tickIndex, feeIndex)
		if !sharesFound {
			shares = types.Shares{
				Address:     msg.Receiver,
				PairId:      pairId,
				TickIndex:   tickIndex,
				FeeIndex:    feeIndex,
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
		ctx.EventManager().EmitEvent(types.CreateDepositEvent(
			msg.Creator,
			msg.Receiver,
			token0,
			token1,
			fmt.Sprint(msg.TickIndexes[i]),
			fmt.Sprint(msg.FeeIndexes[i]),
			lowerReserve0.Sub(trueAmount0).String(),
			lowerReserve0.Sub(trueAmount0).String(),
			upperReserve1.Sub(trueAmount1).String(),
			lowerReserve0.String(),
			upperReserve1.String(),
		),
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
	return amounts0Deposited, amounts1Deposited, nil
}

// Handles core logic for MsgWithdrawl; calculating and withdrawing reserve0,reserve1 from a specified tick given a specfied number of shares to remove.
// Calculates the amount of reserve0, reserve1 to withdraw based on the percetange of the desired number of shares to remove compared to the total number of shares at the given tick
func (k Keeper) WithdrawCore(goCtx context.Context, msg *types.MsgWithdrawl, token0 string, token1 string, callerAddr sdk.AccAddress, receiverAddr sdk.AccAddress) error {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pairId := k.CreatePairId(token0, token1)
	pair, found := k.GetPairMap(ctx, pairId)
	if !found {
		return sdkerrors.Wrapf(types.ErrValidPairNotFound, "Pair not found")
	}
	totalReserve0ToRemove := sdk.ZeroDec()
	totalReserve1ToRemove := sdk.ZeroDec()

	for i, feeIndex := range msg.FeeIndexes {
		sharesToRemove := msg.SharesToRemove[i]
		tickIndex := msg.TickIndexes[i]

		shareOwner, found := k.GetShares(
			ctx,
			msg.Creator,
			pairId,
			tickIndex,
			feeIndex,
		)
		if !found {
			return sdkerrors.Wrapf(types.ErrValidShareNotFound, "No valid share owner fonnd")
		}
		userSharesOwned := &shareOwner.SharesOwned

		feeValue, found := k.GetFeeList(ctx, feeIndex)
		if !found {
			return sdkerrors.Wrapf(types.ErrValidFeeIndexNotFound, "(%d) does not correspond to a valid fee", feeIndex)
		}
		fee := feeValue.Fee
		lowerTickIndex := tickIndex - fee
		upperTickIndex := tickIndex + fee
		lowerTick, lowerTickFound := k.GetTickMap(ctx, pairId, lowerTickIndex)
		upperTick, upperTickFound := k.GetTickMap(ctx, pairId, upperTickIndex)
		if !lowerTickFound || !upperTickFound {
			return sdkerrors.Wrapf(types.ErrValidTickNotFound, "No tick found at the requested index")
		}

		lowerTickFeeTotalShares := &lowerTick.TickData.Reserve0AndShares[feeIndex].TotalShares
		lowerTickFeeReserve0 := &lowerTick.TickData.Reserve0AndShares[feeIndex].Reserve0
		upperTickFeeReserve1 := &upperTick.TickData.Reserve1[feeIndex]

		// Checks to see if there are some totalShares to withdraw
		// In keeper/verification.go we check this condition for the msg.Creator, thus we know that they also has a valid position in the tick.

		if lowerTickFeeTotalShares.Equal(sdk.ZeroDec()) {
			return sdkerrors.Wrapf(types.ErrValidTickNotFound, "No tick found at the requested index")
		}

		// calculates the amount to withdraw of each token based on a ratio of the amountToRemove to
		// totalShares multiplied by the amount of the respective asset
		sharesToRemove = MinDec(sharesToRemove, *userSharesOwned)
		ownershipRatio := sharesToRemove.Quo(*lowerTickFeeTotalShares)
		reserve1ToRemove := ownershipRatio.Mul(*upperTickFeeReserve1)
		reserve0ToRemove := ownershipRatio.Mul(*lowerTickFeeReserve0)

		// Updates upper/lowerTick based on subtracting the calculated amount from the previous reserve0 and reserve1
		*lowerTickFeeReserve0 = lowerTickFeeReserve0.Sub(reserve0ToRemove)
		*upperTickFeeReserve1 = upperTickFeeReserve1.Sub(reserve1ToRemove)
		*lowerTickFeeTotalShares = lowerTickFeeTotalShares.Sub(sharesToRemove)
		*userSharesOwned = userSharesOwned.Sub(sharesToRemove)

		// adds reserve0ToRemove/reserve1ToRemove to totals
		totalReserve0ToRemove = totalReserve0ToRemove.Add(reserve0ToRemove)
		totalReserve1ToRemove = totalReserve1ToRemove.Add(reserve1ToRemove)

		// sets changes to tick mappings, and share mappings
		k.SetShares(ctx, shareOwner)
		k.SetTickMap(ctx, pairId, upperTick)
		k.SetTickMap(ctx, pairId, lowerTick)

		if totalReserve0ToRemove.GT(sdk.ZeroDec()) {
			k.UpdateTickPointersPostRemoveToken0(goCtx, &pair, &lowerTick)
		}

		if totalReserve1ToRemove.GT(sdk.ZeroDec()) {
			k.UpdateTickPointersPostRemoveToken1(goCtx, &pair, &upperTick)
		}

		// emits event for individiual withdrawl
		ctx.EventManager().EmitEvent(types.CreateWithdrawEvent(
			msg.Creator,
			msg.Receiver,
			token0,
			token1,
			fmt.Sprint(msg.TickIndexes[i]),
			fmt.Sprint(msg.FeeIndexes[i]),
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

	return nil
}

// SWAP FUNCTIONS //////////////////////////////////////////////////////////////

// Handles core logic for the asset 0 to asset1 direction of MsgSwap; faciliates swapping amount0 for some amount of amount1, given a specified pair (token0, token1)
func (k Keeper) Swap0to1(goCtx context.Context, msg *types.MsgSwap, token0 string, token1 string, callerAddr sdk.AccAddress) (sdk.Dec, sdk.Dec, error) {

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
		return sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrValidPairNotFound, "Pair not found")
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
			pair.TokenPair.CurrentTick0To1++
			continue
		}

		var i uint64 = 0

		// iterator for feeList
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
			price_0to1 := k.CalcPrice0To1(pair.TokenPair.CurrentTick0To1)

			// price * amout_left + amount_out < minOut, error we cannot meet minOut threshold
			if price_0to1.Mul(amount_left).Add(amount_out).LT(msg.MinOut) {
				return sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Amount Out is less than minium amount out specified: swap failed")
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
			k.UpdateTickPointersPostAddToken0(goCtx, &pair, &Current0Data)
		}

		k.SetTickMap(ctx, pairId, Current1Data)

		k.UpdateTickPointersPostRemoveToken1(goCtx, &pair, &Current1Data)

		// if feeIndex is equal to the largest index in feeList check for valid limit orders
		if i == feeSize {

			// assigns a new variable err to handle err not being initialized in this conditional
			var err error

			// runs swaps for any limitOrders at the specified tick, updating amount_left, amount_out accordingly
			// passes in the outToken (token1), as this is the direction of the limit order for which we check

			amount_left, amount_out, err = k.SwapLimitOrder0to1(goCtx, pairId, token1, amount_out, amount_left, pair.TokenPair.CurrentTick0To1)

			if err != nil {
				return sdk.ZeroDec(), sdk.ZeroDec(), err
			}
		}
	}

	// Make updates to pairMap containing updates CurrentTick0to1 to the KVStore
	k.SetPairMap(ctx, pair)

	// Check to see if amount_out meets the threshold of minOut
	if amount_out.LT(msg.MinOut) {
		return sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Amount Out is less than minium amount out specified: swap failed")
	}

	ctx.EventManager().EmitEvent(types.CreateSwapEvent(msg.Creator, msg.Receiver,
		token0, token1, msg.TokenIn, msg.AmountIn.String(), amount_out.String(), msg.MinOut.String(),
	))

	// Returns amount_out to keeper/msg.server: Swap
	// @Dev token transfers happen in keeper/msg.server: Swap
	return amount_out, amount_left, nil
}

// Handles core logic for the asset 1 to asset 0 direction of MsgSwap; faciliates swapping amount1 for some amount of amount0, given a specified pair (token0, token1)
func (k Keeper) Swap1to0(goCtx context.Context, msg *types.MsgSwap, token0 string, token1 string, callerAddr sdk.AccAddress) (sdk.Dec, sdk.Dec, error) {

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
		return sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrValidPairNotFound, "Pair not found")
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

		var i uint64 = 0
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
			price_1to0 := k.CalcPrice1To0(pair.TokenPair.CurrentTick1To0)

			// price * amout_left + amount_out < minOut, error we cannot meet minOut threshold
			if price_1to0.Mul(amount_left).Add(amount_out).LT(msg.MinOut) {
				return sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Amount Out is less than minium amount out specified: swap failed")
			}

			// If there is not enough to complete the trade
			if Current0Data.TickData.Reserve0AndShares[i].Reserve0.LT(amount_left.Mul(price_1to0)) {
				amount_out = amount_out.Add(Current0Data.TickData.Reserve0AndShares[i].Reserve0)
				amountInTemp := Current0Data.TickData.Reserve0AndShares[i].Reserve0.Quo(price_1to0)
				amount_left = amount_left.Sub(amountInTemp)
				Current1Data.TickData.Reserve1[i] = Current1Data.TickData.Reserve1[i].Add(amountInTemp)
				Current0Data.TickData.Reserve0AndShares[i].Reserve0 = sdk.ZeroDec()
			} else {
				amountOutTemp := amount_left.Mul(price_1to0)
				amount_out = amount_out.Add(amountOutTemp)
				Current0Data.TickData.Reserve0AndShares[i].Reserve0 = Current0Data.TickData.Reserve0AndShares[i].Reserve0.Sub(amountOutTemp)
				Current1Data.TickData.Reserve1[i] = Current1Data.TickData.Reserve1[i].Add(amount_left)
				amount_left = sdk.ZeroDec()
			}

			//updates feeIndex
			i++

			//Make updates to tickMap containing reserve0/1 data to the KVStore

			k.SetTickMap(ctx, pairId, Current1Data)

			k.UpdateTickPointersPostAddToken1(goCtx, &pair, &Current1Data)

		}

		k.SetTickMap(ctx, pairId, Current0Data)

		k.UpdateTickPointersPostRemoveToken0(goCtx, &pair, &Current0Data)
		// if feeIndex is equal to the largest index in feeList, check for valid limit orders at the specfied tick
		if i == feeSize {

			// assigns a new variable err to handle err not being initialized in this conditional
			var err error
			// runs swaps for any limitOrders at the specified tick, updating amount_left, amount_out accordingly

			// passes in the outToken (token0), as this is the direction of the limit order for which we check
			amount_left, amount_out, err = k.SwapLimitOrder1to0(goCtx, pairId, token0, amount_out, amount_left, pair.TokenPair.CurrentTick1To0)

			if err != nil {
				return sdk.ZeroDec(), sdk.ZeroDec(), err
			}
		}
	}

	// Check to see if amount_out meets the threshold of minOut
	k.SetPairMap(ctx, pair)

	if amount_out.LT(msg.MinOut) {
		return sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Amount Out is less than minium amount out specified: swap failed")
	}

	ctx.EventManager().EmitEvent(types.CreateSwapEvent(msg.Creator, msg.Receiver,
		token0, token1, msg.TokenIn, msg.AmountIn.String(), amount_out.String(), msg.MinOut.String(),
	))

	// Returns amount_out to keeper/msg.server: Swap
	// @Dev token transfers happen in keeper/msg.server: Swap
	return amount_out, amount_left, nil
}

// Swap Limit Orders

// Handles swapping asset 0 for asset 1 through any active limit orders at a specified tick
// Returns amount_out, amount_left, error
func (k Keeper) SwapLimitOrder0to1(goCtx context.Context, pairId string, tokenOut string, amountOut sdk.Dec, amountRemainingTokenIn sdk.Dec, tickIndex int64) (sdk.Dec, sdk.Dec, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// returns price for the given tick and specified direction (0 -> 1)
	priceInToOut := k.CalcPrice0To1(tickIndex)
	priceOutToIn := sdk.OneDec().Quo(priceInToOut)

	// Gets tick for specified tick at currentTick0to1
	tick, tickFound := k.GetTickMap(ctx, pairId, tickIndex)
	if !tickFound {
		return amountRemainingTokenIn, amountOut, nil
	}

	fillTranche := &tick.LimitOrderPool1To0.CurrentLimitOrderKey
	placeTranche := &tick.LimitOrderPool1To0.Count

	reserveData, found := k.GetLimitOrderPoolReserveMap(ctx, pairId, tickIndex, tokenOut, *fillTranche)
	if !found || reserveData.Reserves.Equal(sdk.ZeroDec()) {
		return amountRemainingTokenIn, amountOut, nil
	}
	reservesTokenOut := &reserveData.Reserves

	fillData, found := k.GetLimitOrderPoolFillMap(
		ctx,
		pairId,
		tickIndex,
		tokenOut,
		*fillTranche,
	)
	if !found {
		return sdk.ZeroDec(), sdk.ZeroDec(), fmt.Errorf("should not happen1")
	}
	fillTokenIn := &fillData.FilledReserves

	if reservesTokenOut.LTE(amountRemainingTokenIn.Mul(priceInToOut)) {
		amountOut = amountOut.Add(*reservesTokenOut)
		amountFilledTokenIn := reservesTokenOut.Mul(priceOutToIn)
		amountRemainingTokenIn = amountRemainingTokenIn.Sub(amountFilledTokenIn)
		*reservesTokenOut = sdk.ZeroDec()
		*fillTokenIn = fillTokenIn.Add(amountFilledTokenIn)
		k.SetLimitOrderPoolReserveMap(ctx, reserveData)
		k.SetLimitOrderPoolFillMap(ctx, fillData)

		if *fillTranche != *placeTranche {
			*fillTranche++
			k.SetTickMap(ctx, pairId, tick)

			reserveData, found = k.GetLimitOrderPoolReserveMap(ctx, pairId, tickIndex, tokenOut, *fillTranche)
			fillData, _ = k.GetLimitOrderPoolFillMap(ctx, pairId, tickIndex, tokenOut, *fillTranche)
			if !found {
				return sdk.ZeroDec(), sdk.ZeroDec(), fmt.Errorf("should not happen2")
			}
			reservesTokenOut = &reserveData.Reserves
			fillTokenIn = &fillData.FilledReserves

			if reservesTokenOut.LTE(amountRemainingTokenIn.Mul(priceInToOut)) {
				amountOut = amountOut.Add(*reservesTokenOut)
				amountFilledTokenIn := reservesTokenOut.Mul(priceOutToIn)
				amountRemainingTokenIn = amountRemainingTokenIn.Sub(amountFilledTokenIn)
				*reservesTokenOut = sdk.ZeroDec()
				*fillTokenIn = fillTokenIn.Add(amountFilledTokenIn)
			} else {
				amountFilledTokenOut := amountRemainingTokenIn.Mul(priceInToOut)
				amountOut = amountOut.Add(amountFilledTokenOut)
				*fillTokenIn = fillTokenIn.Add(amountRemainingTokenIn)
				*reservesTokenOut = reservesTokenOut.Sub(amountFilledTokenOut)
				amountRemainingTokenIn = sdk.ZeroDec()
			}
			k.SetLimitOrderPoolReserveMap(ctx, reserveData)
			k.SetLimitOrderPoolFillMap(ctx, fillData)
		}
	} else {
		amountFilledTokenOut := amountRemainingTokenIn.Mul(priceInToOut)
		amountOut = amountOut.Add(amountFilledTokenOut)
		fillData.FilledReserves = fillData.FilledReserves.Add(amountRemainingTokenIn)
		reserveData.Reserves = reserveData.Reserves.Sub(amountFilledTokenOut)
		amountRemainingTokenIn = sdk.ZeroDec()
		k.SetLimitOrderPoolReserveMap(ctx, reserveData)
		k.SetLimitOrderPoolFillMap(ctx, fillData)
	}

	pair, _ := k.GetPairMap(ctx, pairId)
	k.UpdateTickPointersPostRemoveToken1(goCtx, &pair, &tick)

	return amountRemainingTokenIn, amountOut, nil
}

// Handles swapping asset 1 for asset 0 through any active limit orders at a specified tick
// Returns amount_out, amount_left, error

func (k Keeper) SwapLimitOrder1to0(
	goCtx context.Context,
	pairId string,
	tokenOut string,
	amountOut sdk.Dec,
	amountRemainingTokenIn sdk.Dec,
	tickIndex int64,
) (sdk.Dec, sdk.Dec, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// returns price for the given tick and specified direction (0 -> 1)
	priceInToOut := k.CalcPrice1To0(tickIndex)
	priceOutToIn := sdk.OneDec().Quo(priceInToOut)

	// Gets tick for specified tick at currentTick0to1
	tick, tickFound := k.GetTickMap(ctx, pairId, tickIndex)
	if !tickFound {
		return amountRemainingTokenIn, amountOut, nil
	}

	fillTranche := &tick.LimitOrderPool0To1.CurrentLimitOrderKey
	placeTranche := &tick.LimitOrderPool0To1.Count

	reserveData, found := k.GetLimitOrderPoolReserveMap(ctx, pairId, tickIndex, tokenOut, *fillTranche)
	if !found || reserveData.Reserves.Equal(sdk.ZeroDec()) {
		return amountRemainingTokenIn, amountOut, nil
	}
	reservesTokenOut := &reserveData.Reserves

	fillData, found := k.GetLimitOrderPoolFillMap(
		ctx,
		pairId,
		tickIndex,
		tokenOut,
		*fillTranche,
	)
	if !found {
		return sdk.ZeroDec(), sdk.ZeroDec(), fmt.Errorf("should not happen3")
	}
	fillTokenIn := &fillData.FilledReserves

	if reservesTokenOut.LTE(amountRemainingTokenIn.Mul(priceInToOut)) {
		amountOut = amountOut.Add(*reservesTokenOut)
		amountFilledTokenIn := reservesTokenOut.Mul(priceOutToIn)
		amountRemainingTokenIn = amountRemainingTokenIn.Sub(amountFilledTokenIn)
		*reservesTokenOut = sdk.ZeroDec()
		*fillTokenIn = fillTokenIn.Add(amountFilledTokenIn)
		k.SetLimitOrderPoolReserveMap(ctx, reserveData)
		k.SetLimitOrderPoolFillMap(ctx, fillData)

		if *fillTranche != *placeTranche {
			*fillTranche++
			k.SetTickMap(ctx, pairId, tick)

			reserveData, found = k.GetLimitOrderPoolReserveMap(ctx, pairId, tickIndex, tokenOut, *fillTranche)
			fillData, _ = k.GetLimitOrderPoolFillMap(ctx, pairId, tickIndex, tokenOut, *fillTranche)
			if !found {
				return sdk.ZeroDec(), sdk.ZeroDec(), fmt.Errorf("should not happen4")
			}
			reservesTokenOut = &reserveData.Reserves
			fillTokenIn = &fillData.FilledReserves

			if reservesTokenOut.LTE(amountRemainingTokenIn.Mul(priceInToOut)) {
				amountOut = amountOut.Add(*reservesTokenOut)
				amountFilledTokenIn := reservesTokenOut.Mul(priceOutToIn)
				amountRemainingTokenIn = amountRemainingTokenIn.Sub(amountFilledTokenIn)
				*reservesTokenOut = sdk.ZeroDec()
				*fillTokenIn = fillTokenIn.Add(amountFilledTokenIn)
			} else {
				amountFilledTokenOut := amountRemainingTokenIn.Mul(priceInToOut)
				amountOut = amountOut.Add(amountFilledTokenOut)
				*fillTokenIn = fillTokenIn.Add(amountRemainingTokenIn)
				*reservesTokenOut = reservesTokenOut.Sub(amountFilledTokenOut)
				amountRemainingTokenIn = sdk.ZeroDec()
			}
			k.SetLimitOrderPoolReserveMap(ctx, reserveData)
			k.SetLimitOrderPoolFillMap(ctx, fillData)
		}
	} else {
		amountFilledTokenOut := amountRemainingTokenIn.Mul(priceInToOut)
		amountOut = amountOut.Add(amountFilledTokenOut)
		fillData.FilledReserves = fillData.FilledReserves.Add(amountRemainingTokenIn)
		reserveData.Reserves = reserveData.Reserves.Sub(amountFilledTokenOut)
		amountRemainingTokenIn = sdk.ZeroDec()
		k.SetLimitOrderPoolReserveMap(ctx, reserveData)
		k.SetLimitOrderPoolFillMap(ctx, fillData)
	}

	pair, _ := k.GetPairMap(ctx, pairId)
	k.UpdateTickPointersPostRemoveToken0(goCtx, &pair, &tick)

	return amountRemainingTokenIn, amountOut, nil
}

// Place, Cancel & Withdraw  Limit Order ////////////////////////////////////////////////////////////////////

// Handles MsgPlaceLimitOrder, initializing (tick, pair) data structures if needed, calculating and storing information for a new limit order at a specific tick
func (k Keeper) PlaceLimitOrderCore(goCtx context.Context, msg *types.MsgPlaceLimitOrder, token0 string, token1 string, callerAddr sdk.AccAddress) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// checks if pair is initialized, if not intialize it and return pairId

	pair := k.GetOrInitPair(goCtx, token0, token1)
	pairId := pair.PairId
	tick := k.GetOrInitTick(goCtx, pair.PairId, msg.TickIndex)

	var fillTrancheIndex *uint64
	var placeTrancheIndex *uint64

	if msg.TokenIn == token0 {
		if msg.TickIndex > pair.TokenPair.CurrentTick0To1 && pair.MaxTick != math.MinInt64 {
			return sdkerrors.Wrapf(types.ErrValidPairNotFound, "Cannot deposit amount 0 at a tick greater than the CurrentTick0to1")
		}
		fillTrancheIndex = &tick.LimitOrderPool0To1.CurrentLimitOrderKey
		placeTrancheIndex = &tick.LimitOrderPool0To1.Count
	} else {
		if msg.TickIndex < pair.TokenPair.CurrentTick1To0 && pair.MinTick != math.MaxInt64 {
			return sdkerrors.Wrapf(types.ErrValidPairNotFound, "Cannot deposit amount 1 at a tick less than the CurrentTick0to1")
		}
		fillTrancheIndex = &tick.LimitOrderPool1To0.CurrentLimitOrderKey
		placeTrancheIndex = &tick.LimitOrderPool1To0.Count
	}

	FillData := k.GetOrInitTickTrancheFillMap(goCtx, pairId, msg.TickIndex, *placeTrancheIndex, msg.TokenIn)
	ReserveData, UserShareData, TotalSharesData := k.GetOrInitLimitOrderMaps(goCtx, pairId, msg.TickIndex, msg.TokenIn, *fillTrancheIndex, msg.Receiver)
	if FillData.FilledReserves.GT(sdk.ZeroDec()) {
		*placeTrancheIndex++
		if FillData.FilledReserves.Equal(TotalSharesData.TotalShares) {
			*fillTrancheIndex++
		}
		k.SetTickMap(ctx, pairId, tick)
		FillData = k.GetOrInitTickTrancheFillMap(goCtx, pairId, msg.TickIndex, *placeTrancheIndex, msg.TokenIn)
		ReserveData, UserShareData, TotalSharesData = k.GetOrInitLimitOrderMaps(goCtx, pairId, msg.TickIndex, msg.TokenIn, *fillTrancheIndex, msg.Receiver)
	}

	// Adds amountIn to ReserveData
	ReserveData.Reserves = ReserveData.Reserves.Add(msg.AmountIn)

	// Adds newShares to User's shares owned
	UserShareData.SharesOwned = UserShareData.SharesOwned.Add(msg.AmountIn)

	// Adds newShares to totalShares
	TotalSharesData.TotalShares = TotalSharesData.TotalShares.Add(msg.AmountIn)

	// Set Fill, Reserve, UserShares, and TotalShares maps in KVStore
	k.SetLimitOrderPoolFillMap(ctx, FillData)
	k.SetLimitOrderPoolReserveMap(ctx, ReserveData)
	k.SetLimitOrderPoolUserShareMap(ctx, UserShareData)
	k.SetLimitOrderPoolTotalSharesMap(ctx, TotalSharesData)
	k.SetPairMap(ctx, pair)

	if msg.TokenIn == token0 {
		k.UpdateTickPointersPostAddToken0(goCtx, &pair, &tick)
	} else if msg.TokenIn == token1 {
		k.UpdateTickPointersPostAddToken1(goCtx, &pair, &tick)
	}

	// Sends AmountIn from Address to Module
	if msg.AmountIn.GT(sdk.ZeroDec()) {
		coin0 := sdk.NewCoin(msg.TokenIn, sdk.NewIntFromBigInt(msg.AmountIn.BigInt()))
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin0}); err != nil {
			return err
		}
	}

	ctx.EventManager().EmitEvent(types.CreatePlaceLimitOrderEvent(msg.Creator, msg.Receiver,
		token0, token1, msg.TokenIn, msg.AmountIn.String(), msg.AmountIn.String(), strconv.Itoa(int(*fillTrancheIndex)),
	))

	return nil
}

// Handles MsgCancelLimitOrder, removing a specifed number of shares from a limit order and returning the respective amount in terms of the reserve to the user
func (k Keeper) CancelLimitOrderCore(goCtx context.Context, msg *types.MsgCancelLimitOrder, token0 string, token1 string, callerAddr sdk.AccAddress, receiverAddr sdk.AccAddress) error {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// PairId for token0, token1 ("token0/token1")
	pairId := k.CreatePairId(token0, token1)
	// Retrives TickMap object from KVStore
	tick, tickFound := k.GetTickMap(ctx, pairId, msg.TickIndex)

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

	pair, _ := k.GetPairMap(ctx, pairId)
	if msg.KeyToken == token0 {
		k.UpdateTickPointersPostRemoveToken0(goCtx, &pair, &tick)
	} else {
		k.UpdateTickPointersPostRemoveToken1(goCtx, &pair, &tick)
	}

	return nil
}

// Handles MsgWithdrawFilledLimitOrder, calculates and sends filled liqudity from module to user for a limit order based on amount wished to receive.
func (k Keeper) WithdrawFilledLimitOrderCore(
	goCtx context.Context,
	msg *types.MsgWithdrawFilledLimitOrder,
	token0 string,
	token1 string,
	callerAddr sdk.AccAddress,
	receiverAddr sdk.AccAddress,
) error {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pairId := k.CreatePairId(token0, token1)
	tick, found := k.GetTickMap(ctx, pairId, msg.TickIndex)
	if !found {
		return sdkerrors.Wrapf(types.ErrValidTickNotFound, "Valid tick not found ")
	}

	orderTokenIn := msg.KeyToken
	var orderTokenOut string
	if msg.KeyToken == token0 {
		orderTokenOut = token1
	} else {
		orderTokenOut = token0
	}
	trancheIndex := msg.Key
	tickIndex := msg.TickIndex

	FillData, FillDataFound := k.GetLimitOrderPoolFillMap(ctx, pairId, tickIndex, orderTokenIn, trancheIndex)
	UserShareData, UserShareDataFound := k.GetLimitOrderPoolUserShareMap(ctx, pairId, tickIndex, orderTokenIn, trancheIndex, msg.Creator)
	TotalSharesData, TotalSharesDataFound := k.GetLimitOrderPoolTotalSharesMap(ctx, pairId, tickIndex, orderTokenIn, trancheIndex)
	if !FillDataFound || !UserShareDataFound || !TotalSharesDataFound {
		return sdkerrors.Wrapf(types.ErrValidLimitOrderMapsNotFound, "Valid mappings for limit order withdraw not found")
	}
	amountFilledLimitTokenOut := FillData.FilledReserves
	userSharesLimitTokenIn := &UserShareData.SharesOwned
	totalSharesLimitTokenIn := TotalSharesData.TotalShares

	UserSharesWithdrawnData, UserSharesWithdrawnDataFound := k.GetLimitOrderPoolUserSharesWithdrawn(
		ctx,
		pairId,
		tickIndex,
		orderTokenIn,
		trancheIndex,
		msg.Creator,
	)
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
	userSharesWithdrawnLimitTokenIn := &UserSharesWithdrawnData.SharesWithdrawn

	var priceLimitInToOut sdk.Dec
	if orderTokenIn == token0 {
		priceLimitInToOut = k.CalcPrice0To1(tick.TickIndex)
	} else {
		priceLimitInToOut = k.CalcPrice1To0(tick.TickIndex)
	}
	priceLimitOutToIn := sdk.OneDec().Quo(priceLimitInToOut)
	amountFilledLimitTokenIn := amountFilledLimitTokenOut.Mul(priceLimitOutToIn)
	ratioFilledTokenIn := amountFilledLimitTokenIn.Quo(totalSharesLimitTokenIn)
	sharesFilledTokenIn := userSharesLimitTokenIn.Mul(ratioFilledTokenIn)
	sharesToRemoveTokenIn := sharesFilledTokenIn.Sub(*userSharesWithdrawnLimitTokenIn)
	amountOut := sharesToRemoveTokenIn.Mul(priceLimitInToOut)

	*userSharesWithdrawnLimitTokenIn = userSharesWithdrawnLimitTokenIn.Add(sharesToRemoveTokenIn)
	k.SetLimitOrderPoolUserSharesWithdrawn(ctx, UserSharesWithdrawnData)

	// Sends amountOut from module address to msg.Receiver account address
	if amountOut.GT(sdk.ZeroDec()) {
		coinOut := sdk.NewCoin(orderTokenOut, sdk.NewIntFromBigInt(amountOut.BigInt()))
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
