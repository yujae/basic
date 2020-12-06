package algorism

type Node struct {
	next  *Node
	value int
}

//func main() {
//
//}

func (root *Node) AddNode(value int) {
	tail := root
	for tail.next != nil {
		tail = tail.next
	}
	tail.next = &Node{value: value}
}
