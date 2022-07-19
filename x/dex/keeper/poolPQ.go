package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"strconv"
	//"fmt"
)

func Len(pools []*types.Pool) int32 { return int32(len(pools)) }

func Swap(pools []*types.Pool, i, j int32) {
	pools[i], pools[j] = pools[j], pools[i]
	pools[i].Index = i
	pools[j].Index = j
}

func Less(pools []*types.Pool, i, j int32) bool {

	priceI, error := strconv.ParseFloat(pools[i].Price, 64)

	feeI, error := strconv.ParseFloat(pools[i].Fee, 64)

	priceJ, error := strconv.ParseFloat(pools[j].Price, 64)

	feeJ, error := strconv.ParseFloat(pools[j].Fee, 64)

	_ = error
	return (priceI * (1 - feeI)) > (priceJ * (1 - feeJ))

}

func Push(pools *([]*types.Pool), newPool types.Pool) {
	n := int32(len(*pools))
	newPool.Index = n
	*pools = append(*pools, &newPool)
}

func Pop(pools *([]*types.Pool)) types.Pool {
	old := *pools
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	*pools = old[0 : n-1]
	return *item
}

// update modifies the priority and value of an Item in the queue.
func (k Keeper) Update(pools *([]*types.Pool), pool *types.Pool, reserveA string, reserveB string, fee string, totalShares string, price string) {
	pool.ReserveA = reserveA
	pool.ReserveB = reserveB
	pool.Fee = fee
	pool.Price = price
	pool.TotalShares = totalShares
	Fix(pools, pool.Index)
}

func (k Keeper) Init(pools *([]*types.Pool)) {
	// heapify
	n := Len(*pools)
	for i := n/2 - 1; i >= 0; i-- {
		down(pools, i, n)
	}
}

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func (k Keeper) Push(pools *([]*types.Pool), newPool *types.Pool) {
	Push(pools, *newPool)
	up(pools, Len(*pools)-1)
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to Remove(h, 0).
func (k Keeper) Pop(pools *([]*types.Pool)) types.Pool {
	n := Len(*pools) - 1
	Swap(*pools, 0, n)
	down(pools, 0, n)
	return Pop(pools)
}

// Remove removes and returns the element at index i from the heap.
// The complexity is O(log n) where n = h.Len().
func (k Keeper) Remove(pools *([]*types.Pool), i int32) types.Pool {
	n := Len(*pools) - 1
	if n != i {
		Swap(*pools, i, n)
		if !down(pools, i, n) {
			up(pools, i)
		}
	}
	return Pop(pools)
}

// Fix re-establishes the heap ordering after the element at index i has changed its value.
// Changing the value of the element at index i and then calling Fix is equivalent to,
// but less expensive than, calling Remove(h, i) followed by a Push of the new value.
// The complexity is O(log n) where n = h.Len().
func Fix(pools *([]*types.Pool), i int32) {
	if !down(pools, i, Len(*pools)) {
		up(pools, i)
	}
}

func up(pools *([]*types.Pool), j int32) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !Less(*pools, j, i) {
			break
		}
		Swap(*pools, i, j)
		j = i
	}
}

func down(pools *([]*types.Pool), i0, n int32) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && Less(*pools, j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !Less(*pools, j, i) {
			break
		}
		Swap(*pools, i, j)
		i = j
	}
	return i > i0
}


func (k Keeper) getPool(pools *([]*types.Pool), Fee, Price string) (types.Pool, bool) {

	for _, s := range *pools {
		if s.Fee == Fee && s.Price == Price {
			return s, true
		}
	}
	return nil, false
}