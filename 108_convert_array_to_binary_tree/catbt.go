package convertarraytobinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	med := len(nums) / 2

	root := TreeNode{
		Val: nums[med],
	}
	root.Left = sortedArrayToBST(nums[:med])
	root.Right = sortedArrayToBST(nums[med+1:])
	return &root
}
