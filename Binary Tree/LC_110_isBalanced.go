package Binary_Tree

import "Go-AC/cig"

// 平衡二叉树
//Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isBalanced(root *cig.TreeNode) bool {
	if root == nil {
		return true
	}
	left := getH(root.Left)
	right := getH(root.Right)

	if left - right > 1 || right - left > 1 {
		return false
	}

	return isBalanced(root.Right) && isBalanced(root.Left)
}

func getH(root *cig.TreeNode) int {
	if root == nil {
		return 0
	}
	left := getH(root.Left)
	right := getH(root.Right)

	return max(left, right) + 1
}
