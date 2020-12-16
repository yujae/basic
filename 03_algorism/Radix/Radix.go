package Radix

import "fmt"

/*
	1. 가장 큰 수를 구함
	2. 가장 큰 수의 자릿수를 구함 (1000단위 = 4자리)
	3. 일의 자리부터 천의 자리까지
	   각 숫자를 순회하며 임의의 공간을 이용해 정렬
	   ex) 원래 배열의 전체 숫자를
          - 일의 자리에 있는 수를 비교해 임의의 공간에 추가 & 임의의 공간의 숫자들을 순서대로 원래 배열에 입력
	      - 십의 자리에 "
	      - 백의 자리에 "
	      - 천의 자리에 "
*/

func Radix(nums []int) {
	var bucket = [10][]int{}

	max := max(nums)
	maxDigit := maxDigit(max)

	for i := 1; i <= maxDigit; i++ {
		for j, v := range nums {
			bucket[digitValue(nums[j], i)] = append(bucket[digitValue(nums[j], i)], v)
		}

		nums = nil
		for j, v := range bucket {
			for _, w := range v {
				nums = append(nums, w)
			}
			bucket[j] = nil
		}
	}

	fmt.Println(nums)
}

func max(nums []int) int {
	max := 0
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}

func maxDigit(num int) int {
	maxDigit := 0
	for num != 0 {
		num /= 10
		maxDigit++
	}
	return maxDigit
}

func digitValue(num int, digit int) int {
	value := num % (tenJegop(digit)) / tenJegop(digit-1)
	return value
}

func tenJegop(count int) int {
	value := 10
	if count == 0 {
		value = 1
	}
	for i := 0; i < count-1; i++ {
		value = value * 10
	}
	return value
}
