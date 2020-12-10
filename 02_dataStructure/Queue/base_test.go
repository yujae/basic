package Queue

import (
	"fmt"
	"testing"
)

func TestBaseQueue(t *testing.T) {
	queue := &[]int{}

	Push(queue, 1)
	Push(queue, 2)
	Push(queue, 3)
	Push(queue, 4)
	fmt.Println(Pop(queue))
	fmt.Println(Pop(queue))
	fmt.Println(Pop(queue))
	fmt.Println(Pop(queue))
}
