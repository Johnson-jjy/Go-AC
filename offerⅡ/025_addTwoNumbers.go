package offer_

// 链表中的两数相加

// 核心思想: 链表反转 + 链表相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverse(l1)
	l2 = reverse(l2)

	carry := 0
	dummy := &ListNode {
		-1,
		nil,
	}
	for l1 != nil || l2 != nil || carry != 0 {
		n1 := 0
		n2 := 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		cur := carry + n1 + n2
		node := &ListNode {
			cur % 10,
			dummy.Next,
		}
		dummy.Next = node
		carry = cur / 10
	}

	return dummy.Next
}
