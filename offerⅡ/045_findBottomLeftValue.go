package offer_

// 二叉树最底层最左边的值
// 解一: 层序遍历, 每次记录当前层的最左边值
func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return -1
	}

	queue := make([]*TreeNode, 1)
	queue[0] = root

	res := -1
	for len(queue) > 0 {
		size := len(queue)
		res = queue[0].Val

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

// 解二: 前序遍历; 待补充 特点:每次都能记录每层的最左边的节点