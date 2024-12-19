package rbt

import (
	"fmt"

	"github.com/emirpasic/gods/trees/redblacktree"
)

// ExecuteRBT runs the Red-Black Tree logic as a demo
func ExecuteRBT() {
	tree := redblacktree.NewWithIntComparator()
	tree.Put(1, "one")
	tree.Put(2, "two")
	tree.Put(3, "three")
	value, found := tree.Get(2)
	if found {
		fmt.Println("RBT: Found value for key 2:", value) // Output: two
	} else {
		fmt.Println("RBT: Key 2 not found")
	}
}
