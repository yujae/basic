package Heap

import (
	"errors"
)

/*
	1. 사용처
      - 최대/최소값 찾기, 정렬
      - Priorty Queue (우선 순위 높은 것이 먼저 나옴, ex. 응급실 대기열)
	2. 규칙
	  - 최대 힙 : 부모가 제일 크거나 같음
			10
		  /   \
		 9     8
		/ \   / \
	   5   6 7   8
	  - 최소 힙 : 부모가 제일 작거나 같음
			1
		  /   \
		 2     3
		/ \   / \
	   4   5 6   7

	  - Push (복잡도 : O(log2N))
	    : ex.최대 힙
		  마지막 자식에 값 추가
          부모로 올라가며 비교하여 자식이 클 경우, 스왑
	  - Pop (복잡도 : O(log2N))
	    : ex.최대 힙
		  최상단 팝, 마지막 자식을 루트로 올림
		  자식으로 내려가며 비교하여 자식이 클 경우, (자식들 중 가장 큰 것으로) 스왑
	  ★ 정렬 : 모든 수를 팝하게 되면 정렬되어 값이 나옴
        정렬 복잡도 = Push 10번, Pop 10번 => O(2Nlog2N) ==> N이 엄청 크면 2는 의미가 없어짐 따라서 O(Nlog2N)으로 표기
	3. 구현
	  - Tree로 구현하기 복잡 (Linked List로는 Tree의 맨 끝을 알기 어렵고 sibling도 알아야 함)
	  - Slice로 구현하기 적합
	  - N 번째 수의  Left 자식 : 2N+1
        N 번째 수의 Right 자식 : 2N+2
        N 번째 수의 부모 : (N-1)/2
        ex.
        s |9|7|8|3|4|6|7|
           0 1 2 3 4 5 6
        9의  Left 자식 : 2 * 0 + 1	=> s[1] = 7
	    9의 Right 자식 : 2 * 0 + 2	=> s[2] = 8
        6의 부모 : (5-1)/2			=> s[2] = 8
*/

type MaxHeap struct {
	list []int
}

type MinHeap struct {
	list []int
}

func (h *MaxHeap) Push(val int) {
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

func (h *MaxHeap) Pop() (int, error) {
	if len(h.list) == 0 {
		return 0, errors.New("no more")
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

		if leftChildIdx > maxIdx {
			break
		}
		// Left 자식만 있는 상태
		if maxIdx < rightChildIdx && h.list[idx] <= h.list[leftChildIdx] {
			h.list[idx], h.list[leftChildIdx] = h.list[leftChildIdx], h.list[idx]
			idx = leftChildIdx
			break
		}

		if rightChildIdx > maxIdx {
			break
		}
		if h.list[leftChildIdx] <= h.list[rightChildIdx] && h.list[idx] <= h.list[rightChildIdx] {
			h.list[idx], h.list[rightChildIdx] = h.list[rightChildIdx], h.list[idx]
			idx = rightChildIdx
		} else if h.list[idx] <= h.list[leftChildIdx] && h.list[rightChildIdx] <= h.list[leftChildIdx] {
			h.list[idx], h.list[leftChildIdx] = h.list[leftChildIdx], h.list[idx]
			idx = leftChildIdx
		} else {
			break
		}
	}

	return first, nil
}

func (h *MinHeap) Push(val int) {
	h.list = append(h.list, val)

	idx := len(h.list) - 1
	for idx > 0 {
		parentIdx := (idx - 1) / 2
		parent := h.list[parentIdx]

		if parent > val {
			h.list[parentIdx], h.list[idx] = h.list[idx], h.list[parentIdx]
		}

		idx = (idx - 1) / 2
	}
}

func (h *MinHeap) Pop() (int, error) {
	if len(h.list) == 0 {
		return 0, errors.New("no more")
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

		if leftChildIdx > maxIdx {
			break
		}
		// Left 자식만 있는 상태
		if maxIdx < rightChildIdx && h.list[idx] >= h.list[leftChildIdx] {
			h.list[idx], h.list[leftChildIdx] = h.list[leftChildIdx], h.list[idx]
			break
		}

		if rightChildIdx > maxIdx {
			break
		}
		// Left, Right 자식 둘다 있는 상태
		if h.list[leftChildIdx] >= h.list[rightChildIdx] && h.list[idx] >= h.list[rightChildIdx] {
			h.list[idx], h.list[rightChildIdx] = h.list[rightChildIdx], h.list[idx]
			idx = rightChildIdx
		} else if h.list[leftChildIdx] <= h.list[rightChildIdx] && h.list[leftChildIdx] <= h.list[idx] {
			h.list[idx], h.list[leftChildIdx] = h.list[leftChildIdx], h.list[idx]
			idx = leftChildIdx
		} else {
			break
		}
	}

	return first, nil
}
