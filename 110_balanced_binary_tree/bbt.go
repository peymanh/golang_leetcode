package balancedbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return x * (-1)
	} else {
		return x
	}
}

func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return false
	}

	if abs(maxDepth(root.Right)-maxDepth(root.Left)) <= 1 {
		return IsBalanced(root.Right) && IsBalanced(root.Left)
	}
	return false
}
