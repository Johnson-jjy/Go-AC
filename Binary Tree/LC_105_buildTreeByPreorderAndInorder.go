package Binary_Tree

import "Go-AC/cig"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 抓住前序和中序的特点，结合go传切片即可
func buildTree(preorder []int, inorder []int) *cig.TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// 注意,这列preorder数组中是int值,故需要新构建一个节点
	root := &cig.TreeNode{
		preorder[0],
		nil,
		nil,
	}
	if len(preorder) == 1 {
		return root
	}
	index := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root.Val { // 注意这里的比较别比错了
			index = i
			break
		}
	}
	// 注意这里不是新初始化,而是直接赋值即可
	root.Left = buildTree(preorder[1 : index + 1], inorder[: index])
	root.Right = buildTree(preorder[index + 1:], inorder[index + 1:]) // 切片的切分细节不要出岔子

	return root
}


