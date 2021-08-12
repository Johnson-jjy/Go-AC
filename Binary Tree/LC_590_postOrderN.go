package Binary_Tree

// 迭代版本 -> 每次都把相应的值操作完才退栈 -> 不要忘了要操作完,也不要忘了退栈!
func postorder1(root *Node) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	stack := make([]*Node, 0)

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
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
			tmp := node.Children[0]
			node.Children = node.Children[1:]
			root = tmp
		} else {
			res = append(res, node.Val)
			stack = stack[:len(stack) - 1]
			root = nil
		}
	}

	return res
}

// 迭代1.5: 欣赏
func postorder(root *Node) []int {
	var res = []int{}
	if root == nil{
		return res
	}
	var stack = []*Node{root}
	for 0 < len(stack) {
		root = stack[len(stack)-1]
		res = append(res, root.Val)
		stack = stack[:len(stack)-1]
		l := len(root.Children)
		for i := 0; i < l; i++ {
			stack = append(stack, root.Children[i]) //入栈
		}
	}

	l := len(res) - 1
	for i := 0; i < l/2+1; i++ {
		res[i], res[l-i] = res[l-i], res[i]
	}
	return res
}

// 递归版本
func postorder2(root *Node) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	var post func(root *Node)
	post = func(root *Node) {
		if root == nil {
			return
		}

		for _, node := range root.Children {
			post(node)
		}
		res = append(res, root.Val)
	}

	post(root)

	return res
}
