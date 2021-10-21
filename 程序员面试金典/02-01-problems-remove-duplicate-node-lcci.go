package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeDuplicateNodes(head *ListNode) *ListNode {
	m := make(map[int]struct{})
	dummyHead := &ListNode{}

	pre := dummyHead

	for head != nil {
		if _, ok := m[head.Val]; !ok {
			m[head.Val] = struct{}{}
			pre.Next = head
			pre = head
		} else {
			pre.Next = head.Next
		}
		head = head.Next
	}
	return dummyHead.Next

}
