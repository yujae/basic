package Recursion

import "fmt"

func Print_one(n int) {
	if n >= 0 {
		fmt.Println("In first function : ", n)
		Print_two(n - 1)
	}
	fmt.Println("out..In first function : ", n)
}

func Print_two(n int) {
	if n >= 0 {
		fmt.Println("In second function : ", n)
		Print_one(n - 1)
	}
	fmt.Println("out..In second function : ", n)
}
