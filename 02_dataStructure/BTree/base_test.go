package BTree

import "testing"

/*
		0
      /   \
     1     2
    / \   / \
   3   4 5   6
*/

func TestAddNode(t *testing.T) {
	root := &TreeNode{val: 0}
	root.AddNode(1, "left")
	root.AddNode(2, "right")

	root.left.AddNode(3, "left")
	root.left.AddNode(4, "right")

	root.right.AddNode(5, "left")
	root.right.AddNode(6, "right")

	root.Print()
}
