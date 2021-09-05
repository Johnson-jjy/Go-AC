package other

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func GetMinTimeCost( root *TreeNode,  failId int ,  timeCost []int ) int {
	// write code here
	if failId >= len(timeCost) || root == nil {
		return 0
	}
	storeT := make([]int, 0)
	queue := make([]*TreeNode, 1)
	queue[0] = root
	tree := make([]*TreeNode, len(timeCost))
	level := make([][]int, 0)

	for len(queue) > 0 {
		size := len(queue)
		cur := timeCost[queue[0].Val]
		curLever := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue[i]
			tree[node.Val] = node
			curLever[i] = node.Val
			if node.Left == nil && node.Right == nil {
				continue
			}
			if timeCost[node.Val] < cur {
				cur = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		storeT = append(storeT, cur)
		queue = queue[size:]
		level = append(level, curLever)
	}
	res := timeCost[failId]

	start := getStart(failId, level)
	h := getHeight(tree[failId])
	for i := start + 1; i < h; i++ {
		res += storeT[i]
	}

	return res
}

func getStart(target int, arr [][]int) int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] == target {
				return i
			}
		}
	}
	return -1
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := getHeight(root.Left)
	right := getHeight(root.Right)

	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}