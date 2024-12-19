package bst

import "fmt"

// Define a Node structure
type Node struct {
	Value  int
	Left   *Node
	Right  *Node
	Height int
}

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

// getTreeDepth calculates the depth of the tree
func GetTreeHeight(node *Node) int {
	if node == nil {
		return 0
	}
	leftDepth := GetTreeHeight(node.Left)
	rightDepth := GetTreeHeight(node.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

// Binary Search
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

// 中序遍历：左子树 -> 当前节点 -> 右子树
func inorderTraversal(node *Node) {
	if node == nil {
		return
	}
	inorderTraversal(node.Left)
	fmt.Printf("%d ", node.Value)
	inorderTraversal(node.Right)
}

// ExecuteBST runs the BST logic as a demo
func ExecuteBST() {
	root := &Node{Value: 20}
	// 插入多个节点
	values := []int{10, 30, 5, 15, 25, 35, 2, 7, 12, 17, 22, 7, 200}
	for _, v := range values {
		root = root.Insert(v)
	}
	// 打印树的高度
	fmt.Println("BST: Tree Height =>", GetTreeHeight(root))
	// 搜索一些节点
	searchValues := []int{15, 7, 40, 22}
	for _, v := range searchValues {
		fmt.Printf("BST: Search for %d => %v\n", v, root.Search(v))
	}

	// 打印树的中序遍历（从小到大的排序）
	fmt.Print("BST: Inorder Traversal => ")
	inorderTraversal(root)
	fmt.Println()
	PrintTree(root)
}
