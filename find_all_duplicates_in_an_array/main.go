package main

import (
	"fmt"
	"math"
)

func findDuplicates2(nums []int) []int {
	m := make(map[int]int)
	var arr []int
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	for j := 0; j < len(nums); j++ {
		if m[nums[j]] == 2 {
			arr = append(arr, nums[j])
			delete(m, nums[j])
		}
	}
	return arr
}

func findDuplicates3(nums []int) []int {
	res := []int{}
	fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		fmt.Printf("i=%d\n", i)
		fmt.Printf("nums[i]=%d\n", nums[i])
		num := int(math.Abs(float64(nums[i])))
		fmt.Printf("%d=int(math.Abs(float64(nums[%d])))\n", num, i)
		fmt.Printf("nums[num-1]=%d\n", nums[num-1])
		if nums[num-1] < 0 {
			res = append(res, num)
		} else {
			nums[num-1] *= -1
		}
		fmt.Println(nums)
	}
	return res
}

func findDuplicates(nums []int) []int {
	var res []int

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	for i := 0; i < len(nums); i++ {
		num := abs(nums[i])
		if nums[num-1] < 0 {
			res = append(res, num)
		} else {
			nums[num-1] *= -1
		}
	}
	return res
}

func main() {
	fmt.Println(findDuplicates([]int{1, 1, 2, 3}))
	fmt.Println(findDuplicates([]int{4, 1, 4, 3, 3, 2}))
}
