package offer

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 此处用的递归实现;三种解法详见-> list.LC_92
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tail := head.Next
	newHead := reverseList(tail)
	tail.Next = head
	head.Next = nil
	return newHead
}
