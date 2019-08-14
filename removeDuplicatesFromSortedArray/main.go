package main

import (
	"fmt"
	"leetcode/removeDuplicatesFromSortedArray/code"
)

func main() {
	fmt.Println(code.RemoveDuplicates([]int{1, 1, 2}))
	fmt.Println(code.RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
