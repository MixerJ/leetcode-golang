package main

import (
	"fmt"
	"leetcode/convertSortedArrayToBinarySearchTree/code"
)

func bfs(node *code.TreeNode, depth int, res *[][]int) {
	if node == nil {
		return
	}
	if depth > len(*res) {
		*res = append(*res, []int{})
	}
	(*res)[depth-1] = append((*res)[depth-1], node.Val)
	bfs(node.Left, depth+1, res)
	bfs(node.Right, depth+1, res)
}

func levelOrderBottomTwo(root *code.TreeNode) [][]int {
	var res [][]int
	bfs(root, 1, &res)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func main() {
	node := code.SortedArrayToBST([]int{-10,-3,0,5,9})
	fmt.Println(levelOrderBottomTwo(node))
}