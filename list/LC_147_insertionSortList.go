package list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 选择排序: 链表中插入可不swap --> 与数组不同,从前往后扫而非从后往前
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{ // 巧妙解决插入在头结点的情况
		-1,
		head,
	}
	cur := head.Next
	lastSorted := head
	for cur != nil {
		if cur.Val > lastSorted.Val { // 先判断,减少扫描次数
			cur = cur.Next
			lastSorted = lastSorted.Next // 也要相应前移
		} else {
			pre := dummy
			for pre.Next.Val < cur.Val {
				pre = pre.Next
			}
			// 一下逻辑交换理清楚
			lastSorted.Next = cur.Next
			cur.Next = pre.Next
			pre.Next = cur
			cur = lastSorted.Next
		}
	}

	return dummy.Next
}
