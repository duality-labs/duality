package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		LimitOrderTrancheUserList:     []LimitOrderTrancheUser{},
		TickLiquidityList:             []TickLiquidity{},
		InactiveLimitOrderTrancheList: []LimitOrderTranche{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in LimitOrderTrancheUser
	LimitOrderTrancheUserIndexMap := make(map[string]struct{})

	for _, elem := range gs.LimitOrderTrancheUserList {
		index := string(LimitOrderTrancheUserKey(elem.Address, elem.TrancheKey))
		if _, ok := LimitOrderTrancheUserIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for LimitOrderTrancheUser")
		}
		LimitOrderTrancheUserIndexMap[index] = struct{}{}
	}

	// Check for duplicated index in tickLiquidity
	tickLiquidityIndexMap := make(map[string]struct{})

	for _, elem := range gs.TickLiquidityList {
		var index string
		switch liquidity := elem.Liquidity.(type) {
		case *TickLiquidity_PoolReserves:
			index = string(TickLiquidityKey(
				liquidity.PoolReserves.PairId,
				liquidity.PoolReserves.TokenIn,
				liquidity.PoolReserves.TickIndex,
				LiquidityTypePoolReserves,
				liquidity.PoolReserves.Fee))
		case *TickLiquidity_LimitOrderTranche:
			index = string(TickLiquidityKey(
				liquidity.LimitOrderTranche.PairId,
				liquidity.LimitOrderTranche.TokenIn,
				liquidity.LimitOrderTranche.TickIndex,
				LiquidityTypeLimitOrder,
				liquidity.LimitOrderTranche.TrancheKey))
		}
		if _, ok := tickLiquidityIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for tickLiquidity")
		}
		tickLiquidityIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in inactiveLimitOrderTranche
	inactiveLimitOrderTrancheKeyMap := make(map[string]struct{})

	for _, elem := range gs.InactiveLimitOrderTrancheList {
		index := string(InactiveLimitOrderTrancheKey(elem.PairId, elem.TokenIn, elem.TickIndex, elem.TrancheKey))
		if _, ok := inactiveLimitOrderTrancheKeyMap[index]; ok {
			return fmt.Errorf("duplicated index for inactiveLimitOrderTranche")
		}
		inactiveLimitOrderTrancheKeyMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
