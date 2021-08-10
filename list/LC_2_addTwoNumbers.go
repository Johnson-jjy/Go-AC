package list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 还是加法类题目 -> 注意进位,链表移动等
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{ // 注:此处不能声明 -> 仅声明则为默认的nil
		-1,
		nil,
	}
	pre := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		l := 0
		r := 0
		if l1 != nil {
			l = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			r = l2.Val
			l2 = l2.Next
		}
		cur := l + r + carry
		node := &ListNode {
			cur%10,
			nil,
		}
		pre.Next = node
		pre = pre.Next
		carry = cur/10
	}

	return dummy.Next
}
