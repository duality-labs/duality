package main

import "fmt"

type Node struct {
	Children map[byte]*Node
	Value    byte
	Count    int
}

func NewNode(value byte) *Node {
	return &Node{
		Children: make(map[byte]*Node),
		Value:    value,
		Count:    0,
	}
}

func (n *Node) Insert(key []byte, weight int) {
	curNode := n
	curNode.Count += weight
	for _, c := range key {
		if sub, ok := curNode.Children[c]; ok {
			curNode = sub
		} else {
			nextNode := NewNode(c)
			curNode.Children[c] = nextNode
			curNode = nextNode
		}
		curNode.Count += weight
	}
}

func (n *Node) MaxLen() int {
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

func (n *Node) PrintStats() {
	pruned := n.PruneSmallNodes(1)
	cur := pruned
	for len(cur.Children) > 0 {
		for _, v := range cur.Children {
			fmt.Printf("%d ", v.Value)
			cur = v
		}
	}
	fmt.Printf("\n")
}

func (n *Node) PruneSmallNodes(keep int) *Node {
	heap := NewMaxList(keep)
	for _, n := range n.Children {
		heap.Insert(n)
	}
	top := NewNode(n.Value)
	for _, e := range heap.Array {
		top.Children[e.Value] = e.PruneSmallNodes(keep)
	}
	return top
}
