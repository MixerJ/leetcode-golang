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

func isSymmetricCycle(root *TreeNode) bool {
	if root == nil {
		return true
	}
	for root.Left != nil && root.Right != nil {
		//if
	}
	return false
}

func symmetricRecursive(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil && right != nil || left != nil && right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return symmetricRecursive(left.Left, right.Right) && symmetricRecursive(left.Right, right.Left)
}

func isSymmetricRecursive(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Right == nil && root.Left == nil {
		return true
	}
	if root.Left == nil && root.Right != nil || root.Left != nil && root.Right == nil {
		return false
	}
	if root.Right.Val != root.Left.Val {
		return false
	}
	return symmetricRecursive(root.Left.Left, root.Right.Right) && symmetricRecursive(root.Left.Right, root.Right.Left)
}

func IsSymmetric(root *TreeNode) bool {
	//return isSymmetricCycle(root)
	return isSymmetricRecursive(root)
}
