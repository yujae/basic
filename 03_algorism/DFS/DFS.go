package DFS

import "fmt"

type TreeNode struct {
	val      int
	children []*TreeNode
}

func (node *TreeNode) AddNode(val int) {
	node.children = append(node.children, &TreeNode{val: val})
}

func (node *TreeNode) DFS() {
	fmt.Println(node.val)
	for i := 0; i < len(node.children); i++ {
		node.children[i].DFS()
	}
}

/*
	0. root 노드 푸쉬 (stack 최상단)
	1. stack 최상단 팝 & 자식 노드들 있으면 푸쉬
	2. 1 반복
*/
func (node *TreeNode) StackDFS() {
	var stack []*TreeNode
	stack = append(stack, node) // push

	for len(stack) > 0 {
		var top *TreeNode
		top, stack = stack[len(stack)-1], stack[:len(stack)-1] // pop
		fmt.Println(top.val)
		for i := 0; i < len(top.children); i++ {
			stack = append(stack, top.children[i]) // push
		}
	}
}
