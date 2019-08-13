package main

import (
	"fmt"
	"leetcode/mergeTwoSortedLists/code"
)

func main() {
	l1 := code.ListNode{
		Val: 1,
		Next: &code.ListNode{
			Val: 2,
			Next: &code.ListNode{
				//Val:3,
				Val: 4,
			},
		},
	}
	l2 := code.ListNode{
		Val: 1,
		Next: &code.ListNode{
			Val: 3,
			Next: &code.ListNode{
				Val: 4,
			},
		},
	}
	l3 := code.MergeTwoLists(&l1, &l2)
	for l3.Next != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}
}
