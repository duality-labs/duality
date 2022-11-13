package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Checks if a tick has reserves0 at any fee tier
func (m *TickDataType) HasToken0() bool {
	for _, s := range m.Reserve0AndShares {
		if s.Reserve0.GT(sdk.ZeroDec()) {
			return true
		}
	}
	return false
}

// Checks if a tick has reserve1 at any fee tier
func (m *TickDataType) HasToken1() bool {
	for _, s := range m.Reserve1 {
		if s.GT(sdk.ZeroDec()) {
			return true
		}
	}
	return false
}

// Checks if a tick has shares at any fee tier
func (m *TickDataType) HasShares() bool {
	for _, s := range m.Reserve0AndShares {
		if s.TotalShares.GT(sdk.ZeroDec()) {
			return true
		}
	}
	return false
}
