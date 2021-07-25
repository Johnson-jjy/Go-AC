package Binary_Tree

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 思路: 在543求二叉树的直径基础上的进一步加强
// 注: 1. res124要取最小,因为节点值可以为负数; 2. 对于要更新全局变量的情况,则肯定要新写函数进行递归
var res124 int

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res124 = math.MinInt32

	getMax(root)

	return res124
}

func getMax(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := max(0, getMax(root.Left))
	right := max(0, getMax(root.Right))

	res = max(res, left + right + root.Val)
	return max(left, right) + root.Val
}
