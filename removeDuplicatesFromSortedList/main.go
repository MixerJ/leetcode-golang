package main

import (
	"fmt"
	"leetcode/removeDuplicatesFromSortedList/code"
)

func readNode(listNode *code.ListNode) {
	var nums []int
	for listNode != nil {
		nums = append(nums, listNode.Val)
		listNode = listNode.Next
	}
	fmt.Printf("vals: %v\n", nums)
}

func main() {
	readNode(code.DeleteDuplicates(&code.ListNode{
		Val:1,
		Next:&code.ListNode{
			Val:1,
			Next:&code.ListNode{
				Val:2,
			},
		},
	}))
	readNode(code.DeleteDuplicates(&code.ListNode{
		Val:1,
		Next:&code.ListNode{
			Val:1,
			Next:&code.ListNode{
				Val:1,
			},
		},
	}))
	readNode(code.DeleteDuplicates(&code.ListNode{
		Val:1,
		Next:&code.ListNode{
			Val:1,
			Next:&code.ListNode{
				Val:2,
				Next:&code.ListNode{
					Val:3,
					Next:&code.ListNode{
						Val:3,
					},
				},
			},
		},
	}))
}