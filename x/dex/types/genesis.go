package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		//TicksList:      []Ticks{},
		PairsList: []Pairs{},
		//IndexQueueList: []IndexQueue{},
		SharesList: []Shares{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {

	// // Check for duplicated index in ticks
	// ticksIndexMap := make(map[string]struct{})

	// for _, elem := range gs.TicksList {
	// 	index := string(TicksKey(elem.Price, elem.Fee, elem.Direction, elem.OrderType))
	// 	if _, ok := ticksIndexMap[index]; ok {
	// 		return fmt.Errorf("duplicated index for ticks")
	// 	}
	// 	ticksIndexMap[index] = struct{}{}
	// }

	// Check for duplicated index in pairs
	pairsIndexMap := make(map[string]struct{})

	for _, elem := range gs.PairsList {
		index := string(PairsKey(elem.Token0, elem.Token1))
		if _, ok := pairsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for pairs")
		}

		pairsIndexMap[index] = struct{}{}
	}

	// // Check for duplicated index in IndexQueue
	// IndexQueueIndexMap := make(map[string]struct{})

	// for _, elem := range gs.IndexQueueList {
	// 	index := string(IndexQueueKey(elem.Index))
	// 	if _, ok := IndexQueueIndexMap[index]; ok {
	// 		return fmt.Errorf("duplicated index for IndexQueue")
	// 	}
	// 	IndexQueueIndexMap[index] = struct{}{}
	//}

	// Check for duplicated index in shares
	sharesIndexMap := make(map[string]struct{})

	for _, elem := range gs.SharesList {
		index := string(SharesKey(elem.Address, elem.Price, elem.Fee, elem.OrderType))
		if _, ok := sharesIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for shares")
		}
		sharesIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
