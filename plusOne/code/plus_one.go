package code

func plusOne(digits []int) []int {
	var plus int
	var digitsLength = len(digits)
	var returnDigits []int
	digits[digitsLength-1] += 1
	if digits[digitsLength-1] < 10 {
		return digits
	} else {
		for i := digitsLength - 1; i >= 0; i-- {
			if digits[i] + plus > 9 {
				digitsNum := digits[i]
				digits[i] = (digitsNum + plus) % 10
				plus = (digitsNum + plus) / 10
			} else {
				if plus > 0 {
					digits[i] += plus
				}
				plus = 0
			}
		}
		if plus > 0 {
			returnDigits = append(returnDigits, plus)
		}
	}
	returnDigits = append(returnDigits, digits...)
	return returnDigits
}
func PlusOne(digits []int) []int {
	return plusOne(digits)
}
