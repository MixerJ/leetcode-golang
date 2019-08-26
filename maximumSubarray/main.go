package main

import (
	"fmt"
	"leetcode/maximumSubarray/code"
)

func main(){
	fmt.Println(code.MaxSubArray([]int{
		-2,1,-3,4,-1,2,1,-5,4,
	}) == 6)
	fmt.Println(code.MaxSubArray([]int{
		-1,
	}) == -1)
	fmt.Println(code.MaxSubArray([]int{
		-2, -1, -4,
	}))
	fmt.Println(code.MaxSubArray([]int{
		-3,-2,-2,-3,
	}))
}
