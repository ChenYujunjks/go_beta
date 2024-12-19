package main

import "fmt"

func (n *Node) Insert(value int) *Node {
	if n == nil {
		return &Node{Value: value} // 递归基准条件：当前节点为空，创建新节点
	}
	if value < n.Value {
		n.Left = n.Left.Insert(value) // 递归调用，向左子树插入
	} else if value > n.Value {
		n.Right = n.Right.Insert(value) // 递归调用，向右子树插入
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
