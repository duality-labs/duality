package keeper

import (
	"context"
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/duality-labs/duality/x/dex/utils"
)

// NOTE: Currently we are using TruncateInt in multiple places for converting Decs back into sdk.Ints.
// This may create some accounting anomalies but seems preferable to other alternatives.
// See full ADR here: https://www.notion.so/dualityxyz/A-Modest-Proposal-For-Truncating-696a919d59254876a617f82fb9567895

// Handles core logic for MsgDeposit, checking and initializing data structures (tick, pair), calculating
// shares based on amount deposited, and sending funds to moduleAddress.
func (k Keeper) DepositCore(
	goCtx context.Context,
	token0 string,
	token1 string,
	callerAddr sdk.AccAddress,
	receiverAddr sdk.AccAddress,
	amounts0 []sdk.Int,
	amounts1 []sdk.Int,
	tickIndices []int64,
	fees []uint64,
	options []*types.DepositOptions,
) (amounts0Deposit, amounts1Deposit []sdk.Int, sharesIssued sdk.Coins, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pairID := CreatePairID(token0, token1)
	totalAmountReserve0 := sdk.ZeroInt()
	totalAmountReserve1 := sdk.ZeroInt()
	amounts0Deposited := make([]sdk.Int, len(amounts0))
	amounts1Deposited := make([]sdk.Int, len(amounts1))
	sharesIssued = sdk.Coins{}

	for i := 0; i < len(amounts0); i++ {
		amounts0Deposited[i] = sdk.ZeroInt()
		amounts1Deposited[i] = sdk.ZeroInt()
	}

	for i, amount0 := range amounts0 {
		amount1 := amounts1[i]
		tickIndex := tickIndices[i]
		fee := fees[i]
		autoswap := options[i].Autoswap

		feeUInt := utils.MustSafeUint64(fee)
		lowerTickIndex := tickIndex - feeUInt
		upperTickIndex := tickIndex + feeUInt
		// behind enemy lines checks
		// TODO: Allow user to deposit "behind enemy lines"
		if amount0.IsPositive() && k.IsBehindEnemyLines(ctx, pairID, pairID.Token0, lowerTickIndex) {
			return nil, nil, nil, types.ErrDepositBehindPairLiquidity
		}
		// TODO: Allow user to deposit "behind enemy lines"
		if amount1.IsPositive() && k.IsBehindEnemyLines(ctx, pairID, pairID.Token1, upperTickIndex) {
			return nil, nil, nil, types.ErrDepositBehindPairLiquidity
		}

		pool, err := k.GetOrInitPool(
			ctx,
			pairID,
			tickIndex,
			fee,
		)
		if err != nil {
			return nil, nil, nil, err
		}

		existingShares := k.bankKeeper.GetSupply(ctx, pool.GetDepositDenom()).Amount

		inAmount0, inAmount1, outShares := pool.Deposit(amount0, amount1, existingShares, autoswap)

		k.SavePool(ctx, pool)

		if inAmount0.IsZero() && inAmount1.IsZero() {
			return nil, nil, nil, types.ErrZeroTrueDeposit
		}

		if outShares.IsZero() {
			return nil, nil, nil, types.ErrDepositShareUnderflow
		}

		if err := k.MintShares(ctx, receiverAddr, outShares); err != nil {
			return nil, nil, nil, err
		}
		sharesIssued = sharesIssued.Add(outShares)

		amounts0Deposited[i] = inAmount0
		amounts1Deposited[i] = inAmount1
		totalAmountReserve0 = totalAmountReserve0.Add(inAmount0)
		totalAmountReserve1 = totalAmountReserve1.Add(inAmount1)

		ctx.EventManager().EmitEvent(types.CreateDepositEvent(
			callerAddr.String(),
			receiverAddr.String(),
			token0,
			token1,
			fmt.Sprint(tickIndices[i]),
			fmt.Sprint(fees[i]),
			pool.GetLowerReserve0().Sub(inAmount0).String(),
			pool.GetUpperReserve1().Sub(inAmount1).String(),
			pool.GetLowerReserve0().String(),
			pool.GetUpperReserve1().String(),
			outShares.String(),
		))
	}

	if totalAmountReserve0.IsPositive() {
		coin0 := sdk.NewCoin(token0, totalAmountReserve0)
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin0}); err != nil {
			return nil, nil, nil, err
		}
	}

	if totalAmountReserve1.IsPositive() {
		coin1 := sdk.NewCoin(token1, totalAmountReserve1)
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin1}); err != nil {
			return nil, nil, nil, err
		}
	}

	return amounts0Deposited, amounts1Deposited, sharesIssued, nil
}

