package offer

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	left := getHeight(root.Left)
	right := getHeight(root.Right)

	if abs(left, right) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := getHeight(root.Left)
	right := getHeight(root.Right)

	return max(left, right) + 1
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
