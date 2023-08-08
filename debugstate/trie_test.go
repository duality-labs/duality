package main_test

import (
	"testing"

	. "github.com/duality-labs/debugstate"
)

func TestTrieNodeCreation(t *testing.T) {
	n := NewTrieNode('A')
	if n.Value != 'A' || len(n.Children) != 0 || n.Count != 0 {
		t.Errorf("NewTrieNode('A') - got: %v; want value: 'A', children: 0, count: 0", n)
	}
}

func TestTrieNodeInsertion(t *testing.T) {
	n := NewTrieNode('A')
	n.Insert([]byte("BCD"), 1)
	if n.Count != 1 {
		t.Errorf("Root TrieNode count after insertion - got: %d; want: 1", n.Count)
	}
	if child, ok := n.Children['B']; !ok {
		t.Errorf("Child 'B' not found after insertion")
	} else if child.Value != 'B' || len(child.Children) != 1 || child.Count != 1 {
		t.Errorf("TrieNode 'B' properties - got: %v; want value: 'B', children: 1, count: 1", child)
	}
}

func TestTrieNodePruneSmallTrieNodes(t *testing.T) {
	n := NewTrieNode('A')
	n.Insert([]byte("BCE"), 1)
	n.Insert([]byte("BCD"), 1)
	n.PrintStats()
}

func TestTrieNodePrintStats(t *testing.T) {
	// This function needs a way to capture the output if you want to test it directly.
	// An alternative way could be to refactor the function to return a string instead of directly printing to stdout.
}
