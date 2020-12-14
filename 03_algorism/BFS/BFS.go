package BFS

import "fmt"

type TreeNode struct {
	val      int
	children []*TreeNode
}

func (node *TreeNode) AddNode(val int) {
	node.children = append(node.children, &TreeNode{val: val})
}

/*
	0. root 노드 푸쉬 (queue 최앞단)
	1. queue 최앞단 팝 & 자식 노드들 있으면 푸쉬
	2. 1 반복
*/
func (node *TreeNode) QueueBFS() {
	var queue []*TreeNode
	queue = append(queue, node) // push

	for len(queue) > 0 {
		var first *TreeNode
		first, queue = queue[0], queue[1:] // pop
		fmt.Println(first.val)

		for i := 0; i < len(first.children); i++ {
			queue = append(queue, first.children[i]) // push
		}
	}
}
