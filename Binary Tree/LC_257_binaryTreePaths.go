package Binary_Tree

import "strconv"

// 二叉树的所有路径
// 解一: DFS
// 难点: 哪些时候path（string）的状态需要回溯？ -> 理解,go里的切片,虽依然传值;但底层指向同一地址;会修改值;但对结构的修改(如添加,删除等)则不会被保留
func binaryTreePaths1(root *TreeNode) []string {
	res := make([]string, 0)
	if root == nil {
		return res
	}


	var dfs func(root *TreeNode, cur string)
	dfs = func(root *TreeNode, cur string) {
		if cur == "" {
			cur = strconv.Itoa(root.Val)
		} else {
			cur = cur + "->"
			cur = cur + strconv.Itoa(root.Val)
		}
		if root.Left == nil && root.Right == nil {
			tmp := cur
			res = append(res, tmp)
			return
		}
		if root.Left != nil {
			dfs(root.Left, cur)
		}
		if root.Right != nil {
			dfs(root.Right, cur)
		}
	}

	cur := ""
	dfs(root, cur)

	return res
}

// 以下解法还需要自己再实现一次! 包括对112和113两道题,再实现一下!
// 解二: BFS -> 非层序遍历时的BFS,关键点都在于每次仅对当前节点进行处理即可;
func binaryTreePaths2(root *TreeNode) []string {
	paths := []string{}
	if root == nil {
		return paths
	}
	// 维护两个队列
	nodeList := []*TreeNode{}
	pathList := []string{}
	nodeList = append(nodeList, root)
	pathList = append(pathList, strconv.Itoa(root.Val))

	for i:=0; i<len(nodeList);i++ {
		node, path := nodeList[i],pathList[i]
		if node.Left == nil && node.Right == nil {
			paths = append(paths, path)
			continue
		}
		if node.Left != nil {
			nodeList = append(nodeList, node.Left)
			pathList = append(pathList, path + "->" + strconv.Itoa(node.Left.Val))
		}
		if node.Right != nil {
			nodeList = append(nodeList, node.Right)
			pathList = append(pathList, path + "->" + strconv.Itoa(node.Right.Val))
		}
	}
	return paths
}
