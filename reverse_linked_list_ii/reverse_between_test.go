package reverse_linked_list_ii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseBetween(t *testing.T) {
	tests := []struct {
		head   *ListNode
		left   int
		right  int
		result *ListNode
	}{
		{
			head: &ListNode{
				1,
				&ListNode{
					2,
					&ListNode{
						3,
						&ListNode{
							4,
							&ListNode{
								5,
								nil,
							},
						},
					},
				},
			},
			left:  1,
			right: 2,
			result: &ListNode{
				1,
				&ListNode{
					2,
					&ListNode{
						3,
						&ListNode{
							4,
							&ListNode{
								5,
								nil,
							},
						},
					},
				},
			},
		},
		{
			head: &ListNode{
				1,
				&ListNode{
					2,
					&ListNode{
						3,
						nil,
					},
				},
			},
			left:  2,
			right: 3,
			result: &ListNode{
				1,
				&ListNode{
					3,
					&ListNode{
						2,
						nil,
					},
				},
			},
		},

		{
			head: &ListNode{
				3,
				&ListNode{
					5,
					nil,
				},
			},
			left:  1,
			right: 2,
			result: &ListNode{
				5,
				&ListNode{
					3,
					nil,
				},
			},
		},
	}

	for i, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, reverseBetween(tc.head, tc.left, tc.right), tc.result)
		})
	}
}
