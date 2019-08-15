package main

import (
	"fmt"
	"leetcode/removeElement/code"
)

func main() {
	fmt.Println(code.RemoveElement([]int{3, 2, 2, 3}, 3))
	fmt.Println(code.RemoveElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
