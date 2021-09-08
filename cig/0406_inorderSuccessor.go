package cig

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 后继者: 注意这种题目必须先保存head(或者说pre)的情况,然后对于每一个节点,都需要对其左右子节点进行重新连接 -> 本质类似于把二叉树结构改为链表结构
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var res *TreeNode
	var pre *TreeNode

	var inorder func(root *TreeNode) bool
	inorder = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		if inorder(root.Left) {
			return true
		}
		if pre == nil {
			pre = root
		} else {
			if pre == p {
				res = root
				return true
			}
			pre.Right = root
			root.Left = pre
			pre = root
		}
		if inorder(root.Right) {
			return true
		}
		return false
	}

	inorder(root)

	return res
}

// 本题可作为关于go指针的探讨题

// 以下代码会出问题,可拿来探究指针奥秘
//func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
//	if root == nil || p == nil {
//		return nil
//	}
//	var ans *TreeNode
//	var pre *TreeNode
//	findInorderSuccer(root, p, pre, ans)
//
//	return ans
//}
//
//func findInorderSuccer(root *TreeNode, target *TreeNode, pre *TreeNode, ans *TreeNode) {
//	if root == nil {
//		return
//	}
//	findInorderSuccer(root.Left, target, pre, ans)
//	if pre != nil {
//		fmt.Printf("%d\n", pre.Val)
//	}
//	if pre != nil && pre.Val == target.Val {
//		ans = root
//		return
//	}
//	pre = root
//	fmt.Printf("%v\n", pre)
//	findInorderSuccer(root.Right, target, pre, ans)
//}


// 即使给ans赋值了以后,实际上最终退出的时候依然没有返回值
//var pre *TreeNode
//
//func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
//	if root == nil || p == nil {
//		return nil
//	}
//	var ans *TreeNode
//
//	findInorderSuccer(root, p, ans)
//
//	return ans
//}
//
//func findInorderSuccer(root *TreeNode, target *TreeNode, ans *TreeNode) {
//	if root == nil {
//		return
//	}
//	findInorderSuccer(root.Left, target, ans)
//	fmt.Printf("1:%v\n", pre)
//	if pre == target {
//		fmt.Printf("3:%v\n", ans)
//		ans = root
//		fmt.Printf("4:%v\n", ans)
//		return
//	}
//	pre = root
//	fmt.Printf("2:%v\n", pre)
//	findInorderSuccer(root.Right, target, ans)
//}


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 思路:利用二叉搜索树的性质找p;针对p后序的两种情况:
// 1. 有右子树时则后序节点为右子树中的最小值,可非递归查询找到
// 2. 无右子树时则后序节点为他最近的左子树父亲节点
var pre *TreeNode

func inorderSuccessor1(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil || p == nil {
		return nil
	}

	pre = nil
	//var target *TreeNode
	target := findtarget(root , p)
	if target == nil {
		return nil
	}
	if target.Right != nil {
		target = target.Right
		for target.Left != nil {
			target = target.Left
		}
		return target
	}
	return pre
}

func findtarget(root *TreeNode, t *TreeNode) *TreeNode {
	if root == nil || root == t {
		return root
	}

	if t.Val > root.Val {
		return findtarget(root.Right, t)
	}

	pre = root
	return findtarget(root.Left, t)
}