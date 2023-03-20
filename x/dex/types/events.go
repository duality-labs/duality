package types

import (
	"strconv"

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

func CreateDepositEvent(creator string, receiver string, token0 string, token1 string, tickIndex string, feeIndex string, oldReserve0 string, oldReserve1 string, newReserve0 string, newReserve1 string, sharesMinted string, otherAttrs ...sdk.Attribute) sdk.Event {
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
		sdk.NewAttribute(DepositEventSharesMinted, sharesMinted),
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

func CreateWithdrawEvent(creator string, receiver string, token0 string, token1 string, tickIndex string, feeIndex string, oldReserve0 string, oldReserve1 string, newReserve0 string, newReserve1 string, sharesRemoved string, otherAttrs ...sdk.Attribute) sdk.Event {
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
		sdk.NewAttribute(WithdrawEventSharesRemoved, sharesRemoved),
	)
}

func createSwapEvent(creator string, receiver string, tokenIn string, tokenOut string, amountIn string, amountOut string, otherAttrs ...sdk.Attribute) sdk.Event {
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
	attrs = append(attrs, otherAttrs...)
	return sdk.NewEvent(sdk.EventTypeMessage, attrs...)
}

func CreateSwapEvent(creator string, receiver string, tokenIn string, tokenOut string, amountIn string, amountOut string, otherAttrs ...sdk.Attribute) sdk.Event {
	return createSwapEvent(
		creator,
		receiver,
		tokenIn,
		tokenOut,
		amountIn,
		amountOut,
	)
}

func createPlaceLimitOrderEvent(creator string, receiver string, tokenIn string, tokenOut string, amountIn string, shares string, trancheKey string, otherAttrs ...sdk.Attribute) sdk.Event {
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
	attrs = append(attrs, otherAttrs...)
	return sdk.NewEvent(sdk.EventTypeMessage, attrs...)
}

func CreatePlaceLimitOrderEvent(creator string, receiver string, tokenIn string, tokenOut string, amountIn string, shares string, currentLimitOrderKey string, otherAttrs ...sdk.Attribute) sdk.Event {
	return createPlaceLimitOrderEvent(
		creator,
		receiver,
		tokenIn,
		tokenOut,
		amountIn,
		shares,
		currentLimitOrderKey,
	)
}

func withdrawFilledLimitOrderEvent(creator string, token0 string, token1 string, tokenKey string, key string, amountOut string, otherAttrs ...sdk.Attribute) sdk.Event {
	attrs := []sdk.Attribute{
		sdk.NewAttribute(sdk.AttributeKeyModule, "dex"),
		sdk.NewAttribute(sdk.AttributeKeyAction, WithdrawFilledLimitOrderEventKey),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventCreator, creator),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventToken0, token0),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventToken1, token1),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventTokenKey, tokenKey),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventLimitOrderKey, key),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventAmountOut, amountOut),
	}
	attrs = append(attrs, otherAttrs...)
	return sdk.NewEvent(sdk.EventTypeMessage, attrs...)
}

func WithdrawFilledLimitOrderEvent(creator string, token0 string, token1 string, tokenKey string, key string, amountOut string, otherAttrs ...sdk.Attribute) sdk.Event {
	return withdrawFilledLimitOrderEvent(
		creator,
		token0,
		token1,
		tokenKey,
		key,
		amountOut,
	)
}

func GoodTilPurgeHitLimitEvent(gas sdk.Gas, otherAttrs ...sdk.Attribute) sdk.Event {
	attrs := []sdk.Attribute{
		sdk.NewAttribute(sdk.AttributeKeyModule, "dex"),
		sdk.NewAttribute(sdk.AttributeKeyAction, GoodTilPurgeHitGasLimitEventKey),
		sdk.NewAttribute(GoodTilPurgeHitGasLimitEventGas, strconv.FormatUint(gas, 10)),
	}
	attrs = append(attrs, otherAttrs...)
	return sdk.NewEvent(sdk.EventTypeMessage, attrs...)

}

func cancelLimitOrderEvent(creator string, token0 string, token1 string, tokenKey string, key string, amountOut string, otherAttrs ...sdk.Attribute) sdk.Event {
	attrs := []sdk.Attribute{
		sdk.NewAttribute(sdk.AttributeKeyModule, "dex"),
		sdk.NewAttribute(sdk.AttributeKeyAction, WithdrawFilledLimitOrderEventKey),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventCreator, creator),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventToken0, token0),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventToken1, token1),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventTokenKey, tokenKey),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventLimitOrderKey, key),
		sdk.NewAttribute(WithdrawFilledLimitOrderEventAmountOut, amountOut),
	}
	attrs = append(attrs, otherAttrs...)
	return sdk.NewEvent(sdk.EventTypeMessage, attrs...)
}

func CancelLimitOrderEvent(creator string, token0 string, token1 string, tokenKey string, key string, amountOut string, otherAttrs ...sdk.Attribute) sdk.Event {
	return withdrawFilledLimitOrderEvent(
		creator,
		token0,
		token1,
		tokenKey,
		key,
		amountOut,
	)
}
