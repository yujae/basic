package main

import "fmt"

func main() {
	first := []int{1, 2}
	Print(first)

	second := append(first, 3) // first의 값이 새 공간에 복사되고 공간확보 필요시 2배씩 증가함
	Print(first)
	Print(second)
}

func Print(s []int) {
	fmt.Printf("%p : %d, %d\n", s, len(s), cap(s))
}
