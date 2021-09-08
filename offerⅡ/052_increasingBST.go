package offer_

//Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 展平二叉搜索树
// 其实还是说:1head,1pre,每次改变时要改root的left
func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var pre *TreeNode
	var head *TreeNode

	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		if head == nil {
			head = root
			pre = root
		} else {
			pre.Right = root
			root.Left = nil
			pre = root
		}
		inorder(root.Right)
	}

	inorder(root)

	return head
}
