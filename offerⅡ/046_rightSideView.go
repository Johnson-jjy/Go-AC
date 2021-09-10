package offer_

// 二叉树的右侧视图
// 解一: BFS -- 基操

// 解二: DFS -- 前序遍历思路 + deep记层数 可以得到本层的第一个节点
// 欲最左则先left;欲最优则先right
func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	var dfs func(root *TreeNode, deep int)
	dfs = func(root *TreeNode, deep int) {
		if root == nil {
			return
		}
		if deep > len(res) {
			res = append(res, root.Val)
		}
		dfs(root.Right, deep + 1)
		dfs(root.Left, deep + 1)
	}

	dfs(root, 1)

	return res
}
