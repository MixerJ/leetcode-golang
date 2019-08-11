package code

import "math"

var multiple = 10

func reverse(x int) int {
	//递归
	// residue := x % 10
	// quotient := x / 10
	// if quotient == 0 {
	//     multiple = 10
	//     return x
	// }
	// returnData := residue * multiple + reverse(quotient)
	// multiple = multiple * 10
	// if returnData < math.MinInt32 || returnData > math.MaxInt32 {
	//     return 0
	// }
	// return returnData

	// 遍历
	var rev int
	for x != 0 {
		pop := x % 10
		x /= 10
		if rev > math.MaxInt32/10 || (rev == math.MaxInt32/10 && pop > 7) {
			return 0
		}
		if rev < math.MinInt32/10 || (rev == math.MinInt32/10 && pop < -8) {
			return 0
		}
		rev = rev*10 + pop
	}
	return rev
}

func Reverse(x int) int {
	return reverse(x)
}
