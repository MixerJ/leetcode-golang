package code

func longestPalindrome(s string) string {
	var longestString string
	var sLength = len(s)
	if sLength == 0 {
		return longestString
	}
	for k := range s {
		for i := sLength; i > 0; i-- {
			if k <= i && s[k] == s[i-1] {
				if isPalindrome(s[k:i]) {
					if i-k > len(longestString) {
						longestString = s[k:i]
					}
					if len(longestString) >= len(s[k:sLength]) {
						return longestString
					}
				} else {
					continue
				}
			}
		}
	}
	return longestString
}

func isPalindrome(s string) bool {
	var sLength = len(s)
	for k := range s {
		if k == sLength-k-1 {
			break
		}
		if s[k] != s[sLength-k-1] {
			return false
		}
	}
	return true
}

func longestPalindromeTwo(s string) string {
	var longestString string
	var sLength = len(s)
	if sLength == 0 {
		return longestString
	}
	for k := range s {
		var count int
		for i := sLength; i > 0; i-- {
			if k <= i {
				if k+count < sLength-1 && s[k+count]==s[i-1] {
					count++
				}
			} else {
				if len(s[k+count:i]) > len(longestString) {
					longestString = s[k+count:i]
				}
			}
		}
	}
	return longestString
}

func LongestPalindrome(s string) string {
	//return longestPalindrome(s)
	return longestPalindromeTwo(s)
}
