package offer

// 对称二叉树
// 解一: 递归
func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return check(root.Left, root.Right)
}

func checkSymmetric(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	return checkSymmetric(p.Left, q.Right) && checkSymmetric(p.Right, q.Left)
}

// 解二: BFS -> 不要拘泥于size; 每次只取要操作的节点进行判断即可
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