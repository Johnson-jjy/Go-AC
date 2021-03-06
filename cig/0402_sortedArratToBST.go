package cig

// 思路:每次二分找头结点;另可利用go的切片性质直接递归 -> 注意直接递归时,需处理的右数组应该是mid+1
var store []int

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	store = make([]int, len(nums))
	copy(store, nums) // 使用内置函数copy,这样每次不需要传入切片(go均为传值操作，会拷贝)
	left := 0
	right := len(nums) - 1
	mid := left + (right - left)/2
	root := &TreeNode{
		nums[mid],
		nil,
		nil,
	}
	root.Left = getNode(left, mid - 1, root)
	root.Right = getNode(mid + 1, right, root)
	return root
}

func getNode(start int, end int, root *TreeNode) *TreeNode {
	if start > end {
		return nil
	}
	if start == end {
		node := &TreeNode{
			store[start],
			nil,
			nil,
		}
		return node
	}
	mid := start + (end - start)/2
	node := &TreeNode{
		store[mid],
		nil,
		nil,
	}
	//fmt.Printf("%d\n", node.Val)
	node.Left = getNode(start, mid - 1, node)
	node.Right = getNode(mid + 1, end, node)
	return node
}

// go切片直接递归
func sortedArrayToBST2(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	//fmt.Println(nums)
	mid := len(nums)/2
	cur := &TreeNode {
		nums[mid],
		nil,
		nil,
	}
	//fmt.Println(cur.Val)
	cur.Left = sortedArrayToBST(nums[:mid])
	cur.Right = sortedArrayToBST(nums[mid + 1:])

	return cur
}