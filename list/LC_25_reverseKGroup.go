package list

// 思路1:基于92基础思路,对k个一组的链表进行截断再反转操作,标注出相关特殊点即可
// 思路2:基于92进阶思路,先求全链表长度,再对最后需要截断的地方做出标记,之前的只需要依次反转即可
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}

	dummy := &ListNode {
		-1,
		nil,
	}
	dummy.Next = head
	start := dummy.Next
	oldTail := dummy
	//cur := dummy.Next
	pre := dummy

	for pre != nil {
		count := 0
		for count < k && pre.Next != nil { // 注意此处是pre.Next,否则会出现nextStart为空的情况
			pre = pre.Next
			count++
		}
		if count != k {
			break
		}
		nextStart := pre.Next
		pre.Next = nil
		newhead, newtail := reverseNormal(start)
		oldTail.Next = newhead
		newtail.Next = nextStart
		oldTail = newtail
		start = nextStart
		pre = newtail
	}

	return dummy.Next
}


func reverseKGroup2(head *ListNode, k int) *ListNode {
	lenth := 0
	for cur := head; cur != nil; lenth++{
		cur = cur.Next
	}

	lenth -= lenth % k // 对于要反转的部分进行划分

	prev := new(ListNode)
	newHead := prev
	prev.Next = head
	cur := prev.Next


	for i:=1; cur!=nil && i < lenth; i++{
		if i % k == 0{ // 对于不用反转的地方进行连接操作
			prev = cur
			cur = prev.Next
			continue
		}

		// 其余部分都正常反转即可
		next := cur.Next
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}

	return newHead.Next
}
