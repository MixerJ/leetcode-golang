package main

import (
	"fmt"
	"leetcode/longestCommonPrefix/code"
)
func main() {
	fmt.Println(code.LongestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(code.LongestCommonPrefix([]string{"dog", "racecar", "car"}))
}