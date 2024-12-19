package main

import "fmt"

// Define a Node structure
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Insert a value into BST
func (n *Node) Insert(value int) *Node {
	if n == nil {
		return &Node{Value: value}
	}
	if value < n.Value {
		n.Left = n.Left.Insert(value)
	} else if value > n.Value {
		n.Right = n.Right.Insert(value)
	}
	return n
}

// Search for a value
func (n *Node) Search(value int) bool {
	if n == nil {
		return false
	}
	if value == n.Value {
		return true
	} else if value < n.Value {
		return n.Left.Search(value)
	} else {
		return n.Right.Search(value)
	}
}

// ExecuteBST runs the BST logic as a demo
func ExecuteBST() {
	root := &Node{Value: 10}
	root = root.Insert(5)
	root = root.Insert(15)
	fmt.Println("BST: Search for 15:", root.Search(15)) // Output: true
}
