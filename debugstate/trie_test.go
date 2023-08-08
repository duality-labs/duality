package main_test

import (
	"testing"

	. "github.com/duality-labs/debugstate"
)

func TestNodeCreation(t *testing.T) {
	n := NewNode('A')
	if n.Value != 'A' || len(n.Children) != 0 || n.Count != 0 {
		t.Errorf("NewNode('A') - got: %v; want value: 'A', children: 0, count: 0", n)
	}
}

func TestNodeInsertion(t *testing.T) {
	n := NewNode('A')
	n.Insert([]byte("BCD"))
	if n.Count != 1 {
		t.Errorf("Root node count after insertion - got: %d; want: 1", n.Count)
	}
	if child, ok := n.Children['B']; !ok {
		t.Errorf("Child 'B' not found after insertion")
	} else if child.Value != 'B' || len(child.Children) != 1 || child.Count != 1 {
		t.Errorf("Node 'B' properties - got: %v; want value: 'B', children: 1, count: 1", child)
	}
}

func TestNodePruneSmallNodes(t *testing.T) {
	n := NewNode('A')
	n.Insert([]byte("BCE"))
	n.Insert([]byte("BCD"))
	pruned := n.PruneSmallNodes(2)
	if _, ok := pruned.Children['E']; ok {
		t.Errorf("Child 'E' found after pruning, should have been pruned")
	}
}

func TestNodePrintStats(t *testing.T) {
	// This function needs a way to capture the output if you want to test it directly.
	// An alternative way could be to refactor the function to return a string instead of directly printing to stdout.
}
