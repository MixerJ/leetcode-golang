package code

func climbStairs(n int) int {
	var nums = make([]int, n+3)
	nums[0] = 1
	nums[1] = 1
	for i := 2; i <= n; i++ {
		nums[i] = nums[i -1] + nums[i-2]

	}
	return nums[n]
}
func climbStairsTwo(n int) int {
	first:= 1
	second := 1
	for i := 2; i <= n; i++ {
		third := first + second
		first = second
		second = third
	}
	return second
}
func ClimbStairs(n int) int {
	//return climbStairs(n)
	return climbStairsTwo(n)
}
