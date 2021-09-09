package offer_

// 从根节点到叶节点的路径数字之和
// 解一: 前序遍历(其实应该本质上顺序不影响) -- 类似于求minDepth;
// 只有当本节点是叶子节点时才做处理,同理,对非nil的情况则直接不往下递归
func sumNumbers1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	var dfs func(root *TreeNode, cur int)
	dfs = func(root *TreeNode, cur int) {
		cur = cur * 10 + root.Val

		if root.Left == nil && root.Right == nil {
			res += cur
			return
		}

		if root.Left != nil {
			dfs(root.Left, cur)
		}
		if root.Right != nil {
			dfs(root.Right, cur)
		}
	}

	dfs(root, 0)

	return res
}

// 解二: BFS -- 除了保存节点以外, 对值用另外的队列保存,此为常见双队列保存操作
func sumNumbers2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	queue := make([]*TreeNode, 1)
	queue[0] = root
	numQ := make([]int, 1)
	numQ[0] = root.Val

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			node := queue[i]
			cur := numQ[i]
			if node.Left == nil && node.Right == nil {
				res += cur
				continue
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
				numQ = append(numQ, cur * 10 + node.Left.Val)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				numQ = append(numQ, cur * 10 + node.Right.Val)
			}
		}

		queue = queue[size:]
		numQ = numQ[size:]
	}

	return res
}