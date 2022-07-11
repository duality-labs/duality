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

// Deposit Event Attributes
const (
	DepositEventKey          = "NewDeposit"
	DepositEventCreator      = "Creator"
	DepositEventToken0       = "Token0"
	DepositEventToken1       = "Token1"
	DepositEventPrice        = "Price"
	DepositEventFee          = "Fee"
	DepositEventReceiver     = "Receiver"
	DepositEventOldReserves0 = "OldReserves0"
	DepositEventOldReserves1 = "OldReserves1"
	DepositEventNewReserves0 = "NewReserves0"
	DepositEventNewReserves1 = "NewReserves1"
	DepositEventSharesMinted = "SharesMinted"
)

// Withdraw Event Attributes
const (
	WithdrawEventKey          = "NewWithdraw"
	WithdrawEventCreator      = "Creator"
	WithdrawEventToken0       = "Token0"
	WithdrawEventToken1       = "Token1"
	WithdrawEventPrice        = "Price"
	WithdrawEventFee          = "Fee"
	WithdrawEventOldReserves0 = "OldReserves0"
	WithdrawEventOldReserves1 = "OldReserves1"
	WithdrawEventNewReserves0 = "NewReserves0"
	WithdrawEventNewReserves1 = "NewReserves1"
	WithdrawEventReceiver     = "Receiver"
	WithdrawEventAmounts0      = "Shares0"
	WithdrawEventAmounts1      = "Shares1"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
