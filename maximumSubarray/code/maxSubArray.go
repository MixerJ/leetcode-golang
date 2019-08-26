package code

// 计算特别耗时,内存消耗还行
func maxSubArrayOne(nums []int) int {
	var sumNum int
	var preSumNum = nums[0]
	var maxArray []int
	var count int
	for i := count; i < len(nums); i++ {
		sumNum += nums[i]
		maxArray = append(maxArray, nums[i])
		if preSumNum < sumNum {
			preSumNum = sumNum
		}
		if i+1 == len(nums) {
			sumNum = 0
			maxArray = []int{}
			count++
			i = count - 1
		}
	}
	return preSumNum
}

// 预留更好的计算方式
func maxSubArrayTwo(nums []int) int {
	var (
		maxSum = nums[0]
		//上次求和值（需大于0）
		lastSum = 0
	)

	for _, val := range nums {
		if lastSum > 0 {
			lastSum = lastSum + val
		} else {
			lastSum = val
		}
		if maxSum < lastSum {
			maxSum = lastSum
		}
	}
	return maxSum
}

func MaxSubArray(nums []int) int {
	return maxSubArrayOne(nums)
	//return maxSubArrayTwo(nums)
}
