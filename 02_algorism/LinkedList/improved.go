package LinkedList

import "fmt"

type LinkedList struct {
	root *Node
	tail *Node
}

func (l *LinkedList) AddNode(value int) {
	if l.root == nil {
		l.root = &Node{value: value}
		l.tail = l.root
		return
	} else {

	}

	l.tail.next = &Node{value: value}
	l.tail = l.tail.next
}

func (l *LinkedList) RemoveNode(node *Node) {
	if node == l.root {
		l.root = node.next
		node.next = nil
		return
	}

	prev := l.root
	for prev.next != node {
		prev = prev.next
	}

	if node == l.tail {
		prev.next = nil
		l.tail = prev
	} else {
		prev.next = prev.next.next
	}
	node.next = nil
}

func (l *LinkedList) PrintNodes() {
	node := l.root
	for node.next != nil {
		fmt.Printf("%d -> ", node.value)
		node = node.next
	}
	fmt.Printf("%d\n", node.value)
}
