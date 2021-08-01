package two_pointers

// 双指针: 注意理解思路->左右MAX中小的那一个才能做基准值;当前值依附由来的基准值;每次更新基准值
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	res := 0
	lH := height[0]
	rH := height[len(height) - 1]
	l := 0
	r := len(height) - 1

	for l <= r {
		lH = max(lH, height[l])
		rH = max(rH, height[r])
		if lH < rH {
			res += lH - height[l]
			l++
		} else {
			res += rH - height[r]
			r--
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
