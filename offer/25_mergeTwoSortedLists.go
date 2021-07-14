package offer

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 解1: 构造了新链表,未对原链表进行操作
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{
		-1,
		nil,
	}
	cur := dummy
	for l1 != nil && l2 != nil {
		val1 := l1.Val
		val2 := l2.Val
		var node *ListNode
		if val1 < val2 {
			node = &ListNode{
				val1,
				nil,
			}
			l1 = l1.Next
		} else {
			node = &ListNode{
				val2,
				nil,
			}
			l2 = l2.Next
		}
		cur.Next = node
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}

	return dummy.Next
}

// 解2: 在原链表的基础上进行了操作
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	var cur *ListNode
	if l1.Val < l2.Val {
		cur = l1
		l1 = l1.Next
	} else {
		cur = l2
		l2 = l2.Next
	}
	head := cur

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}

	return head
}