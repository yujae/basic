package Stack

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	stack := &DoubleLinkedList{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
