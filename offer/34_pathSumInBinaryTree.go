package offer

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 思路: 递归即可
// 注: 看清题目,明确所指的路径是什么
// 注: 传的是复制,故不需要再做回溯
// 另: 可用闭包的方式写,继而不用定义相应的全局变量
var res [][]int

func pathSum(root *TreeNode, target int) [][]int {
	res = make([][]int, 0)
	if  root == nil {
		return res
	}
	path := make([]int, 0)
	backtrack_(root, path, target)

	return res
}

func backtrack_(root *TreeNode, path []int, total int) {
	if root == nil {
		return
	}
	total -= root.Val
	path = append(path, root.Val)
	//fmt.Printf("%v\t%v\n", root.Val, path)
	if root.Left == nil && root.Right == nil {
		if total == 0 {
			cur := make([]int, len(path))
			copy(cur, path)
			res = append(res, cur)
		}
		return
	}

	backtrack_(root.Left, path, total)
	backtrack_(root.Right, path, total)
}
