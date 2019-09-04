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

func sortedArrayToBST(nums []int) *TreeNode {
	//     二分法
	numsLength := len(nums)
	if numsLength == 0 {
		return nil
	}
	halfLength := numsLength / 2
	return &TreeNode{
		Val:nums[halfLength],
		Left:sortedArrayToBST(nums[:halfLength]),
		Right:sortedArrayToBST(nums[halfLength+1:]),
	}
}

func SortedArrayToBST(nums []int) *TreeNode {
	return sortedArrayToBST(nums)
}