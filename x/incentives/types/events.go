package types

// Incentive module event types.
const (
	TypeEvtCreateGauge  = "create_gauge"
	TypeEvtAddToGauge   = "add_to_gauge"
	TypeEvtDistribution = "distribution"

	AttributeGaugeID     = "gauge_id"
	AttributeLockedDenom = "denom"
	AttributeReceiver    = "receiver"
	AttributeAmount      = "amount"

	TypeEvtLockTokens      = "lock_tokens"
	TypeEvtAddTokensToLock = "add_tokens_to_lock"
	TypeEvtBeginUnlockAll  = "begin_unlock_all"
	TypeEvtBeginUnlock     = "begin_unlock"

	AttributeLockID         = "period_lock_id"
	AttributeLockOwner      = "owner"
	AttributeLockAmount     = "amount"
	AttributeLockDuration   = "duration"
	AttributeLockUnlockTime = "unlock_time"
	AttributeUnlockedCoins  = "unlocked_coins"
)
