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

//需要记录层次
//或者记录堆栈信息
func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var nodes = []*TreeNode{root}
	for len(nodes) > 0 {
		var nextNodes []*TreeNode
		var tmpRes = [][]int{{},}
		for _, node := range nodes {
			if node.Left != nil {
				nextNodes = append(nextNodes, node.Left)
			}
			if node.Right != nil {
				nextNodes = append(nextNodes, node.Right)
			}
			tmpRes[0] = append(tmpRes[0], node.Val)
		}
		nodes = nextNodes
		res = append(tmpRes, res...)
	}
	return res
}

func bfs(node *TreeNode, depth int, res *[][]int) {
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

func levelOrderBottomTwo(root *TreeNode) [][]int {
	var res [][]int
	bfs(root, 1, &res)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func LevelOrderBottom(root *TreeNode) [][]int {
	//return levelOrderBottom(root)
	return levelOrderBottomTwo(root)
}
