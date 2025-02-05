package main

import (
	"fmt"

	balancedbinarytree "golang_leetcode/110_balanced_binary_tree"
)

func main() {
	r1 := balancedbinarytree.TreeNode{
		Val: 1,
	}
	r2r := balancedbinarytree.TreeNode{
		Val: 2,
	}
	// r2l := balancedbinarytree.TreeNode{
	// 	Val: 2,
	// }
	// r3r := balancedbinarytree.TreeNode{
	// 	Val: 3,
	// }
	// r3l := balancedbinarytree.TreeNode{
	// 	Val: 3,
	// }
	// r4r := balancedbinarytree.TreeNode{
	// 	Val: 4,
	// }
	// r4l := balancedbinarytree.TreeNode{
	// 	Val: 4,
	// }

	r1.Right = &r2r
	// r1.Left = &r2l

	// r2l.Left = &r3l
	// r2l.Right = &r3r

	// r3l.Left = &r4l
	// r3l.Right = &r4r

	r := balancedbinarytree.IsBalanced(&r1)
	fmt.Println(r)
}
