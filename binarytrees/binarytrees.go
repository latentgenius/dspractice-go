package binarytrees

import (
	"fmt"
)

// Node is a node in a binary tree
type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

// IsEmpty returns true if a binary tree is empty, false otherwise
func (n *Node) IsEmpty() bool {
	return n == nil
}

// HasNoChildren returns true if the current node has no children, false
// otherwise
func (n *Node) HasNoChildren() bool {
	return n.Left.IsEmpty() && n.Right.IsEmpty()
}

// Insert adds a node into a binary tree
// creates a new node if the given binary tree is empty
func Insert(n *Node, data int) *Node {
	if n.IsEmpty() {
		return NewNode(data)
	} else if data < n.Data {
		n.Left = Insert(n.Left, data)
	} else if data > n.Data {
		n.Right = Insert(n.Right, data)
	}
	return n
}

// NewNode returns a node with the given data
func NewNode(data int) *Node {
	return &Node{
		Data:  data,
		Left:  nil,
		Right: nil,
	}
}

// MaxDepth returns the maximum depth of the tree
func (n *Node) MaxDepth() int {
	if n.IsEmpty() {
		return 0
	}
	lDepth := n.Left.MaxDepth()
	rDepth := n.Right.MaxDepth()

	if lDepth < rDepth {
		return rDepth + 1
	}
	return lDepth + 1
}

// MinValue returns the minimum data value found in the tree
func (n *Node) MinValue() int {
	if n.IsEmpty() {
		fmt.Println("Empty tree")

	}
	if n.Left.IsEmpty() {
		return n.Data
	}
	return n.Left.MinValue()
}

// PrintTree prints the elements of the tree in-order
func (n *Node) PrintTree() {
	if n.IsEmpty() {
		return
	}
	n.Left.PrintTree()
	fmt.Print(n.Data, " ")
	n.Right.PrintTree()
}

// PrintPostOrder prints the tree post-order
func (n *Node) PrintPostOrder() {
	if n.IsEmpty() {
		return
	}
	n.Left.PrintPostOrder()
	n.Right.PrintPostOrder()
	fmt.Print(n.Data, " ")
}

// HasPathSum returns true if the given binary tree has a branch
// whose elements add up to the given sum
func (n *Node) HasPathSum(sum int) bool {
	if n.IsEmpty() {
		return sum == 0
	}
	subSum := sum - n.Data
	return n.Left.HasPathSum(subSum) || n.Right.HasPathSum(subSum)
}

// PrintPaths prints all the root-to-leaf paths of a given binary tree,
// one per line
func (n *Node) PrintPaths() {
	var pathSlc = make([]int, 0, n.MaxDepth())
	f(pathSlc, n)

}

// printSlc is a helper function for f
func printSlc(slc []int) {
	for _, x := range slc {
		fmt.Print(x, " ")
	}
	fmt.Println()
}

// f is a helper function for PrintPaths
func f(slc []int, n *Node) {
	// node contains data?
	if n.IsEmpty() {
		return
	}
	// add current node data to slice
	slc = append(slc, n.Data)
	// n is a leaf node?
	if n.HasNoChildren() {
		printSlc(slc)
		return
	}
	f(slc, n.Left)
	f(slc, n.Right)
}

// Mirror returns a mirror image of the given tree
// with the children of every node swapped left-to-right
func (n *Node) Mirror() *Node {
	if n.IsEmpty() {
		return nil
	}
	newTree := NewNode(n.Data)
	newTree.Left = n.Left.Mirror()
	newTree.Right = n.Right.Mirror()

	newTree.Left, newTree.Right = newTree.Right, newTree.Left
	return newTree
}

// Double adds a clone of each node as its left child in each subtree
func (n *Node) Double() {
	if n.IsEmpty() {
		return
	}
	nClone := NewNode(n.Data)
	n.Left.Double()
	n.Right.Double()
	nClone.Left, n.Left = n.Left, nClone
}

// SameTree compares the current tree to the given tree. Returns true if both
// are similar in structure and members, false otherwise
func (n *Node) SameTree(otherTree *Node) bool {
	if n.IsEmpty() && otherTree.IsEmpty() {
		return true
	}
	if n.IsEmpty() || otherTree.IsEmpty() {
		return false
	}
	if n.Data == otherTree.Data {
		return n.Left.SameTree(otherTree.Left) && n.Right.SameTree(otherTree.Right)
	}
	return false
}