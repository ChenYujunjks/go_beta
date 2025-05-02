package main

import (
	"algorithm/tree/bst"
)

func ExecuteBSTWithPrint() {
	root := &bst.Node{Value: 10}
	root = root.Insert(5)
	root = root.Insert(222)
	root = root.Insert(122)
	root = root.Insert(422)
	root = root.Insert(15)
	root = root.Insert(3)
	root = root.Insert(7)
	root = root.Insert(12)
	root = root.Insert(17)

	// Print the tree structure
	bst.PrintTree(root)
}

func main() {
	//ExecuteBSTWithPrint()
	bst.ExecuteBST()

	// ExecuteAVL()
	// ExecuteRBT()
}
