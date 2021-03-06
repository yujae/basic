package DFS

import "testing"

/*
       1
    /  |  \
   2   3   4
  /\   /\  / \
 5  6 7  8 9  10

깊이우선탐색
결과 : 1 2 5 6 3 7 8 4 9 10
*/

func TestDFS(t *testing.T) {
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

	root.DFS()
}

func TestStackDFS(t *testing.T) {
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

	root.StackDFS()
}
