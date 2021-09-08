package offer

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 从上到下打印二叉树Ⅱ
// 同一 -> 依然是BFS模板题,唯一的不同在于对每一层的变量进行保存
// 注意: 每次要用一个新数组保存,否则追加到res数组的结果会被覆盖
func levelOrder2(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := make([]*TreeNode, 1)
	queue[0] = root

	for len(queue) != 0 {
		size := len(queue)
		cur := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			cur[i] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, cur)
	}

	return res
}
