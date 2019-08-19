package code

func searchInsert(nums []int, target int) int {
	for key, value := range nums {
		if value == target {
			return key
		}
		if target < value {
			return key
		}
	}
	return len(nums)
}

func SearchInsert(nums []int, target int) int {
	return searchInsert(nums, target)
}
