package code

import "fmt"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	s := countAndSay(n-1) + "*"
	var res, count = "", 1
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			count += 1
		} else {
			res += fmt.Sprintf("%d", count) + string(s[i])
			count = 1
		}
	}
	return res
}

func CountAndSay(n int) string {
	return countAndSay(n)
}
