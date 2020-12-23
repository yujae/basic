package Heap

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	s := &MaxHeap{}
	for i := 0; i < 100; i++ {
		s.Push(rand.Intn(100))
	}

	for i := 0; i < 100; i++ {
		val, err := s.Pop()
		if err != nil {
			break
		}
		fmt.Printf("%d ", val)
	}
}

func TestMinHeap(t *testing.T) {
	s := &MinHeap{}
	for i := 0; i < 100; i++ {
		s.Push(rand.Intn(100))
	}

	for i := 0; i < 100; i++ {
		val, err := s.Pop()
		if err != nil {
			break
		}
		fmt.Printf("%d ", val)
	}
}

func TestNth(t *testing.T) {
	nums := []int{-3, -1, 0, 44, 3, 2, 4, 5, 10}
	Nth := 5
	s := &MinHeap{}

	for _, v := range nums {
		s.Push(v)
		if len(s.list) > Nth {
			_, err := s.Pop()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	fmt.Println(s.list[0])
}

func TestNth2(t *testing.T) {
	var nums []int
	Nth := 5
	s := &MinHeap{}

	for i := 1; i <= 100000000; i++ {
		nums = append(nums, i)
	}

	for _, v := range nums {
		s.Push(v)
		if len(s.list) > Nth {
			_, err := s.Pop()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	fmt.Println(s.list[0])
}

func TestNth3(t *testing.T) {
	var nums []int
	Nth := 5
	s := &MaxHeap{}

	for i := 1; i <= 100000000; i++ {
		nums = append(nums, i)
	}

	for _, v := range nums {
		s.Push(v)
	}

	for i := 0; i < Nth; i++ {
		val, _ := s.Pop()
		fmt.Println(val)
	}
}
