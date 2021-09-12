package offer_

// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

// 合并排序链表
// 解一: 归并 -- 注意,必须对len == 1时进行处理,(本质即left <= right)的情况;否则利用mid划分会进入无限循环
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	mid := len(lists)/2
	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:])
	return merge(left, right)
}

func merge(left *ListNode, right *ListNode) *ListNode {
	dummy := &ListNode{
		-1,
		nil,
	}
	pre := dummy

	for left != nil && right != nil {
		if left.Val < right.Val {
			pre.Next = left
			left = left.Next
		} else {
			pre.Next = right
			right = right.Next
		}
		pre = pre.Next
	}
	if left != nil {
		pre.Next = left
	}
	if right != nil {
		pre.Next = right
	}

	return dummy.Next
}
