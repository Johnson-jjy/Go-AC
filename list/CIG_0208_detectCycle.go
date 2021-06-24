package list

func detectCycle2(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			cur := head
			for cur != slow {
				slow = slow.Next
				cur = cur.Next
			}
			return cur
		}
	}

	return nil
}
