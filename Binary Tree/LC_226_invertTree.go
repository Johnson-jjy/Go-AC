package Binary_Tree

// 翻转二叉树
// 解法一: 递归
func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	tmp := invertTree1(root.Left)
	root.Left = invertTree1(root.Right)
	root.Right = tmp

	return root
}

// 解法二: BFS待补充