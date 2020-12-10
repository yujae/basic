package DoubleLinkedList

import (
	"testing"
)

func TestAddNode(t *testing.T) {
	list := &DoubleLinkedList{}

	for i := 1; i <= 10; i++ {
		list.AddNode(i)
	}
	list.PrintNodes()
}

func TestRemoveNode(t *testing.T) {
	list := &DoubleLinkedList{}

	for i := 1; i <= 10; i++ {
		list.AddNode(i)
	}
	list.PrintNodes()

	list.RemoveNode(list.root.next.next)
	list.PrintNodes()
}
