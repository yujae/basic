package Stack

import (
	"fmt"
	"testing"
)

func TestBaseStack(t *testing.T) {
	stack := &[]int{}

	Push(stack, 1)
	Push(stack, 2)
	Push(stack, 3)
	Push(stack, 4)
	fmt.Println(Pop(stack))
	fmt.Println(Pop(stack))
	fmt.Println(Pop(stack))
	fmt.Println(Pop(stack))
}
