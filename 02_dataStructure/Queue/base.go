package Queue

/* 	base : Slice 로 구현
   	improved : Double Linked List 로 구현

	Slice : 데이터 추가 시, cap 증가가 필요할 때 O(N)이고 그 외엔 O(1)
			데이터 삭제 시, O(1)
	List : 데이터 추가/삭제 시, O(1)

	Cache 사용 관점에서는 Slice 가 유리할 있음
	Cache 상관없다면 List 가 유리
*/

func Push(s *[]int, val int) {
	*s = append(*s, val)
}

func Pop(s *[]int) int {
	var first int
	first, *s = (*s)[0], (*s)[1:]
	return first
}
