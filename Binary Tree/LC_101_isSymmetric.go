package Binary_Tree

// 对称二叉树
func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return check(root.Left, root.Right)
}

func check(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return check(left.Left, right.Right) && check(left.Right, right.Left)
}

// 解二: BFS -> 对每次要操作的node进行判断和添加
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := []*TreeNode{root, root}

	for len(queue) > 0 {
		p := queue[0]
		q := queue[1]
		queue = queue[2:]

		if q == nil && p == nil {
			continue
		}
		if q == nil || p == nil {
			return false
		}
		if q.Val == p.Val {
			queue = append(queue, q.Left, p.Right)
			queue = append(queue, q.Right, p.Left)
		} else {
			return false
		}
	}

	return true
}