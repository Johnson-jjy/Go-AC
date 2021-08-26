package main

import "fmt"

// 自动佳适出租车订单分配
// N个点构成环,可顺可逆,N的作用应该是取余数
// 想办法处理K个乘客(也就是K行数据)

/*
	思路:按题目给的方式把输入的K个数组排序(基于:起始时间--上车ID--下车ID) -> 这里其实应该是可以直接用库函数
	然后从第二行开始,判断此时有多少组,最后取最大
	复杂度O^2 -> 10^8 -> 感觉可能有些数据过不了
*/

// 注: 输入输出每个语言不同,这里不处理

func main() {
	N := 50
	K := 5
	array := make([][]int, 5)
	array[0] = []int{0, 0, 15}
	array[1] = []int{10, 10, 11}
	array[2] = []int{15, 20, 40}
	array[3] = []int{20, 30, 40}
	array[4] = []int{30, 40, 10}
	// 以上为题目数据->array是个二维数组

	//排序的逻辑不会,拍成上面这样的的话(Nlogn)下面这个应该没问题

	result := 0
	for i := 1; i < K; i++ {
		thisTerm := 1
		for j := i - 1; j >= 0; j-- {
			// 从前一站点一直遍历到最早出发站点
			cur := (array[j][1] + (array[i][0] - array[j][0])/5) % N // j站点当前位置 -> 防环
			if cur < array[j][2] { // j还在运行
				thisTerm++
			}
		}
		// 每一次更新
		if result < thisTerm {
			result = thisTerm
		}
	}

	fmt.Printf("result: %d\n", result)
}
