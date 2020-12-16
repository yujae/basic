package BSTree

import (
	"fmt"
	"testing"
)

/*
		5
      /   \
     3     8
    / \   / \
   2  4  7  10
        /   /
       6   9
*/

func TestBSTree(t *testing.T) {
	root := &TreeNode{val: 5}
	root.AddNode(3)
	root.AddNode(2)
	root.AddNode(4)
	root.AddNode(8)
	root.AddNode(7)
	root.AddNode(6)
	root.AddNode(10)
	root.AddNode(9)

	root.PrintNodes()
	root.PrintNodesWithQueue()
	fmt.Println()
	root.PrintNodesWithStack()
}
