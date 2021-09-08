package offer

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 从上到下打印二叉树Ⅲ
// 放入cur的顺序调整为逆序即可
// 另: 勿忘对于切片的添加操作 -- append...
func levelOrder3(root *TreeNode) [][]int {
	res := make([][]int, 0)

	if root == nil {
		return res
	}

	queue := make([]*TreeNode, 1)
	queue[0] = root

	flag := 0
	for len(queue) != 0 {
		size := len(queue)
		cur := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			if flag == 0 {
				cur[i] = node.Val
			} else {
				cur[size - 1 - i] = node.Val
			}
		}
		flag = (flag + 1) % 2
		res = append(res, cur)
	}

	return res
}

// 另注: 如果要用一维切片ans存储最后的结果,则用如下语句
// ans = append(ans, cur...) -> 此时可用append连接两个切片