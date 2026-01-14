package palindrome_number

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		Input  int
		Output bool
	}{
		{
			12321,
			true,
		},
		{
			121,
			true,
		},
		{
			-121,
			false,
		},
		{
			10,
			false,
		},
	}

	for i, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, isPalindrome(tc.Input), tc.Output)
		})
	}
}
