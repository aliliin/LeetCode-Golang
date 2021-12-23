package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 遍历完一条链表后，跳转到另一条链表继续遍历，遇到第一个相同的节点就是目标节点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	l1, l2 := headA, headB
	for l1 != l2 {
		if l1 == nil {
			l1 = headB
		} else {
			l1 = l1.Next
		}
		if l2 == nil {
			l2 = headA
		} else {
			l2 = l2.Next
		}
	}
	return l1

}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	currA := headA
	currB := headB
	lenA, lenB := 0, 0
	// 求A，B的长度
	for currA != nil {
		currA = currA.Next
		lenA++
	}
	for currB != nil {
		currB = currB.Next
		lenB++
	}
	var step int
	var fast, slow *ListNode
	// 请求长度差，并且让更长的链表先走相差的长度
	if lenA > lenB {
		step = lenA - lenB
		fast, slow = headA, headB
	} else {
		step = lenB - lenA
		fast, slow = headB, headA
	}
	for i := 0; i < step; i++ {
		fast = fast.Next
	}
	// 遍历两个链表遇到相同则跳出遍历
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

func main() {

}
