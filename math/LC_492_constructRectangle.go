package math

// 构造矩形

// 简单模拟题, 注意循环终止条件
func constructRectangle(area int) []int {
	res := make([]int, 2)

	for i := 1; i * i <= area; i++ {
		if area % i == 0 {
			res[0] = area/i
			res[1] = i
		}
	}

	return res
}
