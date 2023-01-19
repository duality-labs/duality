package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		TradingPairList:             []TradingPair{},
		TokensList:                  []Tokens{},
		TokenMapList:                []TokenMap{},
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
	// Check for duplicated index in TradingPair
	TradingPairIndexMap := make(map[string]struct{})

	for _, elem := range gs.TradingPairList {
		index := string(TradingPairKey(elem.PairId))
		if _, ok := TradingPairIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for TradingPair")
		}
		TradingPairIndexMap[index] = struct{}{}
	}
	// Check for duplicated ID in tokens
	tokensIdMap := make(map[uint64]bool)
	tokensCount := gs.GetTokensCount()
	for _, elem := range gs.TokensList {
		if _, ok := tokensIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for tokens")
		}
		if elem.Id >= tokensCount {
			return fmt.Errorf("tokens id should be lower or equal than the last id")
		}
		tokensIdMap[elem.Id] = true
	}
	// Check for duplicated index in tokenMap
	tokenMapIndexMap := make(map[string]struct{})

	for _, elem := range gs.TokenMapList {
		index := string(TokenMapKey(elem.Address))
		if _, ok := tokenMapIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for tokenMap")
		}
		tokenMapIndexMap[index] = struct{}{}
	}
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
		index := string(LimitOrderTrancheUserKey(elem.PairId, elem.TickIndex, elem.Token, elem.Count, elem.Address))
		if _, ok := LimitOrderTrancheUserIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for LimitOrderTrancheUser")
		}
		LimitOrderTrancheUserIndexMap[index] = struct{}{}
	}

	// Check for duplicated index in tickLiquidity
	tickLiquidityIndexMap := make(map[string]struct{})

	for _, elem := range gs.TickLiquidityList {
		index := string(TickLiquidityKey(elem.PairId, elem.TokenIn, elem.TickIndex, elem.LiquidityType, elem.LiquidityIndex))
		if _, ok := tickLiquidityIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for tickLiquidity")
		}
		tickLiquidityIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in filledLimitOrderTranche
	filledLimitOrderTrancheIndexMap := make(map[string]struct{})

	for _, elem := range gs.FilledLimitOrderTrancheList {
		index := string(FilledLimitOrderTrancheKey(elem.PairId, elem.TokenIn, elem.TickIndex, elem.TrancheIndex))
		if _, ok := filledLimitOrderTrancheIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for filledLimitOrderTranche")
		}
		filledLimitOrderTrancheIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
