package avl

import (
	"algorithm/tree/bst"
	"fmt"
)

// Utility to get node height
func height(n *bst.Node) int {
	if n == nil {
		return 0
	}
	return n.Height
}

// Update node height
func (n *bst.Node) updateHeight() {
	lh, rh := height(n.Left), height(n.Right)
	if lh > rh {
		n.Height = lh + 1
	} else {
		n.Height = rh + 1
	}
}

// Calculate balance factor
func (n *bst.Node) balanceFactor() int {
	return height(n.Left) - height(n.Right)
}

// Placeholder for rotation and insertion logic
func (n *bst.Node) Insert_AVL(value int) *bst.Node {
	// AVL Insertion logic should be added here, including rotations
	// For now, this is a placeholder for demonstration.
	return n
}

// ExecuteAVL runs the AVL tree logic as a demo
func ExecuteAVL() {
	fmt.Println("AVL: Demo execution (rotation logic not implemented)")
}
