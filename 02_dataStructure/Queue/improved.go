package Queue

type BothSideNode struct {
	prev  *BothSideNode
	next  *BothSideNode
	value int
}

type DoubleLinkedList struct {
	root *BothSideNode
	tail *BothSideNode
}

func (list *DoubleLinkedList) Push(value int) {
	if list.root == nil {
		list.root = &BothSideNode{value: value}
		list.tail = list.root
		return
	}

	list.tail.next = &BothSideNode{value: value}
	list.tail.next.prev = list.tail
	list.tail = list.tail.next
}

func (list *DoubleLinkedList) Pop() int {
	first := list.root.value
	list.RemoveNode(list.root)
	return first
}

func (list *DoubleLinkedList) RemoveNode(node *BothSideNode) {
	if node == list.root {
		list.root = node.next
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
}
