package list

// 反转链表Ⅱ

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




// 08/05 bug free:保留好相关节点即可 -> 也可以直接换,不扫描
func reverseBetween2(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	dummy := &ListNode{
		-1,
		nil,
	}
	dummy.Next = head
	pre := dummy
	for i := 1; i < left; i++ {
		pre = pre.Next
	}
	start := pre.Next
	pre.Next = nil
	oldtail := pre
	pre = start
	for i := left; i < right; i++ {
		pre = pre.Next
	}
	var end *ListNode
	if pre.Next != nil {
		end = pre.Next
		pre.Next = nil
	}
	newhead, tail := reverseNormal(start)
	oldtail.Next = newhead
	tail.Next = end

	return dummy.Next
}

func reverseNormal(head *ListNode) (newhead *ListNode, tail *ListNode) {
	if head == nil || head.Next == nil {
		return head, head
	}

	next := head.Next
	head.Next = nil
	newhead, tail = reverseNormal(next)
	tail.Next = head
	tail = head

	return
}