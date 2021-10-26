package Binary_Tree

// 合并二叉树
// DFS: 根据签名含义来皆可
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// 注意: 1. 这里可不用判&&的情况,会被包含; 2. 不是只生成一个新的node, 而是继承所有
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	node := &TreeNode{root1.Val + root2.Val, nil, nil,}
	node.Left = mergeTrees(root1.Left, root2.Left)
	node.Right = mergeTrees(root1.Right, root2.Right)

	return node
}

// BFS: 3个队列保存相关属性