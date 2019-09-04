package code

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	var res []float64
	if root == nil {
		return res
	}
	var nodes = []*TreeNode{root}
	for len(nodes) > 0 {
		var nextNodes []*TreeNode
		var sum int
		var count int
		for _, node := range nodes {
			if node.Left != nil {
				nextNodes = append(nextNodes, node.Left)
			}
			if node.Right != nil {
				nextNodes = append(nextNodes, node.Right)
			}
			sum += node.Val
			count++
		}
		nodes = nextNodes
		res = append(res, float64(sum) / float64(count))
	}
	return res
}

func AverageOfLevels(root *TreeNode) []float64 {
	return averageOfLevels(root)
}
