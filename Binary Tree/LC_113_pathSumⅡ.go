package Binary_Tree

// 路径总和Ⅱ
// 较一进阶的地方只在需要保存路径 -> 每次需要新声明一个tmp, 再做相应的路径添加
// 同一: 注意要判断到了叶子节点,再做相关操作
func pathSum(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	var dfs func(root *TreeNode, cur []int, sum int)
	dfs = func(root *TreeNode, cur []int, sum int) {
		if root == nil {
			return
		}
		sum += root.Val
		cur = append(cur, root.Val)
		if root.Left == nil && root.Right == nil {
			if sum == targetSum {
				tmp := make([]int, len(cur))
				copy(tmp, cur)
				res = append(res, tmp)
			}
			return
		}

		dfs(root.Left, cur, sum)
		dfs(root.Right, cur, sum)
	}

	cur := make([]int, 0)
	dfs(root, cur, 0)

	return res
}
