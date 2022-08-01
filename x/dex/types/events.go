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

func CreateDepositEvent(creator string, token0 string, token1 string,  price string, fee string, oldReserves0 string, oldReserves1 string, newReserves0 string, newReserves1 string, otherAttrs ...sdk.Attribute) sdk.Event {
	return createDepositEvent(
		creator,
		token0,
		token1,
		price,
		fee,
		sdk.NewAttribute(DepositEventOldReserves0, oldReserves0),
		sdk.NewAttribute(DepositEventOldReserves1, oldReserves1),
		sdk.NewAttribute(DepositEventNewReserves0, newReserves0),
		sdk.NewAttribute(DepositEventNewReserves1, newReserves1),
	)
}

func createWithdrawEvent(creator string, token0 string, token1 string, price string, fee string, otherAttrs ...sdk.Attribute) sdk.Event {
	attrs := []sdk.Attribute{
		sdk.NewAttribute(sdk.AttributeKeyModule, "dex"),
		sdk.NewAttribute(sdk.AttributeKeyAction, WithdrawEventKey),
		sdk.NewAttribute(WithdrawEventCreator, creator),
		sdk.NewAttribute(WithdrawEventToken0, token0),
		sdk.NewAttribute(WithdrawEventToken1, token1),
		sdk.NewAttribute(WithdrawEventPrice, price),
		sdk.NewAttribute(WithdrawEventFee, fee),
	}
	attrs = append(attrs, otherAttrs...)
	return sdk.NewEvent(sdk.EventTypeMessage, attrs...)
}

func CreateWithdrawEvent(creator string, token0 string, token1 string,  price string, fee string, oldReserves0 string, oldReserves1 string, newReserves0 string, newReserves1 string, otherAttrs ...sdk.Attribute) sdk.Event {
	return createWithdrawEvent(
		creator,
		token0,
		token1,
		price,
		fee,
		sdk.NewAttribute(WithdrawEventOldReserves0, oldReserves0),
		sdk.NewAttribute(WithdrawEventOldReserves1, oldReserves1),
		sdk.NewAttribute(WithdrawEventNewReserves0, newReserves0),
		sdk.NewAttribute(WithdrawEventNewReserves1, newReserves1),
	)
}