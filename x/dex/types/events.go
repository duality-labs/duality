package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func createDepositEvent(creator string, token0 string, token1 string, priceIndex string, feeIndex string, otherAttrs ...sdk.Attribute) sdk.Event {
	attrs := []sdk.Attribute{
		sdk.NewAttribute(sdk.AttributeKeyModule, "dex"),
		sdk.NewAttribute(sdk.AttributeKeyAction, DepositEventKey),
		sdk.NewAttribute(DepositEventCreator, creator),
		sdk.NewAttribute(DepositEventToken0, token0),
		sdk.NewAttribute(DepositEventToken1, token1),
		sdk.NewAttribute(DepositEventPrice, priceIndex),
		sdk.NewAttribute(DepositEventFeeIndex, feeIndex),
	}
	attrs = append(attrs, otherAttrs...)
	return sdk.NewEvent(sdk.EventTypeMessage, attrs...)
}

func CreateDepositEvent(creator string, token0 string, token1 string, priceIndex string, feeIndex string, tokenDirection string, oldReserve0 string, oldReserve1 string, newReserve0 string, newReserve1 string, otherAttrs ...sdk.Attribute) sdk.Event {
	return createDepositEvent(
		creator,
		token0,
		token1,
		priceIndex,
		feeIndex,

		sdk.NewAttribute(DepositEventOldReserves0, oldReserve0),
		sdk.NewAttribute(DepositEventNewReserves0, newReserve0),
		sdk.NewAttribute(DepositEventOldReserves1, oldReserve1),
		sdk.NewAttribute(DepositEventNewReserves1, newReserve1),
	)
}

func createWithdrawEvent(creator string, token0 string, token1 string, priceIndex string, feeIndex string, otherAttrs ...sdk.Attribute) sdk.Event {
	attrs := []sdk.Attribute{
		sdk.NewAttribute(sdk.AttributeKeyModule, "dex"),
		sdk.NewAttribute(sdk.AttributeKeyAction, WithdrawEventKey),
		sdk.NewAttribute(WithdrawEventCreator, creator),
		sdk.NewAttribute(WithdrawEventToken0, token0),
		sdk.NewAttribute(WithdrawEventToken1, token1),
		sdk.NewAttribute(WithdrawEventPrice, priceIndex),
		sdk.NewAttribute(WithdrawEventFee, feeIndex),
	}
	attrs = append(attrs, otherAttrs...)
	return sdk.NewEvent(sdk.EventTypeMessage, attrs...)
}

func CreateWithdrawEvent(creator string, token0 string, token1 string, priceIndex string, feeIndex string, oldReserve0 string, oldReserve1 string, newReserve0 string, newReserve1 string, otherAttrs ...sdk.Attribute) sdk.Event {
	return createWithdrawEvent(
		creator,
		token0,
		token1,
		priceIndex,
		feeIndex,
		sdk.NewAttribute(WithdrawEventOldReserve0, oldReserve0),
		sdk.NewAttribute(WithdrawEventOldReserve1, oldReserve1),
		sdk.NewAttribute(WithdrawEventNewReserve0, newReserve0),
		sdk.NewAttribute(WithdrawEventNewReserve1, newReserve1),
	)
}
