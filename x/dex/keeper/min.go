package keeper

func (k Keeper) min(a, b uint) uint {
	if a < b {
		return a
	}
	return b
}
