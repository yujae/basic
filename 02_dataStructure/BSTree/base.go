package BSTree

import "fmt"

/*
	Binary Search Tree
	- 부모는 자식을 좌우 한개씩 갖음
	- 부모보다 작은 것은 좌, 부모보다 큰 것은 우
	- 중복될 수 없다
*/

type TreeNode struct {
	val   int
	dir   string
	left  *TreeNode
	right *TreeNode
}

func (node *TreeNode) AddNode(val int) {
	n := &TreeNode{val: val}
	if node.val >= val {
		if node.left != nil {
			node.left.AddNode(val)
			return
		}
		node.left = n
		node.left.dir = "left"
	} else {
		if node.right != nil {
			node.right.AddNode(val)
			return
		}
		node.right = n
		node.right.dir = "right"
	}
}

// 결국 재귀호출도 스택
func (node *TreeNode) PrintNodes() {
	fmt.Println(node.val, node.dir)
	if node.left != nil {
		node.left.PrintNodes()
	}
	if node.right != nil {
		node.right.PrintNodes()
	}
}

// BFS
func (node *TreeNode) PrintNodesWithQueue() {
	var q []*TreeNode
	var first *TreeNode

	q = append(q, node)
	for len(q) > 0 {
		first, q = q[0], q[1:]
		fmt.Printf("%d ", first.val)

		if first.left != nil {
			q = append(q, first.left)
		}
		if first.right != nil {
			q = append(q, first.right)
		}
	}
}

// DFS
func (node *TreeNode) PrintNodesWithStack() {
	var s []*TreeNode
	var last *TreeNode

	s = append(s, node)
	for len(s) > 0 {
		last, s = s[len(s)-1], s[:len(s)-1]
		fmt.Printf("%d ", last.val)

		if last.right != nil {
			s = append(s, last.right)
		}
		if last.left != nil {
			s = append(s, last.left)
		}
	}
}
