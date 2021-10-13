package offer_

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 删除链表的倒数第n个节点
// 解法一: 先统计list有多少个节点; 进而将倒数转化成正数
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		-1,
		head,
	}
	pre := dummy
	length := 0
	cur := pre.Next
	for cur != nil {
		cur = cur.Next
		length++
	}
	target := length - n + 1
	for target > 1 {
		pre = pre.Next
		target--
	}
	pre.Next = pre.Next.Next
	return dummy.Next
}

// 解法二: 双指针, 一前一后, 一次遍历进行定位
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		-1,
		head,
	}
	pre := dummy
	cur := pre.Next // 这两行的设定可画图确认
	for n > 1 {
		cur = cur.Next
		n--
	}
	for cur.Next != nil {
		pre = pre.Next
		cur = cur.Next
	}
	pre.Next = pre.Next.Next
	return dummy.Next
}