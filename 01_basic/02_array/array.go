package main

import "fmt"

func main() {
	var strs [5]string
	strs[0] = "AB"
	strs[1] = "C"
	strs[2] = "D"

	fmt.Println(strs[0])
	fmt.Println(strs[1])
	fmt.Println(strs[2])
	fmt.Println(strs[3]) // 값을 대입하지 않은 것은 ""으로 초기화 됨

	var nums [4]int
	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	fmt.Println(nums) // 값을 대입하지 않은 것은 0으로 초기화 됨

	var arr = [3]int{3, 2, 1} // 리터럴
	fmt.Printf("배열값 : ", arr)
	fmt.Printf("배열길이 : ", len(arr))
}
