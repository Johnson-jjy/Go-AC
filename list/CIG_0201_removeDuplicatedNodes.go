package list

// 核心思想：用map保存出现过的结点
func removeDuplicateNodes1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	pre, cur := head, head.Next

	// 此处也可用bool，本题更节省空间
	store := map[int]*ListNode{}
	store[pre.Val] = pre

	for cur != nil {
		_, contain := store[cur.Val]
		if !contain {
			store[cur.Val] = cur
			pre = pre.Next
			cur = cur.Next
		} else {
			cur = cur.Next
			pre.Next = cur
		}
	}
	// 对最后一个结点的处理，可能它之后仍有重复值，故需要指向Nil
	pre.Next = nil

	return head
}

// 核心思想： 不使用额外空间，暴力搜索，类似deleteNode
func removeDuplicateNodes2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head

	for cur != nil {
		pre := cur
		check := cur.Next
		for check != nil {
			if check.Val != cur.Val {
				pre = pre.Next
				check = check.Next
			} else {
				check = check.Next
				pre.Next = check
			}
		}
		// 注意，cur继续走
		cur = cur.Next
	}

	return head
}