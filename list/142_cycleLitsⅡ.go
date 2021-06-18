package list

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		//确认有环后，从头结点出发找成环点，可画图理解
		if fast == slow {
			start := head
			for start != slow {
				slow = slow.Next
				start = start.Next
			}
			return start
		}
	}

	return nil
}
