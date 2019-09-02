package main

import (
	"fmt"
	"leetcode/symmetricTree/code"
)

func main() {
	fmt.Println(code.IsSymmetric(
		&code.TreeNode{
			Val: 1,
			Left: &code.TreeNode{
				Val: 2,
				Left:&code.TreeNode{Val:3},
				Right:&code.TreeNode{Val:4},
			},
			Right: &code.TreeNode{
				Val: 2,
				Left:&code.TreeNode{Val:4},
				Right:&code.TreeNode{Val:3},
			},
		},
	) == true)
	fmt.Println(code.IsSymmetric(
		&code.TreeNode{
			Val: 1,
			Left: &code.TreeNode{
				Val: 2,
				Right:&code.TreeNode{Val:3},
			},
			Right: &code.TreeNode{
				Val: 2,
				Right:&code.TreeNode{Val:3},
			},
		},
	) == false)
	fmt.Println(code.IsSymmetric(
		&code.TreeNode{
			Val: 1,
		},
	) == true)
	fmt.Println(code.IsSymmetric(
		&code.TreeNode{
			Val: 1,
			Left:&code.TreeNode{Val:2},
			Right:&code.TreeNode{Val:3},
		},
	) == false)
	fmt.Println(code.IsSymmetric(
		&code.TreeNode{
			Val: 2,
			Left: &code.TreeNode{
				Val: 3,
				Left:&code.TreeNode{Val:4},
				Right:&code.TreeNode{Val:5},
			},
			Right: &code.TreeNode{
				Val: 3,
				Left:&code.TreeNode{Val:4},
				Right:&code.TreeNode{Val:5},
			},
		},
	) == false)
}
