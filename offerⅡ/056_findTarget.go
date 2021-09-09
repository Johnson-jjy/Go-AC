package offer_

// 二叉搜索树中两个节点的值
// 解一: 中序遍历保存数组(有序), 再双指针找答案即可
func findTarget1(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	store := make([]int, 0)

	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		store = append(store, root.Val)
		inorder(root.Right)
	}

	inorder(root)

	i := 0
	j := len(store) - 1
	for i < j {
		cur := store[i] + store[j]
		if cur == k {
			return true
		} else if cur > k {
			j--
		} else {
			i++
		}
	}

	return false
}

// 解二: 类似与两数之和; 在遍历的过程中找答案即可
func findTarget2(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}

	store := make(map[int]bool, 0)

	var inorder func(root *TreeNode) bool
	inorder = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		if inorder(root.Left) {
			return true
		}
		need := k - root.Val
		if store[need] {
			//fmt.Println(need, root.Val)
			return true
		}
		store[root.Val] = true
		if inorder(root.Right) {
			return true
		}
		return false
	}

	if inorder(root) {
		return true
	}

	return false
}