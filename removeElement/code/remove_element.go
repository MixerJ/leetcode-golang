package code

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	for key := 0; key < len(nums); key++ {
		if nums[key] == val {
			if key+1 > len(nums) {
				break
			}
			nums = append(nums[:key], nums[key+1:]...)
			key -= 1
		}
	}
	return len(nums)
}

func RemoveElement(nums []int, val int) int {
	return removeElement(nums, val)
}
