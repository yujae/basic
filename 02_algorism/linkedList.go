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

func RemoveNode(prev *Node) {
	prev.next = prev.next.next
}

func PrintNodes(root *Node) {
	node := root
	for node.next != nil {
		fmt.Printf("%d -> ", node.value)
		node = node.next
	}
	fmt.Printf("%d\n", node.value)
}
