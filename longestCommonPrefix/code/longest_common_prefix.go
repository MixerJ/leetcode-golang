package code

func longestCommonPrefix(strs []string) string {
	var commonPrefix string
	if len(strs) == 0 {
		return commonPrefix
	}
	oneString := strs[0]
	for i := 1; i < len(strs); i++ {
		var oneLength = len(oneString)
		if oneLength > len(strs[i]) {
			oneLength = len(strs[i])
		}
		var prefix string
		for j := 0; j < oneLength; j++ {
			if string(oneString[j]) != string(strs[i][j]) {
				if i == 1 && j == 0 {
					return ""
				}
				break
			}
			prefix += string(oneString[j])
		}
		oneString = prefix
	}
	commonPrefix = oneString
	return commonPrefix
}

func LongestCommonPrefix(strs []string) string {
	return longestCommonPrefix(strs)
}

