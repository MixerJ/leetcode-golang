package code

import (
	"leetcode/intReverse/code"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var rev = code.Reverse(x)
	if x == rev {
		return true
	}
	return false
}

func IsPalindrome(x int) bool {
	return isPalindrome(x)
}
