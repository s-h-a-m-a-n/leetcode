package string_to_integer_atoi

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{"42", 42},
		{"   -042", -42},
		{"1337c0d3", 1337},
		{"0-1", 0},
		{"words and 987", 0},
		{"  -words and 987", 0},
		{"9223372036854775808", 2147483647},
	}

	for i, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, myAtoi(tc.input), tc.output)
		})
	}
}
