package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
		// LIFO, 여러 파일 클로즈시 사용하면 좋음
	}
}
