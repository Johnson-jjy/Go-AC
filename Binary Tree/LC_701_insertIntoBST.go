package Binary_Tree

// 二叉搜索树中的插入操作
// 法一: 递归 -- 理解递归函数的意义, 空节点时可返回当前节点, 其他情况寻空插入
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		node := &TreeNode {
			val,
			nil,
			nil,
		}
		return node
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	} else {
		root.Left = insertIntoBST(root.Left, val)
	}
	return root
}

// 法二: 迭代 -- 模拟查找BST的过程, 找到待插入处插入即可.
func insertIntoBST2(root *TreeNode, val int) *TreeNode {
	if root == nil {
		node := &TreeNode {
			val,
			nil,
			nil,
		}
		return node
	}
	cur := root
	for {
		if cur.Val < val {
			if cur.Right == nil {
				cur.Right = &TreeNode {val, nil, nil,}
				break
			} else {
				cur = cur.Right
			}
		} else {
			if cur.Left == nil {
				cur.Left = &TreeNode {val, nil, nil,}
				break
			} else {
				cur = cur.Left
			}
		}
	}

	return root
}