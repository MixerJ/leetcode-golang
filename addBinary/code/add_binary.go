package code

func formatString(s string, zeroNum int) string {
	for i := 0; i < zeroNum; i++ {
		s = "0" + s
	}
	return s
}

func addBinary(a string, b string) string {
	aLength := len(a)
	bLength := len(b)
	var returnString string
	if aLength > bLength {
		b = formatString(b, aLength-bLength)
		bLength += aLength-bLength
	}
	if aLength < bLength {
		a = formatString(a, bLength-aLength)
		aLength += bLength-aLength
	}
	var plus int
	for i := aLength - 1; i >= 0; i-- {
		aString := string(a[i])
		bString := string(b[i])
		if aString == "1" && bString == "1" {
			if plus == 0 {
				plus = 1
				returnString = "0" + returnString
				continue
			} else {
				plus = 1
				returnString = "1" + returnString
				continue
			}
		} else if aString == "1" || bString == "1" {
			if plus == 1 {
				plus = 1
				returnString = "0" + returnString
				continue
			} else {
				returnString = "1" + returnString
			}
		} else {
			if plus == 1 {
				plus = 0
				returnString = "1" + returnString
				continue
			} else {
				returnString = "0" + returnString
			}
		}
	}
	if plus == 1 {
		returnString = "1" + returnString
	}
	return returnString
}

func addBinaryTwo(a,b string) string {
	if a == "" || b == "" {
		return a + b
	}
	if string(a[len(a)-1]) == "0" && string(b[len(b)-1]) == "0" {
		return addBinaryTwo(a[:len(a)-1], b[:len(b)-1]) + "0"
	} else if string(a[len(a)-1]) == "1" && string(b[len(b)-1]) == "1"{
		return addBinaryTwo(a[:len(a)-1], addBinaryTwo(b[:len(b)-1], "1")) + "0"
	} else {
		return addBinaryTwo(a[:len(a)-1], b[:len(b)-1]) + "1"
	}
}
func AddBinary(a string, b string) string {
	//return addBinary(a, b)
	return addBinaryTwo(a,b)
}
