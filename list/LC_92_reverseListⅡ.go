package list

//注1：一次遍历解决问题
//注2：反转链表的三种解决方式：a、递归； b、改指针指向； c、交换节点；本解法在反转部分使用了c；用b同理，但需多保存几个节点信息
//注3：一定需要留意如何节点的移动次数，多画图
func reverseBetween(head *ListNode, left, right int) *ListNode {
	// 设置 dummyNode 是这一类问题的一般做法
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head
	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummyNode.Next
}

