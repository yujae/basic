package algorism

import (
	"fmt"
	"testing"
)

func TestNode_AddNode(t *testing.T) {
	root := &Node{value: 0}
	root.AddNode(1)
	root.AddNode(2)

	node := root
	for node.next != nil {
		fmt.Println(node.value)
		node = node.next
	}
	fmt.Println(node.value)
}
