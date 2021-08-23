package main

import (
	"fmt"
	"sort"
)

func main() {
	//input := bufio.NewScanner(os.Stdin)
	//arr := make([][]int, 0)
	//for input.Scan() {
	//	nums := strings.Split(input.Text(), " ")
	//	cur := make([]int, 2)
	//	for i := 0; i < len(nums); i++ {
	//		num, _ := strconv.Atoi(nums[i])
	//		cur[i] = num
	//	}
	//	arr = append(arr, cur)
	//}
	var n, L int
	fmt.Scan(&n, &L)
	arr := make([][2]int, n) // 列确定就直接写
	//fmt.Println(L)

	for i := 0; i < n; i++ {
		//cur := make([]int, 2)
		fmt.Scan(&arr[i][0])
		fmt.Scan(&arr[i][1])
	}

	//fmt.Println(arr)
	// 左端点从小到大, 右端点从大到小
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0] || (arr[i][0] == arr[j][0] && arr[i][1] > arr[j][1])
	})

	// 这个判断条件无法判断最后
	//if arr[0][0] > 0 || arr[len(arr) - 1][1] < L {
	//	fmt.Println(-1)
	//}
	//if arr[0][0] != 0 {
	//	fmt.Println(-1)
	//	return
	//}

	res := 0
	rightMost := 0
	i := 0

	// 和跳跃游戏Ⅱ的区别在于,有可能无法到达
	for i < n {
		res++
		curMost := rightMost
		//j := i
		if arr[i][0] > curMost {
			fmt.Println(-1)
			return
		}
		//fmt.Println(res, "?", curMost)
		// 勿忘此处还需要对j作判断
		for i < n && arr[i][0] <= curMost { // 注意这是二维矩阵而非传统的一维
			rightMost = max(rightMost, arr[i][1]) // 注意此处也和一维不同 -> 注意不是arr[j][0] + arr[j][1];仅仅只是arr[j][1]
			if rightMost >= L { // 看清题意:L才是右端限制
				//fmt.Println(arr[j])
				fmt.Println(res)
				return
			}
			//j++
			i++
		}
		//fmt.Println(rightMost)
		// 注意: 此处不能照搬, i非距离,i为坐标下标
		//i = j // 需理解: 刚进循环时,i为本次其实坐标,curMost为最远可达;故更新时应将i改为curMost+1作下一可达位置
	}

	fmt.Println(-1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
