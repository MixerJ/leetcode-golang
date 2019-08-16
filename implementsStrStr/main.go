package main

import (
	"fmt"
	"leetcode/implementsStrStr/code"
)

func main() {
	fmt.Println(code.StrStr("hello", "ll") == 2)
	fmt.Println(code.StrStr("aaaaa", "bba") == -1)
	fmt.Println(code.StrStr("aaaaa", "") == 0)
	fmt.Println(code.StrStr("a", "a") == 0)
	fmt.Println(code.StrStr("", "a") == -1)
	fmt.Println(code.StrStr("", "") == 0)
	fmt.Println(code.StrStr("aaa", "aaaa") == -1)
	fmt.Println(code.StrStr("mississippi", "issip") == 4)
}
