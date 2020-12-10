package DoubleLinkedList

import "fmt"

type Node struct {
	prev  *Node
	next  *Node
	value int
}

type DoubleLinkedList struct {
	root *Node
	tail *Node
}

func (list *DoubleLinkedList) AddNode(value int) {
	if list.root == nil {
		list.root = &Node{value: value}
		list.tail = list.root
		return
	}

	list.tail.next = &Node{value: value}
	list.tail.next.prev = list.tail
	list.tail = list.tail.next
}

func (list *DoubleLinkedList) RemoveNode(node *Node) {
	if node == list.root {
		list.root = node.next
		node.next.prev = nil
		node.next = nil
		return
	}

	if node == list.tail {
		node.prev.next = nil
		list.tail = node.prev
		node.next = nil
	}

	node.prev.next = node.next
	node.next.prev = node.prev
	node = nil
	//node = nil
	//node.prev = nil
	//node.next = nil
}

func (list *DoubleLinkedList) PrintNodes() {
	node := list.root
	for node.next != nil {
		fmt.Printf("%d -> ", node.value)
		node = node.next
	}
	fmt.Printf("%d\n", node.value)

	node = list.tail
	for node.prev != nil {
		fmt.Printf("%d -> ", node.value)
		node = node.prev
	}
	fmt.Printf("%d\n", node.value)
}
