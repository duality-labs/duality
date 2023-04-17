package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type StakeI interface {
	GetOwner() string
	Amount() sdk.Coins
}
