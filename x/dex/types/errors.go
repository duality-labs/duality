package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/dex module sentinel errors
var (
	ErrInvalidTradingPair                 = sdkerrors.Register(ModuleName, 1102, "Invalid token pair:")   // "%s<>%s", tokenA, tokenB
	ErrInsufficientShares                 = sdkerrors.Register(ModuleName, 1104, "Insufficient shares:")  // "%s does not have %s shares of type %s", address, shares, sharesId
	ErrValidTickNotFound                  = sdkerrors.Register(ModuleName, 1106, "Valid tick not found:") // "%d", tickIndex
	ErrValidPairNotFound                  = sdkerrors.Register(ModuleName, 1107, "Valid pair not found.")
	ErrValidFeeIndexNotFound              = sdkerrors.Register(ModuleName, 1108, "Valid FeeIndex not found.") // "%d", feeIndex
	ErrUnbalancedTxArray                  = sdkerrors.Register(ModuleName, 1110, "Transaction input arrays are not of the same length.")
	ErrValidLimitOrderTrancheNotFound     = sdkerrors.Register(ModuleName, 1111, "Limit order trache not found:")                                // "%d", trancheKey
	ErrCancelEmptyLimitOrder              = sdkerrors.Register(ModuleName, 1112, "Cannot cancel additional liquidity from limit order tranche:") // "%d", tranche.TrancheKey
	ErrSlippageLimitReached               = sdkerrors.Register(ModuleName, 1114, "Slippage limit reached, minOut could not be satisfied.")
	ErrDepositBehindPairLiquidity         = sdkerrors.Register(ModuleName, 1115, "Depositing behind the opposite token pair's liquidity is currently not allowed")
	ErrPlaceLimitOrderBehindPairLiquidity = sdkerrors.Register(ModuleName, 1116, "Placing a limit order behind the opposite token pair's liquidity is currently not allowed")
	ErrTickOutsideRange                   = sdkerrors.Register(ModuleName, 1117, "Supplying a tick outside the range of [-1048575, 1048575] is not allowed")
	ErrInvalidDepositShares               = sdkerrors.Register(ModuleName, 1118, "Denom is not an instance of Duality Pool Shares.")
	ErrInvalidPairIdStr                   = sdkerrors.Register(ModuleName, 1119, "PairId does not conform to pattern TokenA<>TokenB.")
	ErrZeroDeposit                        = sdkerrors.Register(ModuleName, 1120, "At least one deposit amount must be > 0.")
	ErrZeroTrueDeposit                    = sdkerrors.Register(ModuleName, 1121, "Cannot deposit double-sided liquidity in tick with prexisting single-sided liquidity.")
	ErrNotEnoughLimitOrderShares          = sdkerrors.Register(ModuleName, 1122, "Not enough limit order shares.")
	ErrValidLimitOrderTrancheUserNotFound = sdkerrors.Register(ModuleName, 1123, "Limit order trache user not found:") // "tranche %d, user %s", trancheKey, address
	ErrWithdrawEmptyLimitOrder            = sdkerrors.Register(ModuleName, 1124, "Cannot withdraw additional liqudity from this limit order at this time.")
	ErrZeroSwap                           = sdkerrors.Register(ModuleName, 1125, "Amount in must be > 0 for swap.")
	ErrInvalidKeyToken                    = sdkerrors.Register(ModuleName, 1126, "KeyToken not in specified pair.")
	ErrInvalidTokenIn                     = sdkerrors.Register(ModuleName, 1127, "TokenIn not in specified pair.")
	ErrActiveLimitOrderNotFound           = sdkerrors.Register(ModuleName, 1128, "No active limit found. It does not exist or has already been filled")
	ErrZeroWithdraw                       = sdkerrors.Register(ModuleName, 1129, "Withdraw amount must be > 0.")
	ErrZeroLimitOrder                     = sdkerrors.Register(ModuleName, 1130, "Limit order amount must be > 0.")
	ErrNegativeMinOut                     = sdkerrors.Register(ModuleName, 1131, "MinOut must be >= 0.")
	ErrNegativeLimitPrice                 = sdkerrors.Register(ModuleName, 1132, "LimitPrice must be > 0.")
)
