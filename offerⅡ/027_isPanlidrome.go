package offer_

// 回文链表

// 找中点 + 反转链表
func isPalindrome(head *ListNode) bool {
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
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	return true
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{
		-1,
		head,
	}
	cur := head
	for cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = dummy.Next
		dummy.Next = next
	}

	return dummy.Next
}
