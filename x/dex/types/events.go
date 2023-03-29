package types

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func createDepositEvent(
	creator,
	receiver,
	token0,
	token1,
	tickIndex,
	fee string,
	otherAttrs ...sdk.Attribute,
) sdk.Event {
	attrs := []sdk.Attribute{
		sdk.NewAttribute(sdk.AttributeKeyModule, "dex"),
		sdk.NewAttribute(sdk.AttributeKeyAction, DepositEventKey),
		sdk.NewAttribute(DepositEventCreator, creator),
		sdk.NewAttribute(DepositEventReceiver, receiver),
		sdk.NewAttribute(DepositEventToken0, token0),
		sdk.NewAttribute(DepositEventToken1, token1),
		sdk.NewAttribute(DepositEventPrice, tickIndex),
		sdk.NewAttribute(DepositEventFee, fee),
	}
	attrs = append(attrs, otherAttrs...)
	return sdk.NewEvent(sdk.EventTypeMessage, attrs...)
}

func CreateDepositEvent(
	creator,
	receiver,
	token0,
	token1,
	tickIndex,
	fee,
	oldReserve0,
	oldReserve1,
	newReserve0,
	newReserve1,
	sharesMinted string,
	otherAttrs ...sdk.Attribute,
) sdk.Event {
	return createDepositEvent(
		creator,
		receiver,
		token0,
		token1,
		tickIndex,
		fee,

		sdk.NewAttribute(DepositEventOldReserves0, oldReserve0),
		sdk.NewAttribute(DepositEventNewReserves0, newReserve0),
		sdk.NewAttribute(DepositEventOldReserves1, oldReserve1),
		sdk.NewAttribute(DepositEventNewReserves1, newReserve1),
		sdk.NewAttribute(DepositEventSharesMinted, sharesMinted),
	)
}

func createDepositFailedEvent(creator, receiver, token0, token1, tickIndex, fee string, otherAttrs ...sdk.Attribute) sdk.Event {
	attrs := []sdk.Attribute{
		sdk.NewAttribute(sdk.AttributeKeyModule, "dex"),
		sdk.NewAttribute(sdk.AttributeKeyAction, DepositFailEventKey),
		sdk.NewAttribute(DepositFailEventCreator, creator),
		sdk.NewAttribute(DepositFailEventReceiver, receiver),
		sdk.NewAttribute(DepositFailEventToken0, token0),
		sdk.NewAttribute(DepositFailEventToken1, token1),
		sdk.NewAttribute(DepositFailEventPrice, tickIndex),
		sdk.NewAttribute(DepositFailEventFee, fee),
	}
	attrs = append(attrs, otherAttrs...)
	return sdk.NewEvent(sdk.EventTypeMessage, attrs...)
}

func CreateDepositFailedEvent(
	creator,
	receiver,
	token0,
	token1,
	tickIndex,
	fee,
	oldReserve0,
	oldReserve1,
	amount0,
	amount1 string,
	otherAttrs ...sdk.Attribute,
) sdk.Event {
	return createDepositEvent(
		creator,
		receiver,
		token0,
		token1,
		tickIndex,
		fee,

		sdk.NewAttribute(DepositEventOldReserves0, oldReserve0),
		sdk.NewAttribute(DepositEventOldReserves1, oldReserve1),
		sdk.NewAttribute(DepositFailAmountToDeposit0, amount0),
		sdk.NewAttribute(DepositFailAmountToDeposit1, amount1),
	)
}

func createWithdrawEvent(
	creator,
	receiver,
	token0,
	token1,
	tickIndex,
	fee string,
	otherAttrs ...sdk.Attribute,
) sdk.Event {
	attrs := []sdk.Attribute{
		sdk.NewAttribute(sdk.AttributeKeyModule, "dex"),
		sdk.NewAttribute(sdk.AttributeKeyAction, WithdrawEventKey),
		sdk.NewAttribute(WithdrawEventCreator, creator),
		sdk.NewAttribute(WithdrawEventReceiver, receiver),
		sdk.NewAttribute(WithdrawEventToken0, token0),
		sdk.NewAttribute(WithdrawEventToken1, token1),
		sdk.NewAttribute(WithdrawEventPrice, tickIndex),
		sdk.NewAttribute(WithdrawEventFee, fee),
	}
	attrs = append(attrs, otherAttrs...)
	return sdk.NewEvent(sdk.EventTypeMessage, attrs...)
}

func CreateWithdrawEvent(
	creator,
	receiver,
	token0,
	token1,
	tickIndex,
	fee,
	oldReserve0,
	oldReserve1,
	newReserve0,
	newReserve1,
	sharesRemoved string,
	otherAttrs ...sdk.Attribute,
) sdk.Event {
	return createWithdrawEvent(
		creator,
		receiver,
		token0,
		token1,
		tickIndex,
		fee,
		sdk.NewAttribute(WithdrawEventOldReserves0, oldReserve0),
		sdk.NewAttribute(WithdrawEventOldReserves1, oldReserve1),
		sdk.NewAttribute(WithdrawEventNewReserves0, newReserve0),
		sdk.NewAttribute(WithdrawEventNewReserves1, newReserve1),
		sdk.NewAttribute(WithdrawEventSharesRemoved, sharesRemoved),
	)
}

