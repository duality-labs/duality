package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	//"fmt"
)

/////////////////////////////////////////////////////////
// 1 to 0 Priority Queue Functions
/////////////////////////////////////////////////////////

func Len1to0(pools []*types.Pool) int32 { return int32(len(pools)) }

func Swap1to0(pools []*types.Pool, i, j int32) {
	pools[i], pools[j] = pools[j], pools[i]
	pools[i].Index = i
	pools[j].Index = j
}

func Less1to0(pools []*types.Pool, i, j int32) bool {

	return  (pools[i].Price.Mul(pools[i].Fee)).Quo(sdk.NewDecWithPrec(10000, 18)).LT( (pools[j].Price.Mul(pools[j].Fee)).Quo(sdk.NewDecWithPrec(10000, 18)))  

}

func Push1to0(pools *([]*types.Pool), newPool types.Pool) {
	n := int32(len(*pools))
	newPool.Index = n
	*pools = append(*pools, &newPool)
}

func Pop1to0(pools *([]*types.Pool)) types.Pool {
	old := *pools
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	*pools = old[0 : n-1]
	return *item
}

// update modifies the priority and value of an Item in the queue.
func (k Keeper) Update1to0(pools *([]*types.Pool), pool *types.Pool, reserve0 , reserve1 , fee, totalShares , price sdk.Dec) {
	pool.Reserve0 = reserve0
	pool.Reserve1 = reserve1
	pool.Fee = fee
	pool.Price = price
	pool.TotalShares = totalShares
	Fix1to0(pools, pool.Index)
}

func (k Keeper) Init1to0(pools *([]*types.Pool)) {
	// heapify
	n := Len1to0(*pools)
	for i := n/2 - 1; i >= 0; i-- {
		down1to0(pools, i, n)
	}
}

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func (k Keeper) Push1to0(pools *([]*types.Pool), newPool *types.Pool) {
	Push1to0(pools, *newPool)
	up1to0(pools, Len1to0(*pools)-1)
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to Remove(h, 0).
func (k Keeper) Pop1to0(pools *([]*types.Pool)) types.Pool {
	n := Len1to0(*pools) - 1
	Swap1to0(*pools, 0, n)
	down1to0(pools, 0, n)
	return Pop1to0(pools)
}

// Remove removes and returns the element at index i from the heap.
// The complexity is O(log n) where n = h.Len().
func (k Keeper) Remove1to0(pools *([]*types.Pool), i int32) types.Pool {
	n := Len1to0(*pools) - 1
	if n != i {
		Swap1to0(*pools, i, n)
		if !down1to0(pools, i, n) {
			up1to0(pools, i)
		}
	}
	return Pop1to0(pools)
}

// Fix re-establishes the heap ordering after the element at index i has changed its value.
// Changing the value of the element at index i and then calling Fix is equivalent to,
// but less expensive than, calling Remove(h, i) followed by a Push of the new value.
// The complexity is O(log n) where n = h.Len().
func Fix1to0(pools *([]*types.Pool), i int32) {
	if !down1to0(pools, i, Len1to0(*pools)) {
		up1to0(pools, i)
	}
}

func up1to0(pools *([]*types.Pool), j int32) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !Less1to0(*pools, j, i) {
			break
		}
		Swap1to0(*pools, i, j)
		j = i
	}
}

func down1to0(pools *([]*types.Pool), i0, n int32) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && Less1to0(*pools, j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !Less1to0(*pools, j, i) {
			break
		}
		Swap1to0(*pools, i, j)
		i = j
	}
	return i > i0
}

/////////////////////////////////////////////////////////
// 0 to 1 Priority Queue Functions
/////////////////////////////////////////////////////////

func Len0to1(pools []*types.Pool) int32 { return int32(len(pools)) }

func Swap0to1(pools []*types.Pool, i, j int32) {
	pools[i], pools[j] = pools[j], pools[i]
	pools[i].Index = i
	pools[j].Index = j
}

func Less0to1(pools []*types.Pool, i, j int32) bool {

	
	return  (pools[i].Fee).Quo( pools[i].Price.Mul(sdk.NewDecWithPrec(10000, 18))).LT( (pools[j].Fee).Quo(pools[i].Price.Mul(sdk.NewDecWithPrec(10000, 18)))) 

}

func Push0to1(pools *([]*types.Pool), newPool types.Pool) {
	n := int32(len(*pools))
	newPool.Index = n
	*pools = append(*pools, &newPool)
}

func Pop0to1(pools *([]*types.Pool)) types.Pool {
	old := *pools
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	*pools = old[0 : n-1]
	return *item
}

// update modifies the priority and value of an Item in the queue.
func (k Keeper) Update0to1(pools *([]*types.Pool), pool *types.Pool, reserve0 , reserve1 , fee, totalShares , price sdk.Dec) {
	pool.Reserve0 = reserve0
	pool.Reserve1 = reserve1
	pool.Fee = fee
	pool.Price = price
	pool.TotalShares = totalShares
	Fix0to1(pools, pool.Index)
}

func (k Keeper) Init0to1(pools *([]*types.Pool)) {
	// heapify
	n := Len0to1(*pools)
	for i := n/2 - 1; i >= 0; i-- {
		down0to1(pools, i, n)
	}
}

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func (k Keeper) Push0to1(pools *([]*types.Pool), newPool *types.Pool) {
	Push0to1(pools, *newPool)
	up0to1(pools, Len0to1(*pools)-1)
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to Remove(h, 0).
func (k Keeper) Pop0to1(pools *([]*types.Pool)) types.Pool {
	n := Len0to1(*pools) - 1
	Swap0to1(*pools, 0, n)
	down0to1(pools, 0, n)
	return Pop0to1(pools)
}

// Remove removes and returns the element at index i from the heap.
// The complexity is O(log n) where n = h.Len().
func (k Keeper) Remove0to1(pools *([]*types.Pool), i int32) types.Pool {
	n := Len0to1(*pools) - 1
	if n != i {
		Swap0to1(*pools, i, n)
		if !down0to1(pools, i, n) {
			up0to1(pools, i)
		}
	}
	return Pop1to0(pools)
}

// Fix re-establishes the heap ordering after the element at index i has changed its value.
// Changing the value of the element at index i and then calling Fix is equivalent to,
// but less expensive than, calling Remove(h, i) followed by a Push of the new value.
// The complexity is O(log n) where n = h.Len().
func Fix0to1(pools *([]*types.Pool), i int32) {
	if !down0to1(pools, i, Len0to1(*pools)) {
		up0to1(pools, i)
	}
}

func up0to1(pools *([]*types.Pool), j int32) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !Less0to1(*pools, j, i) {
			break
		}
		Swap0to1(*pools, i, j)
		j = i
	}
}

func down0to1(pools *([]*types.Pool), i0, n int32) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && Less0to1(*pools, j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !Less0to1(*pools, j, i) {
			break
		}
		Swap0to1(*pools, i, j)
		i = j
	}
	return i > i0
}

////////////////////////////////////////////////////////
// Other Helper Functions
////////////////////////////////////////////////////////

func (k Keeper) getPool(pools *([]*types.Pool), Fee, Price sdk.Dec) (types.Pool, bool) {

	for _, s := range *pools {
		if s.Fee == Fee && s.Price == Price {
			return *s, true
		}
	}
	return types.Pool{}, false
}