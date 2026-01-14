package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2[:n]...)
	sort.Ints(nums1)
}

func main() {
	var nums1 []int
	nums1 = []int{4, 5, 6, 0, 0, 0}
	merge(nums1, 3, []int{1, 2, 3}, 3)
	fmt.Println(nums1)
	nums1 = []int{1, 2, 3, 0, 0, 0}
	merge(nums1, 3, []int{2, 5, 6}, 3)
	fmt.Println(nums1)
	nums1 = []int{1}
	merge(nums1, 1, []int{}, 0)
	fmt.Println(nums1)
	nums1 = []int{}
	merge(nums1, 0, []int{1}, 1)
	fmt.Println(nums1)

}
