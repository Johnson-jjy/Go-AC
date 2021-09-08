package offer

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 树的子结构
// 注: 当函数对输入值得判断和逻辑中对递归函数输入值的判断不同时,就需要新的函数而不能再自身递归了
// 本题即需要一个新的递归函数以判断传入的两个节点是否满足包含关系
// 深入理解check! --> 它必须一脉相承地判断!
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	return (check(A, B)) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func check(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}

	return check(A.Left, B.Left) && check(A.Right, B.Right)
}

// 0908错误版本: 错误原因--check函数的构想有错误;判断是否是子结构需要一直判断,不能中断
//func isSubStructure(A *TreeNode, B *TreeNode) bool {
//	if B == nil {
//		return false
//	}
//	return check(A, B)
//}
//
//func check(p *TreeNode, q *TreeNode) bool {
//	if q == nil {
//		return true
//	}
//	if p == nil {
//		return false
//	}
//
//	if p.Val == q.Val {
//		return check(p.Left, q.Left) && check(p.Right, q.Right)
//	}
//
//	return check(p.Left, q) || check(p.Right, q)
//}