package main_test

import (
	"testing"

	. "github.com/duality-labs/debugstate"
)

func TestHeapInsertAndExtract(t *testing.T) {
	h := NewMaxList(3)
	node1 := &Node{Value: 1}
	node2 := &Node{Value: 2}
	node3 := &Node{Value: 3}

	h.Insert(node1)
	h.Insert(node2)
	h.Insert(node3)

	if len(h.Array) != 3 {
		t.Errorf("Heap size after inserts = %d; want 3", len(h.Array))
	}

	max := h.Array[0]
	if max != node3 {
		t.Errorf("extractMax() = %v; want %v", max, node3)
	}
}

func TestHeapInsertWhenFull(t *testing.T) {
	h := NewMaxList(2)
	node1 := &Node{Value: 1}
	node2 := &Node{Value: 2}
	node3 := &Node{Value: 3}

	h.Insert(node1)
	h.Insert(node2)

	if len(h.Array) != 2 {
		t.Errorf("Heap size after inserts = %d; want 2", len(h.Array))
	}

	h.Insert(node3)

	if len(h.Array) != 2 {
		t.Errorf("Heap size after trying to insert to full heap = %d; want 2", len(h.Array))
	}
}

func TestHeapExtractMaxMultiple(t *testing.T) {
	h := NewMaxList(3)
	node1 := &Node{Value: 1}
	node2 := &Node{Value: 2}
	node3 := &Node{Value: 3}

	h.Insert(node1)
	h.Insert(node2)
	h.Insert(node3)

	max1 := h.Array[0]
	max2 := h.Array[1]
	max3 := h.Array[2]

	if max1 != node3 {
		t.Errorf("First extractMax() = %v; want %v", max1, node3)
	}

	if max2 != node2 {
		t.Errorf("Second extractMax() = %v; want %v", max2, node2)
	}

	if max3 != node1 {
		t.Errorf("Third extractMax() = %v; want %v", max3, node1)
	}
}

func TestHeapify(t *testing.T) {
	h := NewMaxList(4)
	node1 := &Node{Value: 1}
	node2 := &Node{Value: 2}
	node3 := &Node{Value: 3}
	node4 := &Node{Value: 4}

	h.Insert(node2)
	h.Insert(node3)
	h.Insert(node1)
	h.Insert(node4)

	max := h.Array[0]

	if max != node4 {
		t.Errorf("extractMax() = %v; want %v", max, node4)
	}
}
