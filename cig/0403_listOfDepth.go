package cig

// 特定深度节点链表 -- 本质层序遍历
// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

//Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

// BFS + 链表，注意不同的节点需要进行不同的声明
func listOfDepth(tree *TreeNode) []*ListNode {
	if tree == nil {
		return nil
	}
	res := make([]*ListNode, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, tree)

	for len(queue) != 0 {
		size := len(queue)
		cur := &ListNode{
			-1,
			nil,
		}
		dummy := cur
		for i := 0; i < size; i++ {
			node := queue[0]
			curNode := &ListNode{
				node.Val,
				nil,
			}
			queue = queue[1:len(queue)]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			cur.Next = curNode
			cur = curNode
		}
		res = append(res, dummy.Next)
	}

	return res
}
