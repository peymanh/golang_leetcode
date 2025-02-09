package binarytreeinordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTraversal(root *TreeNode) []int {
	nodes := []int{}
	if root == nil {
		return nodes
	}

	nodes = append(nodes, InorderTraversal(root.Left)...)
	nodes = append(nodes, root.Val)
	nodes = append(nodes, InorderTraversal(root.Right)...)
	return nodes

}
