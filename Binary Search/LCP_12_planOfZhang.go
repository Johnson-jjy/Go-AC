package Binary_Search

func minTime(time []int, m int) int {
	if m >= len(time) {
		return 0
	}

	left := 1
	right := 1000000000

	ans := right
	for left <= right {
		mid := left + (right - left) / 2

		//fmt.Printf("mid:%d\n", mid)
		if lessM(time, mid, m) {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}

func lessM(arr[] int, cur int, m int) bool {
	count := 1
	curTime := 0
	maxTime := 0
	use := false
	for i := 0; i < len(arr); i++ {
		if maxTime < arr[i] {
			maxTime = arr[i]
		}
		if curTime + arr[i] > cur {
			if !use {
				curTime = curTime + arr[i] - maxTime
				if curTime > cur {
					return false
				}
				use = true
			} else {
				curTime = 0
				count++
				i--
				maxTime = 0
				use = false
			}
		} else {
			curTime += arr[i]
		}
	}

	//fmt.Printf("count:%d\n", count)
	return count <= m
}
