package list

// 链表中间节点

// 双指针: 快慢指针; 注意初始化
func middleNode(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
