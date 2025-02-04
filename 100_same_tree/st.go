package sametree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p != nil && q != nil {
		return (p.Val == q.Val) && IsSameTree(p.Right, q.Right) && IsSameTree(p.Left, q.Left)
	} else {
		return p == q
	}
}
