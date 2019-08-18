package code

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1.Val == 0 && l1.Next == nil {
		return l2
	}
	if l2.Val == 0 && l2.Next == nil {
		return l1
	}
	var pre int
	var saveNode = &ListNode{}
	returnNode := saveNode
	for l1 != nil || l2 != nil || pre != 0 {
		saveNode = &ListNode{}
		var sum int
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += pre
		pre = sum / 10
		saveNode.Next = &ListNode{Val: sum % 10, Next: nil}
		saveNode = saveNode.Next
	}
	return returnNode.Next
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbers(l1, l2)
}
