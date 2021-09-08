package cig

import (
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 合法二叉搜索树
// 主要思想借鉴了CIG_0404: 每次都求当前节点对应左子树的最大值和右子树的最小值
// 注意:如果左节点或者右节点为空,就不要再递归了,再递归无法得到合理返回值(即使返回最大或者最小也都不妥)
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left != nil {
		leftMax := getLM(root.Left)
		if leftMax >= root.Val {
			return false
		}
	}
	if root.Right != nil {
		rightMax := getRM(root.Right)
		if root.Val >= rightMax {
			return false
		}
	}

	return isValidBST(root.Left) && isValidBST(root.Right)
}

func getLM(root *TreeNode) int {
	for root.Right != nil {
		root = root.Right
	}
	return root.Val
}

func getRM(root *TreeNode) int {
	for root.Left != nil {
		root = root.Left
	}

	return root.Val
}


// 此处每次传入的都是当前值，并对当前值的左右边界进行判断
func isValidBST2(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}
