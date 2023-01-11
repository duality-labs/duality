package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/dex module sentinel errors
var (
	ErrNoSpendableCoins                   = sdkerrors.Register(ModuleName, 1100, "No Spendable Coins found: %s")
	ErrNotEnoughCoins                     = sdkerrors.Register(ModuleName, 1101, "Not enough Spendable Coins found: %s")
	ErrInvalidTradingPair                 = sdkerrors.Register(ModuleName, 1102, "Invalid Token Pair: (%s, %s)")
	ErrInvalidTokenListSize               = sdkerrors.Register(ModuleName, 1103, "Invalid Array: Array Tokens0 size does not equal Array Tokens1")
	ErrNotEnoughShares                    = sdkerrors.Register(ModuleName, 1104, "Not enough shares are owned by:  %s")
	ErrNotEnoughLimitOrderShares          = sdkerrors.Register(ModuleName, 1104, "Not enough limit order shares are owned by:  %s")
	ErrValidShareNotFound                 = sdkerrors.Register(ModuleName, 1105, "Valid share not found")
	ErrValidTickNotFound                  = sdkerrors.Register(ModuleName, 1106, "Valid tick not found")
	ErrValidPairNotFound                  = sdkerrors.Register(ModuleName, 1107, "Valid pair not found")
	ErrValidFeeIndexNotFound              = sdkerrors.Register(ModuleName, 1108, "Valid FeeIndex not found: %s ")
	ErrUnbalancedTxArray                  = sdkerrors.Register(ModuleName, 1110, "Transaction Input Arrays are not of the same length")
	ErrValidLimitOrderTrancheNotFound     = sdkerrors.Register(ModuleName, 1111, "Limit Order Trache not found")
	ErrValidLimitOrderTrancheUserNotFound = sdkerrors.Register(ModuleName, 1111, "Limit Order Trache User not found")
	ErrCancelEmptyLimitOrder              = sdkerrors.Register(ModuleName, 1112, "Cannot cancel additional liqudity from this limit order at this time")
	ErrAllDepositsFailed                  = sdkerrors.Register(ModuleName, 1113, "MsgDeposit Failed: All inputted deposits fail to complete")
	ErrSlippageLimitReached               = sdkerrors.Register(ModuleName, 1114, "Slippage limit reached.")
	ErrDepositBehindPairLiquidity         = sdkerrors.Register(ModuleName, 1115, "Depositing behind the opposite token pair's liquidity is currently not allowed")
	ErrPlaceLimitOrderBehindPairLiquidity = sdkerrors.Register(ModuleName, 1116, "Placing a limit order behind the opposite token pair's liquidity is currently not allowed")
	ErrTickOutsideRange                   = sdkerrors.Register(ModuleName, 1117, "Supplying a tick outside the range of [-1048575, 1048575] is not allowed")
	ErrInvalidDepositShares               = sdkerrors.Register(ModuleName, 1118, "Denom is not an instance of Duality Pool Shares")
	ErrInvalidPairIdStr                   = sdkerrors.Register(ModuleName, 1119, "PairId does not conform to pattern [TokenA]<>[TokenB]: %s")
	ErrZeroDeposit                        = sdkerrors.Register(ModuleName, 1120, "Deposit amount cannot be 0, 0")
	ErrZeroTrueDeposit                    = sdkerrors.Register(ModuleName, 1120, "Cannot deposit double-sided liquidity in tick with one-sided liquidity")
)
