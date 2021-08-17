package greedy

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// 监控二叉树
// 一个核心思想:叶子节点不放监控器->那么nil应该是已覆盖的状态;继而一层一层讨论
// 最后需要对根节点是否覆盖进行讨论
func minCameraCover(root *TreeNode) int {

	// 0 无覆盖
	// 1 有摄像头
	// 2 有覆盖

	res := 0
	var doSearch func(*TreeNode) int
	doSearch = func(root *TreeNode) int {
		if root == nil {
			return 2
		}
		left := doSearch(root.Left)
		right := doSearch(root.Right)

		//  只要有一个子节点是无覆盖的，就必须要再当前节点安装摄像头
		if left == 0 || right == 0 {
			res += 1
			return 1
		}
		// 然后下面的判断子节点就不会有0了

		// 只要有两个子节点都被覆盖了，那当前节点就是无覆盖的
		if left == 2 && right == 2 {
			return 0
		}
		// 其余情况说明当前子节点至少有1个是有摄像头的，可能会有1个是有覆盖的，但是肯定没有无覆盖的，所以只需要返回有覆盖就行

		return 2

	}

	if doSearch(root) == 0 {
		res += 1
	}
	return res
}
