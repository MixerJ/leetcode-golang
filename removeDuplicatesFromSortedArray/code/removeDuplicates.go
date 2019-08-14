package code

func removeDuplicates(nums []int) int {
	var preNum int
	if len(nums) == 1 {
		return len(nums)
	}
	for key := 0; key < len(nums); key++ {
		if preNum == nums[key] && key != 0 {
			if key+1 > len(nums) {
				break
			}
			nums = append(nums[:key], nums[key+1:]...)
			key -= 1
		}
		preNum = nums[key]
	}
	return len(nums)
}

func RemoveDuplicates(nums []int) int {
	return removeDuplicates(nums)
}
