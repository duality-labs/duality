package types

func NewPoolMetadata(
	id uint64,
	pairId *PairID,
	normalizedCenterTickIndex int64,
	fee uint64,
) *PoolMetadata {
	return &PoolMetadata{
		Id:                        id,
		PairId:                    pairId,
		NormalizedCenterTickIndex: normalizedCenterTickIndex,
		Fee:                       fee,
	}
}

func (p *PoolMetadata) Denom() string {
	if p.Id == 0 {
		panic("Cannot call denom on PoolMetadata where poolID is not set")
	}
	return NewDepositDenom(p.Id)
}
