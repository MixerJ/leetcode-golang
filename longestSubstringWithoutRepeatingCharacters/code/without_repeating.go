package code

func lengthOfLongestSubstring(s string) int {
	var returnLength int
	var sLength = len(s)
	if sLength == 0 {
		return 0
	}
	var substringLength = 1
	for i := 0; i < sLength; i++ {
		var preString = string(s[i])
		for j := i + 1; j < sLength; j++ {
			jString := string(s[j])
			if preString == jString {
				returnLength = substringLength
				substringLength = 1
				if j == sLength-1 {
					sLength = j
					i = j - 1
				}
				break
			} else {
				substringLength++
			}
		}
	}
	if substringLength > returnLength {
		returnLength = substringLength
	}
	return returnLength
}

func lengthOfLongestSubstringTwo(s string) int {
	var returnLength int
	var sLength = len(s)
	var substringLength = 1
	for i := 0; i < sLength; i++ {
		var preString = string(s[i])
		for j := i + 1; j < sLength; j++ {

		}
	}
	return returnLength
}
func LengthOfLongestSubstring(s string) int {
	//return lengthOfLongestSubstring(s)
	return lengthOfLongestSubstringTwo(s)
}
