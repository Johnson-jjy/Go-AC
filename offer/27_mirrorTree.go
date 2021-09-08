package offer

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 二叉树的镜像
// 解法一: 递归
func mirrorTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	temp := mirrorTree1(root.Left)
	root.Left = mirrorTree1(root.Right)
	root.Right = temp
	return root
}

// 解法二: 层序遍历,注意,不要拘泥于对size的使用
func mirrorTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	stack := make([]*TreeNode, 1)
	stack[0] = root
	for len(stack) != 0 {
		cur := stack[len(stack) - 1]
		stack = stack[0 : len(stack) - 1]
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		temp := cur.Left
		cur.Left = cur.Right
		cur.Right = temp
	}
	return root
}
