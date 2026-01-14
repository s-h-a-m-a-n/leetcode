package reverse_linked_list_ii

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left, right int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	if left == right {
		return head
	}

	var res, reverse, after, before *ListNode
	i := 1
	for head != nil {
		switch {
		case i < left:
			if before == nil {
				before = &ListNode{Val: head.Val}
				res = before
			} else {
				before.Next = &ListNode{Val: head.Val}
				before = before.Next
			}
		case i > right:
			if reverse != nil {
				if res == nil {
					res = reverse
				} else {
					before.Next = reverse
				}
				reverse = nil
			}

			after.Next = &ListNode{Val: head.Val}
			after = after.Next
		default:
			reverse = &ListNode{
				Val:  head.Val,
				Next: reverse,
			}
			if i == left {
				after = reverse
			}
		}

		i++
		head = head.Next
	}
	if reverse != nil {
		if res == nil {
			res = reverse
		} else {
			before.Next = reverse
		}
		reverse = nil
	}

	return res
}
