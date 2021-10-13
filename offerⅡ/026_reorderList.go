package offer_

// 重排链表

// 本质还是 找中点 + 反转 -> 然后依据规则重排
func reorderList(head *ListNode)  {
	if head == nil || head.Next == nil {
		return
	}
	dummy := &ListNode {
		-1,
		head,
	}
	slow := dummy
	fast := dummy
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	start := slow.Next
	slow.Next = nil
	l2 := reverse(start)
	l1 := head
	for l2 != nil {
		n1 := l1.Next
		n2 := l2.Next
		l1.Next = l2
		l2.Next = n1
		l1 = n1
		l2 = n2
	}

	return
}
