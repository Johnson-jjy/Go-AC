package offer_

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 反转链表
// 法一: 迭代, 修改指针指向
func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

// 法2: 迭代, 类似头插法
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{
		-1,
		head,
	}
	cur := head
	for cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = dummy.Next
		dummy.Next = next
	}

	return dummy.Next
}

// 法3: 递归 -- 重点在于语义
func reverseList3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	tail := head.Next
	newHead := reverseList(tail) // 反转下一个点
	tail.Next = head
	head.Next = nil // 不能少

	return newHead
}