package list

// 结点定位手法
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	left, right := headA, headB

	for left != nil && right != nil {
		left = left.Next
		right = right.Next
	}
	if left == nil {
		for right != nil {
			headB = headB.Next
			right = right.Next
		}
	}else if right == nil {
		for left != nil {
			headA = headA.Next
			left = left.Next
		}
	}

	for headA != headB && headA != nil {
		headA = headA.Next
		headB = headB.Next
	}

	if headA == headB {
		return headA
	}
	return nil
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	//分别指向两个链表头节点
	ptr1, ptr2 := headA, headB
	//这里有两种情况，两个指针同时移动
	//1: 两个链表长度相同，有交点的话，直接定位到交点，如果没有交点，也会同时指向尾节点nil，循环断出，返回nil
	//2: 两个链表长度不同，遍历到尾节点后交替指向另一个头节点继续遍历，有交点的定位到交点，没有交点的遍历到尾部返回nil
	for ptr1 != ptr2 {
		if ptr1 == nil {
			ptr1 = headB
		} else {
			ptr1 = ptr1.Next
		}

		if ptr2 == nil {
			ptr2 = headA
		} else {
			ptr2 = ptr2.Next
		}
	}
	return ptr1
}
