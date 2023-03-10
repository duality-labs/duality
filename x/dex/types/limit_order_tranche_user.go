package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (l LimitOrderTrancheUser) IsEmpty() bool {
	sharesRemoved := l.SharesCancelled.Add(l.SharesWithdrawn)
	return sharesRemoved.Equal(l.SharesOwned) && l.ReservesFromSwap.IsZero()
}

func (l *LimitOrderTrancheUser) WithdrawSwapReserves() sdk.Int {
	amountOut := l.ReservesFromSwap
	l.ReservesFromSwap = sdk.ZeroInt()
	return amountOut

}
