package best_time_to_buy_and_sell_stock_ii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		Input  []int
		Output int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 7},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{7, 6, 4, 5, 1}, 1},
		{[]int{1, 9, 6, 9, 1, 7, 1, 1, 5, 9, 9, 9}, 25},
	}

	for i, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, maxProfit(tc.Input), tc.Output)
		})
	}
}
