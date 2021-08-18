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
func zigzagLevelOrder(root *cig.TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := make([]*cig.TreeNode, 1)
	queue[0] = root
	flag := 0

	for len(queue) > 0 {
		size := len(queue)
		//fmt.Println(size)
		cur := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue[i]
			if flag == 0 {
				cur[i] = node.Val
			} else if flag == 1{
				cur[size - i - 1] = node.Val
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		flag = (flag + 1) % 2
		queue = queue[size:]
		res = append(res, cur)
	}

	return res
}
