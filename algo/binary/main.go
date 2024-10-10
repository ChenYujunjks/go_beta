package main

import "fmt"

// Node represents a node in the binary tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Insert adds a value to the binary tree
func (n *Node) Insert(value int) {
	if value <= n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
}

// InOrderTraversal prints the values of the tree in-order
func (n *Node) InOrderTraversal() {
	if n == nil {
		return
	}
	n.Left.InOrderTraversal()
	fmt.Println(n.Value)
	n.Right.InOrderTraversal()
}

func main() {
	root := &Node{Value: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(12)
	root.Insert(18)

	fmt.Println("In-order traversal of the binary tree:")
	root.InOrderTraversal()
}
