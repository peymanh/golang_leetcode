package invertbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Right == nil && root.Left == nil {
		return root
	}

	root.Right, root.Left = InvertTree(root.Left), InvertTree(root.Right)

	return root
}
