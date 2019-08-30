package code

func mySqrt(x int) int {
	var sqrtX int
	for i := 1; i <= x/2+1; i++ {
		if i*i == x {
			sqrtX = i
			break
		} else if x > (i-1)*(i-1) && x < i*i {
			sqrtX = i - 1
			break
		} else {
			continue
		}
	}
	return sqrtX
}
func MySqrt(x int) int {
	return mySqrt(x)
}
