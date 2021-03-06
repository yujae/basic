package Tree

import (
	"testing"
)

func TestTreeNode_AddNode(t *testing.T) {
	num := 1
	root := &TreeNode{val: num}

	for i := 0; i < 3; i++ {
		num++
		root.AddNode(num)
	}

	for k := 0; k < len(root.children); k++ {
		for j := 0; j < 2; j++ {
			num++
			root.children[k].AddNode(num)
		}
	}

	root.Print()
}
