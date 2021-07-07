package list

//Definition for a Node.
type Node struct {
    Val int
    Next *Node
    Random *Node
}


func copyRandomList1(head *Node) *Node {
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil {
		copy := &Node{
			cur.Val,
			cur.Next,
			nil,
		}
		cur.Next = copy
		cur = copy.Next
	}

	cur = head
	for cur != nil {
		copy := cur.Next
		if cur.Random != nil {
			copy.Random = cur.Random.Next
		}
		cur = copy.Next
	}

	cur = head
	newHead := head.Next
	for cur != nil {
		copy := cur.Next
		cur.Next = copy.Next
		cur = copy.Next
		if cur != nil {
			copy.Next = cur.Next
		}
	}

	return newHead
}

func copyRandomList2(head *Node) *Node {
	if head == nil {
		return nil
	}
	store := make(map[*Node]*Node, 0)
	cur := head

	for cur != nil {
		store[cur] = &Node{cur.Val, nil, nil,}
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		node := store[cur]
		next := store[cur.Next]
		node.Next = next
		random := store[cur.Random]
		node.Random = random
		cur = cur.Next
	}

	return store[head]
}