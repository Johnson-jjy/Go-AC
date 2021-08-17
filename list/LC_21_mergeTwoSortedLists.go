package list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 本质思想类似归并排序
// 注意:两链表归并时,若某一已经nil,则根据情况直接pre.Next指向即可,而不是继续走for循环,且走for循环时注意别赋值错了
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode {
		-1,
		nil,
	}
	pre := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			pre.Next = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			l2 = l2.Next
		}
		pre = pre.Next
	}
	// 以下代码突出一个丑陋:既多走了for,还在for里将pre的值赋错了
	// for l != nil {
	//     pre = l
	//     //fmt.Println(pre.Val)
	//     pre = pre.Next
	//     l = l.Next
	// }
	// for r != nil {
	//     pre = r
	//     //fmt.Println(pre.Val)
	//     pre = pre.Next
	//     r = r.Next
	// }
	if l1 != nil {
		pre.Next = l1
	} else {
		pre.Next = l2
	}

	return dummy.Next
}
