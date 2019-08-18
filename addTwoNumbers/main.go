package main

import (
	"fmt"
	"leetcode/addTwoNumbers/code"
)

func readListNode(l *code.ListNode) {
	var list []int
	for l != nil {
		list = append(list, l.Val)
		l = l.Next
	}
	fmt.Println(list)
}

func main() {
	readListNode(code.AddTwoNumbers(
		&code.ListNode{
			Val: 2,
			Next: &code.ListNode{
				Val: 4,
				Next: &code.ListNode{
					Val: 3,
				},
			},
		},
		&code.ListNode{
			Val: 5,
			Next: &code.ListNode{
				Val: 6,
				Next: &code.ListNode{
					Val: 4,
				},
			},
		},
	))
	readListNode(code.AddTwoNumbers(
		&code.ListNode{
			Val: 5,
		},
		&code.ListNode{
			Val: 5,
		},
	))
	readListNode(code.AddTwoNumbers(
		&code.ListNode{
			Val: 1,
		},
		&code.ListNode{
			Val: 9,
			Next:&code.ListNode{
				Val:9,
			},
		},
	))
}
