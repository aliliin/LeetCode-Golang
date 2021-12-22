package main

func isPalindrome(head *ListNode) bool {
	slow := head
	fast := head
	var prev *ListNode

	for fast != nil && fast.Next != nil {
		curOld := slow
		slow = slow.Next
		fast = fast.Next.Next
		curOld.Next = prev
		prev = curOld
	}

	if fast != nil {
		slow = slow.Next
	}

	for slow != nil {
		if slow.Val != prev.Val {
			return false
		}
		slow = slow.Next
		prev = prev.Next
	}
	return true

}
