package main

import "fmt"

func main() {
	var a, b = 10, 5
	var result int

	go func() {
		result = a + b
	}()

	// main과 go루틴과 별개로 돌아서 result가 계산되기 전에 끝날수 있음
	fmt.Printf("두 수의 합은 %d입니다.", result)
}
