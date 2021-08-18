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

// 类105
func buildTree106(inorder []int, postorder []int) *cig.TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	root := &cig.TreeNode{
		postorder[len(postorder) - 1],
		nil,
		nil,
	}
	if len(postorder) == 1 {
		return root
	}
	index := 0
	for i := 0; i < len(inorder); i++ {
		if root.Val == inorder[i] {
			index = i
			break
		}
	}

	root.Left = buildTree106(inorder[:index], postorder[:index])
	root.Right = buildTree106(inorder[index + 1:], postorder[index: len(postorder)-1])

	return root
}
