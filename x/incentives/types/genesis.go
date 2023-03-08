package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		IncentivePlanList: []IncentivePlan{},
		UserStakeList:     []UserStake{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in incentivePlan
	incentivePlanIndexMap := make(map[string]struct{})

	for _, elem := range gs.IncentivePlanList {
		index := string(IncentivePlanKey(elem.Index))
		if _, ok := incentivePlanIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for incentivePlan")
		}
		incentivePlanIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in userStake
	userStakeIndexMap := make(map[string]struct{})

	for _, elem := range gs.UserStakeList {
		index := string(UserStakeKey(elem.Index))
		if _, ok := userStakeIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for userStake")
		}
		userStakeIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
