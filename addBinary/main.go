package main

import (
	"fmt"
	"leetcode/addBinary/code"
)

func main() {
	fmt.Println(code.AddBinary("11", "1") == "100")
	fmt.Println(code.AddBinary("1", "11") == "100")
	fmt.Println(code.AddBinary("11", "11") == "110")
	fmt.Println(code.AddBinary("1010", "1011") == "10101")
}
