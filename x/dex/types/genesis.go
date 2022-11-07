package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		TickObjectList:                        []TickObject{},
		PairObjectList:                        []PairObject{},
		TokensList:                            []Tokens{},
		TokenObjectList:                       []TokenObject{},
		SharesList:                            []Shares{},
		FeeListList:                           []FeeList{},
		LimitOrderPoolUserShareObjectList:     []LimitOrderPoolUserShareObject{},
		LimitOrderPoolUserSharesWithdrawnList: []LimitOrderPoolUserSharesWithdrawn{},
		LimitOrderPoolTotalSharesObjectList:   []LimitOrderPoolTotalSharesObject{},
		LimitOrderPoolReserveObjectList:       []LimitOrderPoolReserveObject{},
		LimitOrderPoolFillObjectList:          []LimitOrderPoolFillObject{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in tickObject
	tickObjectIndexMap := make(map[string]struct{})

	for _, elem := range gs.TickObjectList {
		index := string(TickObjectKey(elem.PairId, elem.TickIndex))
		if _, ok := tickObjectIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for tickObject")
		}
		tickObjectIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in pairObject
	pairObjectIndexMap := make(map[string]struct{})

	for _, elem := range gs.PairObjectList {
		index := string(PairObjectKey(elem.PairId))
		if _, ok := pairObjectIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for pairObject")
		}
		pairObjectIndexMap[index] = struct{}{}
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
	// Check for duplicated index in tokenObject
	tokenObjectIndexMap := make(map[string]struct{})

	for _, elem := range gs.TokenObjectList {
		index := string(TokenObjectKey(elem.Address))
		if _, ok := tokenObjectIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for tokenObject")
		}
		tokenObjectIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in shares
	sharesIndexMap := make(map[string]struct{})

	for _, elem := range gs.SharesList {
		index := string(SharesKey(elem.Address, elem.PairId, elem.TickIndex, elem.FeeIndex))
		if _, ok := sharesIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for shares")
		}
		sharesIndexMap[index] = struct{}{}
	}
	// Check for duplicated ID in feeList
	feeListIdMap := make(map[uint64]bool)
	feeListCount := gs.GetFeeListCount()
	for _, elem := range gs.FeeListList {
		if _, ok := feeListIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for feeList")
		}
		if elem.Id >= feeListCount {
			return fmt.Errorf("feeList id should be lower or equal than the last id")
		}
		feeListIdMap[elem.Id] = true
	}
	// Check for duplicated index in limitOrderPoolUserShareObject
	limitOrderPoolUserShareObjectIndexMap := make(map[string]struct{})

	for _, elem := range gs.LimitOrderPoolUserShareObjectList {
		index := string(LimitOrderPoolUserShareObjectKey(elem.PairId, elem.TickIndex, elem.Token, elem.Count, elem.Address))
		if _, ok := limitOrderPoolUserShareObjectIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for limitOrderPoolUserShareObject")
		}
		limitOrderPoolUserShareObjectIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in limitOrderPoolUserSharesWithdrawn
	limitOrderPoolUserSharesWithdrawnIndexMap := make(map[string]struct{})

	for _, elem := range gs.LimitOrderPoolUserSharesWithdrawnList {
		index := string(LimitOrderPoolUserSharesWithdrawnKey(elem.PairId, elem.TickIndex, elem.Token, elem.Count, elem.Address))
		if _, ok := limitOrderPoolUserSharesWithdrawnIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for limitOrderPoolUserSharesWithdrawn")
		}
		limitOrderPoolUserSharesWithdrawnIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in limitOrderPoolTotalSharesObject
	limitOrderPoolTotalSharesObjectIndexMap := make(map[string]struct{})

	for _, elem := range gs.LimitOrderPoolTotalSharesObjectList {
		index := string(LimitOrderPoolTotalSharesObjectKey(elem.PairId, elem.TickIndex, elem.Token, elem.Count))
		if _, ok := limitOrderPoolTotalSharesObjectIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for limitOrderPoolTotalSharesObject")
		}
		limitOrderPoolTotalSharesObjectIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in limitOrderPoolReserveObject
	limitOrderPoolReserveObjectIndexMap := make(map[string]struct{})

	for _, elem := range gs.LimitOrderPoolReserveObjectList {
		index := string(LimitOrderPoolReserveObjectKey(elem.PairId, elem.TickIndex, elem.Token, elem.Count))
		if _, ok := limitOrderPoolReserveObjectIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for limitOrderPoolReserveObject")
		}
		limitOrderPoolReserveObjectIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in limitOrderPoolFillObject
	limitOrderPoolFillObjectIndexMap := make(map[string]struct{})

	for _, elem := range gs.LimitOrderPoolFillObjectList {
		index := string(LimitOrderPoolFillObjectKey(elem.PairId, elem.TickIndex, elem.Token, elem.Count))
		if _, ok := limitOrderPoolFillObjectIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for limitOrderPoolFillObject")
		}
		limitOrderPoolFillObjectIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
