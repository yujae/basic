package main

import "fmt"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	go MakeTire(ch1, ch2)
	go MakeEngine(ch2, ch3)

	ch1 <- "Start Work"
	fmt.Println(<-ch3)
}

func MakeTire(ch1 chan string, ch2 chan string) {
	v := <-ch1
	v = v + "-> MakeTire"
	ch2 <- v
}

func MakeEngine(ch2 chan string, ch3 chan string) {
	v := <-ch2
	v = v + "-> MakeEngine"
	ch3 <- v
}
