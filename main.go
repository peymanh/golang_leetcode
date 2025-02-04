package main

import (
	"fmt"

	linkedlistcycle "golang_leetcode/141_linked_list_cycle"
)

func main() {
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

	r := linkedlistcycle.HasCycle(&l1)
	fmt.Println(r)
}
