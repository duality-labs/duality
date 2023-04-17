package types

// query endpoints supported by the stakeup QueryServer.
const (
	QueryModuleBalance                    = "module_balance"
	QueryModuleStakedAmount               = "module_staked_amount"
	QueryAccountUnstakeableCoins          = "account_unstakeable_coins"
	QueryAccountStakedCoins               = "account_staked_coins"
	QueryAccountStakedPastTime            = "account_staked_pastime"
	QueryAccountUnstakedBeforeTime        = "account_unstaked_beforetime"
	QueryAccountStakedPastTimeDenom       = "account_staked_denom_pastime"
	QueryStakedByID                       = "staked_by_id"
	QueryAccountStakedLongerDuration      = "account_staked_longer_than_duration"
	QueryAccountStakedLongerDurationDenom = "account_staked_longer_than_duration_denom"
	QueryAccountStakedDuration            = "account_staked_duration"
)
