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

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Right == nil && root.Left == nil {
		return 1
	}
	var left = maxDepth(root.Left) + 1
	var right = maxDepth(root.Right) + 1
	if left > right {
		return left
	} else {
		return right
	}
}

func MaxDepth(root *TreeNode) int {
	return maxDepth(root)
}
