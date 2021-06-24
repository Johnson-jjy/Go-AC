package list

// 注意对进位的处理
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cur := 0
	left := 0
	right := 0
	carry := 0

	dummy := &ListNode {
		-1,
		nil,
	}
	pre := dummy

	for l1 != nil || l2 != nil {
		cur = 0
		if l1 == nil {
			left = 0
		} else {
			left = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			right = 0
		} else {
			right = l2.Val
			l2 = l2.Next
		}
		cur = left + right + carry
		carry = cur / 10
		cur %= 10
		node := &ListNode {
			cur,
			nil,
		}
		pre.Next = node
		pre = node
	}
	if carry > 0 {
		node := &ListNode {
			carry,
			nil,
		}
		pre.Next = node
	}

	return dummy.Next
}
