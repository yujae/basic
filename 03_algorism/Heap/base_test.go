package Heap

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	s := &MaxHeap{}
	s.Push(3)
	s.Push(4)
	s.Push(8)
	s.Push(7)
	s.Push(9)
	s.Push(6)
	s.Push(7)

	fmt.Println(s)

	for i := 0; i < 100; i++ {
		s.Pop()
	}
}
