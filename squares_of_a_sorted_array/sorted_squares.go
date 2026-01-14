package squares_of_a_sorted_array

import "math"

func sortedSquares(nums []int) []int {
	ret := make([]int, len(nums))
	l := 0
	r := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if math.Abs(float64(nums[l])) <= math.Abs(float64(nums[r])) {
			ret[i] = nums[r] * nums[r]
			r--
		} else {
			ret[i] = nums[l] * nums[l]
			l++
		}
	}

	return ret
}
