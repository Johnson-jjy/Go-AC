package Binary_Tree

//Definition for a Node.
type MyNode struct {
    Val int
    Left *MyNode
    Right *MyNode
    Next *MyNode
}


func connect(root *MyNode) *MyNode {
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
