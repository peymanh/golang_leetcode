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

/* If you wanna test
l1 := linkedlistcycle.ListNode{
	Val: 1,
}

l2 := linkedlistcycle.ListNode{
	Val: 2,
}

l3 := linkedlistcycle.ListNode{
	Val: 3,
}

l4 := linkedlistcycle.ListNode{
	Val: 4,
}

l1.Next = &l2
l2.Next = &l3
l3.Next = &l4
l4.Next = &l2

*/
