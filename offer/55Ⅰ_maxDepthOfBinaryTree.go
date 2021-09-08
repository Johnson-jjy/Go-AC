package offer

// 二叉树的深度
// 求深度时返回max值+1即可; 注意要和求最短区分!
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	return max(left, right) + 1
}
