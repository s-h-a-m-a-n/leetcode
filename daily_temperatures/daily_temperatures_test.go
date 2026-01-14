package daily_temperatures

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		Input  []int
		Output []int
	}{
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		//{[]int{99, 99, 99, 99, 99, 100, 76, 73}, 4},
		//{[]int{80, 80, 80, 34, 80, 80, 34, 80, 80, 80}, 0},
		{[]int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{[]int{30, 60, 90}, []int{1, 1, 0}},
	}

	for i, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, dailyTemperatures(tc.Input), tc.Output)
		})
	}
}
