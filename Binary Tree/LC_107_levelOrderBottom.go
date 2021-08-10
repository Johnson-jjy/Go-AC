package Binary_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 注:本质上还是BFS遍历,但添加顺序不同,难点在于对append的使用
// 解1: append头插入 -> 对于第一个参数,将cur升维,用匿名初始化法; 对第二个参数,将res"降维",后接"..."
// 解2: 像BFS一样正常添加,然后对于结果数组进行逆序
func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)

	if root == nil {
		return res
	}

	queue := make([]*TreeNode, 1)
	queue[0] = root

	for len(queue) != 0 {
		size := len(queue)
		cur := make([]int, size)

		for i := 0; i < size; i++ {
			node := queue[i]
			cur[i] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		queue = queue[size:]
		res = append([][]int{cur}, res...) // 重点! -> 如何向二维数组头插入一维数组
	}

	return res
}


// 注:正序添加,最后逆序,也可解
//for i := 0; i < len(levelOrder) / 2; i++ {
//	levelOrder[i], levelOrder[len(levelOrder) - 1 - i] = levelOrder[len(levelOrder) - 1 - i], levelOrder[i]
//}
