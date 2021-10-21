package Binary_Tree

// 二叉搜索树中的搜索
// 解一: 递归
func searchBST1(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < val {
		return searchBST1(root.Right, val)
	}

	if root.Val > val {
		return searchBST1(root.Left, val)
	}

	return root
}

// 解二: 迭代
func searchBST2(root *TreeNode, val int) *TreeNode {
	for root != nil && root.Val != val {
		if root.Val < val {
			root = root.Right
		} else {
			root = root.Left
		}
	}

	return root
}

