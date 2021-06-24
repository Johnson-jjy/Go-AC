package list

// 快慢指针，常见
func kthToLast(head *ListNode, k int) int {
	flag := head

	for k > 0 {
		flag = flag.Next
		k--
	}

	// dummy技巧，常见
	dummy := &ListNode{
		-1,
		head,
	}
	pre := dummy

	for flag != nil {
		flag = flag.Next
		pre = pre.Next
	}

	return pre.Next.Val
}
