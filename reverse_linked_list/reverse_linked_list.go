package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var res *ListNode
	for head.Next != nil {
		res = &ListNode{
			Val:  head.Val,
			Next: res,
		}
		head = head.Next
	}
	return &ListNode{
		Val:  head.Val,
		Next: res,
	}
}
