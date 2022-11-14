package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func createDepositEvent(creator string, receiver string, token0 string, token1 string, tickIndex string, feeIndex string, otherAttrs ...sdk.Attribute) sdk.Event {
	attrs := []sdk.Attribute{
		sdk.NewAttribute(sdk.AttributeKeyModule, "dex"),
		sdk.NewAttribute(sdk.AttributeKeyAction, DepositEventKey),
		sdk.NewAttribute(DepositEventCreator, creator),
		sdk.NewAttribute(DepositEventReceiver, receiver),
		sdk.NewAttribute(DepositEventToken0, token0),
		sdk.NewAttribute(DepositEventToken1, token1),
		sdk.NewAttribute(DepositEventPrice, tickIndex),
		sdk.NewAttribute(DepositEventFeeIndex, feeIndex),
	}
	attrs = append(attrs, otherAttrs...)
	return sdk.NewEvent(sdk.EventTypeMessage, attrs...)
}

func CreateDepositEvent(creator string, receiver string, token0 string, token1 string, tickIndex string, feeIndex string, tokenDirection string, oldReserve0 string, oldReserve1 string, newReserve0 string, newReserve1 string, otherAttrs ...sdk.Attribute) sdk.Event {
	return createDepositEvent(
		creator,
		receiver,
		token0,
		token1,
		tickIndex,
		feeIndex,

		sdk.NewAttribute(DepositEventOldReserves0, oldReserve0),
		sdk.NewAttribute(DepositEventNewReserves0, newReserve0),
		sdk.NewAttribute(DepositEventOldReserves1, oldReserve1),
		sdk.NewAttribute(DepositEventNewReserves1, newReserve1),
	)
}

func createDepositFailedEvent(creator string, receiver string, token0 string, token1 string, tickIndex string, feeIndex string, otherAttrs ...sdk.Attribute) sdk.Event {
	attrs := []sdk.Attribute{
		sdk.NewAttribute(sdk.AttributeKeyModule, "dex"),
		sdk.NewAttribute(sdk.AttributeKeyAction, DepositFailEventKey),
		sdk.NewAttribute(DepositFailEventCreator, creator),
		sdk.NewAttribute(DepositFailEventReceiver, receiver),
		sdk.NewAttribute(DepositFailEventToken0, token0),
		sdk.NewAttribute(DepositFailEventToken1, token1),
		sdk.NewAttribute(DepositFailEventPrice, tickIndex),
		sdk.NewAttribute(DepositFailEventFeeIndex, feeIndex),
	}
	attrs = append(attrs, otherAttrs...)
	return sdk.NewEvent(sdk.EventTypeMessage, attrs...)
}

func CreateDepositFailedEvent(creator string, receiver string, token0 string, token1 string, tickIndex string, feeIndex string, oldReserve0 string, oldReserve1 string, amount0 string, amount1 string, otherAttrs ...sdk.Attribute) sdk.Event {
	return createDepositEvent(
		creator,
		receiver,
		token0,
		token1,
		tickIndex,
		feeIndex,

		sdk.NewAttribute(DepositEventOldReserves0, oldReserve0),
		sdk.NewAttribute(DepositEventOldReserves1, oldReserve1),
		sdk.NewAttribute(DepositFailAmountToDeposit0, amount0),
		sdk.NewAttribute(DepositFailAmountToDeposit1, amount1),
	)
}

func createWithdrawEvent(creator string, receiver string, token0 string, token1 string, tickIndex string, feeIndex string, otherAttrs ...sdk.Attribute) sdk.Event {
	attrs := []sdk.Attribute{
		sdk.NewAttribute(sdk.AttributeKeyModule, "dex"),
		sdk.NewAttribute(sdk.AttributeKeyAction, WithdrawEventKey),
		sdk.NewAttribute(WithdrawEventCreator, creator),
		sdk.NewAttribute(WithdrawEventReceiver, receiver),
		sdk.NewAttribute(WithdrawEventToken0, token0),
		sdk.NewAttribute(WithdrawEventToken1, token1),
		sdk.NewAttribute(WithdrawEventPrice, tickIndex),
		sdk.NewAttribute(WithdrawEventFee, feeIndex),
	}
	attrs = append(attrs, otherAttrs...)
	return sdk.NewEvent(sdk.EventTypeMessage, attrs...)
}

func CreateWithdrawEvent(creator string, receiver string, token0 string, token1 string, tickIndex string, feeIndex string, oldReserve0 string, oldReserve1 string, newReserve0 string, newReserve1 string, otherAttrs ...sdk.Attribute) sdk.Event {
	return createWithdrawEvent(
		creator,
		receiver,
		token0,
		token1,
		tickIndex,
		feeIndex,
		sdk.NewAttribute(WithdrawEventOldReserves0, oldReserve0),
		sdk.NewAttribute(WithdrawEventOldReserves1, oldReserve1),
		sdk.NewAttribute(WithdrawEventNewReserves0, newReserve0),
		sdk.NewAttribute(WithdrawEventNewReserves1, newReserve1),
	)
}

