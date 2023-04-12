package twoSum

func twoSum(nums []int, target int) []int {
	//runtime
	l := len(nums)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if target == nums[i]+nums[j] {
				return []int{i, j}
			}
		}
	}
	//memory
	//m := make(map[int]int)
	//for i, val := range nums {
	//	if j, ok := m[target-val]; ok {
	//		return []int{j, i}
	//	}
	//	m[val] = i
	//}

	return []int{}
}
