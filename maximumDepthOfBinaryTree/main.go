package main

import (
	"fmt"
	"leetcode/maximumDepthOfBinaryTree/code"
)

func main() {
	fmt.Println(code.MaxDepth(&code.TreeNode{
		Val:  3,
		Left: &code.TreeNode{Val: 9},
		Right: &code.TreeNode{
			Val:   20,
			Left:  &code.TreeNode{Val: 15},
			Right: &code.TreeNode{Val: 7},
		},
	}))
}
