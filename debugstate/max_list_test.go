package main_test

import (
	"testing"

	. "github.com/duality-labs/debugstate"
)

func TestHeapInsertAndExtract(t *testing.T) {
	h := NewMaxList(3)
	trienode1 := &TrieNode{Value: 1}
	trienode2 := &TrieNode{Value: 2}
	trienode3 := &TrieNode{Value: 3}

	h.Insert(trienode1)
	h.Insert(trienode2)
	h.Insert(trienode3)

	if len(h.Array) != 3 {
		t.Errorf("Heap size after inserts = %d; want 3", len(h.Array))
	}

	max := h.Array[0]
	if max != trienode3 {
		t.Errorf("extractMax() = %v; want %v", max, trienode3)
	}
}

func TestHeapInsertWhenFull(t *testing.T) {
	h := NewMaxList(2)
	trienode1 := &TrieNode{Value: 1}
	trienode2 := &TrieNode{Value: 2}
	trienode3 := &TrieNode{Value: 3}

	h.Insert(trienode1)
	h.Insert(trienode2)

	if len(h.Array) != 2 {
		t.Errorf("Heap size after inserts = %d; want 2", len(h.Array))
	}

	h.Insert(trienode3)

	if len(h.Array) != 2 {
		t.Errorf("Heap size after trying to insert to full heap = %d; want 2", len(h.Array))
	}
}

func TestHeapExtractMaxMultiple(t *testing.T) {
	h := NewMaxList(3)
	trienode1 := &TrieNode{Value: 1}
	trienode2 := &TrieNode{Value: 2}
	trienode3 := &TrieNode{Value: 3}

	h.Insert(trienode1)
	h.Insert(trienode2)
	h.Insert(trienode3)

	max1 := h.Array[0]
	max2 := h.Array[1]
	max3 := h.Array[2]

	if max1 != trienode3 {
		t.Errorf("First extractMax() = %v; want %v", max1, trienode3)
	}

	if max2 != trienode2 {
		t.Errorf("Second extractMax() = %v; want %v", max2, trienode2)
	}

	if max3 != trienode1 {
		t.Errorf("Third extractMax() = %v; want %v", max3, trienode1)
	}
}

func TestHeapify(t *testing.T) {
	h := NewMaxList(4)
	trienode1 := &TrieNode{Value: 1}
	trienode2 := &TrieNode{Value: 2}
	trienode3 := &TrieNode{Value: 3}
	trienode4 := &TrieNode{Value: 4}

	h.Insert(trienode2)
	h.Insert(trienode3)
	h.Insert(trienode1)
	h.Insert(trienode4)

	max := h.Array[0]

	if max != trienode4 {
		t.Errorf("extractMax() = %v; want %v", max, trienode4)
	}
}
