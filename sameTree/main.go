package main

import (
	"fmt"
	"leetcode/sameTree/code"
)

func main() {
	fmt.Println(code.IsSameTree(
		&code.TreeNode{Val: 1, Left: &code.TreeNode{Val: 2}, Right: &code.TreeNode{Val: 3}},
		&code.TreeNode{Val: 1, Left: &code.TreeNode{Val: 2}, Right: &code.TreeNode{Val: 3}},
	) == true)
	fmt.Println(code.IsSameTree(
		&code.TreeNode{Val: 0},
		&code.TreeNode{Val: 1},
	) == false)
	fmt.Println(code.IsSameTree(
		&code.TreeNode{Val: 1, Left: &code.TreeNode{Val: 2}},
		&code.TreeNode{Val: 1, Right: &code.TreeNode{Val: 2}},
	) == false)
	fmt.Println(code.IsSameTree(
		&code.TreeNode{Val: 1, Left: &code.TreeNode{Val: 2}, Right: &code.TreeNode{Val: 1}},
		&code.TreeNode{Val: 1, Left: &code.TreeNode{Val: 1}, Right: &code.TreeNode{Val: 2}},
	) == false)
	fmt.Println(code.IsSameTree(
		&code.TreeNode{Val: 1, Left: &code.TreeNode{Val: 2}, Right: &code.TreeNode{Val: 1}},
		&code.TreeNode{Val: 1, Left: &code.TreeNode{Val: 1}, Right: &code.TreeNode{Val: 2}},
	) == false)
	fmt.Println(code.IsSameTree(
		&code.TreeNode{Val: 10, Left: &code.TreeNode{Val: 5}, Right: &code.TreeNode{Val: 15}},
		&code.TreeNode{Val: 10, Left: &code.TreeNode{Val: 5, Right: &code.TreeNode{Val: 15}}},
	) == false)
}
