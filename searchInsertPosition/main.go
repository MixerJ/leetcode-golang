package main

import (
	"fmt"
	"leetcode/searchInsertPosition/code"
)

func main() {
	fmt.Println(code.SearchInsert([]int{
		1, 3, 5, 6,
	}, 5) == 2)
	fmt.Println(code.SearchInsert([]int{
		1, 3, 5, 6,
	}, 2) == 1)
	fmt.Println(code.SearchInsert([]int{
		1, 3, 5, 6,
	}, 7) == 4)
	fmt.Println(code.SearchInsert([]int{
		1, 3, 5, 6,
	}, 0) == 0)
}
