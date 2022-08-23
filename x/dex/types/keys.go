package types

const (
	// ModuleName defines the module name
	ModuleName = "dex"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_dex"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// Deposit Event Attributes
const (
	DepositEventKey          = "NewDeposit"
	DepositEventCreator      = "Creator"
	DepositEventToken0       = "Token0"
	DepositEventToken1       = "Token1"
	DepositEventPrice        = "Price"
	DepositEventFee          = "Fee"
	DepositEventReceiver     = "Receiver"
	DepositTokenDirection    = "TokenDirection"
	DepositEventOldReserves  = "OldReserves"
	DepositEventNewReserves  = "NewReserves"
	DepositEventSharesMinted = "SharesMinted"
)

// Withdraw Event Attributes
const (
	WithdrawEventKey           = "NewWithdraw"
	WithdrawEventCreator       = "Creator"
	WithdrawEventToken0        = "Token0"
	WithdrawEventToken1        = "Token1"
	WithdrawEventPrice         = "Price"
	WithdrawEventFee           = "Fee"
	WithdrawEventReceiver      = "Receiver"
	WithdrawEventOldReserve0   = "OldReserve0"
	WithdrawEventOldReserve1   = "OldReserve0"
	WithdrawEventNewReserve0   = "NewReserve0"
	WithdrawEventNewReserve1   = "NewReserve1"
	WithdrawEventSharesRemoved = "SharesRemoved"
)
