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
	RemoveNode(root.next.next)
	PrintNodes(root)

	fmt.Println("---------------------")

	RemoveNode(root)
	RemoveNode(root)
	PrintNodes(root)
}
