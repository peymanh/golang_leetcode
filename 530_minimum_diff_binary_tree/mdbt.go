package minimumdiffbinarytree

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getMinimumDifference(root *TreeNode) int {
	values := []int{}

	getAllNodes(root, &values)

	slices.Sort(values)

	min := abs(values[0] - values[1])

	for i := 0; i < len(values)-1; i++ {
		if abs(values[i]-values[i+1]) < min {
			min = abs(values[i] - values[i+1])
		}
	}

	return min
}

func getAllNodes(root *TreeNode, values *[]int) {
	if root == nil {
		return
	}
	*values = append(*values, root.Val)

	getAllNodes(root.Left, values)
	getAllNodes(root.Right, values)
}
