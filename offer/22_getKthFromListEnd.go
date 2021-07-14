package offer

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 快慢指针:善用哑节点,注意和LC-19区别
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{
		-1,
		head,
	}
	cur := dummy
	pre := dummy

	for k > 0 {
		cur = cur.Next
		k--
	}
	for cur.Next != nil {
		pre = pre.Next
		cur = cur.Next
	}

	return pre.Next
}
