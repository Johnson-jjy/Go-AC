package list

type ListNode struct {
	Val int
	Next *ListNode
}

//迭代版本
func reverseList1(head *ListNode) *ListNode {
	var prev *ListNode //默认为空，若此处声明为pre、cur、next则头结点的Next始终不指向空（除非加if判据）
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

//递归版本
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	head.Next = nil
	newHead := reverseList2(next)
	next.Next = head

	return newHead
}