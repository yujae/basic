package main

import (
	"fmt"
	"math"
)

func main() {
	s := func(x float64, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(s(3, 4))
}
