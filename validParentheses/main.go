package main

import (
	"fmt"
	"leetcode/validParentheses/code"
)

func main() {
	fmt.Println(code.IsValid("()"))
	fmt.Println(code.IsValid("()[]{}"))
	fmt.Println(code.IsValid("(]"))
	fmt.Println(code.IsValid("([)]"))
	fmt.Println(code.IsValid("{[]}"))
	fmt.Println(code.IsValid("{{"))
}
