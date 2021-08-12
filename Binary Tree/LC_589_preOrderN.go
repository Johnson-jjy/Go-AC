package Binary_Tree

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

// 解法一: 非递归 -> 核心思想: 弹出栈时对该节点的所有操作已结束(遍历各子节点与加入res数组)
func preorder1(root *Node) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	stack := make([]*Node, 0)

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			res = append(res, root.Val)
			if len(root.Children) > 0 {
				tmp := root.Children[0]
				root.Children = root.Children[1:]
				root = tmp
			} else {
				root = nil
			}
		}
		node := stack[len(stack) - 1]
		if len(node.Children) > 0 {
			root = node.Children[0]
			node.Children = node.Children[1:]
		} else {
			stack = stack[:len(stack) - 1]
		}
	}

	return res
}

// 迭代1.5 -> 更清晰, 欣赏
func preorder(root *Node) []int {
	var res []int
	var stack = []*Node{root}
	for 0 < len(stack) {
		for root != nil {
			res = append(res, root.Val) //前序输出
			if 0 == len(root.Children) {
				break
			}
			for i := len(root.Children) - 1; 0 < i; i-- { // 注意此处顺序: 从右往左而不是从左往右
				stack = append(stack, root.Children[i]) //入栈
			}
			root = root.Children[0]
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return res
}

// 解法二: 递归: 类似二叉树
func preorder2(root *Node) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	var pre func(root *Node)
	pre = func(root *Node) {
		if root == nil {
			return
		}
		res = append(res, root.Val)

		for _, node := range root.Children {
			pre(node)
		}
	}

	pre(root)

	return res
}
