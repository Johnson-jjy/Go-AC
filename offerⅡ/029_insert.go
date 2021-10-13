package offer_

// 排序的循环链表

// 注1: 对初值为nil时的特殊处理
// 注2: 插入的逻辑(难点) -- 1) x恰好在中间(可在循环过程中插入); 2)x为最大值 或 x为最小值 -> 关键在于找出临界点 --> 且为最小值时还需要特别记录Pre值
// 以下0版本为把逻辑弄复杂了的情况 --> 实际上只需要记录max的位置; 无论node最大还是最小,都在max的后一个
func insert0(aNode *ListNode, x int) *ListNode {
	node := &ListNode{
		x,
		nil,
	}

	if aNode == nil {
		aNode = node
		node.Next = node
		return aNode
	}

	cur := aNode.Next
	minNode := cur
	maxNode := cur
	pre := aNode
	minNodePre := pre
	for {
		//fmt.Println(cur)
		if cur.Val <= x && cur.Next.Val >= x {
			node.Next = cur.Next
			cur.Next = node
			break
		}
		if cur.Val <= minNode.Val {
			minNode = cur
			minNodePre = pre
		}
		if cur.Val >= maxNode.Val {
			maxNode = cur
		}
		cur = cur.Next
		pre = pre.Next
		if cur == aNode.Next {
			break
		}
	}

	//fmt.Println(minNode, maxNode)
	if node.Next == nil {
		if x >= maxNode.Val {
			node.Next = maxNode.Next
			maxNode.Next = node
		} else {
			node.Next = minNodePre.Next
			minNodePre.Next = node
		}
	}

	return aNode
}

func insert(aNode *ListNode, x int) *ListNode {
	node := &ListNode{
		x,
		nil,
	}

	if aNode == nil {
		aNode = node
		node.Next = node
		return aNode
	}

	cur := aNode.Next
	maxNode := cur
	for {
		//fmt.Println(cur)
		if cur.Val <= x && cur.Next.Val >= x {
			node.Next = cur.Next
			cur.Next = node
			break
		}
		if cur.Val >= maxNode.Val {
			maxNode = cur
		}
		cur = cur.Next
		if cur == aNode.Next {
			break
		}
	}

	//fmt.Println(minNode, maxNode)
	if node.Next == nil {
		node.Next = maxNode.Next
		maxNode.Next = node
	}

	return aNode
}