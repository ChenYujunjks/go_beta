package bst

import (
	"fmt"
	"strconv"
	"strings"
)

// buildTreeLevels builds each level of the tree as a string
func buildTreeLevels(node *Node, depth, level int, res *[]string, pos int, width int) {
	if node == nil {
		return
	}

	// Add the current node's value to the appropriate position
	line := (*res)[level]
	valueStr := strconv.Itoa(node.Value)
	nodePos := pos + (width / 2)
	line = line[:nodePos] + valueStr + line[nodePos+len(valueStr):]
	(*res)[level] = line

	// Connect left and right children
	if node.Left != nil {
		(*res)[level+1] = (*res)[level+1][:pos+width/4] + "/" + (*res)[level+1][pos+width/4+1:]
		buildTreeLevels(node.Left, depth, level+2, res, pos, width/2)
	}
	if node.Right != nil {
		(*res)[level+1] = (*res)[level+1][:pos+3*width/4] + "\\" + (*res)[level+1][pos+3*width/4+1:]
		buildTreeLevels(node.Right, depth, level+2, res, pos+width/2, width/2)
	}
}

// PrintTree uses / and \ to visualize the tree structure with proper alignment
func PrintTree(root *Node) {
	depth := GetTreeHeight(root)
	width := int(1<<(depth+1)) - 1 // Maximum width of the tree representation
	res := make([]string, depth*2-1)

	// Initialize empty lines
	for i := range res {
		res[i] = strings.Repeat(" ", width)
	}

	buildTreeLevels(root, depth, 0, &res, 0, width)
	for _, line := range res {
		fmt.Println(line)
	}
}

