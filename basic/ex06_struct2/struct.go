package main

import "fmt"

type Point2 struct {
	x, y int
}

func (p Point2) init() Point2 {
	p.x = 1
	p.y = 2
	return p
}

func (p *Point2) initWithRef() {
	p.x = 3
	p.y = 4
}

func main() {
	var p *Point2 = &Point2{4, 5}
	fmt.Println(p)

	w := p.init()
	fmt.Println(w)
	p.initWithRef()
	fmt.Println(p)
}
