package list

// 快慢指针经典用法--找终点 + 反转链表
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	flag := reverse(slow)

	cur := head
	for cur != nil && flag != nil {
		if cur.Val != flag.Val {
			return false
		}
		cur = cur.Next
		flag = flag.Next
	}

	return true
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head
	next := cur
	var pre *ListNode
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}
