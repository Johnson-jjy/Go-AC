package offer

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	cur1 := headA
	cur2 := headB

	for cur1 != nil && cur2 != nil {
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	if cur1 == nil {
		cur1 = headB
	} else {
		cur2 = headA
	}
	for cur1 != nil && cur2 != nil {
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	if cur1 == nil {
		cur1 = headB
	} else {
		cur2 = headA
	}
	for cur1 != nil && cur2 != nil {
		if cur1 == cur2 {
			return cur1
		}
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	return nil
}
