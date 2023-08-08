package main

type MaxList struct {
	Array []*Node
	Size  int
}

func NewMaxList(keep int) MaxList {
	return MaxList{
		Array: []*Node{},
		Size:  keep,
	}
}

func (n *MaxList) Insert(item *Node) {
	i := len(n.Array) - 1
	for ; 0 <= i; i-- {
		e := n.Array[i]
		if e.Value > item.Value {
			break
		}
	}
	loc := i + 1
	if loc < len(n.Array) {
		nextLen := len(n.Array) + 1
		if nextLen > n.Size {
			nextLen = n.Size
		}
		next := make([]*Node, 0, nextLen)
		next = append(next, n.Array[0:loc]...)
		next = append(next, item)
		if len(n.Array) == n.Size {
			next = append(next, n.Array[loc:len(n.Array)-1]...)
		} else {
			next = append(next, n.Array[loc:len(n.Array)]...)
		}
		n.Array = next
	} else {
		if len(n.Array) < n.Size {
			next := append(n.Array, item)
			n.Array = next
		}
	}
}
