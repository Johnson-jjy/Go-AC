package array

// 回旋镖的数量: 直接扫,hash存即可
func numberOfBoomerangs(points [][]int) int {
	n := len(points)

	res := 0

	for i := 0; i < n; i++ {
		store := make(map[int]int, 0)
		for j := 0; j < n; j++ {
			cur := getDistant(points[i], points[j])
			store[cur]++
		}
		for _, v := range store {
			if v >= 2 {
				res += v * (v - 1)
			}
		}
	}

	return res
}

func getDistant(a []int, b []int) int {
	dX := abs1(a[0], b[0])
	dY := abs1(a[1], b[1])

	return dX * dX + dY * dY
}

func abs1(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}