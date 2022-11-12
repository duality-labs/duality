package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewACoin(amt sdk.Int) sdk.Coin {
	return sdk.NewCoin("TokenA", amt)
}

func NewBCoin(amt sdk.Int) sdk.Coin {
	return sdk.NewCoin("TokenB", amt)
}

func ConvInt(amt string) sdk.Int {
	IntAmt, err := sdk.NewIntFromString(amt)

	_ = err
	return IntAmt
}

func NewDec(a int) sdk.Dec {
	return sdk.NewDecFromInt(sdk.NewIntFromUint64(uint64(a)))
}
