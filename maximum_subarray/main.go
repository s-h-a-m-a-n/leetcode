package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	ret := nums[0]
	sum := nums[0]

	for _, n := range nums[1:] {
		sum += n
		ret = max(sum, ret)

		if sum < 0 {
			sum = 0
			continue
		}
	}

	return ret
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
	fmt.Println(maxSubArray([]int{1}))
	fmt.Println(maxSubArray([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(maxSubArray([]int{-12, -100, -2}))
}
