package main

import (
	"fmt"
	"leetcode/climbingStairs/code"
)

func main() {
	fmt.Println(code.ClimbStairs(0) == 1)
	fmt.Println(code.ClimbStairs(1) == 1)
	fmt.Println(code.ClimbStairs(2) == 2)
	fmt.Println(code.ClimbStairs(3) == 3)
	fmt.Println(code.ClimbStairs(4) == 5)
	fmt.Println(code.ClimbStairs(5) == 8)
	fmt.Println(code.ClimbStairs(6) == 13)
}
