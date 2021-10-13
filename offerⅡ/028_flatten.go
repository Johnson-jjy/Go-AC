package offer_

// 展开双向多级链表

// Definition for a Node.
type Node struct {
    Val int
    Prev *Node
    Next *Node
    Child *Node
}

// 注1: 修改节点指向时, 注意child指向空
// 注2: do-while形式
func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	dummy := &Node {
		-1,
		nil,
		root,
		nil,
	}
	store := make([]*Node, 0)
	cur := root
	for {
		if cur.Child != nil {
			next := cur.Next
			if next != nil {
				store = append(store, next)
			}
			cur.Next = cur.Child
			cur.Child.Prev = cur
			cur.Child = nil
		}
		if cur.Next == nil {
			break
		}
		cur = cur.Next
	}
	n := len(store) - 1
	for n >= 0 {
		next := store[n]
		cur.Next = next
		next.Prev = cur
		cur = next
		for cur.Next != nil {
			cur = cur.Next
		}
		n--
	}

	return dummy.Next
}
