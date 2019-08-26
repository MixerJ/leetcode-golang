package main

import (
	"fmt"
	"leetcode/countAndSay/code"
)

func main() {
	fmt.Println(code.CountAndSay(1) == "1")
	fmt.Println(code.CountAndSay(2) == "11")
	fmt.Println(code.CountAndSay(3) == "21")
	fmt.Println(code.CountAndSay(4) == "1211")
	fmt.Println(code.CountAndSay(5) == "111221")
	fmt.Println(code.CountAndSay(6) == "312211")
	fmt.Println(code.CountAndSay(7) == "13112221")
}
