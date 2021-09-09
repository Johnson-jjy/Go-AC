package offer_

import "math"

// 二叉树每层的最大值
// 解一: 标准的BFS
func largestValues(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		curMax := math.MinInt32

		for i := 0; i < size; i++ {
			node := queue[i]
			if node.Val > curMax {
				curMax = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, curMax)
		queue = queue[size:]
	}

	return res
}

// 递归 -- 每一次传入deep