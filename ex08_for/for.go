package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 0
	for sum < 100 {
		sum += 1
	}
	fmt.Println(sum)

	var m = [3]int{9, 8, 7}
	for i, v := range m {
		fmt.Println(i, v)
	}
}
