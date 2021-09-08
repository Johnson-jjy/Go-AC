package cig

// 检查子树
func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
	if t2 == nil {
		return true
	}
	if t1 == nil {
		return false
	}

	return ifSubTree(t1, t2) || checkSubTree(t1.Left, t2) || checkSubTree(t1.Right, t2)
}

func ifSubTree(p *TreeNode, q *TreeNode) bool {
	if q == nil {
		return true
	}
	if p == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	return ifSubTree(p.Left, q.Left) && ifSubTree(p.Right, q.Right)
}
