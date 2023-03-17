package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (l LimitOrderTrancheUser) IsEmpty() bool {
	sharesRemoved := l.SharesCancelled.Add(l.SharesWithdrawn)
	return sharesRemoved.Equal(l.SharesOwned) && l.TakerReserves.IsZero()
}

func (l *LimitOrderTrancheUser) WithdrawTakerReserves() sdk.Int {
	amountOut := l.TakerReserves
	l.TakerReserves = sdk.ZeroInt()
	return amountOut

}
