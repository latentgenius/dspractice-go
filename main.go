package main

import (
	"fmt"

	"dspractice-go/binarytrees"
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
	fmt.Printf("%v\n", a.HasPathSum(3) && a.HasPathSum(5))

	// testing PrintPaths method
	fmt.Println("The branches in the tree:")
	a.PrintPaths()

	// testing Mirror method
	fmt.Println("Original tree:")
	a.PrintTree()
	fmt.Println("\nMirror tree:")
	b := a.Mirror()
	b.PrintTree()
	fmt.Println()

	// testing Double Method
	fmt.Println("Doubling the mirror tree 'b':")
	b.Double()
	fmt.Println("'b' is now:")
	b.PrintTree()
	fmt.Println()

	// testing SameTree method
	fmt.Println("Mirroring 'a' to have 'c' not same as 'a':")
	c := a.Mirror()
	fmt.Print("c is now: ")
	c.PrintTree()
	fmt.Println()
	fmt.Println("Calling same tree on 'a' with 'c' as argument:")
	d := a.SameTree(c)
	fmt.Println(d)

	fmt.Println("Mirroring 'c' to have 'c' same as 'a':")
	c = c.Mirror()
	fmt.Print("c is now: ")
	c.PrintTree()
	fmt.Println()
	fmt.Println("Calling same tree on 'a' with 'c' as argument:")
	d = a.SameTree(c)
	fmt.Println(d)

	// testing CountTrees function
	fmt.Println("Trees with BST with 17 nodes:")
	trees := binarytrees.CountTrees(17)
	fmt.Println(trees)
	// testing CountTreesD function
	fmt.Println("Trees with BST with 17 nodes using dynamic version:")
	trees = binarytrees.CountTreesD(17)
	fmt.Println(trees)
}
