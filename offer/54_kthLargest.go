package offer

// 二叉搜索树的第k大节点
// 解一: 利用中序遍历,只找前K个大的,且穿成一个双向链表
// 注1: 逻辑上:第K大和第K个大的数,注意区别;
// 注2: 因为是第K大,从有往左;故有两处需要调整; 1)中序遍历的顺序,先有后左; 2)串成list链表的顺序,先Right后Left
var pre54 *TreeNode
//var head *TreeNode

func kthLargest1(root *TreeNode, k int) int {
	if root == nil {
		return -1
	}
	res := 0
	count := 0

	var inOrder func(root *TreeNode) bool
	inOrder = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		if inOrder(root.Right) {
			return true
		}
		count++
		if pre54 == nil {
			pre54 = root
		} else {
			pre54.Left = root
			root.Right = pre54
			pre54 = root
		}
		if count == k {
			res = pre54.Val
			return true
		}
		if inOrder(root.Left) {
			return true
		}

		return false
	}

	inOrder(root)

	return res
}


// 解二: 直接中序遍历
// 注: 此处走了一个完全的中序遍历流程,实际上只遍历一部分即可
func kthLargest2(root *TreeNode, k int) int {
	if root == nil {
		return -1
	}
	store := make([]int, 0)

	var inOrder func(root *TreeNode)
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}

		inOrder(root.Left)
		store = append(store, root.Val)
		inOrder(root.Right)
	}

	inOrder(root)

	return store[len(store) - k]
}