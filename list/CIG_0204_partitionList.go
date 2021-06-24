package list

// 对头插法的复习与应用；同时注意维护End结点
func partition(head *ListNode, x int) *ListNode {
	leftStart := &ListNode {
		-1,
		nil,
	}
	leftEnd := leftStart
	midStart := &ListNode {
		-1,
		nil,
	}
	midEnd := midStart
	rightStart := &ListNode {
		-1,
		nil,
	}
	//rightEnd := rightStart

	cur := head
	tmp := &ListNode {
		-1,
		cur,
	}
	for cur != nil {
		tmp = cur
		cur = cur.Next
		if tmp.Val < x {
			if leftStart.Next == nil {
				leftEnd = tmp
				tmp.Next = nil
			}
			tmp.Next = leftStart.Next
			leftStart.Next = tmp
		} else if tmp.Val == x {
			if midStart.Next == nil {
				midEnd = tmp
				tmp.Next = nil
			}
			tmp.Next = midStart.Next
			midStart.Next = tmp
		} else {
			tmp.Next = rightStart.Next
			rightStart.Next = tmp
		}
	}

	// 对不同的情况分类讨论
	if leftStart.Next != nil {
		if midStart.Next != nil {
			leftEnd.Next = midStart.Next
			midEnd.Next = rightStart.Next
		} else {
			leftEnd.Next = rightStart.Next
		}
		return leftStart.Next
	} else {
		if midStart.Next != nil {
			midEnd.Next = rightStart.Next
			return midStart.Next
		} else {
			return rightStart.Next
		}
	}
}
