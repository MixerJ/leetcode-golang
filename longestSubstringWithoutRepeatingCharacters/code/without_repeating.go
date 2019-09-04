package code

func lengthOfLongestSubstring(s string) int {
	var lens, head int
	for k := range s {
		t := indexs(s[head:k], s[k])
		if t == -1 {
			if lens < k-head+1 {
				lens = k - head + 1
			}
		} else {
			head = head + t + 1
		}
	}
	return lens
}

func indexs(s string, b byte) int {
	for k := range s {
		if s[k] == b {
			return k
		}
	}
	return -1
}

func LengthOfLongestSubstring(s string) int {
	return lengthOfLongestSubstring(s)
}
