package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1 * time.Second)
	after := time.After(5 * time.Second)

	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-after:
			fmt.Println("end")
			return
		}
	}
}
