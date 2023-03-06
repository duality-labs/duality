package types

func (l LimitOrderType) IsGTC() bool {
	return l == LimitOrderType_GOOD_TIL_CANCELLED
}

func (l LimitOrderType) IsFoK() bool {
	return l == LimitOrderType_FILL_OR_KILL
}

func (l LimitOrderType) IsIoC() bool {
	return l == LimitOrderType_IMMEDIATE_OR_CANCEL
}

func (l LimitOrderType) IsJIT() bool {
	return l == LimitOrderType_JUST_IN_TIME
}
