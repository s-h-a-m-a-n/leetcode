package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var s []int
	for _, list := range lists {
		for Node := list; Node != nil; Node = Node.Next {
			s = append(s, Node.Val)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(s)))

	var l *ListNode

	for _, v := range s {
		l = &ListNode{
			Val:  v,
			Next: l,
		}
	}

	return l
}

func main() {
	lists := []*ListNode{}

	lists = append(lists, &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}})
	lists = append(lists, &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}})
	lists = append(lists, &ListNode{Val: 2, Next: &ListNode{Val: 6}})

	fmt.Print(mergeKLists(lists))
}
