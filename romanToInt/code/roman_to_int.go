package code

func romanToInt(s string) int {
	var intNum int
	var romanToIntMap = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	sLength := len(s)
	var preRoman string
	for i := 0; i < sLength; i++ {
		roman := string(s[i])
		intNum += romanToIntMap[roman]
		if roman == "X" || roman == "V" {
			if preRoman == "I" {
				intNum -= 2
			}
		} else if roman == "L" || roman == "C" {
			if preRoman == "X" {
				intNum -= 20
			}
		} else if roman == "D" || roman == "M" {
			if preRoman == "C" {
				intNum -= 200
			}
		}
		preRoman = roman
	}
	return intNum
}

func RomanToInt(s string) int {
	return romanToInt(s)
}
