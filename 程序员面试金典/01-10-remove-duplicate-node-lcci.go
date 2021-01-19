package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针
func main(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	slow := head

	for slow != nil {
		fast := slow
		for fast.Next != nil {
			if slow.Val == fast.Next.Val {
				fast.Next = fast.Next.Next
			} else {
				fast = fast.Next
			}
		}
		slow = slow.Next
	}
	return head

}

// map 去重
func removeDuplicateNodes(Head *ListNode) *ListNode {
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
