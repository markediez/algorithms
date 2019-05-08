package main

import (
	"fmt"
)

type Node struct {
	Value      int
	LeftChild  *Node
	RightChild *Node
	Parent     *Node
}

func (self Node) IsFull() bool {
	return self.LeftChild != nil && self.RightChild != nil
}
func (self Node) IsLeaf() bool {
	return self.LeftChild == nil && self.RightChild == nil
}
func (self Node) EQ(rhs *Node) bool {
	return self.Value == rhs.Value
}
func (self Node) GT(rhs *Node) bool {
	return self.Value > rhs.Value
}
func (self Node) LT(rhs *Node) bool {
	return self.Value < rhs.Value
}
func (self Node) GTE(rhs *Node) bool {
	return self.Value >= rhs.Value
}
func (self Node) LTE(rhs *Node) bool {
	return self.Value <= rhs.Value
}
func (self *Node) SwapValue(rhs *Node) {
	tmp := self.Value
	self.Value = rhs.Value
	rhs.Value = tmp
}

type MaxHeap struct {
	Head  *Node
	Tails []*Node
}

func (self *MaxHeap) Push(n int) {
	newNode := &Node{Value: n}
	if self.Head == nil {
		self.Head = newNode
		self.Tails = []*Node{newNode}
	} else {
		self.NewTail(newNode)
		self.BubbleUp(newNode)
	}
}

func (self *MaxHeap) Pop() int {
	out := self.Head.Value
	tail := self.Tails[len(self.Tails) - 1]
	if tail != self.Head {
		self.Head.SwapValue(tail)

		if tail == tail.Parent.LeftChild {
			tail.Parent.LeftChild = nil
		} else {
			tail.Parent.RightChild = nil
		}
		if tail.Parent.IsLeaf() {
			self.Tails = append([]*Node{tail.Parent}, self.Tails...)
		}
		self.BubbleDown(self.Head)
	} else {
		self.Head = nil
	}
	self.Tails = self.Tails[0:len(self.Tails) - 1]	
	
	return out
}

func (self *MaxHeap) BubbleDown(n *Node) {
	if !n.IsLeaf() {
		ret := MaxNode(n)
		if ret != n {
			self.BubbleDown(ret)
		}
	}
}

func MaxNode(n *Node) *Node {
	ret := n
	if n.LeftChild != nil && n.LeftChild.GT(n) {
		n.SwapValue(n.LeftChild)
		ret = n.LeftChild
	}
	if n.RightChild != nil && n.RightChild.GT(n) {
		n.SwapValue(n.RightChild)
		ret = n.RightChild
	}
	
	return ret
}

func (self *MaxHeap) BubbleUp(n *Node) {
	if n != self.Head && n.GT(n.Parent) {
		n.SwapValue(n.Parent)
		self.BubbleUp(n.Parent)
	}
}

func (self *MaxHeap) NewTail(n *Node) {
	if self.Tails[0].IsFull() {
		self.Tails = self.Tails[1:len(self.Tails)]
	}
	
	n.Parent = self.Tails[0]
	if self.Tails[0].LeftChild == nil {
		self.Tails[0].LeftChild = n
	} else {
		self.Tails[0].RightChild = n
	}

	self.Tails = append(self.Tails, n)
}

func (self MaxHeap) Print() {
	var currNode *Node
	toPrint := []*Node{self.Head}
	
	for len(toPrint) > 0 {
		currNode = toPrint[0]
		toPrint = toPrint[1:len(toPrint)]
		
		fmt.Println(currNode.Value)
		if currNode.LeftChild != nil {
			toPrint = append(toPrint, currNode.LeftChild)
		}
		if currNode.RightChild != nil {
			toPrint = append(toPrint, currNode.RightChild)
		}
	}
}

func main() {
	h := MaxHeap{}
	h.Push(7)
	h.Push(3)
	h.Push(5)
	h.Push(19)
	h.Push(27)
	h.Push(25)
	h.Push(6)
	h.Print()
	
	fmt.Printf("\n\n\n")
	for h.Head != nil {
		fmt.Println(h.Pop())
	}
}
