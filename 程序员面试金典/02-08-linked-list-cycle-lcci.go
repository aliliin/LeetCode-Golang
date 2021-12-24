package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 快慢指针
// 我们使用两个指针， fast 与 slow 它们起始都位于链表的头部。随后，slow 指针每次向后移动一个位置，而 fast 指针向后移动两个位置。
// 如果链表中存在环，则 fast 指针最终将再次与 slow 指针在环中相遇。
func detectCycle1(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}

// 我们遍历链表中的每个节点，并将它记录下来；一旦遇到了此前遍历过的节点，就可以判定链表中存在环。借助哈希表可以很方便地实现
func detectCycle(head *ListNode) *ListNode {
	s := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := s[head]; ok {
			return head
		}
		s[head] = struct{}{}
		head = head.Next
	}
	return nil
}

func main() {

}
