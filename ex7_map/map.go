package main

import "fmt"

type Vertex struct {
	lat  float32
	long float32
}

func main() {
	var m1 map[string]Vertex
	m1 = make(map[string]Vertex)
	m1["test"] = Vertex{123, 456}
	fmt.Printf("%v\n", m1)

	var m2 = map[string]Vertex{
		"literal": Vertex{111, 222},
		"test2":   Vertex{333, 444},
	}
	fmt.Printf("%v\n", m2)

	var m3 = map[string]Vertex{
		"literal": {999, 888},
		"test3":   {444, 666},
	}
	fmt.Printf("%v\n\n", m3)

	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Printf("%v\n", m)
	v, present := m["Answer"]
	fmt.Printf("%d %t", v, present)
	delete(m, "Answer")
}
