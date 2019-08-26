package code

func strStr(haystack string, needle string) int {
	var start = -1
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}

	for i := 0; i < len(haystack); i++ {
		start = i
		for key := range needle {
			if i+key >= len(haystack) {
				start = -1
				break
			}
			if needle[key] != haystack[i+key] {
				start = -1
				break
			}
			if len(needle)-1 == key {
				return start
			}
		}
	}
	return start
}
func StrStr(haystack string, needle string) int {
	return strStr(haystack, needle)
}
