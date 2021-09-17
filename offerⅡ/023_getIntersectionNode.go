package offer_

// 两个链表的第一个重合节点
// 类似快慢指针
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	left := headA
	right := headB

	for left != nil && right != nil {
		left = left.Next
		right = right.Next
	}

	if left == nil {
		left = headB
	} else {
		right = headA
	}

	for left != nil && right != nil {
		left = left.Next
		right = right.Next
	}

	if left == nil {
		left = headB
	} else {
		right = headA
	}

	for left != nil {
		if left == right {
			return left
		}
		left = left.Next
		right = right.Next
	}

	return nil
}
