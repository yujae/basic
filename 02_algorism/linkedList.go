package algorism

import "fmt"

type Node struct {
	next  *Node
	value int
}

func NewRoot() *Node {
	root := &Node{value: 0}
	return root
}

func FindTail(root *Node) *Node {
	tail := root
	for tail.next != nil {
		tail = tail.next
	}
	return tail
}

func AddNode(tail *Node, value int) *Node {
	tail.next = &Node{value: value}
	return tail.next
}

func RemoveNode(node *Node, root *Node, tail *Node) (*Node, *Node) {
	if node == root {
		root = node.next
		node.next = nil
		return root, tail
	}

	prev := node
	for prev != nil {
		prev = prev.next
	}

	if node == tail {
		prev.next = nil
		tail = prev
	} else {
		node.next = node.next.next
	}
	return root, tail
}

func PrintNodes(root *Node) {
	node := root
	for node.next != nil {
		fmt.Printf("%d -> ", node.value)
		node = node.next
	}
	fmt.Printf("%d\n", node.value)
}
