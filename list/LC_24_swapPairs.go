package list

// 两两交换链表中的节点
// 注意: 一定要有dummy; 且每次要更新pre; cur的退出条件也不要遗漏了
func swapPairs(head *ListNode) *ListNode {
	cur := head
	dummy := &ListNode {
		-1,
		head,
	}
	pre := dummy

	for cur != nil && cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = cur
		pre.Next = next
		pre = cur
		cur = cur.Next
	}

	return dummy.Next
}
