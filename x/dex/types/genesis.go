package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ShareList: []Share{},
		TickList:  []Tick{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in share
	shareIndexMap := make(map[string]struct{})

	for _, elem := range gs.ShareList {
		index := string(ShareKey(elem.Owner, elem.Token0, elem.Token1, elem.Price, elem.Fee))
		if _, ok := shareIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for share")
		}
		shareIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in tick
	tickIndexMap := make(map[string]struct{})

	for _, elem := range gs.TickList {
		index := string(TickKey(elem.Token0, elem.Token1, elem.Price, elem.Fee))
		if _, ok := tickIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for tick")
		}
		tickIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
