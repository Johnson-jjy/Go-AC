package Binary_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 二叉树的最大深度
// 1.DFS
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	return max(left, right) + 1
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

// 2.BFS
func maxDepth2(root *TreeNode) int {
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
