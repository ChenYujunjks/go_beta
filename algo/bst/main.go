package main

import "fmt"

// Define a Node structure
type Node struct {
	Value  int
	Left   *Node
	Right  *Node
	Height int
}

// PrintTree is a helper function to visualize the BST
func PrintTree(node *Node, prefix string, isLeft bool) {
	if node == nil {
		return
	}

	// Print the current node
	fmt.Printf("%s", prefix)
	if isLeft {
		fmt.Printf("├── ")
	} else {
		fmt.Printf("└── ")
	}
	fmt.Printf("%d\n", node.Value)

	// Prepare the next prefix for left and right children
	newPrefix := prefix
	if isLeft {
		newPrefix += "│   "
	} else {
		newPrefix += "    "
	}

	// Recursively print the left and right subtree
	PrintTree(node.Left, newPrefix, true)
	PrintTree(node.Right, newPrefix, false)
}

// ExecuteBSTWithPrint demonstrates the tree structure and visualization
func ExecuteBSTWithPrint() {
	// Create the BST
	root := &Node{Value: 10}
	root = root.Insert(5)
	root = root.Insert(15)
	root = root.Insert(3)
	root = root.Insert(7)
	root = root.Insert(12)
	root = root.Insert(17)

	// Print the tree structure
	fmt.Println("Binary Search Tree:")
	PrintTree(root, "", false)
}
func main() {
	ExecuteBSTWithPrint()
	//ExecuteBST()

	// ExecuteAVL()

	// ExecuteRBT()
}
