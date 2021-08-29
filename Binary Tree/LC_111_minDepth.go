package Binary_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 二叉树的最小深度
// 1.求最小时建议用BFS,及时退出
func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	queue := make([]*TreeNode, 1)
	queue[0] = root
	for len(queue) > 0 {
		size := len(queue)
		res++
		for i := 0; i < size; i++ {
			node := queue[i]
			if node.Left == nil && node.Right == nil {
				return res
			}
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

// 解二:DFS, 注意,并不是找left和right的最小值就行(max就是找最大值就行,min不同)
// 而是必须取到叶子节点时,才能用min
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := minDepth2(root.Left)
	right := minDepth2(root.Right)
	if left != 0 && right != 0 {
		return min(left, right) + 1
	} else if left == 0 {
		return right + 1
	}
	return left + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
