package offer


// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// 此方法较为笨拙
var pre []int
var in []int

func buildTree1(preorder []int, inorder []int) *TreeNode {
	m := len(preorder)
	if m == 0 {
		return nil
	}
	var root *TreeNode
	index := 0
	pre = make([]int, m)
	copy(pre, preorder)
	in = make([]int, m)
	copy(in, inorder)

	for i := 0; i < m; i++ {
		if preorder[0] == inorder[i] {
			index = i
			root = &TreeNode{
				preorder[0],
				nil,
				nil,
			}
			break
		}
	}

	root.Left = buildBinaryTree(0, index - 1, 1, index)
	root.Right = buildBinaryTree(index + 1, m - 1, index + 1, m - 1)
	return root
}

func buildBinaryTree(inStart, inEnd, preStart, preEnd int) *TreeNode {
	if inStart > inEnd {
		return nil
	}
	root := &TreeNode{
		pre[preStart],
		nil,
		nil,
	}
	index := inStart
	for i := inStart; i <= inEnd; i++ {
		if pre[preStart] == in[i] {
			index = i
			break
		}
	}
	flag := index - inStart
	root.Left = buildBinaryTree(inStart, index - 1, preStart + 1, preStart + flag)
	root.Right = buildBinaryTree(index + 1, inEnd, preStart + flag + 1, preEnd)
	return root
}


// 该方法巧用了Go的切片
func buildTree2(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree2(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree2(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}


// 迭代 + 栈 : 还需再看 --> 思考:可否引申到后序遍历?
func buildTree3(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	stack := []*TreeNode{}
	stack = append(stack, root)
	var inorderIndex int
	for i := 1; i < len(preorder); i++ {
		preorderVal := preorder[i]
		node := stack[len(stack)-1]
		if node.Val != inorder[inorderIndex] {
			node.Left = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Left)
		} else {
			for len(stack) != 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex++
			}
			node.Right = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Right)
		}
	}
	return root
}

