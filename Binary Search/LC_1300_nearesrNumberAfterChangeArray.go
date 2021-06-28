package Binary_Search

// 当发现要找寻的答案存在确切范围后，可考虑二分答案
func findBestValue(arr []int, target int) int {
	m := len(arr)
	max := 0
	// 此处记录最大的值，返回值大于数组最大值是没有意义的
	for i := 0; i < m; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	left := 0
	right := max
	ans := target * m //可用INF，保存最小的绝对值差
	res := max //保存返回值，按题目要求尽量小

	for left <= right {
		mid := left + (right - left)/2
		count := 0
		less := 0
		sum := 0

		for _, num := range arr {
			if num < mid {
				count++
				less += num
			}
		}
		sum = less + mid * (m - count)
		cur := abs(sum, target)
		if cur < ans {
			ans = cur
			res = mid
		} else if cur == ans {
			// 需要考虑ans不变的情况下，可取的更小res值
			if mid < res {
				res = mid
			}
		}
		//fmt.Printf("sum:%d\tless:%d\tans:%d\tres:%d\tleft:%d\tright:%d\tmid:%d\n", sum, less, ans, res, left, right, mid)

		// 让sum从目前方向向target靠拢
		if sum > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return res
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}
