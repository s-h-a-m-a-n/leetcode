package main

import (
	"fmt"
)

func missingNumber(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return len(nums)*(len(nums)+1)/2 - sum
}
func main() {
	fmt.Println(missingNumber([]int{1, 3, 0}))
	fmt.Println(missingNumber([]int{1, 3, 0, 2, 4, 6, 7}))
}
