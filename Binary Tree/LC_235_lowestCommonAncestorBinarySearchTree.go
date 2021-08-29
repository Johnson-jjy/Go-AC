package Binary_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

// 二叉搜索树的最近公共祖先
// 利用二叉树搜索树节点大小性质进行节点位置的判断
func lowestCommonAncestor235(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor235(root.Left, p, q)
	}
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor235(root.Right, p, q)
	}
	return root
}
