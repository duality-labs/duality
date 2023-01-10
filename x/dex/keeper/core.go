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
		price1To0 := CalcPrice1To0(tickIndex)
		feeIndex := msg.FeeIndexes[i]
		fee := feelist[feeIndex].Fee
		curTick0to1 := pair.TokenPair.CurrentTick0To1
		curTick1to0 := pair.TokenPair.CurrentTick1To0
		lowerTickIndex := tickIndex - fee
		upperTickIndex := tickIndex + fee

		// TODO: Allow user to deposit "behind enemy lines"
		if amount0.GT(sdk.ZeroDec()) && curTick0to1 <= lowerTickIndex {
			return nil, nil, types.ErrDepositBehindPairLiquidity
		}

		// TODO: Allow user to deposit "behind enemy lines"
		if amount1.GT(sdk.ZeroDec()) && upperTickIndex <= curTick1to0 {
			return nil, nil, types.ErrDepositBehindPairLiquidity
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

		k.SetShares(ctx, shares)

		totalAmountReserve0 = totalAmountReserve0.Add(trueAmount0)
		totalAmountReserve1 = totalAmountReserve1.Add(trueAmount1)

		ctx.EventManager().EmitEvent(types.CreateDepositEvent(
			msg.Creator,
			msg.Receiver,
			token0,
			token1,
			fmt.Sprint(msg.TickIndexes[i]),
			fmt.Sprint(msg.FeeIndexes[i]),
			lowerReserve0.Sub(trueAmount0).String(),
			upperReserve1.Sub(trueAmount1).String(),
			lowerReserve0.String(),
			upperReserve1.String(),
			sharesMinted.String(),
		),
		)
	}

	if passedDeposit == 0 {
		return nil, nil, sdkerrors.Wrapf(types.ErrAllDepositsFailed, "All deposits failed")
	}

	if totalAmountReserve0.GT(sdk.ZeroDec()) {
		coin0 := sdk.NewCoin(token0, totalAmountReserve0.RoundInt())
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin0}); err != nil {
			return nil, nil, err
		}
	}

	if totalAmountReserve1.GT(sdk.ZeroDec()) {
		coin1 := sdk.NewCoin(token1, totalAmountReserve1.RoundInt())
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
		return types.ErrValidPairNotFound
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
			return types.ErrValidShareNotFound
		}
		userSharesOwned := &shareOwner.SharesOwned

		feeValue, found := k.GetFeeList(ctx, feeIndex)
		if !found {
			return types.ErrValidFeeIndexNotFound
		}
		fee := feeValue.Fee
		lowerTickIndex := tickIndex - fee
		upperTickIndex := tickIndex + fee
		lowerTick, lowerTickFound := k.GetTickMap(ctx, pairId, lowerTickIndex)
		upperTick, upperTickFound := k.GetTickMap(ctx, pairId, upperTickIndex)
		if !lowerTickFound || !upperTickFound {
			return types.ErrValidTickNotFound
		}

		lowerTickFeeTotalShares := &lowerTick.TickData.Reserve0AndShares[feeIndex].TotalShares
		lowerTickFeeReserve0 := &lowerTick.TickData.Reserve0AndShares[feeIndex].Reserve0
		upperTickFeeReserve1 := &upperTick.TickData.Reserve1[feeIndex]
		if lowerTickFeeTotalShares.Equal(sdk.ZeroDec()) {
			return types.ErrNotEnoughShares
		}

		sharesToRemove = MinDec(sharesToRemove, *userSharesOwned)
		ownershipRatio := sharesToRemove.Quo(*lowerTickFeeTotalShares)
		reserve1ToRemove := ownershipRatio.Mul(*upperTickFeeReserve1)
		reserve0ToRemove := ownershipRatio.Mul(*lowerTickFeeReserve0)

		*lowerTickFeeReserve0 = lowerTickFeeReserve0.Sub(reserve0ToRemove)
		*upperTickFeeReserve1 = upperTickFeeReserve1.Sub(reserve1ToRemove)
		*lowerTickFeeTotalShares = lowerTickFeeTotalShares.Sub(sharesToRemove)
		*userSharesOwned = userSharesOwned.Sub(sharesToRemove)

		totalReserve0ToRemove = totalReserve0ToRemove.Add(reserve0ToRemove)
		totalReserve1ToRemove = totalReserve1ToRemove.Add(reserve1ToRemove)

		k.SetShares(ctx, shareOwner)
		k.SetTickMap(ctx, pairId, upperTick)
		k.SetTickMap(ctx, pairId, lowerTick)

		if totalReserve0ToRemove.GT(sdk.ZeroDec()) {
			k.UpdateTickPointersPostRemoveToken0(goCtx, &pair, &lowerTick)
		}

		if totalReserve1ToRemove.GT(sdk.ZeroDec()) {
			k.UpdateTickPointersPostRemoveToken1(goCtx, &pair, &upperTick)
		}

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
			sharesToRemove.String(),
		))
	}
	k.SetPairMap(ctx, pair)
	if totalReserve0ToRemove.GT(sdk.ZeroDec()) {
		coin0 := sdk.NewCoin(token0, totalReserve0ToRemove.RoundInt())
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
		coin1 := sdk.NewCoin(token1, totalReserve1ToRemove.RoundInt())
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

// Handles core logic for the asset 0 to asset1 direction of MsgSwap; faciliates swapping amount0 for some amount of amount1, given a specified pair (token0, token1)
func (k Keeper) Swap0to1(goCtx context.Context, msg *types.MsgSwap, token0 string, token1 string, callerAddr sdk.AccAddress) (sdk.Dec, sdk.Dec, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pairId := k.CreatePairId(token0, token1)
	feeSize := k.GetFeeListCount(ctx)
	feelist := k.GetAllFeeList(ctx)
	pair, pairFound := k.GetPairMap(ctx, pairId)
	if !pairFound {
		return sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrValidPairNotFound, "Pair not found")
	}
	if pair.TokenPair.CurrentTick0To1 == math.MaxInt64 {
		return sdk.ZeroDec(), sdk.ZeroDec(), types.ErrNotEnoughLiquidity
	}

	amount_left := msg.AmountIn
	amount_out := sdk.ZeroDec()

	// verify that amount left is not zero and that there are additional valid ticks to check
	// for !amount_left.Equal(sdk.ZeroDec()) && pair.TokenPair.CurrentTick0To1 <= pair.MaxTick {
	for !amount_left.Equal(sdk.ZeroDec()) && pair.TokenPair.CurrentTick0To1 <= pair.MaxTick {
		Current1Data, Current1Found := k.GetTickMap(ctx, pairId, pair.TokenPair.CurrentTick0To1)
		if !Current1Found {
			pair.TokenPair.CurrentTick0To1++
			continue
		}

		var i uint64 = 0

		for i < feeSize && !amount_left.Equal(sdk.ZeroDec()) {
			fee := feelist[i].Fee
			Current0Data, found := k.GetTickMap(ctx, pairId, pair.TokenPair.CurrentTick0To1-2*fee)
			if !found {
				i++
				continue
			}

			price_0to1 := CalcPrice0To1(pair.TokenPair.CurrentTick0To1)

			if price_0to1.Mul(amount_left).Add(amount_out).LT(msg.MinOut) {
				return sdk.ZeroDec(), sdk.ZeroDec(), types.ErrNotEnoughLiquidity
			}

			if Current1Data.TickData.Reserve1[i].LT(amount_left.Mul(price_0to1)) {
				amount_out = amount_out.Add(Current1Data.TickData.Reserve1[i])
				amountInTemp := Current1Data.TickData.Reserve1[i].Quo(price_0to1)
				amount_left = amount_left.Sub(amountInTemp)
				Current0Data.TickData.Reserve0AndShares[i].Reserve0 = Current0Data.TickData.Reserve0AndShares[i].Reserve0.Add(amountInTemp)
				Current1Data.TickData.Reserve1[i] = sdk.ZeroDec()
			} else {
				amountOutTemp := amount_left.Mul(price_0to1)
				amount_out = amount_out.Add(amountOutTemp)
				Current0Data.TickData.Reserve0AndShares[i].Reserve0 = Current0Data.TickData.Reserve0AndShares[i].Reserve0.Add(amount_left)
				Current1Data.TickData.Reserve1[i] = Current1Data.TickData.Reserve1[i].Sub(amountOutTemp)
				amount_left = sdk.ZeroDec()
			}

			i++

			//Make updates to tickMap containing reserve0/1 data to the KVStore
			k.SetTickMap(ctx, pairId, Current0Data)
			// TODO: Return to this, maybe is receiving the wrong tick
			k.UpdateTickPointersPostAddToken0(goCtx, &pair, &Current0Data)
		}

		k.SetTickMap(ctx, pairId, Current1Data)
		if i == feeSize && amount_left.GT(sdk.ZeroDec()) {
			var err error
			amount_left, amount_out, err = k.SwapLimitOrder0to1(goCtx, pairId, token1, amount_out, amount_left, pair.TokenPair.CurrentTick0To1)
			// err = fmt.Errorf("dummy error for testing")
			if err != nil {
				return sdk.ZeroDec(), sdk.ZeroDec(), err
			}
		}
		k.UpdateTickPointersPostRemoveToken1(goCtx, &pair, &Current1Data)
	}

	k.SetPairMap(ctx, pair)

	// Check to see if amount_out meets the threshold of minOut
	if amount_out.LT(msg.MinOut) {
		return sdk.ZeroDec(), sdk.ZeroDec(), types.ErrNotEnoughLiquidity
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
	pairId := k.CreatePairId(token0, token1)
	feeSize := k.GetFeeListCount(ctx)
	feelist := k.GetAllFeeList(ctx)
	pair, found := k.GetPairMap(ctx, pairId)
	if !found {
		return sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrValidPairNotFound, "Pair not found")
	}
	if pair.TokenPair.CurrentTick1To0 == math.MinInt64 {
		return sdk.ZeroDec(), sdk.ZeroDec(), types.ErrNotEnoughLiquidity
	}

	amount_left := msg.AmountIn
	amount_out := sdk.ZeroDec()
	for !amount_left.Equal(sdk.ZeroDec()) && pair.TokenPair.CurrentTick1To0 >= pair.MinTick {

		Current0Data, Current0Found := k.GetTickMap(ctx, pairId, pair.TokenPair.CurrentTick1To0)
		if !Current0Found {
			pair.TokenPair.CurrentTick1To0 = pair.TokenPair.CurrentTick1To0 - 1
			continue
		}

		var i uint64 = 0
		for i < feeSize && !amount_left.Equal(sdk.ZeroDec()) {
			fee := feelist[i].Fee

			Current1Data, found := k.GetTickMap(ctx, pairId, pair.TokenPair.CurrentTick1To0+2*fee)
			if !found {
				i++
				continue
			}

			price_1to0 := CalcPrice1To0(pair.TokenPair.CurrentTick1To0)
			if price_1to0.Mul(amount_left).Add(amount_out).LT(msg.MinOut) {
				return sdk.ZeroDec(), sdk.ZeroDec(), types.ErrNotEnoughLiquidity
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

			i++

			k.SetTickMap(ctx, pairId, Current1Data)
			k.UpdateTickPointersPostAddToken1(goCtx, &pair, &Current1Data)
		}

		k.SetTickMap(ctx, pairId, Current0Data)

		if i == feeSize && amount_left.GT(sdk.ZeroDec()) {
			var err error
			amount_left, amount_out, err = k.SwapLimitOrder1to0(goCtx, pairId, token0, amount_out, amount_left, pair.TokenPair.CurrentTick1To0)

			if err != nil {
				return sdk.ZeroDec(), sdk.ZeroDec(), err
			}
		}
		k.UpdateTickPointersPostRemoveToken0(goCtx, &pair, &Current0Data)
	}

	k.SetPairMap(ctx, pair)

	if amount_out.LT(msg.MinOut) {
		return sdk.ZeroDec(), sdk.ZeroDec(), types.ErrNotEnoughLiquidity
	}

	ctx.EventManager().EmitEvent(types.CreateSwapEvent(msg.Creator, msg.Receiver,
		token0, token1, msg.TokenIn, msg.AmountIn.String(), amount_out.String(), msg.MinOut.String(),
	))

	return amount_out, amount_left, nil
}

// Handles swapping asset 0 for asset 1 through any active limit orders at a specified tick
// Returns amount_out, amount_left, error
func (k Keeper) SwapLimitOrder0to1(
	goCtx context.Context,
	pairId string,
	tokenOut string,
	amountOut sdk.Dec,
	amountRemainingTokenIn sdk.Dec,
	tickIndex int64,
) (newAmountRemainingTokenIn sdk.Dec, newAmountOut sdk.Dec, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	priceInToOut := CalcPrice0To1(tickIndex)
	priceOutToIn := sdk.OneDec().Quo(priceInToOut)

	tick, tickFound := k.GetTickMap(ctx, pairId, tickIndex)
	if !tickFound {
		return amountRemainingTokenIn, amountOut, nil
	}

	fillTranche := &tick.LimitOrderTranche1To0.FillTrancheIndex
	placeTranche := &tick.LimitOrderTranche1To0.PlaceTrancheIndex

	for amountRemainingTokenIn.GT(sdk.ZeroDec()) && *fillTranche < *placeTranche {
		amountRemainingTokenIn, amountOut, err = k.SwapLimitOrderTranche(
			goCtx,
			pairId,
			tokenOut,
			amountOut,
			amountRemainingTokenIn,
			tickIndex,
			*fillTranche,
			priceInToOut,
			priceOutToIn,
		)
		if err != nil {
			return sdk.ZeroDec(), sdk.ZeroDec(), err
		}
		if !k.TickTrancheHasToken0(ctx, &tick, *fillTranche) {
			*fillTranche++
			k.SetTickMap(ctx, pairId, tick)
		}
	}

	if amountRemainingTokenIn.GT(sdk.ZeroDec()) {
		amountRemainingTokenIn, amountOut, err = k.SwapLimitOrderTranche(
			goCtx,
			pairId,
			tokenOut,
			amountOut,
			amountRemainingTokenIn,
			tickIndex,
			*fillTranche,
			priceInToOut,
			priceOutToIn,
		)
		if err != nil {
			return sdk.ZeroDec(), sdk.ZeroDec(), err
		}
	}

	pair, _ := k.GetPairMap(ctx, pairId)
	k.UpdateTickPointersPostRemoveToken0(goCtx, &pair, &tick)

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
) (newAmountRemainingTokenIn sdk.Dec, newAmountOut sdk.Dec, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	priceInToOut := CalcPrice1To0(tickIndex)
	priceOutToIn := sdk.OneDec().Quo(priceInToOut)

	tick, tickFound := k.GetTickMap(ctx, pairId, tickIndex)
	if !tickFound {
		return amountRemainingTokenIn, amountOut, nil
	}

	fillTranche := &tick.LimitOrderTranche0To1.FillTrancheIndex
	placeTranche := &tick.LimitOrderTranche0To1.PlaceTrancheIndex

	for amountRemainingTokenIn.GT(sdk.ZeroDec()) && *fillTranche < *placeTranche {
		amountRemainingTokenIn, amountOut, err = k.SwapLimitOrderTranche(
			goCtx,
			pairId,
			tokenOut,
			amountOut,
			amountRemainingTokenIn,
			tickIndex,
			*fillTranche,
			priceInToOut,
			priceOutToIn,
		)
		if err != nil {
			return sdk.ZeroDec(), sdk.ZeroDec(), err
		}
		if !k.TickTrancheHasToken1(ctx, &tick, *fillTranche) {
			*fillTranche++
			k.SetTickMap(ctx, pairId, tick)
		}
	}

	if amountRemainingTokenIn.GT(sdk.ZeroDec()) {
		amountRemainingTokenIn, amountOut, err = k.SwapLimitOrderTranche(
			goCtx,
			pairId,
			tokenOut,
			amountOut,
			amountRemainingTokenIn,
			tickIndex,
			*fillTranche,
			priceInToOut,
			priceOutToIn,
		)
		if err != nil {
			return sdk.ZeroDec(), sdk.ZeroDec(), err
		}
	}

	pair, _ := k.GetPairMap(ctx, pairId)
	k.UpdateTickPointersPostRemoveToken0(goCtx, &pair, &tick)

	return amountRemainingTokenIn, amountOut, nil
}

func (k Keeper) SwapLimitOrderTranche(
	goCtx context.Context,
	pairId string,
	tokenOut string,
	amountOut sdk.Dec,
	amountRemainingTokenIn sdk.Dec,
	tickIndex int64,
	trancheIndex uint64,
	priceInToOut sdk.Dec,
	priceOutToIn sdk.Dec,
) (newAmountRemainingTokenIn sdk.Dec, newAmountOut sdk.Dec, error error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	tranche, found := k.GetLimitOrderTranche(ctx, pairId, tickIndex, tokenOut, trancheIndex)
	if !found {
		return amountRemainingTokenIn, amountOut, nil
	}
	reservesTokenOut := &tranche.ReservesTokenIn
	fillTokenIn := &tranche.ReservesTokenOut
	totalTokenIn := &tranche.TotalTokenOut

	if reservesTokenOut.LTE(amountRemainingTokenIn.Mul(priceInToOut)) {
		amountOut = amountOut.Add(*reservesTokenOut)
		amountFilledTokenIn := reservesTokenOut.Mul(priceOutToIn)
		amountRemainingTokenIn = amountRemainingTokenIn.Sub(amountFilledTokenIn)
		*reservesTokenOut = sdk.ZeroDec()
		*fillTokenIn = fillTokenIn.Add(amountFilledTokenIn)
		*totalTokenIn = totalTokenIn.Add(amountFilledTokenIn)
	} else {
		amountFilledTokenOut := amountRemainingTokenIn.Mul(priceInToOut)
		amountOut = amountOut.Add(amountFilledTokenOut)
		*fillTokenIn = fillTokenIn.Add(amountRemainingTokenIn)
		*totalTokenIn = totalTokenIn.Add(amountRemainingTokenIn)
		*reservesTokenOut = reservesTokenOut.Sub(amountFilledTokenOut)
		amountRemainingTokenIn = sdk.ZeroDec()
	}
	k.SetLimitOrderTranche(ctx, tranche)

	return amountRemainingTokenIn, amountOut, nil
}

// Handles MsgPlaceLimitOrder, initializing (tick, pair) data structures if needed, calculating and storing information for a new limit order at a specific tick
func (k Keeper) PlaceLimitOrderCore(goCtx context.Context, msg *types.MsgPlaceLimitOrder, token0 string, token1 string, callerAddr sdk.AccAddress) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pair := k.GetOrInitPair(goCtx, token0, token1)
	pairId := pair.PairId
	tick := k.GetOrInitTick(goCtx, pair.PairId, msg.TickIndex)

	tickIndex := msg.TickIndex
	tokenIn := msg.TokenIn
	receiver := msg.Receiver

	var fillTrancheIndex *uint64
	var placeTrancheIndex *uint64

	if msg.TokenIn == token0 {
		if msg.TickIndex > pair.TokenPair.CurrentTick0To1 {
			return types.ErrPlaceLimitOrderBehindPairLiquidity
		}
		fillTrancheIndex = &tick.LimitOrderTranche0To1.FillTrancheIndex
		placeTrancheIndex = &tick.LimitOrderTranche0To1.PlaceTrancheIndex
	} else {
		if msg.TickIndex < pair.TokenPair.CurrentTick1To0 {
			return types.ErrPlaceLimitOrderBehindPairLiquidity
		}
		fillTrancheIndex = &tick.LimitOrderTranche1To0.FillTrancheIndex
		placeTrancheIndex = &tick.LimitOrderTranche1To0.PlaceTrancheIndex
	}

	tranche := k.GetOrInitLimitOrderTranche(ctx, pairId, tickIndex, tokenIn, *placeTrancheIndex)
	trancheUser := k.GetOrInitLimitOrderTrancheUser(goCtx, pairId, tickIndex, tokenIn, *placeTrancheIndex, receiver)
	if tranche.ReservesTokenIn.LT(tranche.TotalTokenIn) {
		*placeTrancheIndex++
		k.SetTickMap(ctx, pairId, tick)
		tranche = k.GetOrInitLimitOrderTranche(ctx, pairId, tickIndex, tokenIn, *placeTrancheIndex)
		trancheUser = k.GetOrInitLimitOrderTrancheUser(goCtx, pairId, tickIndex, tokenIn, *placeTrancheIndex, receiver)
	}
	tranche.ReservesTokenIn = tranche.ReservesTokenIn.Add(msg.AmountIn)
	tranche.TotalTokenIn = tranche.TotalTokenIn.Add(msg.AmountIn)
	trancheUser.SharesOwned = trancheUser.SharesOwned.Add(msg.AmountIn)

	k.SetLimitOrderTrancheUser(ctx, trancheUser)
	k.SetLimitOrderTranche(ctx, tranche)
	k.SetPairMap(ctx, pair)

	if msg.TokenIn == token0 {
		k.UpdateTickPointersPostAddToken0(goCtx, &pair, &tick)
	} else if msg.TokenIn == token1 {
		k.UpdateTickPointersPostAddToken1(goCtx, &pair, &tick)
	}

	if msg.AmountIn.GT(sdk.ZeroDec()) {
		coin0 := sdk.NewCoin(msg.TokenIn, msg.AmountIn.RoundInt())
		err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin0})
		if err != nil {
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

	attemptedSharesOut := msg.SharesOut
	pairId := k.CreatePairId(token0, token1)

	tick, tickFound := k.GetTickMap(ctx, pairId, msg.TickIndex)
	if !tickFound {
		return sdkerrors.Wrapf(types.ErrValidTickNotFound, "Valid tick not found ")
	}

	trancheUser, found := k.GetLimitOrderTrancheUser(ctx, pairId, msg.TickIndex, msg.KeyToken, msg.Key, msg.Creator)
	if !found {
		return types.ErrValidLimitOrderMapsNotFound
	}

	tranche, found := k.GetLimitOrderTranche(ctx, pairId, msg.TickIndex, msg.KeyToken, msg.Key)
	if !found {
		return types.ErrValidLimitOrderMapsNotFound
	}

	var priceLimitOutToIn sdk.Dec
	if msg.KeyToken == token0 {
		priceLimitOutToIn = CalcPrice1To0(msg.TickIndex)
	} else {
		priceLimitOutToIn = CalcPrice0To1(msg.TickIndex)
	}

	ratioNotFilled := tranche.TotalTokenIn.Sub(tranche.TotalTokenOut.Mul(priceLimitOutToIn)).Quo(tranche.TotalTokenIn)
	maxUserAllowedToCancel := trancheUser.SharesOwned.Mul(ratioNotFilled)
	totalUserAttemptingToCancel := trancheUser.SharesCancelled.Add(attemptedSharesOut)

	if totalUserAttemptingToCancel.GT(maxUserAllowedToCancel) {
		return sdkerrors.Wrapf(types.ErrCannotWithdrawLimitOrder, "sharesOut is larger than shares Owned at the specified tick")
	}

	if totalUserAttemptingToCancel.Add(trancheUser.SharesWithdrawn).GT(trancheUser.SharesOwned) {
		return sdkerrors.Wrapf(types.ErrCannotWithdrawLimitOrder, "sharesOut is larger than shares Owned at the specified tick")
	}

	trancheUser.SharesCancelled = trancheUser.SharesCancelled.Add(attemptedSharesOut)
	k.SetLimitOrderTrancheUser(ctx, trancheUser)

	tranche.ReservesTokenIn = tranche.ReservesTokenIn.Sub(attemptedSharesOut)
	k.SetLimitOrderTranche(ctx, tranche)

	if attemptedSharesOut.GT(sdk.ZeroDec()) {
		coinOut := sdk.NewCoin(msg.KeyToken, attemptedSharesOut.RoundInt())
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, sdk.Coins{coinOut}); err != nil {
			return err
		}
	} else {
		return sdkerrors.Wrapf(types.ErrCannotWithdrawLimitOrder, "Cannot cancel liqudity from this limit order at this time")
	}

	ctx.EventManager().EmitEvent(types.CancelLimitOrderEvent(msg.Creator, msg.Receiver,
		token0, token1, msg.KeyToken, strconv.Itoa(int(msg.Key)), attemptedSharesOut.String(),
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

	orderTokenIn := msg.KeyToken
	var orderTokenOut string
	if msg.KeyToken == token0 {
		orderTokenOut = token1
	} else {
		orderTokenOut = token0
	}
	trancheIndex := msg.Key
	tickIndex := msg.TickIndex

	tranche, found := k.GetLimitOrderTranche(ctx, pairId, tickIndex, orderTokenIn, trancheIndex)
	if !found {
		return types.ErrValidLimitOrderMapsNotFound
	}

	trancheUser, found := k.GetLimitOrderTrancheUser(
		ctx,
		pairId,
		tickIndex,
		orderTokenIn,
		trancheIndex,
		msg.Creator,
	)
	if !found {
		return types.ErrValidLimitOrderMapsNotFound
	}

	tick, found := k.GetTickMap(ctx, pairId, msg.TickIndex)
	if !found {
		return sdkerrors.Wrapf(types.ErrValidTickNotFound, "Valid tick not found ")
	}

	var priceLimitInToOut sdk.Dec
	var priceLimitOutToIn sdk.Dec
	if orderTokenIn == token0 {
		priceLimitInToOut = CalcPrice0To1(tick.TickIndex)
	} else {
		priceLimitInToOut = CalcPrice1To0(tick.TickIndex)
	}
	priceLimitOutToIn = sdk.OneDec().Quo(priceLimitInToOut)

	ratioFilled := tranche.TotalTokenOut.Mul(priceLimitOutToIn).Quo(tranche.TotalTokenIn)
	maxAllowedToWithdraw := MinDec(
		trancheUser.SharesOwned.Mul(ratioFilled),                 // cannot withdraw more than what's been filled
		trancheUser.SharesOwned.Sub(trancheUser.SharesCancelled), // cannot withdraw more than what you own
	)
	amountOutTokenIn := maxAllowedToWithdraw.Sub(trancheUser.SharesWithdrawn)

	amountOutTokenOut := amountOutTokenIn.Mul(priceLimitInToOut)

	trancheUser.SharesWithdrawn = maxAllowedToWithdraw
	k.SetLimitOrderTrancheUser(ctx, trancheUser)

	tranche.ReservesTokenOut = tranche.ReservesTokenOut.Sub(amountOutTokenOut)
	k.SetLimitOrderTranche(ctx, tranche)

	if amountOutTokenOut.GT(sdk.ZeroDec()) {
		coinOut := sdk.NewCoin(orderTokenOut, amountOutTokenOut.RoundInt())
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, sdk.Coins{coinOut}); err != nil {
			return err
		}
	} else {
		return sdkerrors.Wrapf(types.ErrCannotWithdrawLimitOrder, "Cannot withdraw additional liqudity from this limit order at this time")
	}

	ctx.EventManager().EmitEvent(types.WithdrawFilledLimitOrderEvent(msg.Creator, msg.Receiver,
		token0, token1, msg.KeyToken, strconv.Itoa(int(msg.Key)), amountOutTokenOut.String(),
	))

	return nil
}
