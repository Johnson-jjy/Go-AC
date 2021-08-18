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

// 解法一: 迭代版本
func inorderTraversal1(root *cig.TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	stack := make([]*cig.TreeNode, 0)

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		res = append(res, node.Val)
		root = node.Right
	}

	return res
}

// 解法二: 递归版本
func inorderTraversal2(root *cig.TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	var inOrder func(root *cig.TreeNode)
	inOrder = func(root *cig.TreeNode) {
		if root == nil {
			return
		}
		inOrder(root.Left)
		res = append(res, root.Val)
		inOrder(root.Right)
	}

	inOrder(root)

	return res
}
