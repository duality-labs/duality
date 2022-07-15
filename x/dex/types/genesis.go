package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		TicksList: []Ticks{},
		ShareList: []Share{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in ticks
	ticksIndexMap := make(map[string]struct{})

	for _, elem := range gs.TicksList {
		index := string(TicksKey(elem.Token0, elem.Token1))
		if _, ok := ticksIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for ticks")
		}
		ticksIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in share
	shareIndexMap := make(map[string]struct{})

	for _, elem := range gs.ShareList {
		index := string(ShareKey(elem.Owner, elem.Token0, elem.Token1, elem.Price, elem.Fee))
		if _, ok := shareIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for share")
		}
		shareIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
