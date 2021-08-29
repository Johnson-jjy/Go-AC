package Binary_Tree

// 路径总和
// 解一:DFS -- 每次到叶子节点了判断即可
func hasPathSum1(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	var dfs func(root *TreeNode, cur int) bool

	dfs = func(root *TreeNode, cur int) bool {
		//fmt.Println(root.Val, cur)
		cur += root.Val
		if root.Left == nil && root.Right == nil {
			if cur == targetSum {
				return true
			} else {
				return false
			}
		}

		if root.Left != nil && dfs(root.Left, cur) {
			return true
		}
		if root.Right != nil && dfs(root.Right, cur) {
			return true
		}

		return false
	}
	return dfs(root, 0)
}

// 解二:BFS -- 将DFS逻辑写成BFS即可
func hasPathSum2(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	queue := make([]*TreeNode, 1)
	queue[0] = root
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			if node.Left == nil && node.Right == nil && node.Val == targetSum {
				return true
			}
			if node.Left != nil {
				node.Left.Val += node.Val
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				node.Right.Val += node.Val
				queue = append(queue, node.Right)
			}
		}
		queue = queue[size:]
	}

	return false
}
