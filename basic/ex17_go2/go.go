package main

import "fmt"

func main() {
	var a, b = 10, 5
	var result int

	c := make(chan int)

	go func() {
		println("inside go")
		c <- a + b
	}()

	println("inside main")
	// main과 go루틴과 별개로 돌지만 c로부터 수신되기전까지 대기함
	result = <-c
	fmt.Printf("두 수의 합은 %d입니다.", result)
}
