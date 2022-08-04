package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func CreateSwapEvent(creator string, tokenIn string, tokenOut string, amountIn string, amountOut string, otherAttrs ...sdk.Attribute) sdk.Event {
	attrs := append([]sdk.Attribute{
		sdk.NewAttribute(sdk.AttributeKeyModule, "duality"),
		sdk.NewAttribute(sdk.AttributeKeyAction, SwapEventKey),
		sdk.NewAttribute(SwapEventCreator, creator),
		sdk.NewAttribute(SwapEventTokenIn, tokenIn),
		sdk.NewAttribute(SwapEventTokenOut, tokenOut),
		sdk.NewAttribute(SwapEventAmountIn, amountIn),
		sdk.NewAttribute(SwapEventAmountOut, amountOut),
	}, otherAttrs...)
	return sdk.NewEvent(sdk.EventTypeMessage, attrs...)
}

func CreatePoolSwapEvent(tokenIn string, tokenOut string, amountSwappedIn string, amountSwappedOut string ,newReserve0 string, newReserve1 string, price string, fee string, otherAttrs ...sdk.Attribute) sdk.Event {
	attrs := append([]sdk.Attribute{
		sdk.NewAttribute(sdk.AttributeKeyModule, "duality"),
		sdk.NewAttribute(sdk.AttributeKeyAction, SwapEventKey),
		sdk.NewAttribute(SwapEventTokenIn, tokenIn),
		sdk.NewAttribute(SwapEventTokenOut, tokenOut),
		sdk.NewAttribute(SwapEventAmountIn, amountSwappedIn),
		sdk.NewAttribute(SwapEventAmountOut, amountSwappedOut),
		sdk.NewAttribute(SwapEventNewPoolReserve0, newReserve0),
		sdk.NewAttribute(SwapEventNewPoolReserve1, newReserve1),
		sdk.NewAttribute(SwapEventPriceSwap, price),
		sdk.NewAttribute(SwapEventFeeSwap, fee),
	}, otherAttrs... )
	return sdk.NewEvent(sdk.EventTypeMessage, attrs...)
}