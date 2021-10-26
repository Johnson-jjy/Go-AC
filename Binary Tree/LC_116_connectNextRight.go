package Binary_Tree

// 填充每一个节点的下一个右侧节点指针
//Definition for a Node.
type MyNode struct {
    Val int
    Left *MyNode
    Right *MyNode
    Next *MyNode
}

// 标准层序遍历
// 复杂度: t O(N); s O(N)
func connect1(root *MyNode) *MyNode {
	if root == nil {
		return root
	}
	queue := make([]*MyNode, 1)
	queue[0] = root

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			node := queue[i]
			if i == size - 1 {
				node.Next = nil
			} else {
				node.Next = queue[i + 1]
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		queue = queue[size:]
	}

	return root
}

// 遍历, 不用额外的空间
func connect2(root *MyNode) *MyNode {
	if root == nil {
		return root
	}

	// 每次循环从该层的最左侧节点开始
	for leftmost := root; leftmost.Left != nil; leftmost = leftmost.Left {
		// 通过 Next 遍历这一层节点，为下一层的节点更新 Next 指针
		for node := leftmost; node != nil; node = node.Next {
			// 左节点指向右节点
			node.Left.Next = node.Right

			// 右节点指向下一个左节点
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}
	}

	// 返回根节点
	return root
}

// dfs,
func connect(root *MyNode) *MyNode {
	dfs(root)
	return root
}

func dfs(root *MyNode) {
	if root == nil {
		return
	}
	left := root.Left
	right := root.Right

	// 以root为起点,将整个纵深这段串联起来
	for left != nil {
		left.Next = right
		left = left.Right
		right = right.Left
	}

	// 递归的调用左右节点，完成同样的纵深串联
	dfs(root.Left)
	dfs(root.Right)
}