package offer

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//注: 当函数对输入值得判断和逻辑中对递归函数输入值的判断不同时,就需要新的函数而不能再自身递归了
//本题即需要一个新的递归函数以判断传入的两个节点是否满足包含关系
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
