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
		sdk.NewAttribute(WithdrawEventOldReserve0, oldReserve0),
		sdk.NewAttribute(WithdrawEventOldReserve1, oldReserve1),
		sdk.NewAttribute(WithdrawEventNewReserve0, newReserve0),
		sdk.NewAttribute(WithdrawEventNewReserve1, newReserve1),
	)
}

func createSwapEvent(creator string, receiver string, token0 string, token1 string, tokenIn string, amountIn string, amountOut string, minOut string, otherAttrs ...sdk.Attribute) sdk.Event {
	attrs := []sdk.Attribute{
		sdk.NewAttribute(sdk.AttributeKeyModule, "dex"),
		sdk.NewAttribute(sdk.AttributeKeyAction, SwapEventKey),
		sdk.NewAttribute(WithdrawEventCreator, creator),
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