func createSwapEvent(creator, receiver, tokenIn, tokenOut, amountIn, amountOut string) sdk.Event {
	attrs := []sdk.Attribute{
		sdk.NewAttribute(sdk.AttributeKeyModule, "dex"),
		sdk.NewAttribute(sdk.AttributeKeyAction, SwapEventKey),
		sdk.NewAttribute(SwapEventCreator, creator),
		sdk.NewAttribute(SwapEventReceiver, receiver),
		sdk.NewAttribute(SwapEventTokenIn, tokenIn),
		sdk.NewAttribute(SwapEventTokenOut, tokenOut),
		sdk.NewAttribute(SwapEventAmountIn, amountIn),
		sdk.NewAttribute(SwapEventAmoutOut, amountOut),
	}
	return sdk.NewEvent(sdk.EventTypeMessage, attrs...)
}

func CreateSwapEvent(creator, receiver, tokenIn, tokenOut, amountIn, amountOut string) sdk.Event {
	return createSwapEvent(
		creator,
		receiver,
		tokenIn,
		tokenOut,
		amountIn,
		amountOut,
	)
}

func createPlaceLimitOrderEvent(creator, receiver, tokenIn, tokenOut, amountIn, shares, trancheKey string) sdk.Event {
	attrs := []sdk.Attribute{
		sdk.NewAttribute(sdk.AttributeKeyModule, "dex"),
		sdk.NewAttribute(sdk.AttributeKeyAction, PlaceLimitOrderEventKey),
		sdk.NewAttribute(PlaceLimitOrderEventCreator, creator),
		sdk.NewAttribute(PlaceLimitOrderEventReceiver, receiver),
		sdk.NewAttribute(PlaceLimitOrderEventTokenIn, tokenIn),
		sdk.NewAttribute(PlaceLimitOrderEventTokenOut, tokenOut),
		sdk.NewAttribute(PlaceLimitOrderEventAmountIn, amountIn),
		sdk.NewAttribute(PlaceLimitOrderEventShares, shares),
		sdk.NewAttribute(PlaceLimitOrderEventTrancheKey, trancheKey),
	}
	return sdk.NewEvent(sdk.EventTypeMessage, attrs...)
}

func CreatePlaceLimitOrderEvent(creator, receiver, tokenIn, tokenOut, amountIn, shares, trancheKey string) sdk.Event {
	return createPlaceLimitOrderEvent(
		creator,
		receiver,
		tokenIn,
		tokenOut,
		amountIn,
		shares,
		trancheKey,
	)
}

func withdrawFilledLimitOrderEvent(creator, tokenIn, tokenOut, key, amountOut string) sdk.Event {
	attrs := []sdk.Attribute{
		sdk.NewAttribute(sdk.AttributeKeyModule, "dex"),
		sdk.NewAttribute(sdk.AttributeKeyAction, WithdrawFilledLimitOrderEventKey),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventCreator, creator),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventTokenIn, tokenIn),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventTokenOut, tokenOut),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventTrancheKey, key),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventAmountOut, amountOut),
	}
	return sdk.NewEvent(sdk.EventTypeMessage, attrs...)
}

func WithdrawFilledLimitOrderEvent(creator, tokenIn, tokenOut, key, amountOut string) sdk.Event {
	return withdrawFilledLimitOrderEvent(
		creator,
		tokenIn,
		tokenOut,
		key,
		amountOut,
	)
}

func GoodTilPurgeHitLimitEvent(gas sdk.Gas) sdk.Event {
	attrs := []sdk.Attribute{
		sdk.NewAttribute(sdk.AttributeKeyModule, "dex"),
		sdk.NewAttribute(sdk.AttributeKeyAction, GoodTilPurgeHitGasLimitEventKey),
		sdk.NewAttribute(GoodTilPurgeHitGasLimitEventGas, strconv.FormatUint(gas, 10)),
	}
	return sdk.NewEvent(sdk.EventTypeMessage, attrs...)
}

func cancelLimitOrderEvent(creator, tokenIn, tokenOut, key, amountOut string) sdk.Event {
	attrs := []sdk.Attribute{
		sdk.NewAttribute(sdk.AttributeKeyModule, "dex"),
		sdk.NewAttribute(sdk.AttributeKeyAction, WithdrawFilledLimitOrderEventKey),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventCreator, creator),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventTokenIn, tokenIn),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventTokenOut, tokenOut),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventTrancheKey, key),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventAmountOut, amountOut),
	}
	return sdk.NewEvent(sdk.EventTypeMessage, attrs...)
}

func CancelLimitOrderEvent(creator, tokenIn, tokenOut, key, amountOut string) sdk.Event {
	return cancelLimitOrderEvent(
		creator,
		tokenIn,
		tokenOut,
		key,
		amountOut,
	)
}