// Handles core logic for MsgWithdrawal; calculating and withdrawing reserve0,reserve1 from a specified tick
// given a specfied number of shares to remove.
// Calculates the amount of reserve0, reserve1 to withdraw based on the percentage of the desired
// number of shares to remove compared to the total number of shares at the given tick.
func (k Keeper) WithdrawCore(
	goCtx context.Context,
	token0 string,
	token1 string,
	callerAddr sdk.AccAddress,
	receiverAddr sdk.AccAddress,
	sharesToRemoveList []sdk.Int,
	tickIndices []int64,
	fees []uint64,
) error {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pairID := CreatePairID(token0, token1)
	totalReserve0ToRemove := sdk.ZeroInt()
	totalReserve1ToRemove := sdk.ZeroInt()

	for i, fee := range fees {
		sharesToRemove := sharesToRemoveList[i]
		tickIndex := tickIndices[i]

		pool, err := k.GetOrInitPool(ctx, pairID, tickIndex, fee)
		if err != nil {
			return err
		}

		sharesID := NewDepositDenom(&types.PairID{Token0: token0, Token1: token1}, tickIndex, fee).String()
		totalShares := k.bankKeeper.GetSupply(ctx, sharesID).Amount

		if totalShares.LT(sharesToRemove) {
			return sdkerrors.Wrapf(
				types.ErrInsufficientShares,
				"%s does not have %s shares of type %s",
				callerAddr,
				sharesToRemove,
				sharesID,
			)
		}

		outAmount0, outAmount1 := pool.Withdraw(sharesToRemove, totalShares)
		k.SavePool(ctx, pool)

		if sharesToRemove.IsPositive() {
			if err := k.BurnShares(ctx, callerAddr, sharesToRemove, sharesID); err != nil {
				return err
			}
		}

		totalReserve0ToRemove = totalReserve0ToRemove.Add(outAmount0)
		totalReserve1ToRemove = totalReserve1ToRemove.Add(outAmount1)

		ctx.EventManager().EmitEvent(types.CreateWithdrawEvent(
			callerAddr.String(),
			receiverAddr.String(),
			token0,
			token1,
			fmt.Sprint(tickIndices[i]),
			fmt.Sprint(fees[i]),
			pool.LowerTick0.Reserves.Add(outAmount0).String(),
			pool.UpperTick1.Reserves.Add(outAmount1).String(),
			pool.LowerTick0.Reserves.String(),
			pool.UpperTick1.Reserves.String(),
			sharesToRemove.String(),
		))
	}

	if totalReserve0ToRemove.IsPositive() {
		coin0 := sdk.NewCoin(token0, totalReserve0ToRemove)

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

	// sends totalReserve1ToRemove to receiverAddr
	if totalReserve1ToRemove.IsPositive() {
		coin1 := sdk.NewCoin(token1, totalReserve1ToRemove)
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

// Handles core logic for the asset 0 to asset1 direction of MsgSwap;
// faciliates swapping amount0 for some amount of amount1, given a specified pair (token0, token1).
func (k Keeper) SwapCore(goCtx context.Context,
	tokenIn string,
	tokenOut string,
	amountIn sdk.Int,
	callerAddr sdk.AccAddress,
	receiverAddr sdk.AccAddress,
) (coinOut sdk.Coin, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pairID, err := CreatePairIDFromUnsorted(tokenIn, tokenOut)
	if err != nil {
		return sdk.Coin{}, err
	}

	coinIn, coinOut, err := k.SwapWithCache(ctx, pairID, tokenIn, tokenOut, amountIn, nil)
	if err != nil {
		return sdk.Coin{}, err
	}

	if coinOut.IsZero() {
		return sdk.Coin{}, types.ErrInsufficientLiquidity
	}

	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coinIn}); err != nil {
		return sdk.Coin{}, err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx,
		types.ModuleName,
		receiverAddr,
		sdk.Coins{coinOut})
	if err != nil {
		return sdk.Coin{}, err
	}

	ctx.EventManager().EmitEvent(types.CreateSwapEvent(callerAddr.String(), receiverAddr.String(),
		tokenIn, tokenOut, amountIn.String(), coinOut.Amount.String()))

	return coinOut, nil
}

func (k Keeper) MultiHopSwapCore(
	goCtx context.Context,
	amountIn sdk.Int,
	routes []*types.MultiHopRoute,
	exitLimitPrice sdk.Dec,
	pickBestRoute bool,
	callerAddr sdk.AccAddress,
	receiverAddr sdk.AccAddress,
) (coinOut sdk.Coin, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var routeErrors []error
	initialInCoin := sdk.NewCoin(routes[0].Hops[0], amountIn)
	stepCache := make(map[multihopCacheKey]StepResult)
	var bestRoute struct {
		write   func()
		coinOut sdk.Coin
		route   []string
	}
	bestRoute.coinOut = sdk.Coin{Amount: sdk.ZeroInt()}

	for _, route := range routes {
		routeCoinOut, writeRoute, err := k.RunMultihopRoute(ctx, *route, initialInCoin, exitLimitPrice, stepCache)
		if err != nil {
			routeErrors = append(routeErrors, err)
			continue
		}

		if !pickBestRoute || bestRoute.coinOut.Amount.LT(routeCoinOut.Amount) {
			bestRoute.coinOut = routeCoinOut
			bestRoute.write = writeRoute
			bestRoute.route = route.Hops
		}
		if !pickBestRoute {
			break
		}
	}

	if len(routeErrors) == len(routes) {
		// All routes have failed

		allErr := utils.JoinErrors(types.ErrAllMultiHopRoutesFailed, routeErrors...)

		return sdk.Coin{}, allErr
	}

	bestRoute.write()
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{initialInCoin})
	if err != nil {
		return sdk.Coin{}, err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, sdk.Coins{bestRoute.coinOut})
	if err != nil {
		return sdk.Coin{}, err
	}
	ctx.EventManager().EmitEvent(types.CreateMultihopSwapEvent(callerAddr.String(), receiverAddr.String(),
		initialInCoin.String(), bestRoute.coinOut.String(), strings.Join(bestRoute.route, ",")))

	return coinOut, nil
}

