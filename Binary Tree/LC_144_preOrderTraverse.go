package Binary_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 非递归版: 用栈 -> 用栈模拟递归站, 明确加入res的时间 和 往左右子树的不同便利 是关键
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	stack := make([]*TreeNode, 0)

	for len(stack) > 0 || root != nil {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		root = node.Right
	}

	return res
}

// 解法二: 递归版
var res144 []int

func preorderTraversal1(root *TreeNode) []int {
	res144 = make([]int, 0)

	if root == nil {
		return res144
	}

	preOrder(root)

	return res144
}

func preOrder(root *TreeNode) {
	if root == nil {
		return
	}

	res144 = append(res144, root.Val)
	preOrder(root.Left)
	preOrder(root.Right)
}
