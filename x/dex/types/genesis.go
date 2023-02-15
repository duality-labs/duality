package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		FeeTierList:                 []FeeTier{},
		LimitOrderTrancheUserList:   []LimitOrderTrancheUser{},
		TickLiquidityList:           []TickLiquidity{},
		FilledLimitOrderTrancheList: []FilledLimitOrderTranche{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in FeeTier
	FeeTierIdMap := make(map[uint64]bool)
	FeeTierCount := gs.GetFeeTierCount()
	for _, elem := range gs.FeeTierList {
		if _, ok := FeeTierIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for FeeTier")
		}
		if elem.Id >= FeeTierCount {
			return fmt.Errorf("FeeTier id should be lower or equal than the last id")
		}
		FeeTierIdMap[elem.Id] = true
	}
	// Check for duplicated index in LimitOrderTrancheUser
	LimitOrderTrancheUserIndexMap := make(map[string]struct{})

	for _, elem := range gs.LimitOrderTrancheUserList {
		index := string(LimitOrderTrancheUserKey(elem.PairId, elem.TickIndex, elem.Token, elem.TrancheKey, elem.Address))
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
	// Check for duplicated index in filledLimitOrderTranche
	filledLimitOrderTrancheKeyMap := make(map[string]struct{})

	for _, elem := range gs.FilledLimitOrderTrancheList {
		index := string(FilledLimitOrderTrancheKey(elem.PairId, elem.TokenIn, elem.TickIndex, elem.TrancheKey))
		if _, ok := filledLimitOrderTrancheKeyMap[index]; ok {
			return fmt.Errorf("duplicated index for filledLimitOrderTranche")
		}
		filledLimitOrderTrancheKeyMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
