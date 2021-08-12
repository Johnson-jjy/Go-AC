package Binary_Tree

// 解法一: 迭代版本 -> 注意退栈时机和node的更新 -> 退栈时必须完成"加入res数组" 和 "遍历左子树" "遍历右子树" 共三个操作
func postorderTraversal1(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	stack := make([]*TreeNode, 0)
	var pre *TreeNode

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack) - 1]
		if node.Right == nil || node.Right == pre {
			res = append(res, node.Val)
			stack = stack[:len(stack) - 1]
			pre = node
			root = nil
		} else {
			root = node.Right
		}
	}

	return res
}

// 解法二: 递归版本
func postorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	var postOrder func(root *TreeNode)
	postOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		postOrder(root.Left)
		postOrder(root.Right)
		res = append(res, root.Val)
	}

	postOrder(root)

	return res
}
