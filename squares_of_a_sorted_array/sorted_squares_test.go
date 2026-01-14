package squares_of_a_sorted_array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedSquares(t *testing.T) {
	tests := []struct {
		nums []int
		ret  []int
	}{
		{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
	}

	for i, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, sortedSquares(tc.nums), tc.ret)
		})
	}
}
