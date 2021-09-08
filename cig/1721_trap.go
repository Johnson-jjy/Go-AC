package cig

// 经典接雨水 -> 双指针 -> 理解其本质为由动归而来的变式
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	l := 0
	r := len(height) - 1
	lH := height[l]
	rH := height[r]

	res := 0
	for l < r {
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
