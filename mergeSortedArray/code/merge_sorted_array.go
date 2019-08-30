package code

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	allLength := m + n - 1
	m -= 1
	n -= 1
	for m >= 0 && n >= 0 {
		if nums1[m] < nums2[n] {
			nums1[allLength] = nums2[n]
			n -= 1
		} else {
			nums1[allLength] = nums1[m]
			m -= 1
		}
		allLength -= 1
	}
	if m < 0 && n >= 0 {
		for i := 0; i <= n; i++ {
			nums1[i] = nums2[i]
		}
	}
	fmt.Println(nums1)
}

func Merge(nums1 []int, m int, nums2 []int, n int) {
	merge(nums1, m, nums2, n)
}
