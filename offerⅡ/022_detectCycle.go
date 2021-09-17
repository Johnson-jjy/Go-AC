package offer_

// 链表中环的入口节点
// 快慢指针 -- 双倍
func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}
