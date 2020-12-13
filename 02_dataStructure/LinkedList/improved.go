package LinkedList

import "fmt"

type OneSideNode struct {
	next  *OneSideNode
	value int
}

type LinkedList struct {
	root *OneSideNode
	tail *OneSideNode
}

func (list *LinkedList) AddNode(value int) {
	if list.root == nil {
		list.root = &OneSideNode{value: value}
		list.tail = list.root
		return
	}

	list.tail.next = &OneSideNode{value: value}
	list.tail = list.tail.next
}

func (list *LinkedList) RemoveNode(node *OneSideNode) {
	if node == list.root {
		list.root = node.next
		node.next = nil
		return
	}

	prev := list.root
	for prev.next != node {
		prev = prev.next
	}

	if node == list.tail {
		prev.next = nil
		list.tail = prev
	} else {
		prev.next = prev.next.next
	}
	node.next = nil
}

func (list *LinkedList) PrintNodes() {
	node := list.root
	for node.next != nil {
		fmt.Printf("%d -> ", node.value)
		node = node.next
	}
	fmt.Printf("%d\n", node.value)
}
