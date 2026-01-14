package generate_parentheses

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		Input  int
		Output []string
	}{
		//{1, []string{"()"}},
		//{2, []string{
		//	"(())",
		//	"()()"},
		//},
		{4, []string{
			"((()))", "()(())", "(()())", "(())()", "()()()"},
		},
	}

	for i, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			t.Parallel()
			assert.ElementsMatch(t, generateParenthesis(tc.Input), tc.Output)
		})
	}
}
