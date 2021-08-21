package Binary_Tree

import "Go-AC/cig"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 求二叉树的直径
// 关键是理解题意: 二叉树的直径取自任意节点左右子树深度和的最大值
var res int
func diameterOfBinaryTree(root *cig.TreeNode) int {
	if root == nil {
		return 0
	}
	res = 0
	depth(root)

	return res
}

func depth(root *cig.TreeNode) int {
	if root == nil {
		return 0
	}
	left := depth(root.Left)
	right := depth(root.Right)

	res = max(res, left + right)

	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
