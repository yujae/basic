package BTree

import (
	"fmt"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func (node *TreeNode) AddNode(val int, dir string) {
	newNode := &TreeNode{val: val}
	if dir == "left" {
		node.left = newNode
	}
	if dir == "right" {
		node.right = newNode
	}
}

func (node *TreeNode) Print() {
	fmt.Println(node.val)
	if node.left != nil {
		node.left.Print()
	}

	if node.right != nil {
		node.right.Print()
	}
}
