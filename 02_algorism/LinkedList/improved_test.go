package LinkedList

import (
	"testing"
)

func TestImprovedAddNode(t *testing.T) {
	linkedList := &LinkedList{}

	for i := 1; i <= 10; i++ {
		linkedList.AddNode(i)
	}
	linkedList.PrintNodes()
}

func TestImprovedRemoveNode(t *testing.T) {
	linkedList := &LinkedList{}

	for i := 1; i <= 10; i++ {
		linkedList.AddNode(i)
	}

	linkedList.PrintNodes()

	linkedList.RemoveNode(linkedList.root)
	linkedList.PrintNodes()

	linkedList.RemoveNode(linkedList.root.next)
	linkedList.PrintNodes()

	linkedList.RemoveNode(linkedList.tail)
	linkedList.PrintNodes()
}
