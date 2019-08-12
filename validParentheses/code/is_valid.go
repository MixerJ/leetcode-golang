package code

func isValid(s string) bool {
	sLength := len(s)
	if sLength == 0 {
		return true
	}
	if sLength % 2 != 0 {
		return false
	}
	var stack []string
	for i:=0; i < sLength; i++ {
		iString := string(s[i])
		stackLength := len(stack)
		if stackLength == 0 {
			stack = append(stack, iString)
			continue
		}

		if iString == "]" {
			if stack[stackLength-1] != "[" {
				return false
			} else {
				if stackLength == 1 {
					stack = []string{}
				} else {
					stack = stack[0:stackLength-1]
				}
			}
		} else if iString == "}" {
			if stack[stackLength-1] != "{" {
				return false
			} else {
				if len(stack) == 1 {
					stack = []string{}
				} else {
					stack = stack[0:stackLength-1]
				}
			}
		} else if iString == ")" {
			if stack[stackLength-1] != "(" {
				return false
			} else {
				if stackLength == 1 {
					stack = []string{}
				} else {
					stack = stack[0:stackLength-1]
				}
			}
		} else {
			stack = append(stack, iString)
		}
	}

	if len(stack) != 0 {
		return false
	}
	return true
}
func IsValid(s string) bool {
	return isValid(s)
}