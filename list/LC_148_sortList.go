package list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 思路: 归并排序-->自顶向下和自底向上均可
// 注意: 不能用快排!无法实现从尾到头的扫描以划分区间
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := findMin(head)
	rightHead := p.Next
	p.Next = nil
	left := sortList(head)
	right := sortList(rightHead)

	return merge(left, right)
}

func findMin(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode {
		-1,
		head,
	}
	slow := dummy
	fast := dummy

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func merge(left, right *ListNode) *ListNode {
	dummy := &ListNode {
		-1,
		nil,
	}
	cur := dummy
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
		cur = cur.Next
	}

	for left != nil {
		cur.Next = left
		left = left.Next
		cur = cur.Next
	}
	for right != nil {
		cur.Next = right
		right = right.Next
		cur = cur.Next
	}

	return dummy.Next
}
