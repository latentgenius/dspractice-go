package main

import (
	"fmt"

	"dspractice/binarytrees"
)

func main() {
	build123()
}

func build123() {
	var a *binarytrees.Node

	// insert nodes
	a = binarytrees.Insert(a, 2)
	a = binarytrees.Insert(a, 1)
	a = binarytrees.Insert(a, 3)
	a = binarytrees.Insert(a, 5)
	a = binarytrees.Insert(a, 4)

	// check IsEmpty method
	fmt.Println("Tree is empty?", a.IsEmpty())

	// check MaxDepth method
	fmt.Println("Maximum depth of the tree is", a.MaxDepth())

	// check MinValue method
	fmt.Printf("Minimum value in tree is %d\n", a.MinValue())

	// check PrintTree method
	fmt.Println("Tree printed in-order is:")
	a.PrintTree()
	fmt.Println()

	// check PrintPostOrder method
	fmt.Println("Tree printed post-order is:")
	a.PrintPostOrder()
	fmt.Println()

	// check HasPathSum method
	fmt.Println("Given tree has path sums 3 and 4?")
	fmt.Printf("%v\n", (a.HasPathSum(3) && a.HasPathSum(5)))

	// testing PrintPaths method
	fmt.Println("The branches in the tree:")
	a.PrintPaths()
}
