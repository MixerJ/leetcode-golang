package code

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var preNode = new(ListNode)
	saveNode := preNode
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			preNode.Next = l1
			l1 = l1.Next
		} else {
			preNode.Next = l2
			l2 = l2.Next
		}
		preNode = preNode.Next
	}
	if l2 != nil {
		preNode.Next = l2
	} else if l1 != nil {
		preNode.Next = l1
	}
	return saveNode.Next
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	return mergeTwoLists(l1, l2)
}