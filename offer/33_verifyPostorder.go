package offer

// 二叉搜索树的后序遍历序列
// 解1: 递归,抓住二叉搜素书如何切分数组的关键!
func verifyPostorder(postorder []int) bool {
	n := len(postorder)
	if n < 2 {
		return true
	}
	root := postorder[n - 1]
	flag := n - 1
	for i := 0; i < n; i++ {
		if postorder[i] > root {
			if flag == n - 1 {
				flag = i
			}
		}
		if flag != n - 1 && postorder[i] < root {
			return false
		}
	}

	return verifyPostorder(postorder[:flag]) && verifyPostorder(postorder[flag:n - 1])
}

// 解二: 单调栈,待补充
