package Queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := &DoubleLinkedList{}
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)

	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
}
