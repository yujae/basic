package Tree

import "fmt"

type TreeNode struct {
	val      int
	children []*TreeNode
}

func (node *TreeNode) AddNode(val int) {
	node.children = append(node.children, &TreeNode{val: val})
}

func (node *TreeNode) Print() {
	fmt.Println(node.val)
	for i := 0; i < len(node.children); i++ {
		node.children[i].Print()
	}
}
