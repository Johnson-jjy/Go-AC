//package main
////import . "nc_tools"
//
//type ListNode struct{
//  Val int
//  Next *ListNode
//}
//
//
///**
// * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
// *
// * @param m int整型
// * @param a ListNode类 指向彩带的起点，val表示当前节点的val，next指向下一个节点
// * @return ListNode类一维数组
// */
//func solve( m int ,  a *ListNode) []*ListNode {
//	// write code here
//	res := make([]*ListNode, m)
//	tails := make([]*ListNode, m)
//
//	for a != nil {
//		cur := &ListNode{
//			a.Val,
//			nil,
//		}
//		index := cur.Val % m
//		if res[index] == nil {
//			res[index] = cur
//			tails[index] = cur
//		} else {
//			tails[index].Next = cur
//			tails[index] = cur
//		}
//		a = a.Next
//	}
//
//	return res
//}
