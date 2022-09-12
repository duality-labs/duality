package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func createDepositEvent(creator string, token0 string, token1 string, price string, fee string, otherAttrs ...sdk.Attribute) sdk.Event {
	attrs := []sdk.Attribute{
		sdk.NewAttribute(sdk.AttributeKeyModule, "dex"),
		sdk.NewAttribute(sdk.AttributeKeyAction, DepositEventKey),
		sdk.NewAttribute(DepositEventCreator, creator),
		sdk.NewAttribute(DepositEventToken0, token0),
		sdk.NewAttribute(DepositEventToken1, token1),
		sdk.NewAttribute(DepositEventPrice, price),
		sdk.NewAttribute(DepositEventFee, fee),
	}
	attrs = append(attrs, otherAttrs...)
	return sdk.NewEvent(sdk.EventTypeMessage, attrs...)
}

func CreateDepositEvent(creator string, token0 string, token1 string, price string, fee string, tokenDirection string, oldReserve0 string, oldReserve1 string, newReserve0 string, newReserve1 string, otherAttrs ...sdk.Attribute) sdk.Event {
	return createDepositEvent(
		creator,
		token0,
		token1,
		price,
		fee,

		sdk.NewAttribute(DepositEventOldReserves0, oldReserve0),
		sdk.NewAttribute(DepositEventNewReserves0, newReserve0),
		sdk.NewAttribute(DepositEventOldReserves1, oldReserve1),
		sdk.NewAttribute(DepositEventNewReserves1, newReserve1),
	)
}
