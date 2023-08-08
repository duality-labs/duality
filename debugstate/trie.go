package main

import "fmt"

type TrieNode struct {
	Children map[byte]*TrieNode
	Value    byte
	Count    int
}

func NewTrieNode(value byte) *TrieNode {
	return &TrieNode{
		Children: make(map[byte]*TrieNode),
		Value:    value,
		Count:    0,
	}
}

func (n *TrieNode) Insert(key []byte, weight int) {
	curTrieNode := n
	curTrieNode.Count += weight
	for _, c := range key {
		if sub, ok := curTrieNode.Children[c]; ok {
			curTrieNode = sub
		} else {
			nextTrieNode := NewTrieNode(c)
			curTrieNode.Children[c] = nextTrieNode
			curTrieNode = nextTrieNode
		}
		curTrieNode.Count += weight
	}
}

func (n *TrieNode) MaxLen() int {
	max := 0
	if len(n.Children) == 0 {
		return 0
	}
	for _, c := range n.Children {
		cMax := c.MaxLen()
		if cMax > max {
			max = cMax
		}
	}
	return max + 1
}

func (n *TrieNode) PrintStats() {
	pruned := n.PruneSmallTrieNodes(1)
	cur := pruned

	fmt.Printf(
		"%d | ",
		pruned.Index(
			[]byte{115, 47, 107, 58, 115, 108, 97, 115, 104, 105, 110, 103, 47, 110},
		).Count,
	)
	for len(cur.List) > 0 {
		cur = cur.List[0]
		fmt.Printf("%d ", cur.Value)
	}
	fmt.Printf("\n")
}

type MaxListNode struct {
	List  []*MaxListNode
	Value byte
	Count int
}

func (n *MaxListNode) Index(key []byte) *MaxListNode {
	cur := n
	for _, c := range key {
		found := false
		for _, h := range cur.List {
			if h.Value == c {
				cur = h
				found = true
				break
			}
		}
		if !found {
			return nil
		}
	}
	return cur
}

func (n *TrieNode) PruneSmallTrieNodes(keep int) *MaxListNode {
	heap := NewMaxList(keep)
	for _, n := range n.Children {
		heap.Insert(n)
	}
	top := &MaxListNode{
		Value: n.Value,
		Count: n.Count,
		List:  make([]*MaxListNode, len(heap.Array)),
	}
	for i, e := range heap.Array {
		top.List[i] = e.PruneSmallTrieNodes(keep)
	}
	return top
}
