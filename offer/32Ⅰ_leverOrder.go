package offer

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 从上倒下打印二叉树
// 本题为BFS模板题
// 需要注意的是 -> queue中存TreeNode,res中存int
func levelOrder(root *TreeNode) []int {
	res := make([]int, 0)

	if root == nil {
		return res
	}

	queue := make([]*TreeNode, 1)
	queue[0] = root

	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			res = append(res, node.Val)
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return res
}
