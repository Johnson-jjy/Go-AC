package Binary_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder102(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := make([]*TreeNode, 1)
	queue[0] = root

	for len(queue) > 0 {
		size := len(queue)
		//fmt.Println(size) -> 别忘了缩减queue
		cur := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue[i]
			cur[i] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[size:]
		res = append(res, cur)
	}

	return res
}