// Handles MsgPlaceLimitOrder, initializing (tick, pair) data structures if needed, calculating and
// storing information for a new limit order at a specific tick.
func (k Keeper) PlaceLimitOrderCore(
	goCtx context.Context,
	tokenIn string,
	tokenOut string,
	amountIn sdk.Int,
	tickIndex int64,
	orderType types.LimitOrderType,
	goodTil *time.Time,
	callerAddr sdk.AccAddress,
	receiverAddr sdk.AccAddress,
) (trancheKeyP *string, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pairID, err := CreatePairIDFromUnsorted(tokenIn, tokenOut)
	if err != nil {
		return nil, err
	}

	placeTranche, err := k.GetOrInitPlaceTranche(ctx, pairID, tokenIn, tickIndex, goodTil, orderType)
	if err != nil {
		return nil, err
	}

	trancheKey := placeTranche.TrancheKey
	trancheUser := k.GetOrInitLimitOrderTrancheUser(
		ctx,
		pairID,
		tickIndex,
		tokenIn,
		trancheKey,
		orderType,
		receiverAddr.String(),
	)

	amountLeft, totalIn := amountIn, sdk.ZeroInt()
	// For everything except just-in-time (JIT) orders try to execute as a swap first
	if !orderType.IsJIT() {
		limitPrice := placeTranche.PriceMakerToTaker().ToDec()
		coinInSwap, coinOutSwap, err := k.SwapWithCache(
			ctx,
			pairID,
			tokenIn,
			tokenOut,
			amountIn,
			&limitPrice,
		)
		if err != nil {
			return nil, err
		}

		totalIn = coinInSwap.Amount
		amountLeft = amountLeft.Sub(coinInSwap.Amount)

		trancheUser.TakerReserves = coinOutSwap.Amount
	}

	if amountLeft.IsPositive() && orderType.IsFoK() {
		return nil, types.ErrFoKLimitOrderNotFilled
	}

	sharesIssued := sdk.ZeroInt()
	// FOR GTC, JIT & GoodTil try to place a maker limitOrder with remaining Amount
	if amountLeft.IsPositive() && (orderType.IsGTC() || orderType.IsJIT() || orderType.IsGoodTil()) {
		placeTranche.PlaceMakerLimitOrder(amountLeft)
		trancheUser.SharesOwned = trancheUser.SharesOwned.Add(amountLeft)

		if orderType.HasExpiration() {
			goodTilRecord := NewLimitOrderExpiration(placeTranche)
			k.SetLimitOrderExpiration(ctx, goodTilRecord)
			ctx.GasMeter().ConsumeGas(types.ExpiringLimitOrderGas, "Expiring LimitOrder Fee")
		}

		k.SaveTranche(ctx, placeTranche)

		totalIn = totalIn.Add(amountLeft)
		sharesIssued = amountLeft
	}

	k.SaveTrancheUser(ctx, trancheUser)

	if totalIn.IsPositive() {
		coin0 := sdk.NewCoin(tokenIn, totalIn)

		err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin0})
		if err != nil {
			return nil, err
		}
	}

	ctx.EventManager().EmitEvent(types.CreatePlaceLimitOrderEvent(
		callerAddr.String(),
		receiverAddr.String(),
		tokenIn,
		tokenOut,
		totalIn.String(),
		sharesIssued.String(),
		trancheKey,
	))

	return &trancheKey, nil
}

