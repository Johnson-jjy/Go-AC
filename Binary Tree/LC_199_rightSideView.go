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

// 二叉树的右视图: bfs,对每一层的数字直接存最后一个即可
func rightSideView(root *cig.TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*cig.TreeNode, 1)
	queue[0] = root

	for len(queue) != 0 {
		size := len(queue)
		cur := queue[size - 1]
		res = append(res, cur.Val)
		for i := 0; i < size; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[size:]
	}

	return res
}
