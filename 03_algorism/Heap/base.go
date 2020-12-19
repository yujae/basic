package Heap

import "fmt"

type Heap struct {
	list []int
}

func (h *Heap) Push(val int) {
	h.list = append(h.list, val)

	idx := len(h.list) - 1
	for idx > 0 {
		parentIdx := (idx - 1) / 2
		parent := h.list[parentIdx]

		if parent < val {
			h.list[parentIdx], h.list[idx] = h.list[idx], h.list[parentIdx]
		}

		idx = (idx - 1) / 2
	}
}

func (h *Heap) Pop() {
	if len(h.list) == 0 {
		return
	}

	idx := 0
	first := h.list[idx]
	last := h.list[len(h.list)-1]

	h.list[0] = last
	h.list = h.list[:len(h.list)-1]
	maxIdx := len(h.list) - 1

	for idx < maxIdx {
		leftChildIdx := 2*idx + 1
		rightChildIdx := 2*idx + 2

		if leftChildIdx > maxIdx || rightChildIdx > maxIdx {
			break
		}

		if maxIdx < rightChildIdx && h.list[0] < h.list[leftChildIdx] {
			h.list[idx], h.list[leftChildIdx] = h.list[leftChildIdx], h.list[idx]
			idx = leftChildIdx
			break
		}
		if h.list[leftChildIdx] < h.list[rightChildIdx] && h.list[idx] < h.list[rightChildIdx] {
			h.list[idx], h.list[rightChildIdx] = h.list[rightChildIdx], h.list[idx]
			idx = rightChildIdx
		} else if h.list[0] < h.list[leftChildIdx] {
			h.list[idx], h.list[leftChildIdx] = h.list[leftChildIdx], h.list[idx]
			idx = leftChildIdx
		}

		maxIdx = len(h.list) - 1
	}

	fmt.Printf("%d ", first)
}