// Handles MsgCancelLimitOrder, removing a specified number of shares from a limit order
// and returning the respective amount in terms of the reserve to the user.
func (k Keeper) CancelLimitOrderCore(
	goCtx context.Context,
	trancheKey string,
	callerAddr sdk.AccAddress,
) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	trancheUser, foundTrancheUser := k.GetLimitOrderTrancheUser(ctx, callerAddr.String(), trancheKey)
	if !foundTrancheUser {
		return types.ErrActiveLimitOrderNotFound
	}

	pairID, tickIndex, tokenIn := trancheUser.PairID, trancheUser.TickIndex, trancheUser.Token
	tranche, foundTranche := k.GetLimitOrderTranche(ctx, pairID, tokenIn, tickIndex, trancheKey)
	if !foundTranche {
		return types.ErrActiveLimitOrderNotFound
	}

	amountToCancel := tranche.RemoveTokenIn(trancheUser)
	trancheUser.SharesCancelled = trancheUser.SharesCancelled.Add(amountToCancel)

	if amountToCancel.IsPositive() {
		coinOut := sdk.NewCoin(tokenIn, amountToCancel)

		err := k.bankKeeper.SendCoinsFromModuleToAccount(
			ctx,
			types.ModuleName,
			callerAddr,
			sdk.Coins{coinOut},
		)
		if err != nil {
			return err
		}

		k.SaveTrancheUser(ctx, trancheUser)
		k.SaveTranche(ctx, *tranche)

		if trancheUser.OrderType.HasExpiration() {
			k.RemoveLimitOrderExpiration(ctx, *tranche.ExpirationTime, tranche.Ref())
		}
	} else {
		return sdkerrors.Wrapf(types.ErrCancelEmptyLimitOrder, "%s", tranche.TrancheKey)
	}

	ctx.EventManager().EmitEvent(types.CancelLimitOrderEvent(
		callerAddr.String(),
		tokenIn,
		pairID.MustOppositeToken(tokenIn),
		trancheKey,
		amountToCancel.String(),
	))

	return nil
}

// Handles MsgWithdrawFilledLimitOrder, calculates and sends filled liqudity from module to user
// for a limit order based on amount wished to receive.
func (k Keeper) WithdrawFilledLimitOrderCore(
	goCtx context.Context,
	trancheKey string,
	callerAddr sdk.AccAddress,
) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	trancheUser, foundTrancheUser := k.GetLimitOrderTrancheUser(
		ctx,
		callerAddr.String(),
		trancheKey,
	)
	if !foundTrancheUser {
		return sdkerrors.Wrapf(types.ErrValidLimitOrderTrancheNotFound, "%s", trancheKey)
	}

	pairID, tickIndex, tokenIn := trancheUser.PairID, trancheUser.TickIndex, trancheUser.Token

	tranche, wasFilled, foundTranche := k.FindLimitOrderTranche(ctx, pairID, tickIndex, tokenIn, trancheKey)

	amountOutTokenOut := sdk.ZeroDec()
	remainingTokenIn := sdk.ZeroInt()
	// It's possible that a TrancheUser exists but tranche does not if LO was filled entirely through a swap
	if foundTranche {
		var amountOutTokenIn sdk.Int
		amountOutTokenIn, amountOutTokenOut = tranche.Withdraw(trancheUser)

		if wasFilled {
			// This is only relevant for inactive JIT and GoodTil limit orders
			remainingTokenIn = tranche.RemoveTokenIn(trancheUser)
			k.SaveInactiveTranche(ctx, tranche)
		} else {
			k.SetLimitOrderTranche(ctx, tranche)
		}

		trancheUser.SharesWithdrawn = trancheUser.SharesWithdrawn.Add(amountOutTokenIn)
	}

	takerReserves := trancheUser.WithdrawTakerReserves()
	amountOutTokenOut = amountOutTokenOut.Add(takerReserves.ToDec())

	k.SaveTrancheUser(ctx, trancheUser)

	tokenOut := pairID.MustOppositeToken(tokenIn)
	if amountOutTokenOut.IsPositive() || remainingTokenIn.IsPositive() {
		coinOut := sdk.NewCoin(tokenOut, amountOutTokenOut.TruncateInt())
		coinInRefund := sdk.NewCoin(tokenIn, remainingTokenIn)
		coins := sdk.NewCoins(coinOut, coinInRefund)
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, callerAddr, coins); err != nil {
			return err
		}
	} else {
		return types.ErrWithdrawEmptyLimitOrder
	}

	ctx.EventManager().EmitEvent(types.WithdrawFilledLimitOrderEvent(
		callerAddr.String(),
		tokenIn,
		tokenOut,
		trancheKey,
		amountOutTokenOut.String(),
	))

	return nil
}
