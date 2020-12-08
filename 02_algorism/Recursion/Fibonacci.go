package Recursion

func Fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return Fibonacci(n-2) + Fibonacci(n-1)
}
