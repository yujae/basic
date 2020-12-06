package algorism

import "fmt"

func Channel(n int, sum int, result chan int) {
	sum = sum + n
	if n == 1 {
		result <- sum
		return
	}

	go Channel(n-1, sum, result)
}

func Start() {
	result := make(chan int)
	Channel(4, 0, result)
	answer := <-result
	fmt.Println(answer)
}
