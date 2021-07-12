package offer

// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func reversePrint(head *ListNode) []int {
	stack := make([]int, 0)
	cur := head
	for cur != nil {
		stack = append(stack, cur.Val)
		cur = cur.Next
	}
	m := len(stack)
	res := make([]int, m)
	for i := 0; i < m; i++ {
		num := stack[len(stack)-1]
		res[i] = num
		stack = stack[:len(stack)-1]
	}
	return res
}