package main
//import . "nc_tools"

type ListNode struct{
  Val int
  Next *ListNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param a ListNode类一维数组 指向这些数链的开头
 * @return ListNode类
 */
func solve( a []*ListNode ) *ListNode {
	// write code here
	dummy := &ListNode{
		-1,
		nil,
	}
	pre := dummy
	flag := true
	for flag {
		flag = false
		for i := 0; i < len(a); i++ {
			if a[i] == nil {
				continue
			}
			flag = true
			pre.Next = a[i]
			a[i] = a[i].Next
			pre = pre.Next
		}
	}

	return dummy.Next
}