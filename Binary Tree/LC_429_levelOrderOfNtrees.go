package Binary_Tree

//Definition for a Node.
type Node struct {
    Val int
    Children []*Node
}

// N叉树的层序遍历 -- 本质上和二叉树没有区别 -> bfs <- queue
func levelOrder(root *Node) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := make([]*Node, 1)
	queue[0] = root

	for len(queue) != 0 {
		size := len(queue)
		cur := make([]int, size)

		for i := 0; i < size; i++ {
			node := queue[i]
			children := node.Children
			for j := 0; j < len(children); j++ {
				queue = append(queue, children[j])
			}
			cur[i] = node.Val
		}
		queue = queue[size:]
		res = append(res, cur)
	}

	return res
}
