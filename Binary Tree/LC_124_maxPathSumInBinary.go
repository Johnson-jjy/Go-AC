package Binary_Tree

import (
	"Go-AC/cig"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 二叉树中的最大路径和
// 思路: 在543求二叉树的直径基础上的进一步加强
// 注: 1. res124要取最小,因为节点值可以为负数; 2. 对于要更新全局变量的情况,则肯定要新写函数进行递归
var res124 int

func maxPathSum(root *cig.TreeNode) int {
	if root == nil {
		return 0
	}
	res124 = math.MinInt32

	getMax(root)

	return res124
}

func getMax(root *cig.TreeNode) int {
	if root == nil {
		return 0
	}
	left := max(0, getMax(root.Left)) // 对于小于0的情况,直接不取
	right := max(0, getMax(root.Right))

	res = max(res, left + right + root.Val)
	return max(left, right) + root.Val // 注意此处的返回值, 只需返回left或right + Val即可
}

// 2021/08/20 更新
//var res int
//
//func maxPathSum(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	res = math.MinInt32
//	getSum(root)
//
//	return res
//}
//
//func getSum(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	left := getSum(root.Left)
//	right := getSum(root.Right)
//	sum := root.Val
//	if left > 0 {
//		sum += left
//	}
//	if right > 0 {
//		sum += right
//	}
//	res = max(res, sum)
//
//	if sum < 0 {
//		return 0
//	}
//
//	return root.Val + max(left, right)
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
