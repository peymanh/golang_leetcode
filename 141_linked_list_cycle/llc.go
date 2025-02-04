package linkedlistcycle

import (
	"fmt"
	"slices"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func HasCycle(head *ListNode) bool {
	seen := []*ListNode{}

	seen = append(seen, head)
	for head != nil && head.Next != nil {
		fmt.Println(seen)
		if slices.Contains(seen, head.Next) {
			return true
		}
		seen = append(seen, head)
		head = head.Next
	}
	return false
}
