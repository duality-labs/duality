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