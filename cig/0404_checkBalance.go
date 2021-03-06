package cig

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 检查平衡性
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	lH := getHeight(root.Left)
	rH := getHeight(root.Right)
	if lH - rH > 1 || rH - lH > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}


	leftH := getHeight(root.Left)
	rightH := getHeight(root.Right)

	if leftH >= rightH {
		return leftH + 1
	}
	return rightH + 1
}
