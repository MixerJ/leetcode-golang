package main

import (
	"fmt"
	"leetcode/lengthOfLastWord/code"
)

func main() {
	fmt.Println(code.LengthOfLastWord("Hello World"))
	fmt.Println(code.LengthOfLastWord("Hello World "))
	fmt.Println(code.LengthOfLastWord("Hello World  "))
	fmt.Println(code.LengthOfLastWord("  "))
	fmt.Println(code.LengthOfLastWord("Hello  World"))
}
