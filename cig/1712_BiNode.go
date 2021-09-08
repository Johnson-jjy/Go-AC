package cig

// BiNode
// 坑1: 每次需对head和pre设定初始值nil(否则全局变量后台值不改)
// 坑2: 每次修改当前root的left为nil,而不是pre的left; 即每次必须对当前节点root进行相应修改
// 对坑2的补充: 如果不每次都及时对root节点进行操作,则会出现环
var head *TreeNode
var pre1712 *TreeNode

func convertBiNode(root *TreeNode) *TreeNode {
	head = nil
	pre1712 = nil

	if root == nil {
		return nil
	}
	convert(root)

	return head
}

func convert(root *TreeNode) {
	if root == nil {
		return
	}
	convert(root.Left)
	if head == nil {
		head = root
		head.Left = nil
		pre1712 = root
	} else {
		pre1712.Right = root
		root.Left = nil
		pre1712 = root
	}
	convert(root.Right)
}
