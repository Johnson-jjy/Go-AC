package offer

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 注意dummy节点的设置,cur从dummy开始,可能删头结点
func deleteNode(head *ListNode, val int) *ListNode {
	dummy := &ListNode{
		-1,
		head,
	}
	cur := dummy
	for cur != nil {
		next := cur.Next
		if next.Val == val { // 若可能无节点时还要考虑next==nil
			cur.Next = next.Next
			return dummy.Next
		}
		cur = cur.Next
	}
	return nil
}
