package list

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//善用哑节点技巧
	dummy := &ListNode{0, head}
	first, second := head, dummy

	//判断条件容易搞混，一定要画图
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for ; first != nil; first = first.Next {
		second = second.Next
	}

	//注意初始化second为哑结点，此处已包含了对头结点的删除操作
	second.Next = second.Next.Next
	return dummy.Next
}

