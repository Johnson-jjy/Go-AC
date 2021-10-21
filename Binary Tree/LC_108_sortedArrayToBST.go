package Binary_Tree

// 将有序数组转换为二叉搜索树
// 思路: 取数组中值为根节点, 递归即可
// 复杂度: t:O(N); s:O(logN)
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums)/2
	root := &TreeNode{
		nums[mid],
		sortedArrayToBST(nums[:mid]),
		sortedArrayToBST(nums[mid + 1:]),
	}

	return root
}
