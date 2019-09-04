package main

import (
	"fmt"
	"leetcode/binaryTreeLevelOrderTraversalII/code"
)

func main() {
	fmt.Println(code.LevelOrderBottom(&code.TreeNode{
		Val:3,
		Left:&code.TreeNode{
			Val:9,
		},
		Right:&code.TreeNode{
			Val:20,
			Left:&code.TreeNode{
				Val:15,
			},
			Right:&code.TreeNode{
				Val:7,
			},
		},
	}))
	fmt.Println(code.LevelOrderBottom(&code.TreeNode{
		Val:3,
	}))
	fmt.Println(code.LevelOrderBottom(nil))
}