func createSwapEvent(creator string, receiver string, token0 string, token1 string, tokenIn string, amountIn string, amountOut string, minOut string, otherAttrs ...sdk.Attribute) sdk.Event {
	attrs := []sdk.Attribute{
		sdk.NewAttribute(sdk.AttributeKeyModule, "dex"),
		sdk.NewAttribute(sdk.AttributeKeyAction, SwapEventKey),
		sdk.NewAttribute(SwapEventCreator, creator),
		sdk.NewAttribute(SwapEventReceiver, receiver),
		sdk.NewAttribute(SwapEventToken0, token0),
		sdk.NewAttribute(SwapEventToken1, token1),
		sdk.NewAttribute(SwapEventTokenIn, tokenIn),
		sdk.NewAttribute(SwapEventAmountIn, amountIn),
		sdk.NewAttribute(SwapEventAmoutOut, amountOut),
		sdk.NewAttribute(SwapEventMinOut, minOut),
	}
	attrs = append(attrs, otherAttrs...)
	return sdk.NewEvent(sdk.EventTypeMessage, attrs...)
}

func CreateSwapEvent(creator string, receiver string, token0 string, token1 string, tokenIn string, amountIn string, amountOut string, minOut string, otherAttrs ...sdk.Attribute) sdk.Event {
	return createSwapEvent(
		creator,
		receiver,
		token0,
		token1,
		tokenIn,
		amountIn,
		amountOut,
		minOut,
	)
}

func createPlaceLimitOrderEvent(creator string, receiver string, token0 string, token1 string, tokenIn string, amountIn string, shares string, currentLimitOrderKey string, otherAttrs ...sdk.Attribute) sdk.Event {
	attrs := []sdk.Attribute{
		sdk.NewAttribute(sdk.AttributeKeyModule, "dex"),
		sdk.NewAttribute(sdk.AttributeKeyAction, PlaceLimitOrderEventKey),
		sdk.NewAttribute(PlaceLimitOrderEventCreator, creator),
		sdk.NewAttribute(PlaceLimitOrderEventReceiver, receiver),
		sdk.NewAttribute(PlaceLimitOrderEventToken0, token0),
		sdk.NewAttribute(PlaceLimitOrderEventToken1, token1),
		sdk.NewAttribute(PlaceLimitOrderEventTokenIn, tokenIn),
		sdk.NewAttribute(PlaceLimitOrderEventAmountIn, amountIn),
		sdk.NewAttribute(PlaceLimitOrderEventShares, shares),
		sdk.NewAttribute(PlaceLimitOrderEventCurrentKey, currentLimitOrderKey),
	}
	attrs = append(attrs, otherAttrs...)
	return sdk.NewEvent(sdk.EventTypeMessage, attrs...)
}

func CreatePlaceLimitOrderEvent(creator string, receiver string, token0 string, token1 string, tokenIn string, amountIn string, shares string, currentLimitOrderKey string, otherAttrs ...sdk.Attribute) sdk.Event {
	return createPlaceLimitOrderEvent(
		creator,
		receiver,
		token0,
		token1,
		tokenIn,
		amountIn,
		shares,
		currentLimitOrderKey,
	)
}

func withdrawFilledLimitOrderEvent(creator string, receiver string, token0 string, token1 string, tokenKey string, key string, amountOut string, otherAttrs ...sdk.Attribute) sdk.Event {
	attrs := []sdk.Attribute{
		sdk.NewAttribute(sdk.AttributeKeyModule, "dex"),
		sdk.NewAttribute(sdk.AttributeKeyAction, WithdrawFilledLimitOrderEventKey),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventCreator, creator),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventReceiver, receiver),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventToken0, token0),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventToken1, token1),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventTokenKey, tokenKey),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventLimitOrderKey, key),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventAmountOut, amountOut),
	}
	attrs = append(attrs, otherAttrs...)
	return sdk.NewEvent(sdk.EventTypeMessage, attrs...)
}

func WithdrawFilledLimitOrderEvent(creator string, receiver string, token0 string, token1 string, tokenKey string, key string, amountOut string, otherAttrs ...sdk.Attribute) sdk.Event {
	return withdrawFilledLimitOrderEvent(
		creator,
		receiver,
		token0,
		token1,
		tokenKey,
		key,
		amountOut,
	)
}

func cancelLimitOrderEvent(creator string, receiver string, token0 string, token1 string, tokenKey string, key string, amountOut string, otherAttrs ...sdk.Attribute) sdk.Event {
	attrs := []sdk.Attribute{
		sdk.NewAttribute(sdk.AttributeKeyModule, "dex"),
		sdk.NewAttribute(sdk.AttributeKeyAction, WithdrawFilledLimitOrderEventKey),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventCreator, creator),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventReceiver, receiver),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventToken0, token0),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventToken1, token1),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventTokenKey, tokenKey),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventLimitOrderKey, key),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventAmountOut, amountOut),
	}
	attrs = append(attrs, otherAttrs...)
	return sdk.NewEvent(sdk.EventTypeMessage, attrs...)
}

func CancelLimitOrderEvent(creator string, receiver string, token0 string, token1 string, tokenKey string, key string, amountOut string, otherAttrs ...sdk.Attribute) sdk.Event {
	return withdrawFilledLimitOrderEvent(
		creator,
		receiver,
		token0,
		token1,
		tokenKey,
		key,
		amountOut,
	)
}
