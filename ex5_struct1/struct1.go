package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	v := Point{1, 2}
	v.X = 4
	fmt.Printf("%v\n", v)

	fmt.Println("-------------------------------------2")

	w := &v
	fmt.Println(w)
	fmt.Println(w.X)

	fmt.Println("-------------------------------------3")
	z := new(Point) // var z *Point = new(Point)
	x := z
	fmt.Println(z)
	x.X = 10
	z.X, z.Y = 11, 12
	fmt.Println(x)
	fmt.Println(z)
	fmt.Println(*x)
	fmt.Println(x == z)
}
