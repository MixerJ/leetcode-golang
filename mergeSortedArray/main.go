package main

import "leetcode/mergeSortedArray/code"

func main() {
	code.Merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	code.Merge([]int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3)
	code.Merge([]int{0, 0, 0}, 0, []int{1, 2, 3}, 3)
	code.Merge([]int{0, 0, 0}, 0, []int{}, 0)
	code.Merge([]int{0}, 0, []int{1}, 1)
	code.Merge([]int{4, 0, 0, 0, 0, 0}, 1, []int{1, 2, 3, 5, 6}, 5)
}
