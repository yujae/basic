package Radix

import "fmt"

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
