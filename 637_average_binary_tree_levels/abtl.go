package averagebinarytreelevels

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}

	queue := []*TreeNode{}
	result := []float64{}

	queue = append(queue, root)

	for len(queue) > 0 {
		n := len(queue)
		sum := 0

		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			sum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		avg := float64(sum) / float64(n)
		result = append(result, avg)
	}
	return result

}
