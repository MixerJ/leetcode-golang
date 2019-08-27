package main

import (
	"fmt"
	"leetcode/plusOne/code"
)

func main() {
	fmt.Println(code.PlusOne([]int{1, 2, 3}))
	fmt.Println(code.PlusOne([]int{1, 2, 9}))
	fmt.Println(code.PlusOne([]int{9, 8, 9}))
	fmt.Println(code.PlusOne([]int{9}))
}
