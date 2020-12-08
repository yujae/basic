package algorism

import (
	"fmt"
	"testing"
)

func TestAddNode(t *testing.T) {
	root := NewRoot()
	tail := FindTail(root)
	for tail.next != nil {
		tail = tail.next
	}

	for i := 1; i <= 10; i++ {
		tail = AddNode(tail, i)
	}
	PrintNodes(root)
}

func TestRemoveNode(t *testing.T) {
	root := NewRoot()
	tail := FindTail(root)
	for tail.next != nil {
		tail = tail.next
	}

	for i := 1; i <= 10; i++ {
		tail = AddNode(tail, i)
	}
	PrintNodes(root)
	root, tail = RemoveNode(root.next.next, root, tail)
	PrintNodes(root)

	fmt.Println("---------------------")

	root, tail = RemoveNode(root, root, tail)
	root, tail = RemoveNode(root, root, tail)
	PrintNodes(root)
}
