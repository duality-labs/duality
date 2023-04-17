package types

// Incentive module event types.
const (
	TypeEvtCreateGauge  = "create_gauge"
	TypeEvtAddToGauge   = "add_to_gauge"
	TypeEvtDistribution = "distribution"

	AttributeGaugeID     = "gauge_id"
	AttributeStakedDenom = "denom"
	AttributeReceiver    = "receiver"
	AttributeAmount      = "amount"

	TypeEvtStake            = "stake"
	TypeEvtAddTokensToStake = "add_tokens_to_stake"
	TypeEvtUnstake          = "unstake"

	AttributeStakeID          = "period_stake_id"
	AttributeStakeOwner       = "owner"
	AttributeStakeAmount      = "amount"
	AttributeStakeUnstakeTime = "stake_time"
	AttributeUnstakedCoins    = "unstaked_coins"
)
