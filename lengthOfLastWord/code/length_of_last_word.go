package code

func lengthOfLastWord(s string) int {
	var wordLength int
	var lastWordLength int
	for i := 0; i < len(s); i++ {
		sIString := string(s[i])
		if sIString != " " {
			wordLength++
		}
		if sIString == " " || i == len(s)-1 {
			if wordLength != 0 {
				lastWordLength = wordLength
			}
			wordLength = 0
		}
	}
	return lastWordLength
}

func LengthOfLastWord(s string) int {
	return lengthOfLastWord(s)
}
