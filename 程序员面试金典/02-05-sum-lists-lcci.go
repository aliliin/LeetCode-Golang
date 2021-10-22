package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	res := &ListNode{}
	c := res
	temp := 0

	for l1 != nil || l2 != nil {
		resV := temp
		if l1 != nil {
			resV += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			resV += l2.Val
			l2 = l2.Next
		}

		res.Next = &ListNode{Val: resV % 10}
		temp = resV / 10

		res = res.Next
	}
	if temp != 0 {
		res.Next = &ListNode{Val: temp}
	}
	return c.Next
}
