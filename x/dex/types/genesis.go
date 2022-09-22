package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		TickMapList:   []TickMap{},
		PairMapList:   []PairMap{},
		TokensList:    []Tokens{},
		TokenMapList:  []TokenMap{},
		SharesList:    []Shares{},
		FeeListList:   []FeeList{},
		EdgeRowList:   []EdgeRow{},
		AdjMatrixList: []AdjMatrix{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in tickMap
	tickMapIndexMap := make(map[string]struct{})

	for _, elem := range gs.TickMapList {
		index := string(TickMapKey(elem.TickIndex))
		if _, ok := tickMapIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for tickMap")
		}
		tickMapIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in pairMap
	pairMapIndexMap := make(map[string]struct{})

	for _, elem := range gs.PairMapList {
		index := string(PairMapKey(elem.PairId))
		if _, ok := pairMapIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for pairMap")
		}
		pairMapIndexMap[index] = struct{}{}
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
	// Check for duplicated index in shares
	sharesIndexMap := make(map[string]struct{})

	for _, elem := range gs.SharesList {
		index := string(SharesKey(elem.Address, elem.PairId, elem.PriceIndex, elem.FeeIndex))
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
	// Check for duplicated ID in edgeRow
	edgeRowIdMap := make(map[uint64]bool)
	edgeRowCount := gs.GetEdgeRowCount()
	for _, elem := range gs.EdgeRowList {
		if _, ok := edgeRowIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for edgeRow")
		}
		if elem.Id >= edgeRowCount {
			return fmt.Errorf("edgeRow id should be lower or equal than the last id")
		}
		edgeRowIdMap[elem.Id] = true
	}
	// Check for duplicated ID in adjMatrix
	adjMatrixIdMap := make(map[uint64]bool)
	adjMatrixCount := gs.GetAdjMatrixCount()
	for _, elem := range gs.AdjMatrixList {
		if _, ok := adjMatrixIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for adjMatrix")
		}
		if elem.Id >= adjMatrixCount {
			return fmt.Errorf("adjMatrix id should be lower or equal than the last id")
		}
		adjMatrixIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
